# LanNetworkBrief

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeginIpLong** | Pointer to **int64** | The long value of the beginning IP of gatewaySubnet | [optional] 
**EndIpLong** | Pointer to **int64** | The long value of the ending IP of gatewaySubnet | [optional] 
**GatewaySubnet** | Pointer to **string** | The Gateway Subnet of the lan network | [optional] 
**Id** | Pointer to **string** | The ID of the lan network | [optional] 
**InterfaceIds** | Pointer to **[]string** | A list of the Interface ID for the lan network | [optional] 
**IpRangeEnd** | Pointer to **int64** | The long value of ipaddrEnd | [optional] 
**IpRangeStart** | Pointer to **int64** | The long value of ipaddrStart | [optional] 
**IpaddrEnd** | Pointer to **string** | The ending host IP address of gatewaySubnet | [optional] 
**IpaddrStart** | Pointer to **string** | The starting host IP address of gatewaySubnet | [optional] 
**Name** | Pointer to **string** | The name of the lan network | [optional] 
**Primary** | Pointer to **bool** | Whether is default lan | [optional] 
**Purpose** | Pointer to **string** | The purpose of the lan network | [optional] 
**SiteId** | Pointer to **string** | The site ID of the lan network | [optional] 
**Vlan** | Pointer to **int32** | The vlan number of the lan network | [optional] 
**VlanType** | Pointer to **int32** | The type of vlan of the lan network | [optional] 
**Vlans** | Pointer to **string** | Batch Vlan IDS. | [optional] 

## Methods

### NewLanNetworkBrief

`func NewLanNetworkBrief() *LanNetworkBrief`

NewLanNetworkBrief instantiates a new LanNetworkBrief object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkBriefWithDefaults

`func NewLanNetworkBriefWithDefaults() *LanNetworkBrief`

NewLanNetworkBriefWithDefaults instantiates a new LanNetworkBrief object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeginIpLong

`func (o *LanNetworkBrief) GetBeginIpLong() int64`

GetBeginIpLong returns the BeginIpLong field if non-nil, zero value otherwise.

### GetBeginIpLongOk

`func (o *LanNetworkBrief) GetBeginIpLongOk() (*int64, bool)`

GetBeginIpLongOk returns a tuple with the BeginIpLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginIpLong

`func (o *LanNetworkBrief) SetBeginIpLong(v int64)`

SetBeginIpLong sets BeginIpLong field to given value.

### HasBeginIpLong

`func (o *LanNetworkBrief) HasBeginIpLong() bool`

HasBeginIpLong returns a boolean if a field has been set.

### GetEndIpLong

`func (o *LanNetworkBrief) GetEndIpLong() int64`

GetEndIpLong returns the EndIpLong field if non-nil, zero value otherwise.

### GetEndIpLongOk

`func (o *LanNetworkBrief) GetEndIpLongOk() (*int64, bool)`

GetEndIpLongOk returns a tuple with the EndIpLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIpLong

`func (o *LanNetworkBrief) SetEndIpLong(v int64)`

SetEndIpLong sets EndIpLong field to given value.

### HasEndIpLong

`func (o *LanNetworkBrief) HasEndIpLong() bool`

HasEndIpLong returns a boolean if a field has been set.

### GetGatewaySubnet

`func (o *LanNetworkBrief) GetGatewaySubnet() string`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *LanNetworkBrief) GetGatewaySubnetOk() (*string, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *LanNetworkBrief) SetGatewaySubnet(v string)`

SetGatewaySubnet sets GatewaySubnet field to given value.

### HasGatewaySubnet

`func (o *LanNetworkBrief) HasGatewaySubnet() bool`

HasGatewaySubnet returns a boolean if a field has been set.

### GetId

`func (o *LanNetworkBrief) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanNetworkBrief) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanNetworkBrief) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LanNetworkBrief) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceIds

`func (o *LanNetworkBrief) GetInterfaceIds() []string`

GetInterfaceIds returns the InterfaceIds field if non-nil, zero value otherwise.

### GetInterfaceIdsOk

`func (o *LanNetworkBrief) GetInterfaceIdsOk() (*[]string, bool)`

GetInterfaceIdsOk returns a tuple with the InterfaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIds

`func (o *LanNetworkBrief) SetInterfaceIds(v []string)`

SetInterfaceIds sets InterfaceIds field to given value.

### HasInterfaceIds

`func (o *LanNetworkBrief) HasInterfaceIds() bool`

HasInterfaceIds returns a boolean if a field has been set.

### GetIpRangeEnd

`func (o *LanNetworkBrief) GetIpRangeEnd() int64`

GetIpRangeEnd returns the IpRangeEnd field if non-nil, zero value otherwise.

### GetIpRangeEndOk

`func (o *LanNetworkBrief) GetIpRangeEndOk() (*int64, bool)`

GetIpRangeEndOk returns a tuple with the IpRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangeEnd

`func (o *LanNetworkBrief) SetIpRangeEnd(v int64)`

SetIpRangeEnd sets IpRangeEnd field to given value.

### HasIpRangeEnd

`func (o *LanNetworkBrief) HasIpRangeEnd() bool`

HasIpRangeEnd returns a boolean if a field has been set.

### GetIpRangeStart

`func (o *LanNetworkBrief) GetIpRangeStart() int64`

GetIpRangeStart returns the IpRangeStart field if non-nil, zero value otherwise.

### GetIpRangeStartOk

`func (o *LanNetworkBrief) GetIpRangeStartOk() (*int64, bool)`

GetIpRangeStartOk returns a tuple with the IpRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangeStart

`func (o *LanNetworkBrief) SetIpRangeStart(v int64)`

SetIpRangeStart sets IpRangeStart field to given value.

### HasIpRangeStart

`func (o *LanNetworkBrief) HasIpRangeStart() bool`

HasIpRangeStart returns a boolean if a field has been set.

### GetIpaddrEnd

`func (o *LanNetworkBrief) GetIpaddrEnd() string`

GetIpaddrEnd returns the IpaddrEnd field if non-nil, zero value otherwise.

### GetIpaddrEndOk

`func (o *LanNetworkBrief) GetIpaddrEndOk() (*string, bool)`

GetIpaddrEndOk returns a tuple with the IpaddrEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddrEnd

`func (o *LanNetworkBrief) SetIpaddrEnd(v string)`

SetIpaddrEnd sets IpaddrEnd field to given value.

### HasIpaddrEnd

`func (o *LanNetworkBrief) HasIpaddrEnd() bool`

HasIpaddrEnd returns a boolean if a field has been set.

### GetIpaddrStart

`func (o *LanNetworkBrief) GetIpaddrStart() string`

GetIpaddrStart returns the IpaddrStart field if non-nil, zero value otherwise.

### GetIpaddrStartOk

`func (o *LanNetworkBrief) GetIpaddrStartOk() (*string, bool)`

GetIpaddrStartOk returns a tuple with the IpaddrStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddrStart

`func (o *LanNetworkBrief) SetIpaddrStart(v string)`

SetIpaddrStart sets IpaddrStart field to given value.

### HasIpaddrStart

`func (o *LanNetworkBrief) HasIpaddrStart() bool`

HasIpaddrStart returns a boolean if a field has been set.

### GetName

`func (o *LanNetworkBrief) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanNetworkBrief) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanNetworkBrief) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LanNetworkBrief) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimary

`func (o *LanNetworkBrief) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *LanNetworkBrief) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *LanNetworkBrief) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *LanNetworkBrief) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetPurpose

`func (o *LanNetworkBrief) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *LanNetworkBrief) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *LanNetworkBrief) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *LanNetworkBrief) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSiteId

`func (o *LanNetworkBrief) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *LanNetworkBrief) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *LanNetworkBrief) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *LanNetworkBrief) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetVlan

`func (o *LanNetworkBrief) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LanNetworkBrief) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LanNetworkBrief) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *LanNetworkBrief) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *LanNetworkBrief) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *LanNetworkBrief) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *LanNetworkBrief) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *LanNetworkBrief) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.

### GetVlans

`func (o *LanNetworkBrief) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *LanNetworkBrief) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *LanNetworkBrief) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *LanNetworkBrief) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


