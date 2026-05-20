# ModifyLanNetworkBrief

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewaySubnet** | **string** | The Gateway Subnet of the lan network | 
**Id** | **string** | The ID of the lan network | 
**IpaddrEnd** | **string** | The ending host IP address of gatewaySubnet | 
**IpaddrStart** | **string** | The starting host IP address of gatewaySubnet | 
**Name** | **string** | The name of the lan network | 
**SiteId** | **string** | The site ID of the lan network | 
**Vlan** | Pointer to **int32** | The vlan number of the lan network | [optional] 

## Methods

### NewModifyLanNetworkBrief

`func NewModifyLanNetworkBrief(gatewaySubnet string, id string, ipaddrEnd string, ipaddrStart string, name string, siteId string, ) *ModifyLanNetworkBrief`

NewModifyLanNetworkBrief instantiates a new ModifyLanNetworkBrief object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyLanNetworkBriefWithDefaults

`func NewModifyLanNetworkBriefWithDefaults() *ModifyLanNetworkBrief`

NewModifyLanNetworkBriefWithDefaults instantiates a new ModifyLanNetworkBrief object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewaySubnet

`func (o *ModifyLanNetworkBrief) GetGatewaySubnet() string`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *ModifyLanNetworkBrief) GetGatewaySubnetOk() (*string, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *ModifyLanNetworkBrief) SetGatewaySubnet(v string)`

SetGatewaySubnet sets GatewaySubnet field to given value.


### GetId

`func (o *ModifyLanNetworkBrief) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifyLanNetworkBrief) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifyLanNetworkBrief) SetId(v string)`

SetId sets Id field to given value.


### GetIpaddrEnd

`func (o *ModifyLanNetworkBrief) GetIpaddrEnd() string`

GetIpaddrEnd returns the IpaddrEnd field if non-nil, zero value otherwise.

### GetIpaddrEndOk

`func (o *ModifyLanNetworkBrief) GetIpaddrEndOk() (*string, bool)`

GetIpaddrEndOk returns a tuple with the IpaddrEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddrEnd

`func (o *ModifyLanNetworkBrief) SetIpaddrEnd(v string)`

SetIpaddrEnd sets IpaddrEnd field to given value.


### GetIpaddrStart

`func (o *ModifyLanNetworkBrief) GetIpaddrStart() string`

GetIpaddrStart returns the IpaddrStart field if non-nil, zero value otherwise.

### GetIpaddrStartOk

`func (o *ModifyLanNetworkBrief) GetIpaddrStartOk() (*string, bool)`

GetIpaddrStartOk returns a tuple with the IpaddrStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddrStart

`func (o *ModifyLanNetworkBrief) SetIpaddrStart(v string)`

SetIpaddrStart sets IpaddrStart field to given value.


### GetName

`func (o *ModifyLanNetworkBrief) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyLanNetworkBrief) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyLanNetworkBrief) SetName(v string)`

SetName sets Name field to given value.


### GetSiteId

`func (o *ModifyLanNetworkBrief) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ModifyLanNetworkBrief) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ModifyLanNetworkBrief) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetVlan

`func (o *ModifyLanNetworkBrief) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *ModifyLanNetworkBrief) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *ModifyLanNetworkBrief) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *ModifyLanNetworkBrief) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


