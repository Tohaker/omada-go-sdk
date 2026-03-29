# TopologyRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientType** | Pointer to **string** | Specific type of other node when it is a client. | [optional] 
**IsRoot** | Pointer to **bool** | Whether the Node is root. | [optional] 
**Mac** | Pointer to **string** | Root Node MAC address, like AA-BB-CC-DD-EE-FF. | [optional] 
**MacList** | Pointer to **[]string** | MultiSwitch Node MAC list. | [optional] 
**Model** | Pointer to **string** | Root Node Model. | [optional] 
**ModelVersion** | Pointer to **string** | Root Node ModelVersion. | [optional] 
**Name** | Pointer to **string** | Root Node name. | [optional] 
**NameList** | Pointer to **[]string** | MultiSwitch Node name list. | [optional] 
**OmadaDeviceType** | Pointer to **string** | Specific type of other node when it is a omada device. | [optional] 
**ShowModel** | Pointer to **string** | Root Node ShowModel. | [optional] 
**SpecificType** | Pointer to **int32** | MultiSwitch Node SpecificType, which can only be one of the following two types: 0:Mlag and 1:Vrrp. | [optional] 
**StackGroup** | Pointer to **bool** | Whether the Node is stackGroup. | [optional] 
**Type** | Pointer to **string** | Root Node Type, which can only be one of the following three types: switch, olt and other. | [optional] 

## Methods

### NewTopologyRoot

`func NewTopologyRoot() *TopologyRoot`

NewTopologyRoot instantiates a new TopologyRoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyRootWithDefaults

`func NewTopologyRootWithDefaults() *TopologyRoot`

NewTopologyRootWithDefaults instantiates a new TopologyRoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientType

`func (o *TopologyRoot) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *TopologyRoot) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *TopologyRoot) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *TopologyRoot) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetIsRoot

`func (o *TopologyRoot) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *TopologyRoot) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *TopologyRoot) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *TopologyRoot) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### GetMac

`func (o *TopologyRoot) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *TopologyRoot) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *TopologyRoot) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *TopologyRoot) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMacList

`func (o *TopologyRoot) GetMacList() []string`

GetMacList returns the MacList field if non-nil, zero value otherwise.

### GetMacListOk

`func (o *TopologyRoot) GetMacListOk() (*[]string, bool)`

GetMacListOk returns a tuple with the MacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacList

`func (o *TopologyRoot) SetMacList(v []string)`

SetMacList sets MacList field to given value.

### HasMacList

`func (o *TopologyRoot) HasMacList() bool`

HasMacList returns a boolean if a field has been set.

### GetModel

`func (o *TopologyRoot) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TopologyRoot) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TopologyRoot) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TopologyRoot) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *TopologyRoot) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *TopologyRoot) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *TopologyRoot) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *TopologyRoot) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *TopologyRoot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyRoot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyRoot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopologyRoot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameList

`func (o *TopologyRoot) GetNameList() []string`

GetNameList returns the NameList field if non-nil, zero value otherwise.

### GetNameListOk

`func (o *TopologyRoot) GetNameListOk() (*[]string, bool)`

GetNameListOk returns a tuple with the NameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameList

`func (o *TopologyRoot) SetNameList(v []string)`

SetNameList sets NameList field to given value.

### HasNameList

`func (o *TopologyRoot) HasNameList() bool`

HasNameList returns a boolean if a field has been set.

### GetOmadaDeviceType

`func (o *TopologyRoot) GetOmadaDeviceType() string`

GetOmadaDeviceType returns the OmadaDeviceType field if non-nil, zero value otherwise.

### GetOmadaDeviceTypeOk

`func (o *TopologyRoot) GetOmadaDeviceTypeOk() (*string, bool)`

GetOmadaDeviceTypeOk returns a tuple with the OmadaDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadaDeviceType

`func (o *TopologyRoot) SetOmadaDeviceType(v string)`

SetOmadaDeviceType sets OmadaDeviceType field to given value.

### HasOmadaDeviceType

`func (o *TopologyRoot) HasOmadaDeviceType() bool`

HasOmadaDeviceType returns a boolean if a field has been set.

### GetShowModel

`func (o *TopologyRoot) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *TopologyRoot) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *TopologyRoot) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *TopologyRoot) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSpecificType

`func (o *TopologyRoot) GetSpecificType() int32`

GetSpecificType returns the SpecificType field if non-nil, zero value otherwise.

### GetSpecificTypeOk

`func (o *TopologyRoot) GetSpecificTypeOk() (*int32, bool)`

GetSpecificTypeOk returns a tuple with the SpecificType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificType

`func (o *TopologyRoot) SetSpecificType(v int32)`

SetSpecificType sets SpecificType field to given value.

### HasSpecificType

`func (o *TopologyRoot) HasSpecificType() bool`

HasSpecificType returns a boolean if a field has been set.

### GetStackGroup

`func (o *TopologyRoot) GetStackGroup() bool`

GetStackGroup returns the StackGroup field if non-nil, zero value otherwise.

### GetStackGroupOk

`func (o *TopologyRoot) GetStackGroupOk() (*bool, bool)`

GetStackGroupOk returns a tuple with the StackGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackGroup

`func (o *TopologyRoot) SetStackGroup(v bool)`

SetStackGroup sets StackGroup field to given value.

### HasStackGroup

`func (o *TopologyRoot) HasStackGroup() bool`

HasStackGroup returns a boolean if a field has been set.

### GetType

`func (o *TopologyRoot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopologyRoot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopologyRoot) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopologyRoot) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


