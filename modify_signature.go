package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/modifysignature.html

const ApiModifySignature = "https://api.weixin.qq.com/cgi-bin/account/modifysignature"

func (client *WeChatClient) ModifySignature() {

}

type ModifySignatureRequest struct {
	AccessToken string `position:"query" name:"access_token"`
	Signature string `position:"body" name:"signature"`
}
