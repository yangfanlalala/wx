package wx

import "net/http"

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/set_pre_fetchdatasetting.html

const FetchDataSettingActionGet = "get"
const FetchDataSettingActionSetPreFetch = "set_pre_fetch"
const FetchDataSettingActionSetPeriodFetch = "set_period_fetch"
const ApiFetchDataSetting = "https://api.weixin.qq.com/wxa/fetchdatasetting"

func (client *WeChatClient) FetchDataSettingSetPreFetch(data FetchDataSettingSetPreFetchRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiFetchDataSetting).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &CommonResponse{}
	if err := client.DoRequest(req, rsp); err != nil {
		return err
	}
	if err := rsp.Error(); err != nil {
		return err
	}
	return nil
}

type FetchDataSettingSetPreFetchRequest struct {
	AccessToken     string `position:"query" name:"access_token" json:"-"`
	Action          string `position:"body" name:"action" json:"action"`
	IsPreFetchOpen  bool   `position:"body" name:"is_pre_fetch_open" json:"is_pre_fetch_open"`
	PreFetchType    int64  `position:"body" name:"pre_fetch_type" json:"pre_fetch_type"`
	PreFetchURL     string `position:"body" name:"pre_fetch_url" json:"pre_fetch_url"`
	PreEnv          string `position:"body" name:"pre_env" json:"pre_env"`
	PreFunctionName string `position:"body" name:"pre_function_name" json:"pre_function_name"`
}

type FetchDataSettingSetPreFetchResponse struct {
}
