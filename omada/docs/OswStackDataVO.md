# OswStackDataVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbnormalReason** | Pointer to **int32** |  | [optional] 
**CustomLagIds** | Pointer to **[]int32** | The lag ids that has some vlan not included in port vlan. | [optional] 
**CustomStandardPorts** | Pointer to **[]string** | The standard ports that has some vlan not included in port vlan. | [optional] 
**DevCap** | Pointer to [**OswDevCapVO**](OswDevCapVO.md) |  | [optional] 
**Lags** | Pointer to [**[]OswLagStatusVO**](OswLagStatusVO.md) |  | [optional] 
**MasterMac** | Pointer to **string** |  | [optional] 
**Member** | Pointer to **[]map[string]interface{}** |  | [optional] 
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

### NewOswStackDataVO

`func NewOswStackDataVO() *OswStackDataVO`

NewOswStackDataVO instantiates a new OswStackDataVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackDataVOWithDefaults

`func NewOswStackDataVOWithDefaults() *OswStackDataVO`

NewOswStackDataVOWithDefaults instantiates a new OswStackDataVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbnormalReason

`func (o *OswStackDataVO) GetAbnormalReason() int32`

GetAbnormalReason returns the AbnormalReason field if non-nil, zero value otherwise.

### GetAbnormalReasonOk

`func (o *OswStackDataVO) GetAbnormalReasonOk() (*int32, bool)`

GetAbnormalReasonOk returns a tuple with the AbnormalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalReason

`func (o *OswStackDataVO) SetAbnormalReason(v int32)`

SetAbnormalReason sets AbnormalReason field to given value.

### HasAbnormalReason

`func (o *OswStackDataVO) HasAbnormalReason() bool`

HasAbnormalReason returns a boolean if a field has been set.

### GetCustomLagIds

`func (o *OswStackDataVO) GetCustomLagIds() []int32`

GetCustomLagIds returns the CustomLagIds field if non-nil, zero value otherwise.

### GetCustomLagIdsOk

`func (o *OswStackDataVO) GetCustomLagIdsOk() (*[]int32, bool)`

GetCustomLagIdsOk returns a tuple with the CustomLagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLagIds

`func (o *OswStackDataVO) SetCustomLagIds(v []int32)`

SetCustomLagIds sets CustomLagIds field to given value.

### HasCustomLagIds

`func (o *OswStackDataVO) HasCustomLagIds() bool`

HasCustomLagIds returns a boolean if a field has been set.

### GetCustomStandardPorts

`func (o *OswStackDataVO) GetCustomStandardPorts() []string`

GetCustomStandardPorts returns the CustomStandardPorts field if non-nil, zero value otherwise.

### GetCustomStandardPortsOk

`func (o *OswStackDataVO) GetCustomStandardPortsOk() (*[]string, bool)`

GetCustomStandardPortsOk returns a tuple with the CustomStandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStandardPorts

`func (o *OswStackDataVO) SetCustomStandardPorts(v []string)`

SetCustomStandardPorts sets CustomStandardPorts field to given value.

### HasCustomStandardPorts

`func (o *OswStackDataVO) HasCustomStandardPorts() bool`

HasCustomStandardPorts returns a boolean if a field has been set.

### GetDevCap

`func (o *OswStackDataVO) GetDevCap() OswDevCapVO`

GetDevCap returns the DevCap field if non-nil, zero value otherwise.

### GetDevCapOk

`func (o *OswStackDataVO) GetDevCapOk() (*OswDevCapVO, bool)`

GetDevCapOk returns a tuple with the DevCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevCap

`func (o *OswStackDataVO) SetDevCap(v OswDevCapVO)`

SetDevCap sets DevCap field to given value.

### HasDevCap

`func (o *OswStackDataVO) HasDevCap() bool`

HasDevCap returns a boolean if a field has been set.

### GetLags

`func (o *OswStackDataVO) GetLags() []OswLagStatusVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *OswStackDataVO) GetLagsOk() (*[]OswLagStatusVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *OswStackDataVO) SetLags(v []OswLagStatusVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *OswStackDataVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetMasterMac

`func (o *OswStackDataVO) GetMasterMac() string`

GetMasterMac returns the MasterMac field if non-nil, zero value otherwise.

### GetMasterMacOk

`func (o *OswStackDataVO) GetMasterMacOk() (*string, bool)`

GetMasterMacOk returns a tuple with the MasterMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterMac

`func (o *OswStackDataVO) SetMasterMac(v string)`

SetMasterMac sets MasterMac field to given value.

### HasMasterMac

`func (o *OswStackDataVO) HasMasterMac() bool`

HasMasterMac returns a boolean if a field has been set.

### GetMember

`func (o *OswStackDataVO) GetMember() []map[string]interface{}`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *OswStackDataVO) GetMemberOk() (*[]map[string]interface{}, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *OswStackDataVO) SetMember(v []map[string]interface{})`

SetMember sets Member field to given value.

### HasMember

`func (o *OswStackDataVO) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetModel

`func (o *OswStackDataVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswStackDataVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswStackDataVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswStackDataVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswStackDataVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswStackDataVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswStackDataVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswStackDataVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetShowModel

`func (o *OswStackDataVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OswStackDataVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OswStackDataVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OswStackDataVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStackId

`func (o *OswStackDataVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswStackDataVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswStackDataVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswStackDataVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *OswStackDataVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *OswStackDataVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *OswStackDataVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *OswStackDataVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStackStatus

`func (o *OswStackDataVO) GetStackStatus() int32`

GetStackStatus returns the StackStatus field if non-nil, zero value otherwise.

### GetStackStatusOk

`func (o *OswStackDataVO) GetStackStatusOk() (*int32, bool)`

GetStackStatusOk returns a tuple with the StackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackStatus

`func (o *OswStackDataVO) SetStackStatus(v int32)`

SetStackStatus sets StackStatus field to given value.

### HasStackStatus

`func (o *OswStackDataVO) HasStackStatus() bool`

HasStackStatus returns a boolean if a field has been set.

### GetSupportAutoAddVlan

`func (o *OswStackDataVO) GetSupportAutoAddVlan() bool`

GetSupportAutoAddVlan returns the SupportAutoAddVlan field if non-nil, zero value otherwise.

### GetSupportAutoAddVlanOk

`func (o *OswStackDataVO) GetSupportAutoAddVlanOk() (*bool, bool)`

GetSupportAutoAddVlanOk returns a tuple with the SupportAutoAddVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAutoAddVlan

`func (o *OswStackDataVO) SetSupportAutoAddVlan(v bool)`

SetSupportAutoAddVlan sets SupportAutoAddVlan field to given value.

### HasSupportAutoAddVlan

`func (o *OswStackDataVO) HasSupportAutoAddVlan() bool`

HasSupportAutoAddVlan returns a boolean if a field has been set.

### GetSupportCustomDhcpOption

`func (o *OswStackDataVO) GetSupportCustomDhcpOption() bool`

GetSupportCustomDhcpOption returns the SupportCustomDhcpOption field if non-nil, zero value otherwise.

### GetSupportCustomDhcpOptionOk

`func (o *OswStackDataVO) GetSupportCustomDhcpOptionOk() (*bool, bool)`

GetSupportCustomDhcpOptionOk returns a tuple with the SupportCustomDhcpOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomDhcpOption

`func (o *OswStackDataVO) SetSupportCustomDhcpOption(v bool)`

SetSupportCustomDhcpOption sets SupportCustomDhcpOption field to given value.

### HasSupportCustomDhcpOption

`func (o *OswStackDataVO) HasSupportCustomDhcpOption() bool`

HasSupportCustomDhcpOption returns a boolean if a field has been set.

### GetSupportDhcpRange

`func (o *OswStackDataVO) GetSupportDhcpRange() bool`

GetSupportDhcpRange returns the SupportDhcpRange field if non-nil, zero value otherwise.

### GetSupportDhcpRangeOk

`func (o *OswStackDataVO) GetSupportDhcpRangeOk() (*bool, bool)`

GetSupportDhcpRangeOk returns a tuple with the SupportDhcpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDhcpRange

`func (o *OswStackDataVO) SetSupportDhcpRange(v bool)`

SetSupportDhcpRange sets SupportDhcpRange field to given value.

### HasSupportDhcpRange

`func (o *OswStackDataVO) HasSupportDhcpRange() bool`

HasSupportDhcpRange returns a boolean if a field has been set.

### GetSupportLayout

`func (o *OswStackDataVO) GetSupportLayout() bool`

GetSupportLayout returns the SupportLayout field if non-nil, zero value otherwise.

### GetSupportLayoutOk

`func (o *OswStackDataVO) GetSupportLayoutOk() (*bool, bool)`

GetSupportLayoutOk returns a tuple with the SupportLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLayout

`func (o *OswStackDataVO) SetSupportLayout(v bool)`

SetSupportLayout sets SupportLayout field to given value.

### HasSupportLayout

`func (o *OswStackDataVO) HasSupportLayout() bool`

HasSupportLayout returns a boolean if a field has been set.

### GetSupportVrf

`func (o *OswStackDataVO) GetSupportVrf() bool`

GetSupportVrf returns the SupportVrf field if non-nil, zero value otherwise.

### GetSupportVrfOk

`func (o *OswStackDataVO) GetSupportVrfOk() (*bool, bool)`

GetSupportVrfOk returns a tuple with the SupportVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVrf

`func (o *OswStackDataVO) SetSupportVrf(v bool)`

SetSupportVrf sets SupportVrf field to given value.

### HasSupportVrf

`func (o *OswStackDataVO) HasSupportVrf() bool`

HasSupportVrf returns a boolean if a field has been set.

### GetUnSelectedablePorts

`func (o *OswStackDataVO) GetUnSelectedablePorts() []PortVO`

GetUnSelectedablePorts returns the UnSelectedablePorts field if non-nil, zero value otherwise.

### GetUnSelectedablePortsOk

`func (o *OswStackDataVO) GetUnSelectedablePortsOk() (*[]PortVO, bool)`

GetUnSelectedablePortsOk returns a tuple with the UnSelectedablePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSelectedablePorts

`func (o *OswStackDataVO) SetUnSelectedablePorts(v []PortVO)`

SetUnSelectedablePorts sets UnSelectedablePorts field to given value.

### HasUnSelectedablePorts

`func (o *OswStackDataVO) HasUnSelectedablePorts() bool`

HasUnSelectedablePorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


