# Ipv4PptpOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionMode** | **int32** | It should be a value as follows: 0: Connect Automatically; 1: Connect Manually; 2: Time-based. | 
**EndTime** | Pointer to **string** | It is required when [linkType] is 2. For example, 12:30. | [optional] 
**IpFromIsp** | **bool** | Get IP address from ISP. | 
**Ipv4Connection2** | [**Ipv4Connection2OpenApiVO**](Ipv4Connection2OpenApiVO.md) |  | 
**MssClampingType** | Pointer to **int32** | 0: Disable, 1: Auto, 2: Custom | [optional] 
**MssClampingValue** | Pointer to **int32** | (Optional) It is required when [mssClampingType] is 2, which ranges from 532 ~ 1452. | [optional] 
**Mtu** | **int32** | 576-1500, default:1420 | 
**Password** | **string** |  | 
**PrimaryDns** | Pointer to **string** | Primary DNS | [optional] 
**RedialInterval** | Pointer to **int32** | It is required when [linkType] is 0. Unit: Second | [optional] 
**SecondaryDns** | Pointer to **string** | Secondary DNS | [optional] 
**StartTime** | Pointer to **string** | It is required when [linkType] is 2. For example, 12:30. | [optional] 
**UserName** | **string** |  | 

## Methods

### NewIpv4PptpOpenApiVO

`func NewIpv4PptpOpenApiVO(connectionMode int32, ipFromIsp bool, ipv4Connection2 Ipv4Connection2OpenApiVO, mtu int32, password string, userName string, ) *Ipv4PptpOpenApiVO`

NewIpv4PptpOpenApiVO instantiates a new Ipv4PptpOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv4PptpOpenApiVOWithDefaults

`func NewIpv4PptpOpenApiVOWithDefaults() *Ipv4PptpOpenApiVO`

NewIpv4PptpOpenApiVOWithDefaults instantiates a new Ipv4PptpOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionMode

`func (o *Ipv4PptpOpenApiVO) GetConnectionMode() int32`

GetConnectionMode returns the ConnectionMode field if non-nil, zero value otherwise.

### GetConnectionModeOk

`func (o *Ipv4PptpOpenApiVO) GetConnectionModeOk() (*int32, bool)`

GetConnectionModeOk returns a tuple with the ConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionMode

`func (o *Ipv4PptpOpenApiVO) SetConnectionMode(v int32)`

SetConnectionMode sets ConnectionMode field to given value.


### GetEndTime

`func (o *Ipv4PptpOpenApiVO) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Ipv4PptpOpenApiVO) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Ipv4PptpOpenApiVO) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Ipv4PptpOpenApiVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetIpFromIsp

`func (o *Ipv4PptpOpenApiVO) GetIpFromIsp() bool`

GetIpFromIsp returns the IpFromIsp field if non-nil, zero value otherwise.

### GetIpFromIspOk

`func (o *Ipv4PptpOpenApiVO) GetIpFromIspOk() (*bool, bool)`

GetIpFromIspOk returns a tuple with the IpFromIsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFromIsp

`func (o *Ipv4PptpOpenApiVO) SetIpFromIsp(v bool)`

SetIpFromIsp sets IpFromIsp field to given value.


### GetIpv4Connection2

`func (o *Ipv4PptpOpenApiVO) GetIpv4Connection2() Ipv4Connection2OpenApiVO`

GetIpv4Connection2 returns the Ipv4Connection2 field if non-nil, zero value otherwise.

### GetIpv4Connection2Ok

`func (o *Ipv4PptpOpenApiVO) GetIpv4Connection2Ok() (*Ipv4Connection2OpenApiVO, bool)`

GetIpv4Connection2Ok returns a tuple with the Ipv4Connection2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Connection2

`func (o *Ipv4PptpOpenApiVO) SetIpv4Connection2(v Ipv4Connection2OpenApiVO)`

SetIpv4Connection2 sets Ipv4Connection2 field to given value.


### GetMssClampingType

`func (o *Ipv4PptpOpenApiVO) GetMssClampingType() int32`

GetMssClampingType returns the MssClampingType field if non-nil, zero value otherwise.

### GetMssClampingTypeOk

`func (o *Ipv4PptpOpenApiVO) GetMssClampingTypeOk() (*int32, bool)`

GetMssClampingTypeOk returns a tuple with the MssClampingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssClampingType

`func (o *Ipv4PptpOpenApiVO) SetMssClampingType(v int32)`

SetMssClampingType sets MssClampingType field to given value.

### HasMssClampingType

`func (o *Ipv4PptpOpenApiVO) HasMssClampingType() bool`

HasMssClampingType returns a boolean if a field has been set.

### GetMssClampingValue

`func (o *Ipv4PptpOpenApiVO) GetMssClampingValue() int32`

GetMssClampingValue returns the MssClampingValue field if non-nil, zero value otherwise.

### GetMssClampingValueOk

`func (o *Ipv4PptpOpenApiVO) GetMssClampingValueOk() (*int32, bool)`

GetMssClampingValueOk returns a tuple with the MssClampingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssClampingValue

`func (o *Ipv4PptpOpenApiVO) SetMssClampingValue(v int32)`

SetMssClampingValue sets MssClampingValue field to given value.

### HasMssClampingValue

`func (o *Ipv4PptpOpenApiVO) HasMssClampingValue() bool`

HasMssClampingValue returns a boolean if a field has been set.

### GetMtu

`func (o *Ipv4PptpOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *Ipv4PptpOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *Ipv4PptpOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.


### GetPassword

`func (o *Ipv4PptpOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Ipv4PptpOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Ipv4PptpOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPrimaryDns

`func (o *Ipv4PptpOpenApiVO) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *Ipv4PptpOpenApiVO) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *Ipv4PptpOpenApiVO) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *Ipv4PptpOpenApiVO) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetRedialInterval

`func (o *Ipv4PptpOpenApiVO) GetRedialInterval() int32`

GetRedialInterval returns the RedialInterval field if non-nil, zero value otherwise.

### GetRedialIntervalOk

`func (o *Ipv4PptpOpenApiVO) GetRedialIntervalOk() (*int32, bool)`

GetRedialIntervalOk returns a tuple with the RedialInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedialInterval

`func (o *Ipv4PptpOpenApiVO) SetRedialInterval(v int32)`

SetRedialInterval sets RedialInterval field to given value.

### HasRedialInterval

`func (o *Ipv4PptpOpenApiVO) HasRedialInterval() bool`

HasRedialInterval returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *Ipv4PptpOpenApiVO) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *Ipv4PptpOpenApiVO) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *Ipv4PptpOpenApiVO) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *Ipv4PptpOpenApiVO) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetStartTime

`func (o *Ipv4PptpOpenApiVO) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Ipv4PptpOpenApiVO) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Ipv4PptpOpenApiVO) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Ipv4PptpOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUserName

`func (o *Ipv4PptpOpenApiVO) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Ipv4PptpOpenApiVO) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Ipv4PptpOpenApiVO) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


