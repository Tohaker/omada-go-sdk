# ApGroupDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApNum** | Pointer to **int64** | Number of APs in this group | [optional] 
**Id** | Pointer to **string** | AP Group ID | [optional] 
**Name** | Pointer to **string** | AP Group Name | [optional] 
**RemainingBinding** | Pointer to **map[string]int32** | Number of SSID remaining bindings for this group | [optional] 

## Methods

### NewApGroupDetailVO

`func NewApGroupDetailVO() *ApGroupDetailVO`

NewApGroupDetailVO instantiates a new ApGroupDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApGroupDetailVOWithDefaults

`func NewApGroupDetailVOWithDefaults() *ApGroupDetailVO`

NewApGroupDetailVOWithDefaults instantiates a new ApGroupDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApNum

`func (o *ApGroupDetailVO) GetApNum() int64`

GetApNum returns the ApNum field if non-nil, zero value otherwise.

### GetApNumOk

`func (o *ApGroupDetailVO) GetApNumOk() (*int64, bool)`

GetApNumOk returns a tuple with the ApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApNum

`func (o *ApGroupDetailVO) SetApNum(v int64)`

SetApNum sets ApNum field to given value.

### HasApNum

`func (o *ApGroupDetailVO) HasApNum() bool`

HasApNum returns a boolean if a field has been set.

### GetId

`func (o *ApGroupDetailVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApGroupDetailVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApGroupDetailVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApGroupDetailVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApGroupDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApGroupDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApGroupDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApGroupDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemainingBinding

`func (o *ApGroupDetailVO) GetRemainingBinding() map[string]int32`

GetRemainingBinding returns the RemainingBinding field if non-nil, zero value otherwise.

### GetRemainingBindingOk

`func (o *ApGroupDetailVO) GetRemainingBindingOk() (*map[string]int32, bool)`

GetRemainingBindingOk returns a tuple with the RemainingBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingBinding

`func (o *ApGroupDetailVO) SetRemainingBinding(v map[string]int32)`

SetRemainingBinding sets RemainingBinding field to given value.

### HasRemainingBinding

`func (o *ApGroupDetailVO) HasRemainingBinding() bool`

HasRemainingBinding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


