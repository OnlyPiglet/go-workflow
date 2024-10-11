package model

import (
	"gorm.io/gorm"
)

// ExecutionHistory ExecutionHistory
// 执行流历史纪录
type ExecutionHistory struct {
	gorm.Model
	Rev             int             `json:"rev"`
	ProcInstID      int             `json:"procInstID" `
	ProcInstHistory ProcInstHistory `json:"procInstHistory" gorm:"references:ID;foreignKey:ProcInstID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	ProcDefID       uint            `json:"procDefID"`
	ProcDefName     string          `json:"procDefName"`
	// NodeInfos 执行流经过的所有节点
	NodeInfos string `gorm:"size:4000" json:"nodeInfos"`
	IsActive  int8   `json:"isActive"`
	StartTime string `json:"startTime"`
}

func (t *ExecutionHistory) TableName() string {
	return "execution_history"
}

// CopyExecutionToHistoryByProcInstIDTx CopyExecutionToHistoryByProcInstIDTx
func CopyExecutionToHistoryByProcInstIDTx(procInstID int, tx *gorm.DB) error {
	return tx.Exec("insert into execution_history select * from execution where proc_inst_id=?", procInstID).Error
}
