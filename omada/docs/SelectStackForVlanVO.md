# SelectStackForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lags** | Pointer to [**[]SelectStackLagForVlanVO**](SelectStackLagForVlanVO.md) | Lags | [optional] 
**ManuallySelect** | Pointer to **bool** | Whether the stack is manually selected. | [optional] 
**MasterMac** | Pointer to **string** | Master mac | [optional] 
**Members** | Pointer to [**[]SelectStackMemberForVlanVO**](SelectStackMemberForVlanVO.md) | Stack Members | [optional] 
**ReplacedDevice** | Pointer to **bool** | It indicates whether the stack is the dhcp server device that is replaced in the first step | [optional] 
**StackId** | Pointer to **string** | Stack ID | [optional] 
**StackName** | Pointer to **string** | Stack Name | [optional] 
**Status** | Pointer to **int32** | Stack Status should be a value as follows: 0: normal; 1: abnormal; 2: stack not ready | [optional] 

## Methods

### NewSelectStackForVlanVO

`func NewSelectStackForVlanVO() *SelectStackForVlanVO`

NewSelectStackForVlanVO instantiates a new SelectStackForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectStackForVlanVOWithDefaults

`func NewSelectStackForVlanVOWithDefaults() *SelectStackForVlanVO`

NewSelectStackForVlanVOWithDefaults instantiates a new SelectStackForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLags

`func (o *SelectStackForVlanVO) GetLags() []SelectStackLagForVlanVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *SelectStackForVlanVO) GetLagsOk() (*[]SelectStackLagForVlanVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *SelectStackForVlanVO) SetLags(v []SelectStackLagForVlanVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *SelectStackForVlanVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetManuallySelect

`func (o *SelectStackForVlanVO) GetManuallySelect() bool`

GetManuallySelect returns the ManuallySelect field if non-nil, zero value otherwise.

### GetManuallySelectOk

`func (o *SelectStackForVlanVO) GetManuallySelectOk() (*bool, bool)`

GetManuallySelectOk returns a tuple with the ManuallySelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallySelect

`func (o *SelectStackForVlanVO) SetManuallySelect(v bool)`

SetManuallySelect sets ManuallySelect field to given value.

### HasManuallySelect

`func (o *SelectStackForVlanVO) HasManuallySelect() bool`

HasManuallySelect returns a boolean if a field has been set.

### GetMasterMac

`func (o *SelectStackForVlanVO) GetMasterMac() string`

GetMasterMac returns the MasterMac field if non-nil, zero value otherwise.

### GetMasterMacOk

`func (o *SelectStackForVlanVO) GetMasterMacOk() (*string, bool)`

GetMasterMacOk returns a tuple with the MasterMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterMac

`func (o *SelectStackForVlanVO) SetMasterMac(v string)`

SetMasterMac sets MasterMac field to given value.

### HasMasterMac

`func (o *SelectStackForVlanVO) HasMasterMac() bool`

HasMasterMac returns a boolean if a field has been set.

### GetMembers

`func (o *SelectStackForVlanVO) GetMembers() []SelectStackMemberForVlanVO`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *SelectStackForVlanVO) GetMembersOk() (*[]SelectStackMemberForVlanVO, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *SelectStackForVlanVO) SetMembers(v []SelectStackMemberForVlanVO)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *SelectStackForVlanVO) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetReplacedDevice

`func (o *SelectStackForVlanVO) GetReplacedDevice() bool`

GetReplacedDevice returns the ReplacedDevice field if non-nil, zero value otherwise.

### GetReplacedDeviceOk

`func (o *SelectStackForVlanVO) GetReplacedDeviceOk() (*bool, bool)`

GetReplacedDeviceOk returns a tuple with the ReplacedDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedDevice

`func (o *SelectStackForVlanVO) SetReplacedDevice(v bool)`

SetReplacedDevice sets ReplacedDevice field to given value.

### HasReplacedDevice

`func (o *SelectStackForVlanVO) HasReplacedDevice() bool`

HasReplacedDevice returns a boolean if a field has been set.

### GetStackId

`func (o *SelectStackForVlanVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *SelectStackForVlanVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *SelectStackForVlanVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *SelectStackForVlanVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *SelectStackForVlanVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *SelectStackForVlanVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *SelectStackForVlanVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *SelectStackForVlanVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStatus

`func (o *SelectStackForVlanVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SelectStackForVlanVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SelectStackForVlanVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SelectStackForVlanVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


