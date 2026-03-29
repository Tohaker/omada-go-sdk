# OswPortsSettingCountersVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site to which the port belongs | [optional] 
**ConfigMlagDad** | Pointer to **bool** | Indicates whether the current port is configured as M-LAG dad Link | [optional] 
**ConfigMlagPeerLink** | Pointer to **bool** | Indicates whether the current port is configured as M-LAG peer Link | [optional] 
**ConfigStack** | Pointer to **bool** | Indicates whether the current port is configured as a stack port (joined a stack aggregation group) | [optional] 
**ConnectedStatus** | Pointer to **int32** | Port connected status should be a value as follows: 0: Connected; 1: Disconnected; 2: Disable | [optional] 
**Disable** | Pointer to **bool** | Indicates whether to disable the port | [optional] 
**Duplex** | Pointer to **int32** | Duplex should be a value as follows: 0: Auto; 1: Half; 2: Full | [optional] 
**LagSetting** | Pointer to [**OswStackLagVO**](OswStackLagVO.md) |  | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed should be a value as follows: 0: auto; 1: 10M; 2: 100M; 3: 1000M; 4: 10G | [optional] 
**MadUsed** | Pointer to **bool** | Mad Used | [optional] 
**Operation** | Pointer to **string** | Operation should be a value as follows: SWITCHING; MIRRORING; AGGREGATING | [optional] 
**Poe** | Pointer to **int32** | PoE switch should be a value as follows: 0: Off; 1: 802.3at/af | [optional] 
**PoeDisplayType** | Pointer to **int32** | PoeDisplayType should be a value as follows: -1: Not Support POE; 0: Support POE; 1: POE(4W); 2: POE(7W); 3: POE(15.4W); 4: POE+(30W); 5: POE++(45W); 6: POE++(60W); 7: POE++(75W); 8: POE++(90W); 9: POE++(100W). | [optional] 
**Port** | Pointer to **int32** | Port | [optional] 
**PortName** | Pointer to **string** | Port Name | [optional] 
**PortStatus** | Pointer to [**OswPortStatusVO**](OswPortStatusVO.md) |  | [optional] 
**RxBroadcast** | Pointer to **int64** | Total RX Broadcast Packets | [optional] 
**RxBytes** | Pointer to **int64** | Total Rx Bytes | [optional] 
**RxErrors** | Pointer to **int64** | Total RX Errors | [optional] 
**RxFrames** | Pointer to **int64** | Total RX Packets | [optional] 
**RxMulticast** | Pointer to **int64** | Total RX Multicast Packets | [optional] 
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
**TxBroadcast** | Pointer to **int64** | Total TX Broadcast Packets | [optional] 
**TxBytes** | Pointer to **int64** | Total Tx Bytes | [optional] 
**TxErrors** | Pointer to **int64** | Total TX Errors | [optional] 
**TxFrames** | Pointer to **int64** | Total TX Packets | [optional] 
**TxMulticast** | Pointer to **int64** | Total TX Multicast Packets | [optional] 
**Type** | Pointer to **int32** | Type should be a value as follows: 1: Copper; 2: Combo; 3: SFP | [optional] 
**Unit** | Pointer to **int32** | Stack Unit to which the port belongs | [optional] 

## Methods

### NewOswPortsSettingCountersVO

`func NewOswPortsSettingCountersVO() *OswPortsSettingCountersVO`

NewOswPortsSettingCountersVO instantiates a new OswPortsSettingCountersVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPortsSettingCountersVOWithDefaults

`func NewOswPortsSettingCountersVOWithDefaults() *OswPortsSettingCountersVO`

NewOswPortsSettingCountersVOWithDefaults instantiates a new OswPortsSettingCountersVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OswPortsSettingCountersVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OswPortsSettingCountersVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OswPortsSettingCountersVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OswPortsSettingCountersVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetConfigMlagDad

`func (o *OswPortsSettingCountersVO) GetConfigMlagDad() bool`

GetConfigMlagDad returns the ConfigMlagDad field if non-nil, zero value otherwise.

### GetConfigMlagDadOk

`func (o *OswPortsSettingCountersVO) GetConfigMlagDadOk() (*bool, bool)`

GetConfigMlagDadOk returns a tuple with the ConfigMlagDad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagDad

`func (o *OswPortsSettingCountersVO) SetConfigMlagDad(v bool)`

SetConfigMlagDad sets ConfigMlagDad field to given value.

### HasConfigMlagDad

`func (o *OswPortsSettingCountersVO) HasConfigMlagDad() bool`

HasConfigMlagDad returns a boolean if a field has been set.

### GetConfigMlagPeerLink

`func (o *OswPortsSettingCountersVO) GetConfigMlagPeerLink() bool`

GetConfigMlagPeerLink returns the ConfigMlagPeerLink field if non-nil, zero value otherwise.

### GetConfigMlagPeerLinkOk

`func (o *OswPortsSettingCountersVO) GetConfigMlagPeerLinkOk() (*bool, bool)`

GetConfigMlagPeerLinkOk returns a tuple with the ConfigMlagPeerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagPeerLink

`func (o *OswPortsSettingCountersVO) SetConfigMlagPeerLink(v bool)`

SetConfigMlagPeerLink sets ConfigMlagPeerLink field to given value.

### HasConfigMlagPeerLink

`func (o *OswPortsSettingCountersVO) HasConfigMlagPeerLink() bool`

HasConfigMlagPeerLink returns a boolean if a field has been set.

### GetConfigStack

`func (o *OswPortsSettingCountersVO) GetConfigStack() bool`

GetConfigStack returns the ConfigStack field if non-nil, zero value otherwise.

### GetConfigStackOk

`func (o *OswPortsSettingCountersVO) GetConfigStackOk() (*bool, bool)`

GetConfigStackOk returns a tuple with the ConfigStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStack

`func (o *OswPortsSettingCountersVO) SetConfigStack(v bool)`

SetConfigStack sets ConfigStack field to given value.

### HasConfigStack

`func (o *OswPortsSettingCountersVO) HasConfigStack() bool`

HasConfigStack returns a boolean if a field has been set.

### GetConnectedStatus

`func (o *OswPortsSettingCountersVO) GetConnectedStatus() int32`

GetConnectedStatus returns the ConnectedStatus field if non-nil, zero value otherwise.

### GetConnectedStatusOk

`func (o *OswPortsSettingCountersVO) GetConnectedStatusOk() (*int32, bool)`

GetConnectedStatusOk returns a tuple with the ConnectedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedStatus

`func (o *OswPortsSettingCountersVO) SetConnectedStatus(v int32)`

SetConnectedStatus sets ConnectedStatus field to given value.

### HasConnectedStatus

`func (o *OswPortsSettingCountersVO) HasConnectedStatus() bool`

HasConnectedStatus returns a boolean if a field has been set.

### GetDisable

`func (o *OswPortsSettingCountersVO) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *OswPortsSettingCountersVO) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *OswPortsSettingCountersVO) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *OswPortsSettingCountersVO) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDuplex

`func (o *OswPortsSettingCountersVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswPortsSettingCountersVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswPortsSettingCountersVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswPortsSettingCountersVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLagSetting

`func (o *OswPortsSettingCountersVO) GetLagSetting() OswStackLagVO`

GetLagSetting returns the LagSetting field if non-nil, zero value otherwise.

### GetLagSettingOk

`func (o *OswPortsSettingCountersVO) GetLagSettingOk() (*OswStackLagVO, bool)`

GetLagSettingOk returns a tuple with the LagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagSetting

`func (o *OswPortsSettingCountersVO) SetLagSetting(v OswStackLagVO)`

SetLagSetting sets LagSetting field to given value.

### HasLagSetting

`func (o *OswPortsSettingCountersVO) HasLagSetting() bool`

HasLagSetting returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswPortsSettingCountersVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswPortsSettingCountersVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswPortsSettingCountersVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswPortsSettingCountersVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetMadUsed

`func (o *OswPortsSettingCountersVO) GetMadUsed() bool`

GetMadUsed returns the MadUsed field if non-nil, zero value otherwise.

### GetMadUsedOk

`func (o *OswPortsSettingCountersVO) GetMadUsedOk() (*bool, bool)`

GetMadUsedOk returns a tuple with the MadUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadUsed

`func (o *OswPortsSettingCountersVO) SetMadUsed(v bool)`

SetMadUsed sets MadUsed field to given value.

### HasMadUsed

`func (o *OswPortsSettingCountersVO) HasMadUsed() bool`

HasMadUsed returns a boolean if a field has been set.

### GetOperation

`func (o *OswPortsSettingCountersVO) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *OswPortsSettingCountersVO) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *OswPortsSettingCountersVO) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *OswPortsSettingCountersVO) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPoe

`func (o *OswPortsSettingCountersVO) GetPoe() int32`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *OswPortsSettingCountersVO) GetPoeOk() (*int32, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *OswPortsSettingCountersVO) SetPoe(v int32)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *OswPortsSettingCountersVO) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetPoeDisplayType

`func (o *OswPortsSettingCountersVO) GetPoeDisplayType() int32`

GetPoeDisplayType returns the PoeDisplayType field if non-nil, zero value otherwise.

### GetPoeDisplayTypeOk

`func (o *OswPortsSettingCountersVO) GetPoeDisplayTypeOk() (*int32, bool)`

GetPoeDisplayTypeOk returns a tuple with the PoeDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeDisplayType

`func (o *OswPortsSettingCountersVO) SetPoeDisplayType(v int32)`

SetPoeDisplayType sets PoeDisplayType field to given value.

### HasPoeDisplayType

`func (o *OswPortsSettingCountersVO) HasPoeDisplayType() bool`

HasPoeDisplayType returns a boolean if a field has been set.

### GetPort

`func (o *OswPortsSettingCountersVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswPortsSettingCountersVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswPortsSettingCountersVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswPortsSettingCountersVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortName

`func (o *OswPortsSettingCountersVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *OswPortsSettingCountersVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *OswPortsSettingCountersVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *OswPortsSettingCountersVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPortStatus

`func (o *OswPortsSettingCountersVO) GetPortStatus() OswPortStatusVO`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *OswPortsSettingCountersVO) GetPortStatusOk() (*OswPortStatusVO, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *OswPortsSettingCountersVO) SetPortStatus(v OswPortStatusVO)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *OswPortsSettingCountersVO) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetRxBroadcast

`func (o *OswPortsSettingCountersVO) GetRxBroadcast() int64`

GetRxBroadcast returns the RxBroadcast field if non-nil, zero value otherwise.

### GetRxBroadcastOk

`func (o *OswPortsSettingCountersVO) GetRxBroadcastOk() (*int64, bool)`

GetRxBroadcastOk returns a tuple with the RxBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBroadcast

`func (o *OswPortsSettingCountersVO) SetRxBroadcast(v int64)`

SetRxBroadcast sets RxBroadcast field to given value.

### HasRxBroadcast

`func (o *OswPortsSettingCountersVO) HasRxBroadcast() bool`

HasRxBroadcast returns a boolean if a field has been set.

### GetRxBytes

`func (o *OswPortsSettingCountersVO) GetRxBytes() int64`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *OswPortsSettingCountersVO) GetRxBytesOk() (*int64, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *OswPortsSettingCountersVO) SetRxBytes(v int64)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *OswPortsSettingCountersVO) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### GetRxErrors

`func (o *OswPortsSettingCountersVO) GetRxErrors() int64`

GetRxErrors returns the RxErrors field if non-nil, zero value otherwise.

### GetRxErrorsOk

`func (o *OswPortsSettingCountersVO) GetRxErrorsOk() (*int64, bool)`

GetRxErrorsOk returns a tuple with the RxErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrors

`func (o *OswPortsSettingCountersVO) SetRxErrors(v int64)`

SetRxErrors sets RxErrors field to given value.

### HasRxErrors

`func (o *OswPortsSettingCountersVO) HasRxErrors() bool`

HasRxErrors returns a boolean if a field has been set.

### GetRxFrames

`func (o *OswPortsSettingCountersVO) GetRxFrames() int64`

GetRxFrames returns the RxFrames field if non-nil, zero value otherwise.

### GetRxFramesOk

`func (o *OswPortsSettingCountersVO) GetRxFramesOk() (*int64, bool)`

GetRxFramesOk returns a tuple with the RxFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxFrames

`func (o *OswPortsSettingCountersVO) SetRxFrames(v int64)`

SetRxFrames sets RxFrames field to given value.

### HasRxFrames

`func (o *OswPortsSettingCountersVO) HasRxFrames() bool`

HasRxFrames returns a boolean if a field has been set.

### GetRxMulticast

`func (o *OswPortsSettingCountersVO) GetRxMulticast() int64`

GetRxMulticast returns the RxMulticast field if non-nil, zero value otherwise.

### GetRxMulticastOk

`func (o *OswPortsSettingCountersVO) GetRxMulticastOk() (*int64, bool)`

GetRxMulticastOk returns a tuple with the RxMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMulticast

`func (o *OswPortsSettingCountersVO) SetRxMulticast(v int64)`

SetRxMulticast sets RxMulticast field to given value.

### HasRxMulticast

`func (o *OswPortsSettingCountersVO) HasRxMulticast() bool`

HasRxMulticast returns a boolean if a field has been set.

### GetStackId

`func (o *OswPortsSettingCountersVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswPortsSettingCountersVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswPortsSettingCountersVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswPortsSettingCountersVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *OswPortsSettingCountersVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *OswPortsSettingCountersVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *OswPortsSettingCountersVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *OswPortsSettingCountersVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStackPortsGroupIndex

`func (o *OswPortsSettingCountersVO) GetStackPortsGroupIndex() int32`

GetStackPortsGroupIndex returns the StackPortsGroupIndex field if non-nil, zero value otherwise.

### GetStackPortsGroupIndexOk

`func (o *OswPortsSettingCountersVO) GetStackPortsGroupIndexOk() (*int32, bool)`

GetStackPortsGroupIndexOk returns a tuple with the StackPortsGroupIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPortsGroupIndex

`func (o *OswPortsSettingCountersVO) SetStackPortsGroupIndex(v int32)`

SetStackPortsGroupIndex sets StackPortsGroupIndex field to given value.

### HasStackPortsGroupIndex

`func (o *OswPortsSettingCountersVO) HasStackPortsGroupIndex() bool`

HasStackPortsGroupIndex returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswPortsSettingCountersVO) GetStandardPort() OswStandPortVO`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswPortsSettingCountersVO) GetStandardPortOk() (*OswStandPortVO, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswPortsSettingCountersVO) SetStandardPort(v OswStandPortVO)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswPortsSettingCountersVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetSupportPoe

`func (o *OswPortsSettingCountersVO) GetSupportPoe() bool`

GetSupportPoe returns the SupportPoe field if non-nil, zero value otherwise.

### GetSupportPoeOk

`func (o *OswPortsSettingCountersVO) GetSupportPoeOk() (*bool, bool)`

GetSupportPoeOk returns a tuple with the SupportPoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPoe

`func (o *OswPortsSettingCountersVO) SetSupportPoe(v bool)`

SetSupportPoe sets SupportPoe field to given value.

### HasSupportPoe

`func (o *OswPortsSettingCountersVO) HasSupportPoe() bool`

HasSupportPoe returns a boolean if a field has been set.

### GetSwitchMac

`func (o *OswPortsSettingCountersVO) GetSwitchMac() string`

GetSwitchMac returns the SwitchMac field if non-nil, zero value otherwise.

### GetSwitchMacOk

`func (o *OswPortsSettingCountersVO) GetSwitchMacOk() (*string, bool)`

GetSwitchMacOk returns a tuple with the SwitchMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMac

`func (o *OswPortsSettingCountersVO) SetSwitchMac(v string)`

SetSwitchMac sets SwitchMac field to given value.

### HasSwitchMac

`func (o *OswPortsSettingCountersVO) HasSwitchMac() bool`

HasSwitchMac returns a boolean if a field has been set.

### GetSwitchName

`func (o *OswPortsSettingCountersVO) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *OswPortsSettingCountersVO) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *OswPortsSettingCountersVO) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *OswPortsSettingCountersVO) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetSwitchStatusCategory

`func (o *OswPortsSettingCountersVO) GetSwitchStatusCategory() int32`

GetSwitchStatusCategory returns the SwitchStatusCategory field if non-nil, zero value otherwise.

### GetSwitchStatusCategoryOk

`func (o *OswPortsSettingCountersVO) GetSwitchStatusCategoryOk() (*int32, bool)`

GetSwitchStatusCategoryOk returns a tuple with the SwitchStatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchStatusCategory

`func (o *OswPortsSettingCountersVO) SetSwitchStatusCategory(v int32)`

SetSwitchStatusCategory sets SwitchStatusCategory field to given value.

### HasSwitchStatusCategory

`func (o *OswPortsSettingCountersVO) HasSwitchStatusCategory() bool`

HasSwitchStatusCategory returns a boolean if a field has been set.

### GetSwitchSupportPoe

`func (o *OswPortsSettingCountersVO) GetSwitchSupportPoe() int32`

GetSwitchSupportPoe returns the SwitchSupportPoe field if non-nil, zero value otherwise.

### GetSwitchSupportPoeOk

`func (o *OswPortsSettingCountersVO) GetSwitchSupportPoeOk() (*int32, bool)`

GetSwitchSupportPoeOk returns a tuple with the SwitchSupportPoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchSupportPoe

`func (o *OswPortsSettingCountersVO) SetSwitchSupportPoe(v int32)`

SetSwitchSupportPoe sets SwitchSupportPoe field to given value.

### HasSwitchSupportPoe

`func (o *OswPortsSettingCountersVO) HasSwitchSupportPoe() bool`

HasSwitchSupportPoe returns a boolean if a field has been set.

### GetSwitchType

`func (o *OswPortsSettingCountersVO) GetSwitchType() int32`

GetSwitchType returns the SwitchType field if non-nil, zero value otherwise.

### GetSwitchTypeOk

`func (o *OswPortsSettingCountersVO) GetSwitchTypeOk() (*int32, bool)`

GetSwitchTypeOk returns a tuple with the SwitchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchType

`func (o *OswPortsSettingCountersVO) SetSwitchType(v int32)`

SetSwitchType sets SwitchType field to given value.

### HasSwitchType

`func (o *OswPortsSettingCountersVO) HasSwitchType() bool`

HasSwitchType returns a boolean if a field has been set.

### GetTagIds

`func (o *OswPortsSettingCountersVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OswPortsSettingCountersVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OswPortsSettingCountersVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OswPortsSettingCountersVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagName

`func (o *OswPortsSettingCountersVO) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *OswPortsSettingCountersVO) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *OswPortsSettingCountersVO) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *OswPortsSettingCountersVO) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetTxBroadcast

`func (o *OswPortsSettingCountersVO) GetTxBroadcast() int64`

GetTxBroadcast returns the TxBroadcast field if non-nil, zero value otherwise.

### GetTxBroadcastOk

`func (o *OswPortsSettingCountersVO) GetTxBroadcastOk() (*int64, bool)`

GetTxBroadcastOk returns a tuple with the TxBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBroadcast

`func (o *OswPortsSettingCountersVO) SetTxBroadcast(v int64)`

SetTxBroadcast sets TxBroadcast field to given value.

### HasTxBroadcast

`func (o *OswPortsSettingCountersVO) HasTxBroadcast() bool`

HasTxBroadcast returns a boolean if a field has been set.

### GetTxBytes

`func (o *OswPortsSettingCountersVO) GetTxBytes() int64`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *OswPortsSettingCountersVO) GetTxBytesOk() (*int64, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *OswPortsSettingCountersVO) SetTxBytes(v int64)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *OswPortsSettingCountersVO) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### GetTxErrors

`func (o *OswPortsSettingCountersVO) GetTxErrors() int64`

GetTxErrors returns the TxErrors field if non-nil, zero value otherwise.

### GetTxErrorsOk

`func (o *OswPortsSettingCountersVO) GetTxErrorsOk() (*int64, bool)`

GetTxErrorsOk returns a tuple with the TxErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrors

`func (o *OswPortsSettingCountersVO) SetTxErrors(v int64)`

SetTxErrors sets TxErrors field to given value.

### HasTxErrors

`func (o *OswPortsSettingCountersVO) HasTxErrors() bool`

HasTxErrors returns a boolean if a field has been set.

### GetTxFrames

`func (o *OswPortsSettingCountersVO) GetTxFrames() int64`

GetTxFrames returns the TxFrames field if non-nil, zero value otherwise.

### GetTxFramesOk

`func (o *OswPortsSettingCountersVO) GetTxFramesOk() (*int64, bool)`

GetTxFramesOk returns a tuple with the TxFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxFrames

`func (o *OswPortsSettingCountersVO) SetTxFrames(v int64)`

SetTxFrames sets TxFrames field to given value.

### HasTxFrames

`func (o *OswPortsSettingCountersVO) HasTxFrames() bool`

HasTxFrames returns a boolean if a field has been set.

### GetTxMulticast

`func (o *OswPortsSettingCountersVO) GetTxMulticast() int64`

GetTxMulticast returns the TxMulticast field if non-nil, zero value otherwise.

### GetTxMulticastOk

`func (o *OswPortsSettingCountersVO) GetTxMulticastOk() (*int64, bool)`

GetTxMulticastOk returns a tuple with the TxMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMulticast

`func (o *OswPortsSettingCountersVO) SetTxMulticast(v int64)`

SetTxMulticast sets TxMulticast field to given value.

### HasTxMulticast

`func (o *OswPortsSettingCountersVO) HasTxMulticast() bool`

HasTxMulticast returns a boolean if a field has been set.

### GetType

`func (o *OswPortsSettingCountersVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswPortsSettingCountersVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswPortsSettingCountersVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OswPortsSettingCountersVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *OswPortsSettingCountersVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswPortsSettingCountersVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswPortsSettingCountersVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswPortsSettingCountersVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


