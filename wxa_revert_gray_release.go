package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/revertgrayrelease.html

const ApiRevertGrayRelease = "https://api.weixin.qq.com/wxa/revertgrayrelease"

func (client *WeChatClient) RevertGrayRelease() {

}

type RevertGrayReleaseRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}
