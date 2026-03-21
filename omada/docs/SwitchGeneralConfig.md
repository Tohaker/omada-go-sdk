# SwitchGeneralConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jumbo** | Pointer to **int32** | It should be within the range of 1518–9216 Bytes. | [optional] 
**LagHashAlg** | Pointer to **int32** | It should be a value as follows: 0: SRC MAC; 1: DST MAC; 2: SRC MAC + DST MAC; 3: SRC IP; 4: DST IP; 5: SRC IP + DST IP | [optional] 
**LedSetting** | Pointer to **int32** | Led setting should be a value as follows: 0:off; 1:on; 2:Use Site Settings | [optional] 
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**Name** | Pointer to **string** | Device name should contain 1 to 128 characters. | [optional] 
**Sdm** | Pointer to [**OswSdmTemplateVO**](OswSdmTemplateVO.md) |  | [optional] 
**TagIds** | Pointer to **[]string** | Tag IDs | [optional] 

## Methods

### NewSwitchGeneralConfig

`func NewSwitchGeneralConfig() *SwitchGeneralConfig`

NewSwitchGeneralConfig instantiates a new SwitchGeneralConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchGeneralConfigWithDefaults

`func NewSwitchGeneralConfigWithDefaults() *SwitchGeneralConfig`

NewSwitchGeneralConfigWithDefaults instantiates a new SwitchGeneralConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJumbo

`func (o *SwitchGeneralConfig) GetJumbo() int32`

GetJumbo returns the Jumbo field if non-nil, zero value otherwise.

### GetJumboOk

`func (o *SwitchGeneralConfig) GetJumboOk() (*int32, bool)`

GetJumboOk returns a tuple with the Jumbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumbo

`func (o *SwitchGeneralConfig) SetJumbo(v int32)`

SetJumbo sets Jumbo field to given value.

### HasJumbo

`func (o *SwitchGeneralConfig) HasJumbo() bool`

HasJumbo returns a boolean if a field has been set.

### GetLagHashAlg

`func (o *SwitchGeneralConfig) GetLagHashAlg() int32`

GetLagHashAlg returns the LagHashAlg field if non-nil, zero value otherwise.

### GetLagHashAlgOk

`func (o *SwitchGeneralConfig) GetLagHashAlgOk() (*int32, bool)`

GetLagHashAlgOk returns a tuple with the LagHashAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagHashAlg

`func (o *SwitchGeneralConfig) SetLagHashAlg(v int32)`

SetLagHashAlg sets LagHashAlg field to given value.

### HasLagHashAlg

`func (o *SwitchGeneralConfig) HasLagHashAlg() bool`

HasLagHashAlg returns a boolean if a field has been set.

### GetLedSetting

`func (o *SwitchGeneralConfig) GetLedSetting() int32`

GetLedSetting returns the LedSetting field if non-nil, zero value otherwise.

### GetLedSettingOk

`func (o *SwitchGeneralConfig) GetLedSettingOk() (*int32, bool)`

GetLedSettingOk returns a tuple with the LedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedSetting

`func (o *SwitchGeneralConfig) SetLedSetting(v int32)`

SetLedSetting sets LedSetting field to given value.

### HasLedSetting

`func (o *SwitchGeneralConfig) HasLedSetting() bool`

HasLedSetting returns a boolean if a field has been set.

### GetLocation

`func (o *SwitchGeneralConfig) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SwitchGeneralConfig) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SwitchGeneralConfig) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SwitchGeneralConfig) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *SwitchGeneralConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwitchGeneralConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwitchGeneralConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SwitchGeneralConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSdm

`func (o *SwitchGeneralConfig) GetSdm() OswSdmTemplateVO`

GetSdm returns the Sdm field if non-nil, zero value otherwise.

### GetSdmOk

`func (o *SwitchGeneralConfig) GetSdmOk() (*OswSdmTemplateVO, bool)`

GetSdmOk returns a tuple with the Sdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdm

`func (o *SwitchGeneralConfig) SetSdm(v OswSdmTemplateVO)`

SetSdm sets Sdm field to given value.

### HasSdm

`func (o *SwitchGeneralConfig) HasSdm() bool`

HasSdm returns a boolean if a field has been set.

### GetTagIds

`func (o *SwitchGeneralConfig) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *SwitchGeneralConfig) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *SwitchGeneralConfig) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *SwitchGeneralConfig) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


