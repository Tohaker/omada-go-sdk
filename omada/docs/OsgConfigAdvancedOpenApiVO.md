# OsgConfigAdvancedOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EchoServer** | Pointer to **string** | Echo Server should be a domain name or IP address. | [optional] 
**HwOffloadEnable** | Pointer to **bool** | Hardware Offload enabled or not. | [optional] 
**LbSetting2g** | Pointer to [**ApLoadBalanceOpenApiVO**](ApLoadBalanceOpenApiVO.md) |  | [optional] 
**LbSetting5g** | Pointer to [**ApLoadBalanceOpenApiVO**](ApLoadBalanceOpenApiVO.md) |  | [optional] 
**LbSetting5g1** | Pointer to [**ApLoadBalanceOpenApiVO**](ApLoadBalanceOpenApiVO.md) |  | [optional] 
**LbSetting5g2** | Pointer to [**ApLoadBalanceOpenApiVO**](ApLoadBalanceOpenApiVO.md) |  | [optional] 
**LbSetting6g** | Pointer to [**ApLoadBalanceOpenApiVO**](ApLoadBalanceOpenApiVO.md) |  | [optional] 
**LldpEnable** | Pointer to **bool** | LLDP enabled or not. Deprecated, filed lldpSetting is recommended | [optional] 
**LldpSetting** | Pointer to **int32** | LLDP setting. 0: Off, 1: On, 2: Follow site. | [optional] 
**OfdmaEnable2g** | Pointer to **bool** | Enable or disable OFDMA of 2.4 GHz. | [optional] 
**OfdmaEnable5g** | Pointer to **bool** | Enable or disable OFDMA of 5 GHz. | [optional] 
**OfdmaEnable5g1** | Pointer to **bool** | Enable or disable OFDMA of 5 GHz-1. | [optional] 
**OfdmaEnable5g2** | Pointer to **bool** | Enable or disable OFDMA of 5 GHz-2. | [optional] 
**OfdmaEnable6g** | Pointer to **bool** | Enable or disable OFDMA of 6 GHz. | [optional] 
**PoeSettings** | Pointer to [**[]OsgPortPoeOpenApiVO**](OsgPortPoeOpenApiVO.md) | Port Poe setting list. | [optional] 
**QosSetting2g** | Pointer to [**ApQosOpenApiVO**](ApQosOpenApiVO.md) |  | [optional] 
**QosSetting5g** | Pointer to [**ApQosOpenApiVO**](ApQosOpenApiVO.md) |  | [optional] 
**QosSetting5g1** | Pointer to [**ApQosOpenApiVO**](ApQosOpenApiVO.md) |  | [optional] 
**QosSetting5g2** | Pointer to [**ApQosOpenApiVO**](ApQosOpenApiVO.md) |  | [optional] 
**QosSetting6g** | Pointer to [**ApQosOpenApiVO**](ApQosOpenApiVO.md) |  | [optional] 
**RssiSetting2g** | Pointer to [**ApRssiThresholdOpenApiVO**](ApRssiThresholdOpenApiVO.md) |  | [optional] 
**RssiSetting5g** | Pointer to [**ApRssiThresholdOpenApiVO**](ApRssiThresholdOpenApiVO.md) |  | [optional] 
**RssiSetting5g1** | Pointer to [**ApRssiThresholdOpenApiVO**](ApRssiThresholdOpenApiVO.md) |  | [optional] 
**RssiSetting5g2** | Pointer to [**ApRssiThresholdOpenApiVO**](ApRssiThresholdOpenApiVO.md) |  | [optional] 
**RssiSetting6g** | Pointer to [**ApRssiThresholdOpenApiVO**](ApRssiThresholdOpenApiVO.md) |  | [optional] 

## Methods

### NewOsgConfigAdvancedOpenApiVO

`func NewOsgConfigAdvancedOpenApiVO() *OsgConfigAdvancedOpenApiVO`

NewOsgConfigAdvancedOpenApiVO instantiates a new OsgConfigAdvancedOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgConfigAdvancedOpenApiVOWithDefaults

`func NewOsgConfigAdvancedOpenApiVOWithDefaults() *OsgConfigAdvancedOpenApiVO`

NewOsgConfigAdvancedOpenApiVOWithDefaults instantiates a new OsgConfigAdvancedOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEchoServer

`func (o *OsgConfigAdvancedOpenApiVO) GetEchoServer() string`

GetEchoServer returns the EchoServer field if non-nil, zero value otherwise.

### GetEchoServerOk

`func (o *OsgConfigAdvancedOpenApiVO) GetEchoServerOk() (*string, bool)`

GetEchoServerOk returns a tuple with the EchoServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEchoServer

`func (o *OsgConfigAdvancedOpenApiVO) SetEchoServer(v string)`

SetEchoServer sets EchoServer field to given value.

### HasEchoServer

`func (o *OsgConfigAdvancedOpenApiVO) HasEchoServer() bool`

HasEchoServer returns a boolean if a field has been set.

### GetHwOffloadEnable

`func (o *OsgConfigAdvancedOpenApiVO) GetHwOffloadEnable() bool`

GetHwOffloadEnable returns the HwOffloadEnable field if non-nil, zero value otherwise.

### GetHwOffloadEnableOk

`func (o *OsgConfigAdvancedOpenApiVO) GetHwOffloadEnableOk() (*bool, bool)`

GetHwOffloadEnableOk returns a tuple with the HwOffloadEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwOffloadEnable

`func (o *OsgConfigAdvancedOpenApiVO) SetHwOffloadEnable(v bool)`

SetHwOffloadEnable sets HwOffloadEnable field to given value.

### HasHwOffloadEnable

`func (o *OsgConfigAdvancedOpenApiVO) HasHwOffloadEnable() bool`

HasHwOffloadEnable returns a boolean if a field has been set.

### GetLbSetting2g

`func (o *OsgConfigAdvancedOpenApiVO) GetLbSetting2g() ApLoadBalanceOpenApiVO`

GetLbSetting2g returns the LbSetting2g field if non-nil, zero value otherwise.

### GetLbSetting2gOk

`func (o *OsgConfigAdvancedOpenApiVO) GetLbSetting2gOk() (*ApLoadBalanceOpenApiVO, bool)`

GetLbSetting2gOk returns a tuple with the LbSetting2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbSetting2g

`func (o *OsgConfigAdvancedOpenApiVO) SetLbSetting2g(v ApLoadBalanceOpenApiVO)`

SetLbSetting2g sets LbSetting2g field to given value.

### HasLbSetting2g

`func (o *OsgConfigAdvancedOpenApiVO) HasLbSetting2g() bool`

HasLbSetting2g returns a boolean if a field has been set.

### GetLbSetting5g

`func (o *OsgConfigAdvancedOpenApiVO) GetLbSetting5g() ApLoadBalanceOpenApiVO`

GetLbSetting5g returns the LbSetting5g field if non-nil, zero value otherwise.

### GetLbSetting5gOk

`func (o *OsgConfigAdvancedOpenApiVO) GetLbSetting5gOk() (*ApLoadBalanceOpenApiVO, bool)`

GetLbSetting5gOk returns a tuple with the LbSetting5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbSetting5g

`func (o *OsgConfigAdvancedOpenApiVO) SetLbSetting5g(v ApLoadBalanceOpenApiVO)`

SetLbSetting5g sets LbSetting5g field to given value.

### HasLbSetting5g

`func (o *OsgConfigAdvancedOpenApiVO) HasLbSetting5g() bool`

HasLbSetting5g returns a boolean if a field has been set.

### GetLbSetting5g1

`func (o *OsgConfigAdvancedOpenApiVO) GetLbSetting5g1() ApLoadBalanceOpenApiVO`

GetLbSetting5g1 returns the LbSetting5g1 field if non-nil, zero value otherwise.

### GetLbSetting5g1Ok

`func (o *OsgConfigAdvancedOpenApiVO) GetLbSetting5g1Ok() (*ApLoadBalanceOpenApiVO, bool)`

GetLbSetting5g1Ok returns a tuple with the LbSetting5g1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbSetting5g1

`func (o *OsgConfigAdvancedOpenApiVO) SetLbSetting5g1(v ApLoadBalanceOpenApiVO)`

SetLbSetting5g1 sets LbSetting5g1 field to given value.

### HasLbSetting5g1

`func (o *OsgConfigAdvancedOpenApiVO) HasLbSetting5g1() bool`

HasLbSetting5g1 returns a boolean if a field has been set.

### GetLbSetting5g2

`func (o *OsgConfigAdvancedOpenApiVO) GetLbSetting5g2() ApLoadBalanceOpenApiVO`

GetLbSetting5g2 returns the LbSetting5g2 field if non-nil, zero value otherwise.

### GetLbSetting5g2Ok

`func (o *OsgConfigAdvancedOpenApiVO) GetLbSetting5g2Ok() (*ApLoadBalanceOpenApiVO, bool)`

GetLbSetting5g2Ok returns a tuple with the LbSetting5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbSetting5g2

`func (o *OsgConfigAdvancedOpenApiVO) SetLbSetting5g2(v ApLoadBalanceOpenApiVO)`

SetLbSetting5g2 sets LbSetting5g2 field to given value.

### HasLbSetting5g2

`func (o *OsgConfigAdvancedOpenApiVO) HasLbSetting5g2() bool`

HasLbSetting5g2 returns a boolean if a field has been set.

### GetLbSetting6g

`func (o *OsgConfigAdvancedOpenApiVO) GetLbSetting6g() ApLoadBalanceOpenApiVO`

GetLbSetting6g returns the LbSetting6g field if non-nil, zero value otherwise.

### GetLbSetting6gOk

`func (o *OsgConfigAdvancedOpenApiVO) GetLbSetting6gOk() (*ApLoadBalanceOpenApiVO, bool)`

GetLbSetting6gOk returns a tuple with the LbSetting6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbSetting6g

`func (o *OsgConfigAdvancedOpenApiVO) SetLbSetting6g(v ApLoadBalanceOpenApiVO)`

SetLbSetting6g sets LbSetting6g field to given value.

### HasLbSetting6g

`func (o *OsgConfigAdvancedOpenApiVO) HasLbSetting6g() bool`

HasLbSetting6g returns a boolean if a field has been set.

### GetLldpEnable

`func (o *OsgConfigAdvancedOpenApiVO) GetLldpEnable() bool`

GetLldpEnable returns the LldpEnable field if non-nil, zero value otherwise.

### GetLldpEnableOk

`func (o *OsgConfigAdvancedOpenApiVO) GetLldpEnableOk() (*bool, bool)`

GetLldpEnableOk returns a tuple with the LldpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnable

`func (o *OsgConfigAdvancedOpenApiVO) SetLldpEnable(v bool)`

SetLldpEnable sets LldpEnable field to given value.

### HasLldpEnable

`func (o *OsgConfigAdvancedOpenApiVO) HasLldpEnable() bool`

HasLldpEnable returns a boolean if a field has been set.

### GetLldpSetting

`func (o *OsgConfigAdvancedOpenApiVO) GetLldpSetting() int32`

GetLldpSetting returns the LldpSetting field if non-nil, zero value otherwise.

### GetLldpSettingOk

`func (o *OsgConfigAdvancedOpenApiVO) GetLldpSettingOk() (*int32, bool)`

GetLldpSettingOk returns a tuple with the LldpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpSetting

`func (o *OsgConfigAdvancedOpenApiVO) SetLldpSetting(v int32)`

SetLldpSetting sets LldpSetting field to given value.

### HasLldpSetting

`func (o *OsgConfigAdvancedOpenApiVO) HasLldpSetting() bool`

HasLldpSetting returns a boolean if a field has been set.

### GetOfdmaEnable2g

`func (o *OsgConfigAdvancedOpenApiVO) GetOfdmaEnable2g() bool`

GetOfdmaEnable2g returns the OfdmaEnable2g field if non-nil, zero value otherwise.

### GetOfdmaEnable2gOk

`func (o *OsgConfigAdvancedOpenApiVO) GetOfdmaEnable2gOk() (*bool, bool)`

GetOfdmaEnable2gOk returns a tuple with the OfdmaEnable2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfdmaEnable2g

`func (o *OsgConfigAdvancedOpenApiVO) SetOfdmaEnable2g(v bool)`

SetOfdmaEnable2g sets OfdmaEnable2g field to given value.

### HasOfdmaEnable2g

`func (o *OsgConfigAdvancedOpenApiVO) HasOfdmaEnable2g() bool`

HasOfdmaEnable2g returns a boolean if a field has been set.

### GetOfdmaEnable5g

`func (o *OsgConfigAdvancedOpenApiVO) GetOfdmaEnable5g() bool`

GetOfdmaEnable5g returns the OfdmaEnable5g field if non-nil, zero value otherwise.

### GetOfdmaEnable5gOk

`func (o *OsgConfigAdvancedOpenApiVO) GetOfdmaEnable5gOk() (*bool, bool)`

GetOfdmaEnable5gOk returns a tuple with the OfdmaEnable5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfdmaEnable5g

`func (o *OsgConfigAdvancedOpenApiVO) SetOfdmaEnable5g(v bool)`

SetOfdmaEnable5g sets OfdmaEnable5g field to given value.

### HasOfdmaEnable5g

`func (o *OsgConfigAdvancedOpenApiVO) HasOfdmaEnable5g() bool`

HasOfdmaEnable5g returns a boolean if a field has been set.

### GetOfdmaEnable5g1

`func (o *OsgConfigAdvancedOpenApiVO) GetOfdmaEnable5g1() bool`

GetOfdmaEnable5g1 returns the OfdmaEnable5g1 field if non-nil, zero value otherwise.

### GetOfdmaEnable5g1Ok

`func (o *OsgConfigAdvancedOpenApiVO) GetOfdmaEnable5g1Ok() (*bool, bool)`

GetOfdmaEnable5g1Ok returns a tuple with the OfdmaEnable5g1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfdmaEnable5g1

`func (o *OsgConfigAdvancedOpenApiVO) SetOfdmaEnable5g1(v bool)`

SetOfdmaEnable5g1 sets OfdmaEnable5g1 field to given value.

### HasOfdmaEnable5g1

`func (o *OsgConfigAdvancedOpenApiVO) HasOfdmaEnable5g1() bool`

HasOfdmaEnable5g1 returns a boolean if a field has been set.

### GetOfdmaEnable5g2

`func (o *OsgConfigAdvancedOpenApiVO) GetOfdmaEnable5g2() bool`

GetOfdmaEnable5g2 returns the OfdmaEnable5g2 field if non-nil, zero value otherwise.

### GetOfdmaEnable5g2Ok

`func (o *OsgConfigAdvancedOpenApiVO) GetOfdmaEnable5g2Ok() (*bool, bool)`

GetOfdmaEnable5g2Ok returns a tuple with the OfdmaEnable5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfdmaEnable5g2

`func (o *OsgConfigAdvancedOpenApiVO) SetOfdmaEnable5g2(v bool)`

SetOfdmaEnable5g2 sets OfdmaEnable5g2 field to given value.

### HasOfdmaEnable5g2

`func (o *OsgConfigAdvancedOpenApiVO) HasOfdmaEnable5g2() bool`

HasOfdmaEnable5g2 returns a boolean if a field has been set.

### GetOfdmaEnable6g

`func (o *OsgConfigAdvancedOpenApiVO) GetOfdmaEnable6g() bool`

GetOfdmaEnable6g returns the OfdmaEnable6g field if non-nil, zero value otherwise.

### GetOfdmaEnable6gOk

`func (o *OsgConfigAdvancedOpenApiVO) GetOfdmaEnable6gOk() (*bool, bool)`

GetOfdmaEnable6gOk returns a tuple with the OfdmaEnable6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfdmaEnable6g

`func (o *OsgConfigAdvancedOpenApiVO) SetOfdmaEnable6g(v bool)`

SetOfdmaEnable6g sets OfdmaEnable6g field to given value.

### HasOfdmaEnable6g

`func (o *OsgConfigAdvancedOpenApiVO) HasOfdmaEnable6g() bool`

HasOfdmaEnable6g returns a boolean if a field has been set.

### GetPoeSettings

`func (o *OsgConfigAdvancedOpenApiVO) GetPoeSettings() []OsgPortPoeOpenApiVO`

GetPoeSettings returns the PoeSettings field if non-nil, zero value otherwise.

### GetPoeSettingsOk

`func (o *OsgConfigAdvancedOpenApiVO) GetPoeSettingsOk() (*[]OsgPortPoeOpenApiVO, bool)`

GetPoeSettingsOk returns a tuple with the PoeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeSettings

`func (o *OsgConfigAdvancedOpenApiVO) SetPoeSettings(v []OsgPortPoeOpenApiVO)`

SetPoeSettings sets PoeSettings field to given value.

### HasPoeSettings

`func (o *OsgConfigAdvancedOpenApiVO) HasPoeSettings() bool`

HasPoeSettings returns a boolean if a field has been set.

### GetQosSetting2g

`func (o *OsgConfigAdvancedOpenApiVO) GetQosSetting2g() ApQosOpenApiVO`

GetQosSetting2g returns the QosSetting2g field if non-nil, zero value otherwise.

### GetQosSetting2gOk

`func (o *OsgConfigAdvancedOpenApiVO) GetQosSetting2gOk() (*ApQosOpenApiVO, bool)`

GetQosSetting2gOk returns a tuple with the QosSetting2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSetting2g

`func (o *OsgConfigAdvancedOpenApiVO) SetQosSetting2g(v ApQosOpenApiVO)`

SetQosSetting2g sets QosSetting2g field to given value.

### HasQosSetting2g

`func (o *OsgConfigAdvancedOpenApiVO) HasQosSetting2g() bool`

HasQosSetting2g returns a boolean if a field has been set.

### GetQosSetting5g

`func (o *OsgConfigAdvancedOpenApiVO) GetQosSetting5g() ApQosOpenApiVO`

GetQosSetting5g returns the QosSetting5g field if non-nil, zero value otherwise.

### GetQosSetting5gOk

`func (o *OsgConfigAdvancedOpenApiVO) GetQosSetting5gOk() (*ApQosOpenApiVO, bool)`

GetQosSetting5gOk returns a tuple with the QosSetting5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSetting5g

`func (o *OsgConfigAdvancedOpenApiVO) SetQosSetting5g(v ApQosOpenApiVO)`

SetQosSetting5g sets QosSetting5g field to given value.

### HasQosSetting5g

`func (o *OsgConfigAdvancedOpenApiVO) HasQosSetting5g() bool`

HasQosSetting5g returns a boolean if a field has been set.

### GetQosSetting5g1

`func (o *OsgConfigAdvancedOpenApiVO) GetQosSetting5g1() ApQosOpenApiVO`

GetQosSetting5g1 returns the QosSetting5g1 field if non-nil, zero value otherwise.

### GetQosSetting5g1Ok

`func (o *OsgConfigAdvancedOpenApiVO) GetQosSetting5g1Ok() (*ApQosOpenApiVO, bool)`

GetQosSetting5g1Ok returns a tuple with the QosSetting5g1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSetting5g1

`func (o *OsgConfigAdvancedOpenApiVO) SetQosSetting5g1(v ApQosOpenApiVO)`

SetQosSetting5g1 sets QosSetting5g1 field to given value.

### HasQosSetting5g1

`func (o *OsgConfigAdvancedOpenApiVO) HasQosSetting5g1() bool`

HasQosSetting5g1 returns a boolean if a field has been set.

### GetQosSetting5g2

`func (o *OsgConfigAdvancedOpenApiVO) GetQosSetting5g2() ApQosOpenApiVO`

GetQosSetting5g2 returns the QosSetting5g2 field if non-nil, zero value otherwise.

### GetQosSetting5g2Ok

`func (o *OsgConfigAdvancedOpenApiVO) GetQosSetting5g2Ok() (*ApQosOpenApiVO, bool)`

GetQosSetting5g2Ok returns a tuple with the QosSetting5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSetting5g2

`func (o *OsgConfigAdvancedOpenApiVO) SetQosSetting5g2(v ApQosOpenApiVO)`

SetQosSetting5g2 sets QosSetting5g2 field to given value.

### HasQosSetting5g2

`func (o *OsgConfigAdvancedOpenApiVO) HasQosSetting5g2() bool`

HasQosSetting5g2 returns a boolean if a field has been set.

### GetQosSetting6g

`func (o *OsgConfigAdvancedOpenApiVO) GetQosSetting6g() ApQosOpenApiVO`

GetQosSetting6g returns the QosSetting6g field if non-nil, zero value otherwise.

### GetQosSetting6gOk

`func (o *OsgConfigAdvancedOpenApiVO) GetQosSetting6gOk() (*ApQosOpenApiVO, bool)`

GetQosSetting6gOk returns a tuple with the QosSetting6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSetting6g

`func (o *OsgConfigAdvancedOpenApiVO) SetQosSetting6g(v ApQosOpenApiVO)`

SetQosSetting6g sets QosSetting6g field to given value.

### HasQosSetting6g

`func (o *OsgConfigAdvancedOpenApiVO) HasQosSetting6g() bool`

HasQosSetting6g returns a boolean if a field has been set.

### GetRssiSetting2g

`func (o *OsgConfigAdvancedOpenApiVO) GetRssiSetting2g() ApRssiThresholdOpenApiVO`

GetRssiSetting2g returns the RssiSetting2g field if non-nil, zero value otherwise.

### GetRssiSetting2gOk

`func (o *OsgConfigAdvancedOpenApiVO) GetRssiSetting2gOk() (*ApRssiThresholdOpenApiVO, bool)`

GetRssiSetting2gOk returns a tuple with the RssiSetting2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiSetting2g

`func (o *OsgConfigAdvancedOpenApiVO) SetRssiSetting2g(v ApRssiThresholdOpenApiVO)`

SetRssiSetting2g sets RssiSetting2g field to given value.

### HasRssiSetting2g

`func (o *OsgConfigAdvancedOpenApiVO) HasRssiSetting2g() bool`

HasRssiSetting2g returns a boolean if a field has been set.

### GetRssiSetting5g

`func (o *OsgConfigAdvancedOpenApiVO) GetRssiSetting5g() ApRssiThresholdOpenApiVO`

GetRssiSetting5g returns the RssiSetting5g field if non-nil, zero value otherwise.

### GetRssiSetting5gOk

`func (o *OsgConfigAdvancedOpenApiVO) GetRssiSetting5gOk() (*ApRssiThresholdOpenApiVO, bool)`

GetRssiSetting5gOk returns a tuple with the RssiSetting5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiSetting5g

`func (o *OsgConfigAdvancedOpenApiVO) SetRssiSetting5g(v ApRssiThresholdOpenApiVO)`

SetRssiSetting5g sets RssiSetting5g field to given value.

### HasRssiSetting5g

`func (o *OsgConfigAdvancedOpenApiVO) HasRssiSetting5g() bool`

HasRssiSetting5g returns a boolean if a field has been set.

### GetRssiSetting5g1

`func (o *OsgConfigAdvancedOpenApiVO) GetRssiSetting5g1() ApRssiThresholdOpenApiVO`

GetRssiSetting5g1 returns the RssiSetting5g1 field if non-nil, zero value otherwise.

### GetRssiSetting5g1Ok

`func (o *OsgConfigAdvancedOpenApiVO) GetRssiSetting5g1Ok() (*ApRssiThresholdOpenApiVO, bool)`

GetRssiSetting5g1Ok returns a tuple with the RssiSetting5g1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiSetting5g1

`func (o *OsgConfigAdvancedOpenApiVO) SetRssiSetting5g1(v ApRssiThresholdOpenApiVO)`

SetRssiSetting5g1 sets RssiSetting5g1 field to given value.

### HasRssiSetting5g1

`func (o *OsgConfigAdvancedOpenApiVO) HasRssiSetting5g1() bool`

HasRssiSetting5g1 returns a boolean if a field has been set.

### GetRssiSetting5g2

`func (o *OsgConfigAdvancedOpenApiVO) GetRssiSetting5g2() ApRssiThresholdOpenApiVO`

GetRssiSetting5g2 returns the RssiSetting5g2 field if non-nil, zero value otherwise.

### GetRssiSetting5g2Ok

`func (o *OsgConfigAdvancedOpenApiVO) GetRssiSetting5g2Ok() (*ApRssiThresholdOpenApiVO, bool)`

GetRssiSetting5g2Ok returns a tuple with the RssiSetting5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiSetting5g2

`func (o *OsgConfigAdvancedOpenApiVO) SetRssiSetting5g2(v ApRssiThresholdOpenApiVO)`

SetRssiSetting5g2 sets RssiSetting5g2 field to given value.

### HasRssiSetting5g2

`func (o *OsgConfigAdvancedOpenApiVO) HasRssiSetting5g2() bool`

HasRssiSetting5g2 returns a boolean if a field has been set.

### GetRssiSetting6g

`func (o *OsgConfigAdvancedOpenApiVO) GetRssiSetting6g() ApRssiThresholdOpenApiVO`

GetRssiSetting6g returns the RssiSetting6g field if non-nil, zero value otherwise.

### GetRssiSetting6gOk

`func (o *OsgConfigAdvancedOpenApiVO) GetRssiSetting6gOk() (*ApRssiThresholdOpenApiVO, bool)`

GetRssiSetting6gOk returns a tuple with the RssiSetting6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiSetting6g

`func (o *OsgConfigAdvancedOpenApiVO) SetRssiSetting6g(v ApRssiThresholdOpenApiVO)`

SetRssiSetting6g sets RssiSetting6g field to given value.

### HasRssiSetting6g

`func (o *OsgConfigAdvancedOpenApiVO) HasRssiSetting6g() bool`

HasRssiSetting6g returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


