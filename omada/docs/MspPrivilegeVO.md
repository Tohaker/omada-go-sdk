# MspPrivilegeVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** |  | [optional] 
**Customers** | Pointer to [**[]CustomerVO**](CustomerVO.md) |  | [optional] 
**FavoriteCustomers** | Pointer to **[]string** |  | [optional] 
**LastCustomer** | Pointer to **string** |  | [optional] 

## Methods

### NewMspPrivilegeVO

`func NewMspPrivilegeVO() *MspPrivilegeVO`

NewMspPrivilegeVO instantiates a new MspPrivilegeVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspPrivilegeVOWithDefaults

`func NewMspPrivilegeVOWithDefaults() *MspPrivilegeVO`

NewMspPrivilegeVOWithDefaults instantiates a new MspPrivilegeVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *MspPrivilegeVO) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *MspPrivilegeVO) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *MspPrivilegeVO) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *MspPrivilegeVO) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetCustomers

`func (o *MspPrivilegeVO) GetCustomers() []CustomerVO`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *MspPrivilegeVO) GetCustomersOk() (*[]CustomerVO, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *MspPrivilegeVO) SetCustomers(v []CustomerVO)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *MspPrivilegeVO) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetFavoriteCustomers

`func (o *MspPrivilegeVO) GetFavoriteCustomers() []string`

GetFavoriteCustomers returns the FavoriteCustomers field if non-nil, zero value otherwise.

### GetFavoriteCustomersOk

`func (o *MspPrivilegeVO) GetFavoriteCustomersOk() (*[]string, bool)`

GetFavoriteCustomersOk returns a tuple with the FavoriteCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteCustomers

`func (o *MspPrivilegeVO) SetFavoriteCustomers(v []string)`

SetFavoriteCustomers sets FavoriteCustomers field to given value.

### HasFavoriteCustomers

`func (o *MspPrivilegeVO) HasFavoriteCustomers() bool`

HasFavoriteCustomers returns a boolean if a field has been set.

### GetLastCustomer

`func (o *MspPrivilegeVO) GetLastCustomer() string`

GetLastCustomer returns the LastCustomer field if non-nil, zero value otherwise.

### GetLastCustomerOk

`func (o *MspPrivilegeVO) GetLastCustomerOk() (*string, bool)`

GetLastCustomerOk returns a tuple with the LastCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCustomer

`func (o *MspPrivilegeVO) SetLastCustomer(v string)`

SetLastCustomer sets LastCustomer field to given value.

### HasLastCustomer

`func (o *MspPrivilegeVO) HasLastCustomer() bool`

HasLastCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


