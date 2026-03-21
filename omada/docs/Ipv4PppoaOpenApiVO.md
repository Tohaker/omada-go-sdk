# Ipv4PppoaOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionMode** | **int32** | 0: Connect Automatically; 1: Connect Manually; 2: Time-based. | 
**EndTime** | Pointer to **string** | It is required when [linkType] is 2. For example, 12:30. | [optional] 
**IpFromIsp** | **bool** | Get IP address from ISP | 
**Mru** | **int32** | 576-1500, default:1492 | 
**MssClampingType** | **int32** | It should be a value as follows: 0: Disable, 1: Auto, 2: Custom | 
**MssClampingValue** | Pointer to **int32** | (Optional) It is required when [mssClampingType] is 2, which ranges from 532 ~ 1452. | [optional] 
**Mtu** | **int32** | 576-1500, default:1492 | 
**Password** | **string** |  | 
**PrimaryDns** | Pointer to **string** | Primary DNS | [optional] 
**RedialInterval** | Pointer to **int32** | It is required when [linkType] is 0. Unit: Second. | [optional] 
**SecondaryDns** | Pointer to **string** | Secondary DNS | [optional] 
**ServiceName** | Pointer to **string** | Service Name. Keep it blank unless your ISP requires you to configure it. | [optional] 
**StartTime** | Pointer to **string** | It is required when [linkType] is 2. For example, 12:30. | [optional] 
**UserName** | **string** |  | 

## Methods

### NewIpv4PppoaOpenApiVO

`func NewIpv4PppoaOpenApiVO(connectionMode int32, ipFromIsp bool, mru int32, mssClampingType int32, mtu int32, password string, userName string, ) *Ipv4PppoaOpenApiVO`

NewIpv4PppoaOpenApiVO instantiates a new Ipv4PppoaOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv4PppoaOpenApiVOWithDefaults

`func NewIpv4PppoaOpenApiVOWithDefaults() *Ipv4PppoaOpenApiVO`

NewIpv4PppoaOpenApiVOWithDefaults instantiates a new Ipv4PppoaOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionMode

`func (o *Ipv4PppoaOpenApiVO) GetConnectionMode() int32`

GetConnectionMode returns the ConnectionMode field if non-nil, zero value otherwise.

### GetConnectionModeOk

`func (o *Ipv4PppoaOpenApiVO) GetConnectionModeOk() (*int32, bool)`

GetConnectionModeOk returns a tuple with the ConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionMode

`func (o *Ipv4PppoaOpenApiVO) SetConnectionMode(v int32)`

SetConnectionMode sets ConnectionMode field to given value.


### GetEndTime

`func (o *Ipv4PppoaOpenApiVO) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Ipv4PppoaOpenApiVO) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Ipv4PppoaOpenApiVO) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Ipv4PppoaOpenApiVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetIpFromIsp

`func (o *Ipv4PppoaOpenApiVO) GetIpFromIsp() bool`

GetIpFromIsp returns the IpFromIsp field if non-nil, zero value otherwise.

### GetIpFromIspOk

`func (o *Ipv4PppoaOpenApiVO) GetIpFromIspOk() (*bool, bool)`

GetIpFromIspOk returns a tuple with the IpFromIsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFromIsp

`func (o *Ipv4PppoaOpenApiVO) SetIpFromIsp(v bool)`

SetIpFromIsp sets IpFromIsp field to given value.


### GetMru

`func (o *Ipv4PppoaOpenApiVO) GetMru() int32`

GetMru returns the Mru field if non-nil, zero value otherwise.

### GetMruOk

`func (o *Ipv4PppoaOpenApiVO) GetMruOk() (*int32, bool)`

GetMruOk returns a tuple with the Mru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMru

`func (o *Ipv4PppoaOpenApiVO) SetMru(v int32)`

SetMru sets Mru field to given value.


### GetMssClampingType

`func (o *Ipv4PppoaOpenApiVO) GetMssClampingType() int32`

GetMssClampingType returns the MssClampingType field if non-nil, zero value otherwise.

### GetMssClampingTypeOk

`func (o *Ipv4PppoaOpenApiVO) GetMssClampingTypeOk() (*int32, bool)`

GetMssClampingTypeOk returns a tuple with the MssClampingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssClampingType

`func (o *Ipv4PppoaOpenApiVO) SetMssClampingType(v int32)`

SetMssClampingType sets MssClampingType field to given value.


### GetMssClampingValue

`func (o *Ipv4PppoaOpenApiVO) GetMssClampingValue() int32`

GetMssClampingValue returns the MssClampingValue field if non-nil, zero value otherwise.

### GetMssClampingValueOk

`func (o *Ipv4PppoaOpenApiVO) GetMssClampingValueOk() (*int32, bool)`

GetMssClampingValueOk returns a tuple with the MssClampingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssClampingValue

`func (o *Ipv4PppoaOpenApiVO) SetMssClampingValue(v int32)`

SetMssClampingValue sets MssClampingValue field to given value.

### HasMssClampingValue

`func (o *Ipv4PppoaOpenApiVO) HasMssClampingValue() bool`

HasMssClampingValue returns a boolean if a field has been set.

### GetMtu

`func (o *Ipv4PppoaOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *Ipv4PppoaOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *Ipv4PppoaOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.


### GetPassword

`func (o *Ipv4PppoaOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Ipv4PppoaOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Ipv4PppoaOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPrimaryDns

`func (o *Ipv4PppoaOpenApiVO) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *Ipv4PppoaOpenApiVO) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *Ipv4PppoaOpenApiVO) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *Ipv4PppoaOpenApiVO) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetRedialInterval

`func (o *Ipv4PppoaOpenApiVO) GetRedialInterval() int32`

GetRedialInterval returns the RedialInterval field if non-nil, zero value otherwise.

### GetRedialIntervalOk

`func (o *Ipv4PppoaOpenApiVO) GetRedialIntervalOk() (*int32, bool)`

GetRedialIntervalOk returns a tuple with the RedialInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedialInterval

`func (o *Ipv4PppoaOpenApiVO) SetRedialInterval(v int32)`

SetRedialInterval sets RedialInterval field to given value.

### HasRedialInterval

`func (o *Ipv4PppoaOpenApiVO) HasRedialInterval() bool`

HasRedialInterval returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *Ipv4PppoaOpenApiVO) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *Ipv4PppoaOpenApiVO) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *Ipv4PppoaOpenApiVO) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *Ipv4PppoaOpenApiVO) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetServiceName

`func (o *Ipv4PppoaOpenApiVO) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Ipv4PppoaOpenApiVO) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Ipv4PppoaOpenApiVO) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *Ipv4PppoaOpenApiVO) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStartTime

`func (o *Ipv4PppoaOpenApiVO) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Ipv4PppoaOpenApiVO) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Ipv4PppoaOpenApiVO) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Ipv4PppoaOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUserName

`func (o *Ipv4PppoaOpenApiVO) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Ipv4PppoaOpenApiVO) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Ipv4PppoaOpenApiVO) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


