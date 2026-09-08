# ApGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApNum** | Pointer to **int64** | Number of APs in this group | [optional] 
**Id** | Pointer to **string** | AP group ID | [optional] 
**Name** | Pointer to **string** | ap group name should contain 1 to 128 characters. | [optional] 
**Primary** | Pointer to **bool** | Whether it is the default ap group | [optional] 
**RemainingBinding** | Pointer to **map[string]int32** | Number of SSID remaining bindings for this group,  0:2g, 1:5g, 2:6g | [optional] 
**SsidNameList** | Pointer to **[]string** | SSID name list bound to this group | [optional] 

## Methods

### NewApGroupOpenApiVO

`func NewApGroupOpenApiVO() *ApGroupOpenApiVO`

NewApGroupOpenApiVO instantiates a new ApGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApGroupOpenApiVOWithDefaults

`func NewApGroupOpenApiVOWithDefaults() *ApGroupOpenApiVO`

NewApGroupOpenApiVOWithDefaults instantiates a new ApGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApNum

`func (o *ApGroupOpenApiVO) GetApNum() int64`

GetApNum returns the ApNum field if non-nil, zero value otherwise.

### GetApNumOk

`func (o *ApGroupOpenApiVO) GetApNumOk() (*int64, bool)`

GetApNumOk returns a tuple with the ApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApNum

`func (o *ApGroupOpenApiVO) SetApNum(v int64)`

SetApNum sets ApNum field to given value.

### HasApNum

`func (o *ApGroupOpenApiVO) HasApNum() bool`

HasApNum returns a boolean if a field has been set.

### GetId

`func (o *ApGroupOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApGroupOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApGroupOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApGroupOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApGroupOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApGroupOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApGroupOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApGroupOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimary

`func (o *ApGroupOpenApiVO) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *ApGroupOpenApiVO) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *ApGroupOpenApiVO) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *ApGroupOpenApiVO) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetRemainingBinding

`func (o *ApGroupOpenApiVO) GetRemainingBinding() map[string]int32`

GetRemainingBinding returns the RemainingBinding field if non-nil, zero value otherwise.

### GetRemainingBindingOk

`func (o *ApGroupOpenApiVO) GetRemainingBindingOk() (*map[string]int32, bool)`

GetRemainingBindingOk returns a tuple with the RemainingBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingBinding

`func (o *ApGroupOpenApiVO) SetRemainingBinding(v map[string]int32)`

SetRemainingBinding sets RemainingBinding field to given value.

### HasRemainingBinding

`func (o *ApGroupOpenApiVO) HasRemainingBinding() bool`

HasRemainingBinding returns a boolean if a field has been set.

### GetSsidNameList

`func (o *ApGroupOpenApiVO) GetSsidNameList() []string`

GetSsidNameList returns the SsidNameList field if non-nil, zero value otherwise.

### GetSsidNameListOk

`func (o *ApGroupOpenApiVO) GetSsidNameListOk() (*[]string, bool)`

GetSsidNameListOk returns a tuple with the SsidNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNameList

`func (o *ApGroupOpenApiVO) SetSsidNameList(v []string)`

SetSsidNameList sets SsidNameList field to given value.

### HasSsidNameList

`func (o *ApGroupOpenApiVO) HasSsidNameList() bool`

HasSsidNameList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


