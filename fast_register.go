package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/fast_registration_of_mini_program.html

const ApiFastRegister = "https://api.weixin.qq.com/cgi-bin/account/fastregister?access_token=TOKEN" //复用公众号主体注册小程序

func (client WeChatClient) FastRegister() {

}

type FastRegisterRequest struct {
	AccessToken string `position:"query" name:"access_token"`
	Ticket      string `position:"ticket" name:"ticket"`
}
