package business

const (
	ApiTmplGetCategory = "https://api.weixin.qq.com/wxaapi/newtmpl/getcategory" // 获取小程序账号的类目
	ApiTmplGetPubTemplateTitles = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles" // 获取模板标题列表
	ApiTmplGetPubTemplateKeyWords = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords" // 获取模板标题下的关键词库
	ApiTmplAddTemplate = "https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate" // 组合模板并添加到个人模板库
	ApiTmplGetTemplate = "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate" // 获取帐号下的模板列表
	ApiTmplDeleteTemplate = "https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate" // 删除帐号下的某个模板
	ApiTmplSend = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send" // 发送订阅消息
)