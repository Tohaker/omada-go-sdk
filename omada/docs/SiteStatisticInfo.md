# SiteStatisticInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceAccount** | Pointer to [**DeviceAccountSettingVO**](DeviceAccountSettingVO.md) |  | [optional] 
**EapHealth** | Pointer to [**HealthStatVO**](HealthStatVO.md) |  | [optional] 
**GatewayHealth** | Pointer to **int32** |  | [optional] 
**GatewayStatus** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IssueEvent** | Pointer to [**AnomalyStatVO**](AnomalyStatVO.md) |  | [optional] 
**Lan** | Pointer to **bool** |  | [optional] 
**LanDeviceConnectedNum** | Pointer to **int64** |  | [optional] 
**LanDeviceDisconnectedNum** | Pointer to **int64** |  | [optional] 
**LanGuestNum** | Pointer to **int64** |  | [optional] 
**LanUserNum** | Pointer to **int64** |  | [optional] 
**OltDeviceConnectedNum** | Pointer to **int64** |  | [optional] 
**OltDeviceDisconnectedNum** | Pointer to **int64** |  | [optional] 
**PreConfigApNum** | Pointer to **int64** |  | [optional] 
**PreConfigOltNum** | Pointer to **int64** |  | [optional] 
**PreConfigOsgNum** | Pointer to **int64** |  | [optional] 
**PreConfigOswNum** | Pointer to **int64** |  | [optional] 
**SwitchHealth** | Pointer to [**HealthStatVO**](HealthStatVO.md) |  | [optional] 
**WanHealth** | Pointer to [**WanHealthStatVO**](WanHealthStatVO.md) |  | [optional] 
**WiredClientHealth** | Pointer to [**HealthStatVO**](HealthStatVO.md) |  | [optional] 
**WiredUpgrade** | Pointer to **bool** |  | [optional] 
**WirelessClientHealth** | Pointer to [**HealthStatVO**](HealthStatVO.md) |  | [optional] 
**WirelessUpgrade** | Pointer to **bool** |  | [optional] 
**Wlan** | Pointer to **bool** |  | [optional] 
**WlanDeviceConnectedNum** | Pointer to **int64** |  | [optional] 
**WlanDeviceDisconnectedNum** | Pointer to **int64** |  | [optional] 
**WlanDeviceIsolatedNum** | Pointer to **int64** |  | [optional] 
**WlanGuestNum** | Pointer to **int64** |  | [optional] 
**WlanHealth** | Pointer to **int32** |  | [optional] 
**WlanUserNum** | Pointer to **int64** |  | [optional] 

## Methods

### NewSiteStatisticInfo

`func NewSiteStatisticInfo() *SiteStatisticInfo`

NewSiteStatisticInfo instantiates a new SiteStatisticInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteStatisticInfoWithDefaults

`func NewSiteStatisticInfoWithDefaults() *SiteStatisticInfo`

NewSiteStatisticInfoWithDefaults instantiates a new SiteStatisticInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceAccount

`func (o *SiteStatisticInfo) GetDeviceAccount() DeviceAccountSettingVO`

GetDeviceAccount returns the DeviceAccount field if non-nil, zero value otherwise.

### GetDeviceAccountOk

`func (o *SiteStatisticInfo) GetDeviceAccountOk() (*DeviceAccountSettingVO, bool)`

GetDeviceAccountOk returns a tuple with the DeviceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAccount

`func (o *SiteStatisticInfo) SetDeviceAccount(v DeviceAccountSettingVO)`

SetDeviceAccount sets DeviceAccount field to given value.

### HasDeviceAccount

`func (o *SiteStatisticInfo) HasDeviceAccount() bool`

HasDeviceAccount returns a boolean if a field has been set.

### GetEapHealth

`func (o *SiteStatisticInfo) GetEapHealth() HealthStatVO`

GetEapHealth returns the EapHealth field if non-nil, zero value otherwise.

### GetEapHealthOk

`func (o *SiteStatisticInfo) GetEapHealthOk() (*HealthStatVO, bool)`

GetEapHealthOk returns a tuple with the EapHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapHealth

`func (o *SiteStatisticInfo) SetEapHealth(v HealthStatVO)`

SetEapHealth sets EapHealth field to given value.

### HasEapHealth

`func (o *SiteStatisticInfo) HasEapHealth() bool`

HasEapHealth returns a boolean if a field has been set.

### GetGatewayHealth

`func (o *SiteStatisticInfo) GetGatewayHealth() int32`

GetGatewayHealth returns the GatewayHealth field if non-nil, zero value otherwise.

### GetGatewayHealthOk

`func (o *SiteStatisticInfo) GetGatewayHealthOk() (*int32, bool)`

GetGatewayHealthOk returns a tuple with the GatewayHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayHealth

`func (o *SiteStatisticInfo) SetGatewayHealth(v int32)`

SetGatewayHealth sets GatewayHealth field to given value.

### HasGatewayHealth

`func (o *SiteStatisticInfo) HasGatewayHealth() bool`

HasGatewayHealth returns a boolean if a field has been set.

### GetGatewayStatus

`func (o *SiteStatisticInfo) GetGatewayStatus() int32`

GetGatewayStatus returns the GatewayStatus field if non-nil, zero value otherwise.

### GetGatewayStatusOk

`func (o *SiteStatisticInfo) GetGatewayStatusOk() (*int32, bool)`

GetGatewayStatusOk returns a tuple with the GatewayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayStatus

`func (o *SiteStatisticInfo) SetGatewayStatus(v int32)`

SetGatewayStatus sets GatewayStatus field to given value.

### HasGatewayStatus

`func (o *SiteStatisticInfo) HasGatewayStatus() bool`

HasGatewayStatus returns a boolean if a field has been set.

### GetId

`func (o *SiteStatisticInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteStatisticInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteStatisticInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SiteStatisticInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueEvent

`func (o *SiteStatisticInfo) GetIssueEvent() AnomalyStatVO`

GetIssueEvent returns the IssueEvent field if non-nil, zero value otherwise.

### GetIssueEventOk

`func (o *SiteStatisticInfo) GetIssueEventOk() (*AnomalyStatVO, bool)`

GetIssueEventOk returns a tuple with the IssueEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueEvent

`func (o *SiteStatisticInfo) SetIssueEvent(v AnomalyStatVO)`

SetIssueEvent sets IssueEvent field to given value.

### HasIssueEvent

`func (o *SiteStatisticInfo) HasIssueEvent() bool`

HasIssueEvent returns a boolean if a field has been set.

### GetLan

`func (o *SiteStatisticInfo) GetLan() bool`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *SiteStatisticInfo) GetLanOk() (*bool, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *SiteStatisticInfo) SetLan(v bool)`

SetLan sets Lan field to given value.

### HasLan

`func (o *SiteStatisticInfo) HasLan() bool`

HasLan returns a boolean if a field has been set.

### GetLanDeviceConnectedNum

`func (o *SiteStatisticInfo) GetLanDeviceConnectedNum() int64`

GetLanDeviceConnectedNum returns the LanDeviceConnectedNum field if non-nil, zero value otherwise.

### GetLanDeviceConnectedNumOk

`func (o *SiteStatisticInfo) GetLanDeviceConnectedNumOk() (*int64, bool)`

GetLanDeviceConnectedNumOk returns a tuple with the LanDeviceConnectedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanDeviceConnectedNum

`func (o *SiteStatisticInfo) SetLanDeviceConnectedNum(v int64)`

SetLanDeviceConnectedNum sets LanDeviceConnectedNum field to given value.

### HasLanDeviceConnectedNum

`func (o *SiteStatisticInfo) HasLanDeviceConnectedNum() bool`

HasLanDeviceConnectedNum returns a boolean if a field has been set.

### GetLanDeviceDisconnectedNum

`func (o *SiteStatisticInfo) GetLanDeviceDisconnectedNum() int64`

GetLanDeviceDisconnectedNum returns the LanDeviceDisconnectedNum field if non-nil, zero value otherwise.

### GetLanDeviceDisconnectedNumOk

`func (o *SiteStatisticInfo) GetLanDeviceDisconnectedNumOk() (*int64, bool)`

GetLanDeviceDisconnectedNumOk returns a tuple with the LanDeviceDisconnectedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanDeviceDisconnectedNum

`func (o *SiteStatisticInfo) SetLanDeviceDisconnectedNum(v int64)`

SetLanDeviceDisconnectedNum sets LanDeviceDisconnectedNum field to given value.

### HasLanDeviceDisconnectedNum

`func (o *SiteStatisticInfo) HasLanDeviceDisconnectedNum() bool`

HasLanDeviceDisconnectedNum returns a boolean if a field has been set.

### GetLanGuestNum

`func (o *SiteStatisticInfo) GetLanGuestNum() int64`

GetLanGuestNum returns the LanGuestNum field if non-nil, zero value otherwise.

### GetLanGuestNumOk

`func (o *SiteStatisticInfo) GetLanGuestNumOk() (*int64, bool)`

GetLanGuestNumOk returns a tuple with the LanGuestNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanGuestNum

`func (o *SiteStatisticInfo) SetLanGuestNum(v int64)`

SetLanGuestNum sets LanGuestNum field to given value.

### HasLanGuestNum

`func (o *SiteStatisticInfo) HasLanGuestNum() bool`

HasLanGuestNum returns a boolean if a field has been set.

### GetLanUserNum

`func (o *SiteStatisticInfo) GetLanUserNum() int64`

GetLanUserNum returns the LanUserNum field if non-nil, zero value otherwise.

### GetLanUserNumOk

`func (o *SiteStatisticInfo) GetLanUserNumOk() (*int64, bool)`

GetLanUserNumOk returns a tuple with the LanUserNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanUserNum

`func (o *SiteStatisticInfo) SetLanUserNum(v int64)`

SetLanUserNum sets LanUserNum field to given value.

### HasLanUserNum

`func (o *SiteStatisticInfo) HasLanUserNum() bool`

HasLanUserNum returns a boolean if a field has been set.

### GetOltDeviceConnectedNum

`func (o *SiteStatisticInfo) GetOltDeviceConnectedNum() int64`

GetOltDeviceConnectedNum returns the OltDeviceConnectedNum field if non-nil, zero value otherwise.

### GetOltDeviceConnectedNumOk

`func (o *SiteStatisticInfo) GetOltDeviceConnectedNumOk() (*int64, bool)`

GetOltDeviceConnectedNumOk returns a tuple with the OltDeviceConnectedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOltDeviceConnectedNum

`func (o *SiteStatisticInfo) SetOltDeviceConnectedNum(v int64)`

SetOltDeviceConnectedNum sets OltDeviceConnectedNum field to given value.

### HasOltDeviceConnectedNum

`func (o *SiteStatisticInfo) HasOltDeviceConnectedNum() bool`

HasOltDeviceConnectedNum returns a boolean if a field has been set.

### GetOltDeviceDisconnectedNum

`func (o *SiteStatisticInfo) GetOltDeviceDisconnectedNum() int64`

GetOltDeviceDisconnectedNum returns the OltDeviceDisconnectedNum field if non-nil, zero value otherwise.

### GetOltDeviceDisconnectedNumOk

`func (o *SiteStatisticInfo) GetOltDeviceDisconnectedNumOk() (*int64, bool)`

GetOltDeviceDisconnectedNumOk returns a tuple with the OltDeviceDisconnectedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOltDeviceDisconnectedNum

`func (o *SiteStatisticInfo) SetOltDeviceDisconnectedNum(v int64)`

SetOltDeviceDisconnectedNum sets OltDeviceDisconnectedNum field to given value.

### HasOltDeviceDisconnectedNum

`func (o *SiteStatisticInfo) HasOltDeviceDisconnectedNum() bool`

HasOltDeviceDisconnectedNum returns a boolean if a field has been set.

### GetPreConfigApNum

`func (o *SiteStatisticInfo) GetPreConfigApNum() int64`

GetPreConfigApNum returns the PreConfigApNum field if non-nil, zero value otherwise.

### GetPreConfigApNumOk

`func (o *SiteStatisticInfo) GetPreConfigApNumOk() (*int64, bool)`

GetPreConfigApNumOk returns a tuple with the PreConfigApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConfigApNum

`func (o *SiteStatisticInfo) SetPreConfigApNum(v int64)`

SetPreConfigApNum sets PreConfigApNum field to given value.

### HasPreConfigApNum

`func (o *SiteStatisticInfo) HasPreConfigApNum() bool`

HasPreConfigApNum returns a boolean if a field has been set.

### GetPreConfigOltNum

`func (o *SiteStatisticInfo) GetPreConfigOltNum() int64`

GetPreConfigOltNum returns the PreConfigOltNum field if non-nil, zero value otherwise.

### GetPreConfigOltNumOk

`func (o *SiteStatisticInfo) GetPreConfigOltNumOk() (*int64, bool)`

GetPreConfigOltNumOk returns a tuple with the PreConfigOltNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConfigOltNum

`func (o *SiteStatisticInfo) SetPreConfigOltNum(v int64)`

SetPreConfigOltNum sets PreConfigOltNum field to given value.

### HasPreConfigOltNum

`func (o *SiteStatisticInfo) HasPreConfigOltNum() bool`

HasPreConfigOltNum returns a boolean if a field has been set.

### GetPreConfigOsgNum

`func (o *SiteStatisticInfo) GetPreConfigOsgNum() int64`

GetPreConfigOsgNum returns the PreConfigOsgNum field if non-nil, zero value otherwise.

### GetPreConfigOsgNumOk

`func (o *SiteStatisticInfo) GetPreConfigOsgNumOk() (*int64, bool)`

GetPreConfigOsgNumOk returns a tuple with the PreConfigOsgNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConfigOsgNum

`func (o *SiteStatisticInfo) SetPreConfigOsgNum(v int64)`

SetPreConfigOsgNum sets PreConfigOsgNum field to given value.

### HasPreConfigOsgNum

`func (o *SiteStatisticInfo) HasPreConfigOsgNum() bool`

HasPreConfigOsgNum returns a boolean if a field has been set.

### GetPreConfigOswNum

`func (o *SiteStatisticInfo) GetPreConfigOswNum() int64`

GetPreConfigOswNum returns the PreConfigOswNum field if non-nil, zero value otherwise.

### GetPreConfigOswNumOk

`func (o *SiteStatisticInfo) GetPreConfigOswNumOk() (*int64, bool)`

GetPreConfigOswNumOk returns a tuple with the PreConfigOswNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConfigOswNum

`func (o *SiteStatisticInfo) SetPreConfigOswNum(v int64)`

SetPreConfigOswNum sets PreConfigOswNum field to given value.

### HasPreConfigOswNum

`func (o *SiteStatisticInfo) HasPreConfigOswNum() bool`

HasPreConfigOswNum returns a boolean if a field has been set.

### GetSwitchHealth

`func (o *SiteStatisticInfo) GetSwitchHealth() HealthStatVO`

GetSwitchHealth returns the SwitchHealth field if non-nil, zero value otherwise.

### GetSwitchHealthOk

`func (o *SiteStatisticInfo) GetSwitchHealthOk() (*HealthStatVO, bool)`

GetSwitchHealthOk returns a tuple with the SwitchHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchHealth

`func (o *SiteStatisticInfo) SetSwitchHealth(v HealthStatVO)`

SetSwitchHealth sets SwitchHealth field to given value.

### HasSwitchHealth

`func (o *SiteStatisticInfo) HasSwitchHealth() bool`

HasSwitchHealth returns a boolean if a field has been set.

### GetWanHealth

`func (o *SiteStatisticInfo) GetWanHealth() WanHealthStatVO`

GetWanHealth returns the WanHealth field if non-nil, zero value otherwise.

### GetWanHealthOk

`func (o *SiteStatisticInfo) GetWanHealthOk() (*WanHealthStatVO, bool)`

GetWanHealthOk returns a tuple with the WanHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanHealth

`func (o *SiteStatisticInfo) SetWanHealth(v WanHealthStatVO)`

SetWanHealth sets WanHealth field to given value.

### HasWanHealth

`func (o *SiteStatisticInfo) HasWanHealth() bool`

HasWanHealth returns a boolean if a field has been set.

### GetWiredClientHealth

`func (o *SiteStatisticInfo) GetWiredClientHealth() HealthStatVO`

GetWiredClientHealth returns the WiredClientHealth field if non-nil, zero value otherwise.

### GetWiredClientHealthOk

`func (o *SiteStatisticInfo) GetWiredClientHealthOk() (*HealthStatVO, bool)`

GetWiredClientHealthOk returns a tuple with the WiredClientHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredClientHealth

`func (o *SiteStatisticInfo) SetWiredClientHealth(v HealthStatVO)`

SetWiredClientHealth sets WiredClientHealth field to given value.

### HasWiredClientHealth

`func (o *SiteStatisticInfo) HasWiredClientHealth() bool`

HasWiredClientHealth returns a boolean if a field has been set.

### GetWiredUpgrade

`func (o *SiteStatisticInfo) GetWiredUpgrade() bool`

GetWiredUpgrade returns the WiredUpgrade field if non-nil, zero value otherwise.

### GetWiredUpgradeOk

`func (o *SiteStatisticInfo) GetWiredUpgradeOk() (*bool, bool)`

GetWiredUpgradeOk returns a tuple with the WiredUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredUpgrade

`func (o *SiteStatisticInfo) SetWiredUpgrade(v bool)`

SetWiredUpgrade sets WiredUpgrade field to given value.

### HasWiredUpgrade

`func (o *SiteStatisticInfo) HasWiredUpgrade() bool`

HasWiredUpgrade returns a boolean if a field has been set.

### GetWirelessClientHealth

`func (o *SiteStatisticInfo) GetWirelessClientHealth() HealthStatVO`

GetWirelessClientHealth returns the WirelessClientHealth field if non-nil, zero value otherwise.

### GetWirelessClientHealthOk

`func (o *SiteStatisticInfo) GetWirelessClientHealthOk() (*HealthStatVO, bool)`

GetWirelessClientHealthOk returns a tuple with the WirelessClientHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessClientHealth

`func (o *SiteStatisticInfo) SetWirelessClientHealth(v HealthStatVO)`

SetWirelessClientHealth sets WirelessClientHealth field to given value.

### HasWirelessClientHealth

`func (o *SiteStatisticInfo) HasWirelessClientHealth() bool`

HasWirelessClientHealth returns a boolean if a field has been set.

### GetWirelessUpgrade

`func (o *SiteStatisticInfo) GetWirelessUpgrade() bool`

GetWirelessUpgrade returns the WirelessUpgrade field if non-nil, zero value otherwise.

### GetWirelessUpgradeOk

`func (o *SiteStatisticInfo) GetWirelessUpgradeOk() (*bool, bool)`

GetWirelessUpgradeOk returns a tuple with the WirelessUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessUpgrade

`func (o *SiteStatisticInfo) SetWirelessUpgrade(v bool)`

SetWirelessUpgrade sets WirelessUpgrade field to given value.

### HasWirelessUpgrade

`func (o *SiteStatisticInfo) HasWirelessUpgrade() bool`

HasWirelessUpgrade returns a boolean if a field has been set.

### GetWlan

`func (o *SiteStatisticInfo) GetWlan() bool`

GetWlan returns the Wlan field if non-nil, zero value otherwise.

### GetWlanOk

`func (o *SiteStatisticInfo) GetWlanOk() (*bool, bool)`

GetWlanOk returns a tuple with the Wlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlan

`func (o *SiteStatisticInfo) SetWlan(v bool)`

SetWlan sets Wlan field to given value.

### HasWlan

`func (o *SiteStatisticInfo) HasWlan() bool`

HasWlan returns a boolean if a field has been set.

### GetWlanDeviceConnectedNum

`func (o *SiteStatisticInfo) GetWlanDeviceConnectedNum() int64`

GetWlanDeviceConnectedNum returns the WlanDeviceConnectedNum field if non-nil, zero value otherwise.

### GetWlanDeviceConnectedNumOk

`func (o *SiteStatisticInfo) GetWlanDeviceConnectedNumOk() (*int64, bool)`

GetWlanDeviceConnectedNumOk returns a tuple with the WlanDeviceConnectedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanDeviceConnectedNum

`func (o *SiteStatisticInfo) SetWlanDeviceConnectedNum(v int64)`

SetWlanDeviceConnectedNum sets WlanDeviceConnectedNum field to given value.

### HasWlanDeviceConnectedNum

`func (o *SiteStatisticInfo) HasWlanDeviceConnectedNum() bool`

HasWlanDeviceConnectedNum returns a boolean if a field has been set.

### GetWlanDeviceDisconnectedNum

`func (o *SiteStatisticInfo) GetWlanDeviceDisconnectedNum() int64`

GetWlanDeviceDisconnectedNum returns the WlanDeviceDisconnectedNum field if non-nil, zero value otherwise.

### GetWlanDeviceDisconnectedNumOk

`func (o *SiteStatisticInfo) GetWlanDeviceDisconnectedNumOk() (*int64, bool)`

GetWlanDeviceDisconnectedNumOk returns a tuple with the WlanDeviceDisconnectedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanDeviceDisconnectedNum

`func (o *SiteStatisticInfo) SetWlanDeviceDisconnectedNum(v int64)`

SetWlanDeviceDisconnectedNum sets WlanDeviceDisconnectedNum field to given value.

### HasWlanDeviceDisconnectedNum

`func (o *SiteStatisticInfo) HasWlanDeviceDisconnectedNum() bool`

HasWlanDeviceDisconnectedNum returns a boolean if a field has been set.

### GetWlanDeviceIsolatedNum

`func (o *SiteStatisticInfo) GetWlanDeviceIsolatedNum() int64`

GetWlanDeviceIsolatedNum returns the WlanDeviceIsolatedNum field if non-nil, zero value otherwise.

### GetWlanDeviceIsolatedNumOk

`func (o *SiteStatisticInfo) GetWlanDeviceIsolatedNumOk() (*int64, bool)`

GetWlanDeviceIsolatedNumOk returns a tuple with the WlanDeviceIsolatedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanDeviceIsolatedNum

`func (o *SiteStatisticInfo) SetWlanDeviceIsolatedNum(v int64)`

SetWlanDeviceIsolatedNum sets WlanDeviceIsolatedNum field to given value.

### HasWlanDeviceIsolatedNum

`func (o *SiteStatisticInfo) HasWlanDeviceIsolatedNum() bool`

HasWlanDeviceIsolatedNum returns a boolean if a field has been set.

### GetWlanGuestNum

`func (o *SiteStatisticInfo) GetWlanGuestNum() int64`

GetWlanGuestNum returns the WlanGuestNum field if non-nil, zero value otherwise.

### GetWlanGuestNumOk

`func (o *SiteStatisticInfo) GetWlanGuestNumOk() (*int64, bool)`

GetWlanGuestNumOk returns a tuple with the WlanGuestNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanGuestNum

`func (o *SiteStatisticInfo) SetWlanGuestNum(v int64)`

SetWlanGuestNum sets WlanGuestNum field to given value.

### HasWlanGuestNum

`func (o *SiteStatisticInfo) HasWlanGuestNum() bool`

HasWlanGuestNum returns a boolean if a field has been set.

### GetWlanHealth

`func (o *SiteStatisticInfo) GetWlanHealth() int32`

GetWlanHealth returns the WlanHealth field if non-nil, zero value otherwise.

### GetWlanHealthOk

`func (o *SiteStatisticInfo) GetWlanHealthOk() (*int32, bool)`

GetWlanHealthOk returns a tuple with the WlanHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanHealth

`func (o *SiteStatisticInfo) SetWlanHealth(v int32)`

SetWlanHealth sets WlanHealth field to given value.

### HasWlanHealth

`func (o *SiteStatisticInfo) HasWlanHealth() bool`

HasWlanHealth returns a boolean if a field has been set.

### GetWlanUserNum

`func (o *SiteStatisticInfo) GetWlanUserNum() int64`

GetWlanUserNum returns the WlanUserNum field if non-nil, zero value otherwise.

### GetWlanUserNumOk

`func (o *SiteStatisticInfo) GetWlanUserNumOk() (*int64, bool)`

GetWlanUserNumOk returns a tuple with the WlanUserNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanUserNum

`func (o *SiteStatisticInfo) SetWlanUserNum(v int64)`

SetWlanUserNum sets WlanUserNum field to given value.

### HasWlanUserNum

`func (o *SiteStatisticInfo) HasWlanUserNum() bool`

HasWlanUserNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


