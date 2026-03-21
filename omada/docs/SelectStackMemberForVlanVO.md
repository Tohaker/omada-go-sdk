# SelectStackMemberForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedInAdvanced** | Pointer to **bool** | Whether the device is added in advanced. | [optional] 
**Compatible** | Pointer to **int32** | Device firmware and controller compatibility type.Compatible should be a value as follows: 0:COMPATIBLE;1:HIGH_MAJOR_VER;2:LOW_MAJOR_VER;3:HIGH_MINOR_VER;4:LOW_MINOR_VER;7:HIGH_COMPONENT_VER;10:DEVICE_NOT_COMPATIBLE;11:HIGH_ADOPT_COMMPONENT;12:DEVICE_CATEGORY_NOT_COMPATIBLE;14:DEVICE_NOT_COMPATIBLE_IN_CLUSTER | [optional] 
**Mac** | Pointer to **string** | Device Mac | [optional] 
**ModeVersion** | Pointer to **string** | Device Model Version | [optional] 
**Model** | Pointer to **string** | Device Model | [optional] 
**Name** | Pointer to **string** | Device Name | [optional] 
**Ports** | Pointer to [**[]SelectStackMemberPortForVlanVO**](SelectStackMemberPortForVlanVO.md) | Ports | [optional] 
**StatusCategory** | Pointer to **int32** | Device status category, 0: Disconnected, 1: Connected, 2: Pending,3: Heartbeat Missed, 4: Isolated | [optional] 
**Unit** | Pointer to **int32** | Stack member&#39;s unit | [optional] 

## Methods

### NewSelectStackMemberForVlanVO

`func NewSelectStackMemberForVlanVO() *SelectStackMemberForVlanVO`

NewSelectStackMemberForVlanVO instantiates a new SelectStackMemberForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectStackMemberForVlanVOWithDefaults

`func NewSelectStackMemberForVlanVOWithDefaults() *SelectStackMemberForVlanVO`

NewSelectStackMemberForVlanVOWithDefaults instantiates a new SelectStackMemberForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedInAdvanced

`func (o *SelectStackMemberForVlanVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *SelectStackMemberForVlanVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *SelectStackMemberForVlanVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *SelectStackMemberForVlanVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetCompatible

`func (o *SelectStackMemberForVlanVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *SelectStackMemberForVlanVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *SelectStackMemberForVlanVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *SelectStackMemberForVlanVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetMac

`func (o *SelectStackMemberForVlanVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SelectStackMemberForVlanVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SelectStackMemberForVlanVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SelectStackMemberForVlanVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModeVersion

`func (o *SelectStackMemberForVlanVO) GetModeVersion() string`

GetModeVersion returns the ModeVersion field if non-nil, zero value otherwise.

### GetModeVersionOk

`func (o *SelectStackMemberForVlanVO) GetModeVersionOk() (*string, bool)`

GetModeVersionOk returns a tuple with the ModeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeVersion

`func (o *SelectStackMemberForVlanVO) SetModeVersion(v string)`

SetModeVersion sets ModeVersion field to given value.

### HasModeVersion

`func (o *SelectStackMemberForVlanVO) HasModeVersion() bool`

HasModeVersion returns a boolean if a field has been set.

### GetModel

`func (o *SelectStackMemberForVlanVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SelectStackMemberForVlanVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SelectStackMemberForVlanVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SelectStackMemberForVlanVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *SelectStackMemberForVlanVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelectStackMemberForVlanVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelectStackMemberForVlanVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelectStackMemberForVlanVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *SelectStackMemberForVlanVO) GetPorts() []SelectStackMemberPortForVlanVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *SelectStackMemberForVlanVO) GetPortsOk() (*[]SelectStackMemberPortForVlanVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *SelectStackMemberForVlanVO) SetPorts(v []SelectStackMemberPortForVlanVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *SelectStackMemberForVlanVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetStatusCategory

`func (o *SelectStackMemberForVlanVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *SelectStackMemberForVlanVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *SelectStackMemberForVlanVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *SelectStackMemberForVlanVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetUnit

`func (o *SelectStackMemberForVlanVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *SelectStackMemberForVlanVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *SelectStackMemberForVlanVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *SelectStackMemberForVlanVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


