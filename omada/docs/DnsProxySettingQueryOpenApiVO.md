# DnsProxySettingQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsOverrideSetting** | Pointer to [**DnsOverrideSettingOpenApiVO**](DnsOverrideSettingOpenApiVO.md) |  | [optional] 
**DnsSecSetting** | Pointer to [**DnsSecSettingOpenApiVO**](DnsSecSettingOpenApiVO.md) |  | [optional] 
**DohSetting** | Pointer to [**DohSettingOpenApiVO**](DohSettingOpenApiVO.md) |  | [optional] 
**DotSetting** | Pointer to [**DotSettingOpenApiVO**](DotSettingOpenApiVO.md) |  | [optional] 
**Enable** | Pointer to **bool** | DNS proxy setting enable status. | [optional] 
**ExistDnsOverride** | Pointer to **bool** | Whether DNS Override has been configured in DNS Proxy. | [optional] 
**SupportDnsOverride** | Pointer to **bool** | Whether DNS Override setting is supported in DNS Proxy. | [optional] 
**Type** | Pointer to **int32** | DNS proxy setting type. Type should be a value as follows: 0: DNSSEC, 1: DoH, 2: DoT, 3: DNS Override | [optional] 

## Methods

### NewDnsProxySettingQueryOpenApiVO

`func NewDnsProxySettingQueryOpenApiVO() *DnsProxySettingQueryOpenApiVO`

NewDnsProxySettingQueryOpenApiVO instantiates a new DnsProxySettingQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsProxySettingQueryOpenApiVOWithDefaults

`func NewDnsProxySettingQueryOpenApiVOWithDefaults() *DnsProxySettingQueryOpenApiVO`

NewDnsProxySettingQueryOpenApiVOWithDefaults instantiates a new DnsProxySettingQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsOverrideSetting

`func (o *DnsProxySettingQueryOpenApiVO) GetDnsOverrideSetting() DnsOverrideSettingOpenApiVO`

GetDnsOverrideSetting returns the DnsOverrideSetting field if non-nil, zero value otherwise.

### GetDnsOverrideSettingOk

`func (o *DnsProxySettingQueryOpenApiVO) GetDnsOverrideSettingOk() (*DnsOverrideSettingOpenApiVO, bool)`

GetDnsOverrideSettingOk returns a tuple with the DnsOverrideSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsOverrideSetting

`func (o *DnsProxySettingQueryOpenApiVO) SetDnsOverrideSetting(v DnsOverrideSettingOpenApiVO)`

SetDnsOverrideSetting sets DnsOverrideSetting field to given value.

### HasDnsOverrideSetting

`func (o *DnsProxySettingQueryOpenApiVO) HasDnsOverrideSetting() bool`

HasDnsOverrideSetting returns a boolean if a field has been set.

### GetDnsSecSetting

`func (o *DnsProxySettingQueryOpenApiVO) GetDnsSecSetting() DnsSecSettingOpenApiVO`

GetDnsSecSetting returns the DnsSecSetting field if non-nil, zero value otherwise.

### GetDnsSecSettingOk

`func (o *DnsProxySettingQueryOpenApiVO) GetDnsSecSettingOk() (*DnsSecSettingOpenApiVO, bool)`

GetDnsSecSettingOk returns a tuple with the DnsSecSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecSetting

`func (o *DnsProxySettingQueryOpenApiVO) SetDnsSecSetting(v DnsSecSettingOpenApiVO)`

SetDnsSecSetting sets DnsSecSetting field to given value.

### HasDnsSecSetting

`func (o *DnsProxySettingQueryOpenApiVO) HasDnsSecSetting() bool`

HasDnsSecSetting returns a boolean if a field has been set.

### GetDohSetting

`func (o *DnsProxySettingQueryOpenApiVO) GetDohSetting() DohSettingOpenApiVO`

GetDohSetting returns the DohSetting field if non-nil, zero value otherwise.

### GetDohSettingOk

`func (o *DnsProxySettingQueryOpenApiVO) GetDohSettingOk() (*DohSettingOpenApiVO, bool)`

GetDohSettingOk returns a tuple with the DohSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDohSetting

`func (o *DnsProxySettingQueryOpenApiVO) SetDohSetting(v DohSettingOpenApiVO)`

SetDohSetting sets DohSetting field to given value.

### HasDohSetting

`func (o *DnsProxySettingQueryOpenApiVO) HasDohSetting() bool`

HasDohSetting returns a boolean if a field has been set.

### GetDotSetting

`func (o *DnsProxySettingQueryOpenApiVO) GetDotSetting() DotSettingOpenApiVO`

GetDotSetting returns the DotSetting field if non-nil, zero value otherwise.

### GetDotSettingOk

`func (o *DnsProxySettingQueryOpenApiVO) GetDotSettingOk() (*DotSettingOpenApiVO, bool)`

GetDotSettingOk returns a tuple with the DotSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDotSetting

`func (o *DnsProxySettingQueryOpenApiVO) SetDotSetting(v DotSettingOpenApiVO)`

SetDotSetting sets DotSetting field to given value.

### HasDotSetting

`func (o *DnsProxySettingQueryOpenApiVO) HasDotSetting() bool`

HasDotSetting returns a boolean if a field has been set.

### GetEnable

`func (o *DnsProxySettingQueryOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DnsProxySettingQueryOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DnsProxySettingQueryOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DnsProxySettingQueryOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetExistDnsOverride

`func (o *DnsProxySettingQueryOpenApiVO) GetExistDnsOverride() bool`

GetExistDnsOverride returns the ExistDnsOverride field if non-nil, zero value otherwise.

### GetExistDnsOverrideOk

`func (o *DnsProxySettingQueryOpenApiVO) GetExistDnsOverrideOk() (*bool, bool)`

GetExistDnsOverrideOk returns a tuple with the ExistDnsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistDnsOverride

`func (o *DnsProxySettingQueryOpenApiVO) SetExistDnsOverride(v bool)`

SetExistDnsOverride sets ExistDnsOverride field to given value.

### HasExistDnsOverride

`func (o *DnsProxySettingQueryOpenApiVO) HasExistDnsOverride() bool`

HasExistDnsOverride returns a boolean if a field has been set.

### GetSupportDnsOverride

`func (o *DnsProxySettingQueryOpenApiVO) GetSupportDnsOverride() bool`

GetSupportDnsOverride returns the SupportDnsOverride field if non-nil, zero value otherwise.

### GetSupportDnsOverrideOk

`func (o *DnsProxySettingQueryOpenApiVO) GetSupportDnsOverrideOk() (*bool, bool)`

GetSupportDnsOverrideOk returns a tuple with the SupportDnsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDnsOverride

`func (o *DnsProxySettingQueryOpenApiVO) SetSupportDnsOverride(v bool)`

SetSupportDnsOverride sets SupportDnsOverride field to given value.

### HasSupportDnsOverride

`func (o *DnsProxySettingQueryOpenApiVO) HasSupportDnsOverride() bool`

HasSupportDnsOverride returns a boolean if a field has been set.

### GetType

`func (o *DnsProxySettingQueryOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsProxySettingQueryOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsProxySettingQueryOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DnsProxySettingQueryOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


