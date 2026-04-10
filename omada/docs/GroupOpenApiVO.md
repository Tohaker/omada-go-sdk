# GroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildIn** | Pointer to **bool** | Is this profile a built-in profile | [optional] 
**Count** | Pointer to **int32** | Count of list entries | [optional] 
**CountryList** | Pointer to **[]string** | Country list. Valid when [type] is 5 | [optional] 
**Description** | Pointer to **string** | Description. Valid when [type] is 5 | [optional] 
**DomainName** | Pointer to **[]string** | Domain name. Valid when [type] is 7 | [optional] 
**DomainNamePort** | Pointer to [**[]DomainOpenApiVO**](DomainOpenApiVO.md) | Domain info. Handle situations where there are ports,  [type] values of 7 is required | [optional] 
**GroupId** | Pointer to **string** | Group profile ID | [optional] 
**IpList** | Pointer to [**[]IPSubnetsOpenApiVO**](IPSubnetsOpenApiVO.md) | IP subnet info list. [type] value of 0 or 1 is required | [optional] 
**Ipv6List** | Pointer to [**[]IPv6SubnetsOpenApiVO**](IPv6SubnetsOpenApiVO.md) | IPv6 subnet info list. [type] value of 3 or 4 is required | [optional] 
**MacAddressList** | Pointer to [**[]MacAddressOpenApiVO**](MacAddressOpenApiVO.md) | MAC address list. Valid when [type] is 2 | [optional] 
**Name** | Pointer to **string** | Group profile name | [optional] 
**PortList** | Pointer to **[]string** | Port list. Valid when [portType] is 0 | [optional] 
**PortMaskList** | Pointer to [**[]PortMaskOpenApiVO**](PortMaskOpenApiVO.md) | Port mask list. [portType] value of 1 are is required | [optional] 
**PortType** | Pointer to **int32** | Port type. 0: port range 1: port mask. Valid when [type] is 1 or 4 | [optional] 
**Type** | Pointer to **int32** | Type of group profile. 0: IP Group; 1: IP Port Group; 2: MAC Group; 3: IPv6 Group; 4: IPv6 Port Group; 5: Country Group; 7: Domain Group; | [optional] 

## Methods

### NewGroupOpenApiVO

`func NewGroupOpenApiVO() *GroupOpenApiVO`

NewGroupOpenApiVO instantiates a new GroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupOpenApiVOWithDefaults

`func NewGroupOpenApiVOWithDefaults() *GroupOpenApiVO`

NewGroupOpenApiVOWithDefaults instantiates a new GroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildIn

`func (o *GroupOpenApiVO) GetBuildIn() bool`

GetBuildIn returns the BuildIn field if non-nil, zero value otherwise.

### GetBuildInOk

`func (o *GroupOpenApiVO) GetBuildInOk() (*bool, bool)`

GetBuildInOk returns a tuple with the BuildIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildIn

`func (o *GroupOpenApiVO) SetBuildIn(v bool)`

SetBuildIn sets BuildIn field to given value.

### HasBuildIn

`func (o *GroupOpenApiVO) HasBuildIn() bool`

HasBuildIn returns a boolean if a field has been set.

### GetCount

`func (o *GroupOpenApiVO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GroupOpenApiVO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GroupOpenApiVO) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GroupOpenApiVO) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCountryList

`func (o *GroupOpenApiVO) GetCountryList() []string`

GetCountryList returns the CountryList field if non-nil, zero value otherwise.

### GetCountryListOk

`func (o *GroupOpenApiVO) GetCountryListOk() (*[]string, bool)`

GetCountryListOk returns a tuple with the CountryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryList

`func (o *GroupOpenApiVO) SetCountryList(v []string)`

SetCountryList sets CountryList field to given value.

### HasCountryList

`func (o *GroupOpenApiVO) HasCountryList() bool`

HasCountryList returns a boolean if a field has been set.

### GetDescription

`func (o *GroupOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomainName

`func (o *GroupOpenApiVO) GetDomainName() []string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *GroupOpenApiVO) GetDomainNameOk() (*[]string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *GroupOpenApiVO) SetDomainName(v []string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *GroupOpenApiVO) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNamePort

`func (o *GroupOpenApiVO) GetDomainNamePort() []DomainOpenApiVO`

GetDomainNamePort returns the DomainNamePort field if non-nil, zero value otherwise.

### GetDomainNamePortOk

`func (o *GroupOpenApiVO) GetDomainNamePortOk() (*[]DomainOpenApiVO, bool)`

GetDomainNamePortOk returns a tuple with the DomainNamePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNamePort

`func (o *GroupOpenApiVO) SetDomainNamePort(v []DomainOpenApiVO)`

SetDomainNamePort sets DomainNamePort field to given value.

### HasDomainNamePort

`func (o *GroupOpenApiVO) HasDomainNamePort() bool`

HasDomainNamePort returns a boolean if a field has been set.

### GetGroupId

`func (o *GroupOpenApiVO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupOpenApiVO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupOpenApiVO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupOpenApiVO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetIpList

`func (o *GroupOpenApiVO) GetIpList() []IPSubnetsOpenApiVO`

GetIpList returns the IpList field if non-nil, zero value otherwise.

### GetIpListOk

`func (o *GroupOpenApiVO) GetIpListOk() (*[]IPSubnetsOpenApiVO, bool)`

GetIpListOk returns a tuple with the IpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpList

`func (o *GroupOpenApiVO) SetIpList(v []IPSubnetsOpenApiVO)`

SetIpList sets IpList field to given value.

### HasIpList

`func (o *GroupOpenApiVO) HasIpList() bool`

HasIpList returns a boolean if a field has been set.

### GetIpv6List

`func (o *GroupOpenApiVO) GetIpv6List() []IPv6SubnetsOpenApiVO`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *GroupOpenApiVO) GetIpv6ListOk() (*[]IPv6SubnetsOpenApiVO, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *GroupOpenApiVO) SetIpv6List(v []IPv6SubnetsOpenApiVO)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *GroupOpenApiVO) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetMacAddressList

`func (o *GroupOpenApiVO) GetMacAddressList() []MacAddressOpenApiVO`

GetMacAddressList returns the MacAddressList field if non-nil, zero value otherwise.

### GetMacAddressListOk

`func (o *GroupOpenApiVO) GetMacAddressListOk() (*[]MacAddressOpenApiVO, bool)`

GetMacAddressListOk returns a tuple with the MacAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressList

`func (o *GroupOpenApiVO) SetMacAddressList(v []MacAddressOpenApiVO)`

SetMacAddressList sets MacAddressList field to given value.

### HasMacAddressList

`func (o *GroupOpenApiVO) HasMacAddressList() bool`

HasMacAddressList returns a boolean if a field has been set.

### GetName

`func (o *GroupOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortList

`func (o *GroupOpenApiVO) GetPortList() []string`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *GroupOpenApiVO) GetPortListOk() (*[]string, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *GroupOpenApiVO) SetPortList(v []string)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *GroupOpenApiVO) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetPortMaskList

`func (o *GroupOpenApiVO) GetPortMaskList() []PortMaskOpenApiVO`

GetPortMaskList returns the PortMaskList field if non-nil, zero value otherwise.

### GetPortMaskListOk

`func (o *GroupOpenApiVO) GetPortMaskListOk() (*[]PortMaskOpenApiVO, bool)`

GetPortMaskListOk returns a tuple with the PortMaskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMaskList

`func (o *GroupOpenApiVO) SetPortMaskList(v []PortMaskOpenApiVO)`

SetPortMaskList sets PortMaskList field to given value.

### HasPortMaskList

`func (o *GroupOpenApiVO) HasPortMaskList() bool`

HasPortMaskList returns a boolean if a field has been set.

### GetPortType

`func (o *GroupOpenApiVO) GetPortType() int32`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *GroupOpenApiVO) GetPortTypeOk() (*int32, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *GroupOpenApiVO) SetPortType(v int32)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *GroupOpenApiVO) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetType

`func (o *GroupOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *GroupOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


