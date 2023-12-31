package wx

const ICPSuccessStatus int32 = 0

const (
	ErrCodeICPOK                      int32 = 0
	ErrCodeICPSymtemException         int32 = -1
	ErrCodeICPReject                  int32 = 86202
	ErrCodeICPServiceException        int32 = 86301
	ErrCodeICPImageTextMismatch       int32 = 86302
	ErrCodeICPImageCertifiateValidity int32 = 86308
	ErrCodeICPApprovaling             int32 = 86324
	ErrCodeICPUncompletingOrder       int32 = 86358
)

var ICPStatuses = map[int32]string{
	2:    "平台审核中",
	3:    "平台审核驳回",
	4:    "管局审核中",
	5:    "管局审核驳回",
	6:    "已备案",
	1024: "未备案",
	1025: "未备案 && 小程序信息未填",
	1026: "未备案 && 小程序类目未填",
	1027: "未备案 && 小程序信息未填 && 小程序类目未填",
	1028: "未备案 && 小程序未认证",
	1029: "未备案 && 小程序信息未填 && 小程序未认证",
	1030: "未备案 && 小程序类目未填 && 小程序未认证",
	1031: "未备案 && 小程序信息未填 && 小程序类目未填 && 小程序未认证",
}
