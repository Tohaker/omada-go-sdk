# MspUserDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **bool** | Whether this user want to receive alert emails | [optional] 
**AllCustomer** | Pointer to **bool** | Whether msp user has all customer permission, including new created customer | [optional] 
**CustomerIds** | Pointer to **[]string** | User customer privilege list | [optional] 
**CustomerRoleId** | Pointer to **string** | Msp user&#39;s customer role ID when visit customer | [optional] 
**CustomerRoleName** | Pointer to **string** | Msp user&#39;s customer role name when visit customer | [optional] 
**Email** | Pointer to **string** | User email | [optional] 
**EndTime** | Pointer to **int64** | The end time of the user&#39;s validity period. time range: end timestamp (Millisecond). | [optional] 
**Id** | Pointer to **string** | User ID | [optional] 
**MspId** | Pointer to **string** | MSP ID | [optional] 
**Name** | Pointer to **string** | User name | [optional] 
**ParentUserId** | Pointer to **string** | User&#39;s parent user ID | [optional] 
**RoleId** | Pointer to **string** | Msp user role ID in MSP | [optional] 
**RoleName** | Pointer to **string** | User bind role name | [optional] 
**ShowTree** | Pointer to **bool** | Whether msp user has Sub-users. | [optional] 
**StartTime** | Pointer to **int64** | The start time of the user&#39;s validity period. time range: start timestamp (Millisecond). | [optional] 
**TemporaryEnable** | Pointer to **bool** | Whether the user wants to enable the temporary worker permission | [optional] 
**TemporaryValidity** | Pointer to **int32** | Whether the temporary user is still valid | [optional] 
**Type** | Pointer to **int32** | Type of user should be a value as follows: 0:local user; 1: cloud user | [optional] 
**Verified** | Pointer to **bool** | Whether this cloud user has verified | [optional] 

## Methods

### NewMspUserDetailVO

`func NewMspUserDetailVO() *MspUserDetailVO`

NewMspUserDetailVO instantiates a new MspUserDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspUserDetailVOWithDefaults

`func NewMspUserDetailVOWithDefaults() *MspUserDetailVO`

NewMspUserDetailVOWithDefaults instantiates a new MspUserDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *MspUserDetailVO) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *MspUserDetailVO) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *MspUserDetailVO) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *MspUserDetailVO) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetAllCustomer

`func (o *MspUserDetailVO) GetAllCustomer() bool`

GetAllCustomer returns the AllCustomer field if non-nil, zero value otherwise.

### GetAllCustomerOk

`func (o *MspUserDetailVO) GetAllCustomerOk() (*bool, bool)`

GetAllCustomerOk returns a tuple with the AllCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllCustomer

`func (o *MspUserDetailVO) SetAllCustomer(v bool)`

SetAllCustomer sets AllCustomer field to given value.

### HasAllCustomer

`func (o *MspUserDetailVO) HasAllCustomer() bool`

HasAllCustomer returns a boolean if a field has been set.

### GetCustomerIds

`func (o *MspUserDetailVO) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *MspUserDetailVO) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *MspUserDetailVO) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *MspUserDetailVO) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### GetCustomerRoleId

`func (o *MspUserDetailVO) GetCustomerRoleId() string`

GetCustomerRoleId returns the CustomerRoleId field if non-nil, zero value otherwise.

### GetCustomerRoleIdOk

`func (o *MspUserDetailVO) GetCustomerRoleIdOk() (*string, bool)`

GetCustomerRoleIdOk returns a tuple with the CustomerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleId

`func (o *MspUserDetailVO) SetCustomerRoleId(v string)`

SetCustomerRoleId sets CustomerRoleId field to given value.

### HasCustomerRoleId

`func (o *MspUserDetailVO) HasCustomerRoleId() bool`

HasCustomerRoleId returns a boolean if a field has been set.

### GetCustomerRoleName

`func (o *MspUserDetailVO) GetCustomerRoleName() string`

GetCustomerRoleName returns the CustomerRoleName field if non-nil, zero value otherwise.

### GetCustomerRoleNameOk

`func (o *MspUserDetailVO) GetCustomerRoleNameOk() (*string, bool)`

GetCustomerRoleNameOk returns a tuple with the CustomerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleName

`func (o *MspUserDetailVO) SetCustomerRoleName(v string)`

SetCustomerRoleName sets CustomerRoleName field to given value.

### HasCustomerRoleName

`func (o *MspUserDetailVO) HasCustomerRoleName() bool`

HasCustomerRoleName returns a boolean if a field has been set.

### GetEmail

`func (o *MspUserDetailVO) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MspUserDetailVO) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MspUserDetailVO) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MspUserDetailVO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEndTime

`func (o *MspUserDetailVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MspUserDetailVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MspUserDetailVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MspUserDetailVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *MspUserDetailVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MspUserDetailVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MspUserDetailVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MspUserDetailVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMspId

`func (o *MspUserDetailVO) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *MspUserDetailVO) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *MspUserDetailVO) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *MspUserDetailVO) HasMspId() bool`

HasMspId returns a boolean if a field has been set.

### GetName

`func (o *MspUserDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MspUserDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MspUserDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MspUserDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentUserId

`func (o *MspUserDetailVO) GetParentUserId() string`

GetParentUserId returns the ParentUserId field if non-nil, zero value otherwise.

### GetParentUserIdOk

`func (o *MspUserDetailVO) GetParentUserIdOk() (*string, bool)`

GetParentUserIdOk returns a tuple with the ParentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUserId

`func (o *MspUserDetailVO) SetParentUserId(v string)`

SetParentUserId sets ParentUserId field to given value.

### HasParentUserId

`func (o *MspUserDetailVO) HasParentUserId() bool`

HasParentUserId returns a boolean if a field has been set.

### GetRoleId

`func (o *MspUserDetailVO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *MspUserDetailVO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *MspUserDetailVO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *MspUserDetailVO) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *MspUserDetailVO) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *MspUserDetailVO) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *MspUserDetailVO) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *MspUserDetailVO) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetShowTree

`func (o *MspUserDetailVO) GetShowTree() bool`

GetShowTree returns the ShowTree field if non-nil, zero value otherwise.

### GetShowTreeOk

`func (o *MspUserDetailVO) GetShowTreeOk() (*bool, bool)`

GetShowTreeOk returns a tuple with the ShowTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTree

`func (o *MspUserDetailVO) SetShowTree(v bool)`

SetShowTree sets ShowTree field to given value.

### HasShowTree

`func (o *MspUserDetailVO) HasShowTree() bool`

HasShowTree returns a boolean if a field has been set.

### GetStartTime

`func (o *MspUserDetailVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MspUserDetailVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MspUserDetailVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MspUserDetailVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *MspUserDetailVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *MspUserDetailVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *MspUserDetailVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *MspUserDetailVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.

### GetTemporaryValidity

`func (o *MspUserDetailVO) GetTemporaryValidity() int32`

GetTemporaryValidity returns the TemporaryValidity field if non-nil, zero value otherwise.

### GetTemporaryValidityOk

`func (o *MspUserDetailVO) GetTemporaryValidityOk() (*int32, bool)`

GetTemporaryValidityOk returns a tuple with the TemporaryValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryValidity

`func (o *MspUserDetailVO) SetTemporaryValidity(v int32)`

SetTemporaryValidity sets TemporaryValidity field to given value.

### HasTemporaryValidity

`func (o *MspUserDetailVO) HasTemporaryValidity() bool`

HasTemporaryValidity returns a boolean if a field has been set.

### GetType

`func (o *MspUserDetailVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MspUserDetailVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MspUserDetailVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *MspUserDetailVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVerified

`func (o *MspUserDetailVO) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *MspUserDetailVO) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *MspUserDetailVO) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *MspUserDetailVO) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


