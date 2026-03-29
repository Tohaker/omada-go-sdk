# ModifyLocalUserOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyToAllPortals** | Pointer to **bool** | Is the localuser effective for all portals, including all newly created portals | [optional] 
**BindingType** | **int32** | MAC binding type should be a value as follows: 0: no binding; 1: static binding; 2: dynamic binding. | 
**DailyLimit** | Pointer to [**DailyAuthTimeOpenApiVO**](DailyAuthTimeOpenApiVO.md) |  | [optional] 
**DailyLimitEnable** | Pointer to **bool** | Whether to enable localuser daily time limit | [optional] 
**Enable** | **bool** | Whether to enable. | 
**ExpirationTime** | **int64** | Expiration timestamp. Unit:ms. | 
**Logout** | Pointer to **bool** | local user logout. enable local user logout | [optional] 
**MacAddress** | Pointer to **string** | Mac address,the value is only available when the macType is static binding or dynamic binding. | [optional] 
**MaxUsers** | **int32** | The maximum number of users online at the same time when the MAC binding type is No Binding. It cannot be modified after initialization. MaxUsers should be within the range of 1–2048. | 
**Name** | Pointer to **string** | Name should contain 1 to 128 characters, with no spaces at the beginning and end, and spaces in the middle | [optional] 
**Password** | **string** | Password should contain 1 to 128 characters. | 
**Phone** | Pointer to **string** | Phone number should contain 1 to 20 characters. | [optional] 
**Portals** | **[]string** | Bound portal ID list. Portal can be created using &#39;Add portal&#39; interface, and portal ID can be obtained from &#39;Get portal list in a site&#39; interface | 
**RateLimit** | [**RateLimitOpenApiVO**](RateLimitOpenApiVO.md) |  | 
**TrafficLimit** | Pointer to **int64** | Traffic limit in MB. It should be within the range of 1–10485760. | [optional] 
**TrafficLimitEnable** | **bool** | Whether to enable traffic limit. | 
**TrafficLimitFrequency** | Pointer to **int32** | Frequency of traffic limit should be a value as follows: 0: total; 1: daily; 2: weekly; 3: monthly. | [optional] 

## Methods

### NewModifyLocalUserOpenApiVO

`func NewModifyLocalUserOpenApiVO(bindingType int32, enable bool, expirationTime int64, maxUsers int32, password string, portals []string, rateLimit RateLimitOpenApiVO, trafficLimitEnable bool, ) *ModifyLocalUserOpenApiVO`

NewModifyLocalUserOpenApiVO instantiates a new ModifyLocalUserOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyLocalUserOpenApiVOWithDefaults

`func NewModifyLocalUserOpenApiVOWithDefaults() *ModifyLocalUserOpenApiVO`

NewModifyLocalUserOpenApiVOWithDefaults instantiates a new ModifyLocalUserOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyToAllPortals

`func (o *ModifyLocalUserOpenApiVO) GetApplyToAllPortals() bool`

GetApplyToAllPortals returns the ApplyToAllPortals field if non-nil, zero value otherwise.

### GetApplyToAllPortalsOk

`func (o *ModifyLocalUserOpenApiVO) GetApplyToAllPortalsOk() (*bool, bool)`

GetApplyToAllPortalsOk returns a tuple with the ApplyToAllPortals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToAllPortals

`func (o *ModifyLocalUserOpenApiVO) SetApplyToAllPortals(v bool)`

SetApplyToAllPortals sets ApplyToAllPortals field to given value.

### HasApplyToAllPortals

`func (o *ModifyLocalUserOpenApiVO) HasApplyToAllPortals() bool`

HasApplyToAllPortals returns a boolean if a field has been set.

### GetBindingType

`func (o *ModifyLocalUserOpenApiVO) GetBindingType() int32`

GetBindingType returns the BindingType field if non-nil, zero value otherwise.

### GetBindingTypeOk

`func (o *ModifyLocalUserOpenApiVO) GetBindingTypeOk() (*int32, bool)`

GetBindingTypeOk returns a tuple with the BindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingType

`func (o *ModifyLocalUserOpenApiVO) SetBindingType(v int32)`

SetBindingType sets BindingType field to given value.


### GetDailyLimit

`func (o *ModifyLocalUserOpenApiVO) GetDailyLimit() DailyAuthTimeOpenApiVO`

GetDailyLimit returns the DailyLimit field if non-nil, zero value otherwise.

### GetDailyLimitOk

`func (o *ModifyLocalUserOpenApiVO) GetDailyLimitOk() (*DailyAuthTimeOpenApiVO, bool)`

GetDailyLimitOk returns a tuple with the DailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimit

`func (o *ModifyLocalUserOpenApiVO) SetDailyLimit(v DailyAuthTimeOpenApiVO)`

SetDailyLimit sets DailyLimit field to given value.

### HasDailyLimit

`func (o *ModifyLocalUserOpenApiVO) HasDailyLimit() bool`

HasDailyLimit returns a boolean if a field has been set.

### GetDailyLimitEnable

`func (o *ModifyLocalUserOpenApiVO) GetDailyLimitEnable() bool`

GetDailyLimitEnable returns the DailyLimitEnable field if non-nil, zero value otherwise.

### GetDailyLimitEnableOk

`func (o *ModifyLocalUserOpenApiVO) GetDailyLimitEnableOk() (*bool, bool)`

GetDailyLimitEnableOk returns a tuple with the DailyLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimitEnable

`func (o *ModifyLocalUserOpenApiVO) SetDailyLimitEnable(v bool)`

SetDailyLimitEnable sets DailyLimitEnable field to given value.

### HasDailyLimitEnable

`func (o *ModifyLocalUserOpenApiVO) HasDailyLimitEnable() bool`

HasDailyLimitEnable returns a boolean if a field has been set.

### GetEnable

`func (o *ModifyLocalUserOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ModifyLocalUserOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ModifyLocalUserOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetExpirationTime

`func (o *ModifyLocalUserOpenApiVO) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *ModifyLocalUserOpenApiVO) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *ModifyLocalUserOpenApiVO) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.


### GetLogout

`func (o *ModifyLocalUserOpenApiVO) GetLogout() bool`

GetLogout returns the Logout field if non-nil, zero value otherwise.

### GetLogoutOk

`func (o *ModifyLocalUserOpenApiVO) GetLogoutOk() (*bool, bool)`

GetLogoutOk returns a tuple with the Logout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogout

`func (o *ModifyLocalUserOpenApiVO) SetLogout(v bool)`

SetLogout sets Logout field to given value.

### HasLogout

`func (o *ModifyLocalUserOpenApiVO) HasLogout() bool`

HasLogout returns a boolean if a field has been set.

### GetMacAddress

`func (o *ModifyLocalUserOpenApiVO) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ModifyLocalUserOpenApiVO) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ModifyLocalUserOpenApiVO) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ModifyLocalUserOpenApiVO) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMaxUsers

`func (o *ModifyLocalUserOpenApiVO) GetMaxUsers() int32`

GetMaxUsers returns the MaxUsers field if non-nil, zero value otherwise.

### GetMaxUsersOk

`func (o *ModifyLocalUserOpenApiVO) GetMaxUsersOk() (*int32, bool)`

GetMaxUsersOk returns a tuple with the MaxUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsers

`func (o *ModifyLocalUserOpenApiVO) SetMaxUsers(v int32)`

SetMaxUsers sets MaxUsers field to given value.


### GetName

`func (o *ModifyLocalUserOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyLocalUserOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyLocalUserOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModifyLocalUserOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *ModifyLocalUserOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModifyLocalUserOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModifyLocalUserOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPhone

`func (o *ModifyLocalUserOpenApiVO) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ModifyLocalUserOpenApiVO) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ModifyLocalUserOpenApiVO) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ModifyLocalUserOpenApiVO) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPortals

`func (o *ModifyLocalUserOpenApiVO) GetPortals() []string`

GetPortals returns the Portals field if non-nil, zero value otherwise.

### GetPortalsOk

`func (o *ModifyLocalUserOpenApiVO) GetPortalsOk() (*[]string, bool)`

GetPortalsOk returns a tuple with the Portals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortals

`func (o *ModifyLocalUserOpenApiVO) SetPortals(v []string)`

SetPortals sets Portals field to given value.


### GetRateLimit

`func (o *ModifyLocalUserOpenApiVO) GetRateLimit() RateLimitOpenApiVO`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *ModifyLocalUserOpenApiVO) GetRateLimitOk() (*RateLimitOpenApiVO, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *ModifyLocalUserOpenApiVO) SetRateLimit(v RateLimitOpenApiVO)`

SetRateLimit sets RateLimit field to given value.


### GetTrafficLimit

`func (o *ModifyLocalUserOpenApiVO) GetTrafficLimit() int64`

GetTrafficLimit returns the TrafficLimit field if non-nil, zero value otherwise.

### GetTrafficLimitOk

`func (o *ModifyLocalUserOpenApiVO) GetTrafficLimitOk() (*int64, bool)`

GetTrafficLimitOk returns a tuple with the TrafficLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimit

`func (o *ModifyLocalUserOpenApiVO) SetTrafficLimit(v int64)`

SetTrafficLimit sets TrafficLimit field to given value.

### HasTrafficLimit

`func (o *ModifyLocalUserOpenApiVO) HasTrafficLimit() bool`

HasTrafficLimit returns a boolean if a field has been set.

### GetTrafficLimitEnable

`func (o *ModifyLocalUserOpenApiVO) GetTrafficLimitEnable() bool`

GetTrafficLimitEnable returns the TrafficLimitEnable field if non-nil, zero value otherwise.

### GetTrafficLimitEnableOk

`func (o *ModifyLocalUserOpenApiVO) GetTrafficLimitEnableOk() (*bool, bool)`

GetTrafficLimitEnableOk returns a tuple with the TrafficLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitEnable

`func (o *ModifyLocalUserOpenApiVO) SetTrafficLimitEnable(v bool)`

SetTrafficLimitEnable sets TrafficLimitEnable field to given value.


### GetTrafficLimitFrequency

`func (o *ModifyLocalUserOpenApiVO) GetTrafficLimitFrequency() int32`

GetTrafficLimitFrequency returns the TrafficLimitFrequency field if non-nil, zero value otherwise.

### GetTrafficLimitFrequencyOk

`func (o *ModifyLocalUserOpenApiVO) GetTrafficLimitFrequencyOk() (*int32, bool)`

GetTrafficLimitFrequencyOk returns a tuple with the TrafficLimitFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitFrequency

`func (o *ModifyLocalUserOpenApiVO) SetTrafficLimitFrequency(v int32)`

SetTrafficLimitFrequency sets TrafficLimitFrequency field to given value.

### HasTrafficLimitFrequency

`func (o *ModifyLocalUserOpenApiVO) HasTrafficLimitFrequency() bool`

HasTrafficLimitFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


