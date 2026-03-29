# Access

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accesses** | Pointer to [**[]AccessItem**](AccessItem.md) |  | [optional] 

## Methods

### NewAccess

`func NewAccess() *Access`

NewAccess instantiates a new Access object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessWithDefaults

`func NewAccessWithDefaults() *Access`

NewAccessWithDefaults instantiates a new Access object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccesses

`func (o *Access) GetAccesses() []AccessItem`

GetAccesses returns the Accesses field if non-nil, zero value otherwise.

### GetAccessesOk

`func (o *Access) GetAccessesOk() (*[]AccessItem, bool)`

GetAccessesOk returns a tuple with the Accesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccesses

`func (o *Access) SetAccesses(v []AccessItem)`

SetAccesses sets Accesses field to given value.

### HasAccesses

`func (o *Access) HasAccesses() bool`

HasAccesses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


