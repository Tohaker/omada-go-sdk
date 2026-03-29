# AutoEffectDeviceForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpMode** | Pointer to **int32** | It should be a value as follows, 0:None  1:DHCP Server 2:DHCP Relay | [optional] 
**Mac** | Pointer to **string** | Device Mac | [optional] 
**StackId** | Pointer to **string** | Stack Id, only valid when the device is stack. | [optional] 

## Methods

### NewAutoEffectDeviceForVlanVO

`func NewAutoEffectDeviceForVlanVO() *AutoEffectDeviceForVlanVO`

NewAutoEffectDeviceForVlanVO instantiates a new AutoEffectDeviceForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoEffectDeviceForVlanVOWithDefaults

`func NewAutoEffectDeviceForVlanVOWithDefaults() *AutoEffectDeviceForVlanVO`

NewAutoEffectDeviceForVlanVOWithDefaults instantiates a new AutoEffectDeviceForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpMode

`func (o *AutoEffectDeviceForVlanVO) GetDhcpMode() int32`

GetDhcpMode returns the DhcpMode field if non-nil, zero value otherwise.

### GetDhcpModeOk

`func (o *AutoEffectDeviceForVlanVO) GetDhcpModeOk() (*int32, bool)`

GetDhcpModeOk returns a tuple with the DhcpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpMode

`func (o *AutoEffectDeviceForVlanVO) SetDhcpMode(v int32)`

SetDhcpMode sets DhcpMode field to given value.

### HasDhcpMode

`func (o *AutoEffectDeviceForVlanVO) HasDhcpMode() bool`

HasDhcpMode returns a boolean if a field has been set.

### GetMac

`func (o *AutoEffectDeviceForVlanVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *AutoEffectDeviceForVlanVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *AutoEffectDeviceForVlanVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *AutoEffectDeviceForVlanVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStackId

`func (o *AutoEffectDeviceForVlanVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *AutoEffectDeviceForVlanVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *AutoEffectDeviceForVlanVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *AutoEffectDeviceForVlanVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


