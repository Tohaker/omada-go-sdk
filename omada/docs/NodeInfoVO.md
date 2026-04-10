# NodeInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentNode** | Pointer to **string** |  | [optional] 
**ManagedDevices** | Pointer to **int32** |  | [optional] 
**ManagedSites** | Pointer to **int32** |  | [optional] 
**NodeName** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SystemTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewNodeInfoVO

`func NewNodeInfoVO() *NodeInfoVO`

NewNodeInfoVO instantiates a new NodeInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeInfoVOWithDefaults

`func NewNodeInfoVOWithDefaults() *NodeInfoVO`

NewNodeInfoVOWithDefaults instantiates a new NodeInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentNode

`func (o *NodeInfoVO) GetCurrentNode() string`

GetCurrentNode returns the CurrentNode field if non-nil, zero value otherwise.

### GetCurrentNodeOk

`func (o *NodeInfoVO) GetCurrentNodeOk() (*string, bool)`

GetCurrentNodeOk returns a tuple with the CurrentNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNode

`func (o *NodeInfoVO) SetCurrentNode(v string)`

SetCurrentNode sets CurrentNode field to given value.

### HasCurrentNode

`func (o *NodeInfoVO) HasCurrentNode() bool`

HasCurrentNode returns a boolean if a field has been set.

### GetManagedDevices

`func (o *NodeInfoVO) GetManagedDevices() int32`

GetManagedDevices returns the ManagedDevices field if non-nil, zero value otherwise.

### GetManagedDevicesOk

`func (o *NodeInfoVO) GetManagedDevicesOk() (*int32, bool)`

GetManagedDevicesOk returns a tuple with the ManagedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDevices

`func (o *NodeInfoVO) SetManagedDevices(v int32)`

SetManagedDevices sets ManagedDevices field to given value.

### HasManagedDevices

`func (o *NodeInfoVO) HasManagedDevices() bool`

HasManagedDevices returns a boolean if a field has been set.

### GetManagedSites

`func (o *NodeInfoVO) GetManagedSites() int32`

GetManagedSites returns the ManagedSites field if non-nil, zero value otherwise.

### GetManagedSitesOk

`func (o *NodeInfoVO) GetManagedSitesOk() (*int32, bool)`

GetManagedSitesOk returns a tuple with the ManagedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedSites

`func (o *NodeInfoVO) SetManagedSites(v int32)`

SetManagedSites sets ManagedSites field to given value.

### HasManagedSites

`func (o *NodeInfoVO) HasManagedSites() bool`

HasManagedSites returns a boolean if a field has been set.

### GetNodeName

`func (o *NodeInfoVO) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *NodeInfoVO) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *NodeInfoVO) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *NodeInfoVO) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetRole

`func (o *NodeInfoVO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NodeInfoVO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NodeInfoVO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *NodeInfoVO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *NodeInfoVO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeInfoVO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeInfoVO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeInfoVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemTime

`func (o *NodeInfoVO) GetSystemTime() int64`

GetSystemTime returns the SystemTime field if non-nil, zero value otherwise.

### GetSystemTimeOk

`func (o *NodeInfoVO) GetSystemTimeOk() (*int64, bool)`

GetSystemTimeOk returns a tuple with the SystemTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTime

`func (o *NodeInfoVO) SetSystemTime(v int64)`

SetSystemTime sets SystemTime field to given value.

### HasSystemTime

`func (o *NodeInfoVO) HasSystemTime() bool`

HasSystemTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


