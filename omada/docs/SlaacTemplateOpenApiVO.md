# SlaacTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnsv6** | Pointer to **int32** | DHCP Name Server, should be a value as follows: 0: \&quot;auto\&quot;; 1: \&quot;manual\&quot; | [optional] 
**PreId** | Pointer to **int32** | Prefix ID should be within the range of 0-127 | [optional] 
**Prefix** | Pointer to **string** | Address prefix | [optional] 
**PriDns** | Pointer to **string** | Primary DHCP Name Server, only effective for DNSv6 \&quot;manual\&quot; | [optional] 
**SndDns** | Pointer to **string** | Secondary DHCP Name Server, only effective for dnsv6 \&quot;manual\&quot; | [optional] 

## Methods

### NewSlaacTemplateOpenApiVO

`func NewSlaacTemplateOpenApiVO() *SlaacTemplateOpenApiVO`

NewSlaacTemplateOpenApiVO instantiates a new SlaacTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlaacTemplateOpenApiVOWithDefaults

`func NewSlaacTemplateOpenApiVOWithDefaults() *SlaacTemplateOpenApiVO`

NewSlaacTemplateOpenApiVOWithDefaults instantiates a new SlaacTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsv6

`func (o *SlaacTemplateOpenApiVO) GetDnsv6() int32`

GetDnsv6 returns the Dnsv6 field if non-nil, zero value otherwise.

### GetDnsv6Ok

`func (o *SlaacTemplateOpenApiVO) GetDnsv6Ok() (*int32, bool)`

GetDnsv6Ok returns a tuple with the Dnsv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsv6

`func (o *SlaacTemplateOpenApiVO) SetDnsv6(v int32)`

SetDnsv6 sets Dnsv6 field to given value.

### HasDnsv6

`func (o *SlaacTemplateOpenApiVO) HasDnsv6() bool`

HasDnsv6 returns a boolean if a field has been set.

### GetPreId

`func (o *SlaacTemplateOpenApiVO) GetPreId() int32`

GetPreId returns the PreId field if non-nil, zero value otherwise.

### GetPreIdOk

`func (o *SlaacTemplateOpenApiVO) GetPreIdOk() (*int32, bool)`

GetPreIdOk returns a tuple with the PreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreId

`func (o *SlaacTemplateOpenApiVO) SetPreId(v int32)`

SetPreId sets PreId field to given value.

### HasPreId

`func (o *SlaacTemplateOpenApiVO) HasPreId() bool`

HasPreId returns a boolean if a field has been set.

### GetPrefix

`func (o *SlaacTemplateOpenApiVO) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *SlaacTemplateOpenApiVO) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *SlaacTemplateOpenApiVO) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *SlaacTemplateOpenApiVO) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPriDns

`func (o *SlaacTemplateOpenApiVO) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *SlaacTemplateOpenApiVO) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *SlaacTemplateOpenApiVO) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *SlaacTemplateOpenApiVO) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *SlaacTemplateOpenApiVO) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *SlaacTemplateOpenApiVO) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *SlaacTemplateOpenApiVO) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *SlaacTemplateOpenApiVO) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


