# ApConfigResultRadioSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApRadioConfigResultSetting** | Pointer to [**ApRadioConfigResultSettingsVO**](ApRadioConfigResultSettingsVO.md) |  | [optional] 
**ErrorCode** | Pointer to **int32** | error code. | [optional] 
**Msg** | Pointer to **string** | error msg | [optional] 

## Methods

### NewApConfigResultRadioSettingVO

`func NewApConfigResultRadioSettingVO() *ApConfigResultRadioSettingVO`

NewApConfigResultRadioSettingVO instantiates a new ApConfigResultRadioSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApConfigResultRadioSettingVOWithDefaults

`func NewApConfigResultRadioSettingVOWithDefaults() *ApConfigResultRadioSettingVO`

NewApConfigResultRadioSettingVOWithDefaults instantiates a new ApConfigResultRadioSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApRadioConfigResultSetting

`func (o *ApConfigResultRadioSettingVO) GetApRadioConfigResultSetting() ApRadioConfigResultSettingsVO`

GetApRadioConfigResultSetting returns the ApRadioConfigResultSetting field if non-nil, zero value otherwise.

### GetApRadioConfigResultSettingOk

`func (o *ApConfigResultRadioSettingVO) GetApRadioConfigResultSettingOk() (*ApRadioConfigResultSettingsVO, bool)`

GetApRadioConfigResultSettingOk returns a tuple with the ApRadioConfigResultSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApRadioConfigResultSetting

`func (o *ApConfigResultRadioSettingVO) SetApRadioConfigResultSetting(v ApRadioConfigResultSettingsVO)`

SetApRadioConfigResultSetting sets ApRadioConfigResultSetting field to given value.

### HasApRadioConfigResultSetting

`func (o *ApConfigResultRadioSettingVO) HasApRadioConfigResultSetting() bool`

HasApRadioConfigResultSetting returns a boolean if a field has been set.

### GetErrorCode

`func (o *ApConfigResultRadioSettingVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ApConfigResultRadioSettingVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ApConfigResultRadioSettingVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ApConfigResultRadioSettingVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *ApConfigResultRadioSettingVO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ApConfigResultRadioSettingVO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ApConfigResultRadioSettingVO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ApConfigResultRadioSettingVO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


