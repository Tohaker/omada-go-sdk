# OswPortsSettingPoeVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site to which the port belongs | [optional] 
**ConfigMlagDad** | Pointer to **bool** | Indicates whether the current port is configured as M-LAG dad Link | [optional] 
**ConfigMlagPeerLink** | Pointer to **bool** | Indicates whether the current port is configured as M-LAG peer Link | [optional] 
**ConfigStack** | Pointer to **bool** | Indicates whether the current port is configured as a stack port (joined a stack aggregation group) | [optional] 
**ConnectedStatus** | Pointer to **int32** | Port connected status should be a value as follows: 0: Connected; 1: Disconnected; 2: Disable | [optional] 
**Current** | Pointer to **float64** | Poe Current | [optional] 
**Disable** | Pointer to **bool** | Indicates whether to disable the port | [optional] 
**Duplex** | Pointer to **int32** | Duplex should be a value as follows: 0: Auto; 1: Half; 2: Full | [optional] 
**LagSetting** | Pointer to [**OswStackLagVO**](OswStackLagVO.md) |  | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed should be a value as follows: 0: auto; 1: 10M; 2: 100M; 3: 1000M; 4: 10G | [optional] 
**MadUsed** | Pointer to **bool** | Mad Used | [optional] 
**Operation** | Pointer to **string** | Operation should be a value as follows: SWITCHING; MIRRORING; AGGREGATING | [optional] 
**PdClass** | Pointer to **string** | Poe PD Class | [optional] 
**Poe** | Pointer to **int32** | PoE switch should be a value as follows: 0: Off; 1: 802.3at/af | [optional] 
**PoeDisplayType** | Pointer to **int32** | PoeDisplayType should be a value as follows: -1: Not Support POE; 0: Support POE; 1: POE(4W); 2: POE(7W); 3: POE(15.4W); 4: POE+(30W); 5: POE++(45W); 6: POE++(60W); 7: POE++(75W); 8: POE++(90W); 9: POE++(100W). | [optional] 
**PoeStatus** | Pointer to **float64** | Poe Status | [optional] 
**Port** | Pointer to **int32** | Port | [optional] 
**PortName** | Pointer to **string** | Port Name | [optional] 
**PortStatus** | Pointer to [**OswPortStatusVO**](OswPortStatusVO.md) |  | [optional] 
**Power** | Pointer to **float64** | Poe Power | [optional] 
**StackId** | Pointer to **string** | Stack ID to which the port belongs | [optional] 
**StackName** | Pointer to **string** | Stack Name to which the port belongs | [optional] 
**StackPortsGroupIndex** | Pointer to **int32** | Number of the stacking port aggregation group to join | [optional] 
**StandardPort** | Pointer to [**OswStandPortVO**](OswStandPortVO.md) |  | [optional] 
**SupportPoe** | Pointer to **bool** | Indicates whether PoE is supported | [optional] 
**SwitchMac** | Pointer to **string** | Switch Mac to which the port belongs | [optional] 
**SwitchName** | Pointer to **string** | Switch Name to which the port belongs | [optional] 
**SwitchStatusCategory** | Pointer to **int32** | Category of switch status, switchStatusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated.Only &#39;Disconnected&#39; and &#39;Connected&#39; and &#39;Heartbeat&#39; switches ports will be displayed. | [optional] 
**SwitchSupportPoe** | Pointer to **int32** | Indicates whether the switch supports PoE, switchSupportPoe should be a value as follows: 0:Not Support;1:Support. | [optional] 
**SwitchType** | Pointer to **int32** | Switch Type to which the port belongs | [optional] 
**TagIds** | Pointer to **[]string** | Port label ID List | [optional] 
**TagName** | Pointer to **string** | Port label Name | [optional] 
**Type** | Pointer to **int32** | Type should be a value as follows: 1: Copper; 2: Combo; 3: SFP | [optional] 
**Unit** | Pointer to **int32** | Stack Unit to which the port belongs | [optional] 
**Voltage** | Pointer to **float64** | Poe Voltage | [optional] 

## Methods

### NewOswPortsSettingPoeVO

`func NewOswPortsSettingPoeVO() *OswPortsSettingPoeVO`

NewOswPortsSettingPoeVO instantiates a new OswPortsSettingPoeVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPortsSettingPoeVOWithDefaults

`func NewOswPortsSettingPoeVOWithDefaults() *OswPortsSettingPoeVO`

NewOswPortsSettingPoeVOWithDefaults instantiates a new OswPortsSettingPoeVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OswPortsSettingPoeVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OswPortsSettingPoeVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OswPortsSettingPoeVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OswPortsSettingPoeVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetConfigMlagDad

`func (o *OswPortsSettingPoeVO) GetConfigMlagDad() bool`

GetConfigMlagDad returns the ConfigMlagDad field if non-nil, zero value otherwise.

### GetConfigMlagDadOk

`func (o *OswPortsSettingPoeVO) GetConfigMlagDadOk() (*bool, bool)`

GetConfigMlagDadOk returns a tuple with the ConfigMlagDad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagDad

`func (o *OswPortsSettingPoeVO) SetConfigMlagDad(v bool)`

SetConfigMlagDad sets ConfigMlagDad field to given value.

### HasConfigMlagDad

`func (o *OswPortsSettingPoeVO) HasConfigMlagDad() bool`

HasConfigMlagDad returns a boolean if a field has been set.

### GetConfigMlagPeerLink

`func (o *OswPortsSettingPoeVO) GetConfigMlagPeerLink() bool`

GetConfigMlagPeerLink returns the ConfigMlagPeerLink field if non-nil, zero value otherwise.

### GetConfigMlagPeerLinkOk

`func (o *OswPortsSettingPoeVO) GetConfigMlagPeerLinkOk() (*bool, bool)`

GetConfigMlagPeerLinkOk returns a tuple with the ConfigMlagPeerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagPeerLink

`func (o *OswPortsSettingPoeVO) SetConfigMlagPeerLink(v bool)`

SetConfigMlagPeerLink sets ConfigMlagPeerLink field to given value.

### HasConfigMlagPeerLink

`func (o *OswPortsSettingPoeVO) HasConfigMlagPeerLink() bool`

HasConfigMlagPeerLink returns a boolean if a field has been set.

### GetConfigStack

`func (o *OswPortsSettingPoeVO) GetConfigStack() bool`

GetConfigStack returns the ConfigStack field if non-nil, zero value otherwise.

### GetConfigStackOk

`func (o *OswPortsSettingPoeVO) GetConfigStackOk() (*bool, bool)`

GetConfigStackOk returns a tuple with the ConfigStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStack

`func (o *OswPortsSettingPoeVO) SetConfigStack(v bool)`

SetConfigStack sets ConfigStack field to given value.

### HasConfigStack

`func (o *OswPortsSettingPoeVO) HasConfigStack() bool`

HasConfigStack returns a boolean if a field has been set.

### GetConnectedStatus

`func (o *OswPortsSettingPoeVO) GetConnectedStatus() int32`

GetConnectedStatus returns the ConnectedStatus field if non-nil, zero value otherwise.

### GetConnectedStatusOk

`func (o *OswPortsSettingPoeVO) GetConnectedStatusOk() (*int32, bool)`

GetConnectedStatusOk returns a tuple with the ConnectedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedStatus

`func (o *OswPortsSettingPoeVO) SetConnectedStatus(v int32)`

SetConnectedStatus sets ConnectedStatus field to given value.

### HasConnectedStatus

`func (o *OswPortsSettingPoeVO) HasConnectedStatus() bool`

HasConnectedStatus returns a boolean if a field has been set.

### GetCurrent

`func (o *OswPortsSettingPoeVO) GetCurrent() float64`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *OswPortsSettingPoeVO) GetCurrentOk() (*float64, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *OswPortsSettingPoeVO) SetCurrent(v float64)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *OswPortsSettingPoeVO) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetDisable

`func (o *OswPortsSettingPoeVO) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *OswPortsSettingPoeVO) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *OswPortsSettingPoeVO) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *OswPortsSettingPoeVO) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDuplex

`func (o *OswPortsSettingPoeVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswPortsSettingPoeVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswPortsSettingPoeVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswPortsSettingPoeVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLagSetting

`func (o *OswPortsSettingPoeVO) GetLagSetting() OswStackLagVO`

GetLagSetting returns the LagSetting field if non-nil, zero value otherwise.

### GetLagSettingOk

`func (o *OswPortsSettingPoeVO) GetLagSettingOk() (*OswStackLagVO, bool)`

GetLagSettingOk returns a tuple with the LagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagSetting

`func (o *OswPortsSettingPoeVO) SetLagSetting(v OswStackLagVO)`

SetLagSetting sets LagSetting field to given value.

### HasLagSetting

`func (o *OswPortsSettingPoeVO) HasLagSetting() bool`

HasLagSetting returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswPortsSettingPoeVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswPortsSettingPoeVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswPortsSettingPoeVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswPortsSettingPoeVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetMadUsed

`func (o *OswPortsSettingPoeVO) GetMadUsed() bool`

GetMadUsed returns the MadUsed field if non-nil, zero value otherwise.

### GetMadUsedOk

`func (o *OswPortsSettingPoeVO) GetMadUsedOk() (*bool, bool)`

GetMadUsedOk returns a tuple with the MadUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadUsed

`func (o *OswPortsSettingPoeVO) SetMadUsed(v bool)`

SetMadUsed sets MadUsed field to given value.

### HasMadUsed

`func (o *OswPortsSettingPoeVO) HasMadUsed() bool`

HasMadUsed returns a boolean if a field has been set.

### GetOperation

`func (o *OswPortsSettingPoeVO) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *OswPortsSettingPoeVO) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *OswPortsSettingPoeVO) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *OswPortsSettingPoeVO) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPdClass

`func (o *OswPortsSettingPoeVO) GetPdClass() string`

GetPdClass returns the PdClass field if non-nil, zero value otherwise.

### GetPdClassOk

`func (o *OswPortsSettingPoeVO) GetPdClassOk() (*string, bool)`

GetPdClassOk returns a tuple with the PdClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdClass

`func (o *OswPortsSettingPoeVO) SetPdClass(v string)`

SetPdClass sets PdClass field to given value.

### HasPdClass

`func (o *OswPortsSettingPoeVO) HasPdClass() bool`

HasPdClass returns a boolean if a field has been set.

### GetPoe

`func (o *OswPortsSettingPoeVO) GetPoe() int32`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *OswPortsSettingPoeVO) GetPoeOk() (*int32, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *OswPortsSettingPoeVO) SetPoe(v int32)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *OswPortsSettingPoeVO) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetPoeDisplayType

`func (o *OswPortsSettingPoeVO) GetPoeDisplayType() int32`

GetPoeDisplayType returns the PoeDisplayType field if non-nil, zero value otherwise.

### GetPoeDisplayTypeOk

`func (o *OswPortsSettingPoeVO) GetPoeDisplayTypeOk() (*int32, bool)`

GetPoeDisplayTypeOk returns a tuple with the PoeDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeDisplayType

`func (o *OswPortsSettingPoeVO) SetPoeDisplayType(v int32)`

SetPoeDisplayType sets PoeDisplayType field to given value.

### HasPoeDisplayType

`func (o *OswPortsSettingPoeVO) HasPoeDisplayType() bool`

HasPoeDisplayType returns a boolean if a field has been set.

### GetPoeStatus

`func (o *OswPortsSettingPoeVO) GetPoeStatus() float64`

GetPoeStatus returns the PoeStatus field if non-nil, zero value otherwise.

### GetPoeStatusOk

`func (o *OswPortsSettingPoeVO) GetPoeStatusOk() (*float64, bool)`

GetPoeStatusOk returns a tuple with the PoeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeStatus

`func (o *OswPortsSettingPoeVO) SetPoeStatus(v float64)`

SetPoeStatus sets PoeStatus field to given value.

### HasPoeStatus

`func (o *OswPortsSettingPoeVO) HasPoeStatus() bool`

HasPoeStatus returns a boolean if a field has been set.

### GetPort

`func (o *OswPortsSettingPoeVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswPortsSettingPoeVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswPortsSettingPoeVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswPortsSettingPoeVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortName

`func (o *OswPortsSettingPoeVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *OswPortsSettingPoeVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *OswPortsSettingPoeVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *OswPortsSettingPoeVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPortStatus

`func (o *OswPortsSettingPoeVO) GetPortStatus() OswPortStatusVO`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *OswPortsSettingPoeVO) GetPortStatusOk() (*OswPortStatusVO, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *OswPortsSettingPoeVO) SetPortStatus(v OswPortStatusVO)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *OswPortsSettingPoeVO) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetPower

`func (o *OswPortsSettingPoeVO) GetPower() float64`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *OswPortsSettingPoeVO) GetPowerOk() (*float64, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *OswPortsSettingPoeVO) SetPower(v float64)`

SetPower sets Power field to given value.

### HasPower

`func (o *OswPortsSettingPoeVO) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetStackId

`func (o *OswPortsSettingPoeVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswPortsSettingPoeVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswPortsSettingPoeVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswPortsSettingPoeVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *OswPortsSettingPoeVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *OswPortsSettingPoeVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *OswPortsSettingPoeVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *OswPortsSettingPoeVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStackPortsGroupIndex

`func (o *OswPortsSettingPoeVO) GetStackPortsGroupIndex() int32`

GetStackPortsGroupIndex returns the StackPortsGroupIndex field if non-nil, zero value otherwise.

### GetStackPortsGroupIndexOk

`func (o *OswPortsSettingPoeVO) GetStackPortsGroupIndexOk() (*int32, bool)`

GetStackPortsGroupIndexOk returns a tuple with the StackPortsGroupIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPortsGroupIndex

`func (o *OswPortsSettingPoeVO) SetStackPortsGroupIndex(v int32)`

SetStackPortsGroupIndex sets StackPortsGroupIndex field to given value.

### HasStackPortsGroupIndex

`func (o *OswPortsSettingPoeVO) HasStackPortsGroupIndex() bool`

HasStackPortsGroupIndex returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswPortsSettingPoeVO) GetStandardPort() OswStandPortVO`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswPortsSettingPoeVO) GetStandardPortOk() (*OswStandPortVO, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswPortsSettingPoeVO) SetStandardPort(v OswStandPortVO)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswPortsSettingPoeVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetSupportPoe

`func (o *OswPortsSettingPoeVO) GetSupportPoe() bool`

GetSupportPoe returns the SupportPoe field if non-nil, zero value otherwise.

### GetSupportPoeOk

`func (o *OswPortsSettingPoeVO) GetSupportPoeOk() (*bool, bool)`

GetSupportPoeOk returns a tuple with the SupportPoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPoe

`func (o *OswPortsSettingPoeVO) SetSupportPoe(v bool)`

SetSupportPoe sets SupportPoe field to given value.

### HasSupportPoe

`func (o *OswPortsSettingPoeVO) HasSupportPoe() bool`

HasSupportPoe returns a boolean if a field has been set.

### GetSwitchMac

`func (o *OswPortsSettingPoeVO) GetSwitchMac() string`

GetSwitchMac returns the SwitchMac field if non-nil, zero value otherwise.

### GetSwitchMacOk

`func (o *OswPortsSettingPoeVO) GetSwitchMacOk() (*string, bool)`

GetSwitchMacOk returns a tuple with the SwitchMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMac

`func (o *OswPortsSettingPoeVO) SetSwitchMac(v string)`

SetSwitchMac sets SwitchMac field to given value.

### HasSwitchMac

`func (o *OswPortsSettingPoeVO) HasSwitchMac() bool`

HasSwitchMac returns a boolean if a field has been set.

### GetSwitchName

`func (o *OswPortsSettingPoeVO) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *OswPortsSettingPoeVO) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *OswPortsSettingPoeVO) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *OswPortsSettingPoeVO) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetSwitchStatusCategory

`func (o *OswPortsSettingPoeVO) GetSwitchStatusCategory() int32`

GetSwitchStatusCategory returns the SwitchStatusCategory field if non-nil, zero value otherwise.

### GetSwitchStatusCategoryOk

`func (o *OswPortsSettingPoeVO) GetSwitchStatusCategoryOk() (*int32, bool)`

GetSwitchStatusCategoryOk returns a tuple with the SwitchStatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchStatusCategory

`func (o *OswPortsSettingPoeVO) SetSwitchStatusCategory(v int32)`

SetSwitchStatusCategory sets SwitchStatusCategory field to given value.

### HasSwitchStatusCategory

`func (o *OswPortsSettingPoeVO) HasSwitchStatusCategory() bool`

HasSwitchStatusCategory returns a boolean if a field has been set.

### GetSwitchSupportPoe

`func (o *OswPortsSettingPoeVO) GetSwitchSupportPoe() int32`

GetSwitchSupportPoe returns the SwitchSupportPoe field if non-nil, zero value otherwise.

### GetSwitchSupportPoeOk

`func (o *OswPortsSettingPoeVO) GetSwitchSupportPoeOk() (*int32, bool)`

GetSwitchSupportPoeOk returns a tuple with the SwitchSupportPoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchSupportPoe

`func (o *OswPortsSettingPoeVO) SetSwitchSupportPoe(v int32)`

SetSwitchSupportPoe sets SwitchSupportPoe field to given value.

### HasSwitchSupportPoe

`func (o *OswPortsSettingPoeVO) HasSwitchSupportPoe() bool`

HasSwitchSupportPoe returns a boolean if a field has been set.

### GetSwitchType

`func (o *OswPortsSettingPoeVO) GetSwitchType() int32`

GetSwitchType returns the SwitchType field if non-nil, zero value otherwise.

### GetSwitchTypeOk

`func (o *OswPortsSettingPoeVO) GetSwitchTypeOk() (*int32, bool)`

GetSwitchTypeOk returns a tuple with the SwitchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchType

`func (o *OswPortsSettingPoeVO) SetSwitchType(v int32)`

SetSwitchType sets SwitchType field to given value.

### HasSwitchType

`func (o *OswPortsSettingPoeVO) HasSwitchType() bool`

HasSwitchType returns a boolean if a field has been set.

### GetTagIds

`func (o *OswPortsSettingPoeVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OswPortsSettingPoeVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OswPortsSettingPoeVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OswPortsSettingPoeVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagName

`func (o *OswPortsSettingPoeVO) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *OswPortsSettingPoeVO) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *OswPortsSettingPoeVO) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *OswPortsSettingPoeVO) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetType

`func (o *OswPortsSettingPoeVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswPortsSettingPoeVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswPortsSettingPoeVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OswPortsSettingPoeVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *OswPortsSettingPoeVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswPortsSettingPoeVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswPortsSettingPoeVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswPortsSettingPoeVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetVoltage

`func (o *OswPortsSettingPoeVO) GetVoltage() float64`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *OswPortsSettingPoeVO) GetVoltageOk() (*float64, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *OswPortsSettingPoeVO) SetVoltage(v float64)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *OswPortsSettingPoeVO) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


