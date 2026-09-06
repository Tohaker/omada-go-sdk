# OswStackDataVOOswDataVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbnormalReason** | Pointer to **int32** |  | [optional] 
**CustomLagIds** | Pointer to **[]int32** | The lag ids that has some vlan not included in port vlan. | [optional] 
**CustomStandardPorts** | Pointer to **[]string** | The standard ports that has some vlan not included in port vlan. | [optional] 
**DevCap** | Pointer to [**OswDevCapVO**](OswDevCapVO.md) |  | [optional] 
**Lags** | Pointer to [**[]OswLagStatusVO**](OswLagStatusVO.md) |  | [optional] 
**MasterMac** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**[]OswDataVO**](OswDataVO.md) |  | [optional] 
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

### NewOswStackDataVOOswDataVO

`func NewOswStackDataVOOswDataVO() *OswStackDataVOOswDataVO`

NewOswStackDataVOOswDataVO instantiates a new OswStackDataVOOswDataVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackDataVOOswDataVOWithDefaults

`func NewOswStackDataVOOswDataVOWithDefaults() *OswStackDataVOOswDataVO`

NewOswStackDataVOOswDataVOWithDefaults instantiates a new OswStackDataVOOswDataVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbnormalReason

`func (o *OswStackDataVOOswDataVO) GetAbnormalReason() int32`

GetAbnormalReason returns the AbnormalReason field if non-nil, zero value otherwise.

### GetAbnormalReasonOk

`func (o *OswStackDataVOOswDataVO) GetAbnormalReasonOk() (*int32, bool)`

GetAbnormalReasonOk returns a tuple with the AbnormalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalReason

`func (o *OswStackDataVOOswDataVO) SetAbnormalReason(v int32)`

SetAbnormalReason sets AbnormalReason field to given value.

### HasAbnormalReason

`func (o *OswStackDataVOOswDataVO) HasAbnormalReason() bool`

HasAbnormalReason returns a boolean if a field has been set.

### GetCustomLagIds

`func (o *OswStackDataVOOswDataVO) GetCustomLagIds() []int32`

GetCustomLagIds returns the CustomLagIds field if non-nil, zero value otherwise.

### GetCustomLagIdsOk

`func (o *OswStackDataVOOswDataVO) GetCustomLagIdsOk() (*[]int32, bool)`

GetCustomLagIdsOk returns a tuple with the CustomLagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLagIds

`func (o *OswStackDataVOOswDataVO) SetCustomLagIds(v []int32)`

SetCustomLagIds sets CustomLagIds field to given value.

### HasCustomLagIds

`func (o *OswStackDataVOOswDataVO) HasCustomLagIds() bool`

HasCustomLagIds returns a boolean if a field has been set.

### GetCustomStandardPorts

`func (o *OswStackDataVOOswDataVO) GetCustomStandardPorts() []string`

GetCustomStandardPorts returns the CustomStandardPorts field if non-nil, zero value otherwise.

### GetCustomStandardPortsOk

`func (o *OswStackDataVOOswDataVO) GetCustomStandardPortsOk() (*[]string, bool)`

GetCustomStandardPortsOk returns a tuple with the CustomStandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStandardPorts

`func (o *OswStackDataVOOswDataVO) SetCustomStandardPorts(v []string)`

SetCustomStandardPorts sets CustomStandardPorts field to given value.

### HasCustomStandardPorts

`func (o *OswStackDataVOOswDataVO) HasCustomStandardPorts() bool`

HasCustomStandardPorts returns a boolean if a field has been set.

### GetDevCap

`func (o *OswStackDataVOOswDataVO) GetDevCap() OswDevCapVO`

GetDevCap returns the DevCap field if non-nil, zero value otherwise.

### GetDevCapOk

`func (o *OswStackDataVOOswDataVO) GetDevCapOk() (*OswDevCapVO, bool)`

GetDevCapOk returns a tuple with the DevCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevCap

`func (o *OswStackDataVOOswDataVO) SetDevCap(v OswDevCapVO)`

SetDevCap sets DevCap field to given value.

### HasDevCap

`func (o *OswStackDataVOOswDataVO) HasDevCap() bool`

HasDevCap returns a boolean if a field has been set.

### GetLags

`func (o *OswStackDataVOOswDataVO) GetLags() []OswLagStatusVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *OswStackDataVOOswDataVO) GetLagsOk() (*[]OswLagStatusVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *OswStackDataVOOswDataVO) SetLags(v []OswLagStatusVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *OswStackDataVOOswDataVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetMasterMac

`func (o *OswStackDataVOOswDataVO) GetMasterMac() string`

GetMasterMac returns the MasterMac field if non-nil, zero value otherwise.

### GetMasterMacOk

`func (o *OswStackDataVOOswDataVO) GetMasterMacOk() (*string, bool)`

GetMasterMacOk returns a tuple with the MasterMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterMac

`func (o *OswStackDataVOOswDataVO) SetMasterMac(v string)`

SetMasterMac sets MasterMac field to given value.

### HasMasterMac

`func (o *OswStackDataVOOswDataVO) HasMasterMac() bool`

HasMasterMac returns a boolean if a field has been set.

### GetMember

`func (o *OswStackDataVOOswDataVO) GetMember() []OswDataVO`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *OswStackDataVOOswDataVO) GetMemberOk() (*[]OswDataVO, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *OswStackDataVOOswDataVO) SetMember(v []OswDataVO)`

SetMember sets Member field to given value.

### HasMember

`func (o *OswStackDataVOOswDataVO) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetModel

`func (o *OswStackDataVOOswDataVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswStackDataVOOswDataVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswStackDataVOOswDataVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswStackDataVOOswDataVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswStackDataVOOswDataVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswStackDataVOOswDataVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswStackDataVOOswDataVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswStackDataVOOswDataVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetShowModel

`func (o *OswStackDataVOOswDataVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OswStackDataVOOswDataVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OswStackDataVOOswDataVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OswStackDataVOOswDataVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStackId

`func (o *OswStackDataVOOswDataVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswStackDataVOOswDataVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswStackDataVOOswDataVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswStackDataVOOswDataVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *OswStackDataVOOswDataVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *OswStackDataVOOswDataVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *OswStackDataVOOswDataVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *OswStackDataVOOswDataVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStackStatus

`func (o *OswStackDataVOOswDataVO) GetStackStatus() int32`

GetStackStatus returns the StackStatus field if non-nil, zero value otherwise.

### GetStackStatusOk

`func (o *OswStackDataVOOswDataVO) GetStackStatusOk() (*int32, bool)`

GetStackStatusOk returns a tuple with the StackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackStatus

`func (o *OswStackDataVOOswDataVO) SetStackStatus(v int32)`

SetStackStatus sets StackStatus field to given value.

### HasStackStatus

`func (o *OswStackDataVOOswDataVO) HasStackStatus() bool`

HasStackStatus returns a boolean if a field has been set.

### GetSupportAutoAddVlan

`func (o *OswStackDataVOOswDataVO) GetSupportAutoAddVlan() bool`

GetSupportAutoAddVlan returns the SupportAutoAddVlan field if non-nil, zero value otherwise.

### GetSupportAutoAddVlanOk

`func (o *OswStackDataVOOswDataVO) GetSupportAutoAddVlanOk() (*bool, bool)`

GetSupportAutoAddVlanOk returns a tuple with the SupportAutoAddVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAutoAddVlan

`func (o *OswStackDataVOOswDataVO) SetSupportAutoAddVlan(v bool)`

SetSupportAutoAddVlan sets SupportAutoAddVlan field to given value.

### HasSupportAutoAddVlan

`func (o *OswStackDataVOOswDataVO) HasSupportAutoAddVlan() bool`

HasSupportAutoAddVlan returns a boolean if a field has been set.

### GetSupportCustomDhcpOption

`func (o *OswStackDataVOOswDataVO) GetSupportCustomDhcpOption() bool`

GetSupportCustomDhcpOption returns the SupportCustomDhcpOption field if non-nil, zero value otherwise.

### GetSupportCustomDhcpOptionOk

`func (o *OswStackDataVOOswDataVO) GetSupportCustomDhcpOptionOk() (*bool, bool)`

GetSupportCustomDhcpOptionOk returns a tuple with the SupportCustomDhcpOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomDhcpOption

`func (o *OswStackDataVOOswDataVO) SetSupportCustomDhcpOption(v bool)`

SetSupportCustomDhcpOption sets SupportCustomDhcpOption field to given value.

### HasSupportCustomDhcpOption

`func (o *OswStackDataVOOswDataVO) HasSupportCustomDhcpOption() bool`

HasSupportCustomDhcpOption returns a boolean if a field has been set.

### GetSupportDhcpRange

`func (o *OswStackDataVOOswDataVO) GetSupportDhcpRange() bool`

GetSupportDhcpRange returns the SupportDhcpRange field if non-nil, zero value otherwise.

### GetSupportDhcpRangeOk

`func (o *OswStackDataVOOswDataVO) GetSupportDhcpRangeOk() (*bool, bool)`

GetSupportDhcpRangeOk returns a tuple with the SupportDhcpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDhcpRange

`func (o *OswStackDataVOOswDataVO) SetSupportDhcpRange(v bool)`

SetSupportDhcpRange sets SupportDhcpRange field to given value.

### HasSupportDhcpRange

`func (o *OswStackDataVOOswDataVO) HasSupportDhcpRange() bool`

HasSupportDhcpRange returns a boolean if a field has been set.

### GetSupportLayout

`func (o *OswStackDataVOOswDataVO) GetSupportLayout() bool`

GetSupportLayout returns the SupportLayout field if non-nil, zero value otherwise.

### GetSupportLayoutOk

`func (o *OswStackDataVOOswDataVO) GetSupportLayoutOk() (*bool, bool)`

GetSupportLayoutOk returns a tuple with the SupportLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLayout

`func (o *OswStackDataVOOswDataVO) SetSupportLayout(v bool)`

SetSupportLayout sets SupportLayout field to given value.

### HasSupportLayout

`func (o *OswStackDataVOOswDataVO) HasSupportLayout() bool`

HasSupportLayout returns a boolean if a field has been set.

### GetSupportVrf

`func (o *OswStackDataVOOswDataVO) GetSupportVrf() bool`

GetSupportVrf returns the SupportVrf field if non-nil, zero value otherwise.

### GetSupportVrfOk

`func (o *OswStackDataVOOswDataVO) GetSupportVrfOk() (*bool, bool)`

GetSupportVrfOk returns a tuple with the SupportVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVrf

`func (o *OswStackDataVOOswDataVO) SetSupportVrf(v bool)`

SetSupportVrf sets SupportVrf field to given value.

### HasSupportVrf

`func (o *OswStackDataVOOswDataVO) HasSupportVrf() bool`

HasSupportVrf returns a boolean if a field has been set.

### GetUnSelectedablePorts

`func (o *OswStackDataVOOswDataVO) GetUnSelectedablePorts() []PortVO`

GetUnSelectedablePorts returns the UnSelectedablePorts field if non-nil, zero value otherwise.

### GetUnSelectedablePortsOk

`func (o *OswStackDataVOOswDataVO) GetUnSelectedablePortsOk() (*[]PortVO, bool)`

GetUnSelectedablePortsOk returns a tuple with the UnSelectedablePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSelectedablePorts

`func (o *OswStackDataVOOswDataVO) SetUnSelectedablePorts(v []PortVO)`

SetUnSelectedablePorts sets UnSelectedablePorts field to given value.

### HasUnSelectedablePorts

`func (o *OswStackDataVOOswDataVO) HasUnSelectedablePorts() bool`

HasUnSelectedablePorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


