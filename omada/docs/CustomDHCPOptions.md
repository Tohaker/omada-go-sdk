# CustomDHCPOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Custom DHCP option code | [optional] 
**Type** | Pointer to **int32** | Type should be a value as follows: 0: \&quot;String\&quot;; 1: \&quot;IP Address\&quot;; 2: \&quot;Hex Array\&quot; | [optional] 
**Value** | Pointer to **string** | Value | [optional] 

## Methods

### NewCustomDHCPOptions

`func NewCustomDHCPOptions() *CustomDHCPOptions`

NewCustomDHCPOptions instantiates a new CustomDHCPOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDHCPOptionsWithDefaults

`func NewCustomDHCPOptionsWithDefaults() *CustomDHCPOptions`

NewCustomDHCPOptionsWithDefaults instantiates a new CustomDHCPOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CustomDHCPOptions) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustomDHCPOptions) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustomDHCPOptions) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *CustomDHCPOptions) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *CustomDHCPOptions) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomDHCPOptions) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomDHCPOptions) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CustomDHCPOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *CustomDHCPOptions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomDHCPOptions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomDHCPOptions) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomDHCPOptions) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


