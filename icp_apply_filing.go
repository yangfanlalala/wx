package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/applyIcpFiling.html

const ApiICPApplyFiling = "https://api.weixin.qq.com/wxa/icp/apply_icp_filing"

func (client *WeChatClient) ICPApplyFiling(data *ICPApplyFilingRequest) error {
	req := &CommonRequest{}
	req.WithURL(ApiICPApplyFiling).
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

func (client *WeChatClient) BuildICPApplyFilingRequest() *ICPApplyFilingRequest {
	return &ICPApplyFilingRequest{}
}

type ICPSubject struct {
	BaseInfo        *ICPSubjectBaseInfo        `json:"base_info"`                   // 主体基本信息
	PersonalInfo    *ICPSubjectPersonalInfo    `json:"personal_info,omitempty"`     // 个人主体额外信息 （非必填）
	OrganizeInfo    *ICPSubjectOrganizeInfo    `json:"organiaze_info"`              // 主体额外信息（个人备案时，如果存在与主体负责人信息相同的字段，则填入相同的值）
	PrincipalInfo   *ICPSubjectPrincipalInfo   `json:"principal_info"`              // 主体负责人信息
	LegalPersonInfo *ICPSubjectLegalPersonInfo `json:"legal_person_info,omitempty"` // 法人信息（非个人备案，且主体负责人不是法人时，必填）
}

type ICPSubjectBaseInfo struct {
	Type         int32  `json:"type"`          // 主体性质，示例值：5
	Name         string `json:"name"`          // 主办单位名称，示例值："张三"
	Province     string `json:"province"`      // 备案省份，使用省份代码，示例值："110000"(参考：获取区域信息接口)
	City         string `json:"city"`          // 备案城市，使用城市代码，示例值："110100"(参考：获取区域信息接口)
	District     string `json:"district"`      // 备案县区，使用县区代码，示例值："110105"(参考：获取区域信息接口)
	Address      string `json:"address"`       // 通讯地址，必须属于备案省市区，地址开头的省市区不用填入，例如：通信地址为“北京市朝阳区高碑店路181号1栋12345室”时，只需要填写 "高碑店路181号1栋12345室" 即可
	Comment      string `json:"comment"`       // 主体信息备注，根据需要，如实填写
	RecordNumber string `json:"record_number"` // 主体备案号，示例值：粤B2-20090059（申请小程序备案时不用填写，查询已备案详情时会返回）
}

type ICPSubjectPersonalInfo struct {
	ResidencePermit string `json:"residence_permit"` // 临时居住证明照片 media_id，个人备案且非本省人员，需要提供居住证、暂住证、社保证明、房产证等临时居住证明
}

type ICPSubjectOrganizeInfo struct {
	CertificateType    int32  `json:"certificate_type"`    // 主体证件类型，示例值：2(参考：获取证件类型接口)
	CertificateNumber  string `json:"certificate_number"`  //主体证件号码，示例值："110105199001011234"
	CertificateAddress string `json:"certificate_address"` //主体证件住所，示例值："北京市朝阳区高碑店路181号1栋12345室"
	CertificatePhoto   string `json:"certificate_photo"`   //	主体证件照片 media_id，如果小程序主体为非个人类型，则必填，
}

type ICPSubjectPrincipalInfo struct {
	Name                         string `json:"name"`                             // 负责人姓名
	Mobile                       string `json:"mobile"`                           // 负责人联系方式
	Email                        string `json:"email"`                            // 负责人电子邮件
	EmergencyContact             string `json:"emergency_contact"`                // 负责人应急联系方式
	CertificateType              int32  `json:"certificate_type"`                 // 负责人证件类型，示例值：2(参考：获取证件类型接口，此处只能填入单位性质属于个人的证件类型)
	CertificateNumber            string `json:"certificate_number"`               // 负责人证件号码
	CertificateValidityDateStart string `json:"	certificate_validity_date_start"` // 负责人证件有效期起始日期，格式为 YYYYmmdd，示例值："20230815"
	CertificateValidityDate_end  string `json:"certificate_validity_date_end"`    // 负责人证件有效期终止日期，格式为 YYYYmmdd，如证件长期有效，请填写 "长期"，示例值："20330815"
	CertificatePhotoFront        string `json:"certificate_photo_front"`          // 负责人证件正面照片 media_id（身份证为人像面）
	CertificatePhotoBack         string `json:"certificate_photo_back"`           // 负责人证件背面照片 media_id（身份证为国徽面）
	AuthorizationLetter          string `json:"	authorization_letter"`            // 授权书 media_id，当主体负责人不是法人时需要主体负责人授权书，当小程序负责人不是法人时需要小程序负责人授权书
	VerifyTaskId                 string `json:"verify_task_id"`                   // 	扫脸认证任务id(扫脸认证接口返回的task_id)，仅小程序负责人需要扫脸，主体负责人无需扫脸，
}

type ICPSubjectLegalPersonInfo struct {
	Name              string `json:"name"`               // 法人代表姓名
	CertificateNumber string `json:"certificate_number"` // 法人证件号码
}

type ICPApplets struct {
	BaseInfo      *ICPAppletsBaseInfo      `json:"base_info"`
	PrincipalInfo *ICPAppletsPrincipalInfo `json:"principal_info"`
}

type ICPAppletsNRLXDetail struct {
	Type  int32  `json:"type"`  // 前置审批类型，示例值：2(参考：获取前置审批项接口)
	Code  string `json:"code"`  // 前置审批号，如果前置审批类型不是“以上都不涉及”，则必填，示例值："粤-12345号"
	Media string `json:"media"` // 前置审批媒体材料 media_id，如果前置审批类型不是“以上都不涉及”，则必填
}

type ICPAppletsBaseInfo struct {
	APPID               string                  `json:"appid"`                 // 小程序ID，不用填写，后台自动拉取
	Name                string                  `json:"name"`                  // 小程序名称，不用填写，后台自动拉取
	ServiceContentTypes []int32                 `json:"service_content_types"` // 小程序服务内容类型，只能填写二级服务内容类型，最多5个，示例值：[3, 4](参考：获取小程序服务类型接口)
	NRLXDetail          []*ICPAppletsNRLXDetail `json:"nrlx_details"`          // 前置审批项，列表中不能存在重复的前置审批类型id，如不涉及前置审批项，也需要填“以上都不涉及”
	Comment             string                  `json:"comment"`               // 请具体描述小程序实际经营内容、主要服务内容，该信息为主管部门审核重要依据，备注内容字数限制20-200字，请认真填写。（特殊备注要求请查看注意事项）
	RecodeNumber        string                  `json:"record_number"`         // 小程序备案号，示例值：粤B2-20090059-1626X（申请小程序备案时不用填写，查询已备案详情时会返回）
}

type ICPAppletsPrincipalInfo struct {
	Name                         string `json:"name"`                            // 负责人姓名，示例值："张三"
	Mobile                       string `json:"mobile"`                          // 负责人联系方式，示例值："13012344321"
	Email                        string `json:"email"`                           // 负责人电子邮件，示例值："zhangsan@zhangsancorp.com"
	EmergencyContact             string `json:"emergency_contact"`               // 负责人应急联系方式，示例值："17743211234"
	CertificateType              int32  `json:"certificate_type"`                // 负责人证件类型，示例值：2(参考：获取证件类型接口，此处只能填入单位性质属于个人的证件类型)
	CertificateNumber            string `json:"certificate_number"`              // 负责人证件号码，示例值："110105199001011234"
	CertificateValidityDateStart string `json:"certificate_validity_date_start"` // 负责人证件有效期起始日期，格式为 YYYYmmdd，示例值："20230815"
	CertificateValidityDateEnd   string `json:"certificate_validity_date_end"`   // 负责人证件有效期终止日期，格式为 YYYYmmdd，如证件长期有效，请填写 "长期"，示例值："20330815"
	CertificatePhotoFront        string `json:"certificate_photo_front"`         // 负责人证件正面照片 media_id（身份证为人像面）
	CertificatePhotoBack         string `json:"certificate_photo_back"`          // 负责人证件背面照片 media_id（身份证为国徽面）
	AuthorizationLetter          string `json:"authorization_letter"`            // 授权书 media_id，当主体负责人不是法人时需要主体负责人授权书，当小程序负责人不是法人时需要小程序负责人授权书，
	VerifyTaskID                 string `json:"verify_task_id"`                  // 扫脸认证任务id(扫脸认证接口返回的task_id)，仅小程序负责人需要扫脸，主体负责人无需扫脸
}

type ICPMaterials struct {
	CommitmentLetter                 []string `json:"commitment_letter"`                  // 互联网信息服务承诺书 media_id，最多上传1个
	BusinessNameChangeLetter         []string `json:"business_name_change_letter"`        // 主体更名函 media_id(非个人类型，且发生过更名时需要上传)，最多上传1个
	PartyBuildingConfirmationLetter  []string `json:"party_building_confirmation_letter"` // 党建确认函 media_id，最多上传1个
	PromiseVideo                     []string `json:"promise_video"`                      // 承诺视频 media_id，最多上传1个
	AuthenticityResponsibilityLetter []string `json:"authenticity_responsibility_letter"` // 网站备案信息真实性责任告知书 media_id，最多上传1个
	AUthenticityCommitmentLetter     []string `json:"authenticity_commitment_letter"`     // 小程序备案信息真实性承诺书 media_id，最多上传1个
	WebsiteConstructionProposal      []string `json:"website_construction_proposal"`      // 小程序建设方案书 media_id，最多上传1个
	SubjectOtherMaterials            []string `json:"subject_other_materials"`            // 主体其它附件 media_id，最多上传10个
	AppletsOtherMaterials            []string `json:"applets_other_materials"`            // 小程序其它附件 media_id，最多上传10个
	HoldingCertificatePhoto          []string `json:"holding_certificate_photo"`          // 手持证件照 media_id，最多上传1个
}

type ICPApplyFilingRequest struct {
	AccessToken  string        `position:"query" name:"access_token" json:"-"`
	ICPSubject   *ICPSubject   `position:"body" name:"icp_subject" json:"icp_subject"`
	ICPApplets   *ICPApplets   `position:"body" name:"icp_applets" json:"icp_applets"`
	ICPMaterials *ICPMaterials `position:"body" name:"icp_materials" json:"icp_materials"`
}

type ICPApplyFilingResponse struct {
}
