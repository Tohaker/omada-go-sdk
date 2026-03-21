# DnsSecSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplyPolicy** | **int32** | Bogus DNS reply policy type. ReplyPolicy should be a value as follows: 0: Pass, 1: Drop | 
**Servers** | **[]string** | DNS Server IP list, Up to 2 entries are allowed for the server list | 

## Methods

### NewDnsSecSettingOpenApiVO

`func NewDnsSecSettingOpenApiVO(replyPolicy int32, servers []string, ) *DnsSecSettingOpenApiVO`

NewDnsSecSettingOpenApiVO instantiates a new DnsSecSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsSecSettingOpenApiVOWithDefaults

`func NewDnsSecSettingOpenApiVOWithDefaults() *DnsSecSettingOpenApiVO`

NewDnsSecSettingOpenApiVOWithDefaults instantiates a new DnsSecSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplyPolicy

`func (o *DnsSecSettingOpenApiVO) GetReplyPolicy() int32`

GetReplyPolicy returns the ReplyPolicy field if non-nil, zero value otherwise.

### GetReplyPolicyOk

`func (o *DnsSecSettingOpenApiVO) GetReplyPolicyOk() (*int32, bool)`

GetReplyPolicyOk returns a tuple with the ReplyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyPolicy

`func (o *DnsSecSettingOpenApiVO) SetReplyPolicy(v int32)`

SetReplyPolicy sets ReplyPolicy field to given value.


### GetServers

`func (o *DnsSecSettingOpenApiVO) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DnsSecSettingOpenApiVO) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DnsSecSettingOpenApiVO) SetServers(v []string)`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


