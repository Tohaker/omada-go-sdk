# VoipEmergencyNumberSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmergencyNumberEnable** | **bool** | Whether to enable the emergency number. | 
**EmergencyNumbers** | Pointer to **[]string** | Emergency number list. | [optional] 
**NoOperationTime** | Pointer to **int32** | No operation time should be within the range of 2-8s. | [optional] 

## Methods

### NewVoipEmergencyNumberSetting

`func NewVoipEmergencyNumberSetting(emergencyNumberEnable bool, ) *VoipEmergencyNumberSetting`

NewVoipEmergencyNumberSetting instantiates a new VoipEmergencyNumberSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoipEmergencyNumberSettingWithDefaults

`func NewVoipEmergencyNumberSettingWithDefaults() *VoipEmergencyNumberSetting`

NewVoipEmergencyNumberSettingWithDefaults instantiates a new VoipEmergencyNumberSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmergencyNumberEnable

`func (o *VoipEmergencyNumberSetting) GetEmergencyNumberEnable() bool`

GetEmergencyNumberEnable returns the EmergencyNumberEnable field if non-nil, zero value otherwise.

### GetEmergencyNumberEnableOk

`func (o *VoipEmergencyNumberSetting) GetEmergencyNumberEnableOk() (*bool, bool)`

GetEmergencyNumberEnableOk returns a tuple with the EmergencyNumberEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyNumberEnable

`func (o *VoipEmergencyNumberSetting) SetEmergencyNumberEnable(v bool)`

SetEmergencyNumberEnable sets EmergencyNumberEnable field to given value.


### GetEmergencyNumbers

`func (o *VoipEmergencyNumberSetting) GetEmergencyNumbers() []string`

GetEmergencyNumbers returns the EmergencyNumbers field if non-nil, zero value otherwise.

### GetEmergencyNumbersOk

`func (o *VoipEmergencyNumberSetting) GetEmergencyNumbersOk() (*[]string, bool)`

GetEmergencyNumbersOk returns a tuple with the EmergencyNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyNumbers

`func (o *VoipEmergencyNumberSetting) SetEmergencyNumbers(v []string)`

SetEmergencyNumbers sets EmergencyNumbers field to given value.

### HasEmergencyNumbers

`func (o *VoipEmergencyNumberSetting) HasEmergencyNumbers() bool`

HasEmergencyNumbers returns a boolean if a field has been set.

### GetNoOperationTime

`func (o *VoipEmergencyNumberSetting) GetNoOperationTime() int32`

GetNoOperationTime returns the NoOperationTime field if non-nil, zero value otherwise.

### GetNoOperationTimeOk

`func (o *VoipEmergencyNumberSetting) GetNoOperationTimeOk() (*int32, bool)`

GetNoOperationTimeOk returns a tuple with the NoOperationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOperationTime

`func (o *VoipEmergencyNumberSetting) SetNoOperationTime(v int32)`

SetNoOperationTime sets NoOperationTime field to given value.

### HasNoOperationTime

`func (o *VoipEmergencyNumberSetting) HasNoOperationTime() bool`

HasNoOperationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


