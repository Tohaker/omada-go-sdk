# OswStackDataVOSwitchDataVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbnormalReason** | Pointer to **int32** |  | [optional] 
**CustomLagIds** | Pointer to **[]int32** | The lag ids that has some vlan not included in port vlan. | [optional] 
**CustomStandardPorts** | Pointer to **[]string** | The standard ports that has some vlan not included in port vlan. | [optional] 
**DevCap** | Pointer to [**OswDevCapVO**](OswDevCapVO.md) |  | [optional] 
**Lags** | Pointer to [**[]OswLagStatusVO**](OswLagStatusVO.md) |  | [optional] 
**MasterMac** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**[]SwitchDataVO**](SwitchDataVO.md) |  | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**ShowModel** | Pointer to **string** | Model complex shown in the front end.Ap：model+(country)+modelVersion,EAP225(EU) v3.0  Gateway/Switch：model+modelVersion,Osg v3.0 | [optional] 
**StackId** | Pointer to **string** |  | [optional] 
**StackName** | Pointer to **string** |  | [optional] 
**StackStatus** | Pointer to **int32** |  | [optional] 
**SupportAutoAddVlan** | Pointer to **bool** |  | [optional] 
**SupportCustomDhcpOption** | Pointer to **bool** |  | [optional] 
**SupportDhcpRange** | Pointer to **bool** |  | [optional] 
**SupportLayout** | Pointer to **bool** | Whether the device supports reporting port layout information. | [optional] 
**SupportVrf** | Pointer to **bool** |  | [optional] 
**UnSelectedablePorts** | Pointer to [**[]PortVO**](PortVO.md) | The unSelectedable ports of the device. | [optional] 

## Methods

### NewOswStackDataVOSwitchDataVO

`func NewOswStackDataVOSwitchDataVO() *OswStackDataVOSwitchDataVO`

NewOswStackDataVOSwitchDataVO instantiates a new OswStackDataVOSwitchDataVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackDataVOSwitchDataVOWithDefaults

`func NewOswStackDataVOSwitchDataVOWithDefaults() *OswStackDataVOSwitchDataVO`

NewOswStackDataVOSwitchDataVOWithDefaults instantiates a new OswStackDataVOSwitchDataVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbnormalReason

`func (o *OswStackDataVOSwitchDataVO) GetAbnormalReason() int32`

GetAbnormalReason returns the AbnormalReason field if non-nil, zero value otherwise.

### GetAbnormalReasonOk

`func (o *OswStackDataVOSwitchDataVO) GetAbnormalReasonOk() (*int32, bool)`

GetAbnormalReasonOk returns a tuple with the AbnormalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalReason

`func (o *OswStackDataVOSwitchDataVO) SetAbnormalReason(v int32)`

SetAbnormalReason sets AbnormalReason field to given value.

### HasAbnormalReason

`func (o *OswStackDataVOSwitchDataVO) HasAbnormalReason() bool`

HasAbnormalReason returns a boolean if a field has been set.

### GetCustomLagIds

`func (o *OswStackDataVOSwitchDataVO) GetCustomLagIds() []int32`

GetCustomLagIds returns the CustomLagIds field if non-nil, zero value otherwise.

### GetCustomLagIdsOk

`func (o *OswStackDataVOSwitchDataVO) GetCustomLagIdsOk() (*[]int32, bool)`

GetCustomLagIdsOk returns a tuple with the CustomLagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLagIds

`func (o *OswStackDataVOSwitchDataVO) SetCustomLagIds(v []int32)`

SetCustomLagIds sets CustomLagIds field to given value.

### HasCustomLagIds

`func (o *OswStackDataVOSwitchDataVO) HasCustomLagIds() bool`

HasCustomLagIds returns a boolean if a field has been set.

### GetCustomStandardPorts

`func (o *OswStackDataVOSwitchDataVO) GetCustomStandardPorts() []string`

GetCustomStandardPorts returns the CustomStandardPorts field if non-nil, zero value otherwise.

### GetCustomStandardPortsOk

`func (o *OswStackDataVOSwitchDataVO) GetCustomStandardPortsOk() (*[]string, bool)`

GetCustomStandardPortsOk returns a tuple with the CustomStandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStandardPorts

`func (o *OswStackDataVOSwitchDataVO) SetCustomStandardPorts(v []string)`

SetCustomStandardPorts sets CustomStandardPorts field to given value.

### HasCustomStandardPorts

`func (o *OswStackDataVOSwitchDataVO) HasCustomStandardPorts() bool`

HasCustomStandardPorts returns a boolean if a field has been set.

### GetDevCap

`func (o *OswStackDataVOSwitchDataVO) GetDevCap() OswDevCapVO`

GetDevCap returns the DevCap field if non-nil, zero value otherwise.

### GetDevCapOk

`func (o *OswStackDataVOSwitchDataVO) GetDevCapOk() (*OswDevCapVO, bool)`

GetDevCapOk returns a tuple with the DevCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevCap

`func (o *OswStackDataVOSwitchDataVO) SetDevCap(v OswDevCapVO)`

SetDevCap sets DevCap field to given value.

### HasDevCap

`func (o *OswStackDataVOSwitchDataVO) HasDevCap() bool`

HasDevCap returns a boolean if a field has been set.

### GetLags

`func (o *OswStackDataVOSwitchDataVO) GetLags() []OswLagStatusVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *OswStackDataVOSwitchDataVO) GetLagsOk() (*[]OswLagStatusVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *OswStackDataVOSwitchDataVO) SetLags(v []OswLagStatusVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *OswStackDataVOSwitchDataVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetMasterMac

`func (o *OswStackDataVOSwitchDataVO) GetMasterMac() string`

GetMasterMac returns the MasterMac field if non-nil, zero value otherwise.

### GetMasterMacOk

`func (o *OswStackDataVOSwitchDataVO) GetMasterMacOk() (*string, bool)`

GetMasterMacOk returns a tuple with the MasterMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterMac

`func (o *OswStackDataVOSwitchDataVO) SetMasterMac(v string)`

SetMasterMac sets MasterMac field to given value.

### HasMasterMac

`func (o *OswStackDataVOSwitchDataVO) HasMasterMac() bool`

HasMasterMac returns a boolean if a field has been set.

### GetMember

`func (o *OswStackDataVOSwitchDataVO) GetMember() []SwitchDataVO`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *OswStackDataVOSwitchDataVO) GetMemberOk() (*[]SwitchDataVO, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *OswStackDataVOSwitchDataVO) SetMember(v []SwitchDataVO)`

SetMember sets Member field to given value.

### HasMember

`func (o *OswStackDataVOSwitchDataVO) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetModel

`func (o *OswStackDataVOSwitchDataVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswStackDataVOSwitchDataVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswStackDataVOSwitchDataVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswStackDataVOSwitchDataVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswStackDataVOSwitchDataVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswStackDataVOSwitchDataVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswStackDataVOSwitchDataVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswStackDataVOSwitchDataVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetShowModel

`func (o *OswStackDataVOSwitchDataVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OswStackDataVOSwitchDataVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OswStackDataVOSwitchDataVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OswStackDataVOSwitchDataVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStackId

`func (o *OswStackDataVOSwitchDataVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswStackDataVOSwitchDataVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswStackDataVOSwitchDataVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswStackDataVOSwitchDataVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *OswStackDataVOSwitchDataVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *OswStackDataVOSwitchDataVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *OswStackDataVOSwitchDataVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *OswStackDataVOSwitchDataVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStackStatus

`func (o *OswStackDataVOSwitchDataVO) GetStackStatus() int32`

GetStackStatus returns the StackStatus field if non-nil, zero value otherwise.

### GetStackStatusOk

`func (o *OswStackDataVOSwitchDataVO) GetStackStatusOk() (*int32, bool)`

GetStackStatusOk returns a tuple with the StackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackStatus

`func (o *OswStackDataVOSwitchDataVO) SetStackStatus(v int32)`

SetStackStatus sets StackStatus field to given value.

### HasStackStatus

`func (o *OswStackDataVOSwitchDataVO) HasStackStatus() bool`

HasStackStatus returns a boolean if a field has been set.

### GetSupportAutoAddVlan

`func (o *OswStackDataVOSwitchDataVO) GetSupportAutoAddVlan() bool`

GetSupportAutoAddVlan returns the SupportAutoAddVlan field if non-nil, zero value otherwise.

### GetSupportAutoAddVlanOk

`func (o *OswStackDataVOSwitchDataVO) GetSupportAutoAddVlanOk() (*bool, bool)`

GetSupportAutoAddVlanOk returns a tuple with the SupportAutoAddVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAutoAddVlan

`func (o *OswStackDataVOSwitchDataVO) SetSupportAutoAddVlan(v bool)`

SetSupportAutoAddVlan sets SupportAutoAddVlan field to given value.

### HasSupportAutoAddVlan

`func (o *OswStackDataVOSwitchDataVO) HasSupportAutoAddVlan() bool`

HasSupportAutoAddVlan returns a boolean if a field has been set.

### GetSupportCustomDhcpOption

`func (o *OswStackDataVOSwitchDataVO) GetSupportCustomDhcpOption() bool`

GetSupportCustomDhcpOption returns the SupportCustomDhcpOption field if non-nil, zero value otherwise.

### GetSupportCustomDhcpOptionOk

`func (o *OswStackDataVOSwitchDataVO) GetSupportCustomDhcpOptionOk() (*bool, bool)`

GetSupportCustomDhcpOptionOk returns a tuple with the SupportCustomDhcpOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomDhcpOption

`func (o *OswStackDataVOSwitchDataVO) SetSupportCustomDhcpOption(v bool)`

SetSupportCustomDhcpOption sets SupportCustomDhcpOption field to given value.

### HasSupportCustomDhcpOption

`func (o *OswStackDataVOSwitchDataVO) HasSupportCustomDhcpOption() bool`

HasSupportCustomDhcpOption returns a boolean if a field has been set.

### GetSupportDhcpRange

`func (o *OswStackDataVOSwitchDataVO) GetSupportDhcpRange() bool`

GetSupportDhcpRange returns the SupportDhcpRange field if non-nil, zero value otherwise.

### GetSupportDhcpRangeOk

`func (o *OswStackDataVOSwitchDataVO) GetSupportDhcpRangeOk() (*bool, bool)`

GetSupportDhcpRangeOk returns a tuple with the SupportDhcpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDhcpRange

`func (o *OswStackDataVOSwitchDataVO) SetSupportDhcpRange(v bool)`

SetSupportDhcpRange sets SupportDhcpRange field to given value.

### HasSupportDhcpRange

`func (o *OswStackDataVOSwitchDataVO) HasSupportDhcpRange() bool`

HasSupportDhcpRange returns a boolean if a field has been set.

### GetSupportLayout

`func (o *OswStackDataVOSwitchDataVO) GetSupportLayout() bool`

GetSupportLayout returns the SupportLayout field if non-nil, zero value otherwise.

### GetSupportLayoutOk

`func (o *OswStackDataVOSwitchDataVO) GetSupportLayoutOk() (*bool, bool)`

GetSupportLayoutOk returns a tuple with the SupportLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLayout

`func (o *OswStackDataVOSwitchDataVO) SetSupportLayout(v bool)`

SetSupportLayout sets SupportLayout field to given value.

### HasSupportLayout

`func (o *OswStackDataVOSwitchDataVO) HasSupportLayout() bool`

HasSupportLayout returns a boolean if a field has been set.

### GetSupportVrf

`func (o *OswStackDataVOSwitchDataVO) GetSupportVrf() bool`

GetSupportVrf returns the SupportVrf field if non-nil, zero value otherwise.

### GetSupportVrfOk

`func (o *OswStackDataVOSwitchDataVO) GetSupportVrfOk() (*bool, bool)`

GetSupportVrfOk returns a tuple with the SupportVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVrf

`func (o *OswStackDataVOSwitchDataVO) SetSupportVrf(v bool)`

SetSupportVrf sets SupportVrf field to given value.

### HasSupportVrf

`func (o *OswStackDataVOSwitchDataVO) HasSupportVrf() bool`

HasSupportVrf returns a boolean if a field has been set.

### GetUnSelectedablePorts

`func (o *OswStackDataVOSwitchDataVO) GetUnSelectedablePorts() []PortVO`

GetUnSelectedablePorts returns the UnSelectedablePorts field if non-nil, zero value otherwise.

### GetUnSelectedablePortsOk

`func (o *OswStackDataVOSwitchDataVO) GetUnSelectedablePortsOk() (*[]PortVO, bool)`

GetUnSelectedablePortsOk returns a tuple with the UnSelectedablePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSelectedablePorts

`func (o *OswStackDataVOSwitchDataVO) SetUnSelectedablePorts(v []PortVO)`

SetUnSelectedablePorts sets UnSelectedablePorts field to given value.

### HasUnSelectedablePorts

`func (o *OswStackDataVOSwitchDataVO) HasUnSelectedablePorts() bool`

HasUnSelectedablePorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


