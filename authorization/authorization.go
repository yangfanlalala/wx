package authorization

const (
	ApiStartPushTicket = "https://api.weixin.qq.com/cgi-bin/component/api_start_push_ticket" // 开启推送
	ApiComponentToken = "https://api.weixin.qq.com/cgi-bin/component/api_component_token" // 第三方平台access_token
	ApiCreatePreAuthCode = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode" // 预授权码
	ApiQueryAuth = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth" // 使用授权码获取授权信息
	ApiAuthorizerToken = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token" // 获取/刷新接口调用令牌
)
