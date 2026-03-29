# ModifyIPSAllowListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **int32** | The location of the target that can trigger the threat, direction should be a value as follows:0: both 1: source, 2: destination. | [optional] 
**Id** | **string** | Allow list entry ID. | 
**IpAddress** | Pointer to **string** | The value of the trafficType is 0, indicating IP address. | [optional] 
**Network** | Pointer to **string** | The value of the trafficType is 1, indicating LAN network ID. | [optional] 
**Subnet** | Pointer to **string** | The value of the trafficType is 2, indicating subnet. | [optional] 
**TrafficType** | Pointer to **int32** | Exempt the category of objects (targets) that can trigger the threat, trafficType should be a value as follows: 0: IP Address, 1: Network, 2: Subnet. | [optional] 

## Methods

### NewModifyIPSAllowListEntry

`func NewModifyIPSAllowListEntry(id string, ) *ModifyIPSAllowListEntry`

NewModifyIPSAllowListEntry instantiates a new ModifyIPSAllowListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyIPSAllowListEntryWithDefaults

`func NewModifyIPSAllowListEntryWithDefaults() *ModifyIPSAllowListEntry`

NewModifyIPSAllowListEntryWithDefaults instantiates a new ModifyIPSAllowListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *ModifyIPSAllowListEntry) GetDirection() int32`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ModifyIPSAllowListEntry) GetDirectionOk() (*int32, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ModifyIPSAllowListEntry) SetDirection(v int32)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ModifyIPSAllowListEntry) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *ModifyIPSAllowListEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifyIPSAllowListEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifyIPSAllowListEntry) SetId(v string)`

SetId sets Id field to given value.


### GetIpAddress

`func (o *ModifyIPSAllowListEntry) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ModifyIPSAllowListEntry) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ModifyIPSAllowListEntry) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ModifyIPSAllowListEntry) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetNetwork

`func (o *ModifyIPSAllowListEntry) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ModifyIPSAllowListEntry) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ModifyIPSAllowListEntry) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ModifyIPSAllowListEntry) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSubnet

`func (o *ModifyIPSAllowListEntry) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *ModifyIPSAllowListEntry) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *ModifyIPSAllowListEntry) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *ModifyIPSAllowListEntry) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetTrafficType

`func (o *ModifyIPSAllowListEntry) GetTrafficType() int32`

GetTrafficType returns the TrafficType field if non-nil, zero value otherwise.

### GetTrafficTypeOk

`func (o *ModifyIPSAllowListEntry) GetTrafficTypeOk() (*int32, bool)`

GetTrafficTypeOk returns a tuple with the TrafficType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficType

`func (o *ModifyIPSAllowListEntry) SetTrafficType(v int32)`

SetTrafficType sets TrafficType field to given value.

### HasTrafficType

`func (o *ModifyIPSAllowListEntry) HasTrafficType() bool`

HasTrafficType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


