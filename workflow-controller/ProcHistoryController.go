package controller

import (
	"fmt"
	"net/http"

	"github.com/OnlyPiglet/go-workflow/workflow-engine/service"
	"github.com/mumushuiding/util"
)

// FindProcHistory 查询我的审批纪录
func FindProcHistory(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		util.ResponseErr(writer, "只支持POST方法")
		return
	}
	var receiver = service.GetDefaultProcessPageReceiver()
	err := util.Body2Struct(request, &receiver)
	if err != nil {
		util.ResponseErr(writer, err)
		return
	}
	if len(receiver.UserID) == 0 {
		util.Response(writer, "用户userID不能为空", false)
		return
	}
	if len(receiver.Company) == 0 {
		util.Response(writer, "字段 company 不能为空", false)
		return
	}
	result, err := service.FindProcHistory(receiver)
	if err != nil {
		util.ResponseErr(writer, err)
		return
	}
	fmt.Fprintf(writer, result)
}

// StartHistoryByMyself 查询我发起的流程
func StartHistoryByMyself(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		util.ResponseErr(writer, "只支持Post方法！！Only suppoert Post ")
		return
	}
	var receiver = service.GetDefaultProcessPageReceiver()
	err := util.Body2Struct(request, &receiver)
	if err != nil {
		util.ResponseErr(writer, err)
		return
	}
	if len(receiver.UserID) == 0 {
		util.Response(writer, "用户userID不能为空", false)
		return
	}
	if len(receiver.Company) == 0 {
		util.Response(writer, "字段 company 不能为空", false)
		return
	}
	result, err := service.StartHistoryByMyself(receiver)
	if err != nil {
		util.ResponseErr(writer, err)
		return
	}
	fmt.Fprintf(writer, result)
}

// FindProcHistoryNotify 查询抄送我的流程
func FindProcHistoryNotify(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		util.ResponseErr(writer, "只支持POST方法")
	}
	var receiver = service.GetDefaultProcessPageReceiver()
	err := util.Body2Struct(request, &receiver)
	if err != nil {
		util.ResponseErr(writer, err)
		return
	}
	if len(receiver.UserID) == 0 {
		util.Response(writer, "用户userID不能为空", false)
		return
	}
	if len(receiver.Company) == 0 {
		util.Response(writer, "字段 company 不能为空", false)
		return
	}
	result, err := service.FindProcHistoryNotify(receiver)
	if err != nil {
		util.ResponseErr(writer, err)
		return
	}
	fmt.Fprintf(writer, result)
}
