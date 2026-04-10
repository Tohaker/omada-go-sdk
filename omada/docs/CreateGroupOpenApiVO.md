# CreateGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryList** | Pointer to **[]string** | Country list. [type] value of 5 is required. For the values of countryList, refer to &#39;Country Name&#39; section 5.4.1 of the Open API Access Guide | [optional] 
**Description** | Pointer to **string** | Description should contain 1 to 256 characters. [type] value of 5 is required | [optional] 
**DomainNamePort** | Pointer to [**[]DomainOpenApiVO**](DomainOpenApiVO.md) | Domain info. Handle situations where there are ports, [type] value of 7 is required | [optional] 
**IpList** | Pointer to [**[]IPSubnetsOpenApiVO**](IPSubnetsOpenApiVO.md) | IP subnet info list. [type] value of 0 or 1 is required | [optional] 
**Ipv6List** | Pointer to [**[]IPv6SubnetsOpenApiVO**](IPv6SubnetsOpenApiVO.md) | IPv6 subnet info list. [type] value of 3 or 4 is required | [optional] 
**MacAddressList** | Pointer to [**[]CreateMacAddressOpenApiVO**](CreateMacAddressOpenApiVO.md) | MAC address list. [type] value of 2 is required | [optional] 
**Name** | **string** | Group profile name. Name should contain 1 to 64 characters | 
**PortList** | Pointer to **[]string** | Port list. [portType] value of 0 is required. PortList should be within the range of 0-65535, e.g. 80 or 80-100 | [optional] 
**PortMaskList** | Pointer to [**[]PortMaskOpenApiVO**](PortMaskOpenApiVO.md) | Port mask list. [portType] value of 1 is required | [optional] 
**PortType** | Pointer to **int32** | Port type, portType should be a value as follows: 0: port range 1: port mask. [type] value of 1 or 4 is required | [optional] 
**Type** | **int32** | Type of group profile, type should be a value as follows: 0: IP Group; 1: IP Port Group; 2: MAC Group; 3: IPv6 Group; 4: IPv6 Port Group; 5: Country Group; 7: Domain Group; | 

## Methods

### NewCreateGroupOpenApiVO

`func NewCreateGroupOpenApiVO(name string, type_ int32, ) *CreateGroupOpenApiVO`

NewCreateGroupOpenApiVO instantiates a new CreateGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupOpenApiVOWithDefaults

`func NewCreateGroupOpenApiVOWithDefaults() *CreateGroupOpenApiVO`

NewCreateGroupOpenApiVOWithDefaults instantiates a new CreateGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryList

`func (o *CreateGroupOpenApiVO) GetCountryList() []string`

GetCountryList returns the CountryList field if non-nil, zero value otherwise.

### GetCountryListOk

`func (o *CreateGroupOpenApiVO) GetCountryListOk() (*[]string, bool)`

GetCountryListOk returns a tuple with the CountryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryList

`func (o *CreateGroupOpenApiVO) SetCountryList(v []string)`

SetCountryList sets CountryList field to given value.

### HasCountryList

`func (o *CreateGroupOpenApiVO) HasCountryList() bool`

HasCountryList returns a boolean if a field has been set.

### GetDescription

`func (o *CreateGroupOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGroupOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGroupOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGroupOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomainNamePort

`func (o *CreateGroupOpenApiVO) GetDomainNamePort() []DomainOpenApiVO`

GetDomainNamePort returns the DomainNamePort field if non-nil, zero value otherwise.

### GetDomainNamePortOk

`func (o *CreateGroupOpenApiVO) GetDomainNamePortOk() (*[]DomainOpenApiVO, bool)`

GetDomainNamePortOk returns a tuple with the DomainNamePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNamePort

`func (o *CreateGroupOpenApiVO) SetDomainNamePort(v []DomainOpenApiVO)`

SetDomainNamePort sets DomainNamePort field to given value.

### HasDomainNamePort

`func (o *CreateGroupOpenApiVO) HasDomainNamePort() bool`

HasDomainNamePort returns a boolean if a field has been set.

### GetIpList

`func (o *CreateGroupOpenApiVO) GetIpList() []IPSubnetsOpenApiVO`

GetIpList returns the IpList field if non-nil, zero value otherwise.

### GetIpListOk

`func (o *CreateGroupOpenApiVO) GetIpListOk() (*[]IPSubnetsOpenApiVO, bool)`

GetIpListOk returns a tuple with the IpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpList

`func (o *CreateGroupOpenApiVO) SetIpList(v []IPSubnetsOpenApiVO)`

SetIpList sets IpList field to given value.

### HasIpList

`func (o *CreateGroupOpenApiVO) HasIpList() bool`

HasIpList returns a boolean if a field has been set.

### GetIpv6List

`func (o *CreateGroupOpenApiVO) GetIpv6List() []IPv6SubnetsOpenApiVO`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *CreateGroupOpenApiVO) GetIpv6ListOk() (*[]IPv6SubnetsOpenApiVO, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *CreateGroupOpenApiVO) SetIpv6List(v []IPv6SubnetsOpenApiVO)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *CreateGroupOpenApiVO) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetMacAddressList

`func (o *CreateGroupOpenApiVO) GetMacAddressList() []CreateMacAddressOpenApiVO`

GetMacAddressList returns the MacAddressList field if non-nil, zero value otherwise.

### GetMacAddressListOk

`func (o *CreateGroupOpenApiVO) GetMacAddressListOk() (*[]CreateMacAddressOpenApiVO, bool)`

GetMacAddressListOk returns a tuple with the MacAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressList

`func (o *CreateGroupOpenApiVO) SetMacAddressList(v []CreateMacAddressOpenApiVO)`

SetMacAddressList sets MacAddressList field to given value.

### HasMacAddressList

`func (o *CreateGroupOpenApiVO) HasMacAddressList() bool`

HasMacAddressList returns a boolean if a field has been set.

### GetName

`func (o *CreateGroupOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPortList

`func (o *CreateGroupOpenApiVO) GetPortList() []string`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *CreateGroupOpenApiVO) GetPortListOk() (*[]string, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *CreateGroupOpenApiVO) SetPortList(v []string)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *CreateGroupOpenApiVO) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetPortMaskList

`func (o *CreateGroupOpenApiVO) GetPortMaskList() []PortMaskOpenApiVO`

GetPortMaskList returns the PortMaskList field if non-nil, zero value otherwise.

### GetPortMaskListOk

`func (o *CreateGroupOpenApiVO) GetPortMaskListOk() (*[]PortMaskOpenApiVO, bool)`

GetPortMaskListOk returns a tuple with the PortMaskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMaskList

`func (o *CreateGroupOpenApiVO) SetPortMaskList(v []PortMaskOpenApiVO)`

SetPortMaskList sets PortMaskList field to given value.

### HasPortMaskList

`func (o *CreateGroupOpenApiVO) HasPortMaskList() bool`

HasPortMaskList returns a boolean if a field has been set.

### GetPortType

`func (o *CreateGroupOpenApiVO) GetPortType() int32`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *CreateGroupOpenApiVO) GetPortTypeOk() (*int32, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *CreateGroupOpenApiVO) SetPortType(v int32)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *CreateGroupOpenApiVO) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetType

`func (o *CreateGroupOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateGroupOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateGroupOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


