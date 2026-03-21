# LocalUserOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyToAllPortals** | Pointer to **bool** | Is the localuser effective for all portals, including all newly created portals | [optional] 
**BindingType** | Pointer to **int32** | MAC binding type should be a value as follows: 0: no binding; 1: static binding; 2: dynamic binding. | [optional] 
**DailyLimit** | Pointer to [**AuthTimeOpenApiVO**](AuthTimeOpenApiVO.md) |  | [optional] 
**DailyLimitEnable** | Pointer to **bool** | Whether to enable localuser daily time limit | [optional] 
**DailyLimitLeftMs** | Pointer to **int64** | Daily time left, unit is ms, required when parameter [dailyLimitEnable] is true | [optional] 
**DailyLimitMs** | Pointer to **int64** | Daily time limit, unit is ms, required when parameter [dailyLimitEnable] is true | [optional] 
**Enable** | Pointer to **bool** | Local user enable status | [optional] 
**ExpirationTime** | Pointer to **int64** | Expiration time, unit: ms | [optional] 
**Id** | Pointer to **string** | Local user ID | [optional] 
**Logout** | Pointer to **bool** | local user logout. enable local user logout. | [optional] 
**MacAddress** | Pointer to **string** | Mac address,the value is only available when the macType is static binding or dynamic binding. | [optional] 
**MaxUsers** | Pointer to **int32** | The maximum number of users online at the same time when the MAC binding type is No Binding. It cannot be modified after initialization. Value of Maximum Users should be within the range of 1-2048. | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Overtime** | Pointer to **bool** | Whether the current time has exceeded the expirationTime | [optional] 
**Password** | Pointer to **string** | Password should contain 1 to 128 characters | [optional] 
**Phone** | Pointer to **string** | Phone number should contain 1 to 20 characters. | [optional] 
**Portals** | Pointer to **[]string** | Bound portal names. | [optional] 
**RateLimit** | Pointer to [**RateLimitOpenApiVO**](RateLimitOpenApiVO.md) |  | [optional] 
**TrafficLeft** | Pointer to **bool** | Is there any remaining traffic. | [optional] 
**TrafficLimit** | Pointer to **int64** | Traffic limit in MB. The value should be within the range of 1–10485760. | [optional] 
**TrafficLimitEnable** | Pointer to **bool** | Whether to enable traffic limit. | [optional] 
**TrafficLimitFrequency** | Pointer to **int32** | Frequency of traffic limit should be a value as follows: 0: total; 1: daily; 2: weekly; 3: monthly. | [optional] 
**TrafficUsed** | Pointer to **int64** | Used traffic(MB). | [optional] 
**Used** | Pointer to **int32** | Used quantity. | [optional] 
**UserName** | Pointer to **string** | User name should contain 1 to 128 characters | [optional] 

## Methods

### NewLocalUserOpenApiVO

`func NewLocalUserOpenApiVO() *LocalUserOpenApiVO`

NewLocalUserOpenApiVO instantiates a new LocalUserOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalUserOpenApiVOWithDefaults

`func NewLocalUserOpenApiVOWithDefaults() *LocalUserOpenApiVO`

NewLocalUserOpenApiVOWithDefaults instantiates a new LocalUserOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyToAllPortals

`func (o *LocalUserOpenApiVO) GetApplyToAllPortals() bool`

GetApplyToAllPortals returns the ApplyToAllPortals field if non-nil, zero value otherwise.

### GetApplyToAllPortalsOk

`func (o *LocalUserOpenApiVO) GetApplyToAllPortalsOk() (*bool, bool)`

GetApplyToAllPortalsOk returns a tuple with the ApplyToAllPortals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToAllPortals

`func (o *LocalUserOpenApiVO) SetApplyToAllPortals(v bool)`

SetApplyToAllPortals sets ApplyToAllPortals field to given value.

### HasApplyToAllPortals

`func (o *LocalUserOpenApiVO) HasApplyToAllPortals() bool`

HasApplyToAllPortals returns a boolean if a field has been set.

### GetBindingType

`func (o *LocalUserOpenApiVO) GetBindingType() int32`

GetBindingType returns the BindingType field if non-nil, zero value otherwise.

### GetBindingTypeOk

`func (o *LocalUserOpenApiVO) GetBindingTypeOk() (*int32, bool)`

GetBindingTypeOk returns a tuple with the BindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingType

`func (o *LocalUserOpenApiVO) SetBindingType(v int32)`

SetBindingType sets BindingType field to given value.

### HasBindingType

`func (o *LocalUserOpenApiVO) HasBindingType() bool`

HasBindingType returns a boolean if a field has been set.

### GetDailyLimit

`func (o *LocalUserOpenApiVO) GetDailyLimit() AuthTimeOpenApiVO`

GetDailyLimit returns the DailyLimit field if non-nil, zero value otherwise.

### GetDailyLimitOk

`func (o *LocalUserOpenApiVO) GetDailyLimitOk() (*AuthTimeOpenApiVO, bool)`

GetDailyLimitOk returns a tuple with the DailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimit

`func (o *LocalUserOpenApiVO) SetDailyLimit(v AuthTimeOpenApiVO)`

SetDailyLimit sets DailyLimit field to given value.

### HasDailyLimit

`func (o *LocalUserOpenApiVO) HasDailyLimit() bool`

HasDailyLimit returns a boolean if a field has been set.

### GetDailyLimitEnable

`func (o *LocalUserOpenApiVO) GetDailyLimitEnable() bool`

GetDailyLimitEnable returns the DailyLimitEnable field if non-nil, zero value otherwise.

### GetDailyLimitEnableOk

`func (o *LocalUserOpenApiVO) GetDailyLimitEnableOk() (*bool, bool)`

GetDailyLimitEnableOk returns a tuple with the DailyLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimitEnable

`func (o *LocalUserOpenApiVO) SetDailyLimitEnable(v bool)`

SetDailyLimitEnable sets DailyLimitEnable field to given value.

### HasDailyLimitEnable

`func (o *LocalUserOpenApiVO) HasDailyLimitEnable() bool`

HasDailyLimitEnable returns a boolean if a field has been set.

### GetDailyLimitLeftMs

`func (o *LocalUserOpenApiVO) GetDailyLimitLeftMs() int64`

GetDailyLimitLeftMs returns the DailyLimitLeftMs field if non-nil, zero value otherwise.

### GetDailyLimitLeftMsOk

`func (o *LocalUserOpenApiVO) GetDailyLimitLeftMsOk() (*int64, bool)`

GetDailyLimitLeftMsOk returns a tuple with the DailyLimitLeftMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimitLeftMs

`func (o *LocalUserOpenApiVO) SetDailyLimitLeftMs(v int64)`

SetDailyLimitLeftMs sets DailyLimitLeftMs field to given value.

### HasDailyLimitLeftMs

`func (o *LocalUserOpenApiVO) HasDailyLimitLeftMs() bool`

HasDailyLimitLeftMs returns a boolean if a field has been set.

### GetDailyLimitMs

`func (o *LocalUserOpenApiVO) GetDailyLimitMs() int64`

GetDailyLimitMs returns the DailyLimitMs field if non-nil, zero value otherwise.

### GetDailyLimitMsOk

`func (o *LocalUserOpenApiVO) GetDailyLimitMsOk() (*int64, bool)`

GetDailyLimitMsOk returns a tuple with the DailyLimitMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimitMs

`func (o *LocalUserOpenApiVO) SetDailyLimitMs(v int64)`

SetDailyLimitMs sets DailyLimitMs field to given value.

### HasDailyLimitMs

`func (o *LocalUserOpenApiVO) HasDailyLimitMs() bool`

HasDailyLimitMs returns a boolean if a field has been set.

### GetEnable

`func (o *LocalUserOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *LocalUserOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *LocalUserOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *LocalUserOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetExpirationTime

`func (o *LocalUserOpenApiVO) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *LocalUserOpenApiVO) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *LocalUserOpenApiVO) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *LocalUserOpenApiVO) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetId

`func (o *LocalUserOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalUserOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalUserOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocalUserOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogout

`func (o *LocalUserOpenApiVO) GetLogout() bool`

GetLogout returns the Logout field if non-nil, zero value otherwise.

### GetLogoutOk

`func (o *LocalUserOpenApiVO) GetLogoutOk() (*bool, bool)`

GetLogoutOk returns a tuple with the Logout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogout

`func (o *LocalUserOpenApiVO) SetLogout(v bool)`

SetLogout sets Logout field to given value.

### HasLogout

`func (o *LocalUserOpenApiVO) HasLogout() bool`

HasLogout returns a boolean if a field has been set.

### GetMacAddress

`func (o *LocalUserOpenApiVO) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *LocalUserOpenApiVO) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *LocalUserOpenApiVO) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *LocalUserOpenApiVO) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMaxUsers

`func (o *LocalUserOpenApiVO) GetMaxUsers() int32`

GetMaxUsers returns the MaxUsers field if non-nil, zero value otherwise.

### GetMaxUsersOk

`func (o *LocalUserOpenApiVO) GetMaxUsersOk() (*int32, bool)`

GetMaxUsersOk returns a tuple with the MaxUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsers

`func (o *LocalUserOpenApiVO) SetMaxUsers(v int32)`

SetMaxUsers sets MaxUsers field to given value.

### HasMaxUsers

`func (o *LocalUserOpenApiVO) HasMaxUsers() bool`

HasMaxUsers returns a boolean if a field has been set.

### GetName

`func (o *LocalUserOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocalUserOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocalUserOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocalUserOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOvertime

`func (o *LocalUserOpenApiVO) GetOvertime() bool`

GetOvertime returns the Overtime field if non-nil, zero value otherwise.

### GetOvertimeOk

`func (o *LocalUserOpenApiVO) GetOvertimeOk() (*bool, bool)`

GetOvertimeOk returns a tuple with the Overtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvertime

`func (o *LocalUserOpenApiVO) SetOvertime(v bool)`

SetOvertime sets Overtime field to given value.

### HasOvertime

`func (o *LocalUserOpenApiVO) HasOvertime() bool`

HasOvertime returns a boolean if a field has been set.

### GetPassword

`func (o *LocalUserOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LocalUserOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LocalUserOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LocalUserOpenApiVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhone

`func (o *LocalUserOpenApiVO) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *LocalUserOpenApiVO) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *LocalUserOpenApiVO) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *LocalUserOpenApiVO) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPortals

`func (o *LocalUserOpenApiVO) GetPortals() []string`

GetPortals returns the Portals field if non-nil, zero value otherwise.

### GetPortalsOk

`func (o *LocalUserOpenApiVO) GetPortalsOk() (*[]string, bool)`

GetPortalsOk returns a tuple with the Portals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortals

`func (o *LocalUserOpenApiVO) SetPortals(v []string)`

SetPortals sets Portals field to given value.

### HasPortals

`func (o *LocalUserOpenApiVO) HasPortals() bool`

HasPortals returns a boolean if a field has been set.

### GetRateLimit

`func (o *LocalUserOpenApiVO) GetRateLimit() RateLimitOpenApiVO`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *LocalUserOpenApiVO) GetRateLimitOk() (*RateLimitOpenApiVO, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *LocalUserOpenApiVO) SetRateLimit(v RateLimitOpenApiVO)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *LocalUserOpenApiVO) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetTrafficLeft

`func (o *LocalUserOpenApiVO) GetTrafficLeft() bool`

GetTrafficLeft returns the TrafficLeft field if non-nil, zero value otherwise.

### GetTrafficLeftOk

`func (o *LocalUserOpenApiVO) GetTrafficLeftOk() (*bool, bool)`

GetTrafficLeftOk returns a tuple with the TrafficLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLeft

`func (o *LocalUserOpenApiVO) SetTrafficLeft(v bool)`

SetTrafficLeft sets TrafficLeft field to given value.

### HasTrafficLeft

`func (o *LocalUserOpenApiVO) HasTrafficLeft() bool`

HasTrafficLeft returns a boolean if a field has been set.

### GetTrafficLimit

`func (o *LocalUserOpenApiVO) GetTrafficLimit() int64`

GetTrafficLimit returns the TrafficLimit field if non-nil, zero value otherwise.

### GetTrafficLimitOk

`func (o *LocalUserOpenApiVO) GetTrafficLimitOk() (*int64, bool)`

GetTrafficLimitOk returns a tuple with the TrafficLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimit

`func (o *LocalUserOpenApiVO) SetTrafficLimit(v int64)`

SetTrafficLimit sets TrafficLimit field to given value.

### HasTrafficLimit

`func (o *LocalUserOpenApiVO) HasTrafficLimit() bool`

HasTrafficLimit returns a boolean if a field has been set.

### GetTrafficLimitEnable

`func (o *LocalUserOpenApiVO) GetTrafficLimitEnable() bool`

GetTrafficLimitEnable returns the TrafficLimitEnable field if non-nil, zero value otherwise.

### GetTrafficLimitEnableOk

`func (o *LocalUserOpenApiVO) GetTrafficLimitEnableOk() (*bool, bool)`

GetTrafficLimitEnableOk returns a tuple with the TrafficLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitEnable

`func (o *LocalUserOpenApiVO) SetTrafficLimitEnable(v bool)`

SetTrafficLimitEnable sets TrafficLimitEnable field to given value.

### HasTrafficLimitEnable

`func (o *LocalUserOpenApiVO) HasTrafficLimitEnable() bool`

HasTrafficLimitEnable returns a boolean if a field has been set.

### GetTrafficLimitFrequency

`func (o *LocalUserOpenApiVO) GetTrafficLimitFrequency() int32`

GetTrafficLimitFrequency returns the TrafficLimitFrequency field if non-nil, zero value otherwise.

### GetTrafficLimitFrequencyOk

`func (o *LocalUserOpenApiVO) GetTrafficLimitFrequencyOk() (*int32, bool)`

GetTrafficLimitFrequencyOk returns a tuple with the TrafficLimitFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitFrequency

`func (o *LocalUserOpenApiVO) SetTrafficLimitFrequency(v int32)`

SetTrafficLimitFrequency sets TrafficLimitFrequency field to given value.

### HasTrafficLimitFrequency

`func (o *LocalUserOpenApiVO) HasTrafficLimitFrequency() bool`

HasTrafficLimitFrequency returns a boolean if a field has been set.

### GetTrafficUsed

`func (o *LocalUserOpenApiVO) GetTrafficUsed() int64`

GetTrafficUsed returns the TrafficUsed field if non-nil, zero value otherwise.

### GetTrafficUsedOk

`func (o *LocalUserOpenApiVO) GetTrafficUsedOk() (*int64, bool)`

GetTrafficUsedOk returns a tuple with the TrafficUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUsed

`func (o *LocalUserOpenApiVO) SetTrafficUsed(v int64)`

SetTrafficUsed sets TrafficUsed field to given value.

### HasTrafficUsed

`func (o *LocalUserOpenApiVO) HasTrafficUsed() bool`

HasTrafficUsed returns a boolean if a field has been set.

### GetUsed

`func (o *LocalUserOpenApiVO) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *LocalUserOpenApiVO) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *LocalUserOpenApiVO) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *LocalUserOpenApiVO) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUserName

`func (o *LocalUserOpenApiVO) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *LocalUserOpenApiVO) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *LocalUserOpenApiVO) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *LocalUserOpenApiVO) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


