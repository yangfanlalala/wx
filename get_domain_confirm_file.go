package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/thirdparty-management/domain-mgnt/getThirdpartyJumpDomainConfirmFile.html

const ApiGetDomainCofirmFile = "https://api.weixin.qq.com/cgi-bin/component/get_domain_confirmfile"

func (client *WeChatClient) GetDomainConfirmFile(data *GetDomainConfirmFileRequest) (*GetDomainConfirmFileResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiGetDomainCofirmFile).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		GetDomainConfirmFileResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.GetDomainConfirmFileResponse, nil
}

func (client *WeChatClient) BuildGetDomainConfirmFileRequest() *GetDomainConfirmFileRequest {
	return &GetDomainConfirmFileRequest{}
}

type GetDomainConfirmFileRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type GetDomainConfirmFileResponse struct {
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}
