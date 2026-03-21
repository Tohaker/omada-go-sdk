# OswStackDataVOOswDataVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbnormalReason** | Pointer to **int32** |  | [optional] 
**DevCap** | Pointer to [**OswDevCapVO**](OswDevCapVO.md) |  | [optional] 
**Lags** | Pointer to [**[]OswLagStatusVO**](OswLagStatusVO.md) |  | [optional] 
**MasterMac** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**[]OswDataVO**](OswDataVO.md) |  | [optional] 
**StackId** | Pointer to **string** |  | [optional] 
**StackName** | Pointer to **string** |  | [optional] 
**StackStatus** | Pointer to **int32** |  | [optional] 
**SupportCustomDhcpOption** | Pointer to **bool** |  | [optional] 
**SupportDhcpRange** | Pointer to **bool** |  | [optional] 
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


