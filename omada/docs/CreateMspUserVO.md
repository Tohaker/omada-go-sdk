# CreateMspUserVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **bool** | Whether this user want to receive alert emails | [optional] 
**AllCustomer** | **bool** | Whether msp user has all customer permission, including new created customer. | 
**CustomerRoleId** | **string** | Msp user&#39;s customer role ID when visit customer | 
**Customers** | Pointer to **[]string** | User customer privilege list | [optional] 
**Email** | Pointer to **string** | Email of user | [optional] 
**EndTime** | Pointer to **int64** | The end time of the user&#39;s validity period. time range: end timestamp (Millisecond). | [optional] 
**Name** | **string** | User name. When creating cloud user, you should set TP-LINK ID. It should contain 1 to 128 ASCII visible characters and start with letters, numbers, and underscores. | 
**Password** | Pointer to **string** | Password of local user should contain 8 to 128 ASCII visible characters.And password must be a combination of uppercase letters, lowercase letters, numbers, and special symbols. Symbols such as ! # $ % &amp; * @ ^ are supported. | [optional] 
**RoleId** | **string** | Msp role ID of user | 
**StartTime** | Pointer to **int64** | The start time of the user&#39;s validity period. time range: start timestamp (Millisecond). | [optional] 
**TemporaryEnable** | Pointer to **bool** | Whether the user wants to enable the temporary worker permission | [optional] 
**Type** | **int32** | Type of user should be a value as follows: 0:local user; 1: cloud user | 

## Methods

### NewCreateMspUserVO

`func NewCreateMspUserVO(allCustomer bool, customerRoleId string, name string, roleId string, type_ int32, ) *CreateMspUserVO`

NewCreateMspUserVO instantiates a new CreateMspUserVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMspUserVOWithDefaults

`func NewCreateMspUserVOWithDefaults() *CreateMspUserVO`

NewCreateMspUserVOWithDefaults instantiates a new CreateMspUserVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *CreateMspUserVO) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *CreateMspUserVO) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *CreateMspUserVO) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *CreateMspUserVO) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetAllCustomer

`func (o *CreateMspUserVO) GetAllCustomer() bool`

GetAllCustomer returns the AllCustomer field if non-nil, zero value otherwise.

### GetAllCustomerOk

`func (o *CreateMspUserVO) GetAllCustomerOk() (*bool, bool)`

GetAllCustomerOk returns a tuple with the AllCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllCustomer

`func (o *CreateMspUserVO) SetAllCustomer(v bool)`

SetAllCustomer sets AllCustomer field to given value.


### GetCustomerRoleId

`func (o *CreateMspUserVO) GetCustomerRoleId() string`

GetCustomerRoleId returns the CustomerRoleId field if non-nil, zero value otherwise.

### GetCustomerRoleIdOk

`func (o *CreateMspUserVO) GetCustomerRoleIdOk() (*string, bool)`

GetCustomerRoleIdOk returns a tuple with the CustomerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleId

`func (o *CreateMspUserVO) SetCustomerRoleId(v string)`

SetCustomerRoleId sets CustomerRoleId field to given value.


### GetCustomers

`func (o *CreateMspUserVO) GetCustomers() []string`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *CreateMspUserVO) GetCustomersOk() (*[]string, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *CreateMspUserVO) SetCustomers(v []string)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *CreateMspUserVO) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetEmail

`func (o *CreateMspUserVO) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateMspUserVO) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateMspUserVO) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateMspUserVO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEndTime

`func (o *CreateMspUserVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CreateMspUserVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CreateMspUserVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CreateMspUserVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetName

`func (o *CreateMspUserVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMspUserVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMspUserVO) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateMspUserVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateMspUserVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateMspUserVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateMspUserVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRoleId

`func (o *CreateMspUserVO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CreateMspUserVO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CreateMspUserVO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetStartTime

`func (o *CreateMspUserVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateMspUserVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateMspUserVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CreateMspUserVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *CreateMspUserVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *CreateMspUserVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *CreateMspUserVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *CreateMspUserVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.

### GetType

`func (o *CreateMspUserVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateMspUserVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateMspUserVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


