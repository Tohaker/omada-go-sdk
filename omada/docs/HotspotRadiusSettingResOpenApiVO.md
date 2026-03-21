# HotspotRadiusSettingResOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **int32** | RADIUS auth mode, should be a value as follows: 1: PAP; 2: CHAP | [optional] 
**AuthTimeout** | Pointer to [**AuthTimeOpenApiVO**](AuthTimeOpenApiVO.md) |  | [optional] 
**DisconnectReq** | Pointer to **bool** | Whether to support disconnect messages. Only for Omada Local Controller | [optional] 
**NasId** | Pointer to **string** | RADIUS Attribute: NasID, should contain 1 to 64 characters. | [optional] 
**RadiusProfileId** | Pointer to **string** | This field represents radius profile ID. Radius profile can be created using &#39;Create a new Radius profile&#39; (&#39;Create a new Radius profile template&#39;) interface, and radius profile ID can be obtained from &#39;Get Radius profile list&#39; (&#39;Get Radius profile template list&#39;) interface | [optional] 
**ReceiverPort** | Pointer to **int32** | Port for listening to disconnect messages, should be within the range of 1–65535.Only for Omada Local Controller | [optional] 
**ReceiverPortStatus** | Pointer to **int32** | Port binding status, should be a value as follow: 1: Disconnect Requests port status running, 2: Disconnect Requests port status disable. Only for Omada Local Controller | [optional] 

## Methods

### NewHotspotRadiusSettingResOpenApiVO

`func NewHotspotRadiusSettingResOpenApiVO() *HotspotRadiusSettingResOpenApiVO`

NewHotspotRadiusSettingResOpenApiVO instantiates a new HotspotRadiusSettingResOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHotspotRadiusSettingResOpenApiVOWithDefaults

`func NewHotspotRadiusSettingResOpenApiVOWithDefaults() *HotspotRadiusSettingResOpenApiVO`

NewHotspotRadiusSettingResOpenApiVOWithDefaults instantiates a new HotspotRadiusSettingResOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *HotspotRadiusSettingResOpenApiVO) GetAuthMode() int32`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *HotspotRadiusSettingResOpenApiVO) GetAuthModeOk() (*int32, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *HotspotRadiusSettingResOpenApiVO) SetAuthMode(v int32)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *HotspotRadiusSettingResOpenApiVO) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthTimeout

`func (o *HotspotRadiusSettingResOpenApiVO) GetAuthTimeout() AuthTimeOpenApiVO`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *HotspotRadiusSettingResOpenApiVO) GetAuthTimeoutOk() (*AuthTimeOpenApiVO, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *HotspotRadiusSettingResOpenApiVO) SetAuthTimeout(v AuthTimeOpenApiVO)`

SetAuthTimeout sets AuthTimeout field to given value.

### HasAuthTimeout

`func (o *HotspotRadiusSettingResOpenApiVO) HasAuthTimeout() bool`

HasAuthTimeout returns a boolean if a field has been set.

### GetDisconnectReq

`func (o *HotspotRadiusSettingResOpenApiVO) GetDisconnectReq() bool`

GetDisconnectReq returns the DisconnectReq field if non-nil, zero value otherwise.

### GetDisconnectReqOk

`func (o *HotspotRadiusSettingResOpenApiVO) GetDisconnectReqOk() (*bool, bool)`

GetDisconnectReqOk returns a tuple with the DisconnectReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectReq

`func (o *HotspotRadiusSettingResOpenApiVO) SetDisconnectReq(v bool)`

SetDisconnectReq sets DisconnectReq field to given value.

### HasDisconnectReq

`func (o *HotspotRadiusSettingResOpenApiVO) HasDisconnectReq() bool`

HasDisconnectReq returns a boolean if a field has been set.

### GetNasId

`func (o *HotspotRadiusSettingResOpenApiVO) GetNasId() string`

GetNasId returns the NasId field if non-nil, zero value otherwise.

### GetNasIdOk

`func (o *HotspotRadiusSettingResOpenApiVO) GetNasIdOk() (*string, bool)`

GetNasIdOk returns a tuple with the NasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasId

`func (o *HotspotRadiusSettingResOpenApiVO) SetNasId(v string)`

SetNasId sets NasId field to given value.

### HasNasId

`func (o *HotspotRadiusSettingResOpenApiVO) HasNasId() bool`

HasNasId returns a boolean if a field has been set.

### GetRadiusProfileId

`func (o *HotspotRadiusSettingResOpenApiVO) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *HotspotRadiusSettingResOpenApiVO) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *HotspotRadiusSettingResOpenApiVO) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.

### HasRadiusProfileId

`func (o *HotspotRadiusSettingResOpenApiVO) HasRadiusProfileId() bool`

HasRadiusProfileId returns a boolean if a field has been set.

### GetReceiverPort

`func (o *HotspotRadiusSettingResOpenApiVO) GetReceiverPort() int32`

GetReceiverPort returns the ReceiverPort field if non-nil, zero value otherwise.

### GetReceiverPortOk

`func (o *HotspotRadiusSettingResOpenApiVO) GetReceiverPortOk() (*int32, bool)`

GetReceiverPortOk returns a tuple with the ReceiverPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverPort

`func (o *HotspotRadiusSettingResOpenApiVO) SetReceiverPort(v int32)`

SetReceiverPort sets ReceiverPort field to given value.

### HasReceiverPort

`func (o *HotspotRadiusSettingResOpenApiVO) HasReceiverPort() bool`

HasReceiverPort returns a boolean if a field has been set.

### GetReceiverPortStatus

`func (o *HotspotRadiusSettingResOpenApiVO) GetReceiverPortStatus() int32`

GetReceiverPortStatus returns the ReceiverPortStatus field if non-nil, zero value otherwise.

### GetReceiverPortStatusOk

`func (o *HotspotRadiusSettingResOpenApiVO) GetReceiverPortStatusOk() (*int32, bool)`

GetReceiverPortStatusOk returns a tuple with the ReceiverPortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverPortStatus

`func (o *HotspotRadiusSettingResOpenApiVO) SetReceiverPortStatus(v int32)`

SetReceiverPortStatus sets ReceiverPortStatus field to given value.

### HasReceiverPortStatus

`func (o *HotspotRadiusSettingResOpenApiVO) HasReceiverPortStatus() bool`

HasReceiverPortStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


