package model

import (
	"strings"
	"sync"

	"gorm.io/gorm"
)

// ProcInstHistory ProcInstHistory
type ProcInstHistory struct {
	gorm.Model
	// 流程定义ID
	ProcDefID int `json:"procDefId"`
	// 流程定义名
	ProcDefName string `json:"procDefName"`
	// title 标题
	Title string `json:"title"`
	// 用户部门
	Department string `json:"department"`
	Company    string `json:"company"`
	// 当前节点
	NodeID string `json:"nodeID"`
	// 审批人
	Candidate string `json:"candidate"`
	// 当前任务
	TaskID        int    `json:"taskID"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
	Duration      int64  `json:"duration"`
	StartUserID   string `json:"startUserId"`
	StartUserName string `json:"startUserName"`
	IsFinished    bool   `gorm:"default:false" json:"isFinished"`
}

func (t *ProcInstHistory) TableName() string {
	return "proc_inst_history"
}

// StartHistoryByMyself 查询我发起的流程
func StartHistoryByMyself(userID, company string, pageIndex, pageSize int) ([]*ProcInstHistory, int, error) {
	maps := map[string]interface{}{
		"start_user_id": userID,
		"company":       company,
	}
	return findProcInstsHistory(maps, pageIndex, pageSize)
}
func findProcInstsHistory(maps map[string]interface{}, pageIndex, pageSize int) ([]*ProcInstHistory, int, error) {
	var datas []*ProcInstHistory
	var count int64
	selectDatas := func(in chan<- error, wg *sync.WaitGroup) {
		go func() {
			err := GetDB().Where(maps).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Order("start_time desc").Find(&datas).Error
			in <- err
			wg.Done()
		}()
	}
	selectCount := func(in chan<- error, wg *sync.WaitGroup) {
		err := GetDB().Model(&ProcInstHistory{}).Where(maps).Count(&count).Error
		in <- err
		wg.Done()
	}
	var err1 error
	var wg sync.WaitGroup
	numberOfRoutine := 2
	wg.Add(numberOfRoutine)
	errStream := make(chan error, numberOfRoutine)
	// defer fmt.Println("close channel")
	selectDatas(errStream, &wg)
	selectCount(errStream, &wg)
	wg.Wait()
	defer close(errStream) // 关闭通道
	for i := 0; i < numberOfRoutine; i++ {
		// log.Printf("send: %v", <-errStream)
		if err := <-errStream; err != nil {
			err1 = err
		}
	}
	// fmt.Println("结束")
	return datas, int(count), err1
}

// FindProcHistory 查询历史纪录
func FindProcHistory(userID, company string, pageIndex, pageSize int) ([]*ProcInstHistory, int, error) {
	var datas []*ProcInstHistory
	var count int64
	var err1 error
	var wg sync.WaitGroup
	numberOfRoutine := 2
	errStream := make(chan error, numberOfRoutine)
	selectDatas := func(wg *sync.WaitGroup) {
		go func() {
			err := GetDB().Where("id in (select distinct proc_inst_id from identitylink_history where company=? and user_id=?)", company, userID).
				Offset((pageIndex - 1) * pageSize).Limit(pageSize).
				Order("start_time desc").Find(&datas).Error
			errStream <- err
			wg.Done()
		}()
	}
	selectCount := func(wg *sync.WaitGroup) {
		go func() {
			err := GetDB().Model(&ProcInstHistory{}).
				Where("id in (select distinct proc_inst_id from identitylink_history where company=? and user_id=?)", company, userID).
				Count(&count).Error
			errStream <- err
			wg.Done()
		}()
	}
	wg.Add(numberOfRoutine)
	selectDatas(&wg)
	selectCount(&wg)
	wg.Wait()
	close(errStream)

	for i := 0; i < numberOfRoutine; i++ {
		if err := <-errStream; err != nil {
			err1 = err
		}
	}
	return datas, int(count), err1
}

// SaveProcInstHistory SaveProcInstHistory
func SaveProcInstHistory(p *ProcInst) error {
	return GetDB().Table("proc_inst_history").Create(p).Error
}

// DelProcInstHistoryByID DelProcInstHistoryByID
func DelProcInstHistoryByID(id int) error {
	return GetDB().Where("id=?", id).Delete(&ProcInstHistory{}).Error
}

// SaveProcInstHistoryTx SaveProcInstHistoryTx
func SaveProcInstHistoryTx(p *ProcInst, tx *gorm.DB) error {
	return tx.Table("proc_inst_history").Create(p).Error
}

// FindProcHistoryNotify 查询抄送我的历史纪录
func FindProcHistoryNotify(userID, company string, groups []string, pageIndex, pageSize int) ([]*ProcInstHistory, int, error) {
	var datas []*ProcInstHistory
	var count int64
	var sql string
	if len(groups) != 0 {
		var s []string
		for _, val := range groups {
			s = append(s, "\""+val+"\"")
		}
		sql = "select proc_inst_id from identitylink_history i where i.type='notifier' and i.company='" + company + "' and (i.user_id='" + userID + "' or i.group in (" + strings.Join(s, ",") + "))"
	} else {
		sql = "select proc_inst_id from identitylink_history i where i.type='notifier' and i.company='" + company + "' and i.user_id='" + userID + "'"
	}
	err := GetDB().Where("id in (" + sql + ")").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Order("start_time desc").Find(&datas).Error
	if err != nil {
		return datas, int(count), err
	}
	err = GetDB().Model(&ProcInstHistory{}).Where("id in (" + sql + ")").Count(&count).Error
	if err != nil {
		return nil, int(count), err
	}
	return datas, int(count), err
}
