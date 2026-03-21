# MacFilteringGeneralSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**GatewayDirectionEntity**](GatewayDirectionEntity.md) |  | [optional] 
**Enable** | **bool** | Enable of the MAC filtering general setting. | 
**FilterMode** | Pointer to **int32** | Filter mode should be a value as follows: 0: allow; 1: deny. | [optional] 

## Methods

### NewMacFilteringGeneralSetting

`func NewMacFilteringGeneralSetting(enable bool, ) *MacFilteringGeneralSetting`

NewMacFilteringGeneralSetting instantiates a new MacFilteringGeneralSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacFilteringGeneralSettingWithDefaults

`func NewMacFilteringGeneralSettingWithDefaults() *MacFilteringGeneralSetting`

NewMacFilteringGeneralSettingWithDefaults instantiates a new MacFilteringGeneralSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *MacFilteringGeneralSetting) GetDirection() GatewayDirectionEntity`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MacFilteringGeneralSetting) GetDirectionOk() (*GatewayDirectionEntity, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MacFilteringGeneralSetting) SetDirection(v GatewayDirectionEntity)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MacFilteringGeneralSetting) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetEnable

`func (o *MacFilteringGeneralSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *MacFilteringGeneralSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *MacFilteringGeneralSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetFilterMode

`func (o *MacFilteringGeneralSetting) GetFilterMode() int32`

GetFilterMode returns the FilterMode field if non-nil, zero value otherwise.

### GetFilterModeOk

`func (o *MacFilteringGeneralSetting) GetFilterModeOk() (*int32, bool)`

GetFilterModeOk returns a tuple with the FilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMode

`func (o *MacFilteringGeneralSetting) SetFilterMode(v int32)`

SetFilterMode sets FilterMode field to given value.

### HasFilterMode

`func (o *MacFilteringGeneralSetting) HasFilterMode() bool`

HasFilterMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


