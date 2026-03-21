# OswDataVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpSnoopPorts** | Pointer to [**[]PortVO**](PortVO.md) | The ports dhcp snooping selected. | [optional] 
**Lags** | Pointer to [**[]OswLagVO**](OswLagVO.md) | Lag List | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**NetworkNotAssociatedLags** | Pointer to **[]int32** | The lags not associated with network. | [optional] 
**NetworkNotAssociatedStandardPorts** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | The ports not associated with network. | [optional] 
**PortNum** | Pointer to **int32** | The number of ports of one device. | [optional] 
**Ports** | Pointer to [**[]OswPortVO**](OswPortVO.md) | Port List | [optional] 
**ShowModel** | Pointer to **string** | Model complex shown in the front end.Ap：model+(country)+modelVersion,EAP225(EU) v3.0  Gateway/Switch：model+modelVersion,Osg v3.0 | [optional] 
**StackId** | Pointer to **string** | Stack Id | [optional] 
**StackName** | Pointer to **string** | Stack name | [optional] 
**StackOswData** | Pointer to [**OswStackDataVOOswDataVO**](OswStackDataVOOswDataVO.md) |  | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**UnSelectedablePorts** | Pointer to [**[]PortVO**](PortVO.md) | The unSelectedable ports of the device. | [optional] 
**Unit** | Pointer to **int32** | Unit | [optional] 
**Uplink** | Pointer to [**OswUplinkVO**](OswUplinkVO.md) |  | [optional] 

## Methods

### NewOswDataVO

`func NewOswDataVO() *OswDataVO`

NewOswDataVO instantiates a new OswDataVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswDataVOWithDefaults

`func NewOswDataVOWithDefaults() *OswDataVO`

NewOswDataVOWithDefaults instantiates a new OswDataVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpSnoopPorts

`func (o *OswDataVO) GetDhcpSnoopPorts() []PortVO`

GetDhcpSnoopPorts returns the DhcpSnoopPorts field if non-nil, zero value otherwise.

### GetDhcpSnoopPortsOk

`func (o *OswDataVO) GetDhcpSnoopPortsOk() (*[]PortVO, bool)`

GetDhcpSnoopPortsOk returns a tuple with the DhcpSnoopPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSnoopPorts

`func (o *OswDataVO) SetDhcpSnoopPorts(v []PortVO)`

SetDhcpSnoopPorts sets DhcpSnoopPorts field to given value.

### HasDhcpSnoopPorts

`func (o *OswDataVO) HasDhcpSnoopPorts() bool`

HasDhcpSnoopPorts returns a boolean if a field has been set.

### GetLags

`func (o *OswDataVO) GetLags() []OswLagVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *OswDataVO) GetLagsOk() (*[]OswLagVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *OswDataVO) SetLags(v []OswLagVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *OswDataVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetMac

`func (o *OswDataVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswDataVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswDataVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswDataVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *OswDataVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswDataVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswDataVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswDataVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswDataVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswDataVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswDataVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswDataVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *OswDataVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswDataVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswDataVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswDataVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkNotAssociatedLags

`func (o *OswDataVO) GetNetworkNotAssociatedLags() []int32`

GetNetworkNotAssociatedLags returns the NetworkNotAssociatedLags field if non-nil, zero value otherwise.

### GetNetworkNotAssociatedLagsOk

`func (o *OswDataVO) GetNetworkNotAssociatedLagsOk() (*[]int32, bool)`

GetNetworkNotAssociatedLagsOk returns a tuple with the NetworkNotAssociatedLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkNotAssociatedLags

`func (o *OswDataVO) SetNetworkNotAssociatedLags(v []int32)`

SetNetworkNotAssociatedLags sets NetworkNotAssociatedLags field to given value.

### HasNetworkNotAssociatedLags

`func (o *OswDataVO) HasNetworkNotAssociatedLags() bool`

HasNetworkNotAssociatedLags returns a boolean if a field has been set.

### GetNetworkNotAssociatedStandardPorts

`func (o *OswDataVO) GetNetworkNotAssociatedStandardPorts() []OswStandPortVO`

GetNetworkNotAssociatedStandardPorts returns the NetworkNotAssociatedStandardPorts field if non-nil, zero value otherwise.

### GetNetworkNotAssociatedStandardPortsOk

`func (o *OswDataVO) GetNetworkNotAssociatedStandardPortsOk() (*[]OswStandPortVO, bool)`

GetNetworkNotAssociatedStandardPortsOk returns a tuple with the NetworkNotAssociatedStandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkNotAssociatedStandardPorts

`func (o *OswDataVO) SetNetworkNotAssociatedStandardPorts(v []OswStandPortVO)`

SetNetworkNotAssociatedStandardPorts sets NetworkNotAssociatedStandardPorts field to given value.

### HasNetworkNotAssociatedStandardPorts

`func (o *OswDataVO) HasNetworkNotAssociatedStandardPorts() bool`

HasNetworkNotAssociatedStandardPorts returns a boolean if a field has been set.

### GetPortNum

`func (o *OswDataVO) GetPortNum() int32`

GetPortNum returns the PortNum field if non-nil, zero value otherwise.

### GetPortNumOk

`func (o *OswDataVO) GetPortNumOk() (*int32, bool)`

GetPortNumOk returns a tuple with the PortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNum

`func (o *OswDataVO) SetPortNum(v int32)`

SetPortNum sets PortNum field to given value.

### HasPortNum

`func (o *OswDataVO) HasPortNum() bool`

HasPortNum returns a boolean if a field has been set.

### GetPorts

`func (o *OswDataVO) GetPorts() []OswPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswDataVO) GetPortsOk() (*[]OswPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswDataVO) SetPorts(v []OswPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswDataVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetShowModel

`func (o *OswDataVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OswDataVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OswDataVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OswDataVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStackId

`func (o *OswDataVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswDataVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswDataVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswDataVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *OswDataVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *OswDataVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *OswDataVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *OswDataVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStackOswData

`func (o *OswDataVO) GetStackOswData() OswStackDataVOOswDataVO`

GetStackOswData returns the StackOswData field if non-nil, zero value otherwise.

### GetStackOswDataOk

`func (o *OswDataVO) GetStackOswDataOk() (*OswStackDataVOOswDataVO, bool)`

GetStackOswDataOk returns a tuple with the StackOswData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackOswData

`func (o *OswDataVO) SetStackOswData(v OswStackDataVOOswDataVO)`

SetStackOswData sets StackOswData field to given value.

### HasStackOswData

`func (o *OswDataVO) HasStackOswData() bool`

HasStackOswData returns a boolean if a field has been set.

### GetStatus

`func (o *OswDataVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswDataVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswDataVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswDataVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *OswDataVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswDataVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswDataVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OswDataVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnSelectedablePorts

`func (o *OswDataVO) GetUnSelectedablePorts() []PortVO`

GetUnSelectedablePorts returns the UnSelectedablePorts field if non-nil, zero value otherwise.

### GetUnSelectedablePortsOk

`func (o *OswDataVO) GetUnSelectedablePortsOk() (*[]PortVO, bool)`

GetUnSelectedablePortsOk returns a tuple with the UnSelectedablePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSelectedablePorts

`func (o *OswDataVO) SetUnSelectedablePorts(v []PortVO)`

SetUnSelectedablePorts sets UnSelectedablePorts field to given value.

### HasUnSelectedablePorts

`func (o *OswDataVO) HasUnSelectedablePorts() bool`

HasUnSelectedablePorts returns a boolean if a field has been set.

### GetUnit

`func (o *OswDataVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswDataVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswDataVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswDataVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUplink

`func (o *OswDataVO) GetUplink() OswUplinkVO`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *OswDataVO) GetUplinkOk() (*OswUplinkVO, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *OswDataVO) SetUplink(v OswUplinkVO)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *OswDataVO) HasUplink() bool`

HasUplink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


