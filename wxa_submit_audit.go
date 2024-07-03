package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/submit_audit.html

const ApiWxaSubmitAudit = "https://api.weixin.qq.com/wxa/submit_audit"

func (client *WeChatClient) WxaSubmitAudit(data *WxaSubmitAuditRequest) (*WxaSubmitAuditResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiWxaSubmitAudit).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		WxaSubmitAuditResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.WxaSubmitAuditResponse, nil
}

func (client *WeChatClient) BuildWxaSubmitAuditRequest() *WxaSubmitAuditRequest {
	return &WxaSubmitAuditRequest{}
}

type WxaSubmitAuditRequest struct {
	AccessToken   string                 `position:"query" name:"access_token" json:"-"`
	ItemList      []*WxaSubmitAuditItem  `position:"body" name:"item_list" json:"item_list,omitempty"`
	PreviewInfo   *WxaSubmitAuditPreview `position:"body" name:"preview_info" json:"preview_info,omitempty"`
	VersionDesc   string                 `position:"body" name:"version_desc" json:"version_desc,omitempty"`
	FeedbackInfo  string                 `position:"body" name:"feedback_info" json:"feedback_info,omitempty"`
	FeedbackStuff string                 `position:"body" name:"feedback_stuff" json:"feedback_stuff"`
	UgcDeclare    *WxaSubmitAuditUGC     `position:"body" name:"ugc_declare" json:"ugc_declare,omitempty"`
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
	Title       string `name:"title" json:"title"`
}

type WxaSubmitAuditPreview struct {
	VideoIDList   []string `name:"video_id_list" json:"video_id_list"`
	PictureIDList []string `name:"pic_id_list" json:"picture_id_list"`
}

type WxaSubmitAuditUGC struct {
	Scene          []int64 `name:"scene" json:"scene"`
	OtherSceneDesc string  `name:"other_scene_desc" json:"other_scene_desc"`
	Method         []int64 `name:"method" json:"method"`
	HasAuditTeam   int64   `name:"has_audit_team" json:"has_audit_team"`
	AuditDesc      string  `name:"audit_desc" json:"audit_desc"`
}

type WxaSubmitAuditResponse struct {
	AuditID int64 `json:"auditid"`
}

type AuditUGCKV struct {
	Key   string
	Value int64
}

var (
	AuditUGCScene = []*AuditUGCKV{
		{
			Key:   "不涉及用户生成内容",
			Value: 0,
		},
		{
			Key:   "用户资料",
			Value: 1,
		},
		{
			Key:   "图片",
			Value: 2,
		},
		{
			Key:   "视频",
			Value: 3,
		},
		{
			Key:   "文本",
			Value: 4,
		},
		{
			Key:   "其他",
			Value: 5,
		},
	}
	AuditUGCMethod = []*AuditUGCKV{
		{
			Key:   "使用平台建议的内容安全API",
			Value: 1,
		},
		{
			Key:   "使用其他的内容审核产品",
			Value: 2,
		},
		{
			Key:   "通过人工审核把关",
			Value: 3,
		},
		{
			Key:   "未做内容审核把关",
			Value: 4,
		},
	}
)
