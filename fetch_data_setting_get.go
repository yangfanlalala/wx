package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/get_fetchdatasetting.html

func (client *WeChatClient) FetchDataSettingGet() {

}

type FetchDataSettingGetRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
	Action      string `position:"body" name:"action" json:"action"`
}

type FetchDataSettingGetResponse struct {
	ErrorCode          int64  `json:"errcode"`
	ErrorMessage       string `json:"errmsg"`
	IsPreFetchOpen     bool   `json:"is_pre_fetch_open"`
	PreFetchType       int64  `json:"pre_fetch_type"`
	PreFetchURL        string `json:"pre_fetch_url"`
	PreEnv             string `json:"pre_env"`
	PreFunctionName    string `json:"pre_function_name"`
	IsPeriodFetchOpen  bool   `json:"is_period_fetch_open"`
	PeriodFetchType    int64  `json:"period_fetch_type"`
	PeriodFetchURL     string `json:"period_fetch_url"`
	PeriodEnv          string `json:"period_env"`
	PeriodFunctionName string `json:"period_function_name"`
}
