package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/weapp-wxverify/secwxaapi_wxaauth.html

const ApiSecQueryWxaAuth = "https://api.weixin.qq.com/wxa/sec/queryauth"

var SecApplyStatus = map[int32]string{
	0: "审核单不存在",
	1: "待支付",
	2: "审核中",
	3: "打回重填",
	4: "认证通过",
	5: "认证最终失败（不能再修改）",
}
var SecTaskStatus = map[int32]string{
	0:  "初始状态",
	1:  "任务超时, 24小时内有效",
	2:  "用户授权拒绝",
	3:  "用户授权同意 ",
	4:  "发起人脸流程",
	5:  "人脸认证失败",
	6:  "人脸认证ok",
	7:  "人脸认证后，已经提交手机号码下发验证码",
	8:  "手机验证失败",
	9:  "手机验证成功",
	11: "创建认证审核单失败",
	12: "创建认证审核审核单成功",
	14: "验证失败",
	15: "等待支付",
}

func (client *WeChatClient) SecQueryWxaAuth(data *SecQueryWxaAuthRequest) (*SecQueryWxaAuthResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiSecQueryWxaAuth).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		SecQueryWxaAuthResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.SecQueryWxaAuthResponse, nil
}

func (client *WeChatClient) BuildSecQueryWxaAuthRequest() *SecQueryWxaAuthRequest {
	return &SecQueryWxaAuthRequest{}
}

type SecQueryWxaAuthRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	TaskID      string `position:"body" name:"taskid" json:"taskid"`
}

type SecQueryWxaAuthResponse struct {
	ApplyStatus  int32  `json:"apply_status"`  // 当前任务流程的状态码
	TaskStatus   int32  `json:"task_status"`   //  小程序管理员授权链接
	AuthURL      string `json:"auth_url"`      // 审核单状态，创建认证审核审核单成功后该值有效
	OrderID      int64  `json:"order_id"`      // 小程序后台展示的认证订单号
	APPID        string `json:"appid"`         // 小程序appid
	RefillReason string `json:"refill_reason"` // 当审核单被打回重填(apply_status=3)时有效
	FailReason   string `json:"fail_reason"`   // 审核最终失败的原因(apply_status=5)时有效
}
