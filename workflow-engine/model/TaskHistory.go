package model

import (
	"gorm.io/gorm"
)

// TaskHistory TaskHistory
type TaskHistory struct {
	gorm.Model
	// Company 任务创建人对应的公司
	// Company string `json:"company"`
	// ExecutionID     string `json:"executionID"`
	// 当前执行流所在的节点
	NodeID string `json:"nodeId"`
	Step   int    `json:"step"`
	// 流程实例id
	ProcInstID int    `json:"procInstID"`
	Assignee   string `json:"assignee"`
	CreateTime string `json:"createTime"`
	ClaimTime  string `json:"claimTime"`
	// 还未审批的用户数，等于0代表会签已经全部审批结束，默认值为1
	MemberCount   int8 `json:"memberCount" gorm:"default:1"`
	UnCompleteNum int8 `json:"unCompleteNum" gorm:"default:1"`
	//审批通过数
	AgreeNum int8 `json:"agreeNum"`
	// and 为会签，or为或签，默认为or
	ActType    string `json:"actType" gorm:"default:'or'"`
	IsFinished bool   `gorm:"default:false" json:"isFinished"`
}

func (t *TaskHistory) TableName() string {
	return "task_history"
}

// CopyTaskToHistoryByProInstID CopyTaskToHistoryByProInstID
// 根据procInstID查询结果，并将结果复制到task_history表
func CopyTaskToHistoryByProInstID(procInstID int, tx *gorm.DB) error {
	return tx.Exec("insert into task_history select * from task where proc_inst_id=?", procInstID).Error
}
