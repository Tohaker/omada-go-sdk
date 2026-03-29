# DhcpSettingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpNextServer** | Pointer to **string** | The dhcpNextServer should be valid IP address, which is used in optional set next DHCP server. | [optional] 
**Dhcpns** | Pointer to **string** | Setup DHCP server: \&quot;auto\&quot; or \&quot;manual\&quot; | [optional] 
**Enable** | Pointer to **bool** | When value is true, DHCP server is enabled | [optional] 
**Gateway** | Pointer to **string** | Manual Setup of DHCP Gateway IP | [optional] 
**IpRangePool** | Pointer to [**[]DhcpRangeOpenApiVO**](DhcpRangeOpenApiVO.md) | The list of DHCP Range, which size can&#39;t be more than \&quot;dhcpRangePoolSize\&quot;, \&quot;dhcpRangePoolSize\&quot; can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**Leasetime** | Pointer to **int32** | Leasetime should be within the range of 2-10080 | [optional] 
**Options** | Pointer to [**[]CustomDHCPOptions**](CustomDHCPOptions.md) | User custom DHCP options | [optional] 
**PriDns** | Pointer to **string** | When DHCPs are \&quot;manual\&quot;, primary DNS Server. | [optional] 
**SndDns** | Pointer to **string** | When DHCPs are \&quot;manual\&quot;, second DNS Server. | [optional] 

## Methods

### NewDhcpSettingConfig

`func NewDhcpSettingConfig() *DhcpSettingConfig`

NewDhcpSettingConfig instantiates a new DhcpSettingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpSettingConfigWithDefaults

`func NewDhcpSettingConfigWithDefaults() *DhcpSettingConfig`

NewDhcpSettingConfigWithDefaults instantiates a new DhcpSettingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpNextServer

`func (o *DhcpSettingConfig) GetDhcpNextServer() string`

GetDhcpNextServer returns the DhcpNextServer field if non-nil, zero value otherwise.

### GetDhcpNextServerOk

`func (o *DhcpSettingConfig) GetDhcpNextServerOk() (*string, bool)`

GetDhcpNextServerOk returns a tuple with the DhcpNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpNextServer

`func (o *DhcpSettingConfig) SetDhcpNextServer(v string)`

SetDhcpNextServer sets DhcpNextServer field to given value.

### HasDhcpNextServer

`func (o *DhcpSettingConfig) HasDhcpNextServer() bool`

HasDhcpNextServer returns a boolean if a field has been set.

### GetDhcpns

`func (o *DhcpSettingConfig) GetDhcpns() string`

GetDhcpns returns the Dhcpns field if non-nil, zero value otherwise.

### GetDhcpnsOk

`func (o *DhcpSettingConfig) GetDhcpnsOk() (*string, bool)`

GetDhcpnsOk returns a tuple with the Dhcpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpns

`func (o *DhcpSettingConfig) SetDhcpns(v string)`

SetDhcpns sets Dhcpns field to given value.

### HasDhcpns

`func (o *DhcpSettingConfig) HasDhcpns() bool`

HasDhcpns returns a boolean if a field has been set.

### GetEnable

`func (o *DhcpSettingConfig) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DhcpSettingConfig) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DhcpSettingConfig) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DhcpSettingConfig) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetGateway

`func (o *DhcpSettingConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DhcpSettingConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DhcpSettingConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DhcpSettingConfig) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpRangePool

`func (o *DhcpSettingConfig) GetIpRangePool() []DhcpRangeOpenApiVO`

GetIpRangePool returns the IpRangePool field if non-nil, zero value otherwise.

### GetIpRangePoolOk

`func (o *DhcpSettingConfig) GetIpRangePoolOk() (*[]DhcpRangeOpenApiVO, bool)`

GetIpRangePoolOk returns a tuple with the IpRangePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangePool

`func (o *DhcpSettingConfig) SetIpRangePool(v []DhcpRangeOpenApiVO)`

SetIpRangePool sets IpRangePool field to given value.

### HasIpRangePool

`func (o *DhcpSettingConfig) HasIpRangePool() bool`

HasIpRangePool returns a boolean if a field has been set.

### GetLeasetime

`func (o *DhcpSettingConfig) GetLeasetime() int32`

GetLeasetime returns the Leasetime field if non-nil, zero value otherwise.

### GetLeasetimeOk

`func (o *DhcpSettingConfig) GetLeasetimeOk() (*int32, bool)`

GetLeasetimeOk returns a tuple with the Leasetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasetime

`func (o *DhcpSettingConfig) SetLeasetime(v int32)`

SetLeasetime sets Leasetime field to given value.

### HasLeasetime

`func (o *DhcpSettingConfig) HasLeasetime() bool`

HasLeasetime returns a boolean if a field has been set.

### GetOptions

`func (o *DhcpSettingConfig) GetOptions() []CustomDHCPOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DhcpSettingConfig) GetOptionsOk() (*[]CustomDHCPOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DhcpSettingConfig) SetOptions(v []CustomDHCPOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DhcpSettingConfig) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPriDns

`func (o *DhcpSettingConfig) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *DhcpSettingConfig) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *DhcpSettingConfig) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *DhcpSettingConfig) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *DhcpSettingConfig) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *DhcpSettingConfig) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *DhcpSettingConfig) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *DhcpSettingConfig) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


