// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetConnectionTicketRequest struct {
	// example:
	//
	// ca-etn4zizgaezo9gis9
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// aig-bw1o1gcwvd3e1ipeu
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId      *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// example:
	//
	// 1.0.0.1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// f2463208-ec89-4309-8e8c-8b17acdcab93
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 22.21.2.21
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_"Windows 10 Enterprise LTSC 2019" 10.0 (Build 17763)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 2.0.1-D-20211008.101607
	ClientVersion        *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// example:
	//
	// test.test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// example:
	//
	// v1c4e2ef03d620f0f6cb41634196161219054e12d8aa5a13deb9ed14eebb76d674559115ad2e27a57f6820c1fd33e0ca36
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// /home/test/test.jpg
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceId  *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 09e2b2e6-3181-4c84-9539-6fc9f1c3199e
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 6f41731b-7091-4954-80c8-1d1e0b3ebb48
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1126819517152528
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// A8B35215993FBF283F28D617975204C4
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) SetAppId(v string) *GetConnectionTicketRequest {
	s.AppId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceGroupId(v string) *GetConnectionTicketRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceId(v string) *GetConnectionTicketRequest {
	s.AppInstanceId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppVersion(v string) *GetConnectionTicketRequest {
	s.AppVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetBizRegionId(v string) *GetConnectionTicketRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientId(v string) *GetConnectionTicketRequest {
	s.ClientId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientIp(v string) *GetConnectionTicketRequest {
	s.ClientIp = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientOS(v string) *GetConnectionTicketRequest {
	s.ClientOS = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientVersion(v string) *GetConnectionTicketRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetConnectionProperties(v string) *GetConnectionTicketRequest {
	s.ConnectionProperties = &v
	return s
}

func (s *GetConnectionTicketRequest) SetEndUserId(v string) *GetConnectionTicketRequest {
	s.EndUserId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetLoginRegionId(v string) *GetConnectionTicketRequest {
	s.LoginRegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetLoginToken(v string) *GetConnectionTicketRequest {
	s.LoginToken = &v
	return s
}

func (s *GetConnectionTicketRequest) SetParam(v string) *GetConnectionTicketRequest {
	s.Param = &v
	return s
}

func (s *GetConnectionTicketRequest) SetProductType(v string) *GetConnectionTicketRequest {
	s.ProductType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceId(v string) *GetConnectionTicketRequest {
	s.ResourceId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetSessionId(v string) *GetConnectionTicketRequest {
	s.SessionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTenantId(v string) *GetConnectionTicketRequest {
	s.TenantId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetUuid(v string) *GetConnectionTicketRequest {
	s.Uuid = &v
	return s
}

type GetConnectionTicketResponseBody struct {
	// example:
	//
	// aig-53fvrq1oanz6cxzi3
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// example:
	//
	// ai-gc1gemx6vpa6vlync
	AppInstanceId           *string                                       `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppInstancePersistentId *string                                       `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	BindQueueInfo           *GetConnectionTicketResponseBodyBindQueueInfo `json:"BindQueueInfo,omitempty" xml:"BindQueueInfo,omitempty" type:"Struct"`
	// example:
	//
	// InternalError.TicketGenInternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// v15418e331d8d068c29411646996786785303b8f61fd880aeaa50c5b584443cd9e65cc7eec72acdaad0a844666a3b26dab
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// reenter app instance failed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// Windows
	OsType *string                                `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Policy *GetConnectionTicketResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AD2D0761-1FE5-549D-B169-D3F8D19C6CDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	RetryTimes *int32 `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	// example:
	//
	// f3d1b31c-605e-4d04-a896-015fc9fc03b4
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// Running
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TenantId   *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// n7n9bqZlPrvgUOPSJzfdb$89jWwlVISgrchpY0tOfVYGBBcdoPoH39PVHK63fQTEM14kzajQdWAnHTnSicc35W_eI2LbTSGKquKukwcU7opRwmInhtQH*mlmsZQ3ByOLYVmqI*1hFESs0
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetConnectionTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBody) SetAppInstanceGroupId(v string) *GetConnectionTicketResponseBody {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAppInstanceId(v string) *GetConnectionTicketResponseBody {
	s.AppInstanceId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAppInstancePersistentId(v string) *GetConnectionTicketResponseBody {
	s.AppInstancePersistentId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetBindQueueInfo(v *GetConnectionTicketResponseBodyBindQueueInfo) *GetConnectionTicketResponseBody {
	s.BindQueueInfo = v
	return s
}

func (s *GetConnectionTicketResponseBody) SetCode(v string) *GetConnectionTicketResponseBody {
	s.Code = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetLoginToken(v string) *GetConnectionTicketResponseBody {
	s.LoginToken = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetMessage(v string) *GetConnectionTicketResponseBody {
	s.Message = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetOsType(v string) *GetConnectionTicketResponseBody {
	s.OsType = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetPolicy(v *GetConnectionTicketResponseBodyPolicy) *GetConnectionTicketResponseBody {
	s.Policy = v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRegionId(v string) *GetConnectionTicketResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRequestId(v string) *GetConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRetryTimes(v int32) *GetConnectionTicketResponseBody {
	s.RetryTimes = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskId(v string) *GetConnectionTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskStatus(v string) *GetConnectionTicketResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTenantId(v int64) *GetConnectionTicketResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTicket(v string) *GetConnectionTicketResponseBody {
	s.Ticket = &v
	return s
}

type GetConnectionTicketResponseBodyBindQueueInfo struct {
	Length           *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	Rank             *int32  `json:"Rank,omitempty" xml:"Rank,omitempty"`
	RemainingTimeMin *int32  `json:"RemainingTimeMin,omitempty" xml:"RemainingTimeMin,omitempty"`
	RequestKey       *string `json:"RequestKey,omitempty" xml:"RequestKey,omitempty"`
	TargetId         *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	WaitTimeMin      *int32  `json:"WaitTimeMin,omitempty" xml:"WaitTimeMin,omitempty"`
}

func (s GetConnectionTicketResponseBodyBindQueueInfo) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponseBodyBindQueueInfo) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetLength(v int32) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.Length = &v
	return s
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetRank(v int32) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.Rank = &v
	return s
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetRemainingTimeMin(v int32) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.RemainingTimeMin = &v
	return s
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetRequestKey(v string) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.RequestKey = &v
	return s
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetTargetId(v string) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.TargetId = &v
	return s
}

func (s *GetConnectionTicketResponseBodyBindQueueInfo) SetWaitTimeMin(v int32) *GetConnectionTicketResponseBodyBindQueueInfo {
	s.WaitTimeMin = &v
	return s
}

type GetConnectionTicketResponseBodyPolicy struct {
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" xml:"ResolutionAdaptive,omitempty"`
	ResolutionHeight   *int32  `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	ResolutionWidth    *int32  `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s GetConnectionTicketResponseBodyPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBodyPolicy) SetResolutionAdaptive(v string) *GetConnectionTicketResponseBodyPolicy {
	s.ResolutionAdaptive = &v
	return s
}

func (s *GetConnectionTicketResponseBodyPolicy) SetResolutionHeight(v int32) *GetConnectionTicketResponseBodyPolicy {
	s.ResolutionHeight = &v
	return s
}

func (s *GetConnectionTicketResponseBodyPolicy) SetResolutionWidth(v int32) *GetConnectionTicketResponseBodyPolicy {
	s.ResolutionWidth = &v
	return s
}

type GetConnectionTicketResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConnectionTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConnectionTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponse) SetHeaders(v map[string]*string) *GetConnectionTicketResponse {
	s.Headers = v
	return s
}

func (s *GetConnectionTicketResponse) SetStatusCode(v int32) *GetConnectionTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnectionTicketResponse) SetBody(v *GetConnectionTicketResponseBody) *GetConnectionTicketResponse {
	s.Body = v
	return s
}

type ListLFUAppRequest struct {
	AliUid             *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ApiType            *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	BizRegionId        *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ClientChannel      *string `json:"ClientChannel,omitempty" xml:"ClientChannel,omitempty"`
	ClientId           *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientIp           *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientOS           *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion      *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	EndUserId          *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ExtendsAccessToken *string `json:"ExtendsAccessToken,omitempty" xml:"ExtendsAccessToken,omitempty"`
	IdpId              *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	LoginRegionId      *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	LoginToken         *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId          *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TraceId            *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	Uuid               *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	WyId               *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s ListLFUAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLFUAppRequest) GoString() string {
	return s.String()
}

func (s *ListLFUAppRequest) SetAliUid(v int64) *ListLFUAppRequest {
	s.AliUid = &v
	return s
}

func (s *ListLFUAppRequest) SetApiType(v string) *ListLFUAppRequest {
	s.ApiType = &v
	return s
}

func (s *ListLFUAppRequest) SetBizRegionId(v string) *ListLFUAppRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListLFUAppRequest) SetClientChannel(v string) *ListLFUAppRequest {
	s.ClientChannel = &v
	return s
}

func (s *ListLFUAppRequest) SetClientId(v string) *ListLFUAppRequest {
	s.ClientId = &v
	return s
}

func (s *ListLFUAppRequest) SetClientIp(v string) *ListLFUAppRequest {
	s.ClientIp = &v
	return s
}

func (s *ListLFUAppRequest) SetClientOS(v string) *ListLFUAppRequest {
	s.ClientOS = &v
	return s
}

func (s *ListLFUAppRequest) SetClientVersion(v string) *ListLFUAppRequest {
	s.ClientVersion = &v
	return s
}

func (s *ListLFUAppRequest) SetEndUserId(v string) *ListLFUAppRequest {
	s.EndUserId = &v
	return s
}

func (s *ListLFUAppRequest) SetExtendsAccessToken(v string) *ListLFUAppRequest {
	s.ExtendsAccessToken = &v
	return s
}

func (s *ListLFUAppRequest) SetIdpId(v string) *ListLFUAppRequest {
	s.IdpId = &v
	return s
}

func (s *ListLFUAppRequest) SetLoginRegionId(v string) *ListLFUAppRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ListLFUAppRequest) SetLoginToken(v string) *ListLFUAppRequest {
	s.LoginToken = &v
	return s
}

func (s *ListLFUAppRequest) SetProductType(v string) *ListLFUAppRequest {
	s.ProductType = &v
	return s
}

func (s *ListLFUAppRequest) SetRegionId(v string) *ListLFUAppRequest {
	s.RegionId = &v
	return s
}

func (s *ListLFUAppRequest) SetSessionId(v string) *ListLFUAppRequest {
	s.SessionId = &v
	return s
}

func (s *ListLFUAppRequest) SetTraceId(v string) *ListLFUAppRequest {
	s.TraceId = &v
	return s
}

func (s *ListLFUAppRequest) SetUuid(v string) *ListLFUAppRequest {
	s.Uuid = &v
	return s
}

func (s *ListLFUAppRequest) SetWyId(v string) *ListLFUAppRequest {
	s.WyId = &v
	return s
}

type ListLFUAppResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Count          *int32                        `json:"Count,omitempty" xml:"Count,omitempty"`
	Data           []*ListLFUAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLFUAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLFUAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListLFUAppResponseBody) SetCode(v string) *ListLFUAppResponseBody {
	s.Code = &v
	return s
}

func (s *ListLFUAppResponseBody) SetCount(v int32) *ListLFUAppResponseBody {
	s.Count = &v
	return s
}

func (s *ListLFUAppResponseBody) SetData(v []*ListLFUAppResponseBodyData) *ListLFUAppResponseBody {
	s.Data = v
	return s
}

func (s *ListLFUAppResponseBody) SetHttpStatusCode(v int32) *ListLFUAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLFUAppResponseBody) SetMessage(v string) *ListLFUAppResponseBody {
	s.Message = &v
	return s
}

func (s *ListLFUAppResponseBody) SetRequestId(v string) *ListLFUAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLFUAppResponseBody) SetSuccess(v bool) *ListLFUAppResponseBody {
	s.Success = &v
	return s
}

type ListLFUAppResponseBodyData struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppVersion     *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	IconUrl        *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	IsAuth         *bool   `json:"IsAuth,omitempty" xml:"IsAuth,omitempty"`
	OsType         *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLFUAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListLFUAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLFUAppResponseBodyData) SetAppId(v string) *ListLFUAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetAppName(v string) *ListLFUAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetAppVersion(v string) *ListLFUAppResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetAppVersionName(v string) *ListLFUAppResponseBodyData {
	s.AppVersionName = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetIconUrl(v string) *ListLFUAppResponseBodyData {
	s.IconUrl = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetIsAuth(v bool) *ListLFUAppResponseBodyData {
	s.IsAuth = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetOsType(v string) *ListLFUAppResponseBodyData {
	s.OsType = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetRequestId(v string) *ListLFUAppResponseBodyData {
	s.RequestId = &v
	return s
}

type ListLFUAppResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLFUAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLFUAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLFUAppResponse) GoString() string {
	return s.String()
}

func (s *ListLFUAppResponse) SetHeaders(v map[string]*string) *ListLFUAppResponse {
	s.Headers = v
	return s
}

func (s *ListLFUAppResponse) SetStatusCode(v int32) *ListLFUAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLFUAppResponse) SetBody(v *ListLFUAppResponseBody) *ListLFUAppResponse {
	s.Body = v
	return s
}

type ListPublishedAppInfoRequest struct {
	// example:
	//
	// Microsoft Word
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// 1
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 1
	CategoryType *int64 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// 17b38aaa-761f-44c5-9862-2ad0f5025d15
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 125.80.132.13
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_\\"Windows 10 Enterprise\\" 10.0 (Build 14393)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 2.0.1-D-20211008.101607
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// test.test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// cn-shanghai
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// example:
	//
	// v189fa78c1aff77a0483b16497517322299131027b85bb84bbdc0871988ce8296d8fd891e2fdeaded3bd75f81f639acee8
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OrderParam *string `json:"OrderParam,omitempty" xml:"OrderParam,omitempty"`
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// c261a6a1-e242-4f4b-813c-5fe807e49f03
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SortType  *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListPublishedAppInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAppInfoRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInfoRequest) SetAppName(v string) *ListPublishedAppInfoRequest {
	s.AppName = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetBizRegionId(v string) *ListPublishedAppInfoRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetCategoryId(v int64) *ListPublishedAppInfoRequest {
	s.CategoryId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetCategoryType(v int64) *ListPublishedAppInfoRequest {
	s.CategoryType = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetClientId(v string) *ListPublishedAppInfoRequest {
	s.ClientId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetClientIp(v string) *ListPublishedAppInfoRequest {
	s.ClientIp = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetClientOS(v string) *ListPublishedAppInfoRequest {
	s.ClientOS = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetClientVersion(v string) *ListPublishedAppInfoRequest {
	s.ClientVersion = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetEndUserId(v string) *ListPublishedAppInfoRequest {
	s.EndUserId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetLoginRegionId(v string) *ListPublishedAppInfoRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetLoginToken(v string) *ListPublishedAppInfoRequest {
	s.LoginToken = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetOrderParam(v string) *ListPublishedAppInfoRequest {
	s.OrderParam = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetProductType(v string) *ListPublishedAppInfoRequest {
	s.ProductType = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetSessionId(v string) *ListPublishedAppInfoRequest {
	s.SessionId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetSortType(v string) *ListPublishedAppInfoRequest {
	s.SortType = &v
	return s
}

type ListPublishedAppInfoResponseBody struct {
	// appModels
	AppModels []*ListPublishedAppInfoResponseBodyAppModels `json:"AppModels,omitempty" xml:"AppModels,omitempty" type:"Repeated"`
	// example:
	//
	// 2NVfhLfgy5b3J5iJyoLQ6x4EULMg1hbhgB9NfnvdK9oj5zwxd17j4TuQkZze3RvhEvBinZYjknujF3Q1M
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DB70F8FE-63A3-587B-8560-CEC258E8B944
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPublishedAppInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInfoResponseBody) SetAppModels(v []*ListPublishedAppInfoResponseBodyAppModels) *ListPublishedAppInfoResponseBody {
	s.AppModels = v
	return s
}

func (s *ListPublishedAppInfoResponseBody) SetNextToken(v string) *ListPublishedAppInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublishedAppInfoResponseBody) SetRequestId(v string) *ListPublishedAppInfoResponseBody {
	s.RequestId = &v
	return s
}

type ListPublishedAppInfoResponseBodyAppModels struct {
	// example:
	//
	// img-f37nddbjc1lje14st
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// example:
	//
	// ca-fxwp4koyr5y2sp4mz
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// Microsoft Word
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppThemeColor *string `json:"AppThemeColor,omitempty" xml:"AppThemeColor,omitempty"`
	// example:
	//
	// R2021a
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// v1.0
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	AuthTime       *string `json:"AuthTime,omitempty" xml:"AuthTime,omitempty"`
	// example:
	//
	// 2
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 1
	CategoryType *int64 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// https://app-streaming-icon-prod-shanghai.oss-cn-shanghai.aliyuncs.com/tenant/1973619010349344/1634523814270_Matlab.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// example:
	//
	// True
	IsAuth *bool `json:"IsAuth,omitempty" xml:"IsAuth,omitempty"`
	// example:
	//
	// True
	UsedInSession *bool `json:"UsedInSession,omitempty" xml:"UsedInSession,omitempty"`
}

func (s ListPublishedAppInfoResponseBodyAppModels) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAppInfoResponseBodyAppModels) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppCenterImageId(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppCenterImageId = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppId(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppId = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppName(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppName = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppThemeColor(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppThemeColor = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppVersion(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppVersion = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppVersionName(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppVersionName = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAuthTime(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AuthTime = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetCategoryId(v int64) *ListPublishedAppInfoResponseBodyAppModels {
	s.CategoryId = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetCategoryType(v int64) *ListPublishedAppInfoResponseBodyAppModels {
	s.CategoryType = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetIconUrl(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.IconUrl = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetIsAuth(v bool) *ListPublishedAppInfoResponseBodyAppModels {
	s.IsAuth = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetUsedInSession(v bool) *ListPublishedAppInfoResponseBodyAppModels {
	s.UsedInSession = &v
	return s
}

type ListPublishedAppInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishedAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishedAppInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAppInfoResponse) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInfoResponse) SetHeaders(v map[string]*string) *ListPublishedAppInfoResponse {
	s.Headers = v
	return s
}

func (s *ListPublishedAppInfoResponse) SetStatusCode(v int32) *ListPublishedAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishedAppInfoResponse) SetBody(v *ListPublishedAppInfoResponseBody) *ListPublishedAppInfoResponse {
	s.Body = v
	return s
}

type ListRunningAppsRequest struct {
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// 370b56f8-2812-4b6c-bfa6-2560791cad88
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 22.21.2.32
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_"Windows 10 Enterprise" 10.0 (Build 18363)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 0.1.0-R-20220512.175656
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// test.test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// example:
	//
	// v1124500957832f30b3e716406562071655aa43b2a723ed2be0837815483d54e025db13ba5469f06f2410d0efc4d302e36
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// a863f4c3-2f1d-4971-8cf7-e2b92ae97764
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1735953493960828
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 470E8C12AB78CE9C3F6627DD0409E51D
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListRunningAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRunningAppsRequest) GoString() string {
	return s.String()
}

func (s *ListRunningAppsRequest) SetBizRegionId(v string) *ListRunningAppsRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListRunningAppsRequest) SetClientId(v string) *ListRunningAppsRequest {
	s.ClientId = &v
	return s
}

func (s *ListRunningAppsRequest) SetClientIp(v string) *ListRunningAppsRequest {
	s.ClientIp = &v
	return s
}

func (s *ListRunningAppsRequest) SetClientOS(v string) *ListRunningAppsRequest {
	s.ClientOS = &v
	return s
}

func (s *ListRunningAppsRequest) SetClientVersion(v string) *ListRunningAppsRequest {
	s.ClientVersion = &v
	return s
}

func (s *ListRunningAppsRequest) SetEndUserId(v string) *ListRunningAppsRequest {
	s.EndUserId = &v
	return s
}

func (s *ListRunningAppsRequest) SetLoginRegionId(v string) *ListRunningAppsRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ListRunningAppsRequest) SetLoginToken(v string) *ListRunningAppsRequest {
	s.LoginToken = &v
	return s
}

func (s *ListRunningAppsRequest) SetProductType(v string) *ListRunningAppsRequest {
	s.ProductType = &v
	return s
}

func (s *ListRunningAppsRequest) SetSessionId(v string) *ListRunningAppsRequest {
	s.SessionId = &v
	return s
}

func (s *ListRunningAppsRequest) SetTenantId(v string) *ListRunningAppsRequest {
	s.TenantId = &v
	return s
}

func (s *ListRunningAppsRequest) SetUuid(v string) *ListRunningAppsRequest {
	s.Uuid = &v
	return s
}

type ListRunningAppsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2DC3521C-3820-5EA5-9A9A-00BB7AF4E8E5
	RequestId        *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunningCloudApps []*ListRunningAppsResponseBodyRunningCloudApps `json:"RunningCloudApps,omitempty" xml:"RunningCloudApps,omitempty" type:"Repeated"`
}

func (s ListRunningAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRunningAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRunningAppsResponseBody) SetRequestId(v string) *ListRunningAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRunningAppsResponseBody) SetRunningCloudApps(v []*ListRunningAppsResponseBodyRunningCloudApps) *ListRunningAppsResponseBody {
	s.RunningCloudApps = v
	return s
}

type ListRunningAppsResponseBodyRunningCloudApps struct {
	// example:
	//
	// ca-dln05y44ze6esfl8x
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// aig-dk8p95irk9xs5xi6a
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// example:
	//
	// ai-gc1gemx6vpa6vlync
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// example:
	//
	// alihealth-keeper
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 11.1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// test1.0
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	// example:
	//
	// 87
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// https://app-icon-shanghai.oss-cn-shanghai.aliyuncs.com/tenant/187465/18_bf1.jpg
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1642748400
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListRunningAppsResponseBodyRunningCloudApps) String() string {
	return tea.Prettify(s)
}

func (s ListRunningAppsResponseBodyRunningCloudApps) GoString() string {
	return s.String()
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppId(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppId = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppInstanceGroupId(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppInstanceId(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppInstanceId = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppName(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppName = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppVersion(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppVersion = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppVersionName(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppVersionName = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetDuration(v int64) *ListRunningAppsResponseBodyRunningCloudApps {
	s.Duration = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetIconUrl(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.IconUrl = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetOsType(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.OsType = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetRegionId(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.RegionId = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetStartTime(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.StartTime = &v
	return s
}

type ListRunningAppsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRunningAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRunningAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRunningAppsResponse) GoString() string {
	return s.String()
}

func (s *ListRunningAppsResponse) SetHeaders(v map[string]*string) *ListRunningAppsResponse {
	s.Headers = v
	return s
}

func (s *ListRunningAppsResponse) SetStatusCode(v int32) *ListRunningAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRunningAppsResponse) SetBody(v *ListRunningAppsResponseBody) *ListRunningAppsResponse {
	s.Body = v
	return s
}

type StopAppRequest struct {
	// example:
	//
	// 1924794279035094
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// AnonymousUserAPI
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// example:
	//
	// ca-fxwp4koywsglzvvex
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// aig-89ibriac2wudyph38
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// example:
	//
	// ai-d297eyf83g5niwnjl
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// pc
	ClientChannel *string `json:"ClientChannel,omitempty" xml:"ClientChannel,omitempty"`
	// example:
	//
	// 91b79184-51d0-42ad-8475-78cae95b0aa6
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 22.21.2.79
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_"Windows 10 Enterprise" 10.0 (Build 19042)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 3.1.1-R-20211022.144255
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// test.test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// example:
	//
	// idp-9ie5smicnct2xodn2
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// example:
	//
	// v185fdd7f6d39fa7861981639366085772e150a390a5bb7b43c4e62440d94fc392b945770e1596cebe90085ce0af4d330e
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 597e869d-ea14-4b83-9490-714f68bfe935
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2943802884B27030B6759F9132B26903
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// ac3cb49059261898
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s StopAppRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAppRequest) GoString() string {
	return s.String()
}

func (s *StopAppRequest) SetAliUid(v int64) *StopAppRequest {
	s.AliUid = &v
	return s
}

func (s *StopAppRequest) SetApiType(v string) *StopAppRequest {
	s.ApiType = &v
	return s
}

func (s *StopAppRequest) SetAppId(v string) *StopAppRequest {
	s.AppId = &v
	return s
}

func (s *StopAppRequest) SetAppInstanceGroupId(v string) *StopAppRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *StopAppRequest) SetAppInstanceId(v string) *StopAppRequest {
	s.AppInstanceId = &v
	return s
}

func (s *StopAppRequest) SetBizRegionId(v string) *StopAppRequest {
	s.BizRegionId = &v
	return s
}

func (s *StopAppRequest) SetClientChannel(v string) *StopAppRequest {
	s.ClientChannel = &v
	return s
}

func (s *StopAppRequest) SetClientId(v string) *StopAppRequest {
	s.ClientId = &v
	return s
}

func (s *StopAppRequest) SetClientIp(v string) *StopAppRequest {
	s.ClientIp = &v
	return s
}

func (s *StopAppRequest) SetClientOS(v string) *StopAppRequest {
	s.ClientOS = &v
	return s
}

func (s *StopAppRequest) SetClientVersion(v string) *StopAppRequest {
	s.ClientVersion = &v
	return s
}

func (s *StopAppRequest) SetEndUserId(v string) *StopAppRequest {
	s.EndUserId = &v
	return s
}

func (s *StopAppRequest) SetForceStop(v bool) *StopAppRequest {
	s.ForceStop = &v
	return s
}

func (s *StopAppRequest) SetIdpId(v string) *StopAppRequest {
	s.IdpId = &v
	return s
}

func (s *StopAppRequest) SetLoginRegionId(v string) *StopAppRequest {
	s.LoginRegionId = &v
	return s
}

func (s *StopAppRequest) SetLoginToken(v string) *StopAppRequest {
	s.LoginToken = &v
	return s
}

func (s *StopAppRequest) SetProductType(v string) *StopAppRequest {
	s.ProductType = &v
	return s
}

func (s *StopAppRequest) SetRegionId(v string) *StopAppRequest {
	s.RegionId = &v
	return s
}

func (s *StopAppRequest) SetSessionId(v string) *StopAppRequest {
	s.SessionId = &v
	return s
}

func (s *StopAppRequest) SetUuid(v string) *StopAppRequest {
	s.Uuid = &v
	return s
}

func (s *StopAppRequest) SetWyId(v string) *StopAppRequest {
	s.WyId = &v
	return s
}

type StopAppResponseBody struct {
	// example:
	//
	// 83A9075B-C646-59A9-8232-CAE41AF4B9E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAppResponseBody) GoString() string {
	return s.String()
}

func (s *StopAppResponseBody) SetRequestId(v string) *StopAppResponseBody {
	s.RequestId = &v
	return s
}

type StopAppResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAppResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAppResponse) GoString() string {
	return s.String()
}

func (s *StopAppResponse) SetHeaders(v map[string]*string) *StopAppResponse {
	s.Headers = v
	return s
}

func (s *StopAppResponse) SetStatusCode(v int32) *StopAppResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAppResponse) SetBody(v *StopAppResponseBody) *StopAppResponse {
	s.Body = v
	return s
}

type UnbindRequest struct {
	// example:
	//
	// ca-fxwp4koxs8hopi94e
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aig-e1l4kqqykxt4uzdx9
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId      *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// example:
	//
	// eac19bef-1e45-4190-a03a-4ea74b699ca7
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 139.129.223.122
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_\\"Windows 10 Pro\\" 10.0 (Build 19041)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 2.0.1-D-20220303.171122
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// test.test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// cn-shanghai
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1e9c8e83d83ea11270871640059145702bde8c5be8c6b9a854ffb6a43bd2673c19a5551c83800724e024f488dbfb0b247
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11040139-4fb4-4b35-9b44-6c07c746a43e
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1569416393841402
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s UnbindRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindRequest) GoString() string {
	return s.String()
}

func (s *UnbindRequest) SetAppId(v string) *UnbindRequest {
	s.AppId = &v
	return s
}

func (s *UnbindRequest) SetAppInstanceGroupId(v string) *UnbindRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *UnbindRequest) SetAppInstanceId(v string) *UnbindRequest {
	s.AppInstanceId = &v
	return s
}

func (s *UnbindRequest) SetClientId(v string) *UnbindRequest {
	s.ClientId = &v
	return s
}

func (s *UnbindRequest) SetClientIp(v string) *UnbindRequest {
	s.ClientIp = &v
	return s
}

func (s *UnbindRequest) SetClientOS(v string) *UnbindRequest {
	s.ClientOS = &v
	return s
}

func (s *UnbindRequest) SetClientVersion(v string) *UnbindRequest {
	s.ClientVersion = &v
	return s
}

func (s *UnbindRequest) SetEndUserId(v string) *UnbindRequest {
	s.EndUserId = &v
	return s
}

func (s *UnbindRequest) SetLoginRegionId(v string) *UnbindRequest {
	s.LoginRegionId = &v
	return s
}

func (s *UnbindRequest) SetLoginToken(v string) *UnbindRequest {
	s.LoginToken = &v
	return s
}

func (s *UnbindRequest) SetProductType(v string) *UnbindRequest {
	s.ProductType = &v
	return s
}

func (s *UnbindRequest) SetSessionId(v string) *UnbindRequest {
	s.SessionId = &v
	return s
}

func (s *UnbindRequest) SetTenantId(v int64) *UnbindRequest {
	s.TenantId = &v
	return s
}

type UnbindResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6EBD4046-2202-5FBD-8595-4B631F0C484B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindResponseBody) SetRequestId(v string) *UnbindResponseBody {
	s.RequestId = &v
	return s
}

type UnbindResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindResponse) GoString() string {
	return s.String()
}

func (s *UnbindResponse) SetHeaders(v map[string]*string) *UnbindResponse {
	s.Headers = v
	return s
}

func (s *UnbindResponse) SetStatusCode(v int32) *UnbindResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindResponse) SetBody(v *UnbindResponseBody) *UnbindResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("appstream-center"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取连接信息
//
// @param request - GetConnectionTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectionTicketResponse
func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		body["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		body["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		body["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionProperties)) {
		body["ConnectionProperties"] = request.ConnectionProperties
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		body["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["Param"] = request.Param
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConnectionTicket"),
		Version:     tea.String("2021-09-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取连接信息
//
// @param request - GetConnectionTicketRequest
//
// @return GetConnectionTicketResponse
func (client *Client) GetConnectionTicket(request *GetConnectionTicketRequest) (_result *GetConnectionTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.GetConnectionTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListLFUAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLFUAppResponse
func (client *Client) ListLFUAppWithOptions(request *ListLFUAppRequest, runtime *util.RuntimeOptions) (_result *ListLFUAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		body["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		body["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientChannel)) {
		body["ClientChannel"] = request.ClientChannel
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		body["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		body["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendsAccessToken)) {
		body["ExtendsAccessToken"] = request.ExtendsAccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.IdpId)) {
		body["IdpId"] = request.IdpId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		body["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TraceId)) {
		body["TraceId"] = request.TraceId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	if !tea.BoolValue(util.IsUnset(request.WyId)) {
		body["WyId"] = request.WyId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLFUApp"),
		Version:     tea.String("2021-09-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLFUAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListLFUAppRequest
//
// @return ListLFUAppResponse
func (client *Client) ListLFUApp(request *ListLFUAppRequest) (_result *ListLFUAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLFUAppResponse{}
	_body, _err := client.ListLFUAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 已上架应用列表
//
// @param request - ListPublishedAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishedAppInfoResponse
func (client *Client) ListPublishedAppInfoWithOptions(request *ListPublishedAppInfoRequest, runtime *util.RuntimeOptions) (_result *ListPublishedAppInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryType)) {
		query["CategoryType"] = request.CategoryType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		query["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderParam)) {
		query["OrderParam"] = request.OrderParam
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublishedAppInfo"),
		Version:     tea.String("2021-09-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublishedAppInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 已上架应用列表
//
// @param request - ListPublishedAppInfoRequest
//
// @return ListPublishedAppInfoResponse
func (client *Client) ListPublishedAppInfo(request *ListPublishedAppInfoRequest) (_result *ListPublishedAppInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublishedAppInfoResponse{}
	_body, _err := client.ListPublishedAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 运行中应用列表
//
// @param request - ListRunningAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRunningAppsResponse
func (client *Client) ListRunningAppsWithOptions(request *ListRunningAppsRequest, runtime *util.RuntimeOptions) (_result *ListRunningAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		query["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRunningApps"),
		Version:     tea.String("2021-09-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRunningAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行中应用列表
//
// @param request - ListRunningAppsRequest
//
// @return ListRunningAppsResponse
func (client *Client) ListRunningApps(request *ListRunningAppsRequest) (_result *ListRunningAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRunningAppsResponse{}
	_body, _err := client.ListRunningAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止应用
//
// @param request - StopAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAppResponse
func (client *Client) StopAppWithOptions(request *StopAppRequest, runtime *util.RuntimeOptions) (_result *StopAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		body["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		body["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientChannel)) {
		body["ClientChannel"] = request.ClientChannel
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		body["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		body["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ForceStop)) {
		body["ForceStop"] = request.ForceStop
	}

	if !tea.BoolValue(util.IsUnset(request.IdpId)) {
		body["IdpId"] = request.IdpId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		body["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	if !tea.BoolValue(util.IsUnset(request.WyId)) {
		body["WyId"] = request.WyId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopApp"),
		Version:     tea.String("2021-09-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止应用
//
// @param request - StopAppRequest
//
// @return StopAppResponse
func (client *Client) StopApp(request *StopAppRequest) (_result *StopAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAppResponse{}
	_body, _err := client.StopAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑实例
//
// @param request - UnbindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindResponse
func (client *Client) UnbindWithOptions(request *UnbindRequest, runtime *util.RuntimeOptions) (_result *UnbindResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		body["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		body["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		body["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Unbind"),
		Version:     tea.String("2021-09-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑实例
//
// @param request - UnbindRequest
//
// @return UnbindResponse
func (client *Client) Unbind(request *UnbindRequest) (_result *UnbindResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindResponse{}
	_body, _err := client.UnbindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
