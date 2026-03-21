# DhcpSettingsConfigTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dhcpns** | Pointer to **string** | Setup DHCP server: \&quot;auto\&quot; or \&quot;manual\&quot; | [optional] 
**Enable** | Pointer to **bool** | When value is true, DHCP server is enabled | [optional] 
**Gateway** | Pointer to **string** | Manual Setup of DHCP Gateway IP | [optional] 
**IpRangePool** | Pointer to [**[]DhcpRangeOpenApiVO**](DhcpRangeOpenApiVO.md) | The list of DHCP Range, which size can&#39;t be more than \&quot;dhcpRangePoolSize\&quot;, \&quot;dhcpRangePoolSize\&quot; can be obtained from &#39;Get LAN network template list&#39; interface. | [optional] 
**Leasetime** | Pointer to **int32** | Leasetime should be within the range of 2-10080 | [optional] 
**Options** | Pointer to [**[]CustomDHCPOptions**](CustomDHCPOptions.md) | User custom DHCP options | [optional] 
**PriDns** | Pointer to **string** | When DHCPs are \&quot;manual\&quot;, primary DNS Server. | [optional] 
**SndDns** | Pointer to **string** | When DHCPs are \&quot;manual\&quot;, second DNS Server. | [optional] 

## Methods

### NewDhcpSettingsConfigTemplateOpenApiVO

`func NewDhcpSettingsConfigTemplateOpenApiVO() *DhcpSettingsConfigTemplateOpenApiVO`

NewDhcpSettingsConfigTemplateOpenApiVO instantiates a new DhcpSettingsConfigTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpSettingsConfigTemplateOpenApiVOWithDefaults

`func NewDhcpSettingsConfigTemplateOpenApiVOWithDefaults() *DhcpSettingsConfigTemplateOpenApiVO`

NewDhcpSettingsConfigTemplateOpenApiVOWithDefaults instantiates a new DhcpSettingsConfigTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpns

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetDhcpns() string`

GetDhcpns returns the Dhcpns field if non-nil, zero value otherwise.

### GetDhcpnsOk

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetDhcpnsOk() (*string, bool)`

GetDhcpnsOk returns a tuple with the Dhcpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpns

`func (o *DhcpSettingsConfigTemplateOpenApiVO) SetDhcpns(v string)`

SetDhcpns sets Dhcpns field to given value.

### HasDhcpns

`func (o *DhcpSettingsConfigTemplateOpenApiVO) HasDhcpns() bool`

HasDhcpns returns a boolean if a field has been set.

### GetEnable

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DhcpSettingsConfigTemplateOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DhcpSettingsConfigTemplateOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetGateway

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DhcpSettingsConfigTemplateOpenApiVO) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DhcpSettingsConfigTemplateOpenApiVO) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpRangePool

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetIpRangePool() []DhcpRangeOpenApiVO`

GetIpRangePool returns the IpRangePool field if non-nil, zero value otherwise.

### GetIpRangePoolOk

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetIpRangePoolOk() (*[]DhcpRangeOpenApiVO, bool)`

GetIpRangePoolOk returns a tuple with the IpRangePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangePool

`func (o *DhcpSettingsConfigTemplateOpenApiVO) SetIpRangePool(v []DhcpRangeOpenApiVO)`

SetIpRangePool sets IpRangePool field to given value.

### HasIpRangePool

`func (o *DhcpSettingsConfigTemplateOpenApiVO) HasIpRangePool() bool`

HasIpRangePool returns a boolean if a field has been set.

### GetLeasetime

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetLeasetime() int32`

GetLeasetime returns the Leasetime field if non-nil, zero value otherwise.

### GetLeasetimeOk

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetLeasetimeOk() (*int32, bool)`

GetLeasetimeOk returns a tuple with the Leasetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasetime

`func (o *DhcpSettingsConfigTemplateOpenApiVO) SetLeasetime(v int32)`

SetLeasetime sets Leasetime field to given value.

### HasLeasetime

`func (o *DhcpSettingsConfigTemplateOpenApiVO) HasLeasetime() bool`

HasLeasetime returns a boolean if a field has been set.

### GetOptions

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetOptions() []CustomDHCPOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetOptionsOk() (*[]CustomDHCPOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DhcpSettingsConfigTemplateOpenApiVO) SetOptions(v []CustomDHCPOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DhcpSettingsConfigTemplateOpenApiVO) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPriDns

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *DhcpSettingsConfigTemplateOpenApiVO) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *DhcpSettingsConfigTemplateOpenApiVO) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *DhcpSettingsConfigTemplateOpenApiVO) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *DhcpSettingsConfigTemplateOpenApiVO) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *DhcpSettingsConfigTemplateOpenApiVO) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


