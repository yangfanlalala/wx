package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/queryIcpVerifyTask.html

const ApiICPQueryVerifyTask = "https://api.weixin.qq.com/wxa/icp/query_icp_verifytask"

var ICPFaceStatus = map[int32]string{
	0: "未开始", 1: "等待中", 2: "失败", 3: "成功",
}

func (client *WeChatClient) ICPQueryVerifyTask(data *ICPQueryVerifyTaskRequest) (*ICPQueryVerifyTaskResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiICPQueryVerifyTask).
		WithMethod(http.MethodGet).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		ICPQueryVerifyTaskResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.ICPQueryVerifyTaskResponse, nil
}

func (client *WeChatClient) BuildICPQueryVerifyTaskRequest() *ICPQueryVerifyTaskRequest {
	return &ICPQueryVerifyTaskRequest{}
}

type ICPQueryVerifyTaskRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	TaskID      string `position:"body" name:"task_id" json:"task_id"`
}

type ICPQueryVerifyTaskResponse struct {
	IsFinished bool  `json:"is_finished"`
	FaceStatus int32 `json:"face_status"`
}
