# DeviceOuiModeVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | **string** |  | 
**DeviceName** | Pointer to **string** |  | [optional] 
**LagInfo** | Pointer to [**[]LagInfoVO**](LagInfoVO.md) |  | [optional] 
**LagList** | Pointer to **[]int32** |  | [optional] 
**LanList** | Pointer to **[]string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModelVersion** | Pointer to **string** |  | [optional] 
**PortList** | Pointer to **[]int32** |  | [optional] 
**StackId** | Pointer to **string** |  | [optional] 
**StackName** | Pointer to **string** |  | [optional] 
**StackPortList** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceOuiModeVO

`func NewDeviceOuiModeVO(deviceMac string, ) *DeviceOuiModeVO`

NewDeviceOuiModeVO instantiates a new DeviceOuiModeVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceOuiModeVOWithDefaults

`func NewDeviceOuiModeVOWithDefaults() *DeviceOuiModeVO`

NewDeviceOuiModeVOWithDefaults instantiates a new DeviceOuiModeVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *DeviceOuiModeVO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *DeviceOuiModeVO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *DeviceOuiModeVO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.


### GetDeviceName

`func (o *DeviceOuiModeVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceOuiModeVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceOuiModeVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DeviceOuiModeVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetLagInfo

`func (o *DeviceOuiModeVO) GetLagInfo() []LagInfoVO`

GetLagInfo returns the LagInfo field if non-nil, zero value otherwise.

### GetLagInfoOk

`func (o *DeviceOuiModeVO) GetLagInfoOk() (*[]LagInfoVO, bool)`

GetLagInfoOk returns a tuple with the LagInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagInfo

`func (o *DeviceOuiModeVO) SetLagInfo(v []LagInfoVO)`

SetLagInfo sets LagInfo field to given value.

### HasLagInfo

`func (o *DeviceOuiModeVO) HasLagInfo() bool`

HasLagInfo returns a boolean if a field has been set.

### GetLagList

`func (o *DeviceOuiModeVO) GetLagList() []int32`

GetLagList returns the LagList field if non-nil, zero value otherwise.

### GetLagListOk

`func (o *DeviceOuiModeVO) GetLagListOk() (*[]int32, bool)`

GetLagListOk returns a tuple with the LagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagList

`func (o *DeviceOuiModeVO) SetLagList(v []int32)`

SetLagList sets LagList field to given value.

### HasLagList

`func (o *DeviceOuiModeVO) HasLagList() bool`

HasLagList returns a boolean if a field has been set.

### GetLanList

`func (o *DeviceOuiModeVO) GetLanList() []string`

GetLanList returns the LanList field if non-nil, zero value otherwise.

### GetLanListOk

`func (o *DeviceOuiModeVO) GetLanListOk() (*[]string, bool)`

GetLanListOk returns a tuple with the LanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanList

`func (o *DeviceOuiModeVO) SetLanList(v []string)`

SetLanList sets LanList field to given value.

### HasLanList

`func (o *DeviceOuiModeVO) HasLanList() bool`

HasLanList returns a boolean if a field has been set.

### GetModel

`func (o *DeviceOuiModeVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceOuiModeVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceOuiModeVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceOuiModeVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *DeviceOuiModeVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *DeviceOuiModeVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *DeviceOuiModeVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *DeviceOuiModeVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetPortList

`func (o *DeviceOuiModeVO) GetPortList() []int32`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *DeviceOuiModeVO) GetPortListOk() (*[]int32, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *DeviceOuiModeVO) SetPortList(v []int32)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *DeviceOuiModeVO) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetStackId

`func (o *DeviceOuiModeVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *DeviceOuiModeVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *DeviceOuiModeVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *DeviceOuiModeVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *DeviceOuiModeVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *DeviceOuiModeVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *DeviceOuiModeVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *DeviceOuiModeVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStackPortList

`func (o *DeviceOuiModeVO) GetStackPortList() []string`

GetStackPortList returns the StackPortList field if non-nil, zero value otherwise.

### GetStackPortListOk

`func (o *DeviceOuiModeVO) GetStackPortListOk() (*[]string, bool)`

GetStackPortListOk returns a tuple with the StackPortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPortList

`func (o *DeviceOuiModeVO) SetStackPortList(v []string)`

SetStackPortList sets StackPortList field to given value.

### HasStackPortList

`func (o *DeviceOuiModeVO) HasStackPortList() bool`

HasStackPortList returns a boolean if a field has been set.

### GetType

`func (o *DeviceOuiModeVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceOuiModeVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceOuiModeVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceOuiModeVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


