# SwitchACLEtherTypeEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Default:false | 
**Value** | Pointer to **string** | Value, if enable is true, value must not be null | [optional] 

## Methods

### NewSwitchACLEtherTypeEntity

`func NewSwitchACLEtherTypeEntity(enable bool, ) *SwitchACLEtherTypeEntity`

NewSwitchACLEtherTypeEntity instantiates a new SwitchACLEtherTypeEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchACLEtherTypeEntityWithDefaults

`func NewSwitchACLEtherTypeEntityWithDefaults() *SwitchACLEtherTypeEntity`

NewSwitchACLEtherTypeEntityWithDefaults instantiates a new SwitchACLEtherTypeEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SwitchACLEtherTypeEntity) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SwitchACLEtherTypeEntity) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SwitchACLEtherTypeEntity) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetValue

`func (o *SwitchACLEtherTypeEntity) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SwitchACLEtherTypeEntity) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SwitchACLEtherTypeEntity) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SwitchACLEtherTypeEntity) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


