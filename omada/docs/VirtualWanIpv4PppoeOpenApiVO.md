# VirtualWanIpv4PppoeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns1** | Pointer to **string** | Primary DNS | [optional] 
**Dns2** | Pointer to **string** | Secondary DNS | [optional] 
**EndTime** | Pointer to **string** | It is required when [linkType] is 2. For example, 12:30. | [optional] 
**Gateway** | Pointer to **string** | Gateway IP. | [optional] 
**IpFromIsp** | Pointer to **string** | Get IP address from ISP. | [optional] 
**Ipaddr** | Pointer to **string** | IP address. | [optional] 
**Ipv4Connection2** | Pointer to [**VirtualWanIpv4Connection2OpenApiVO**](VirtualWanIpv4Connection2OpenApiVO.md) |  | [optional] 
**LinkType** | Pointer to **string** | Connection Mode. Parameter [linkType] should be as follows: auto: Connect Automatically; demand: Connect Manually; time: Time-based. | [optional] 
**Mru** | Pointer to **int32** | Parameter [mru] should be a value between 576 and 1492 | [optional] 
**MssClampingType** | Pointer to **int32** |  | [optional] 
**MssClampingValue** | Pointer to **int32** |  | [optional] 
**Mtu** | Pointer to **int32** | Parameter [mtu] should be a value between 576 and 1492 | [optional] 
**Netmask** | Pointer to **string** | Subnet mask. | [optional] 
**Password** | Pointer to **string** | Password. Parameter [password] should contain 1 to 255 ASCII characters. | [optional] 
**RedialInterval** | Pointer to **int32** | It is required when [linkType] is auto. Unit: Second | [optional] 
**Service** | Pointer to **string** | Service name. Keep parameter [service] blank unless your ISP requires you to configure it. | [optional] 
**StartTime** | Pointer to **string** | It is required when [linkType] is 2. For example, 12:30. | [optional] 
**UserName** | Pointer to **string** | Username. Parameter [userName] should contain 1 to 255 ASCII characters. | [optional] 

## Methods

### NewVirtualWanIpv4PppoeOpenApiVO

`func NewVirtualWanIpv4PppoeOpenApiVO() *VirtualWanIpv4PppoeOpenApiVO`

NewVirtualWanIpv4PppoeOpenApiVO instantiates a new VirtualWanIpv4PppoeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanIpv4PppoeOpenApiVOWithDefaults

`func NewVirtualWanIpv4PppoeOpenApiVOWithDefaults() *VirtualWanIpv4PppoeOpenApiVO`

NewVirtualWanIpv4PppoeOpenApiVOWithDefaults instantiates a new VirtualWanIpv4PppoeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns1

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetDns2

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetDns2() string`

GetDns2 returns the Dns2 field if non-nil, zero value otherwise.

### GetDns2Ok

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetDns2Ok() (*string, bool)`

GetDns2Ok returns a tuple with the Dns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns2

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetDns2(v string)`

SetDns2 sets Dns2 field to given value.

### HasDns2

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasDns2() bool`

HasDns2 returns a boolean if a field has been set.

### GetEndTime

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetGateway

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpFromIsp

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetIpFromIsp() string`

GetIpFromIsp returns the IpFromIsp field if non-nil, zero value otherwise.

### GetIpFromIspOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetIpFromIspOk() (*string, bool)`

GetIpFromIspOk returns a tuple with the IpFromIsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFromIsp

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetIpFromIsp(v string)`

SetIpFromIsp sets IpFromIsp field to given value.

### HasIpFromIsp

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasIpFromIsp() bool`

HasIpFromIsp returns a boolean if a field has been set.

### GetIpaddr

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetIpaddr() string`

GetIpaddr returns the Ipaddr field if non-nil, zero value otherwise.

### GetIpaddrOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetIpaddrOk() (*string, bool)`

GetIpaddrOk returns a tuple with the Ipaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddr

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetIpaddr(v string)`

SetIpaddr sets Ipaddr field to given value.

### HasIpaddr

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasIpaddr() bool`

HasIpaddr returns a boolean if a field has been set.

### GetIpv4Connection2

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetIpv4Connection2() VirtualWanIpv4Connection2OpenApiVO`

GetIpv4Connection2 returns the Ipv4Connection2 field if non-nil, zero value otherwise.

### GetIpv4Connection2Ok

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetIpv4Connection2Ok() (*VirtualWanIpv4Connection2OpenApiVO, bool)`

GetIpv4Connection2Ok returns a tuple with the Ipv4Connection2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Connection2

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetIpv4Connection2(v VirtualWanIpv4Connection2OpenApiVO)`

SetIpv4Connection2 sets Ipv4Connection2 field to given value.

### HasIpv4Connection2

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasIpv4Connection2() bool`

HasIpv4Connection2 returns a boolean if a field has been set.

### GetLinkType

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.

### GetMru

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetMru() int32`

GetMru returns the Mru field if non-nil, zero value otherwise.

### GetMruOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetMruOk() (*int32, bool)`

GetMruOk returns a tuple with the Mru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMru

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetMru(v int32)`

SetMru sets Mru field to given value.

### HasMru

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasMru() bool`

HasMru returns a boolean if a field has been set.

### GetMssClampingType

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetMssClampingType() int32`

GetMssClampingType returns the MssClampingType field if non-nil, zero value otherwise.

### GetMssClampingTypeOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetMssClampingTypeOk() (*int32, bool)`

GetMssClampingTypeOk returns a tuple with the MssClampingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssClampingType

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetMssClampingType(v int32)`

SetMssClampingType sets MssClampingType field to given value.

### HasMssClampingType

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasMssClampingType() bool`

HasMssClampingType returns a boolean if a field has been set.

### GetMssClampingValue

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetMssClampingValue() int32`

GetMssClampingValue returns the MssClampingValue field if non-nil, zero value otherwise.

### GetMssClampingValueOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetMssClampingValueOk() (*int32, bool)`

GetMssClampingValueOk returns a tuple with the MssClampingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssClampingValue

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetMssClampingValue(v int32)`

SetMssClampingValue sets MssClampingValue field to given value.

### HasMssClampingValue

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasMssClampingValue() bool`

HasMssClampingValue returns a boolean if a field has been set.

### GetMtu

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNetmask

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetPassword

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRedialInterval

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetRedialInterval() int32`

GetRedialInterval returns the RedialInterval field if non-nil, zero value otherwise.

### GetRedialIntervalOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetRedialIntervalOk() (*int32, bool)`

GetRedialIntervalOk returns a tuple with the RedialInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedialInterval

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetRedialInterval(v int32)`

SetRedialInterval sets RedialInterval field to given value.

### HasRedialInterval

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasRedialInterval() bool`

HasRedialInterval returns a boolean if a field has been set.

### GetService

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStartTime

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUserName

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *VirtualWanIpv4PppoeOpenApiVO) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *VirtualWanIpv4PppoeOpenApiVO) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *VirtualWanIpv4PppoeOpenApiVO) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


