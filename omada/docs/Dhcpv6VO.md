# Dhcpv6VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnsv6** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**IpRangeEnd** | Pointer to **string** |  | [optional] 
**IpRangeStart** | Pointer to **string** |  | [optional] 
**IpaddrEnd** | Pointer to **string** |  | [optional] 
**IpaddrStart** | Pointer to **string** |  | [optional] 
**Leasetime** | Pointer to **int32** |  | [optional] 
**PriDns** | Pointer to **string** |  | [optional] 
**SndDns** | Pointer to **string** |  | [optional] 
**Subnet** | Pointer to **int32** |  | [optional] 

## Methods

### NewDhcpv6VO

`func NewDhcpv6VO() *Dhcpv6VO`

NewDhcpv6VO instantiates a new Dhcpv6VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpv6VOWithDefaults

`func NewDhcpv6VOWithDefaults() *Dhcpv6VO`

NewDhcpv6VOWithDefaults instantiates a new Dhcpv6VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsv6

`func (o *Dhcpv6VO) GetDnsv6() string`

GetDnsv6 returns the Dnsv6 field if non-nil, zero value otherwise.

### GetDnsv6Ok

`func (o *Dhcpv6VO) GetDnsv6Ok() (*string, bool)`

GetDnsv6Ok returns a tuple with the Dnsv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsv6

`func (o *Dhcpv6VO) SetDnsv6(v string)`

SetDnsv6 sets Dnsv6 field to given value.

### HasDnsv6

`func (o *Dhcpv6VO) HasDnsv6() bool`

HasDnsv6 returns a boolean if a field has been set.

### GetGateway

`func (o *Dhcpv6VO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Dhcpv6VO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Dhcpv6VO) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Dhcpv6VO) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpRangeEnd

`func (o *Dhcpv6VO) GetIpRangeEnd() string`

GetIpRangeEnd returns the IpRangeEnd field if non-nil, zero value otherwise.

### GetIpRangeEndOk

`func (o *Dhcpv6VO) GetIpRangeEndOk() (*string, bool)`

GetIpRangeEndOk returns a tuple with the IpRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangeEnd

`func (o *Dhcpv6VO) SetIpRangeEnd(v string)`

SetIpRangeEnd sets IpRangeEnd field to given value.

### HasIpRangeEnd

`func (o *Dhcpv6VO) HasIpRangeEnd() bool`

HasIpRangeEnd returns a boolean if a field has been set.

### GetIpRangeStart

`func (o *Dhcpv6VO) GetIpRangeStart() string`

GetIpRangeStart returns the IpRangeStart field if non-nil, zero value otherwise.

### GetIpRangeStartOk

`func (o *Dhcpv6VO) GetIpRangeStartOk() (*string, bool)`

GetIpRangeStartOk returns a tuple with the IpRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangeStart

`func (o *Dhcpv6VO) SetIpRangeStart(v string)`

SetIpRangeStart sets IpRangeStart field to given value.

### HasIpRangeStart

`func (o *Dhcpv6VO) HasIpRangeStart() bool`

HasIpRangeStart returns a boolean if a field has been set.

### GetIpaddrEnd

`func (o *Dhcpv6VO) GetIpaddrEnd() string`

GetIpaddrEnd returns the IpaddrEnd field if non-nil, zero value otherwise.

### GetIpaddrEndOk

`func (o *Dhcpv6VO) GetIpaddrEndOk() (*string, bool)`

GetIpaddrEndOk returns a tuple with the IpaddrEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddrEnd

`func (o *Dhcpv6VO) SetIpaddrEnd(v string)`

SetIpaddrEnd sets IpaddrEnd field to given value.

### HasIpaddrEnd

`func (o *Dhcpv6VO) HasIpaddrEnd() bool`

HasIpaddrEnd returns a boolean if a field has been set.

### GetIpaddrStart

`func (o *Dhcpv6VO) GetIpaddrStart() string`

GetIpaddrStart returns the IpaddrStart field if non-nil, zero value otherwise.

### GetIpaddrStartOk

`func (o *Dhcpv6VO) GetIpaddrStartOk() (*string, bool)`

GetIpaddrStartOk returns a tuple with the IpaddrStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddrStart

`func (o *Dhcpv6VO) SetIpaddrStart(v string)`

SetIpaddrStart sets IpaddrStart field to given value.

### HasIpaddrStart

`func (o *Dhcpv6VO) HasIpaddrStart() bool`

HasIpaddrStart returns a boolean if a field has been set.

### GetLeasetime

`func (o *Dhcpv6VO) GetLeasetime() int32`

GetLeasetime returns the Leasetime field if non-nil, zero value otherwise.

### GetLeasetimeOk

`func (o *Dhcpv6VO) GetLeasetimeOk() (*int32, bool)`

GetLeasetimeOk returns a tuple with the Leasetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasetime

`func (o *Dhcpv6VO) SetLeasetime(v int32)`

SetLeasetime sets Leasetime field to given value.

### HasLeasetime

`func (o *Dhcpv6VO) HasLeasetime() bool`

HasLeasetime returns a boolean if a field has been set.

### GetPriDns

`func (o *Dhcpv6VO) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *Dhcpv6VO) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *Dhcpv6VO) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *Dhcpv6VO) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *Dhcpv6VO) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *Dhcpv6VO) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *Dhcpv6VO) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *Dhcpv6VO) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.

### GetSubnet

`func (o *Dhcpv6VO) GetSubnet() int32`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *Dhcpv6VO) GetSubnetOk() (*int32, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *Dhcpv6VO) SetSubnet(v int32)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *Dhcpv6VO) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


