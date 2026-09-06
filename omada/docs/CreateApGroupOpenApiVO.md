# CreateApGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMacs** | Pointer to **[]string** | List of AP device MAC addresses bound to this AP group. Can be empty. | [optional] 
**Name** | **string** | AP group name should contain 1 to 128 characters. | 

## Methods

### NewCreateApGroupOpenApiVO

`func NewCreateApGroupOpenApiVO(name string, ) *CreateApGroupOpenApiVO`

NewCreateApGroupOpenApiVO instantiates a new CreateApGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApGroupOpenApiVOWithDefaults

`func NewCreateApGroupOpenApiVOWithDefaults() *CreateApGroupOpenApiVO`

NewCreateApGroupOpenApiVOWithDefaults instantiates a new CreateApGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMacs

`func (o *CreateApGroupOpenApiVO) GetApMacs() []string`

GetApMacs returns the ApMacs field if non-nil, zero value otherwise.

### GetApMacsOk

`func (o *CreateApGroupOpenApiVO) GetApMacsOk() (*[]string, bool)`

GetApMacsOk returns a tuple with the ApMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMacs

`func (o *CreateApGroupOpenApiVO) SetApMacs(v []string)`

SetApMacs sets ApMacs field to given value.

### HasApMacs

`func (o *CreateApGroupOpenApiVO) HasApMacs() bool`

HasApMacs returns a boolean if a field has been set.

### GetName

`func (o *CreateApGroupOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApGroupOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApGroupOpenApiVO) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


