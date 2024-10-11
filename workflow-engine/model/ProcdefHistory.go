package model

// ProcdefHistory 历史流程定义
type ProcdefHistory struct {
	Procdef
}

func (t *ProcdefHistory) TableName() string {
	return "procdef_history"
}

// Save Save
func (p *ProcdefHistory) Save() (ID int, err error) {
	err = GetDB().Create(p).Error
	if err != nil {
		return 0, err
	}
	return int(p.ID), nil
}
