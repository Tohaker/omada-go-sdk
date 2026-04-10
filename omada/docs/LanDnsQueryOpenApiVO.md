# LanDnsQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Name of Lan Dns Item | [optional] 
**Cname** | Pointer to **string** | When Type is CNAME, set the domain name to which Domain Name and Alias Domain Name need to be mapped. | [optional] 
**CustomTtl** | Pointer to **bool** | When custom TTL is activated, TTL will take effect as custom. Otherwise, TTL will take effect as default(3600). | [optional] 
**DnsServers** | Pointer to **[]string** | When the Type is FORWARD, set the Domain Name and Alias Domain Name to be forwarded to a specific DNS Server, up to two DNS Servers can be configured. | [optional] 
**Domain** | **string** | Enter the domain name. | 
**Enable** | **bool** | off:false, on: true | 
**Id** | Pointer to **string** | The ID of current Lan Dns entry. | [optional] 
**IpAddresses** | Pointer to **[]string** | When the Type is IP, it is the IPv4 address of the returned DNS response. | [optional] 
**Ipv6Addresses** | Pointer to **[]string** | When the Type is IP, it is the IPv6 address of the returned DNS response. | [optional] 
**LanNetworkIds** | Pointer to **[]string** | The ids of Lan Network. | [optional] 
**Name** | **string** | Name of Lan Dns Item | 
**Ttl** | Pointer to **int32** | The amount of time DNS information is allowed to be cached before it expires and needs to be refreshed. It is recommended to use the default TTL (3600) for each record. The range of TTL is between 1 and 86400. | [optional] 
**Type** | **int32** | There are three options, IP:0, CNAME:1, and FORWARD:2. | 

## Methods

### NewLanDnsQueryOpenApiVO

`func NewLanDnsQueryOpenApiVO(domain string, enable bool, name string, type_ int32, ) *LanDnsQueryOpenApiVO`

NewLanDnsQueryOpenApiVO instantiates a new LanDnsQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanDnsQueryOpenApiVOWithDefaults

`func NewLanDnsQueryOpenApiVOWithDefaults() *LanDnsQueryOpenApiVO`

NewLanDnsQueryOpenApiVOWithDefaults instantiates a new LanDnsQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *LanDnsQueryOpenApiVO) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *LanDnsQueryOpenApiVO) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *LanDnsQueryOpenApiVO) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *LanDnsQueryOpenApiVO) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetCname

`func (o *LanDnsQueryOpenApiVO) GetCname() string`

GetCname returns the Cname field if non-nil, zero value otherwise.

### GetCnameOk

`func (o *LanDnsQueryOpenApiVO) GetCnameOk() (*string, bool)`

GetCnameOk returns a tuple with the Cname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCname

`func (o *LanDnsQueryOpenApiVO) SetCname(v string)`

SetCname sets Cname field to given value.

### HasCname

`func (o *LanDnsQueryOpenApiVO) HasCname() bool`

HasCname returns a boolean if a field has been set.

### GetCustomTtl

`func (o *LanDnsQueryOpenApiVO) GetCustomTtl() bool`

GetCustomTtl returns the CustomTtl field if non-nil, zero value otherwise.

### GetCustomTtlOk

`func (o *LanDnsQueryOpenApiVO) GetCustomTtlOk() (*bool, bool)`

GetCustomTtlOk returns a tuple with the CustomTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTtl

`func (o *LanDnsQueryOpenApiVO) SetCustomTtl(v bool)`

SetCustomTtl sets CustomTtl field to given value.

### HasCustomTtl

`func (o *LanDnsQueryOpenApiVO) HasCustomTtl() bool`

HasCustomTtl returns a boolean if a field has been set.

### GetDnsServers

`func (o *LanDnsQueryOpenApiVO) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *LanDnsQueryOpenApiVO) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *LanDnsQueryOpenApiVO) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *LanDnsQueryOpenApiVO) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDomain

`func (o *LanDnsQueryOpenApiVO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LanDnsQueryOpenApiVO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LanDnsQueryOpenApiVO) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEnable

`func (o *LanDnsQueryOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *LanDnsQueryOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *LanDnsQueryOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetId

`func (o *LanDnsQueryOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanDnsQueryOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanDnsQueryOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LanDnsQueryOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddresses

`func (o *LanDnsQueryOpenApiVO) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *LanDnsQueryOpenApiVO) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *LanDnsQueryOpenApiVO) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *LanDnsQueryOpenApiVO) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *LanDnsQueryOpenApiVO) GetIpv6Addresses() []string`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *LanDnsQueryOpenApiVO) GetIpv6AddressesOk() (*[]string, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *LanDnsQueryOpenApiVO) SetIpv6Addresses(v []string)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *LanDnsQueryOpenApiVO) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetLanNetworkIds

`func (o *LanDnsQueryOpenApiVO) GetLanNetworkIds() []string`

GetLanNetworkIds returns the LanNetworkIds field if non-nil, zero value otherwise.

### GetLanNetworkIdsOk

`func (o *LanDnsQueryOpenApiVO) GetLanNetworkIdsOk() (*[]string, bool)`

GetLanNetworkIdsOk returns a tuple with the LanNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkIds

`func (o *LanDnsQueryOpenApiVO) SetLanNetworkIds(v []string)`

SetLanNetworkIds sets LanNetworkIds field to given value.

### HasLanNetworkIds

`func (o *LanDnsQueryOpenApiVO) HasLanNetworkIds() bool`

HasLanNetworkIds returns a boolean if a field has been set.

### GetName

`func (o *LanDnsQueryOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanDnsQueryOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanDnsQueryOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetTtl

`func (o *LanDnsQueryOpenApiVO) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *LanDnsQueryOpenApiVO) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *LanDnsQueryOpenApiVO) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *LanDnsQueryOpenApiVO) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *LanDnsQueryOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LanDnsQueryOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LanDnsQueryOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


