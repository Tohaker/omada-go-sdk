# VlanNetworkAffectingDeviceVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpMode** | Pointer to **int32** | DHCP Mode, Only type is switch will return. | [optional] 
**Mac** | Pointer to **string** | Device Mac. | [optional] 
**Name** | Pointer to **string** | Device Name. | [optional] 
**StackId** | Pointer to **string** | StackId | [optional] 
**Type** | Pointer to **int32** | Device type, 1: gateway  2: switch  3: ap | [optional] 

## Methods

### NewVlanNetworkAffectingDeviceVO

`func NewVlanNetworkAffectingDeviceVO() *VlanNetworkAffectingDeviceVO`

NewVlanNetworkAffectingDeviceVO instantiates a new VlanNetworkAffectingDeviceVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanNetworkAffectingDeviceVOWithDefaults

`func NewVlanNetworkAffectingDeviceVOWithDefaults() *VlanNetworkAffectingDeviceVO`

NewVlanNetworkAffectingDeviceVOWithDefaults instantiates a new VlanNetworkAffectingDeviceVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpMode

`func (o *VlanNetworkAffectingDeviceVO) GetDhcpMode() int32`

GetDhcpMode returns the DhcpMode field if non-nil, zero value otherwise.

### GetDhcpModeOk

`func (o *VlanNetworkAffectingDeviceVO) GetDhcpModeOk() (*int32, bool)`

GetDhcpModeOk returns a tuple with the DhcpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpMode

`func (o *VlanNetworkAffectingDeviceVO) SetDhcpMode(v int32)`

SetDhcpMode sets DhcpMode field to given value.

### HasDhcpMode

`func (o *VlanNetworkAffectingDeviceVO) HasDhcpMode() bool`

HasDhcpMode returns a boolean if a field has been set.

### GetMac

`func (o *VlanNetworkAffectingDeviceVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *VlanNetworkAffectingDeviceVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *VlanNetworkAffectingDeviceVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *VlanNetworkAffectingDeviceVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *VlanNetworkAffectingDeviceVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VlanNetworkAffectingDeviceVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VlanNetworkAffectingDeviceVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VlanNetworkAffectingDeviceVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStackId

`func (o *VlanNetworkAffectingDeviceVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *VlanNetworkAffectingDeviceVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *VlanNetworkAffectingDeviceVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *VlanNetworkAffectingDeviceVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetType

`func (o *VlanNetworkAffectingDeviceVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VlanNetworkAffectingDeviceVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VlanNetworkAffectingDeviceVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *VlanNetworkAffectingDeviceVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


