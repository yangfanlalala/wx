package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatedraftlist.html

const ApiGetTemplateDraftList = "https://api.weixin.qq.com/wxa/gettemplatedraftlist"

func (client *WeChatClient) GetTemplateDraftList(data *GetTemplateDraftListRequest) (*GetTemplateDraftListResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetTemplateDraftList).
		WithMethod(http.MethodGet).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetTemplateDraftListResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetTemplateDraftListResponse, nil
}

func (client *WeChatClient) BuildGetTemplateDraftListRequest() *GetTemplateDraftListRequest {
	return &GetTemplateDraftListRequest{}
}

type GetTemplateDraftListRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type GetTemplateDraftListResponse struct {
	DraftList []struct {
		CreateTime  int64  `json:"create_time"`
		UserVersion string `json:"user_version"`
		UserDesc    string `json:"user_desc"`
		DraftID     string `json:"draft_id"`
	} `json:"draft_list"`
}
