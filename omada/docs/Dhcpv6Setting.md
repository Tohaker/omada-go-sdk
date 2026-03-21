# Dhcpv6Setting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnsv6** | Pointer to **int32** | DHCP Name Server, should be a value as follows: 0: \&quot;auto\&quot;; 1: \&quot;manual\&quot; | [optional] 
**Gateway** | Pointer to **string** | Gateway IPv6 | [optional] 
**IpaddrEnd** | Pointer to **string** | DHCP Range End | [optional] 
**IpaddrStart** | Pointer to **string** | DHCP Range Start | [optional] 
**Leasetime** | Pointer to **int32** | Leasetime should be within the range of 1-11520. Time unit: min(s) | [optional] 
**PriDns** | Pointer to **string** | When DNSv6 is 1: \&quot;manual\&quot;, primary DNS Server | [optional] 
**SndDns** | Pointer to **string** | When DNSv6 is 1: \&quot;manual\&quot;, second DNS Server | [optional] 
**Subnet** | Pointer to **int32** | Netmask IPv6 | [optional] 

## Methods

### NewDhcpv6Setting

`func NewDhcpv6Setting() *Dhcpv6Setting`

NewDhcpv6Setting instantiates a new Dhcpv6Setting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpv6SettingWithDefaults

`func NewDhcpv6SettingWithDefaults() *Dhcpv6Setting`

NewDhcpv6SettingWithDefaults instantiates a new Dhcpv6Setting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsv6

`func (o *Dhcpv6Setting) GetDnsv6() int32`

GetDnsv6 returns the Dnsv6 field if non-nil, zero value otherwise.

### GetDnsv6Ok

`func (o *Dhcpv6Setting) GetDnsv6Ok() (*int32, bool)`

GetDnsv6Ok returns a tuple with the Dnsv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsv6

`func (o *Dhcpv6Setting) SetDnsv6(v int32)`

SetDnsv6 sets Dnsv6 field to given value.

### HasDnsv6

`func (o *Dhcpv6Setting) HasDnsv6() bool`

HasDnsv6 returns a boolean if a field has been set.

### GetGateway

`func (o *Dhcpv6Setting) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Dhcpv6Setting) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Dhcpv6Setting) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Dhcpv6Setting) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpaddrEnd

`func (o *Dhcpv6Setting) GetIpaddrEnd() string`

GetIpaddrEnd returns the IpaddrEnd field if non-nil, zero value otherwise.

### GetIpaddrEndOk

`func (o *Dhcpv6Setting) GetIpaddrEndOk() (*string, bool)`

GetIpaddrEndOk returns a tuple with the IpaddrEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddrEnd

`func (o *Dhcpv6Setting) SetIpaddrEnd(v string)`

SetIpaddrEnd sets IpaddrEnd field to given value.

### HasIpaddrEnd

`func (o *Dhcpv6Setting) HasIpaddrEnd() bool`

HasIpaddrEnd returns a boolean if a field has been set.

### GetIpaddrStart

`func (o *Dhcpv6Setting) GetIpaddrStart() string`

GetIpaddrStart returns the IpaddrStart field if non-nil, zero value otherwise.

### GetIpaddrStartOk

`func (o *Dhcpv6Setting) GetIpaddrStartOk() (*string, bool)`

GetIpaddrStartOk returns a tuple with the IpaddrStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddrStart

`func (o *Dhcpv6Setting) SetIpaddrStart(v string)`

SetIpaddrStart sets IpaddrStart field to given value.

### HasIpaddrStart

`func (o *Dhcpv6Setting) HasIpaddrStart() bool`

HasIpaddrStart returns a boolean if a field has been set.

### GetLeasetime

`func (o *Dhcpv6Setting) GetLeasetime() int32`

GetLeasetime returns the Leasetime field if non-nil, zero value otherwise.

### GetLeasetimeOk

`func (o *Dhcpv6Setting) GetLeasetimeOk() (*int32, bool)`

GetLeasetimeOk returns a tuple with the Leasetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasetime

`func (o *Dhcpv6Setting) SetLeasetime(v int32)`

SetLeasetime sets Leasetime field to given value.

### HasLeasetime

`func (o *Dhcpv6Setting) HasLeasetime() bool`

HasLeasetime returns a boolean if a field has been set.

### GetPriDns

`func (o *Dhcpv6Setting) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *Dhcpv6Setting) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *Dhcpv6Setting) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *Dhcpv6Setting) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *Dhcpv6Setting) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *Dhcpv6Setting) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *Dhcpv6Setting) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *Dhcpv6Setting) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.

### GetSubnet

`func (o *Dhcpv6Setting) GetSubnet() int32`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *Dhcpv6Setting) GetSubnetOk() (*int32, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *Dhcpv6Setting) SetSubnet(v int32)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *Dhcpv6Setting) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


