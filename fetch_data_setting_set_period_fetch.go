package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/set_period_fetchdatasetting.html

func (client *WeChatClient) SetDataSettingSetPeriodFetch(data *SetDataSettingSetPeriodFetchRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiFetchDataSetting).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := struct {
		CommonResponse
		SetDataSettingSetPeriodFetchResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return err
	}
	if err := rsp.Error(); err != nil {
		return err
	}
	return nil
}

func (client *WeChatClient) BuildSetDataSettingSetPeriodFetchRequest() *SetDataSettingSetPeriodFetchRequest {
	return &SetDataSettingSetPeriodFetchRequest{
		Action: "set_period_fetch",
	}
}

type SetDataSettingSetPeriodFetchRequest struct {
	AccessToken        string `position:"query" name:"access_token" json:"-"`
	Action             string `position:"body" name:"action" json:"action"`
	IsPeriodFetchOpen  bool   `position:"body" name:"is_period_fetch_open" json:"is_period_fetch_open"`
	PeriodFetchType    int64  `position:"body" name:"period_fetch_type" json:"period_fetch_type"`
	PeriodFetchURL     string `position:"body" name:"period_fetch_url" json:"period_fetch_url"`
	PeriodEnv          string `position:"body" name:"period_env" json:"period_env"`
	PeriodFunctionName string `position:"body" name:"period_function_name" json:"period_function_name"`
}

type SetDataSettingSetPeriodFetchResponse struct {
}
