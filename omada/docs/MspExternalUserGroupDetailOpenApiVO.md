# MspExternalUserGroupDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllCustomer** | Pointer to **bool** | Whether having all customer permissions. | [optional] 
**CustomerRoleId** | Pointer to **string** | Customer role ID. | [optional] 
**CustomerRoleName** | Pointer to **string** | Customer role name. | [optional] 
**CustomerRoleType** | Pointer to **int32** | Customer role type. | [optional] 
**Customers** | Pointer to [**[]OmadacInfoOpenApiVO**](OmadacInfoOpenApiVO.md) | The customers which can be accessed. Required when allCustomer is false. | [optional] 
**EndTime** | Pointer to **int64** | The end time of the user&#39;s validity period. time range: end timestamp (Millisecond). | [optional] 
**Id** | Pointer to **string** | Msp external user group ID. | [optional] 
**Name** | Pointer to **string** | Msp external user group name. | [optional] 
**RoleId** | Pointer to **string** | Msp role ID. | [optional] 
**RoleName** | Pointer to **string** | Msp role name. | [optional] 
**RoleType** | Pointer to **int32** | Msp role type. | [optional] 
**StartTime** | Pointer to **int64** | The start time of the user&#39;s validity period. time range: start timestamp (Millisecond). | [optional] 
**TemporaryEnable** | Pointer to **bool** | Whether the user wants to enable the temporary worker permission | [optional] 
**TemporaryValidity** | Pointer to **int32** | Whether the temporary user is still valid | [optional] 

## Methods

### NewMspExternalUserGroupDetailOpenApiVO

`func NewMspExternalUserGroupDetailOpenApiVO() *MspExternalUserGroupDetailOpenApiVO`

NewMspExternalUserGroupDetailOpenApiVO instantiates a new MspExternalUserGroupDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspExternalUserGroupDetailOpenApiVOWithDefaults

`func NewMspExternalUserGroupDetailOpenApiVOWithDefaults() *MspExternalUserGroupDetailOpenApiVO`

NewMspExternalUserGroupDetailOpenApiVOWithDefaults instantiates a new MspExternalUserGroupDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllCustomer

`func (o *MspExternalUserGroupDetailOpenApiVO) GetAllCustomer() bool`

GetAllCustomer returns the AllCustomer field if non-nil, zero value otherwise.

### GetAllCustomerOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetAllCustomerOk() (*bool, bool)`

GetAllCustomerOk returns a tuple with the AllCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllCustomer

`func (o *MspExternalUserGroupDetailOpenApiVO) SetAllCustomer(v bool)`

SetAllCustomer sets AllCustomer field to given value.

### HasAllCustomer

`func (o *MspExternalUserGroupDetailOpenApiVO) HasAllCustomer() bool`

HasAllCustomer returns a boolean if a field has been set.

### GetCustomerRoleId

`func (o *MspExternalUserGroupDetailOpenApiVO) GetCustomerRoleId() string`

GetCustomerRoleId returns the CustomerRoleId field if non-nil, zero value otherwise.

### GetCustomerRoleIdOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetCustomerRoleIdOk() (*string, bool)`

GetCustomerRoleIdOk returns a tuple with the CustomerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleId

`func (o *MspExternalUserGroupDetailOpenApiVO) SetCustomerRoleId(v string)`

SetCustomerRoleId sets CustomerRoleId field to given value.

### HasCustomerRoleId

`func (o *MspExternalUserGroupDetailOpenApiVO) HasCustomerRoleId() bool`

HasCustomerRoleId returns a boolean if a field has been set.

### GetCustomerRoleName

`func (o *MspExternalUserGroupDetailOpenApiVO) GetCustomerRoleName() string`

GetCustomerRoleName returns the CustomerRoleName field if non-nil, zero value otherwise.

### GetCustomerRoleNameOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetCustomerRoleNameOk() (*string, bool)`

GetCustomerRoleNameOk returns a tuple with the CustomerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleName

`func (o *MspExternalUserGroupDetailOpenApiVO) SetCustomerRoleName(v string)`

SetCustomerRoleName sets CustomerRoleName field to given value.

### HasCustomerRoleName

`func (o *MspExternalUserGroupDetailOpenApiVO) HasCustomerRoleName() bool`

HasCustomerRoleName returns a boolean if a field has been set.

### GetCustomerRoleType

`func (o *MspExternalUserGroupDetailOpenApiVO) GetCustomerRoleType() int32`

GetCustomerRoleType returns the CustomerRoleType field if non-nil, zero value otherwise.

### GetCustomerRoleTypeOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetCustomerRoleTypeOk() (*int32, bool)`

GetCustomerRoleTypeOk returns a tuple with the CustomerRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleType

`func (o *MspExternalUserGroupDetailOpenApiVO) SetCustomerRoleType(v int32)`

SetCustomerRoleType sets CustomerRoleType field to given value.

### HasCustomerRoleType

`func (o *MspExternalUserGroupDetailOpenApiVO) HasCustomerRoleType() bool`

HasCustomerRoleType returns a boolean if a field has been set.

### GetCustomers

`func (o *MspExternalUserGroupDetailOpenApiVO) GetCustomers() []OmadacInfoOpenApiVO`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetCustomersOk() (*[]OmadacInfoOpenApiVO, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *MspExternalUserGroupDetailOpenApiVO) SetCustomers(v []OmadacInfoOpenApiVO)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *MspExternalUserGroupDetailOpenApiVO) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetEndTime

`func (o *MspExternalUserGroupDetailOpenApiVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MspExternalUserGroupDetailOpenApiVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MspExternalUserGroupDetailOpenApiVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *MspExternalUserGroupDetailOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MspExternalUserGroupDetailOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MspExternalUserGroupDetailOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MspExternalUserGroupDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MspExternalUserGroupDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MspExternalUserGroupDetailOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleId

`func (o *MspExternalUserGroupDetailOpenApiVO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *MspExternalUserGroupDetailOpenApiVO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *MspExternalUserGroupDetailOpenApiVO) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *MspExternalUserGroupDetailOpenApiVO) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *MspExternalUserGroupDetailOpenApiVO) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *MspExternalUserGroupDetailOpenApiVO) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetRoleType

`func (o *MspExternalUserGroupDetailOpenApiVO) GetRoleType() int32`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetRoleTypeOk() (*int32, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *MspExternalUserGroupDetailOpenApiVO) SetRoleType(v int32)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *MspExternalUserGroupDetailOpenApiVO) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetStartTime

`func (o *MspExternalUserGroupDetailOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MspExternalUserGroupDetailOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MspExternalUserGroupDetailOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *MspExternalUserGroupDetailOpenApiVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *MspExternalUserGroupDetailOpenApiVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *MspExternalUserGroupDetailOpenApiVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.

### GetTemporaryValidity

`func (o *MspExternalUserGroupDetailOpenApiVO) GetTemporaryValidity() int32`

GetTemporaryValidity returns the TemporaryValidity field if non-nil, zero value otherwise.

### GetTemporaryValidityOk

`func (o *MspExternalUserGroupDetailOpenApiVO) GetTemporaryValidityOk() (*int32, bool)`

GetTemporaryValidityOk returns a tuple with the TemporaryValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryValidity

`func (o *MspExternalUserGroupDetailOpenApiVO) SetTemporaryValidity(v int32)`

SetTemporaryValidity sets TemporaryValidity field to given value.

### HasTemporaryValidity

`func (o *MspExternalUserGroupDetailOpenApiVO) HasTemporaryValidity() bool`

HasTemporaryValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


