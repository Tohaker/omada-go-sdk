# DevicePortVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceModel** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**DeviceModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**DeviceName** | Pointer to **string** | The name of one device. | [optional] 
**DeviceType** | Pointer to **string** | Device type:ap, gateway, switch, olt | [optional] 
**ForbiddenRouterPorts** | Pointer to [**[]PortVO**](PortVO.md) | The collection of forbidden router ports related to one device. | [optional] 
**Mac** | Pointer to **string** | The unique identification of one device. | [optional] 
**OswStack** | Pointer to [**OswStackInfoVO**](OswStackInfoVO.md) |  | [optional] 
**StackId** | Pointer to **string** | The stack id of devices. | [optional] 
**StackName** | Pointer to **string** | The name of the stacking devices. | [optional] 
**StaticRouterPorts** | Pointer to [**[]PortVO**](PortVO.md) | The collection of static router ports related to one device. | [optional] 

## Methods

### NewDevicePortVO

`func NewDevicePortVO() *DevicePortVO`

NewDevicePortVO instantiates a new DevicePortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePortVOWithDefaults

`func NewDevicePortVOWithDefaults() *DevicePortVO`

NewDevicePortVOWithDefaults instantiates a new DevicePortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceModel

`func (o *DevicePortVO) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *DevicePortVO) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *DevicePortVO) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *DevicePortVO) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceModelVersion

`func (o *DevicePortVO) GetDeviceModelVersion() string`

GetDeviceModelVersion returns the DeviceModelVersion field if non-nil, zero value otherwise.

### GetDeviceModelVersionOk

`func (o *DevicePortVO) GetDeviceModelVersionOk() (*string, bool)`

GetDeviceModelVersionOk returns a tuple with the DeviceModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModelVersion

`func (o *DevicePortVO) SetDeviceModelVersion(v string)`

SetDeviceModelVersion sets DeviceModelVersion field to given value.

### HasDeviceModelVersion

`func (o *DevicePortVO) HasDeviceModelVersion() bool`

HasDeviceModelVersion returns a boolean if a field has been set.

### GetDeviceName

`func (o *DevicePortVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DevicePortVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DevicePortVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DevicePortVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *DevicePortVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DevicePortVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DevicePortVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DevicePortVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetForbiddenRouterPorts

`func (o *DevicePortVO) GetForbiddenRouterPorts() []PortVO`

GetForbiddenRouterPorts returns the ForbiddenRouterPorts field if non-nil, zero value otherwise.

### GetForbiddenRouterPortsOk

`func (o *DevicePortVO) GetForbiddenRouterPortsOk() (*[]PortVO, bool)`

GetForbiddenRouterPortsOk returns a tuple with the ForbiddenRouterPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenRouterPorts

`func (o *DevicePortVO) SetForbiddenRouterPorts(v []PortVO)`

SetForbiddenRouterPorts sets ForbiddenRouterPorts field to given value.

### HasForbiddenRouterPorts

`func (o *DevicePortVO) HasForbiddenRouterPorts() bool`

HasForbiddenRouterPorts returns a boolean if a field has been set.

### GetMac

`func (o *DevicePortVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DevicePortVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DevicePortVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DevicePortVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetOswStack

`func (o *DevicePortVO) GetOswStack() OswStackInfoVO`

GetOswStack returns the OswStack field if non-nil, zero value otherwise.

### GetOswStackOk

`func (o *DevicePortVO) GetOswStackOk() (*OswStackInfoVO, bool)`

GetOswStackOk returns a tuple with the OswStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswStack

`func (o *DevicePortVO) SetOswStack(v OswStackInfoVO)`

SetOswStack sets OswStack field to given value.

### HasOswStack

`func (o *DevicePortVO) HasOswStack() bool`

HasOswStack returns a boolean if a field has been set.

### GetStackId

`func (o *DevicePortVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *DevicePortVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *DevicePortVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *DevicePortVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *DevicePortVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *DevicePortVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *DevicePortVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *DevicePortVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStaticRouterPorts

`func (o *DevicePortVO) GetStaticRouterPorts() []PortVO`

GetStaticRouterPorts returns the StaticRouterPorts field if non-nil, zero value otherwise.

### GetStaticRouterPortsOk

`func (o *DevicePortVO) GetStaticRouterPortsOk() (*[]PortVO, bool)`

GetStaticRouterPortsOk returns a tuple with the StaticRouterPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRouterPorts

`func (o *DevicePortVO) SetStaticRouterPorts(v []PortVO)`

SetStaticRouterPorts sets StaticRouterPorts field to given value.

### HasStaticRouterPorts

`func (o *DevicePortVO) HasStaticRouterPorts() bool`

HasStaticRouterPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


