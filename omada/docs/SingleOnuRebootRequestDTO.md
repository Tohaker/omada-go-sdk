# SingleOnuRebootRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentId** | Pointer to **string** | ONT device ID should contain 1 to 20 characters in ASCII code from \\x21 to \\x7e. | [optional] 
**Key** | Pointer to **string** | Identifier of ONU | [optional] 
**MacAddress** | Pointer to **string** | Mac address of ONU | [optional] 

## Methods

### NewSingleOnuRebootRequestDTO

`func NewSingleOnuRebootRequestDTO() *SingleOnuRebootRequestDTO`

NewSingleOnuRebootRequestDTO instantiates a new SingleOnuRebootRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleOnuRebootRequestDTOWithDefaults

`func NewSingleOnuRebootRequestDTOWithDefaults() *SingleOnuRebootRequestDTO`

NewSingleOnuRebootRequestDTOWithDefaults instantiates a new SingleOnuRebootRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentId

`func (o *SingleOnuRebootRequestDTO) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *SingleOnuRebootRequestDTO) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *SingleOnuRebootRequestDTO) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.

### HasEquipmentId

`func (o *SingleOnuRebootRequestDTO) HasEquipmentId() bool`

HasEquipmentId returns a boolean if a field has been set.

### GetKey

`func (o *SingleOnuRebootRequestDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SingleOnuRebootRequestDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SingleOnuRebootRequestDTO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SingleOnuRebootRequestDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMacAddress

`func (o *SingleOnuRebootRequestDTO) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *SingleOnuRebootRequestDTO) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *SingleOnuRebootRequestDTO) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *SingleOnuRebootRequestDTO) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


