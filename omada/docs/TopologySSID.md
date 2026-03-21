# TopologySSID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientNum** | Pointer to **int32** | Client Num. | [optional] 
**Configuring** | Pointer to **bool** | Whether SSID is configuring. | [optional] 
**Id** | Pointer to **string** | SSID ID. | [optional] 
**Name** | Pointer to **string** | SSID Name. | [optional] 
**VlanEnable** | Pointer to **bool** | Whether enable vlan. | [optional] 
**VlanIds** | Pointer to **[]int32** | Vlan IDs. | [optional] 
**WlanName** | Pointer to **string** | WLAN group name | [optional] 

## Methods

### NewTopologySSID

`func NewTopologySSID() *TopologySSID`

NewTopologySSID instantiates a new TopologySSID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologySSIDWithDefaults

`func NewTopologySSIDWithDefaults() *TopologySSID`

NewTopologySSIDWithDefaults instantiates a new TopologySSID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientNum

`func (o *TopologySSID) GetClientNum() int32`

GetClientNum returns the ClientNum field if non-nil, zero value otherwise.

### GetClientNumOk

`func (o *TopologySSID) GetClientNumOk() (*int32, bool)`

GetClientNumOk returns a tuple with the ClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNum

`func (o *TopologySSID) SetClientNum(v int32)`

SetClientNum sets ClientNum field to given value.

### HasClientNum

`func (o *TopologySSID) HasClientNum() bool`

HasClientNum returns a boolean if a field has been set.

### GetConfiguring

`func (o *TopologySSID) GetConfiguring() bool`

GetConfiguring returns the Configuring field if non-nil, zero value otherwise.

### GetConfiguringOk

`func (o *TopologySSID) GetConfiguringOk() (*bool, bool)`

GetConfiguringOk returns a tuple with the Configuring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguring

`func (o *TopologySSID) SetConfiguring(v bool)`

SetConfiguring sets Configuring field to given value.

### HasConfiguring

`func (o *TopologySSID) HasConfiguring() bool`

HasConfiguring returns a boolean if a field has been set.

### GetId

`func (o *TopologySSID) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopologySSID) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopologySSID) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TopologySSID) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TopologySSID) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologySSID) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologySSID) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopologySSID) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVlanEnable

`func (o *TopologySSID) GetVlanEnable() bool`

GetVlanEnable returns the VlanEnable field if non-nil, zero value otherwise.

### GetVlanEnableOk

`func (o *TopologySSID) GetVlanEnableOk() (*bool, bool)`

GetVlanEnableOk returns a tuple with the VlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnable

`func (o *TopologySSID) SetVlanEnable(v bool)`

SetVlanEnable sets VlanEnable field to given value.

### HasVlanEnable

`func (o *TopologySSID) HasVlanEnable() bool`

HasVlanEnable returns a boolean if a field has been set.

### GetVlanIds

`func (o *TopologySSID) GetVlanIds() []int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *TopologySSID) GetVlanIdsOk() (*[]int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *TopologySSID) SetVlanIds(v []int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *TopologySSID) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.

### GetWlanName

`func (o *TopologySSID) GetWlanName() string`

GetWlanName returns the WlanName field if non-nil, zero value otherwise.

### GetWlanNameOk

`func (o *TopologySSID) GetWlanNameOk() (*string, bool)`

GetWlanNameOk returns a tuple with the WlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanName

`func (o *TopologySSID) SetWlanName(v string)`

SetWlanName sets WlanName field to given value.

### HasWlanName

`func (o *TopologySSID) HasWlanName() bool`

HasWlanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


