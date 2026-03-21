# LanNetworkIpv6ConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dhcpv6** | Pointer to [**Dhcpv6VO**](Dhcpv6VO.md) |  | [optional] 
**Enable** | **int32** |  | 
**Passthrough** | Pointer to [**PassthroughVO**](PassthroughVO.md) |  | [optional] 
**Proto** | **string** |  | 
**Ra** | Pointer to [**RouteAdvertisementVO**](RouteAdvertisementVO.md) |  | [optional] 
**Rdnss** | Pointer to [**SlaacVO**](SlaacVO.md) |  | [optional] 
**Slaac** | Pointer to [**SlaacVO**](SlaacVO.md) |  | [optional] 

## Methods

### NewLanNetworkIpv6ConfigVO

`func NewLanNetworkIpv6ConfigVO(enable int32, proto string, ) *LanNetworkIpv6ConfigVO`

NewLanNetworkIpv6ConfigVO instantiates a new LanNetworkIpv6ConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkIpv6ConfigVOWithDefaults

`func NewLanNetworkIpv6ConfigVOWithDefaults() *LanNetworkIpv6ConfigVO`

NewLanNetworkIpv6ConfigVOWithDefaults instantiates a new LanNetworkIpv6ConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpv6

`func (o *LanNetworkIpv6ConfigVO) GetDhcpv6() Dhcpv6VO`

GetDhcpv6 returns the Dhcpv6 field if non-nil, zero value otherwise.

### GetDhcpv6Ok

`func (o *LanNetworkIpv6ConfigVO) GetDhcpv6Ok() (*Dhcpv6VO, bool)`

GetDhcpv6Ok returns a tuple with the Dhcpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6

`func (o *LanNetworkIpv6ConfigVO) SetDhcpv6(v Dhcpv6VO)`

SetDhcpv6 sets Dhcpv6 field to given value.

### HasDhcpv6

`func (o *LanNetworkIpv6ConfigVO) HasDhcpv6() bool`

HasDhcpv6 returns a boolean if a field has been set.

### GetEnable

`func (o *LanNetworkIpv6ConfigVO) GetEnable() int32`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *LanNetworkIpv6ConfigVO) GetEnableOk() (*int32, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *LanNetworkIpv6ConfigVO) SetEnable(v int32)`

SetEnable sets Enable field to given value.


### GetPassthrough

`func (o *LanNetworkIpv6ConfigVO) GetPassthrough() PassthroughVO`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *LanNetworkIpv6ConfigVO) GetPassthroughOk() (*PassthroughVO, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *LanNetworkIpv6ConfigVO) SetPassthrough(v PassthroughVO)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *LanNetworkIpv6ConfigVO) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.

### GetProto

`func (o *LanNetworkIpv6ConfigVO) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *LanNetworkIpv6ConfigVO) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *LanNetworkIpv6ConfigVO) SetProto(v string)`

SetProto sets Proto field to given value.


### GetRa

`func (o *LanNetworkIpv6ConfigVO) GetRa() RouteAdvertisementVO`

GetRa returns the Ra field if non-nil, zero value otherwise.

### GetRaOk

`func (o *LanNetworkIpv6ConfigVO) GetRaOk() (*RouteAdvertisementVO, bool)`

GetRaOk returns a tuple with the Ra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRa

`func (o *LanNetworkIpv6ConfigVO) SetRa(v RouteAdvertisementVO)`

SetRa sets Ra field to given value.

### HasRa

`func (o *LanNetworkIpv6ConfigVO) HasRa() bool`

HasRa returns a boolean if a field has been set.

### GetRdnss

`func (o *LanNetworkIpv6ConfigVO) GetRdnss() SlaacVO`

GetRdnss returns the Rdnss field if non-nil, zero value otherwise.

### GetRdnssOk

`func (o *LanNetworkIpv6ConfigVO) GetRdnssOk() (*SlaacVO, bool)`

GetRdnssOk returns a tuple with the Rdnss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdnss

`func (o *LanNetworkIpv6ConfigVO) SetRdnss(v SlaacVO)`

SetRdnss sets Rdnss field to given value.

### HasRdnss

`func (o *LanNetworkIpv6ConfigVO) HasRdnss() bool`

HasRdnss returns a boolean if a field has been set.

### GetSlaac

`func (o *LanNetworkIpv6ConfigVO) GetSlaac() SlaacVO`

GetSlaac returns the Slaac field if non-nil, zero value otherwise.

### GetSlaacOk

`func (o *LanNetworkIpv6ConfigVO) GetSlaacOk() (*SlaacVO, bool)`

GetSlaacOk returns a tuple with the Slaac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaac

`func (o *LanNetworkIpv6ConfigVO) SetSlaac(v SlaacVO)`

SetSlaac sets Slaac field to given value.

### HasSlaac

`func (o *LanNetworkIpv6ConfigVO) HasSlaac() bool`

HasSlaac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


