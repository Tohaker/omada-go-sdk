# CreateMacAddressOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | **string** | MAC address, should be a valid MAC address format, e.g. AA-BB-CC-DD-11-22 | 
**Name** | Pointer to **string** | MAC address name, name should contain 1 to 128 characters | [optional] 

## Methods

### NewCreateMacAddressOpenApiVO

`func NewCreateMacAddressOpenApiVO(macAddress string, ) *CreateMacAddressOpenApiVO`

NewCreateMacAddressOpenApiVO instantiates a new CreateMacAddressOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMacAddressOpenApiVOWithDefaults

`func NewCreateMacAddressOpenApiVOWithDefaults() *CreateMacAddressOpenApiVO`

NewCreateMacAddressOpenApiVOWithDefaults instantiates a new CreateMacAddressOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *CreateMacAddressOpenApiVO) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *CreateMacAddressOpenApiVO) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *CreateMacAddressOpenApiVO) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetName

`func (o *CreateMacAddressOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMacAddressOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMacAddressOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateMacAddressOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


