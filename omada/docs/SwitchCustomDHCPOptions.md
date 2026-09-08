# SwitchCustomDHCPOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Custom DHCP option code | [optional] 
**Custom** | Pointer to **bool** | Whether is custom by user. | [optional] 
**Name** | Pointer to **string** | Custom DHCP option name. | [optional] 
**Type** | Pointer to **int32** | Type should be a value as follows: 0: \&quot;Hex Array\&quot;; 1: \&quot;String\&quot;; 2: \&quot;IP Address\&quot; | [optional] 
**Value** | Pointer to **string** | Value | [optional] 

## Methods

### NewSwitchCustomDHCPOptions

`func NewSwitchCustomDHCPOptions() *SwitchCustomDHCPOptions`

NewSwitchCustomDHCPOptions instantiates a new SwitchCustomDHCPOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchCustomDHCPOptionsWithDefaults

`func NewSwitchCustomDHCPOptionsWithDefaults() *SwitchCustomDHCPOptions`

NewSwitchCustomDHCPOptionsWithDefaults instantiates a new SwitchCustomDHCPOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SwitchCustomDHCPOptions) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SwitchCustomDHCPOptions) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SwitchCustomDHCPOptions) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *SwitchCustomDHCPOptions) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCustom

`func (o *SwitchCustomDHCPOptions) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *SwitchCustomDHCPOptions) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *SwitchCustomDHCPOptions) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *SwitchCustomDHCPOptions) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetName

`func (o *SwitchCustomDHCPOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwitchCustomDHCPOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwitchCustomDHCPOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SwitchCustomDHCPOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SwitchCustomDHCPOptions) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwitchCustomDHCPOptions) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwitchCustomDHCPOptions) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SwitchCustomDHCPOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *SwitchCustomDHCPOptions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SwitchCustomDHCPOptions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SwitchCustomDHCPOptions) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SwitchCustomDHCPOptions) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


