# DhcpSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpNextServer** | Pointer to **string** | The dhcpNextServer should be valid IP address, which is used in optional set next DHCP server. | [optional] 
**Dhcpns** | Pointer to **string** | Setup DHCP server: \&quot;auto\&quot; or \&quot;manual\&quot; | [optional] 
**Enable** | Pointer to **bool** | When value is true, DHCP server is enabled | [optional] 
**Gateway** | Pointer to **string** | Manual Setup of DHCP Gateway IP | [optional] 
**HostIP** | Pointer to **string** | DHCP Omada Controller IP | [optional] 
**IpRangeEnd** | Pointer to **int64** | The specific format value of Gateway Subnet End IP | [optional] 
**IpRangeStart** | Pointer to **int64** | The specific format value of Gateway Subnet start IP | [optional] 
**IpaddrEnd** | Pointer to **string** | DHCP Range End IP. Must use ipRangePool field If want to configure multiple DHCP Ranges. | [optional] 
**IpaddrStart** | Pointer to **string** | DHCP Range Start IP. Must use ipRangePool field If want to configure multiple DHCP Ranges. | [optional] 
**Leasetime** | Pointer to **int32** | Valid value is from 2 to 2880 | [optional] 
**Option138** | Pointer to **string** | The option138 should be valid IP address, which is used in discovering the devices by the Omada controller. | [optional] 
**Option60** | Pointer to **string** | Option60 should be between 0 and 50, which is used to optionally identify the vendor type and configuration of a DHCP client. | [optional] 
**Option66** | Pointer to **string** | The Option66 should be between 0 and 128, which specifies the TFTP server information and supports a single TFTP server IP address. | [optional] 
**Options** | Pointer to [**[]CustomDHCPOptions**](CustomDHCPOptions.md) | User custom DHCP options | [optional] 
**PriDns** | Pointer to **string** | When DHCPs are \&quot;manual\&quot;, primary DNS Server. | [optional] 
**SndDns** | Pointer to **string** | When DHCPs are \&quot;manual\&quot;, second DNS Server. | [optional] 

## Methods

### NewDhcpSettings

`func NewDhcpSettings() *DhcpSettings`

NewDhcpSettings instantiates a new DhcpSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpSettingsWithDefaults

`func NewDhcpSettingsWithDefaults() *DhcpSettings`

NewDhcpSettingsWithDefaults instantiates a new DhcpSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpNextServer

`func (o *DhcpSettings) GetDhcpNextServer() string`

GetDhcpNextServer returns the DhcpNextServer field if non-nil, zero value otherwise.

### GetDhcpNextServerOk

`func (o *DhcpSettings) GetDhcpNextServerOk() (*string, bool)`

GetDhcpNextServerOk returns a tuple with the DhcpNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpNextServer

`func (o *DhcpSettings) SetDhcpNextServer(v string)`

SetDhcpNextServer sets DhcpNextServer field to given value.

### HasDhcpNextServer

`func (o *DhcpSettings) HasDhcpNextServer() bool`

HasDhcpNextServer returns a boolean if a field has been set.

### GetDhcpns

`func (o *DhcpSettings) GetDhcpns() string`

GetDhcpns returns the Dhcpns field if non-nil, zero value otherwise.

### GetDhcpnsOk

`func (o *DhcpSettings) GetDhcpnsOk() (*string, bool)`

GetDhcpnsOk returns a tuple with the Dhcpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpns

`func (o *DhcpSettings) SetDhcpns(v string)`

SetDhcpns sets Dhcpns field to given value.

### HasDhcpns

`func (o *DhcpSettings) HasDhcpns() bool`

HasDhcpns returns a boolean if a field has been set.

### GetEnable

`func (o *DhcpSettings) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DhcpSettings) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DhcpSettings) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DhcpSettings) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetGateway

`func (o *DhcpSettings) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DhcpSettings) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DhcpSettings) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DhcpSettings) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetHostIP

`func (o *DhcpSettings) GetHostIP() string`

GetHostIP returns the HostIP field if non-nil, zero value otherwise.

### GetHostIPOk

`func (o *DhcpSettings) GetHostIPOk() (*string, bool)`

GetHostIPOk returns a tuple with the HostIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIP

`func (o *DhcpSettings) SetHostIP(v string)`

SetHostIP sets HostIP field to given value.

### HasHostIP

`func (o *DhcpSettings) HasHostIP() bool`

HasHostIP returns a boolean if a field has been set.

### GetIpRangeEnd

`func (o *DhcpSettings) GetIpRangeEnd() int64`

GetIpRangeEnd returns the IpRangeEnd field if non-nil, zero value otherwise.

### GetIpRangeEndOk

`func (o *DhcpSettings) GetIpRangeEndOk() (*int64, bool)`

GetIpRangeEndOk returns a tuple with the IpRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangeEnd

`func (o *DhcpSettings) SetIpRangeEnd(v int64)`

SetIpRangeEnd sets IpRangeEnd field to given value.

### HasIpRangeEnd

`func (o *DhcpSettings) HasIpRangeEnd() bool`

HasIpRangeEnd returns a boolean if a field has been set.

### GetIpRangeStart

`func (o *DhcpSettings) GetIpRangeStart() int64`

GetIpRangeStart returns the IpRangeStart field if non-nil, zero value otherwise.

### GetIpRangeStartOk

`func (o *DhcpSettings) GetIpRangeStartOk() (*int64, bool)`

GetIpRangeStartOk returns a tuple with the IpRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangeStart

`func (o *DhcpSettings) SetIpRangeStart(v int64)`

SetIpRangeStart sets IpRangeStart field to given value.

### HasIpRangeStart

`func (o *DhcpSettings) HasIpRangeStart() bool`

HasIpRangeStart returns a boolean if a field has been set.

### GetIpaddrEnd

`func (o *DhcpSettings) GetIpaddrEnd() string`

GetIpaddrEnd returns the IpaddrEnd field if non-nil, zero value otherwise.

### GetIpaddrEndOk

`func (o *DhcpSettings) GetIpaddrEndOk() (*string, bool)`

GetIpaddrEndOk returns a tuple with the IpaddrEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddrEnd

`func (o *DhcpSettings) SetIpaddrEnd(v string)`

SetIpaddrEnd sets IpaddrEnd field to given value.

### HasIpaddrEnd

`func (o *DhcpSettings) HasIpaddrEnd() bool`

HasIpaddrEnd returns a boolean if a field has been set.

### GetIpaddrStart

`func (o *DhcpSettings) GetIpaddrStart() string`

GetIpaddrStart returns the IpaddrStart field if non-nil, zero value otherwise.

### GetIpaddrStartOk

`func (o *DhcpSettings) GetIpaddrStartOk() (*string, bool)`

GetIpaddrStartOk returns a tuple with the IpaddrStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddrStart

`func (o *DhcpSettings) SetIpaddrStart(v string)`

SetIpaddrStart sets IpaddrStart field to given value.

### HasIpaddrStart

`func (o *DhcpSettings) HasIpaddrStart() bool`

HasIpaddrStart returns a boolean if a field has been set.

### GetLeasetime

`func (o *DhcpSettings) GetLeasetime() int32`

GetLeasetime returns the Leasetime field if non-nil, zero value otherwise.

### GetLeasetimeOk

`func (o *DhcpSettings) GetLeasetimeOk() (*int32, bool)`

GetLeasetimeOk returns a tuple with the Leasetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasetime

`func (o *DhcpSettings) SetLeasetime(v int32)`

SetLeasetime sets Leasetime field to given value.

### HasLeasetime

`func (o *DhcpSettings) HasLeasetime() bool`

HasLeasetime returns a boolean if a field has been set.

### GetOption138

`func (o *DhcpSettings) GetOption138() string`

GetOption138 returns the Option138 field if non-nil, zero value otherwise.

### GetOption138Ok

`func (o *DhcpSettings) GetOption138Ok() (*string, bool)`

GetOption138Ok returns a tuple with the Option138 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption138

`func (o *DhcpSettings) SetOption138(v string)`

SetOption138 sets Option138 field to given value.

### HasOption138

`func (o *DhcpSettings) HasOption138() bool`

HasOption138 returns a boolean if a field has been set.

### GetOption60

`func (o *DhcpSettings) GetOption60() string`

GetOption60 returns the Option60 field if non-nil, zero value otherwise.

### GetOption60Ok

`func (o *DhcpSettings) GetOption60Ok() (*string, bool)`

GetOption60Ok returns a tuple with the Option60 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption60

`func (o *DhcpSettings) SetOption60(v string)`

SetOption60 sets Option60 field to given value.

### HasOption60

`func (o *DhcpSettings) HasOption60() bool`

HasOption60 returns a boolean if a field has been set.

### GetOption66

`func (o *DhcpSettings) GetOption66() string`

GetOption66 returns the Option66 field if non-nil, zero value otherwise.

### GetOption66Ok

`func (o *DhcpSettings) GetOption66Ok() (*string, bool)`

GetOption66Ok returns a tuple with the Option66 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption66

`func (o *DhcpSettings) SetOption66(v string)`

SetOption66 sets Option66 field to given value.

### HasOption66

`func (o *DhcpSettings) HasOption66() bool`

HasOption66 returns a boolean if a field has been set.

### GetOptions

`func (o *DhcpSettings) GetOptions() []CustomDHCPOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DhcpSettings) GetOptionsOk() (*[]CustomDHCPOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DhcpSettings) SetOptions(v []CustomDHCPOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DhcpSettings) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPriDns

`func (o *DhcpSettings) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *DhcpSettings) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *DhcpSettings) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *DhcpSettings) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *DhcpSettings) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *DhcpSettings) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *DhcpSettings) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *DhcpSettings) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


