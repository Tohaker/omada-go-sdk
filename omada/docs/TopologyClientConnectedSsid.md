# TopologyClientConnectedSsid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientVlanId** | Pointer to **int32** | Client Report Vlan Id. | [optional] 
**Id** | Pointer to **string** | SSID ID. | [optional] 
**Name** | Pointer to **string** | SSID Name. | [optional] 
**VlanEnable** | Pointer to **bool** | Whether the SSID enable vlan. | [optional] 
**VlanIds** | Pointer to **[]int32** | Vlan ID while SSID enable vlan. | [optional] 

## Methods

### NewTopologyClientConnectedSsid

`func NewTopologyClientConnectedSsid() *TopologyClientConnectedSsid`

NewTopologyClientConnectedSsid instantiates a new TopologyClientConnectedSsid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyClientConnectedSsidWithDefaults

`func NewTopologyClientConnectedSsidWithDefaults() *TopologyClientConnectedSsid`

NewTopologyClientConnectedSsidWithDefaults instantiates a new TopologyClientConnectedSsid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientVlanId

`func (o *TopologyClientConnectedSsid) GetClientVlanId() int32`

GetClientVlanId returns the ClientVlanId field if non-nil, zero value otherwise.

### GetClientVlanIdOk

`func (o *TopologyClientConnectedSsid) GetClientVlanIdOk() (*int32, bool)`

GetClientVlanIdOk returns a tuple with the ClientVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVlanId

`func (o *TopologyClientConnectedSsid) SetClientVlanId(v int32)`

SetClientVlanId sets ClientVlanId field to given value.

### HasClientVlanId

`func (o *TopologyClientConnectedSsid) HasClientVlanId() bool`

HasClientVlanId returns a boolean if a field has been set.

### GetId

`func (o *TopologyClientConnectedSsid) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopologyClientConnectedSsid) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopologyClientConnectedSsid) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TopologyClientConnectedSsid) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TopologyClientConnectedSsid) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyClientConnectedSsid) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyClientConnectedSsid) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopologyClientConnectedSsid) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVlanEnable

`func (o *TopologyClientConnectedSsid) GetVlanEnable() bool`

GetVlanEnable returns the VlanEnable field if non-nil, zero value otherwise.

### GetVlanEnableOk

`func (o *TopologyClientConnectedSsid) GetVlanEnableOk() (*bool, bool)`

GetVlanEnableOk returns a tuple with the VlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnable

`func (o *TopologyClientConnectedSsid) SetVlanEnable(v bool)`

SetVlanEnable sets VlanEnable field to given value.

### HasVlanEnable

`func (o *TopologyClientConnectedSsid) HasVlanEnable() bool`

HasVlanEnable returns a boolean if a field has been set.

### GetVlanIds

`func (o *TopologyClientConnectedSsid) GetVlanIds() []int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *TopologyClientConnectedSsid) GetVlanIdsOk() (*[]int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *TopologyClientConnectedSsid) SetVlanIds(v []int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *TopologyClientConnectedSsid) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


