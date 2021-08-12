package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/undocodeaudit.html

const ApiUndoCodeAudit = "https://api.weixin.qq.com/wxa/undocodeaudit"

func (client *WeChatClient) UndoCodeAudit() {

}

type UndoCodeAuditRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}
