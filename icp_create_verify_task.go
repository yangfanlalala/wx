package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/createIcpVerifyTask.html

const ApiICPCreateVerifyTask = "https://api.weixin.qq.com/wxa/icp/create_icp_verifytask"

func (client *WeChatClient) ICPCreateVerifyTask(data *ICPCreateVerifyTaskRequest) (*ICPCreateVerifyTaskResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiICPCreateVerifyTask).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ICPCreateVerifyTaskResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ICPCreateVerifyTaskResponse, nil
}

func (client *WeChatClient) BuildICPCreateVerifyTaskRequest() *ICPCreateVerifyTaskRequest {
	return &ICPCreateVerifyTaskRequest{}
}

type ICPCreateVerifyTaskRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type ICPCreateVerifyTaskResponse struct {
	TaskID string `json:"task_id"`
}
