# ControllerRoleVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddAdoptDevice** | Pointer to **int32** | Add and adopt devices permission should be a value as follows: 0:block; 2:access | [optional] 
**AddDevices** | Pointer to **int32** | Add devices permission should be a value as follows: 0:block; 2:access | [optional] 
**Adopt** | Pointer to **int32** | Adopt permission should be a value as follows: 0:block; 2:access | [optional] 
**Analyze** | Pointer to **int32** | Tools permission in global view. Only for hardware controller should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**Anomaly** | Pointer to **int32** | Anomaly permission in global view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**Clients** | Pointer to **int32** | Clients permission in site view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**Dashboard** | Pointer to **int32** | Dashboard permission in site view. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**DeviceAccount** | Pointer to **int32** | Device account permission in site view -&gt; site settings should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**DeviceRecovery** | Pointer to **int32** | Device recovery permission in site view. 0:block; 1:view only; 2:modify | [optional] 
**Devices** | Pointer to **int32** | Devices permission in site view and global view. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**ExportData** | Pointer to **int32** | Export data permission in site view should be a value as follows: 0:block; 2:access | [optional] 
**ExportGlobalLog** | Pointer to **int32** | Export global log data permission in global view should be a value as follows: 0:block; 2:access | [optional] 
**FirmwareManager** | Pointer to **int32** | FirmwareManager in global view should be a value as follows: 0:block; 1:view only; 2:modify, only for pro controller | [optional] 
**GlobalCluster** | Pointer to **int32** | Cluster permission in global view. 0:block; 1:view only; 2:modify | [optional] 
**GlobalDashboard** | Pointer to **int32** | Dashboard permission in global view. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**GlobalExportData** | Pointer to **int32** | Export data permission in global view should be a value as follows: 0:block; 2:access | [optional] 
**GlobalLog** | Pointer to **int32** | Log permission in global view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**GlobalMapToken** | Pointer to **int32** | MapToken in global view should be a value as follows: 0:block; 1:view only; 2:modify, only for pro controller | [optional] 
**GlobalSecurity** | Pointer to **int32** | Security permission in global view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**GlobalSetting** | Pointer to **int32** | Settings permission in global view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**GlobalWebhook** | Pointer to **int32** | Webhook in global view should be a value as follows: 0:block; 1:view only; 2:modify, only for pro controller | [optional] 
**Hotspot** | Pointer to **int32** | Hotspot permission should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**Insight** | Pointer to **int32** | Insight permission in site view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**License** | Pointer to **int32** | License permission, it should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**LicenseBind** | Pointer to **int32** | License bind permission should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**Log** | Pointer to **int32** | Log permission in site view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**ManualUpgrade** | Pointer to **int32** | Manual Upgrade permission should be a value as follows: 0:block; 2:access | [optional] 
**Map** | Pointer to **int32** | Map permission in site view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**Network** | Pointer to **int32** | Site network settings permission in site view -&gt; settings should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**Report** | Pointer to **int32** | Network report permission in site view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**Roles** | Pointer to **int32** | Roles permission in global view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**SamlRoles** | Pointer to **int32** | Saml roles permission in global view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**SamlSsos** | Pointer to **int32** | Saml ssos permission in global view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**SamlUsers** | Pointer to **int32** | Saml users permission in global view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**SdWan** | Pointer to **int32** | SD-WAN in global view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**SiteAnalyze** | Pointer to **int32** | Tools in site view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**SiteDevice** | Pointer to **int32** | Device config page permission in site view. 0:block; 1:view only; 2:modify | [optional] 
**SiteHome** | Pointer to **int32** | Site home permission in site view. 0:block; 1:view only; 2:modify | [optional] 
**SiteMaintain** | Pointer to **int32** | Maintenance page permission in site view. 0:block; 1:view only; 2:modify | [optional] 
**SiteNetwork** | Pointer to **int32** | Network config page permission in site view. 0:block; 1:view only; 2:modify | [optional] 
**SiteTemplate** | Pointer to **int32** | SiteTemplate in global view should be a value as follows: 0:block; 1:view only; 2:modify, only for pro controller | [optional] 
**Statics** | Pointer to **int32** | Statics permission in site view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**Users** | Pointer to **int32** | Users permission in global view should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 

## Methods

### NewControllerRoleVO

`func NewControllerRoleVO() *ControllerRoleVO`

NewControllerRoleVO instantiates a new ControllerRoleVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerRoleVOWithDefaults

`func NewControllerRoleVOWithDefaults() *ControllerRoleVO`

NewControllerRoleVOWithDefaults instantiates a new ControllerRoleVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddAdoptDevice

`func (o *ControllerRoleVO) GetAddAdoptDevice() int32`

GetAddAdoptDevice returns the AddAdoptDevice field if non-nil, zero value otherwise.

### GetAddAdoptDeviceOk

`func (o *ControllerRoleVO) GetAddAdoptDeviceOk() (*int32, bool)`

GetAddAdoptDeviceOk returns a tuple with the AddAdoptDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAdoptDevice

`func (o *ControllerRoleVO) SetAddAdoptDevice(v int32)`

SetAddAdoptDevice sets AddAdoptDevice field to given value.

### HasAddAdoptDevice

`func (o *ControllerRoleVO) HasAddAdoptDevice() bool`

HasAddAdoptDevice returns a boolean if a field has been set.

### GetAddDevices

`func (o *ControllerRoleVO) GetAddDevices() int32`

GetAddDevices returns the AddDevices field if non-nil, zero value otherwise.

### GetAddDevicesOk

`func (o *ControllerRoleVO) GetAddDevicesOk() (*int32, bool)`

GetAddDevicesOk returns a tuple with the AddDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddDevices

`func (o *ControllerRoleVO) SetAddDevices(v int32)`

SetAddDevices sets AddDevices field to given value.

### HasAddDevices

`func (o *ControllerRoleVO) HasAddDevices() bool`

HasAddDevices returns a boolean if a field has been set.

### GetAdopt

`func (o *ControllerRoleVO) GetAdopt() int32`

GetAdopt returns the Adopt field if non-nil, zero value otherwise.

### GetAdoptOk

`func (o *ControllerRoleVO) GetAdoptOk() (*int32, bool)`

GetAdoptOk returns a tuple with the Adopt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdopt

`func (o *ControllerRoleVO) SetAdopt(v int32)`

SetAdopt sets Adopt field to given value.

### HasAdopt

`func (o *ControllerRoleVO) HasAdopt() bool`

HasAdopt returns a boolean if a field has been set.

### GetAnalyze

`func (o *ControllerRoleVO) GetAnalyze() int32`

GetAnalyze returns the Analyze field if non-nil, zero value otherwise.

### GetAnalyzeOk

`func (o *ControllerRoleVO) GetAnalyzeOk() (*int32, bool)`

GetAnalyzeOk returns a tuple with the Analyze field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyze

`func (o *ControllerRoleVO) SetAnalyze(v int32)`

SetAnalyze sets Analyze field to given value.

### HasAnalyze

`func (o *ControllerRoleVO) HasAnalyze() bool`

HasAnalyze returns a boolean if a field has been set.

### GetAnomaly

`func (o *ControllerRoleVO) GetAnomaly() int32`

GetAnomaly returns the Anomaly field if non-nil, zero value otherwise.

### GetAnomalyOk

`func (o *ControllerRoleVO) GetAnomalyOk() (*int32, bool)`

GetAnomalyOk returns a tuple with the Anomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomaly

`func (o *ControllerRoleVO) SetAnomaly(v int32)`

SetAnomaly sets Anomaly field to given value.

### HasAnomaly

`func (o *ControllerRoleVO) HasAnomaly() bool`

HasAnomaly returns a boolean if a field has been set.

### GetClients

`func (o *ControllerRoleVO) GetClients() int32`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ControllerRoleVO) GetClientsOk() (*int32, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ControllerRoleVO) SetClients(v int32)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ControllerRoleVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetDashboard

`func (o *ControllerRoleVO) GetDashboard() int32`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *ControllerRoleVO) GetDashboardOk() (*int32, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *ControllerRoleVO) SetDashboard(v int32)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *ControllerRoleVO) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetDeviceAccount

`func (o *ControllerRoleVO) GetDeviceAccount() int32`

GetDeviceAccount returns the DeviceAccount field if non-nil, zero value otherwise.

### GetDeviceAccountOk

`func (o *ControllerRoleVO) GetDeviceAccountOk() (*int32, bool)`

GetDeviceAccountOk returns a tuple with the DeviceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAccount

`func (o *ControllerRoleVO) SetDeviceAccount(v int32)`

SetDeviceAccount sets DeviceAccount field to given value.

### HasDeviceAccount

`func (o *ControllerRoleVO) HasDeviceAccount() bool`

HasDeviceAccount returns a boolean if a field has been set.

### GetDeviceRecovery

`func (o *ControllerRoleVO) GetDeviceRecovery() int32`

GetDeviceRecovery returns the DeviceRecovery field if non-nil, zero value otherwise.

### GetDeviceRecoveryOk

`func (o *ControllerRoleVO) GetDeviceRecoveryOk() (*int32, bool)`

GetDeviceRecoveryOk returns a tuple with the DeviceRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRecovery

`func (o *ControllerRoleVO) SetDeviceRecovery(v int32)`

SetDeviceRecovery sets DeviceRecovery field to given value.

### HasDeviceRecovery

`func (o *ControllerRoleVO) HasDeviceRecovery() bool`

HasDeviceRecovery returns a boolean if a field has been set.

### GetDevices

`func (o *ControllerRoleVO) GetDevices() int32`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ControllerRoleVO) GetDevicesOk() (*int32, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ControllerRoleVO) SetDevices(v int32)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *ControllerRoleVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetExportData

`func (o *ControllerRoleVO) GetExportData() int32`

GetExportData returns the ExportData field if non-nil, zero value otherwise.

### GetExportDataOk

`func (o *ControllerRoleVO) GetExportDataOk() (*int32, bool)`

GetExportDataOk returns a tuple with the ExportData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportData

`func (o *ControllerRoleVO) SetExportData(v int32)`

SetExportData sets ExportData field to given value.

### HasExportData

`func (o *ControllerRoleVO) HasExportData() bool`

HasExportData returns a boolean if a field has been set.

### GetExportGlobalLog

`func (o *ControllerRoleVO) GetExportGlobalLog() int32`

GetExportGlobalLog returns the ExportGlobalLog field if non-nil, zero value otherwise.

### GetExportGlobalLogOk

`func (o *ControllerRoleVO) GetExportGlobalLogOk() (*int32, bool)`

GetExportGlobalLogOk returns a tuple with the ExportGlobalLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportGlobalLog

`func (o *ControllerRoleVO) SetExportGlobalLog(v int32)`

SetExportGlobalLog sets ExportGlobalLog field to given value.

### HasExportGlobalLog

`func (o *ControllerRoleVO) HasExportGlobalLog() bool`

HasExportGlobalLog returns a boolean if a field has been set.

### GetFirmwareManager

`func (o *ControllerRoleVO) GetFirmwareManager() int32`

GetFirmwareManager returns the FirmwareManager field if non-nil, zero value otherwise.

### GetFirmwareManagerOk

`func (o *ControllerRoleVO) GetFirmwareManagerOk() (*int32, bool)`

GetFirmwareManagerOk returns a tuple with the FirmwareManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareManager

`func (o *ControllerRoleVO) SetFirmwareManager(v int32)`

SetFirmwareManager sets FirmwareManager field to given value.

### HasFirmwareManager

`func (o *ControllerRoleVO) HasFirmwareManager() bool`

HasFirmwareManager returns a boolean if a field has been set.

### GetGlobalCluster

`func (o *ControllerRoleVO) GetGlobalCluster() int32`

GetGlobalCluster returns the GlobalCluster field if non-nil, zero value otherwise.

### GetGlobalClusterOk

`func (o *ControllerRoleVO) GetGlobalClusterOk() (*int32, bool)`

GetGlobalClusterOk returns a tuple with the GlobalCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCluster

`func (o *ControllerRoleVO) SetGlobalCluster(v int32)`

SetGlobalCluster sets GlobalCluster field to given value.

### HasGlobalCluster

`func (o *ControllerRoleVO) HasGlobalCluster() bool`

HasGlobalCluster returns a boolean if a field has been set.

### GetGlobalDashboard

`func (o *ControllerRoleVO) GetGlobalDashboard() int32`

GetGlobalDashboard returns the GlobalDashboard field if non-nil, zero value otherwise.

### GetGlobalDashboardOk

`func (o *ControllerRoleVO) GetGlobalDashboardOk() (*int32, bool)`

GetGlobalDashboardOk returns a tuple with the GlobalDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDashboard

`func (o *ControllerRoleVO) SetGlobalDashboard(v int32)`

SetGlobalDashboard sets GlobalDashboard field to given value.

### HasGlobalDashboard

`func (o *ControllerRoleVO) HasGlobalDashboard() bool`

HasGlobalDashboard returns a boolean if a field has been set.

### GetGlobalExportData

`func (o *ControllerRoleVO) GetGlobalExportData() int32`

GetGlobalExportData returns the GlobalExportData field if non-nil, zero value otherwise.

### GetGlobalExportDataOk

`func (o *ControllerRoleVO) GetGlobalExportDataOk() (*int32, bool)`

GetGlobalExportDataOk returns a tuple with the GlobalExportData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalExportData

`func (o *ControllerRoleVO) SetGlobalExportData(v int32)`

SetGlobalExportData sets GlobalExportData field to given value.

### HasGlobalExportData

`func (o *ControllerRoleVO) HasGlobalExportData() bool`

HasGlobalExportData returns a boolean if a field has been set.

### GetGlobalLog

`func (o *ControllerRoleVO) GetGlobalLog() int32`

GetGlobalLog returns the GlobalLog field if non-nil, zero value otherwise.

### GetGlobalLogOk

`func (o *ControllerRoleVO) GetGlobalLogOk() (*int32, bool)`

GetGlobalLogOk returns a tuple with the GlobalLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalLog

`func (o *ControllerRoleVO) SetGlobalLog(v int32)`

SetGlobalLog sets GlobalLog field to given value.

### HasGlobalLog

`func (o *ControllerRoleVO) HasGlobalLog() bool`

HasGlobalLog returns a boolean if a field has been set.

### GetGlobalMapToken

`func (o *ControllerRoleVO) GetGlobalMapToken() int32`

GetGlobalMapToken returns the GlobalMapToken field if non-nil, zero value otherwise.

### GetGlobalMapTokenOk

`func (o *ControllerRoleVO) GetGlobalMapTokenOk() (*int32, bool)`

GetGlobalMapTokenOk returns a tuple with the GlobalMapToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalMapToken

`func (o *ControllerRoleVO) SetGlobalMapToken(v int32)`

SetGlobalMapToken sets GlobalMapToken field to given value.

### HasGlobalMapToken

`func (o *ControllerRoleVO) HasGlobalMapToken() bool`

HasGlobalMapToken returns a boolean if a field has been set.

### GetGlobalSecurity

`func (o *ControllerRoleVO) GetGlobalSecurity() int32`

GetGlobalSecurity returns the GlobalSecurity field if non-nil, zero value otherwise.

### GetGlobalSecurityOk

`func (o *ControllerRoleVO) GetGlobalSecurityOk() (*int32, bool)`

GetGlobalSecurityOk returns a tuple with the GlobalSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSecurity

`func (o *ControllerRoleVO) SetGlobalSecurity(v int32)`

SetGlobalSecurity sets GlobalSecurity field to given value.

### HasGlobalSecurity

`func (o *ControllerRoleVO) HasGlobalSecurity() bool`

HasGlobalSecurity returns a boolean if a field has been set.

### GetGlobalSetting

`func (o *ControllerRoleVO) GetGlobalSetting() int32`

GetGlobalSetting returns the GlobalSetting field if non-nil, zero value otherwise.

### GetGlobalSettingOk

`func (o *ControllerRoleVO) GetGlobalSettingOk() (*int32, bool)`

GetGlobalSettingOk returns a tuple with the GlobalSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSetting

`func (o *ControllerRoleVO) SetGlobalSetting(v int32)`

SetGlobalSetting sets GlobalSetting field to given value.

### HasGlobalSetting

`func (o *ControllerRoleVO) HasGlobalSetting() bool`

HasGlobalSetting returns a boolean if a field has been set.

### GetGlobalWebhook

`func (o *ControllerRoleVO) GetGlobalWebhook() int32`

GetGlobalWebhook returns the GlobalWebhook field if non-nil, zero value otherwise.

### GetGlobalWebhookOk

`func (o *ControllerRoleVO) GetGlobalWebhookOk() (*int32, bool)`

GetGlobalWebhookOk returns a tuple with the GlobalWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalWebhook

`func (o *ControllerRoleVO) SetGlobalWebhook(v int32)`

SetGlobalWebhook sets GlobalWebhook field to given value.

### HasGlobalWebhook

`func (o *ControllerRoleVO) HasGlobalWebhook() bool`

HasGlobalWebhook returns a boolean if a field has been set.

### GetHotspot

`func (o *ControllerRoleVO) GetHotspot() int32`

GetHotspot returns the Hotspot field if non-nil, zero value otherwise.

### GetHotspotOk

`func (o *ControllerRoleVO) GetHotspotOk() (*int32, bool)`

GetHotspotOk returns a tuple with the Hotspot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotspot

`func (o *ControllerRoleVO) SetHotspot(v int32)`

SetHotspot sets Hotspot field to given value.

### HasHotspot

`func (o *ControllerRoleVO) HasHotspot() bool`

HasHotspot returns a boolean if a field has been set.

### GetInsight

`func (o *ControllerRoleVO) GetInsight() int32`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *ControllerRoleVO) GetInsightOk() (*int32, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *ControllerRoleVO) SetInsight(v int32)`

SetInsight sets Insight field to given value.

### HasInsight

`func (o *ControllerRoleVO) HasInsight() bool`

HasInsight returns a boolean if a field has been set.

### GetLicense

`func (o *ControllerRoleVO) GetLicense() int32`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ControllerRoleVO) GetLicenseOk() (*int32, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ControllerRoleVO) SetLicense(v int32)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ControllerRoleVO) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetLicenseBind

`func (o *ControllerRoleVO) GetLicenseBind() int32`

GetLicenseBind returns the LicenseBind field if non-nil, zero value otherwise.

### GetLicenseBindOk

`func (o *ControllerRoleVO) GetLicenseBindOk() (*int32, bool)`

GetLicenseBindOk returns a tuple with the LicenseBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseBind

`func (o *ControllerRoleVO) SetLicenseBind(v int32)`

SetLicenseBind sets LicenseBind field to given value.

### HasLicenseBind

`func (o *ControllerRoleVO) HasLicenseBind() bool`

HasLicenseBind returns a boolean if a field has been set.

### GetLog

`func (o *ControllerRoleVO) GetLog() int32`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *ControllerRoleVO) GetLogOk() (*int32, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *ControllerRoleVO) SetLog(v int32)`

SetLog sets Log field to given value.

### HasLog

`func (o *ControllerRoleVO) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetManualUpgrade

`func (o *ControllerRoleVO) GetManualUpgrade() int32`

GetManualUpgrade returns the ManualUpgrade field if non-nil, zero value otherwise.

### GetManualUpgradeOk

`func (o *ControllerRoleVO) GetManualUpgradeOk() (*int32, bool)`

GetManualUpgradeOk returns a tuple with the ManualUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualUpgrade

`func (o *ControllerRoleVO) SetManualUpgrade(v int32)`

SetManualUpgrade sets ManualUpgrade field to given value.

### HasManualUpgrade

`func (o *ControllerRoleVO) HasManualUpgrade() bool`

HasManualUpgrade returns a boolean if a field has been set.

### GetMap

`func (o *ControllerRoleVO) GetMap() int32`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *ControllerRoleVO) GetMapOk() (*int32, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *ControllerRoleVO) SetMap(v int32)`

SetMap sets Map field to given value.

### HasMap

`func (o *ControllerRoleVO) HasMap() bool`

HasMap returns a boolean if a field has been set.

### GetNetwork

`func (o *ControllerRoleVO) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ControllerRoleVO) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ControllerRoleVO) SetNetwork(v int32)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ControllerRoleVO) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReport

`func (o *ControllerRoleVO) GetReport() int32`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *ControllerRoleVO) GetReportOk() (*int32, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *ControllerRoleVO) SetReport(v int32)`

SetReport sets Report field to given value.

### HasReport

`func (o *ControllerRoleVO) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetRoles

`func (o *ControllerRoleVO) GetRoles() int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ControllerRoleVO) GetRolesOk() (*int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ControllerRoleVO) SetRoles(v int32)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ControllerRoleVO) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSamlRoles

`func (o *ControllerRoleVO) GetSamlRoles() int32`

GetSamlRoles returns the SamlRoles field if non-nil, zero value otherwise.

### GetSamlRolesOk

`func (o *ControllerRoleVO) GetSamlRolesOk() (*int32, bool)`

GetSamlRolesOk returns a tuple with the SamlRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlRoles

`func (o *ControllerRoleVO) SetSamlRoles(v int32)`

SetSamlRoles sets SamlRoles field to given value.

### HasSamlRoles

`func (o *ControllerRoleVO) HasSamlRoles() bool`

HasSamlRoles returns a boolean if a field has been set.

### GetSamlSsos

`func (o *ControllerRoleVO) GetSamlSsos() int32`

GetSamlSsos returns the SamlSsos field if non-nil, zero value otherwise.

### GetSamlSsosOk

`func (o *ControllerRoleVO) GetSamlSsosOk() (*int32, bool)`

GetSamlSsosOk returns a tuple with the SamlSsos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSsos

`func (o *ControllerRoleVO) SetSamlSsos(v int32)`

SetSamlSsos sets SamlSsos field to given value.

### HasSamlSsos

`func (o *ControllerRoleVO) HasSamlSsos() bool`

HasSamlSsos returns a boolean if a field has been set.

### GetSamlUsers

`func (o *ControllerRoleVO) GetSamlUsers() int32`

GetSamlUsers returns the SamlUsers field if non-nil, zero value otherwise.

### GetSamlUsersOk

`func (o *ControllerRoleVO) GetSamlUsersOk() (*int32, bool)`

GetSamlUsersOk returns a tuple with the SamlUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlUsers

`func (o *ControllerRoleVO) SetSamlUsers(v int32)`

SetSamlUsers sets SamlUsers field to given value.

### HasSamlUsers

`func (o *ControllerRoleVO) HasSamlUsers() bool`

HasSamlUsers returns a boolean if a field has been set.

### GetSdWan

`func (o *ControllerRoleVO) GetSdWan() int32`

GetSdWan returns the SdWan field if non-nil, zero value otherwise.

### GetSdWanOk

`func (o *ControllerRoleVO) GetSdWanOk() (*int32, bool)`

GetSdWanOk returns a tuple with the SdWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdWan

`func (o *ControllerRoleVO) SetSdWan(v int32)`

SetSdWan sets SdWan field to given value.

### HasSdWan

`func (o *ControllerRoleVO) HasSdWan() bool`

HasSdWan returns a boolean if a field has been set.

### GetSiteAnalyze

`func (o *ControllerRoleVO) GetSiteAnalyze() int32`

GetSiteAnalyze returns the SiteAnalyze field if non-nil, zero value otherwise.

### GetSiteAnalyzeOk

`func (o *ControllerRoleVO) GetSiteAnalyzeOk() (*int32, bool)`

GetSiteAnalyzeOk returns a tuple with the SiteAnalyze field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteAnalyze

`func (o *ControllerRoleVO) SetSiteAnalyze(v int32)`

SetSiteAnalyze sets SiteAnalyze field to given value.

### HasSiteAnalyze

`func (o *ControllerRoleVO) HasSiteAnalyze() bool`

HasSiteAnalyze returns a boolean if a field has been set.

### GetSiteDevice

`func (o *ControllerRoleVO) GetSiteDevice() int32`

GetSiteDevice returns the SiteDevice field if non-nil, zero value otherwise.

### GetSiteDeviceOk

`func (o *ControllerRoleVO) GetSiteDeviceOk() (*int32, bool)`

GetSiteDeviceOk returns a tuple with the SiteDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteDevice

`func (o *ControllerRoleVO) SetSiteDevice(v int32)`

SetSiteDevice sets SiteDevice field to given value.

### HasSiteDevice

`func (o *ControllerRoleVO) HasSiteDevice() bool`

HasSiteDevice returns a boolean if a field has been set.

### GetSiteHome

`func (o *ControllerRoleVO) GetSiteHome() int32`

GetSiteHome returns the SiteHome field if non-nil, zero value otherwise.

### GetSiteHomeOk

`func (o *ControllerRoleVO) GetSiteHomeOk() (*int32, bool)`

GetSiteHomeOk returns a tuple with the SiteHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteHome

`func (o *ControllerRoleVO) SetSiteHome(v int32)`

SetSiteHome sets SiteHome field to given value.

### HasSiteHome

`func (o *ControllerRoleVO) HasSiteHome() bool`

HasSiteHome returns a boolean if a field has been set.

### GetSiteMaintain

`func (o *ControllerRoleVO) GetSiteMaintain() int32`

GetSiteMaintain returns the SiteMaintain field if non-nil, zero value otherwise.

### GetSiteMaintainOk

`func (o *ControllerRoleVO) GetSiteMaintainOk() (*int32, bool)`

GetSiteMaintainOk returns a tuple with the SiteMaintain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteMaintain

`func (o *ControllerRoleVO) SetSiteMaintain(v int32)`

SetSiteMaintain sets SiteMaintain field to given value.

### HasSiteMaintain

`func (o *ControllerRoleVO) HasSiteMaintain() bool`

HasSiteMaintain returns a boolean if a field has been set.

### GetSiteNetwork

`func (o *ControllerRoleVO) GetSiteNetwork() int32`

GetSiteNetwork returns the SiteNetwork field if non-nil, zero value otherwise.

### GetSiteNetworkOk

`func (o *ControllerRoleVO) GetSiteNetworkOk() (*int32, bool)`

GetSiteNetworkOk returns a tuple with the SiteNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNetwork

`func (o *ControllerRoleVO) SetSiteNetwork(v int32)`

SetSiteNetwork sets SiteNetwork field to given value.

### HasSiteNetwork

`func (o *ControllerRoleVO) HasSiteNetwork() bool`

HasSiteNetwork returns a boolean if a field has been set.

### GetSiteTemplate

`func (o *ControllerRoleVO) GetSiteTemplate() int32`

GetSiteTemplate returns the SiteTemplate field if non-nil, zero value otherwise.

### GetSiteTemplateOk

`func (o *ControllerRoleVO) GetSiteTemplateOk() (*int32, bool)`

GetSiteTemplateOk returns a tuple with the SiteTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplate

`func (o *ControllerRoleVO) SetSiteTemplate(v int32)`

SetSiteTemplate sets SiteTemplate field to given value.

### HasSiteTemplate

`func (o *ControllerRoleVO) HasSiteTemplate() bool`

HasSiteTemplate returns a boolean if a field has been set.

### GetStatics

`func (o *ControllerRoleVO) GetStatics() int32`

GetStatics returns the Statics field if non-nil, zero value otherwise.

### GetStaticsOk

`func (o *ControllerRoleVO) GetStaticsOk() (*int32, bool)`

GetStaticsOk returns a tuple with the Statics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatics

`func (o *ControllerRoleVO) SetStatics(v int32)`

SetStatics sets Statics field to given value.

### HasStatics

`func (o *ControllerRoleVO) HasStatics() bool`

HasStatics returns a boolean if a field has been set.

### GetUsers

`func (o *ControllerRoleVO) GetUsers() int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ControllerRoleVO) GetUsersOk() (*int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ControllerRoleVO) SetUsers(v int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ControllerRoleVO) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


