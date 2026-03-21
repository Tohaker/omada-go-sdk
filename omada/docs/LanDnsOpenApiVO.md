# LanDnsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Name of Lan Dns Item | [optional] 
**Cname** | Pointer to **string** | When Type is CNAME, set the domain name to which Domain Name and Alias Domain Name need to be mapped. | [optional] 
**DnsServers** | Pointer to **[]string** | When the Type is FORWARD, set the Domain Name and Alias Domain Name to be forwarded to a specific DNS Server, up to two DNS Servers can be configured. | [optional] 
**Domain** | **string** | Enter the domain name. | 
**Enable** | **bool** | off:false, on: true | 
**Id** | Pointer to **string** |  | [optional] 
**IpAddresses** | Pointer to **[]string** | When the Type is IP, it is the IPv4 address of the returned DNS response. | [optional] 
**Ipv6Addresses** | Pointer to **[]string** | When the Type is IP, it is the IPv6 address of the returned DNS response. | [optional] 
**LanNetworkIds** | Pointer to **[]string** | The ids of Lan Network. | [optional] 
**Name** | **string** | Name of Lan Dns Item | 
**Type** | **int32** | There are three options, IP:0, CNAME:1, and FORWARD:2. | 

## Methods

### NewLanDnsOpenApiVO

`func NewLanDnsOpenApiVO(domain string, enable bool, name string, type_ int32, ) *LanDnsOpenApiVO`

NewLanDnsOpenApiVO instantiates a new LanDnsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanDnsOpenApiVOWithDefaults

`func NewLanDnsOpenApiVOWithDefaults() *LanDnsOpenApiVO`

NewLanDnsOpenApiVOWithDefaults instantiates a new LanDnsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *LanDnsOpenApiVO) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *LanDnsOpenApiVO) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *LanDnsOpenApiVO) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *LanDnsOpenApiVO) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetCname

`func (o *LanDnsOpenApiVO) GetCname() string`

GetCname returns the Cname field if non-nil, zero value otherwise.

### GetCnameOk

`func (o *LanDnsOpenApiVO) GetCnameOk() (*string, bool)`

GetCnameOk returns a tuple with the Cname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCname

`func (o *LanDnsOpenApiVO) SetCname(v string)`

SetCname sets Cname field to given value.

### HasCname

`func (o *LanDnsOpenApiVO) HasCname() bool`

HasCname returns a boolean if a field has been set.

### GetDnsServers

`func (o *LanDnsOpenApiVO) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *LanDnsOpenApiVO) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *LanDnsOpenApiVO) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *LanDnsOpenApiVO) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDomain

`func (o *LanDnsOpenApiVO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LanDnsOpenApiVO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LanDnsOpenApiVO) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEnable

`func (o *LanDnsOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *LanDnsOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *LanDnsOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetId

`func (o *LanDnsOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanDnsOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanDnsOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LanDnsOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddresses

`func (o *LanDnsOpenApiVO) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *LanDnsOpenApiVO) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *LanDnsOpenApiVO) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *LanDnsOpenApiVO) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *LanDnsOpenApiVO) GetIpv6Addresses() []string`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *LanDnsOpenApiVO) GetIpv6AddressesOk() (*[]string, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *LanDnsOpenApiVO) SetIpv6Addresses(v []string)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *LanDnsOpenApiVO) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetLanNetworkIds

`func (o *LanDnsOpenApiVO) GetLanNetworkIds() []string`

GetLanNetworkIds returns the LanNetworkIds field if non-nil, zero value otherwise.

### GetLanNetworkIdsOk

`func (o *LanDnsOpenApiVO) GetLanNetworkIdsOk() (*[]string, bool)`

GetLanNetworkIdsOk returns a tuple with the LanNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkIds

`func (o *LanDnsOpenApiVO) SetLanNetworkIds(v []string)`

SetLanNetworkIds sets LanNetworkIds field to given value.

### HasLanNetworkIds

`func (o *LanDnsOpenApiVO) HasLanNetworkIds() bool`

HasLanNetworkIds returns a boolean if a field has been set.

### GetName

`func (o *LanDnsOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanDnsOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanDnsOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *LanDnsOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LanDnsOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LanDnsOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


