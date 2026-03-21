# ModifyMspUserVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **bool** | Alert email | [optional] 
**AllCustomer** | **bool** | All customers including new created customer | 
**CustomerRoleId** | **string** | Customer role ID of msp user | 
**Customers** | Pointer to **[]string** | Customer ID list of msp user | [optional] 
**Email** | Pointer to **string** | Email of user | [optional] 
**EndTime** | Pointer to **int64** | The end time of the user&#39;s validity period. time range: end timestamp (Millisecond). | [optional] 
**ForceModify** | Pointer to **bool** | Force modify | [optional] 
**Name** | **string** | User name. When creating cloud user, you should set TP-LINK ID. It should contain 1 to 128 ASCII characters and start with letters, numbers, and underscores. | 
**Password** | Pointer to **string** | Password of local user. User can only modify his child local user&#39;s password. It should contain 8 to 128 ASCII characters.And password must be a combination of uppercase letters, lowercase letters, numbers, and special symbols. Symbols such as ! # $ % &amp; * @ ^ are supported. | [optional] 
**RoleId** | **string** | Role ID of user | 
**StartTime** | Pointer to **int64** | The start time of the user&#39;s validity period. time range: start timestamp (Millisecond). | [optional] 
**TemporaryEnable** | Pointer to **bool** | Whether the user wants to enable the temporary worker permission | [optional] 

## Methods

### NewModifyMspUserVO

`func NewModifyMspUserVO(allCustomer bool, customerRoleId string, name string, roleId string, ) *ModifyMspUserVO`

NewModifyMspUserVO instantiates a new ModifyMspUserVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyMspUserVOWithDefaults

`func NewModifyMspUserVOWithDefaults() *ModifyMspUserVO`

NewModifyMspUserVOWithDefaults instantiates a new ModifyMspUserVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *ModifyMspUserVO) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *ModifyMspUserVO) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *ModifyMspUserVO) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *ModifyMspUserVO) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetAllCustomer

`func (o *ModifyMspUserVO) GetAllCustomer() bool`

GetAllCustomer returns the AllCustomer field if non-nil, zero value otherwise.

### GetAllCustomerOk

`func (o *ModifyMspUserVO) GetAllCustomerOk() (*bool, bool)`

GetAllCustomerOk returns a tuple with the AllCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllCustomer

`func (o *ModifyMspUserVO) SetAllCustomer(v bool)`

SetAllCustomer sets AllCustomer field to given value.


### GetCustomerRoleId

`func (o *ModifyMspUserVO) GetCustomerRoleId() string`

GetCustomerRoleId returns the CustomerRoleId field if non-nil, zero value otherwise.

### GetCustomerRoleIdOk

`func (o *ModifyMspUserVO) GetCustomerRoleIdOk() (*string, bool)`

GetCustomerRoleIdOk returns a tuple with the CustomerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleId

`func (o *ModifyMspUserVO) SetCustomerRoleId(v string)`

SetCustomerRoleId sets CustomerRoleId field to given value.


### GetCustomers

`func (o *ModifyMspUserVO) GetCustomers() []string`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *ModifyMspUserVO) GetCustomersOk() (*[]string, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *ModifyMspUserVO) SetCustomers(v []string)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *ModifyMspUserVO) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetEmail

`func (o *ModifyMspUserVO) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModifyMspUserVO) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModifyMspUserVO) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ModifyMspUserVO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEndTime

`func (o *ModifyMspUserVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ModifyMspUserVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ModifyMspUserVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ModifyMspUserVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetForceModify

`func (o *ModifyMspUserVO) GetForceModify() bool`

GetForceModify returns the ForceModify field if non-nil, zero value otherwise.

### GetForceModifyOk

`func (o *ModifyMspUserVO) GetForceModifyOk() (*bool, bool)`

GetForceModifyOk returns a tuple with the ForceModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceModify

`func (o *ModifyMspUserVO) SetForceModify(v bool)`

SetForceModify sets ForceModify field to given value.

### HasForceModify

`func (o *ModifyMspUserVO) HasForceModify() bool`

HasForceModify returns a boolean if a field has been set.

### GetName

`func (o *ModifyMspUserVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyMspUserVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyMspUserVO) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *ModifyMspUserVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModifyMspUserVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModifyMspUserVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModifyMspUserVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRoleId

`func (o *ModifyMspUserVO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ModifyMspUserVO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ModifyMspUserVO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetStartTime

`func (o *ModifyMspUserVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ModifyMspUserVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ModifyMspUserVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ModifyMspUserVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *ModifyMspUserVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *ModifyMspUserVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *ModifyMspUserVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *ModifyMspUserVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


