# ApConfigResultPortSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApPortConfigResultSettings** | Pointer to [**ApPortConfigResultSettingsVO**](ApPortConfigResultSettingsVO.md) |  | [optional] 
**ErrorCode** | Pointer to **int32** | error code. | [optional] 
**LanPort** | Pointer to **string** | lanPort | [optional] 
**Msg** | Pointer to **string** | error msg | [optional] 

## Methods

### NewApConfigResultPortSettingVO

`func NewApConfigResultPortSettingVO() *ApConfigResultPortSettingVO`

NewApConfigResultPortSettingVO instantiates a new ApConfigResultPortSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApConfigResultPortSettingVOWithDefaults

`func NewApConfigResultPortSettingVOWithDefaults() *ApConfigResultPortSettingVO`

NewApConfigResultPortSettingVOWithDefaults instantiates a new ApConfigResultPortSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApPortConfigResultSettings

`func (o *ApConfigResultPortSettingVO) GetApPortConfigResultSettings() ApPortConfigResultSettingsVO`

GetApPortConfigResultSettings returns the ApPortConfigResultSettings field if non-nil, zero value otherwise.

### GetApPortConfigResultSettingsOk

`func (o *ApConfigResultPortSettingVO) GetApPortConfigResultSettingsOk() (*ApPortConfigResultSettingsVO, bool)`

GetApPortConfigResultSettingsOk returns a tuple with the ApPortConfigResultSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApPortConfigResultSettings

`func (o *ApConfigResultPortSettingVO) SetApPortConfigResultSettings(v ApPortConfigResultSettingsVO)`

SetApPortConfigResultSettings sets ApPortConfigResultSettings field to given value.

### HasApPortConfigResultSettings

`func (o *ApConfigResultPortSettingVO) HasApPortConfigResultSettings() bool`

HasApPortConfigResultSettings returns a boolean if a field has been set.

### GetErrorCode

`func (o *ApConfigResultPortSettingVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ApConfigResultPortSettingVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ApConfigResultPortSettingVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ApConfigResultPortSettingVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetLanPort

`func (o *ApConfigResultPortSettingVO) GetLanPort() string`

GetLanPort returns the LanPort field if non-nil, zero value otherwise.

### GetLanPortOk

`func (o *ApConfigResultPortSettingVO) GetLanPortOk() (*string, bool)`

GetLanPortOk returns a tuple with the LanPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanPort

`func (o *ApConfigResultPortSettingVO) SetLanPort(v string)`

SetLanPort sets LanPort field to given value.

### HasLanPort

`func (o *ApConfigResultPortSettingVO) HasLanPort() bool`

HasLanPort returns a boolean if a field has been set.

### GetMsg

`func (o *ApConfigResultPortSettingVO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ApConfigResultPortSettingVO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ApConfigResultPortSettingVO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ApConfigResultPortSettingVO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


