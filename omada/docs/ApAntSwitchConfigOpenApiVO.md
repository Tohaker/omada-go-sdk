# ApAntSwitchConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntMode** | Pointer to **int32** | antenna mode. 0: Auto; 1: Built-in; 2: Omni; 3:Custom | [optional] 
**AntSetting2g** | Pointer to [**AntSettingVO**](AntSettingVO.md) |  | [optional] 
**AntSetting5g** | Pointer to [**AntSettingVO**](AntSettingVO.md) |  | [optional] 
**AntSetting5g2** | Pointer to [**AntSettingVO**](AntSettingVO.md) |  | [optional] 
**AntSetting6g** | Pointer to [**AntSettingVO**](AntSettingVO.md) |  | [optional] 
**OmniAntState** | Pointer to **int32** | omni antenna install state. 0: Not Installed; 1: Abnormal installation status; 2: Normal. | [optional] 

## Methods

### NewApAntSwitchConfigOpenApiVO

`func NewApAntSwitchConfigOpenApiVO() *ApAntSwitchConfigOpenApiVO`

NewApAntSwitchConfigOpenApiVO instantiates a new ApAntSwitchConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApAntSwitchConfigOpenApiVOWithDefaults

`func NewApAntSwitchConfigOpenApiVOWithDefaults() *ApAntSwitchConfigOpenApiVO`

NewApAntSwitchConfigOpenApiVOWithDefaults instantiates a new ApAntSwitchConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntMode

`func (o *ApAntSwitchConfigOpenApiVO) GetAntMode() int32`

GetAntMode returns the AntMode field if non-nil, zero value otherwise.

### GetAntModeOk

`func (o *ApAntSwitchConfigOpenApiVO) GetAntModeOk() (*int32, bool)`

GetAntModeOk returns a tuple with the AntMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntMode

`func (o *ApAntSwitchConfigOpenApiVO) SetAntMode(v int32)`

SetAntMode sets AntMode field to given value.

### HasAntMode

`func (o *ApAntSwitchConfigOpenApiVO) HasAntMode() bool`

HasAntMode returns a boolean if a field has been set.

### GetAntSetting2g

`func (o *ApAntSwitchConfigOpenApiVO) GetAntSetting2g() AntSettingVO`

GetAntSetting2g returns the AntSetting2g field if non-nil, zero value otherwise.

### GetAntSetting2gOk

`func (o *ApAntSwitchConfigOpenApiVO) GetAntSetting2gOk() (*AntSettingVO, bool)`

GetAntSetting2gOk returns a tuple with the AntSetting2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntSetting2g

`func (o *ApAntSwitchConfigOpenApiVO) SetAntSetting2g(v AntSettingVO)`

SetAntSetting2g sets AntSetting2g field to given value.

### HasAntSetting2g

`func (o *ApAntSwitchConfigOpenApiVO) HasAntSetting2g() bool`

HasAntSetting2g returns a boolean if a field has been set.

### GetAntSetting5g

`func (o *ApAntSwitchConfigOpenApiVO) GetAntSetting5g() AntSettingVO`

GetAntSetting5g returns the AntSetting5g field if non-nil, zero value otherwise.

### GetAntSetting5gOk

`func (o *ApAntSwitchConfigOpenApiVO) GetAntSetting5gOk() (*AntSettingVO, bool)`

GetAntSetting5gOk returns a tuple with the AntSetting5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntSetting5g

`func (o *ApAntSwitchConfigOpenApiVO) SetAntSetting5g(v AntSettingVO)`

SetAntSetting5g sets AntSetting5g field to given value.

### HasAntSetting5g

`func (o *ApAntSwitchConfigOpenApiVO) HasAntSetting5g() bool`

HasAntSetting5g returns a boolean if a field has been set.

### GetAntSetting5g2

`func (o *ApAntSwitchConfigOpenApiVO) GetAntSetting5g2() AntSettingVO`

GetAntSetting5g2 returns the AntSetting5g2 field if non-nil, zero value otherwise.

### GetAntSetting5g2Ok

`func (o *ApAntSwitchConfigOpenApiVO) GetAntSetting5g2Ok() (*AntSettingVO, bool)`

GetAntSetting5g2Ok returns a tuple with the AntSetting5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntSetting5g2

`func (o *ApAntSwitchConfigOpenApiVO) SetAntSetting5g2(v AntSettingVO)`

SetAntSetting5g2 sets AntSetting5g2 field to given value.

### HasAntSetting5g2

`func (o *ApAntSwitchConfigOpenApiVO) HasAntSetting5g2() bool`

HasAntSetting5g2 returns a boolean if a field has been set.

### GetAntSetting6g

`func (o *ApAntSwitchConfigOpenApiVO) GetAntSetting6g() AntSettingVO`

GetAntSetting6g returns the AntSetting6g field if non-nil, zero value otherwise.

### GetAntSetting6gOk

`func (o *ApAntSwitchConfigOpenApiVO) GetAntSetting6gOk() (*AntSettingVO, bool)`

GetAntSetting6gOk returns a tuple with the AntSetting6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntSetting6g

`func (o *ApAntSwitchConfigOpenApiVO) SetAntSetting6g(v AntSettingVO)`

SetAntSetting6g sets AntSetting6g field to given value.

### HasAntSetting6g

`func (o *ApAntSwitchConfigOpenApiVO) HasAntSetting6g() bool`

HasAntSetting6g returns a boolean if a field has been set.

### GetOmniAntState

`func (o *ApAntSwitchConfigOpenApiVO) GetOmniAntState() int32`

GetOmniAntState returns the OmniAntState field if non-nil, zero value otherwise.

### GetOmniAntStateOk

`func (o *ApAntSwitchConfigOpenApiVO) GetOmniAntStateOk() (*int32, bool)`

GetOmniAntStateOk returns a tuple with the OmniAntState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmniAntState

`func (o *ApAntSwitchConfigOpenApiVO) SetOmniAntState(v int32)`

SetOmniAntState sets OmniAntState field to given value.

### HasOmniAntState

`func (o *ApAntSwitchConfigOpenApiVO) HasOmniAntState() bool`

HasOmniAntState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


