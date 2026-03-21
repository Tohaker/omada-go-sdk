# BatchUpdateApAntSwitchOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntMode** | Pointer to **int32** |  | [optional] 
**AntSetting2g** | Pointer to [**AntSwitchRadioConfig**](AntSwitchRadioConfig.md) |  | [optional] 
**AntSetting5g** | Pointer to [**AntSwitchRadioConfig**](AntSwitchRadioConfig.md) |  | [optional] 
**AntSetting5g2** | Pointer to [**AntSwitchRadioConfig**](AntSwitchRadioConfig.md) |  | [optional] 
**AntSetting6g** | Pointer to [**AntSwitchRadioConfig**](AntSwitchRadioConfig.md) |  | [optional] 
**ApMacList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBatchUpdateApAntSwitchOpenApiVO

`func NewBatchUpdateApAntSwitchOpenApiVO() *BatchUpdateApAntSwitchOpenApiVO`

NewBatchUpdateApAntSwitchOpenApiVO instantiates a new BatchUpdateApAntSwitchOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateApAntSwitchOpenApiVOWithDefaults

`func NewBatchUpdateApAntSwitchOpenApiVOWithDefaults() *BatchUpdateApAntSwitchOpenApiVO`

NewBatchUpdateApAntSwitchOpenApiVOWithDefaults instantiates a new BatchUpdateApAntSwitchOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntMode

`func (o *BatchUpdateApAntSwitchOpenApiVO) GetAntMode() int32`

GetAntMode returns the AntMode field if non-nil, zero value otherwise.

### GetAntModeOk

`func (o *BatchUpdateApAntSwitchOpenApiVO) GetAntModeOk() (*int32, bool)`

GetAntModeOk returns a tuple with the AntMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntMode

`func (o *BatchUpdateApAntSwitchOpenApiVO) SetAntMode(v int32)`

SetAntMode sets AntMode field to given value.

### HasAntMode

`func (o *BatchUpdateApAntSwitchOpenApiVO) HasAntMode() bool`

HasAntMode returns a boolean if a field has been set.

### GetAntSetting2g

`func (o *BatchUpdateApAntSwitchOpenApiVO) GetAntSetting2g() AntSwitchRadioConfig`

GetAntSetting2g returns the AntSetting2g field if non-nil, zero value otherwise.

### GetAntSetting2gOk

`func (o *BatchUpdateApAntSwitchOpenApiVO) GetAntSetting2gOk() (*AntSwitchRadioConfig, bool)`

GetAntSetting2gOk returns a tuple with the AntSetting2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntSetting2g

`func (o *BatchUpdateApAntSwitchOpenApiVO) SetAntSetting2g(v AntSwitchRadioConfig)`

SetAntSetting2g sets AntSetting2g field to given value.

### HasAntSetting2g

`func (o *BatchUpdateApAntSwitchOpenApiVO) HasAntSetting2g() bool`

HasAntSetting2g returns a boolean if a field has been set.

### GetAntSetting5g

`func (o *BatchUpdateApAntSwitchOpenApiVO) GetAntSetting5g() AntSwitchRadioConfig`

GetAntSetting5g returns the AntSetting5g field if non-nil, zero value otherwise.

### GetAntSetting5gOk

`func (o *BatchUpdateApAntSwitchOpenApiVO) GetAntSetting5gOk() (*AntSwitchRadioConfig, bool)`

GetAntSetting5gOk returns a tuple with the AntSetting5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntSetting5g

`func (o *BatchUpdateApAntSwitchOpenApiVO) SetAntSetting5g(v AntSwitchRadioConfig)`

SetAntSetting5g sets AntSetting5g field to given value.

### HasAntSetting5g

`func (o *BatchUpdateApAntSwitchOpenApiVO) HasAntSetting5g() bool`

HasAntSetting5g returns a boolean if a field has been set.

### GetAntSetting5g2

`func (o *BatchUpdateApAntSwitchOpenApiVO) GetAntSetting5g2() AntSwitchRadioConfig`

GetAntSetting5g2 returns the AntSetting5g2 field if non-nil, zero value otherwise.

### GetAntSetting5g2Ok

`func (o *BatchUpdateApAntSwitchOpenApiVO) GetAntSetting5g2Ok() (*AntSwitchRadioConfig, bool)`

GetAntSetting5g2Ok returns a tuple with the AntSetting5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntSetting5g2

`func (o *BatchUpdateApAntSwitchOpenApiVO) SetAntSetting5g2(v AntSwitchRadioConfig)`

SetAntSetting5g2 sets AntSetting5g2 field to given value.

### HasAntSetting5g2

`func (o *BatchUpdateApAntSwitchOpenApiVO) HasAntSetting5g2() bool`

HasAntSetting5g2 returns a boolean if a field has been set.

### GetAntSetting6g

`func (o *BatchUpdateApAntSwitchOpenApiVO) GetAntSetting6g() AntSwitchRadioConfig`

GetAntSetting6g returns the AntSetting6g field if non-nil, zero value otherwise.

### GetAntSetting6gOk

`func (o *BatchUpdateApAntSwitchOpenApiVO) GetAntSetting6gOk() (*AntSwitchRadioConfig, bool)`

GetAntSetting6gOk returns a tuple with the AntSetting6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntSetting6g

`func (o *BatchUpdateApAntSwitchOpenApiVO) SetAntSetting6g(v AntSwitchRadioConfig)`

SetAntSetting6g sets AntSetting6g field to given value.

### HasAntSetting6g

`func (o *BatchUpdateApAntSwitchOpenApiVO) HasAntSetting6g() bool`

HasAntSetting6g returns a boolean if a field has been set.

### GetApMacList

`func (o *BatchUpdateApAntSwitchOpenApiVO) GetApMacList() []string`

GetApMacList returns the ApMacList field if non-nil, zero value otherwise.

### GetApMacListOk

`func (o *BatchUpdateApAntSwitchOpenApiVO) GetApMacListOk() (*[]string, bool)`

GetApMacListOk returns a tuple with the ApMacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMacList

`func (o *BatchUpdateApAntSwitchOpenApiVO) SetApMacList(v []string)`

SetApMacList sets ApMacList field to given value.

### HasApMacList

`func (o *BatchUpdateApAntSwitchOpenApiVO) HasApMacList() bool`

HasApMacList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


