package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/weapp-wxverify/secwxaapi_wxaauth.html

const ApiSecWxaAuth = "https://api.weixin.qq.com/wxa/sec/wxaauth"

type SecCustomerType int32
type SecInvoiceType int32

const (
	SecCustomerTypeEnterprise SecCustomerType = 1  //企业
	SecCustomerTypeIndividual SecCustomerType = 12 // 个体工商户
	SecCustomerTypePerson     SecCustomerType = 15 // 个人
	SecInvoiceTypeNone        SecInvoiceType  = 1  // 无需开
	SecInvoiceTypeElectronic  SecInvoiceType  = 2  // 电子发票
	SecInvoiceTypeVat         SecInvoiceType  = 3  // 增值税发票
)

func (client *WeChatClient) SecWxaAuth(data *SecWxaAuthRequest) (*SecWxaAuthResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiSecWxaAuth).
		WithMethod(http.MethodPost).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		SecWxaAuthResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.SecWxaAuthResponse, nil
}

func (client *WeChatClient) BuildSecWxaAuthRequest() *SecWxaAuthRequest {
	return &SecWxaAuthRequest{}
}

type SecWxaAuthRequest struct {
	AccessToken         string          `position:"query" name:"access_token" json:"-"`                           //
	CustomerType        SecCustomerType `position:"body" name:"customer_type" json:"customer_type"`               // 企业为1，个体工商户 为12，个人是15
	TaskID              string          `position:"body" name:"task_id" json:"task_id"`                           // 认证任务id，打回重审调用reauth时为必填
	ContactInfo         *SecContactInfo `position:"body" name:"contact_info" json:"contact_info"`                 // 联系人信息
	InvoiceInfo         *SecInvoiceInfo `position:"body" name:"invoice_info" json:"invoice_info"`                 // 发票信息，如果是服务商代缴模式，不需要改参数
	Qualification       string          `position:"body" name:"qualification" json:"qualification"`               // 非个人类型必填。主体资质材料 media_id 支持jpg,jpeg .bmp.gif .png格式，仅支持一张图片
	QualificationOther  []string        `position:"body" name:"qualifacation_other" json:"qualifacation_other"`   // 主体资质其他证明材料 media_id 支持jpg,jpeg .bmp.gif .png格式，最多上传10张图片
	AccountName         string          `position:"body" name:"account_name" json:"account_name"`                 // 小程序账号名称
	AccountNameType     int32           `position:"body" name:"account_name_type" json:"account_name_type"`       // 小程序账号名称命名类型 1：基于自选词汇命名 2：基于商标命名
	AccountSupplemental []string        `position:"body" name:"account_supplemental" json:"account_supplemental"` // 名称命中关键词-补充材料 media_id 支持jpg,jpeg .bmp.gif .png格式，支持上传多张图片
	PayType             int32           `position:"body" name:"pay_type" json:"pay_type"`                         // 支付方式 1：消耗服务商预购包 2：小程序开发者自行支付
	AuthIdentification  string          `position:"body" name:"auth_identification" json:"auth_identification"`   // 认证类型为个人类型时可以选择要认证的身份，从/wxa/sec/authidentitytree 里获取，填叶节点的name
	AuthIdentMaterial   []string        `position:"body" name:"auth_ident_material" json:"auth_ident_material"`   // 填了auth_identification则必填。身份证明材料 media_id （1）基于不同认证身份上传不同的材料；（2）认证类型=1时选填，支持上传10张图片（3）支持jpg,jpeg .bmp.gif .png格式
	ThirdPartyPhone     string          `position:"body" name:"third_party_phone" json:"third_party_phone"`       // 第三方联系电话
	ServiceAPPID        string          `position:"body" name:"service_appid" json:"service_appid"`               // 选择服务商代缴模式时必填。服务市场appid，该服务市场账号主体必须与服务商账号主体一致
}

// 联系人信息
type SecContactInfo struct {
	Name  string `json:"name"`  // 联系人姓名
	Email string `json:"email"` // 联系人电邮
}

// 电子发票
type SecElectronicInvoice struct {
	ID   string `json:"id"`   // 纳税识别号（15位、17、18或20位）
	Desc string `json:"desc"` // 发票备注（选填）
}

// 直至发票
type SecVatInvoice struct {
	EnterprisePhone   string `json:"enterprise_phone"`   // 企业电话
	ID                string `json:"id"`                 // 纳税识别号（15位、17、18或20位）
	EnterpriseAddress string `json:"enterprise_address"` // 企业注册地址
	BankName          string `json:"bank_name"`          // 企业开户银行
	BankAccount       string `json:"bank_account"`       // 企业银行账号
	MailingAddress    string `json:"mailing_address"`    // 发票邮寄地址邮编
	Address           string `json:"address"`            // 街道地址
	Name              string `json:"name"`               // 联系人
	Phone             string `json:"phone"`              // 联系电话
	Province          string `json:"province"`           // 省份
	City              string `json:"city"`               // 城市
	District          string `json:"district"`           // 县区
	Desc              string `json:"desc"`               // 发票备注（选填）
}

// 发票信息
type SecInvoiceInfo struct {
	InvoiceType  SecInvoiceType        `json:"invoice_info"`  // 发票类型 1: 不开发票 2: 电子发票 3: 增值税专票
	Electronic   *SecElectronicInvoice `json:"electronic"`    // 发票类型=2时必填 电子发票开票信息
	Vat          *SecVatInvoice        `json:"vat"`           // 发票类型=3时必填 增值税专票开票信息
	InvoiceTitle string                `json:"invoice_title"` // 发票抬头，需要和认证主体名称一样
}

type SecWxaAuthResponse struct {
	TaskID  string `json:"taskid"`   // 认证任务ID
	AuthURL string `json:"auth_url"` // 小程序管理员授权链接
}
