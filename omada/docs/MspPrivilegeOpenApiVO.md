# MspPrivilegeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Whether having all customer privilege. | [optional] 
**Customers** | Pointer to **[]string** | The IDs of customer that can be accessed by this user. | [optional] 

## Methods

### NewMspPrivilegeOpenApiVO

`func NewMspPrivilegeOpenApiVO() *MspPrivilegeOpenApiVO`

NewMspPrivilegeOpenApiVO instantiates a new MspPrivilegeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspPrivilegeOpenApiVOWithDefaults

`func NewMspPrivilegeOpenApiVOWithDefaults() *MspPrivilegeOpenApiVO`

NewMspPrivilegeOpenApiVOWithDefaults instantiates a new MspPrivilegeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *MspPrivilegeOpenApiVO) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *MspPrivilegeOpenApiVO) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *MspPrivilegeOpenApiVO) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *MspPrivilegeOpenApiVO) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetCustomers

`func (o *MspPrivilegeOpenApiVO) GetCustomers() []string`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *MspPrivilegeOpenApiVO) GetCustomersOk() (*[]string, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *MspPrivilegeOpenApiVO) SetCustomers(v []string)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *MspPrivilegeOpenApiVO) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


