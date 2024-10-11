package model

import (
	"time"

	"github.com/mumushuiding/util"
	"gorm.io/gorm"
)

// Execution 流程实例（执行流）表
// ProcInstID 流程实例ID
// BusinessKey 启动业务时指定的业务主键
// ProcDefID 流程定义数据的ID
type Execution struct {
	gorm.Model
	Rev         int      `json:"rev"`
	ProcInstID  uint     `json:"procInstID"`
	ProcInst    ProcInst `json:"procInst" gorm:"references:ID;foreignKey:ProcInstID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	ProcDefID   uint     `json:"procDefID"`
	ProcDefName string   `json:"procDefName"`
	// NodeInfos 执行流经过的所有节点
	NodeInfos string `gorm:"size:4000" json:"nodeInfos"`
	IsActive  int8   `json:"isActive"`
	StartTime string `json:"startTime"`
}

func (t *Execution) TableName() string {
	return "execution"
}

// Save save
func (p *Execution) Save() (ID int, err error) {
	err = GetDB().Create(p).Error
	if err != nil {
		return 0, err
	}
	return int(p.ID), nil
}

// SaveTx SaveTx
// 接收外部事务
func (p *Execution) SaveTx(tx *gorm.DB) (ID int, err error) {
	p.StartTime = util.FormatDate(time.Now(), util.YYYY_MM_DD_HH_MM_SS)
	if err := tx.Create(p).Error; err != nil {
		return 0, err
	}
	return int(p.ID), nil
}

// GetExecByProcInst GetExecByProcInst
// 根据流程实例id查询执行流
func GetExecByProcInst(procInstID int) (*Execution, error) {
	var p = &Execution{}
	err := GetDB().Where("proc_inst_id=?", procInstID).Find(p).Error
	// log.Printf("procdef:%v,err:%v", p, err)
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil || p == nil {
		return nil, err
	}
	return p, nil
}

// GetExecNodeInfosByProcInstID GetExecNodeInfosByProcInstID
// 根据流程实例procInstID查询执行流经过的所有节点信息
func GetExecNodeInfosByProcInstID(procInstID int) (string, error) {
	var e = &Execution{}
	err := GetDB().Select("node_infos").Where("proc_inst_id=?", procInstID).Find(e).Error
	// fmt.Println(e)
	if err != nil {
		return "", err
	}
	return e.NodeInfos, nil
}

// ExistsExecByProcInst ExistsExecByProcInst
// 指定流程实例的执行流是否已经存在
func ExistsExecByProcInst(procInst int) (bool, error) {
	e, err := GetExecByProcInst(procInst)
	// var p = &Execution{}
	// err := db.Where("proc_inst_id=?", procInst).Find(p).RecordNotFound
	// log.Printf("errnotfound:%v", err)
	if e != nil {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, nil
}

type User struct {
	gorm.Model
	Name      string
	CompanyID string
	Company   Company `gorm:"references:Code"` // use Code as references
}

func (t *User) TableName() string {
	return "user"
}

type Company struct {
	ID   int
	Code string
	Name string
}

func (t *Company) TableName() string {
	return "company"
}
