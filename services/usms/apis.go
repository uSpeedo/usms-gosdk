// Code is generated by um-model, DO NOT EDIT IT.

package usms

import (
	"github.com/usms-sdk-go/um/request"
	"github.com/usms-sdk-go/um/response"
)

// USMS API Schema

// CreateUSMSSignatureRequest is request schema for CreateUSMSSignature action
type CreateUSMSSignatureRequest struct {
	request.CommonBase

	// [公共参数] 项目ID，不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 签名的资质证明文件类型，需与签名类型保持一致，说明如下：0-三证合一/企业营业执照/组织机构代码证书/社会信用代码证书；1-应用商店后台开发者管理截图；2-备案服务商的备案成功截图(含域名，网站名称，备案号)；3-公众号或小程序的管理界面截图；4-商标注册证书；5-组织机构代码证书、社会信用代码证书；
	CertificateType *int `required:"true"`

	// 短信签名申请原因
	Description *string `required:"true"`

	// 短信签名的资质证明文件，需先进行base64编码格式转换，此处填写转换后的字符串。文件大小不超过4 MB
	File *string `required:"true"`

	// 国内/国际短信。true:国际短信，false:国内短信，若不传值则默认该值为false
	International *bool `required:"false"`

	// 短信签名授权委托文件，需先进行base64编码格式转换，此处填写转换后的字符串。文件大小不超过4 MB；当您是代理并使用第三方的签名时（也即SigPurpose为1-他用），该项为必填项；
	ProxyFile *string `required:"false"`

	// 签名内容
	SigContent *string `required:"true"`

	// 签名用途，0-自用，1-他用；
	SigPurpose *int `required:"true"`

	// 签名类型，说明如下：0-公司或企业的全称或简称；1-App应用的全称或简称；2-工信部备案网站的全称或简称；3-公众号或小程序的全称或简称；4-商标名的全称或简称；5-政府/机关事业单位/其他单位的全称或简称；
	SigType *int `required:"true"`
}

// CreateUSMSSignatureResponse is response schema for CreateUSMSSignature action
type CreateUSMSSignatureResponse struct {
	response.CommonBase

	// 返回状态码描述，如果操作成功，默认返回为空
	Message string

	// 短信签名ID（短信签名申请时的工单ID）
	SigId string
}

// NewCreateUSMSSignatureRequest will create request of CreateUSMSSignature action.
func (c *USMSClient) NewCreateUSMSSignatureRequest() *CreateUSMSSignatureRequest {
	req := &CreateUSMSSignatureRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateUSMSSignature

调用接口CreateUSMSSignature申请短信签名
*/
func (c *USMSClient) CreateUSMSSignature(req *CreateUSMSSignatureRequest) (*CreateUSMSSignatureResponse, error) {
	var err error
	var res CreateUSMSSignatureResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateUSMSSignature", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CreateUSMSTemplateRequest is request schema for CreateUSMSTemplate action
type CreateUSMSTemplateRequest struct {
	request.CommonBase

	// [公共参数] 项目ID，不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 模板变量属性说明
	Instruction *string `required:"false"`

	// 标记是否为国际短信。true:国际短信，false:国内短信，若不传值则默认该值为false
	International *bool `required:"false"`

	// 短信模板用途类型：1-验证码类短信模板；2-系统通知类短信模板；3-会员推广类短信模板；
	Purpose *int `required:"true"`

	// 短信模板申请原因说明，字数不超过128，每个中文、符号、英文、数字等都计为1个字。
	Remark *string `required:"false"`

	// 短信模板内容，说明如下：字数不超过500，每个中文、符号、英文、数组等都计为一个字；模板中的变量填写格式：{N}，其中N为大于1的整数，有多个参数时，建议N从1开始顺次，例如：{1}、{2}等；短信模板禁止仅包括变量的情况；
	Template *string `required:"true"`

	// 短信模板名称，不超过32个字符，每个中文、符号、英文、数字等都计为1个字。
	TemplateName *string `required:"true"`

	// 当Purpose为3时，也即会员推广类短信模板，该项必填。枚举值：TD退订、回T退订、回N退订、回TD退订、退订回T、退订回D、退订回TD、退订回复T、退订回复D、退订回复N、退订回复TD、拒收回T
	UnsubscribeInfo *string `required:"false"`
}

// CreateUSMSTemplateResponse is response schema for CreateUSMSTemplate action
type CreateUSMSTemplateResponse struct {
	response.CommonBase

	// 返回状态码描述，如果操作成功，默认返回为空
	Message string

	// 短信模板ID（短信模板申请时的工单ID）
	TemplateId string
}

// NewCreateUSMSTemplateRequest will create request of CreateUSMSTemplate action.
func (c *USMSClient) NewCreateUSMSTemplateRequest() *CreateUSMSTemplateRequest {
	req := &CreateUSMSTemplateRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateUSMSTemplate

调用接口CreateUSMSTemplate申请短信模板
*/
func (c *USMSClient) CreateUSMSTemplate(req *CreateUSMSTemplateRequest) (*CreateUSMSTemplateResponse, error) {
	var err error
	var res CreateUSMSTemplateResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateUSMSTemplate", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteUSMSSignatureRequest is request schema for DeleteUSMSSignature action
type DeleteUSMSSignatureRequest struct {
	request.CommonBase

	// [公共参数] 项目ID，不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// 签名ID（也即短信签名申请时的工单ID），支持以数组的方式，举例，以SigIds.0、SigIds.1...SigIds.N方式传入
	SigIds []string `required:"true"`
}

// DeleteUSMSSignatureResponse is response schema for DeleteUSMSSignature action
type DeleteUSMSSignatureResponse struct {
	response.CommonBase

	// 返回状态码描述，如果操作成功，默认返回为空
	Message string
}

// NewDeleteUSMSSignatureRequest will create request of DeleteUSMSSignature action.
func (c *USMSClient) NewDeleteUSMSSignatureRequest() *DeleteUSMSSignatureRequest {
	req := &DeleteUSMSSignatureRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteUSMSSignature

调用接口DeleteUSMSSignature删除短信签名
*/
func (c *USMSClient) DeleteUSMSSignature(req *DeleteUSMSSignatureRequest) (*DeleteUSMSSignatureResponse, error) {
	var err error
	var res DeleteUSMSSignatureResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteUSMSSignature", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteUSMSTemplateRequest is request schema for DeleteUSMSTemplate action
type DeleteUSMSTemplateRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// 模板ID（也即短信模板申请时的工单ID），支持以数组的方式，举例，以TemplateIds.0、TemplateIds.1...TemplateIds.N方式传入
	TemplateIds []string `required:"true"`
}

// DeleteUSMSTemplateResponse is response schema for DeleteUSMSTemplate action
type DeleteUSMSTemplateResponse struct {
	response.CommonBase

	// 返回状态码描述，如果操作成功，默认返回为空
	Message string
}

// NewDeleteUSMSTemplateRequest will create request of DeleteUSMSTemplate action.
func (c *USMSClient) NewDeleteUSMSTemplateRequest() *DeleteUSMSTemplateRequest {
	req := &DeleteUSMSTemplateRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteUSMSTemplate

调用接口DeleteUSMSTemplate删除短信模板
*/
func (c *USMSClient) DeleteUSMSTemplate(req *DeleteUSMSTemplateRequest) (*DeleteUSMSTemplateResponse, error) {
	var err error
	var res DeleteUSMSTemplateResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteUSMSTemplate", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetUSMSSendReceiptRequest is request schema for GetUSMSSendReceipt action
type GetUSMSSendReceiptRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"false"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"false"`

	// 发送短信时返回的SessionNo集合，SessionNoSet.0,SessionNoSet.1....格式，单次调用集合数需控制在100个以内
	SessionNoSet []string `required:"true"`
}

// GetUSMSSendReceiptResponse is response schema for GetUSMSSendReceipt action
type GetUSMSSendReceiptResponse struct {
	response.CommonBase

	// 回执信息集合
	Data []ReceiptPerSession

	// 错误描述
	Message string
}

// NewGetUSMSSendReceiptRequest will create request of GetUSMSSendReceipt action.
func (c *USMSClient) NewGetUSMSSendReceiptRequest() *GetUSMSSendReceiptRequest {
	req := &GetUSMSSendReceiptRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetUSMSSendReceipt

调用接口GetUSMSSendReceipt短信发送状态信息
*/
func (c *USMSClient) GetUSMSSendReceipt(req *GetUSMSSendReceiptRequest) (*GetUSMSSendReceiptResponse, error) {
	var err error
	var res GetUSMSSendReceiptResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetUSMSSendReceipt", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// QueryUSMSSignatureRequest is request schema for QueryUSMSSignature action
type QueryUSMSSignatureRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 签名内容；签名ID和签名至少需填写1项；
	SigContent *string `required:"false"`

	// 已申请的短信签名ID（短信签名申请时的工单ID）；签名ID和签名至少需填写1项；
	SigId *string `required:"false"`
}

// QueryUSMSSignatureResponse is response schema for QueryUSMSSignature action
type QueryUSMSSignatureResponse struct {
	response.CommonBase

	// 签名信息
	Data OutSignature

	// 发生错误时，表示具体错误描述
	Message string
}

// NewQueryUSMSSignatureRequest will create request of QueryUSMSSignature action.
func (c *USMSClient) NewQueryUSMSSignatureRequest() *QueryUSMSSignatureRequest {
	req := &QueryUSMSSignatureRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: QueryUSMSSignature

调用接口QueryUSMSSignature查询短信签名申请状态
*/
func (c *USMSClient) QueryUSMSSignature(req *QueryUSMSSignatureRequest) (*QueryUSMSSignatureResponse, error) {
	var err error
	var res QueryUSMSSignatureResponse

	reqCopier := *req

	err = c.Client.InvokeAction("QueryUSMSSignature", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// QueryUSMSTemplateRequest is request schema for QueryUSMSTemplate action
type QueryUSMSTemplateRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 模板ID
	TemplateId *string `required:"true"`
}

// QueryUSMSTemplateResponse is response schema for QueryUSMSTemplate action
type QueryUSMSTemplateResponse struct {
	response.CommonBase

	// 短信模板明细信息，各字段说明详见OutTemplate
	Data OutTemplate

	// 当RetCode不为0时，Message中显示具体错误描述
	Message string
}

// NewQueryUSMSTemplateRequest will create request of QueryUSMSTemplate action.
func (c *USMSClient) NewQueryUSMSTemplateRequest() *QueryUSMSTemplateRequest {
	req := &QueryUSMSTemplateRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: QueryUSMSTemplate

调用接口QueryUSMSTemplate查询短信模板申请状态
*/
func (c *USMSClient) QueryUSMSTemplate(req *QueryUSMSTemplateRequest) (*QueryUSMSTemplateResponse, error) {
	var err error
	var res QueryUSMSTemplateResponse

	reqCopier := *req

	err = c.Client.InvokeAction("QueryUSMSTemplate", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// SendBatchUSMSMessageRequest is request schema for SendBatchUSMSMessage action
type SendBatchUSMSMessageRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 批量发送内容，该参数是json数组的base64编码结果。发送内容json数组中，每个“模板+签名”组合作为一个子项，每个子项内支持多个号码，示例：发送内容json数组（base64编码前）：[{"TemplateId": "UTA20212831C85C", "SigContent": "UCloud", "Target": [{"TemplateParams": ["123456"], "Phone": "18500000123", "ExtendCode": "123", "UserId": "456"} ] } ]   。json数组中各参数的定义："TemplateId":模板ID，"SigContent"短信签名内容，"Target"具体到号码粒度的发送内容。"Target"中的具体字段有："TemplateParams"实际发送的模板参数（若使用的是无参数模板，该参数不能传值），"Phone"手机号码, "ExtendCode"短信扩展码, "UserId"自定义业务标识ID。其中必传参数为"TemplateId", "SigContent", "Target"（"Target"中必传参数为"Phone"）。实际调用本接口时TaskContent传值（发送内容base64编码后）为：W3siVGVtcGxhdGVJZCI6ICJVVEEyMDIxMjgzMUM4NUMiLCAiU2lnQ29udGVudCI6ICJVQ2xvdWQiLCAiVGFyZ2V0IjogW3siVGVtcGxhdGVQYXJhbXMiOiBbIjEyMzQ1NiJdLCAiUGhvbmUiOiAiMTg1MDAwMDAxMjMiLCAiRXh0ZW5kQ29kZSI6ICIxMjMiLCAiVXNlcklkIjogIjQ1NiJ9IF0gfSBdIA==
	TaskContent *string `required:"true"`
}

// SendBatchUSMSMessageResponse is response schema for SendBatchUSMSMessage action
type SendBatchUSMSMessageResponse struct {
	response.CommonBase

	// 未发送成功的详情，返回码非0时该字段有效，可根据该字段数据重发
	FailContent []BatchInfo

	// 发生错误时表示错误描述
	Message string

	// 本次请求Uuid
	ReqUuid string

	// 本次提交发送任务的唯一ID，可根据该值查询本次发送的短信列表。注：成功提交短信数大于0时，才返回该字段
	SessionNo string

	// 成功提交短信（未拆分）条数
	SuccessCount int
}

// NewSendBatchUSMSMessageRequest will create request of SendBatchUSMSMessage action.
func (c *USMSClient) NewSendBatchUSMSMessageRequest() *SendBatchUSMSMessageRequest {
	req := &SendBatchUSMSMessageRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: SendBatchUSMSMessage

调用SendBatchUSMSMessage接口批量发送短信
*/
func (c *USMSClient) SendBatchUSMSMessage(req *SendBatchUSMSMessageRequest) (*SendBatchUSMSMessageResponse, error) {
	var err error
	var res SendBatchUSMSMessageResponse

	reqCopier := *req

	err = c.Client.InvokeAction("SendBatchUSMSMessage", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// SendUSMSMessageRequest is request schema for SendUSMSMessage action
type SendUSMSMessageRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 短信扩展码，格式为阿拉伯数字串，默认不开通，如需开通请联系 UCloud技术支持
	ExtendCode *string `required:"false"`

	// 电话号码数组，电话号码格式为(60)1xxxxxxxx，()中为国际长途区号(如中国为86或0086，两种格式都支持)，后面为电话号码.若不传入国际区号，如1851623xxxx，则默认为国内手机号
	PhoneNumbers []string `required:"true"`

	// 短信签名内容，请到[USMS控制台](https://console.ucloud.cn/usms)的签名管理页面查看；使用的短信签名必须是已申请并且通过审核；
	SigContent *string `required:"false"`

	// 模板ID（也即短信模板申请时的工单ID），请到[USMS控制台](https://console.ucloud.cn/usms)的模板管理页面查看；使用的短信模板必须是已申请并通过审核；
	TemplateId *string `required:"true"`

	// 模板可变参数，以数组的方式填写，举例，TemplateParams.0，TemplateParams.1，... 若模板中无可变参数，则该项可不填写；若模板中有可变参数，则该项为必填项，参数个数需与变量个数保持一致，否则无法发送；
	TemplateParams []string `required:"false"`

	// 自定义的业务标识ID，字符串（ 长度不能超过32 位），不支持 单引号、表情包符号等特殊字符
	UserId *string `required:"false"`
}

// SendUSMSMessageResponse is response schema for SendUSMSMessage action
type SendUSMSMessageResponse struct {
	response.CommonBase

	// 发生错误时表示错误描述
	Message string

	// 本次提交发送的短信的唯一ID，可根据该值查询本次发送的短信列表
	SessionNo string

	// 本次提交的自定义业务标识ID，仅当发送时传入有效的UserId，才返回该字段。
	UserId string
}

// NewSendUSMSMessageRequest will create request of SendUSMSMessage action.
func (c *USMSClient) NewSendUSMSMessageRequest() *SendUSMSMessageRequest {
	req := &SendUSMSMessageRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: SendUSMSMessage

调用接口SendUSMSMessage发送短信
*/
func (c *USMSClient) SendUSMSMessage(req *SendUSMSMessageRequest) (*SendUSMSMessageResponse, error) {
	var err error
	var res SendUSMSMessageResponse

	reqCopier := *req

	err = c.Client.InvokeAction("SendUSMSMessage", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// UpdateUSMSSignatureRequest is request schema for UpdateUSMSSignature action
type UpdateUSMSSignatureRequest struct {
	request.CommonBase

	// [公共参数] 项目ID，不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 签名的资质证明文件类型，需与签名类型保持一致，说明如下：0-三证合一/企业营业执照/组织机构代码证书/社会信用代码证书；1-应用商店后台开发者管理截图；2-备案服务商的备案成功截图(含域名，网站名称，备案号)；3-公众号或小程序的管理界面截图；4-商标注册证书；5-组织机构代码证书、社会信用代码证书；
	CertificateType *int `required:"false"`

	// 短信签名的资质证明文件URL，若未更改审核材料，则该处使用已上传审核材料的URL链接，否则使用File参数
	Document *string `required:"false"`

	// 短信签名的资质证明文件内容，需先进行base64编码格式转换，此处填写转换后的字符串。文件大小不超过4 MB。内容格式如下: [file type];[code type],[base64]  如：image/jpeg;base64,5YaF5a65
	File *string `required:"false"`

	// 短信签名授权委托文件URL，若未更改授权委托文件，则该处填写已上传的授权委托文件的URL链接，否则使用ProxyFile参数
	ProxyDoc *string `required:"false"`

	// 短信签名授权委托文件内容，需先进行base64编码格式转换，此处填写转换后的字符串。文件大小不超过4 MB；当您是代理并使用第三方的签名时（也即SigPurpose为1-他用），该项为必填项；格式和File类似。
	ProxyFile *string `required:"false"`

	// 新的短信签名内容；长度为2-12个字符, 可包含中文、数字和符号；无需填写【】或[]，系统会自动添加
	SigContent *string `required:"true"`

	// 签名ID（也即短信签名申请时的工单ID），支持以数组的方式，举例，以SigIds.0、SigIds.1...SigIds.N方式传入
	SigId *string `required:"true"`

	// 签名用途，0-自用，1-他用；
	SigPurpose *int `required:"true"`

	// 签名类型，说明如下：0-公司或企业的全称或简称；1-App应用的全称或简称；2-工信部备案网站的全称或简称；3-公众号或小程序的全称或简称；4-商标名的全称或简称；5-政府/机关事业单位/其他单位的全称或简称；
	SigType *int `required:"true"`
}

// UpdateUSMSSignatureResponse is response schema for UpdateUSMSSignature action
type UpdateUSMSSignatureResponse struct {
	response.CommonBase

	// 返回状态码描述，如果操作成功，默认返回为空
	Message string
}

// NewUpdateUSMSSignatureRequest will create request of UpdateUSMSSignature action.
func (c *USMSClient) NewUpdateUSMSSignatureRequest() *UpdateUSMSSignatureRequest {
	req := &UpdateUSMSSignatureRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: UpdateUSMSSignature

调用接口UpdateUSMSSignature修改未通过审核的短信签名，并重新提交审核
*/
func (c *USMSClient) UpdateUSMSSignature(req *UpdateUSMSSignatureRequest) (*UpdateUSMSSignatureResponse, error) {
	var err error
	var res UpdateUSMSSignatureResponse

	reqCopier := *req

	err = c.Client.InvokeAction("UpdateUSMSSignature", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// UpdateUSMSTemplateRequest is request schema for UpdateUSMSTemplate action
type UpdateUSMSTemplateRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// 模板变量属性说明
	Instruction *string `required:"false"`

	// 短信模板申请原因说明，字数不超过128，每个中文、符号、英文、数字等都计为1个字。
	Remark *string `required:"false"`

	// 新的模板内容。模板名称和模板内容必须提供一个，否则会报错。小于等于600个字
	Template *string `required:"true"`

	// 短信模板ID
	TemplateId *string `required:"true"`

	// 新的模板名称。小于等于32个字，每个中文、英文、数组、符合都计为一个字
	TemplateName *string `required:"false"`
}

// UpdateUSMSTemplateResponse is response schema for UpdateUSMSTemplate action
type UpdateUSMSTemplateResponse struct {
	response.CommonBase

	// 发生错误时表示错误描述
	Message string
}

// NewUpdateUSMSTemplateRequest will create request of UpdateUSMSTemplate action.
func (c *USMSClient) NewUpdateUSMSTemplateRequest() *UpdateUSMSTemplateRequest {
	req := &UpdateUSMSTemplateRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: UpdateUSMSTemplate

调用接口UpdateUSMSTemplate修改未通过审核的短信模板，并重新提交审核
*/
func (c *USMSClient) UpdateUSMSTemplate(req *UpdateUSMSTemplateRequest) (*UpdateUSMSTemplateResponse, error) {
	var err error
	var res UpdateUSMSTemplateResponse

	reqCopier := *req

	err = c.Client.InvokeAction("UpdateUSMSTemplate", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}