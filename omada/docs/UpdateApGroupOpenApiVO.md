# UpdateApGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddApMacs** | Pointer to **[]string** | List of AP device MAC addresses to be added to this AP group. Can be empty. | [optional] 
**Name** | **string** | AP group name should contain 1 to 128 characters. | 
**RemoveApMacs** | Pointer to **[]string** | List of AP device MAC addresses to be removed from this AP group. Can be empty. | [optional] 

## Methods

### NewUpdateApGroupOpenApiVO

`func NewUpdateApGroupOpenApiVO(name string, ) *UpdateApGroupOpenApiVO`

NewUpdateApGroupOpenApiVO instantiates a new UpdateApGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApGroupOpenApiVOWithDefaults

`func NewUpdateApGroupOpenApiVOWithDefaults() *UpdateApGroupOpenApiVO`

NewUpdateApGroupOpenApiVOWithDefaults instantiates a new UpdateApGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddApMacs

`func (o *UpdateApGroupOpenApiVO) GetAddApMacs() []string`

GetAddApMacs returns the AddApMacs field if non-nil, zero value otherwise.

### GetAddApMacsOk

`func (o *UpdateApGroupOpenApiVO) GetAddApMacsOk() (*[]string, bool)`

GetAddApMacsOk returns a tuple with the AddApMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddApMacs

`func (o *UpdateApGroupOpenApiVO) SetAddApMacs(v []string)`

SetAddApMacs sets AddApMacs field to given value.

### HasAddApMacs

`func (o *UpdateApGroupOpenApiVO) HasAddApMacs() bool`

HasAddApMacs returns a boolean if a field has been set.

### GetName

`func (o *UpdateApGroupOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApGroupOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateApGroupOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetRemoveApMacs

`func (o *UpdateApGroupOpenApiVO) GetRemoveApMacs() []string`

GetRemoveApMacs returns the RemoveApMacs field if non-nil, zero value otherwise.

### GetRemoveApMacsOk

`func (o *UpdateApGroupOpenApiVO) GetRemoveApMacsOk() (*[]string, bool)`

GetRemoveApMacsOk returns a tuple with the RemoveApMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveApMacs

`func (o *UpdateApGroupOpenApiVO) SetRemoveApMacs(v []string)`

SetRemoveApMacs sets RemoveApMacs field to given value.

### HasRemoveApMacs

`func (o *UpdateApGroupOpenApiVO) HasRemoveApMacs() bool`

HasRemoveApMacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


