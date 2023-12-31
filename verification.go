package wx

const VerificationApplySuccessStatus int32 = 4

var VerificationTaskStatuses = map[int32]string{
	0:  "初始状态",
	1:  "任务超时",
	2:  "用户授权拒绝",
	3:  "用户授权同意",
	4:  "发起人脸流程",
	5:  "人脸认证失败",
	6:  "人脸认证成功",
	7:  "人脸认证后，已经提交手机号码下发验证码",
	8:  "手机验证失败",
	9:  "手机验证成功",
	11: "创建认证审核单失败",
	12: "创建认证审核审核单成功",
	14: "验证失败",
	15: "等待支付",
}

var VerificationApllyStatuses = map[int32]string{
	0: "审核单不存在",
	1: "待支付",
	2: "审核中",
	3: "打回重填",
	4: "认证通过",
	5: "认证失败（无法修改）",
}
