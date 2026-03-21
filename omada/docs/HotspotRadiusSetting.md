# HotspotRadiusSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | **int32** | RADIUS auth mode, should be a value as follows: 1: PAP; 2: CHAP | 
**AuthTimeout** | [**AuthTimeoutSetting**](AuthTimeoutSetting.md) |  | 
**DisconnectReq** | Pointer to **bool** | Whether to support disconnect messages. | [optional] 
**NasId** | **string** | RADIUS Attribute: NasID, should contain 1 to 64 characters. | 
**RadiusProfileId** | **string** | RADIUS profile ID. | 
**ReceiverPort** | Pointer to **int32** | Port for listening to disconnect messages, should be within the range of 1–65535. | [optional] 

## Methods

### NewHotspotRadiusSetting

`func NewHotspotRadiusSetting(authMode int32, authTimeout AuthTimeoutSetting, nasId string, radiusProfileId string, ) *HotspotRadiusSetting`

NewHotspotRadiusSetting instantiates a new HotspotRadiusSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHotspotRadiusSettingWithDefaults

`func NewHotspotRadiusSettingWithDefaults() *HotspotRadiusSetting`

NewHotspotRadiusSettingWithDefaults instantiates a new HotspotRadiusSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *HotspotRadiusSetting) GetAuthMode() int32`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *HotspotRadiusSetting) GetAuthModeOk() (*int32, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *HotspotRadiusSetting) SetAuthMode(v int32)`

SetAuthMode sets AuthMode field to given value.


### GetAuthTimeout

`func (o *HotspotRadiusSetting) GetAuthTimeout() AuthTimeoutSetting`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *HotspotRadiusSetting) GetAuthTimeoutOk() (*AuthTimeoutSetting, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *HotspotRadiusSetting) SetAuthTimeout(v AuthTimeoutSetting)`

SetAuthTimeout sets AuthTimeout field to given value.


### GetDisconnectReq

`func (o *HotspotRadiusSetting) GetDisconnectReq() bool`

GetDisconnectReq returns the DisconnectReq field if non-nil, zero value otherwise.

### GetDisconnectReqOk

`func (o *HotspotRadiusSetting) GetDisconnectReqOk() (*bool, bool)`

GetDisconnectReqOk returns a tuple with the DisconnectReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectReq

`func (o *HotspotRadiusSetting) SetDisconnectReq(v bool)`

SetDisconnectReq sets DisconnectReq field to given value.

### HasDisconnectReq

`func (o *HotspotRadiusSetting) HasDisconnectReq() bool`

HasDisconnectReq returns a boolean if a field has been set.

### GetNasId

`func (o *HotspotRadiusSetting) GetNasId() string`

GetNasId returns the NasId field if non-nil, zero value otherwise.

### GetNasIdOk

`func (o *HotspotRadiusSetting) GetNasIdOk() (*string, bool)`

GetNasIdOk returns a tuple with the NasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasId

`func (o *HotspotRadiusSetting) SetNasId(v string)`

SetNasId sets NasId field to given value.


### GetRadiusProfileId

`func (o *HotspotRadiusSetting) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *HotspotRadiusSetting) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *HotspotRadiusSetting) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.


### GetReceiverPort

`func (o *HotspotRadiusSetting) GetReceiverPort() int32`

GetReceiverPort returns the ReceiverPort field if non-nil, zero value otherwise.

### GetReceiverPortOk

`func (o *HotspotRadiusSetting) GetReceiverPortOk() (*int32, bool)`

GetReceiverPortOk returns a tuple with the ReceiverPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverPort

`func (o *HotspotRadiusSetting) SetReceiverPort(v int32)`

SetReceiverPort sets ReceiverPort field to given value.

### HasReceiverPort

`func (o *HotspotRadiusSetting) HasReceiverPort() bool`

HasReceiverPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


