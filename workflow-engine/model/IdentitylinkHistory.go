package model

import (
	"gorm.io/gorm"
)

// IdentitylinkHistory IdentitylinkHistory
type IdentitylinkHistory struct {
	gorm.Model
	Group           string          `json:"group,omitempty"`
	Type            string          `json:"type,omitempty"`
	UserID          string          `json:"userid,omitempty"`
	UserName        string          `json:"username,omitempty"`
	TaskID          int             `json:"taskID,omitempty"`
	Step            int             `json:"step"`
	ProcInstID      int             `json:"procInstID,omitempty"`
	ProcInstHistory ProcInstHistory `json:"procInstHistory" gorm:"references:ID;foreignKey:ProcInstID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	Company         string          `json:"company,omitempty"`
	Comment         string          `json:"comment,omitempty"`
}

func (t *IdentitylinkHistory) TableName() string {
	return "identitylink_history"
}

// CopyIdentitylinkToHistoryByProcInstID CopyIdentitylinkToHistoryByProcInstID
func CopyIdentitylinkToHistoryByProcInstID(procInstID int, tx *gorm.DB) error {
	return tx.Exec("insert into identitylink_history select * from identitylink where proc_inst_id=?", procInstID).Error
}

// FindParticipantHistoryByProcInstID FindParticipantHistoryByProcInstID
func FindParticipantHistoryByProcInstID(procInstID int) ([]*IdentitylinkHistory, error) {
	var datas []*IdentitylinkHistory
	err := GetDB().Select("id,user_id,step,comment").Where("proc_inst_id=? and type=?", procInstID, IdentityTypes[PARTICIPANT]).Order("id asc").Find(&datas).Error
	return datas, err
}
