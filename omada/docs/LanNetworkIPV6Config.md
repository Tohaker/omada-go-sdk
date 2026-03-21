# LanNetworkIPV6Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dhcpv6** | Pointer to [**Dhcpv6Setting**](Dhcpv6Setting.md) |  | [optional] 
**Enable** | **int32** | IPv6 enable should be a value as follows: 0: Disable (Default); 1: Enable | 
**Passthrough** | Pointer to [**LanNetworkProtoPassThroughMode**](LanNetworkProtoPassThroughMode.md) |  | [optional] 
**Proto** | **int32** | The IPv6 Connection Type of LAN port. Proto should be a value as follows:  0: \&quot;none\&quot; (Default); 1: \&quot;DHCPv6\&quot;; 2: \&quot;SLAAC+Stateless DHCP\&quot;; 3: \&quot;SLAAC+RDNSS\&quot;; 4: \&quot;passthrough\&quot; | 
**Ra** | Pointer to [**RaSetting**](RaSetting.md) |  | [optional] 
**Rdnss** | Pointer to [**LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode**](LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode.md) |  | [optional] 
**Slaac** | Pointer to [**LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode**](LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode.md) |  | [optional] 

## Methods

### NewLanNetworkIPV6Config

`func NewLanNetworkIPV6Config(enable int32, proto int32, ) *LanNetworkIPV6Config`

NewLanNetworkIPV6Config instantiates a new LanNetworkIPV6Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkIPV6ConfigWithDefaults

`func NewLanNetworkIPV6ConfigWithDefaults() *LanNetworkIPV6Config`

NewLanNetworkIPV6ConfigWithDefaults instantiates a new LanNetworkIPV6Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpv6

`func (o *LanNetworkIPV6Config) GetDhcpv6() Dhcpv6Setting`

GetDhcpv6 returns the Dhcpv6 field if non-nil, zero value otherwise.

### GetDhcpv6Ok

`func (o *LanNetworkIPV6Config) GetDhcpv6Ok() (*Dhcpv6Setting, bool)`

GetDhcpv6Ok returns a tuple with the Dhcpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6

`func (o *LanNetworkIPV6Config) SetDhcpv6(v Dhcpv6Setting)`

SetDhcpv6 sets Dhcpv6 field to given value.

### HasDhcpv6

`func (o *LanNetworkIPV6Config) HasDhcpv6() bool`

HasDhcpv6 returns a boolean if a field has been set.

### GetEnable

`func (o *LanNetworkIPV6Config) GetEnable() int32`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *LanNetworkIPV6Config) GetEnableOk() (*int32, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *LanNetworkIPV6Config) SetEnable(v int32)`

SetEnable sets Enable field to given value.


### GetPassthrough

`func (o *LanNetworkIPV6Config) GetPassthrough() LanNetworkProtoPassThroughMode`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *LanNetworkIPV6Config) GetPassthroughOk() (*LanNetworkProtoPassThroughMode, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *LanNetworkIPV6Config) SetPassthrough(v LanNetworkProtoPassThroughMode)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *LanNetworkIPV6Config) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.

### GetProto

`func (o *LanNetworkIPV6Config) GetProto() int32`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *LanNetworkIPV6Config) GetProtoOk() (*int32, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *LanNetworkIPV6Config) SetProto(v int32)`

SetProto sets Proto field to given value.


### GetRa

`func (o *LanNetworkIPV6Config) GetRa() RaSetting`

GetRa returns the Ra field if non-nil, zero value otherwise.

### GetRaOk

`func (o *LanNetworkIPV6Config) GetRaOk() (*RaSetting, bool)`

GetRaOk returns a tuple with the Ra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRa

`func (o *LanNetworkIPV6Config) SetRa(v RaSetting)`

SetRa sets Ra field to given value.

### HasRa

`func (o *LanNetworkIPV6Config) HasRa() bool`

HasRa returns a boolean if a field has been set.

### GetRdnss

`func (o *LanNetworkIPV6Config) GetRdnss() LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode`

GetRdnss returns the Rdnss field if non-nil, zero value otherwise.

### GetRdnssOk

`func (o *LanNetworkIPV6Config) GetRdnssOk() (*LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode, bool)`

GetRdnssOk returns a tuple with the Rdnss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdnss

`func (o *LanNetworkIPV6Config) SetRdnss(v LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode)`

SetRdnss sets Rdnss field to given value.

### HasRdnss

`func (o *LanNetworkIPV6Config) HasRdnss() bool`

HasRdnss returns a boolean if a field has been set.

### GetSlaac

`func (o *LanNetworkIPV6Config) GetSlaac() LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode`

GetSlaac returns the Slaac field if non-nil, zero value otherwise.

### GetSlaacOk

`func (o *LanNetworkIPV6Config) GetSlaacOk() (*LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode, bool)`

GetSlaacOk returns a tuple with the Slaac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaac

`func (o *LanNetworkIPV6Config) SetSlaac(v LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode)`

SetSlaac sets Slaac field to given value.

### HasSlaac

`func (o *LanNetworkIPV6Config) HasSlaac() bool`

HasSlaac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


