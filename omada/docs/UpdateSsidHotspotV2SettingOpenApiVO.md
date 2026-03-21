# UpdateSsidHotspotV2SettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityIpv4** | Pointer to **int32** | Available type information of IPv4 addresses.&lt;br /&gt;Parameter availabilityIpv4 should be a value as follows: [0: Address type not available; 1: Public IPv4 address available; 2: Port-restricted IPv4 address available; 3: Single NATed private IPv4 address available; 4: Double NATed private IPv4 address available; 5: Port-restricted IPv4 address and single NATed IPv4 address available; 6: Port-restricted IPv4 address and double NATed IPv4 address available; 7: Availability of the address type is not known]. | [optional] 
**AvailabilityIpv6** | Pointer to **int32** | Available type information of IPv6 addresses.&lt;br /&gt; Parameter availabilityIpv6 should be a value as follows: [0: Address type not available; 1: Address type available; 2: Availability of the address type not known]. | [optional] 
**DgafDisable** | Pointer to **bool** | Whether to enable DGAF(downstream group-addressed forwarding) disable mode.&lt;br /&gt;In DGAF disable mode, the AP will not forward downstream multicast and broadcast packets. | [optional] 
**HeSsid** | Pointer to **string** | Homogenous Extended Service Set Identifier, it is used to identify the same type of ESS network set.&lt;br /&gt;Note: HESSID should be consistent with one of the BSSIDs of the APs in the zone. | [optional] 
**HotspotV2Enable** | **bool** | Whether Hotspot2.0 is enabled.&lt;br /&gt;If hotspot2.0 is disabled, other parameters in Hotspot2.0 will be invalid. | 
**Internet** | Pointer to **bool** | Internet access support status (network reachability), which indicates that the network is allowed to access the Internet. | [optional] 
**NetworkType** | Pointer to **int32** | Specify the 802.11u network type.&lt;br /&gt; Parameter networkType should be a value as follows: [0: Private network; 1: Private network with guest access; 2: Chargeable public network; 3: Free public network; 4: Personal device network; 5: Emergency services only network; 14: Test or experimental; 15: Wildcard]. | [optional] 
**OperatorDomain** | Pointer to **string** | Enter the domain name of the hotspot operator.&lt;br /&gt;For example, www.omadanetworks.com. | [optional] 
**OperatorFriendly** | Pointer to **string** | Hotspot network operator friendly name.&lt;br /&gt;Note:Parameter operatorFriendly should contain between 1 and 64 visible ASCII characters, with no Spaces at the beginning and end, and Spaces in between. | [optional] 
**PlmnId** | Pointer to [**[]PlmnIdOpenApiVO**](PlmnIdOpenApiVO.md) | PLMN ID list, enter PLMN ID of 802.11u 3GPP cellular network.&lt;br /&gt;Note: Up to 6 entries are allowed for the PLMN ID list. | [optional] 
**RealmList** | Pointer to [**[]RealmOpenApiVO**](RealmOpenApiVO.md) | Add a profile to identify and describe a NAI (Network Access Identifier) realm accessible using the AP, and the method that this NAI realm uses for authentication.&lt;br /&gt;Note: Up to 10 entries are allowed for the NAI Realm list. | [optional] 
**RoamingConsortiumOi** | Pointer to [**[]RoamingConsortiumOiOpenApiVO**](RoamingConsortiumOiOpenApiVO.md) | Roaming Consortium Oi list, enter the 802.11u roaming organization identifiers.&lt;br /&gt;Note: Up to 3 entries are allowed for the Roaming Consortium Oi list. | [optional] 
**VenueInfo** | Pointer to [**VenueInfoOpenApiVO**](VenueInfoOpenApiVO.md) |  | [optional] 

## Methods

### NewUpdateSsidHotspotV2SettingOpenApiVO

`func NewUpdateSsidHotspotV2SettingOpenApiVO(hotspotV2Enable bool, ) *UpdateSsidHotspotV2SettingOpenApiVO`

NewUpdateSsidHotspotV2SettingOpenApiVO instantiates a new UpdateSsidHotspotV2SettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSsidHotspotV2SettingOpenApiVOWithDefaults

`func NewUpdateSsidHotspotV2SettingOpenApiVOWithDefaults() *UpdateSsidHotspotV2SettingOpenApiVO`

NewUpdateSsidHotspotV2SettingOpenApiVOWithDefaults instantiates a new UpdateSsidHotspotV2SettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityIpv4

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetAvailabilityIpv4() int32`

GetAvailabilityIpv4 returns the AvailabilityIpv4 field if non-nil, zero value otherwise.

### GetAvailabilityIpv4Ok

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetAvailabilityIpv4Ok() (*int32, bool)`

GetAvailabilityIpv4Ok returns a tuple with the AvailabilityIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityIpv4

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetAvailabilityIpv4(v int32)`

SetAvailabilityIpv4 sets AvailabilityIpv4 field to given value.

### HasAvailabilityIpv4

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) HasAvailabilityIpv4() bool`

HasAvailabilityIpv4 returns a boolean if a field has been set.

### GetAvailabilityIpv6

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetAvailabilityIpv6() int32`

GetAvailabilityIpv6 returns the AvailabilityIpv6 field if non-nil, zero value otherwise.

### GetAvailabilityIpv6Ok

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetAvailabilityIpv6Ok() (*int32, bool)`

GetAvailabilityIpv6Ok returns a tuple with the AvailabilityIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityIpv6

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetAvailabilityIpv6(v int32)`

SetAvailabilityIpv6 sets AvailabilityIpv6 field to given value.

### HasAvailabilityIpv6

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) HasAvailabilityIpv6() bool`

HasAvailabilityIpv6 returns a boolean if a field has been set.

### GetDgafDisable

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetDgafDisable() bool`

GetDgafDisable returns the DgafDisable field if non-nil, zero value otherwise.

### GetDgafDisableOk

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetDgafDisableOk() (*bool, bool)`

GetDgafDisableOk returns a tuple with the DgafDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDgafDisable

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetDgafDisable(v bool)`

SetDgafDisable sets DgafDisable field to given value.

### HasDgafDisable

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) HasDgafDisable() bool`

HasDgafDisable returns a boolean if a field has been set.

### GetHeSsid

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetHeSsid() string`

GetHeSsid returns the HeSsid field if non-nil, zero value otherwise.

### GetHeSsidOk

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetHeSsidOk() (*string, bool)`

GetHeSsidOk returns a tuple with the HeSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeSsid

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetHeSsid(v string)`

SetHeSsid sets HeSsid field to given value.

### HasHeSsid

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) HasHeSsid() bool`

HasHeSsid returns a boolean if a field has been set.

### GetHotspotV2Enable

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetHotspotV2Enable() bool`

GetHotspotV2Enable returns the HotspotV2Enable field if non-nil, zero value otherwise.

### GetHotspotV2EnableOk

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetHotspotV2EnableOk() (*bool, bool)`

GetHotspotV2EnableOk returns a tuple with the HotspotV2Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotspotV2Enable

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetHotspotV2Enable(v bool)`

SetHotspotV2Enable sets HotspotV2Enable field to given value.


### GetInternet

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetInternet() bool`

GetInternet returns the Internet field if non-nil, zero value otherwise.

### GetInternetOk

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetInternetOk() (*bool, bool)`

GetInternetOk returns a tuple with the Internet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternet

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetInternet(v bool)`

SetInternet sets Internet field to given value.

### HasInternet

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) HasInternet() bool`

HasInternet returns a boolean if a field has been set.

### GetNetworkType

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetOperatorDomain

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetOperatorDomain() string`

GetOperatorDomain returns the OperatorDomain field if non-nil, zero value otherwise.

### GetOperatorDomainOk

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetOperatorDomainOk() (*string, bool)`

GetOperatorDomainOk returns a tuple with the OperatorDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorDomain

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetOperatorDomain(v string)`

SetOperatorDomain sets OperatorDomain field to given value.

### HasOperatorDomain

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) HasOperatorDomain() bool`

HasOperatorDomain returns a boolean if a field has been set.

### GetOperatorFriendly

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetOperatorFriendly() string`

GetOperatorFriendly returns the OperatorFriendly field if non-nil, zero value otherwise.

### GetOperatorFriendlyOk

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetOperatorFriendlyOk() (*string, bool)`

GetOperatorFriendlyOk returns a tuple with the OperatorFriendly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorFriendly

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetOperatorFriendly(v string)`

SetOperatorFriendly sets OperatorFriendly field to given value.

### HasOperatorFriendly

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) HasOperatorFriendly() bool`

HasOperatorFriendly returns a boolean if a field has been set.

### GetPlmnId

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetPlmnId() []PlmnIdOpenApiVO`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetPlmnIdOk() (*[]PlmnIdOpenApiVO, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetPlmnId(v []PlmnIdOpenApiVO)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetRealmList

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetRealmList() []RealmOpenApiVO`

GetRealmList returns the RealmList field if non-nil, zero value otherwise.

### GetRealmListOk

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetRealmListOk() (*[]RealmOpenApiVO, bool)`

GetRealmListOk returns a tuple with the RealmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmList

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetRealmList(v []RealmOpenApiVO)`

SetRealmList sets RealmList field to given value.

### HasRealmList

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) HasRealmList() bool`

HasRealmList returns a boolean if a field has been set.

### GetRoamingConsortiumOi

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetRoamingConsortiumOi() []RoamingConsortiumOiOpenApiVO`

GetRoamingConsortiumOi returns the RoamingConsortiumOi field if non-nil, zero value otherwise.

### GetRoamingConsortiumOiOk

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetRoamingConsortiumOiOk() (*[]RoamingConsortiumOiOpenApiVO, bool)`

GetRoamingConsortiumOiOk returns a tuple with the RoamingConsortiumOi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingConsortiumOi

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetRoamingConsortiumOi(v []RoamingConsortiumOiOpenApiVO)`

SetRoamingConsortiumOi sets RoamingConsortiumOi field to given value.

### HasRoamingConsortiumOi

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) HasRoamingConsortiumOi() bool`

HasRoamingConsortiumOi returns a boolean if a field has been set.

### GetVenueInfo

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetVenueInfo() VenueInfoOpenApiVO`

GetVenueInfo returns the VenueInfo field if non-nil, zero value otherwise.

### GetVenueInfoOk

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) GetVenueInfoOk() (*VenueInfoOpenApiVO, bool)`

GetVenueInfoOk returns a tuple with the VenueInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueInfo

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) SetVenueInfo(v VenueInfoOpenApiVO)`

SetVenueInfo sets VenueInfo field to given value.

### HasVenueInfo

`func (o *UpdateSsidHotspotV2SettingOpenApiVO) HasVenueInfo() bool`

HasVenueInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


