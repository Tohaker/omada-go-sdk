# ExternalRadiusSettingResOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **int32** | RADIUS auth mode, should be a value as follows: 1: PAP; 2: CHAP | [optional] 
**DisconnectReq** | Pointer to **bool** | Whether to support disconnect messages. Only for Omada Local Controller | [optional] 
**ExternalUrl** | Pointer to **string** | External URL, required when [portalCustom] is 2 | [optional] 
**ExternalUrlScheme** | Pointer to **string** | External URL scheme, required when [portalCustom] is 2, value could be &#39;http&#39; or &#39;https&#39;. | [optional] 
**NasId** | Pointer to **string** | RADIUS Attribute: NasID, should contain 1 to 64 characters. | [optional] 
**PortalCustom** | Pointer to **int32** | Portal customization, should be a value as follows: 1: Local Web Portal; 2: External Web Portal. | [optional] 
**RadiusProfileId** | Pointer to **string** | This field represents radius profile ID. Radius profile can be created using &#39;Create a new Radius profile&#39; (&#39;Create a new Radius profile template&#39;) interface, and radius profile ID can be obtained from &#39;Get Radius profile list&#39; (&#39;Get Radius profile template list&#39;) interface | [optional] 
**ReceiverPort** | Pointer to **int32** | Port for listening to disconnect messages, should be within the range of 1–65535.Only for Omada Local Controller | [optional] 
**ReceiverPortStatus** | Pointer to **int32** | Port binding status, should be a value as follow: 1: Disconnect Requests port status running, 2: Disconnect Requests port status disable. Only for Omada Local Controller | [optional] 

## Methods

### NewExternalRadiusSettingResOpenApiVO

`func NewExternalRadiusSettingResOpenApiVO() *ExternalRadiusSettingResOpenApiVO`

NewExternalRadiusSettingResOpenApiVO instantiates a new ExternalRadiusSettingResOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalRadiusSettingResOpenApiVOWithDefaults

`func NewExternalRadiusSettingResOpenApiVOWithDefaults() *ExternalRadiusSettingResOpenApiVO`

NewExternalRadiusSettingResOpenApiVOWithDefaults instantiates a new ExternalRadiusSettingResOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *ExternalRadiusSettingResOpenApiVO) GetAuthMode() int32`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *ExternalRadiusSettingResOpenApiVO) GetAuthModeOk() (*int32, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *ExternalRadiusSettingResOpenApiVO) SetAuthMode(v int32)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *ExternalRadiusSettingResOpenApiVO) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetDisconnectReq

`func (o *ExternalRadiusSettingResOpenApiVO) GetDisconnectReq() bool`

GetDisconnectReq returns the DisconnectReq field if non-nil, zero value otherwise.

### GetDisconnectReqOk

`func (o *ExternalRadiusSettingResOpenApiVO) GetDisconnectReqOk() (*bool, bool)`

GetDisconnectReqOk returns a tuple with the DisconnectReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectReq

`func (o *ExternalRadiusSettingResOpenApiVO) SetDisconnectReq(v bool)`

SetDisconnectReq sets DisconnectReq field to given value.

### HasDisconnectReq

`func (o *ExternalRadiusSettingResOpenApiVO) HasDisconnectReq() bool`

HasDisconnectReq returns a boolean if a field has been set.

### GetExternalUrl

`func (o *ExternalRadiusSettingResOpenApiVO) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *ExternalRadiusSettingResOpenApiVO) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *ExternalRadiusSettingResOpenApiVO) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *ExternalRadiusSettingResOpenApiVO) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetExternalUrlScheme

`func (o *ExternalRadiusSettingResOpenApiVO) GetExternalUrlScheme() string`

GetExternalUrlScheme returns the ExternalUrlScheme field if non-nil, zero value otherwise.

### GetExternalUrlSchemeOk

`func (o *ExternalRadiusSettingResOpenApiVO) GetExternalUrlSchemeOk() (*string, bool)`

GetExternalUrlSchemeOk returns a tuple with the ExternalUrlScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrlScheme

`func (o *ExternalRadiusSettingResOpenApiVO) SetExternalUrlScheme(v string)`

SetExternalUrlScheme sets ExternalUrlScheme field to given value.

### HasExternalUrlScheme

`func (o *ExternalRadiusSettingResOpenApiVO) HasExternalUrlScheme() bool`

HasExternalUrlScheme returns a boolean if a field has been set.

### GetNasId

`func (o *ExternalRadiusSettingResOpenApiVO) GetNasId() string`

GetNasId returns the NasId field if non-nil, zero value otherwise.

### GetNasIdOk

`func (o *ExternalRadiusSettingResOpenApiVO) GetNasIdOk() (*string, bool)`

GetNasIdOk returns a tuple with the NasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasId

`func (o *ExternalRadiusSettingResOpenApiVO) SetNasId(v string)`

SetNasId sets NasId field to given value.

### HasNasId

`func (o *ExternalRadiusSettingResOpenApiVO) HasNasId() bool`

HasNasId returns a boolean if a field has been set.

### GetPortalCustom

`func (o *ExternalRadiusSettingResOpenApiVO) GetPortalCustom() int32`

GetPortalCustom returns the PortalCustom field if non-nil, zero value otherwise.

### GetPortalCustomOk

`func (o *ExternalRadiusSettingResOpenApiVO) GetPortalCustomOk() (*int32, bool)`

GetPortalCustomOk returns a tuple with the PortalCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalCustom

`func (o *ExternalRadiusSettingResOpenApiVO) SetPortalCustom(v int32)`

SetPortalCustom sets PortalCustom field to given value.

### HasPortalCustom

`func (o *ExternalRadiusSettingResOpenApiVO) HasPortalCustom() bool`

HasPortalCustom returns a boolean if a field has been set.

### GetRadiusProfileId

`func (o *ExternalRadiusSettingResOpenApiVO) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *ExternalRadiusSettingResOpenApiVO) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *ExternalRadiusSettingResOpenApiVO) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.

### HasRadiusProfileId

`func (o *ExternalRadiusSettingResOpenApiVO) HasRadiusProfileId() bool`

HasRadiusProfileId returns a boolean if a field has been set.

### GetReceiverPort

`func (o *ExternalRadiusSettingResOpenApiVO) GetReceiverPort() int32`

GetReceiverPort returns the ReceiverPort field if non-nil, zero value otherwise.

### GetReceiverPortOk

`func (o *ExternalRadiusSettingResOpenApiVO) GetReceiverPortOk() (*int32, bool)`

GetReceiverPortOk returns a tuple with the ReceiverPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverPort

`func (o *ExternalRadiusSettingResOpenApiVO) SetReceiverPort(v int32)`

SetReceiverPort sets ReceiverPort field to given value.

### HasReceiverPort

`func (o *ExternalRadiusSettingResOpenApiVO) HasReceiverPort() bool`

HasReceiverPort returns a boolean if a field has been set.

### GetReceiverPortStatus

`func (o *ExternalRadiusSettingResOpenApiVO) GetReceiverPortStatus() int32`

GetReceiverPortStatus returns the ReceiverPortStatus field if non-nil, zero value otherwise.

### GetReceiverPortStatusOk

`func (o *ExternalRadiusSettingResOpenApiVO) GetReceiverPortStatusOk() (*int32, bool)`

GetReceiverPortStatusOk returns a tuple with the ReceiverPortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverPortStatus

`func (o *ExternalRadiusSettingResOpenApiVO) SetReceiverPortStatus(v int32)`

SetReceiverPortStatus sets ReceiverPortStatus field to given value.

### HasReceiverPortStatus

`func (o *ExternalRadiusSettingResOpenApiVO) HasReceiverPortStatus() bool`

HasReceiverPortStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


