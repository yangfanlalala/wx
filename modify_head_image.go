package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/modifyheadimage.html

const ApiModifyHeadImage = "https://api.weixin.qq.com/cgi-bin/account/modifyheadimage"

func (client *WeChatClient) ModifyHeadImage() {

}

type ModifyHeadImageRequest struct {
	AccessToken string `position:"query" name:"access_token"`
	HeadImageMediaID string `position:"body" name:"head_img_media_id"`
	X1 string `position:"body" name:"x1"`
	Y1 string `position:"body" name:"y1"`
	X2 string `position:"body" name:"x2"`
	Y2 string `position:"body" name:"y2"`
}