# OswPortsSettingOverviewVO

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
**NativeBridgeVlan** | Pointer to **int32** | Native Network Bridge Vlan. | [optional] 
**NativeNetworkId** | Pointer to **string** | Native Network ID, Native Network cannot be selected from Tagged Networks or Untagged Networks. | [optional] 
**NetworkMode** | Pointer to **int32** | Network Mode should be a value as follows: 0: Trunk, 1: Access | [optional] 
**NetworkTagsSetting** | Pointer to **int32** | Network Tags Setting should be a value as follows: 0: Allow All; 1: Block All; 2: Custom | [optional] 
**Operation** | Pointer to **string** | Operation should be a value as follows: SWITCHING; MIRRORING; AGGREGATING | [optional] 
**Poe** | Pointer to **int32** | PoE switch should be a value as follows: 0: Off; 1: 802.3at/af | [optional] 
**PoeDisplayType** | Pointer to **int32** | PoeDisplayType should be a value as follows: -1: Not Support POE; 0: Support POE; 1: POE(4W); 2: POE(7W); 3: POE(15.4W); 4: POE+(30W); 5: POE++(45W); 6: POE++(60W); 7: POE++(75W); 8: POE++(90W); 9: POE++(100W). | [optional] 
**Port** | Pointer to **int32** | Port | [optional] 
**PortName** | Pointer to **string** | Port Name | [optional] 
**PortStatus** | Pointer to [**OswPortStatusVO**](OswPortStatusVO.md) |  | [optional] 
**ProfileId** | Pointer to **string** | Lan Profile ID | [optional] 
**ProfileName** | Pointer to **string** | Lan Profile Name | [optional] 
**ProfileOverrideEnable** | Pointer to **bool** | Indicates whether to enable Profile Override | [optional] 
**ProfileType** | Pointer to **int32** | Lan Profile Type | [optional] 
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

## Methods

### NewOswPortsSettingOverviewVO

`func NewOswPortsSettingOverviewVO() *OswPortsSettingOverviewVO`

NewOswPortsSettingOverviewVO instantiates a new OswPortsSettingOverviewVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPortsSettingOverviewVOWithDefaults

`func NewOswPortsSettingOverviewVOWithDefaults() *OswPortsSettingOverviewVO`

NewOswPortsSettingOverviewVOWithDefaults instantiates a new OswPortsSettingOverviewVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OswPortsSettingOverviewVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OswPortsSettingOverviewVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OswPortsSettingOverviewVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OswPortsSettingOverviewVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetConfigMlagDad

`func (o *OswPortsSettingOverviewVO) GetConfigMlagDad() bool`

GetConfigMlagDad returns the ConfigMlagDad field if non-nil, zero value otherwise.

### GetConfigMlagDadOk

`func (o *OswPortsSettingOverviewVO) GetConfigMlagDadOk() (*bool, bool)`

GetConfigMlagDadOk returns a tuple with the ConfigMlagDad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagDad

`func (o *OswPortsSettingOverviewVO) SetConfigMlagDad(v bool)`

SetConfigMlagDad sets ConfigMlagDad field to given value.

### HasConfigMlagDad

`func (o *OswPortsSettingOverviewVO) HasConfigMlagDad() bool`

HasConfigMlagDad returns a boolean if a field has been set.

### GetConfigMlagPeerLink

`func (o *OswPortsSettingOverviewVO) GetConfigMlagPeerLink() bool`

GetConfigMlagPeerLink returns the ConfigMlagPeerLink field if non-nil, zero value otherwise.

### GetConfigMlagPeerLinkOk

`func (o *OswPortsSettingOverviewVO) GetConfigMlagPeerLinkOk() (*bool, bool)`

GetConfigMlagPeerLinkOk returns a tuple with the ConfigMlagPeerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagPeerLink

`func (o *OswPortsSettingOverviewVO) SetConfigMlagPeerLink(v bool)`

SetConfigMlagPeerLink sets ConfigMlagPeerLink field to given value.

### HasConfigMlagPeerLink

`func (o *OswPortsSettingOverviewVO) HasConfigMlagPeerLink() bool`

HasConfigMlagPeerLink returns a boolean if a field has been set.

### GetConfigStack

`func (o *OswPortsSettingOverviewVO) GetConfigStack() bool`

GetConfigStack returns the ConfigStack field if non-nil, zero value otherwise.

### GetConfigStackOk

`func (o *OswPortsSettingOverviewVO) GetConfigStackOk() (*bool, bool)`

GetConfigStackOk returns a tuple with the ConfigStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStack

`func (o *OswPortsSettingOverviewVO) SetConfigStack(v bool)`

SetConfigStack sets ConfigStack field to given value.

### HasConfigStack

`func (o *OswPortsSettingOverviewVO) HasConfigStack() bool`

HasConfigStack returns a boolean if a field has been set.

### GetConnectedStatus

`func (o *OswPortsSettingOverviewVO) GetConnectedStatus() int32`

GetConnectedStatus returns the ConnectedStatus field if non-nil, zero value otherwise.

### GetConnectedStatusOk

`func (o *OswPortsSettingOverviewVO) GetConnectedStatusOk() (*int32, bool)`

GetConnectedStatusOk returns a tuple with the ConnectedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedStatus

`func (o *OswPortsSettingOverviewVO) SetConnectedStatus(v int32)`

SetConnectedStatus sets ConnectedStatus field to given value.

### HasConnectedStatus

`func (o *OswPortsSettingOverviewVO) HasConnectedStatus() bool`

HasConnectedStatus returns a boolean if a field has been set.

### GetDisable

`func (o *OswPortsSettingOverviewVO) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *OswPortsSettingOverviewVO) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *OswPortsSettingOverviewVO) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *OswPortsSettingOverviewVO) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDuplex

`func (o *OswPortsSettingOverviewVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswPortsSettingOverviewVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswPortsSettingOverviewVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswPortsSettingOverviewVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLagSetting

`func (o *OswPortsSettingOverviewVO) GetLagSetting() OswStackLagVO`

GetLagSetting returns the LagSetting field if non-nil, zero value otherwise.

### GetLagSettingOk

`func (o *OswPortsSettingOverviewVO) GetLagSettingOk() (*OswStackLagVO, bool)`

GetLagSettingOk returns a tuple with the LagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagSetting

`func (o *OswPortsSettingOverviewVO) SetLagSetting(v OswStackLagVO)`

SetLagSetting sets LagSetting field to given value.

### HasLagSetting

`func (o *OswPortsSettingOverviewVO) HasLagSetting() bool`

HasLagSetting returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswPortsSettingOverviewVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswPortsSettingOverviewVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswPortsSettingOverviewVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswPortsSettingOverviewVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetMadUsed

`func (o *OswPortsSettingOverviewVO) GetMadUsed() bool`

GetMadUsed returns the MadUsed field if non-nil, zero value otherwise.

### GetMadUsedOk

`func (o *OswPortsSettingOverviewVO) GetMadUsedOk() (*bool, bool)`

GetMadUsedOk returns a tuple with the MadUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadUsed

`func (o *OswPortsSettingOverviewVO) SetMadUsed(v bool)`

SetMadUsed sets MadUsed field to given value.

### HasMadUsed

`func (o *OswPortsSettingOverviewVO) HasMadUsed() bool`

HasMadUsed returns a boolean if a field has been set.

### GetNativeBridgeVlan

`func (o *OswPortsSettingOverviewVO) GetNativeBridgeVlan() int32`

GetNativeBridgeVlan returns the NativeBridgeVlan field if non-nil, zero value otherwise.

### GetNativeBridgeVlanOk

`func (o *OswPortsSettingOverviewVO) GetNativeBridgeVlanOk() (*int32, bool)`

GetNativeBridgeVlanOk returns a tuple with the NativeBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeBridgeVlan

`func (o *OswPortsSettingOverviewVO) SetNativeBridgeVlan(v int32)`

SetNativeBridgeVlan sets NativeBridgeVlan field to given value.

### HasNativeBridgeVlan

`func (o *OswPortsSettingOverviewVO) HasNativeBridgeVlan() bool`

HasNativeBridgeVlan returns a boolean if a field has been set.

### GetNativeNetworkId

`func (o *OswPortsSettingOverviewVO) GetNativeNetworkId() string`

GetNativeNetworkId returns the NativeNetworkId field if non-nil, zero value otherwise.

### GetNativeNetworkIdOk

`func (o *OswPortsSettingOverviewVO) GetNativeNetworkIdOk() (*string, bool)`

GetNativeNetworkIdOk returns a tuple with the NativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkId

`func (o *OswPortsSettingOverviewVO) SetNativeNetworkId(v string)`

SetNativeNetworkId sets NativeNetworkId field to given value.

### HasNativeNetworkId

`func (o *OswPortsSettingOverviewVO) HasNativeNetworkId() bool`

HasNativeNetworkId returns a boolean if a field has been set.

### GetNetworkMode

`func (o *OswPortsSettingOverviewVO) GetNetworkMode() int32`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *OswPortsSettingOverviewVO) GetNetworkModeOk() (*int32, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *OswPortsSettingOverviewVO) SetNetworkMode(v int32)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *OswPortsSettingOverviewVO) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### GetNetworkTagsSetting

`func (o *OswPortsSettingOverviewVO) GetNetworkTagsSetting() int32`

GetNetworkTagsSetting returns the NetworkTagsSetting field if non-nil, zero value otherwise.

### GetNetworkTagsSettingOk

`func (o *OswPortsSettingOverviewVO) GetNetworkTagsSettingOk() (*int32, bool)`

GetNetworkTagsSettingOk returns a tuple with the NetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTagsSetting

`func (o *OswPortsSettingOverviewVO) SetNetworkTagsSetting(v int32)`

SetNetworkTagsSetting sets NetworkTagsSetting field to given value.

### HasNetworkTagsSetting

`func (o *OswPortsSettingOverviewVO) HasNetworkTagsSetting() bool`

HasNetworkTagsSetting returns a boolean if a field has been set.

### GetOperation

`func (o *OswPortsSettingOverviewVO) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *OswPortsSettingOverviewVO) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *OswPortsSettingOverviewVO) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *OswPortsSettingOverviewVO) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPoe

`func (o *OswPortsSettingOverviewVO) GetPoe() int32`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *OswPortsSettingOverviewVO) GetPoeOk() (*int32, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *OswPortsSettingOverviewVO) SetPoe(v int32)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *OswPortsSettingOverviewVO) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetPoeDisplayType

`func (o *OswPortsSettingOverviewVO) GetPoeDisplayType() int32`

GetPoeDisplayType returns the PoeDisplayType field if non-nil, zero value otherwise.

### GetPoeDisplayTypeOk

`func (o *OswPortsSettingOverviewVO) GetPoeDisplayTypeOk() (*int32, bool)`

GetPoeDisplayTypeOk returns a tuple with the PoeDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeDisplayType

`func (o *OswPortsSettingOverviewVO) SetPoeDisplayType(v int32)`

SetPoeDisplayType sets PoeDisplayType field to given value.

### HasPoeDisplayType

`func (o *OswPortsSettingOverviewVO) HasPoeDisplayType() bool`

HasPoeDisplayType returns a boolean if a field has been set.

### GetPort

`func (o *OswPortsSettingOverviewVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswPortsSettingOverviewVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswPortsSettingOverviewVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswPortsSettingOverviewVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortName

`func (o *OswPortsSettingOverviewVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *OswPortsSettingOverviewVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *OswPortsSettingOverviewVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *OswPortsSettingOverviewVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPortStatus

`func (o *OswPortsSettingOverviewVO) GetPortStatus() OswPortStatusVO`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *OswPortsSettingOverviewVO) GetPortStatusOk() (*OswPortStatusVO, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *OswPortsSettingOverviewVO) SetPortStatus(v OswPortStatusVO)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *OswPortsSettingOverviewVO) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetProfileId

`func (o *OswPortsSettingOverviewVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *OswPortsSettingOverviewVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *OswPortsSettingOverviewVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *OswPortsSettingOverviewVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileName

`func (o *OswPortsSettingOverviewVO) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *OswPortsSettingOverviewVO) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *OswPortsSettingOverviewVO) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *OswPortsSettingOverviewVO) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetProfileOverrideEnable

`func (o *OswPortsSettingOverviewVO) GetProfileOverrideEnable() bool`

GetProfileOverrideEnable returns the ProfileOverrideEnable field if non-nil, zero value otherwise.

### GetProfileOverrideEnableOk

`func (o *OswPortsSettingOverviewVO) GetProfileOverrideEnableOk() (*bool, bool)`

GetProfileOverrideEnableOk returns a tuple with the ProfileOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOverrideEnable

`func (o *OswPortsSettingOverviewVO) SetProfileOverrideEnable(v bool)`

SetProfileOverrideEnable sets ProfileOverrideEnable field to given value.

### HasProfileOverrideEnable

`func (o *OswPortsSettingOverviewVO) HasProfileOverrideEnable() bool`

HasProfileOverrideEnable returns a boolean if a field has been set.

### GetProfileType

`func (o *OswPortsSettingOverviewVO) GetProfileType() int32`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *OswPortsSettingOverviewVO) GetProfileTypeOk() (*int32, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileType

`func (o *OswPortsSettingOverviewVO) SetProfileType(v int32)`

SetProfileType sets ProfileType field to given value.

### HasProfileType

`func (o *OswPortsSettingOverviewVO) HasProfileType() bool`

HasProfileType returns a boolean if a field has been set.

### GetStackId

`func (o *OswPortsSettingOverviewVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswPortsSettingOverviewVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswPortsSettingOverviewVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswPortsSettingOverviewVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *OswPortsSettingOverviewVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *OswPortsSettingOverviewVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *OswPortsSettingOverviewVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *OswPortsSettingOverviewVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStackPortsGroupIndex

`func (o *OswPortsSettingOverviewVO) GetStackPortsGroupIndex() int32`

GetStackPortsGroupIndex returns the StackPortsGroupIndex field if non-nil, zero value otherwise.

### GetStackPortsGroupIndexOk

`func (o *OswPortsSettingOverviewVO) GetStackPortsGroupIndexOk() (*int32, bool)`

GetStackPortsGroupIndexOk returns a tuple with the StackPortsGroupIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPortsGroupIndex

`func (o *OswPortsSettingOverviewVO) SetStackPortsGroupIndex(v int32)`

SetStackPortsGroupIndex sets StackPortsGroupIndex field to given value.

### HasStackPortsGroupIndex

`func (o *OswPortsSettingOverviewVO) HasStackPortsGroupIndex() bool`

HasStackPortsGroupIndex returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswPortsSettingOverviewVO) GetStandardPort() OswStandPortVO`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswPortsSettingOverviewVO) GetStandardPortOk() (*OswStandPortVO, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswPortsSettingOverviewVO) SetStandardPort(v OswStandPortVO)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswPortsSettingOverviewVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetSupportPoe

`func (o *OswPortsSettingOverviewVO) GetSupportPoe() bool`

GetSupportPoe returns the SupportPoe field if non-nil, zero value otherwise.

### GetSupportPoeOk

`func (o *OswPortsSettingOverviewVO) GetSupportPoeOk() (*bool, bool)`

GetSupportPoeOk returns a tuple with the SupportPoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPoe

`func (o *OswPortsSettingOverviewVO) SetSupportPoe(v bool)`

SetSupportPoe sets SupportPoe field to given value.

### HasSupportPoe

`func (o *OswPortsSettingOverviewVO) HasSupportPoe() bool`

HasSupportPoe returns a boolean if a field has been set.

### GetSwitchMac

`func (o *OswPortsSettingOverviewVO) GetSwitchMac() string`

GetSwitchMac returns the SwitchMac field if non-nil, zero value otherwise.

### GetSwitchMacOk

`func (o *OswPortsSettingOverviewVO) GetSwitchMacOk() (*string, bool)`

GetSwitchMacOk returns a tuple with the SwitchMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMac

`func (o *OswPortsSettingOverviewVO) SetSwitchMac(v string)`

SetSwitchMac sets SwitchMac field to given value.

### HasSwitchMac

`func (o *OswPortsSettingOverviewVO) HasSwitchMac() bool`

HasSwitchMac returns a boolean if a field has been set.

### GetSwitchName

`func (o *OswPortsSettingOverviewVO) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *OswPortsSettingOverviewVO) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *OswPortsSettingOverviewVO) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *OswPortsSettingOverviewVO) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetSwitchStatusCategory

`func (o *OswPortsSettingOverviewVO) GetSwitchStatusCategory() int32`

GetSwitchStatusCategory returns the SwitchStatusCategory field if non-nil, zero value otherwise.

### GetSwitchStatusCategoryOk

`func (o *OswPortsSettingOverviewVO) GetSwitchStatusCategoryOk() (*int32, bool)`

GetSwitchStatusCategoryOk returns a tuple with the SwitchStatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchStatusCategory

`func (o *OswPortsSettingOverviewVO) SetSwitchStatusCategory(v int32)`

SetSwitchStatusCategory sets SwitchStatusCategory field to given value.

### HasSwitchStatusCategory

`func (o *OswPortsSettingOverviewVO) HasSwitchStatusCategory() bool`

HasSwitchStatusCategory returns a boolean if a field has been set.

### GetSwitchSupportPoe

`func (o *OswPortsSettingOverviewVO) GetSwitchSupportPoe() int32`

GetSwitchSupportPoe returns the SwitchSupportPoe field if non-nil, zero value otherwise.

### GetSwitchSupportPoeOk

`func (o *OswPortsSettingOverviewVO) GetSwitchSupportPoeOk() (*int32, bool)`

GetSwitchSupportPoeOk returns a tuple with the SwitchSupportPoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchSupportPoe

`func (o *OswPortsSettingOverviewVO) SetSwitchSupportPoe(v int32)`

SetSwitchSupportPoe sets SwitchSupportPoe field to given value.

### HasSwitchSupportPoe

`func (o *OswPortsSettingOverviewVO) HasSwitchSupportPoe() bool`

HasSwitchSupportPoe returns a boolean if a field has been set.

### GetSwitchType

`func (o *OswPortsSettingOverviewVO) GetSwitchType() int32`

GetSwitchType returns the SwitchType field if non-nil, zero value otherwise.

### GetSwitchTypeOk

`func (o *OswPortsSettingOverviewVO) GetSwitchTypeOk() (*int32, bool)`

GetSwitchTypeOk returns a tuple with the SwitchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchType

`func (o *OswPortsSettingOverviewVO) SetSwitchType(v int32)`

SetSwitchType sets SwitchType field to given value.

### HasSwitchType

`func (o *OswPortsSettingOverviewVO) HasSwitchType() bool`

HasSwitchType returns a boolean if a field has been set.

### GetTagIds

`func (o *OswPortsSettingOverviewVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OswPortsSettingOverviewVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OswPortsSettingOverviewVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OswPortsSettingOverviewVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagName

`func (o *OswPortsSettingOverviewVO) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *OswPortsSettingOverviewVO) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *OswPortsSettingOverviewVO) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *OswPortsSettingOverviewVO) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetType

`func (o *OswPortsSettingOverviewVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswPortsSettingOverviewVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswPortsSettingOverviewVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OswPortsSettingOverviewVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *OswPortsSettingOverviewVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswPortsSettingOverviewVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswPortsSettingOverviewVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswPortsSettingOverviewVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


