package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/submit_audit.html

const WxaSubmitAudit = "https://api.weixin.qq.com/wxa/submit_audit"

func (client *WeChatClient) WxaSubmitAudit() {

}

type WxaSubmitAuditRequest struct {
	AccessToken   string                  `name:"access_token" json:"-"`
	ItemList      []WxaSubmitAuditItem    `name:"item_list" json:"item_list"`
	PreviewInfo   []WxaSubmitAuditPreview `name:"preview_info" json:"preview_info"`
	VersionDesc   string                  `name:"version_desc" json:"version_desc"`
	FeedbackInfo  string                  `name:"feedback_info"`
	FeedbackStuff string                  `name:"feedback_stuff" json:"feedback_stuff"`
	UgcDeclare    WxaSubmitAuditUGC       `name:"ugc_declare" json:"ugc_declare"`
}

type WxaSubmitAuditItem struct {
	Address     string `name:"address" json:"address"`
	Tag         string `name:"tag" json:"tag"`
	FirstClass  string `name:"first_class" json:"first_class"`
	SecondClass string `name:"second_class" json:"second_class"`
	ThirdClass  string `name:"third_class" json:"third_class"`
	FirstID     int64  `name:"first_id" json:"first_id"`
	SecondID    int64  `name:"second_id" json:"second_id"`
	ThirdID     int64  `name:"third_id" json:"third_id"`
	Title       int64  `name:"title" json:"title"`
}

type WxaSubmitAuditPreview struct {
	VideoIDList   []string `name:"video_id_list" json:"video_id_list"`
	PictureIDList []string `name:"pic_id_list" json:"picture_id_list"`
}

type WxaSubmitAuditUGC struct {
	Scene          int64  `name:"scene" json:"scene"`
	OtherSceneDesc string `name:"other_scene_desc" json:"other_scene_desc"`
	Method         int64  `name:"method" json:"method"`
	HasAuditTeam   int64  `name:"has_audit_team" json:"has_audit_team"`
	AuditDesc      string `name:"audit_desc" json:"audit_desc"`
}

type WxaSubmitAuditResponse struct {
CommonResponse
	AuditID      int64  `json:"auditid"`
}
