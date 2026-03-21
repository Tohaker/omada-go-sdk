# DnsProxySettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsOverrideSetting** | Pointer to [**DnsOverrideSettingOpenApiVO**](DnsOverrideSettingOpenApiVO.md) |  | [optional] 
**DnsSecSetting** | Pointer to [**DnsSecSettingOpenApiVO**](DnsSecSettingOpenApiVO.md) |  | [optional] 
**DohSetting** | Pointer to [**DohSettingOpenApiVO**](DohSettingOpenApiVO.md) |  | [optional] 
**DotSetting** | Pointer to [**DotSettingOpenApiVO**](DotSettingOpenApiVO.md) |  | [optional] 
**Enable** | **bool** | DNS proxy setting enable status | 
**Type** | Pointer to **int32** | DNS proxy setting type. Type should be a value as follows: 0: DNSSEC, 1: DoH, 2: DoT, 3: DNS Override | [optional] 

## Methods

### NewDnsProxySettingOpenApiVO

`func NewDnsProxySettingOpenApiVO(enable bool, ) *DnsProxySettingOpenApiVO`

NewDnsProxySettingOpenApiVO instantiates a new DnsProxySettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsProxySettingOpenApiVOWithDefaults

`func NewDnsProxySettingOpenApiVOWithDefaults() *DnsProxySettingOpenApiVO`

NewDnsProxySettingOpenApiVOWithDefaults instantiates a new DnsProxySettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsOverrideSetting

`func (o *DnsProxySettingOpenApiVO) GetDnsOverrideSetting() DnsOverrideSettingOpenApiVO`

GetDnsOverrideSetting returns the DnsOverrideSetting field if non-nil, zero value otherwise.

### GetDnsOverrideSettingOk

`func (o *DnsProxySettingOpenApiVO) GetDnsOverrideSettingOk() (*DnsOverrideSettingOpenApiVO, bool)`

GetDnsOverrideSettingOk returns a tuple with the DnsOverrideSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsOverrideSetting

`func (o *DnsProxySettingOpenApiVO) SetDnsOverrideSetting(v DnsOverrideSettingOpenApiVO)`

SetDnsOverrideSetting sets DnsOverrideSetting field to given value.

### HasDnsOverrideSetting

`func (o *DnsProxySettingOpenApiVO) HasDnsOverrideSetting() bool`

HasDnsOverrideSetting returns a boolean if a field has been set.

### GetDnsSecSetting

`func (o *DnsProxySettingOpenApiVO) GetDnsSecSetting() DnsSecSettingOpenApiVO`

GetDnsSecSetting returns the DnsSecSetting field if non-nil, zero value otherwise.

### GetDnsSecSettingOk

`func (o *DnsProxySettingOpenApiVO) GetDnsSecSettingOk() (*DnsSecSettingOpenApiVO, bool)`

GetDnsSecSettingOk returns a tuple with the DnsSecSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecSetting

`func (o *DnsProxySettingOpenApiVO) SetDnsSecSetting(v DnsSecSettingOpenApiVO)`

SetDnsSecSetting sets DnsSecSetting field to given value.

### HasDnsSecSetting

`func (o *DnsProxySettingOpenApiVO) HasDnsSecSetting() bool`

HasDnsSecSetting returns a boolean if a field has been set.

### GetDohSetting

`func (o *DnsProxySettingOpenApiVO) GetDohSetting() DohSettingOpenApiVO`

GetDohSetting returns the DohSetting field if non-nil, zero value otherwise.

### GetDohSettingOk

`func (o *DnsProxySettingOpenApiVO) GetDohSettingOk() (*DohSettingOpenApiVO, bool)`

GetDohSettingOk returns a tuple with the DohSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDohSetting

`func (o *DnsProxySettingOpenApiVO) SetDohSetting(v DohSettingOpenApiVO)`

SetDohSetting sets DohSetting field to given value.

### HasDohSetting

`func (o *DnsProxySettingOpenApiVO) HasDohSetting() bool`

HasDohSetting returns a boolean if a field has been set.

### GetDotSetting

`func (o *DnsProxySettingOpenApiVO) GetDotSetting() DotSettingOpenApiVO`

GetDotSetting returns the DotSetting field if non-nil, zero value otherwise.

### GetDotSettingOk

`func (o *DnsProxySettingOpenApiVO) GetDotSettingOk() (*DotSettingOpenApiVO, bool)`

GetDotSettingOk returns a tuple with the DotSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDotSetting

`func (o *DnsProxySettingOpenApiVO) SetDotSetting(v DotSettingOpenApiVO)`

SetDotSetting sets DotSetting field to given value.

### HasDotSetting

`func (o *DnsProxySettingOpenApiVO) HasDotSetting() bool`

HasDotSetting returns a boolean if a field has been set.

### GetEnable

`func (o *DnsProxySettingOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DnsProxySettingOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DnsProxySettingOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetType

`func (o *DnsProxySettingOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsProxySettingOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsProxySettingOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DnsProxySettingOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


