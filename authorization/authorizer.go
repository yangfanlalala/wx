package authorization

const (
	ApiGetAuthorizerList = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_list" // 拉取所有已授权的帐号信息
	ApiGetAuthorizerOption = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_option" // 获取授权方选项信息
	ApiSetAuthorizerOption = "https://api.weixin.qq.com/cgi-bin/component/api_set_authorizer_option" // 设置授权方选项信息
	ApiGetAuthorizerInfo = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info" // 获取授权方的帐号基本信息
)
