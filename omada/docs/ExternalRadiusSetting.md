# ExternalRadiusSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | **int32** | RADIUS auth mode, should be a value as follows: 1: PAP; 2: CHAP | 
**DisconnectReq** | Pointer to **bool** | Whether to support disconnect messages. | [optional] 
**ExternalUrl** | Pointer to **string** | External URL, required when [portalCustom] is 2 | [optional] 
**ExternalUrlScheme** | Pointer to **string** | External URL scheme, required when [portalCustom] is 2, value could be &#39;http&#39; or &#39;https&#39;. | [optional] 
**NasId** | **string** | RADIUS Attribute: NasID, should contain 1 to 64 characters. | 
**PortalCustom** | **int32** | Portal customization, should be a value as follows: 1: Local Web Portal; 2: External Web Portal. | 
**RadiusProfileId** | **string** | RADIUS profile ID. | 
**ReceiverPort** | Pointer to **int32** | Port for listening to disconnect messages, should be within the range of 1–65535. | [optional] 

## Methods

### NewExternalRadiusSetting

`func NewExternalRadiusSetting(authMode int32, nasId string, portalCustom int32, radiusProfileId string, ) *ExternalRadiusSetting`

NewExternalRadiusSetting instantiates a new ExternalRadiusSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalRadiusSettingWithDefaults

`func NewExternalRadiusSettingWithDefaults() *ExternalRadiusSetting`

NewExternalRadiusSettingWithDefaults instantiates a new ExternalRadiusSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *ExternalRadiusSetting) GetAuthMode() int32`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *ExternalRadiusSetting) GetAuthModeOk() (*int32, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *ExternalRadiusSetting) SetAuthMode(v int32)`

SetAuthMode sets AuthMode field to given value.


### GetDisconnectReq

`func (o *ExternalRadiusSetting) GetDisconnectReq() bool`

GetDisconnectReq returns the DisconnectReq field if non-nil, zero value otherwise.

### GetDisconnectReqOk

`func (o *ExternalRadiusSetting) GetDisconnectReqOk() (*bool, bool)`

GetDisconnectReqOk returns a tuple with the DisconnectReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectReq

`func (o *ExternalRadiusSetting) SetDisconnectReq(v bool)`

SetDisconnectReq sets DisconnectReq field to given value.

### HasDisconnectReq

`func (o *ExternalRadiusSetting) HasDisconnectReq() bool`

HasDisconnectReq returns a boolean if a field has been set.

### GetExternalUrl

`func (o *ExternalRadiusSetting) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *ExternalRadiusSetting) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *ExternalRadiusSetting) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *ExternalRadiusSetting) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetExternalUrlScheme

`func (o *ExternalRadiusSetting) GetExternalUrlScheme() string`

GetExternalUrlScheme returns the ExternalUrlScheme field if non-nil, zero value otherwise.

### GetExternalUrlSchemeOk

`func (o *ExternalRadiusSetting) GetExternalUrlSchemeOk() (*string, bool)`

GetExternalUrlSchemeOk returns a tuple with the ExternalUrlScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrlScheme

`func (o *ExternalRadiusSetting) SetExternalUrlScheme(v string)`

SetExternalUrlScheme sets ExternalUrlScheme field to given value.

### HasExternalUrlScheme

`func (o *ExternalRadiusSetting) HasExternalUrlScheme() bool`

HasExternalUrlScheme returns a boolean if a field has been set.

### GetNasId

`func (o *ExternalRadiusSetting) GetNasId() string`

GetNasId returns the NasId field if non-nil, zero value otherwise.

### GetNasIdOk

`func (o *ExternalRadiusSetting) GetNasIdOk() (*string, bool)`

GetNasIdOk returns a tuple with the NasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasId

`func (o *ExternalRadiusSetting) SetNasId(v string)`

SetNasId sets NasId field to given value.


### GetPortalCustom

`func (o *ExternalRadiusSetting) GetPortalCustom() int32`

GetPortalCustom returns the PortalCustom field if non-nil, zero value otherwise.

### GetPortalCustomOk

`func (o *ExternalRadiusSetting) GetPortalCustomOk() (*int32, bool)`

GetPortalCustomOk returns a tuple with the PortalCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalCustom

`func (o *ExternalRadiusSetting) SetPortalCustom(v int32)`

SetPortalCustom sets PortalCustom field to given value.


### GetRadiusProfileId

`func (o *ExternalRadiusSetting) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *ExternalRadiusSetting) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *ExternalRadiusSetting) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.


### GetReceiverPort

`func (o *ExternalRadiusSetting) GetReceiverPort() int32`

GetReceiverPort returns the ReceiverPort field if non-nil, zero value otherwise.

### GetReceiverPortOk

`func (o *ExternalRadiusSetting) GetReceiverPortOk() (*int32, bool)`

GetReceiverPortOk returns a tuple with the ReceiverPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverPort

`func (o *ExternalRadiusSetting) SetReceiverPort(v int32)`

SetReceiverPort sets ReceiverPort field to given value.

### HasReceiverPort

`func (o *ExternalRadiusSetting) HasReceiverPort() bool`

HasReceiverPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


