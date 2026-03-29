# DhcpSnoopVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllLags** | Pointer to [**[]OswLagVO**](OswLagVO.md) | All lag list in one device. | [optional] 
**AllPorts** | Pointer to [**[]OswPortVO**](OswPortVO.md) | All Port list in one device. | [optional] 
**DeviceModel** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**DeviceModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**DeviceName** | Pointer to **string** | The name of the device. | [optional] 
**DeviceType** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**Devices** | Pointer to [**[]DeviceVO**](DeviceVO.md) | The devices selected to create entries. | [optional] 
**Id** | Pointer to **string** | The primary id of the dhcp snoop. | [optional] 
**Mac** | Pointer to **string** | The mac of the general device. | [optional] 
**Name** | Pointer to **string** | The name of the dhcp snoop. | [optional] 
**OmadacId** | Pointer to **string** | The id of the omada the dhcp snoop belongs to. | [optional] 
**OswStack** | Pointer to [**OswStackInfoVO**](OswStackInfoVO.md) |  | [optional] 
**Ports** | Pointer to [**[]PortVO**](PortVO.md) | The ports selected. | [optional] 
**ShowModel** | Pointer to **string** | The client showModel | [optional] 
**SiteId** | Pointer to **string** | The id of the site the dhcp snoop belongs to. | [optional] 
**StackId** | Pointer to **string** | The stack of the stack device. | [optional] 
**StackName** | Pointer to **string** | The name of the stacking devices. | [optional] 
**Status** | Pointer to **int32** | The connected status of the device. | [optional] 
**UnSelectedablePorts** | Pointer to [**[]PortVO**](PortVO.md) | The unSelectedable ports of the device. | [optional] 

## Methods

### NewDhcpSnoopVO

`func NewDhcpSnoopVO() *DhcpSnoopVO`

NewDhcpSnoopVO instantiates a new DhcpSnoopVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpSnoopVOWithDefaults

`func NewDhcpSnoopVOWithDefaults() *DhcpSnoopVO`

NewDhcpSnoopVOWithDefaults instantiates a new DhcpSnoopVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllLags

`func (o *DhcpSnoopVO) GetAllLags() []OswLagVO`

GetAllLags returns the AllLags field if non-nil, zero value otherwise.

### GetAllLagsOk

`func (o *DhcpSnoopVO) GetAllLagsOk() (*[]OswLagVO, bool)`

GetAllLagsOk returns a tuple with the AllLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLags

`func (o *DhcpSnoopVO) SetAllLags(v []OswLagVO)`

SetAllLags sets AllLags field to given value.

### HasAllLags

`func (o *DhcpSnoopVO) HasAllLags() bool`

HasAllLags returns a boolean if a field has been set.

### GetAllPorts

`func (o *DhcpSnoopVO) GetAllPorts() []OswPortVO`

GetAllPorts returns the AllPorts field if non-nil, zero value otherwise.

### GetAllPortsOk

`func (o *DhcpSnoopVO) GetAllPortsOk() (*[]OswPortVO, bool)`

GetAllPortsOk returns a tuple with the AllPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPorts

`func (o *DhcpSnoopVO) SetAllPorts(v []OswPortVO)`

SetAllPorts sets AllPorts field to given value.

### HasAllPorts

`func (o *DhcpSnoopVO) HasAllPorts() bool`

HasAllPorts returns a boolean if a field has been set.

### GetDeviceModel

`func (o *DhcpSnoopVO) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *DhcpSnoopVO) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *DhcpSnoopVO) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *DhcpSnoopVO) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceModelVersion

`func (o *DhcpSnoopVO) GetDeviceModelVersion() string`

GetDeviceModelVersion returns the DeviceModelVersion field if non-nil, zero value otherwise.

### GetDeviceModelVersionOk

`func (o *DhcpSnoopVO) GetDeviceModelVersionOk() (*string, bool)`

GetDeviceModelVersionOk returns a tuple with the DeviceModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModelVersion

`func (o *DhcpSnoopVO) SetDeviceModelVersion(v string)`

SetDeviceModelVersion sets DeviceModelVersion field to given value.

### HasDeviceModelVersion

`func (o *DhcpSnoopVO) HasDeviceModelVersion() bool`

HasDeviceModelVersion returns a boolean if a field has been set.

### GetDeviceName

`func (o *DhcpSnoopVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DhcpSnoopVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DhcpSnoopVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DhcpSnoopVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *DhcpSnoopVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DhcpSnoopVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DhcpSnoopVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DhcpSnoopVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDevices

`func (o *DhcpSnoopVO) GetDevices() []DeviceVO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *DhcpSnoopVO) GetDevicesOk() (*[]DeviceVO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *DhcpSnoopVO) SetDevices(v []DeviceVO)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *DhcpSnoopVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetId

`func (o *DhcpSnoopVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DhcpSnoopVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DhcpSnoopVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DhcpSnoopVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMac

`func (o *DhcpSnoopVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DhcpSnoopVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DhcpSnoopVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DhcpSnoopVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *DhcpSnoopVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DhcpSnoopVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DhcpSnoopVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DhcpSnoopVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *DhcpSnoopVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *DhcpSnoopVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *DhcpSnoopVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *DhcpSnoopVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetOswStack

`func (o *DhcpSnoopVO) GetOswStack() OswStackInfoVO`

GetOswStack returns the OswStack field if non-nil, zero value otherwise.

### GetOswStackOk

`func (o *DhcpSnoopVO) GetOswStackOk() (*OswStackInfoVO, bool)`

GetOswStackOk returns a tuple with the OswStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswStack

`func (o *DhcpSnoopVO) SetOswStack(v OswStackInfoVO)`

SetOswStack sets OswStack field to given value.

### HasOswStack

`func (o *DhcpSnoopVO) HasOswStack() bool`

HasOswStack returns a boolean if a field has been set.

### GetPorts

`func (o *DhcpSnoopVO) GetPorts() []PortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *DhcpSnoopVO) GetPortsOk() (*[]PortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *DhcpSnoopVO) SetPorts(v []PortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *DhcpSnoopVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetShowModel

`func (o *DhcpSnoopVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *DhcpSnoopVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *DhcpSnoopVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *DhcpSnoopVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSiteId

`func (o *DhcpSnoopVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DhcpSnoopVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DhcpSnoopVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DhcpSnoopVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStackId

`func (o *DhcpSnoopVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *DhcpSnoopVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *DhcpSnoopVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *DhcpSnoopVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *DhcpSnoopVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *DhcpSnoopVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *DhcpSnoopVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *DhcpSnoopVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStatus

`func (o *DhcpSnoopVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DhcpSnoopVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DhcpSnoopVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DhcpSnoopVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUnSelectedablePorts

`func (o *DhcpSnoopVO) GetUnSelectedablePorts() []PortVO`

GetUnSelectedablePorts returns the UnSelectedablePorts field if non-nil, zero value otherwise.

### GetUnSelectedablePortsOk

`func (o *DhcpSnoopVO) GetUnSelectedablePortsOk() (*[]PortVO, bool)`

GetUnSelectedablePortsOk returns a tuple with the UnSelectedablePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSelectedablePorts

`func (o *DhcpSnoopVO) SetUnSelectedablePorts(v []PortVO)`

SetUnSelectedablePorts sets UnSelectedablePorts field to given value.

### HasUnSelectedablePorts

`func (o *DhcpSnoopVO) HasUnSelectedablePorts() bool`

HasUnSelectedablePorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


