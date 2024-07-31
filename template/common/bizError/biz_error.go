package bizError

import "fmt"

type BizError struct {
	Code    int    `json:"code"`
	Message string `json:"messageHandlers"`
}

func (e *BizError) Error() string {
	return fmt.Sprintf("[%v]%v", e.Code, e.Message)
}

func NewBizError(code int, message string) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
	}
}

var (
	UnknowError = NewBizError(-1, "未知错误！")

	CommonInvalidError = NewBizError(-100, "参数异常！")
	CommonUpdateError  = NewBizError(-101, "更新异常！")
	CommonDeleteError  = NewBizError(-102, "删除异常！")
	CommonQueryError   = NewBizError(-103, "查询异常！")

	LoginPassCodeNoneError = NewBizError(100001, "登陆异常！")
	UserContextNoneError   = NewBizError(100002, "用户信息不存在！")
	UserAuthError          = NewBizError(100003, "当前用户无此权限！")

	TaskOrderStatusAlreadyQualified = NewBizError(200001, "当前任务已经完成！")
	TaskOrderStatusNotQualified     = NewBizError(200002, "当前任务未达标！")
	TaskOrderStatusAlreadyReceived  = NewBizError(200101, "当前任务积分已经领取！")

	SigninedAlreadyError = NewBizError(110000, "已经签到过啦！")
)
