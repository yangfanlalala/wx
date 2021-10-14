package wx

// AuthorizerInfo 授权信息
type AuthorizerInfo struct {
	Nickname        string `json:"nick_name"`
	HeadImage       string `json:"head_img"`
	ServiceTypeInfo struct {
		ID int64 `json:"id"`
	} `json:"service_type_info"`
	VerifyTypeInfo struct {
		ID int64 `json:"id"`
	} `json:"verify_type_info"`
	Username        string `json:"user_name"`
	PrincipalName   string `json:"principal_name"` //主体名称
	Alias           string `json:"alias"`          //公众号所设置的微信号
	Signature       string `json:"signature"`      //(小程序)账号介绍
	BusinessInfo    string `json:"business_info"`
	QRCodeURL       string `json:"qrcode_url"`
	MiniProgramInfo struct {
		Network    WxaNetwork    `json:"network"`
		Categories []WxaCategory `json:"categories"`
	} `json:"MiniProgramInfo"`
}

// AuthorizationInfo 授权信息
type AuthorizationInfo struct {
	AuthorizerAppID        string `json:"authorizer_appid"`
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int64  `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
	FuncInfo               []struct {
		FuncScopeCategory struct {
			ID int64
		} `json:"funcscope_category"`
	} `json:"func_info"`
}

// BusinessInfo 开通信息
type BusinessInfo struct {
	OpenStore int64 `json:"open_store"` //是否开通微信门店功能
	OpenScan  int64 `json:"open_scan"`  //是否开通微信扫商品功能
	OpenPay   int64 `json:"open_pay"`   //是否开通微信支付功能
	OpenCard  int64 `json:"open_card"`  //是否开通微信卡券功能
	OpenShake int64 `json:"open_shake"` //是否开通微信摇一摇功能
}

// VerifyInfo 微信认证信息
type VerifyInfo struct {
	QualificationVerify   bool  `json:"qualification_verify"` //是否资质认证，若是，拥有微信认证相关的权限。
	NamingVerify          bool  `json:"naming_verify"`
	AnnualReview          bool  `json:"annual_review"` //是否需要年审
	AnnualReviewBeginTime int64 `json:"annual_review_begin_time"`
	AnnualReviewEndTime   int64 `json:"annual_review_end_time"`
}

// SignatureInfo 功能介绍信息
type SignatureInfo struct {
	Signature       string `json:"signature"`
	ModifyUsedCount int64  `json:"modify_used_count"`
	ModifyQuota     int64  `json:"modify_quota"` //功能介绍修改次数总额度（本月）
}

// HeadImageInfo 头像信息
type HeadImageInfo struct {
	HeadImageURL    string `json:"head_image_url"`
	ModifyUsedCount int64  `json:"modify_used_count"`
	ModifyQuota     int64  `json:"modify_quota"`
}

// NicknameInfo 名称信息
type NicknameInfo struct {
	Nickname        string `json:"nickname"`
	ModifyUsedCount int64  `json:"modify_used_count"`
	ModifyQuota     int64  `json:"modify_quota"`
}

type WxaNetwork struct {
	RequestDomain   []string `json:"RequestDomain"`
	WsRequestDomain []string `json:"WsRequestDomain"`
	UploadDomain    []string `json:"UploadDomain"`
	DownloadDomain  []string `json:"DownloadDomain"`
	BizDomain       []string `json:"BizDomain"`
	UDPDomain       []string `json:"UDPDomain"`
}

type Category struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Level         int64   `json:"level"`
	Father        int64   `json:"father"`
	Children      []int64 `json:"children"`
	SensitiveType int64   `json:"sensitive_type"`
	Qualify       struct {
		ExterList []struct {
			InnerList []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"inner_list"`
		} `json:"exter_list"`
		Remark string `json:"remark"]`
	} `json:"qualify"`
}

type WxaCategory struct {
}

type AppVersion struct {
	AppVersion  int64  `json:"app_version"`
	UserVersion string `json:"user_version"`
	UserDesc    string `json:"user_desc"`
	CommitTime  int64  `json:"commit_time"`
}

type Template struct {
	CreateTime   int64  `json:"create_time"`
	UserVersion  string `json:"user_version"`
	UserDesc     string `json:"user_desc"`
	TemplateID   int64  `json:"template_id"`
	TemplateType int64  `json:"template_type"`
}

type TemplateDraft struct {
	CreateTime  int64  `json:"create_time"`
	UserVersion string `json:"user_version"`
	UserDesc    string `json:"user_desc"`
	DraftID     int64 `json:"draft_id"`
}

type UserPhoneInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode string `json:"countryCode"`
	Watermark struct{
		Appid string `json:"appid"`
	} `json:"watermark"`
}
