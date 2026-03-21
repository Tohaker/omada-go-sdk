# BatchUpdateApVlanOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMacList** | **[]string** | AP mac list | 
**MvlanSetting** | Pointer to [**ApMvlanSettingOpenApiVO**](ApMvlanSettingOpenApiVO.md) |  | [optional] 
**VoipVlanSetting** | Pointer to [**ApVoipVlanSettingOpenApiVO**](ApVoipVlanSettingOpenApiVO.md) |  | [optional] 

## Methods

### NewBatchUpdateApVlanOpenApiVO

`func NewBatchUpdateApVlanOpenApiVO(apMacList []string, ) *BatchUpdateApVlanOpenApiVO`

NewBatchUpdateApVlanOpenApiVO instantiates a new BatchUpdateApVlanOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateApVlanOpenApiVOWithDefaults

`func NewBatchUpdateApVlanOpenApiVOWithDefaults() *BatchUpdateApVlanOpenApiVO`

NewBatchUpdateApVlanOpenApiVOWithDefaults instantiates a new BatchUpdateApVlanOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMacList

`func (o *BatchUpdateApVlanOpenApiVO) GetApMacList() []string`

GetApMacList returns the ApMacList field if non-nil, zero value otherwise.

### GetApMacListOk

`func (o *BatchUpdateApVlanOpenApiVO) GetApMacListOk() (*[]string, bool)`

GetApMacListOk returns a tuple with the ApMacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMacList

`func (o *BatchUpdateApVlanOpenApiVO) SetApMacList(v []string)`

SetApMacList sets ApMacList field to given value.


### GetMvlanSetting

`func (o *BatchUpdateApVlanOpenApiVO) GetMvlanSetting() ApMvlanSettingOpenApiVO`

GetMvlanSetting returns the MvlanSetting field if non-nil, zero value otherwise.

### GetMvlanSettingOk

`func (o *BatchUpdateApVlanOpenApiVO) GetMvlanSettingOk() (*ApMvlanSettingOpenApiVO, bool)`

GetMvlanSettingOk returns a tuple with the MvlanSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvlanSetting

`func (o *BatchUpdateApVlanOpenApiVO) SetMvlanSetting(v ApMvlanSettingOpenApiVO)`

SetMvlanSetting sets MvlanSetting field to given value.

### HasMvlanSetting

`func (o *BatchUpdateApVlanOpenApiVO) HasMvlanSetting() bool`

HasMvlanSetting returns a boolean if a field has been set.

### GetVoipVlanSetting

`func (o *BatchUpdateApVlanOpenApiVO) GetVoipVlanSetting() ApVoipVlanSettingOpenApiVO`

GetVoipVlanSetting returns the VoipVlanSetting field if non-nil, zero value otherwise.

### GetVoipVlanSettingOk

`func (o *BatchUpdateApVlanOpenApiVO) GetVoipVlanSettingOk() (*ApVoipVlanSettingOpenApiVO, bool)`

GetVoipVlanSettingOk returns a tuple with the VoipVlanSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipVlanSetting

`func (o *BatchUpdateApVlanOpenApiVO) SetVoipVlanSetting(v ApVoipVlanSettingOpenApiVO)`

SetVoipVlanSetting sets VoipVlanSetting field to given value.

### HasVoipVlanSetting

`func (o *BatchUpdateApVlanOpenApiVO) HasVoipVlanSetting() bool`

HasVoipVlanSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


