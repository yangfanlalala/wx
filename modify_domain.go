package wx

// api document https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/Server_Address_Configuration.html

const ApiModifyDomain = "https://api.weixin.qq.com/wxa/modify_domain"
const ModifyDomainActionAdd = "add"
const ModifyDomainActionDelete = "delete"
const ModifyDomainActionSet = "set"
const ModifyDomainActionGet = "get"

func (client *WeChatClient) ModifyDomain() {

}

type ModifyDomainRequest struct {
	AccessToken     string   `position:"query" name:"access_token" json:"-"`
	Action          string   `position:"body" name:"action" json:"action"`
	RequestDomain   []string `position:"body" name:"requestdomain" json:"requestdomain"`
	WsRequestDomain []string `position:"body" name:"wsrequestdomain" json:"wsrequestdomain"`
	UploadDomain    []string `position:"body" name:"uploaddomain" json:"uploaddomain"`
	DownloadDomain  []string `position:"body" name:"downloaddomain" json:"downloaddomain"`
}

type ModifyDomainResponse struct {
	CommonResponse
	RequestDomain   []string `position:"body" name:"requestdomain" json:"requestdomain"`
	WsRequestDomain []string `position:"body" name:"wsrequestdomain" json:"wsrequestdomain"`
	UploadDomain    []string `position:"body" name:"uploaddomain" json:"uploaddomain"`
	DownloadDomain  []string `position:"body" name:"downloaddomain" json:"downloaddomain"`
}
