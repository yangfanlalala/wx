package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/set_pre_fetchdatasetting.html

const FetchDataSettingActionGet = "get"
const FetchDataSettingActionSetPreFetch = "set_pre_fetch"
const FetchDataSettingActionSetPeriodFetch = "set_period_fetch"
const ApiFetchDataSetting = "https://api.weixin.qq.com/wxa/fetchdatasetting"

func (client WeChatClient) FetchDataSettingSetPreFetch() {

}

type FetchDataSettingSetPreFetchRequest struct {
	AccessToken string `position:"query" name:"access_token"`
	Action string `position:"body" name:"action"`
	IsPreFetchOpen bool `position:"body" name:"is_pre_fetch_open"`
	PreFetchType int64 `position:"body" name:"pre_fetch_type"`
	PreFetchURL string `position:"body" name:"pre_fetch_url"`
	PreEnv string `position:"body" name:"pre_env"`
	PreFunctionName string `position:"body" name:"pre_function_name"`
}