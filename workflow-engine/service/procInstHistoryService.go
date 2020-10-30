package service

import (
	"github.com/go-workflow/go-workflow/workflow-engine/model"
	"github.com/mumushuiding/util"
)

// FindProcHistory 查询我的审批
func FindProcHistory(receiver *ProcessPageReceiver) (string, error) {
	datas, count, err := findAllProcHistory(receiver)
	if err != nil {
		return "", err
	}
	return util.ToPageJSON(datas, count, receiver.PageIndex, receiver.PageSize)
}

func findAllProcHistory(receiver *ProcessPageReceiver) ([]*model.ProcInstHistory, int, error) {
	var page = util.Page{}
	page.PageRequest(receiver.PageIndex, receiver.PageSize)
	return model.FindProcHistory(receiver.UserID, receiver.Company, receiver.PageIndex, receiver.PageSize)
}

// DelProcInstHistoryByID DelProcInstHistoryByID
func DelProcInstHistoryByID(id int) error {
	return model.DelProcInstHistoryByID(id)
}

// StartHistoryByMyself 查询我发起的流程
func StartHistoryByMyself(receiver *ProcessPageReceiver) (string, error) {
	var page = util.Page{}
	page.PageRequest(receiver.PageIndex, receiver.PageSize)
	datas, count, err := model.StartHistoryByMyself(receiver.UserID, receiver.Company, receiver.PageIndex, receiver.PageSize)
	if err != nil {
		return "", err
	}
	return util.ToPageJSON(datas, count, receiver.PageIndex, receiver.PageSize)
}

// FindProcHistoryNotify 查询抄送我的流程
func FindProcHistoryNotify(receiver *ProcessPageReceiver) (string, error) {
	var page = util.Page{}
	page.PageRequest(receiver.PageIndex, receiver.PageSize)
	datas, count, err := model.FindProcHistoryNotify(receiver.UserID, receiver.Company, receiver.Groups, receiver.PageIndex, receiver.PageSize)
	if err != nil {
		return "", err
	}
	return util.ToPageJSON(datas, count, receiver.PageIndex, receiver.PageSize)
}
