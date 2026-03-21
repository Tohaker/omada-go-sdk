# DnsOverrideSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyNetwork** | Pointer to **[]string** | Specify the effective LAN network to apply DNS Override. Available LAN networks can be obtained from &#39;Get all \&quot;single\&quot;-\&quot;multi\&quot; interface lan network&#39; | [optional] 
**PrimaryDnsServer** | **string** | Specify the primary upstream DNS server. | 
**SecondaryDnsServer** | Pointer to **string** | Specify the secondary upstream DNS server. | [optional] 

## Methods

### NewDnsOverrideSettingOpenApiVO

`func NewDnsOverrideSettingOpenApiVO(primaryDnsServer string, ) *DnsOverrideSettingOpenApiVO`

NewDnsOverrideSettingOpenApiVO instantiates a new DnsOverrideSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsOverrideSettingOpenApiVOWithDefaults

`func NewDnsOverrideSettingOpenApiVOWithDefaults() *DnsOverrideSettingOpenApiVO`

NewDnsOverrideSettingOpenApiVOWithDefaults instantiates a new DnsOverrideSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyNetwork

`func (o *DnsOverrideSettingOpenApiVO) GetApplyNetwork() []string`

GetApplyNetwork returns the ApplyNetwork field if non-nil, zero value otherwise.

### GetApplyNetworkOk

`func (o *DnsOverrideSettingOpenApiVO) GetApplyNetworkOk() (*[]string, bool)`

GetApplyNetworkOk returns a tuple with the ApplyNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyNetwork

`func (o *DnsOverrideSettingOpenApiVO) SetApplyNetwork(v []string)`

SetApplyNetwork sets ApplyNetwork field to given value.

### HasApplyNetwork

`func (o *DnsOverrideSettingOpenApiVO) HasApplyNetwork() bool`

HasApplyNetwork returns a boolean if a field has been set.

### GetPrimaryDnsServer

`func (o *DnsOverrideSettingOpenApiVO) GetPrimaryDnsServer() string`

GetPrimaryDnsServer returns the PrimaryDnsServer field if non-nil, zero value otherwise.

### GetPrimaryDnsServerOk

`func (o *DnsOverrideSettingOpenApiVO) GetPrimaryDnsServerOk() (*string, bool)`

GetPrimaryDnsServerOk returns a tuple with the PrimaryDnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDnsServer

`func (o *DnsOverrideSettingOpenApiVO) SetPrimaryDnsServer(v string)`

SetPrimaryDnsServer sets PrimaryDnsServer field to given value.


### GetSecondaryDnsServer

`func (o *DnsOverrideSettingOpenApiVO) GetSecondaryDnsServer() string`

GetSecondaryDnsServer returns the SecondaryDnsServer field if non-nil, zero value otherwise.

### GetSecondaryDnsServerOk

`func (o *DnsOverrideSettingOpenApiVO) GetSecondaryDnsServerOk() (*string, bool)`

GetSecondaryDnsServerOk returns a tuple with the SecondaryDnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDnsServer

`func (o *DnsOverrideSettingOpenApiVO) SetSecondaryDnsServer(v string)`

SetSecondaryDnsServer sets SecondaryDnsServer field to given value.

### HasSecondaryDnsServer

`func (o *DnsOverrideSettingOpenApiVO) HasSecondaryDnsServer() bool`

HasSecondaryDnsServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


