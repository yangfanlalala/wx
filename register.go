package wx

const FastRegisterNotificationSuccessStatus = 0

var FastRegisterStatuses = map[int32]string{
	0:      "生成任务",
	1:      "任务超时",
	2:      "任务已经被用户拒绝",
	3:      "用户同意创建",
	4:      "已经发起人脸流程",
	5:      "人脸认证失败",
	6:      "人脸认证ok",
	7:      "人脸认证后，已经提交手机号码下发验证码",
	8:      "手机验证失败",
	9:      "手机验证成功",
	10:     "手机号验证成功后，提交了创建",
	11:     "创建失败",
	12:     "创建成功",
	13:     "验证成功",
	14:     "验证失败",
	100001: "已下发的模板消息法人并未确认且已超时（24h），未进行身份证校验",
	100002: "已下发的模板消息法人并未确认且已超时（24h），未进行人脸识别校验",
	100003: "已下发的模板消息法人并未确认且已超时（24h）",
	101:    "工商数据返回：“企业已注销”",
	102:    "工商数据返回：“企业不存在或企业信息未更新”",
	103:    "工商数据返回：“企业法定代表人姓名不一致”",
	104:    "工商数据返回：“企业法定代表人身份证号码不一致”",
	105:    "法定代表人身份证号码，工商数据未更新，请 5-15 个工作日之后尝试",
	1000:   "工商数据返回：“企业信息或法定代表人信息不一致”",
	1001:   "主体创建小程序数量达到上限",
	1002:   "主体违规命中黑名单",
	1003:   "管理员绑定账号数量达到上限",
	1004:   "管理员违规命中黑名单",
	1005:   "管理员手机绑定账号数量达到上限",
	1006:   "管理员手机号违规命中黑名单",
	1008:   "管理员身份证违规命中黑名单",
	-1:     "企业与法人姓名不一致",
}
