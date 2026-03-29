# UpdateTabOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name on this Dashboard page must be unique, it should contain 1 to 64 characters. | [optional] 
**Reset** | Pointer to **bool** | True: reset Default TAB (custom TAB cannot be reset). | [optional] 

## Methods

### NewUpdateTabOpenApiVO

`func NewUpdateTabOpenApiVO() *UpdateTabOpenApiVO`

NewUpdateTabOpenApiVO instantiates a new UpdateTabOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTabOpenApiVOWithDefaults

`func NewUpdateTabOpenApiVOWithDefaults() *UpdateTabOpenApiVO`

NewUpdateTabOpenApiVOWithDefaults instantiates a new UpdateTabOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTabOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTabOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTabOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTabOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReset

`func (o *UpdateTabOpenApiVO) GetReset() bool`

GetReset returns the Reset field if non-nil, zero value otherwise.

### GetResetOk

`func (o *UpdateTabOpenApiVO) GetResetOk() (*bool, bool)`

GetResetOk returns a tuple with the Reset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReset

`func (o *UpdateTabOpenApiVO) SetReset(v bool)`

SetReset sets Reset field to given value.

### HasReset

`func (o *UpdateTabOpenApiVO) HasReset() bool`

HasReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


