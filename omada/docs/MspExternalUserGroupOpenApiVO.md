# MspExternalUserGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllCustomer** | **bool** | Whether having all customer permissions. | 
**CustomerRoleId** | **string** | Customer role ID which can be obtained from &#39;Get customer list&#39; interface. | 
**Customers** | Pointer to **[]string** | The customer IDs that can be accessed. Effective when allCustomer is false. | [optional] 
**EndTime** | Pointer to **int64** | The end time of the user&#39;s validity period. time range: end timestamp (Millisecond). | [optional] 
**Name** | **string** | Msp external user group name should contain 1 to 128 characters. | 
**RoleId** | **string** | Msp role ID which can be obtained from &#39;Get msp role list&#39; interface. | 
**StartTime** | Pointer to **int64** | The start time of the user&#39;s validity period. time range: start timestamp (Millisecond). | [optional] 
**TemporaryEnable** | Pointer to **bool** | Whether the user wants to enable the temporary worker permission | [optional] 

## Methods

### NewMspExternalUserGroupOpenApiVO

`func NewMspExternalUserGroupOpenApiVO(allCustomer bool, customerRoleId string, name string, roleId string, ) *MspExternalUserGroupOpenApiVO`

NewMspExternalUserGroupOpenApiVO instantiates a new MspExternalUserGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspExternalUserGroupOpenApiVOWithDefaults

`func NewMspExternalUserGroupOpenApiVOWithDefaults() *MspExternalUserGroupOpenApiVO`

NewMspExternalUserGroupOpenApiVOWithDefaults instantiates a new MspExternalUserGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllCustomer

`func (o *MspExternalUserGroupOpenApiVO) GetAllCustomer() bool`

GetAllCustomer returns the AllCustomer field if non-nil, zero value otherwise.

### GetAllCustomerOk

`func (o *MspExternalUserGroupOpenApiVO) GetAllCustomerOk() (*bool, bool)`

GetAllCustomerOk returns a tuple with the AllCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllCustomer

`func (o *MspExternalUserGroupOpenApiVO) SetAllCustomer(v bool)`

SetAllCustomer sets AllCustomer field to given value.


### GetCustomerRoleId

`func (o *MspExternalUserGroupOpenApiVO) GetCustomerRoleId() string`

GetCustomerRoleId returns the CustomerRoleId field if non-nil, zero value otherwise.

### GetCustomerRoleIdOk

`func (o *MspExternalUserGroupOpenApiVO) GetCustomerRoleIdOk() (*string, bool)`

GetCustomerRoleIdOk returns a tuple with the CustomerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleId

`func (o *MspExternalUserGroupOpenApiVO) SetCustomerRoleId(v string)`

SetCustomerRoleId sets CustomerRoleId field to given value.


### GetCustomers

`func (o *MspExternalUserGroupOpenApiVO) GetCustomers() []string`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *MspExternalUserGroupOpenApiVO) GetCustomersOk() (*[]string, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *MspExternalUserGroupOpenApiVO) SetCustomers(v []string)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *MspExternalUserGroupOpenApiVO) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetEndTime

`func (o *MspExternalUserGroupOpenApiVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MspExternalUserGroupOpenApiVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MspExternalUserGroupOpenApiVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MspExternalUserGroupOpenApiVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetName

`func (o *MspExternalUserGroupOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MspExternalUserGroupOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MspExternalUserGroupOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetRoleId

`func (o *MspExternalUserGroupOpenApiVO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *MspExternalUserGroupOpenApiVO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *MspExternalUserGroupOpenApiVO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetStartTime

`func (o *MspExternalUserGroupOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MspExternalUserGroupOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MspExternalUserGroupOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MspExternalUserGroupOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *MspExternalUserGroupOpenApiVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *MspExternalUserGroupOpenApiVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *MspExternalUserGroupOpenApiVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *MspExternalUserGroupOpenApiVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


