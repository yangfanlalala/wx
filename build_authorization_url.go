package wx

import "fmt"

const PCAuthorizationURLTpl = "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%s"
const H5AuthorizationURLTpl = "https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%s#wechat_redirect"
const AuthTypeMP = "1"
const AuthTypeWxa = "2"
const AuthTypeMpAndWxa = "3"

func (client *WeChatClient) BuildAuthorizationURL(preCode, redirectURL, authType string) *AuthorizationURL {
	return &AuthorizationURL{
		PC: fmt.Sprintf(PCAuthorizationURLTpl, client.options.AppID, preCode, redirectURL, authType),
		H5: fmt.Sprintf(H5AuthorizationURLTpl, client.options.AppID, preCode, redirectURL, authType),
	}
}

type AuthorizationURL struct {
	PC string `json:"pc"`
	H5 string `json:"h5"`
}
