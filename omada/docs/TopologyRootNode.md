# TopologyRootNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Root Node MAC address, like AA-BB-CC-DD-EE-FF. | [optional] 
**MacList** | Pointer to **[]string** | MultiSwitch Node MAC list. | [optional] 
**SpecificType** | Pointer to **int32** | MultiSwitch Node SpecificType, which can only be one of the following two types: 0:Mlag and 1:Vrrp. | [optional] 
**Type** | Pointer to **string** | Root Node Type, which can only be one of the following three types: switch, olt and other. | [optional] 

## Methods

### NewTopologyRootNode

`func NewTopologyRootNode() *TopologyRootNode`

NewTopologyRootNode instantiates a new TopologyRootNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyRootNodeWithDefaults

`func NewTopologyRootNodeWithDefaults() *TopologyRootNode`

NewTopologyRootNodeWithDefaults instantiates a new TopologyRootNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *TopologyRootNode) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *TopologyRootNode) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *TopologyRootNode) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *TopologyRootNode) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMacList

`func (o *TopologyRootNode) GetMacList() []string`

GetMacList returns the MacList field if non-nil, zero value otherwise.

### GetMacListOk

`func (o *TopologyRootNode) GetMacListOk() (*[]string, bool)`

GetMacListOk returns a tuple with the MacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacList

`func (o *TopologyRootNode) SetMacList(v []string)`

SetMacList sets MacList field to given value.

### HasMacList

`func (o *TopologyRootNode) HasMacList() bool`

HasMacList returns a boolean if a field has been set.

### GetSpecificType

`func (o *TopologyRootNode) GetSpecificType() int32`

GetSpecificType returns the SpecificType field if non-nil, zero value otherwise.

### GetSpecificTypeOk

`func (o *TopologyRootNode) GetSpecificTypeOk() (*int32, bool)`

GetSpecificTypeOk returns a tuple with the SpecificType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificType

`func (o *TopologyRootNode) SetSpecificType(v int32)`

SetSpecificType sets SpecificType field to given value.

### HasSpecificType

`func (o *TopologyRootNode) HasSpecificType() bool`

HasSpecificType returns a boolean if a field has been set.

### GetType

`func (o *TopologyRootNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopologyRootNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopologyRootNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopologyRootNode) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


