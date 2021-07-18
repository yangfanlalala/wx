package authorization

const (
	ApiStartPushTicket = "https://api.weixin.qq.com/cgi-bin/component/api_start_push_ticket"
	// 第三方平台access_token
	ApiComponentToken = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	// 预授权码
	ApiCreatePreAuthCode = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=COMPONENT_ACCESS_TOKEN"
	// 使用授权码获取授权信息
	ApiQueryAuth = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=COMPONENT_ACCESS_TOKEN"
	// 获取/刷新接口调用令牌
	ApiAuthorizerToken = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=COMPONENT_ACCESS_TOKEN"
)
