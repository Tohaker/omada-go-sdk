# ArpDetectVO

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
**DhcpSnoopPorts** | Pointer to [**[]PortVO**](PortVO.md) | The ports dhcp snooping selected. | [optional] 
**Id** | Pointer to **string** | The primary id of the arp detection. | [optional] 
**Impbs** | Pointer to [**[]ImpbVO**](ImpbVO.md) | The impbs selected. | [optional] 
**Mac** | Pointer to **string** | The mac of the general device. | [optional] 
**Name** | Pointer to **string** | The name of the arp detect. | [optional] 
**OmadacId** | Pointer to **string** | The id of the omada the arp detect belongs to. | [optional] 
**OswStack** | Pointer to [**OswStackInfoVO**](OswStackInfoVO.md) |  | [optional] 
**Ports** | Pointer to [**[]PortVO**](PortVO.md) | The ports selected. | [optional] 
**ShowModel** | Pointer to **string** | The client showModel | [optional] 
**SiteId** | Pointer to **string** | The id of the site the arp detect belongs to. | [optional] 
**StackId** | Pointer to **string** | The stack of the stack device. | [optional] 
**StackName** | Pointer to **string** | The name of the stacking devices. | [optional] 
**Status** | Pointer to **int32** | The connected status of the device. | [optional] 
**UnSelectedablePorts** | Pointer to [**[]PortVO**](PortVO.md) | The unSelectedable ports of the device. | [optional] 

## Methods

### NewArpDetectVO

`func NewArpDetectVO() *ArpDetectVO`

NewArpDetectVO instantiates a new ArpDetectVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArpDetectVOWithDefaults

`func NewArpDetectVOWithDefaults() *ArpDetectVO`

NewArpDetectVOWithDefaults instantiates a new ArpDetectVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllLags

`func (o *ArpDetectVO) GetAllLags() []OswLagVO`

GetAllLags returns the AllLags field if non-nil, zero value otherwise.

### GetAllLagsOk

`func (o *ArpDetectVO) GetAllLagsOk() (*[]OswLagVO, bool)`

GetAllLagsOk returns a tuple with the AllLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLags

`func (o *ArpDetectVO) SetAllLags(v []OswLagVO)`

SetAllLags sets AllLags field to given value.

### HasAllLags

`func (o *ArpDetectVO) HasAllLags() bool`

HasAllLags returns a boolean if a field has been set.

### GetAllPorts

`func (o *ArpDetectVO) GetAllPorts() []OswPortVO`

GetAllPorts returns the AllPorts field if non-nil, zero value otherwise.

### GetAllPortsOk

`func (o *ArpDetectVO) GetAllPortsOk() (*[]OswPortVO, bool)`

GetAllPortsOk returns a tuple with the AllPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPorts

`func (o *ArpDetectVO) SetAllPorts(v []OswPortVO)`

SetAllPorts sets AllPorts field to given value.

### HasAllPorts

`func (o *ArpDetectVO) HasAllPorts() bool`

HasAllPorts returns a boolean if a field has been set.

### GetDeviceModel

`func (o *ArpDetectVO) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *ArpDetectVO) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *ArpDetectVO) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *ArpDetectVO) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceModelVersion

`func (o *ArpDetectVO) GetDeviceModelVersion() string`

GetDeviceModelVersion returns the DeviceModelVersion field if non-nil, zero value otherwise.

### GetDeviceModelVersionOk

`func (o *ArpDetectVO) GetDeviceModelVersionOk() (*string, bool)`

GetDeviceModelVersionOk returns a tuple with the DeviceModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModelVersion

`func (o *ArpDetectVO) SetDeviceModelVersion(v string)`

SetDeviceModelVersion sets DeviceModelVersion field to given value.

### HasDeviceModelVersion

`func (o *ArpDetectVO) HasDeviceModelVersion() bool`

HasDeviceModelVersion returns a boolean if a field has been set.

### GetDeviceName

`func (o *ArpDetectVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ArpDetectVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ArpDetectVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ArpDetectVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *ArpDetectVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ArpDetectVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ArpDetectVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ArpDetectVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDevices

`func (o *ArpDetectVO) GetDevices() []DeviceVO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ArpDetectVO) GetDevicesOk() (*[]DeviceVO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ArpDetectVO) SetDevices(v []DeviceVO)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *ArpDetectVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetDhcpSnoopPorts

`func (o *ArpDetectVO) GetDhcpSnoopPorts() []PortVO`

GetDhcpSnoopPorts returns the DhcpSnoopPorts field if non-nil, zero value otherwise.

### GetDhcpSnoopPortsOk

`func (o *ArpDetectVO) GetDhcpSnoopPortsOk() (*[]PortVO, bool)`

GetDhcpSnoopPortsOk returns a tuple with the DhcpSnoopPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSnoopPorts

`func (o *ArpDetectVO) SetDhcpSnoopPorts(v []PortVO)`

SetDhcpSnoopPorts sets DhcpSnoopPorts field to given value.

### HasDhcpSnoopPorts

`func (o *ArpDetectVO) HasDhcpSnoopPorts() bool`

HasDhcpSnoopPorts returns a boolean if a field has been set.

### GetId

`func (o *ArpDetectVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArpDetectVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArpDetectVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArpDetectVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpbs

`func (o *ArpDetectVO) GetImpbs() []ImpbVO`

GetImpbs returns the Impbs field if non-nil, zero value otherwise.

### GetImpbsOk

`func (o *ArpDetectVO) GetImpbsOk() (*[]ImpbVO, bool)`

GetImpbsOk returns a tuple with the Impbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpbs

`func (o *ArpDetectVO) SetImpbs(v []ImpbVO)`

SetImpbs sets Impbs field to given value.

### HasImpbs

`func (o *ArpDetectVO) HasImpbs() bool`

HasImpbs returns a boolean if a field has been set.

### GetMac

`func (o *ArpDetectVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ArpDetectVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ArpDetectVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ArpDetectVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *ArpDetectVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArpDetectVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArpDetectVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArpDetectVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *ArpDetectVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *ArpDetectVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *ArpDetectVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *ArpDetectVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetOswStack

`func (o *ArpDetectVO) GetOswStack() OswStackInfoVO`

GetOswStack returns the OswStack field if non-nil, zero value otherwise.

### GetOswStackOk

`func (o *ArpDetectVO) GetOswStackOk() (*OswStackInfoVO, bool)`

GetOswStackOk returns a tuple with the OswStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswStack

`func (o *ArpDetectVO) SetOswStack(v OswStackInfoVO)`

SetOswStack sets OswStack field to given value.

### HasOswStack

`func (o *ArpDetectVO) HasOswStack() bool`

HasOswStack returns a boolean if a field has been set.

### GetPorts

`func (o *ArpDetectVO) GetPorts() []PortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ArpDetectVO) GetPortsOk() (*[]PortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ArpDetectVO) SetPorts(v []PortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ArpDetectVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetShowModel

`func (o *ArpDetectVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *ArpDetectVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *ArpDetectVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *ArpDetectVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSiteId

`func (o *ArpDetectVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ArpDetectVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ArpDetectVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ArpDetectVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStackId

`func (o *ArpDetectVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *ArpDetectVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *ArpDetectVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *ArpDetectVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *ArpDetectVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *ArpDetectVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *ArpDetectVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *ArpDetectVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStatus

`func (o *ArpDetectVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArpDetectVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArpDetectVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArpDetectVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUnSelectedablePorts

`func (o *ArpDetectVO) GetUnSelectedablePorts() []PortVO`

GetUnSelectedablePorts returns the UnSelectedablePorts field if non-nil, zero value otherwise.

### GetUnSelectedablePortsOk

`func (o *ArpDetectVO) GetUnSelectedablePortsOk() (*[]PortVO, bool)`

GetUnSelectedablePortsOk returns a tuple with the UnSelectedablePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSelectedablePorts

`func (o *ArpDetectVO) SetUnSelectedablePorts(v []PortVO)`

SetUnSelectedablePorts sets UnSelectedablePorts field to given value.

### HasUnSelectedablePorts

`func (o *ArpDetectVO) HasUnSelectedablePorts() bool`

HasUnSelectedablePorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


