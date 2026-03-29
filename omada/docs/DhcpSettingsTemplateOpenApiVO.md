# DhcpSettingsTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpNextServer** | Pointer to **string** | The dhcpNextServer should be valid IP address, which is used in optional set next DHCP server. | [optional] 
**Dhcpns** | Pointer to **string** | Setup DHCP server: \&quot;auto\&quot; or \&quot;manual\&quot; | [optional] 
**Enable** | Pointer to **bool** | When value is true, DHCP server is enabled | [optional] 
**Gateway** | Pointer to **string** | Manual Setup of DHCP Gateway IP | [optional] 
**IpRangeEnd** | Pointer to **int64** | The specific format value of Gateway Subnet End IP | [optional] 
**IpRangePool** | Pointer to [**[]DhcpRangeOpenApiVO**](DhcpRangeOpenApiVO.md) | The list of DHCP Range, which size can&#39;t be more than \&quot;dhcpRangePoolSize\&quot;, \&quot;dhcpRangePoolSize\&quot; can be obtained from &#39;Get LAN network template list&#39; interface. | [optional] 
**IpRangeStart** | Pointer to **int64** | The specific format value of Gateway Subnet start IP | [optional] 
**Leasetime** | Pointer to **int32** | Leasetime should be within the range of 2-10080 | [optional] 
**Options** | Pointer to [**[]CustomDHCPOptions**](CustomDHCPOptions.md) | User custom DHCP options | [optional] 
**PriDns** | Pointer to **string** | When DHCPs are \&quot;manual\&quot;, primary DNS Server. | [optional] 
**SndDns** | Pointer to **string** | When DHCPs are \&quot;manual\&quot;, second DNS Server. | [optional] 

## Methods

### NewDhcpSettingsTemplateOpenApiVO

`func NewDhcpSettingsTemplateOpenApiVO() *DhcpSettingsTemplateOpenApiVO`

NewDhcpSettingsTemplateOpenApiVO instantiates a new DhcpSettingsTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpSettingsTemplateOpenApiVOWithDefaults

`func NewDhcpSettingsTemplateOpenApiVOWithDefaults() *DhcpSettingsTemplateOpenApiVO`

NewDhcpSettingsTemplateOpenApiVOWithDefaults instantiates a new DhcpSettingsTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpNextServer

`func (o *DhcpSettingsTemplateOpenApiVO) GetDhcpNextServer() string`

GetDhcpNextServer returns the DhcpNextServer field if non-nil, zero value otherwise.

### GetDhcpNextServerOk

`func (o *DhcpSettingsTemplateOpenApiVO) GetDhcpNextServerOk() (*string, bool)`

GetDhcpNextServerOk returns a tuple with the DhcpNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpNextServer

`func (o *DhcpSettingsTemplateOpenApiVO) SetDhcpNextServer(v string)`

SetDhcpNextServer sets DhcpNextServer field to given value.

### HasDhcpNextServer

`func (o *DhcpSettingsTemplateOpenApiVO) HasDhcpNextServer() bool`

HasDhcpNextServer returns a boolean if a field has been set.

### GetDhcpns

`func (o *DhcpSettingsTemplateOpenApiVO) GetDhcpns() string`

GetDhcpns returns the Dhcpns field if non-nil, zero value otherwise.

### GetDhcpnsOk

`func (o *DhcpSettingsTemplateOpenApiVO) GetDhcpnsOk() (*string, bool)`

GetDhcpnsOk returns a tuple with the Dhcpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpns

`func (o *DhcpSettingsTemplateOpenApiVO) SetDhcpns(v string)`

SetDhcpns sets Dhcpns field to given value.

### HasDhcpns

`func (o *DhcpSettingsTemplateOpenApiVO) HasDhcpns() bool`

HasDhcpns returns a boolean if a field has been set.

### GetEnable

`func (o *DhcpSettingsTemplateOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DhcpSettingsTemplateOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DhcpSettingsTemplateOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DhcpSettingsTemplateOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetGateway

`func (o *DhcpSettingsTemplateOpenApiVO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DhcpSettingsTemplateOpenApiVO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DhcpSettingsTemplateOpenApiVO) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DhcpSettingsTemplateOpenApiVO) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpRangeEnd

`func (o *DhcpSettingsTemplateOpenApiVO) GetIpRangeEnd() int64`

GetIpRangeEnd returns the IpRangeEnd field if non-nil, zero value otherwise.

### GetIpRangeEndOk

`func (o *DhcpSettingsTemplateOpenApiVO) GetIpRangeEndOk() (*int64, bool)`

GetIpRangeEndOk returns a tuple with the IpRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangeEnd

`func (o *DhcpSettingsTemplateOpenApiVO) SetIpRangeEnd(v int64)`

SetIpRangeEnd sets IpRangeEnd field to given value.

### HasIpRangeEnd

`func (o *DhcpSettingsTemplateOpenApiVO) HasIpRangeEnd() bool`

HasIpRangeEnd returns a boolean if a field has been set.

### GetIpRangePool

`func (o *DhcpSettingsTemplateOpenApiVO) GetIpRangePool() []DhcpRangeOpenApiVO`

GetIpRangePool returns the IpRangePool field if non-nil, zero value otherwise.

### GetIpRangePoolOk

`func (o *DhcpSettingsTemplateOpenApiVO) GetIpRangePoolOk() (*[]DhcpRangeOpenApiVO, bool)`

GetIpRangePoolOk returns a tuple with the IpRangePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangePool

`func (o *DhcpSettingsTemplateOpenApiVO) SetIpRangePool(v []DhcpRangeOpenApiVO)`

SetIpRangePool sets IpRangePool field to given value.

### HasIpRangePool

`func (o *DhcpSettingsTemplateOpenApiVO) HasIpRangePool() bool`

HasIpRangePool returns a boolean if a field has been set.

### GetIpRangeStart

`func (o *DhcpSettingsTemplateOpenApiVO) GetIpRangeStart() int64`

GetIpRangeStart returns the IpRangeStart field if non-nil, zero value otherwise.

### GetIpRangeStartOk

`func (o *DhcpSettingsTemplateOpenApiVO) GetIpRangeStartOk() (*int64, bool)`

GetIpRangeStartOk returns a tuple with the IpRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangeStart

`func (o *DhcpSettingsTemplateOpenApiVO) SetIpRangeStart(v int64)`

SetIpRangeStart sets IpRangeStart field to given value.

### HasIpRangeStart

`func (o *DhcpSettingsTemplateOpenApiVO) HasIpRangeStart() bool`

HasIpRangeStart returns a boolean if a field has been set.

### GetLeasetime

`func (o *DhcpSettingsTemplateOpenApiVO) GetLeasetime() int32`

GetLeasetime returns the Leasetime field if non-nil, zero value otherwise.

### GetLeasetimeOk

`func (o *DhcpSettingsTemplateOpenApiVO) GetLeasetimeOk() (*int32, bool)`

GetLeasetimeOk returns a tuple with the Leasetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasetime

`func (o *DhcpSettingsTemplateOpenApiVO) SetLeasetime(v int32)`

SetLeasetime sets Leasetime field to given value.

### HasLeasetime

`func (o *DhcpSettingsTemplateOpenApiVO) HasLeasetime() bool`

HasLeasetime returns a boolean if a field has been set.

### GetOptions

`func (o *DhcpSettingsTemplateOpenApiVO) GetOptions() []CustomDHCPOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DhcpSettingsTemplateOpenApiVO) GetOptionsOk() (*[]CustomDHCPOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DhcpSettingsTemplateOpenApiVO) SetOptions(v []CustomDHCPOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DhcpSettingsTemplateOpenApiVO) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPriDns

`func (o *DhcpSettingsTemplateOpenApiVO) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *DhcpSettingsTemplateOpenApiVO) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *DhcpSettingsTemplateOpenApiVO) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *DhcpSettingsTemplateOpenApiVO) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *DhcpSettingsTemplateOpenApiVO) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *DhcpSettingsTemplateOpenApiVO) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *DhcpSettingsTemplateOpenApiVO) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *DhcpSettingsTemplateOpenApiVO) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


