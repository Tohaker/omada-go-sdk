# LanNetworkIpv6ConfigTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dhcpv6** | Pointer to [**Dhcpv6Setting**](Dhcpv6Setting.md) |  | [optional] 
**Enable** | **int32** | IPv6 enable should be a value as follows: 0: Disable (Default); 1: Enable | 
**Proto** | **int32** | The IPv6 Connection Type of LAN port. Proto should be a value as follows:  0: \&quot;none\&quot; (Default); 1: \&quot;DHCPv6\&quot;; 2: \&quot;SLAAC+Stateless DHCP\&quot;; 3: \&quot;SLAAC+RDNSS\&quot;; 4: \&quot;passthrough\&quot; | 
**Ra** | Pointer to [**RaSetting**](RaSetting.md) |  | [optional] 
**Rdnss** | Pointer to [**SlaacTemplateOpenApiVO**](SlaacTemplateOpenApiVO.md) |  | [optional] 
**Slaac** | Pointer to [**SlaacTemplateOpenApiVO**](SlaacTemplateOpenApiVO.md) |  | [optional] 

## Methods

### NewLanNetworkIpv6ConfigTemplateOpenApiVO

`func NewLanNetworkIpv6ConfigTemplateOpenApiVO(enable int32, proto int32, ) *LanNetworkIpv6ConfigTemplateOpenApiVO`

NewLanNetworkIpv6ConfigTemplateOpenApiVO instantiates a new LanNetworkIpv6ConfigTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkIpv6ConfigTemplateOpenApiVOWithDefaults

`func NewLanNetworkIpv6ConfigTemplateOpenApiVOWithDefaults() *LanNetworkIpv6ConfigTemplateOpenApiVO`

NewLanNetworkIpv6ConfigTemplateOpenApiVOWithDefaults instantiates a new LanNetworkIpv6ConfigTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpv6

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) GetDhcpv6() Dhcpv6Setting`

GetDhcpv6 returns the Dhcpv6 field if non-nil, zero value otherwise.

### GetDhcpv6Ok

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) GetDhcpv6Ok() (*Dhcpv6Setting, bool)`

GetDhcpv6Ok returns a tuple with the Dhcpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) SetDhcpv6(v Dhcpv6Setting)`

SetDhcpv6 sets Dhcpv6 field to given value.

### HasDhcpv6

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) HasDhcpv6() bool`

HasDhcpv6 returns a boolean if a field has been set.

### GetEnable

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) GetEnable() int32`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) GetEnableOk() (*int32, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) SetEnable(v int32)`

SetEnable sets Enable field to given value.


### GetProto

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) GetProto() int32`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) GetProtoOk() (*int32, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) SetProto(v int32)`

SetProto sets Proto field to given value.


### GetRa

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) GetRa() RaSetting`

GetRa returns the Ra field if non-nil, zero value otherwise.

### GetRaOk

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) GetRaOk() (*RaSetting, bool)`

GetRaOk returns a tuple with the Ra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRa

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) SetRa(v RaSetting)`

SetRa sets Ra field to given value.

### HasRa

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) HasRa() bool`

HasRa returns a boolean if a field has been set.

### GetRdnss

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) GetRdnss() SlaacTemplateOpenApiVO`

GetRdnss returns the Rdnss field if non-nil, zero value otherwise.

### GetRdnssOk

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) GetRdnssOk() (*SlaacTemplateOpenApiVO, bool)`

GetRdnssOk returns a tuple with the Rdnss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdnss

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) SetRdnss(v SlaacTemplateOpenApiVO)`

SetRdnss sets Rdnss field to given value.

### HasRdnss

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) HasRdnss() bool`

HasRdnss returns a boolean if a field has been set.

### GetSlaac

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) GetSlaac() SlaacTemplateOpenApiVO`

GetSlaac returns the Slaac field if non-nil, zero value otherwise.

### GetSlaacOk

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) GetSlaacOk() (*SlaacTemplateOpenApiVO, bool)`

GetSlaacOk returns a tuple with the Slaac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaac

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) SetSlaac(v SlaacTemplateOpenApiVO)`

SetSlaac sets Slaac field to given value.

### HasSlaac

`func (o *LanNetworkIpv6ConfigTemplateOpenApiVO) HasSlaac() bool`

HasSlaac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


