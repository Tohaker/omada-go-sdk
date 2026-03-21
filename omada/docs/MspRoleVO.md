# MspRoleVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MspAddAdoptDevice** | Pointer to **int32** | Msp add and adopt devices. It should be a value as follows: 0:block; 2:access | [optional] 
**MspAddDevices** | Pointer to **int32** | Msp add devices. It should be a value as follows: 0:block; 2:access | [optional] 
**MspAdopt** | Pointer to **int32** | Msp adopt. It should be a value as follows: 0:block; 2:access | [optional] 
**MspCluster** | Pointer to **int32** | msp cluster, 0:block; 1:view only; 2:modify | [optional] 
**MspDashboard** | Pointer to **int32** | Msp dashboard permission. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**MspDevice** | Pointer to **int32** | Msp device. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**MspExportData** | Pointer to **int32** | Export data in msp view. It should be a value as follows: 0:block; 2:access | [optional] 
**MspLicense** | Pointer to **int32** | Msp license. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**MspLicenseBind** | Pointer to **int32** | msp license bind | [optional] 
**MspLog** | Pointer to **int32** | Msp log. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**MspRoles** | Pointer to **int32** | Msp roles. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**MspSamlRoles** | Pointer to **int32** | Saml roles in msp view. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**MspSamlSsos** | Pointer to **int32** | Saml ssos in msp view. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**MspSamlUsers** | Pointer to **int32** | Saml users in msp view. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**MspSetting** | Pointer to **int32** | Msp setting. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**MspSites** | Pointer to **int32** | Msp sites. It should be a value as follows: 0:block; 2:access | [optional] 
**MspUsers** | Pointer to **int32** | Msp users. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 
**MspWebhook** | Pointer to **int32** | Msp webhook. It should be a value as follows: 0:block; 1:view only; 2:modify | [optional] 

## Methods

### NewMspRoleVO

`func NewMspRoleVO() *MspRoleVO`

NewMspRoleVO instantiates a new MspRoleVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspRoleVOWithDefaults

`func NewMspRoleVOWithDefaults() *MspRoleVO`

NewMspRoleVOWithDefaults instantiates a new MspRoleVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMspAddAdoptDevice

`func (o *MspRoleVO) GetMspAddAdoptDevice() int32`

GetMspAddAdoptDevice returns the MspAddAdoptDevice field if non-nil, zero value otherwise.

### GetMspAddAdoptDeviceOk

`func (o *MspRoleVO) GetMspAddAdoptDeviceOk() (*int32, bool)`

GetMspAddAdoptDeviceOk returns a tuple with the MspAddAdoptDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspAddAdoptDevice

`func (o *MspRoleVO) SetMspAddAdoptDevice(v int32)`

SetMspAddAdoptDevice sets MspAddAdoptDevice field to given value.

### HasMspAddAdoptDevice

`func (o *MspRoleVO) HasMspAddAdoptDevice() bool`

HasMspAddAdoptDevice returns a boolean if a field has been set.

### GetMspAddDevices

`func (o *MspRoleVO) GetMspAddDevices() int32`

GetMspAddDevices returns the MspAddDevices field if non-nil, zero value otherwise.

### GetMspAddDevicesOk

`func (o *MspRoleVO) GetMspAddDevicesOk() (*int32, bool)`

GetMspAddDevicesOk returns a tuple with the MspAddDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspAddDevices

`func (o *MspRoleVO) SetMspAddDevices(v int32)`

SetMspAddDevices sets MspAddDevices field to given value.

### HasMspAddDevices

`func (o *MspRoleVO) HasMspAddDevices() bool`

HasMspAddDevices returns a boolean if a field has been set.

### GetMspAdopt

`func (o *MspRoleVO) GetMspAdopt() int32`

GetMspAdopt returns the MspAdopt field if non-nil, zero value otherwise.

### GetMspAdoptOk

`func (o *MspRoleVO) GetMspAdoptOk() (*int32, bool)`

GetMspAdoptOk returns a tuple with the MspAdopt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspAdopt

`func (o *MspRoleVO) SetMspAdopt(v int32)`

SetMspAdopt sets MspAdopt field to given value.

### HasMspAdopt

`func (o *MspRoleVO) HasMspAdopt() bool`

HasMspAdopt returns a boolean if a field has been set.

### GetMspCluster

`func (o *MspRoleVO) GetMspCluster() int32`

GetMspCluster returns the MspCluster field if non-nil, zero value otherwise.

### GetMspClusterOk

`func (o *MspRoleVO) GetMspClusterOk() (*int32, bool)`

GetMspClusterOk returns a tuple with the MspCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspCluster

`func (o *MspRoleVO) SetMspCluster(v int32)`

SetMspCluster sets MspCluster field to given value.

### HasMspCluster

`func (o *MspRoleVO) HasMspCluster() bool`

HasMspCluster returns a boolean if a field has been set.

### GetMspDashboard

`func (o *MspRoleVO) GetMspDashboard() int32`

GetMspDashboard returns the MspDashboard field if non-nil, zero value otherwise.

### GetMspDashboardOk

`func (o *MspRoleVO) GetMspDashboardOk() (*int32, bool)`

GetMspDashboardOk returns a tuple with the MspDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspDashboard

`func (o *MspRoleVO) SetMspDashboard(v int32)`

SetMspDashboard sets MspDashboard field to given value.

### HasMspDashboard

`func (o *MspRoleVO) HasMspDashboard() bool`

HasMspDashboard returns a boolean if a field has been set.

### GetMspDevice

`func (o *MspRoleVO) GetMspDevice() int32`

GetMspDevice returns the MspDevice field if non-nil, zero value otherwise.

### GetMspDeviceOk

`func (o *MspRoleVO) GetMspDeviceOk() (*int32, bool)`

GetMspDeviceOk returns a tuple with the MspDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspDevice

`func (o *MspRoleVO) SetMspDevice(v int32)`

SetMspDevice sets MspDevice field to given value.

### HasMspDevice

`func (o *MspRoleVO) HasMspDevice() bool`

HasMspDevice returns a boolean if a field has been set.

### GetMspExportData

`func (o *MspRoleVO) GetMspExportData() int32`

GetMspExportData returns the MspExportData field if non-nil, zero value otherwise.

### GetMspExportDataOk

`func (o *MspRoleVO) GetMspExportDataOk() (*int32, bool)`

GetMspExportDataOk returns a tuple with the MspExportData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspExportData

`func (o *MspRoleVO) SetMspExportData(v int32)`

SetMspExportData sets MspExportData field to given value.

### HasMspExportData

`func (o *MspRoleVO) HasMspExportData() bool`

HasMspExportData returns a boolean if a field has been set.

### GetMspLicense

`func (o *MspRoleVO) GetMspLicense() int32`

GetMspLicense returns the MspLicense field if non-nil, zero value otherwise.

### GetMspLicenseOk

`func (o *MspRoleVO) GetMspLicenseOk() (*int32, bool)`

GetMspLicenseOk returns a tuple with the MspLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspLicense

`func (o *MspRoleVO) SetMspLicense(v int32)`

SetMspLicense sets MspLicense field to given value.

### HasMspLicense

`func (o *MspRoleVO) HasMspLicense() bool`

HasMspLicense returns a boolean if a field has been set.

### GetMspLicenseBind

`func (o *MspRoleVO) GetMspLicenseBind() int32`

GetMspLicenseBind returns the MspLicenseBind field if non-nil, zero value otherwise.

### GetMspLicenseBindOk

`func (o *MspRoleVO) GetMspLicenseBindOk() (*int32, bool)`

GetMspLicenseBindOk returns a tuple with the MspLicenseBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspLicenseBind

`func (o *MspRoleVO) SetMspLicenseBind(v int32)`

SetMspLicenseBind sets MspLicenseBind field to given value.

### HasMspLicenseBind

`func (o *MspRoleVO) HasMspLicenseBind() bool`

HasMspLicenseBind returns a boolean if a field has been set.

### GetMspLog

`func (o *MspRoleVO) GetMspLog() int32`

GetMspLog returns the MspLog field if non-nil, zero value otherwise.

### GetMspLogOk

`func (o *MspRoleVO) GetMspLogOk() (*int32, bool)`

GetMspLogOk returns a tuple with the MspLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspLog

`func (o *MspRoleVO) SetMspLog(v int32)`

SetMspLog sets MspLog field to given value.

### HasMspLog

`func (o *MspRoleVO) HasMspLog() bool`

HasMspLog returns a boolean if a field has been set.

### GetMspRoles

`func (o *MspRoleVO) GetMspRoles() int32`

GetMspRoles returns the MspRoles field if non-nil, zero value otherwise.

### GetMspRolesOk

`func (o *MspRoleVO) GetMspRolesOk() (*int32, bool)`

GetMspRolesOk returns a tuple with the MspRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspRoles

`func (o *MspRoleVO) SetMspRoles(v int32)`

SetMspRoles sets MspRoles field to given value.

### HasMspRoles

`func (o *MspRoleVO) HasMspRoles() bool`

HasMspRoles returns a boolean if a field has been set.

### GetMspSamlRoles

`func (o *MspRoleVO) GetMspSamlRoles() int32`

GetMspSamlRoles returns the MspSamlRoles field if non-nil, zero value otherwise.

### GetMspSamlRolesOk

`func (o *MspRoleVO) GetMspSamlRolesOk() (*int32, bool)`

GetMspSamlRolesOk returns a tuple with the MspSamlRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspSamlRoles

`func (o *MspRoleVO) SetMspSamlRoles(v int32)`

SetMspSamlRoles sets MspSamlRoles field to given value.

### HasMspSamlRoles

`func (o *MspRoleVO) HasMspSamlRoles() bool`

HasMspSamlRoles returns a boolean if a field has been set.

### GetMspSamlSsos

`func (o *MspRoleVO) GetMspSamlSsos() int32`

GetMspSamlSsos returns the MspSamlSsos field if non-nil, zero value otherwise.

### GetMspSamlSsosOk

`func (o *MspRoleVO) GetMspSamlSsosOk() (*int32, bool)`

GetMspSamlSsosOk returns a tuple with the MspSamlSsos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspSamlSsos

`func (o *MspRoleVO) SetMspSamlSsos(v int32)`

SetMspSamlSsos sets MspSamlSsos field to given value.

### HasMspSamlSsos

`func (o *MspRoleVO) HasMspSamlSsos() bool`

HasMspSamlSsos returns a boolean if a field has been set.

### GetMspSamlUsers

`func (o *MspRoleVO) GetMspSamlUsers() int32`

GetMspSamlUsers returns the MspSamlUsers field if non-nil, zero value otherwise.

### GetMspSamlUsersOk

`func (o *MspRoleVO) GetMspSamlUsersOk() (*int32, bool)`

GetMspSamlUsersOk returns a tuple with the MspSamlUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspSamlUsers

`func (o *MspRoleVO) SetMspSamlUsers(v int32)`

SetMspSamlUsers sets MspSamlUsers field to given value.

### HasMspSamlUsers

`func (o *MspRoleVO) HasMspSamlUsers() bool`

HasMspSamlUsers returns a boolean if a field has been set.

### GetMspSetting

`func (o *MspRoleVO) GetMspSetting() int32`

GetMspSetting returns the MspSetting field if non-nil, zero value otherwise.

### GetMspSettingOk

`func (o *MspRoleVO) GetMspSettingOk() (*int32, bool)`

GetMspSettingOk returns a tuple with the MspSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspSetting

`func (o *MspRoleVO) SetMspSetting(v int32)`

SetMspSetting sets MspSetting field to given value.

### HasMspSetting

`func (o *MspRoleVO) HasMspSetting() bool`

HasMspSetting returns a boolean if a field has been set.

### GetMspSites

`func (o *MspRoleVO) GetMspSites() int32`

GetMspSites returns the MspSites field if non-nil, zero value otherwise.

### GetMspSitesOk

`func (o *MspRoleVO) GetMspSitesOk() (*int32, bool)`

GetMspSitesOk returns a tuple with the MspSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspSites

`func (o *MspRoleVO) SetMspSites(v int32)`

SetMspSites sets MspSites field to given value.

### HasMspSites

`func (o *MspRoleVO) HasMspSites() bool`

HasMspSites returns a boolean if a field has been set.

### GetMspUsers

`func (o *MspRoleVO) GetMspUsers() int32`

GetMspUsers returns the MspUsers field if non-nil, zero value otherwise.

### GetMspUsersOk

`func (o *MspRoleVO) GetMspUsersOk() (*int32, bool)`

GetMspUsersOk returns a tuple with the MspUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspUsers

`func (o *MspRoleVO) SetMspUsers(v int32)`

SetMspUsers sets MspUsers field to given value.

### HasMspUsers

`func (o *MspRoleVO) HasMspUsers() bool`

HasMspUsers returns a boolean if a field has been set.

### GetMspWebhook

`func (o *MspRoleVO) GetMspWebhook() int32`

GetMspWebhook returns the MspWebhook field if non-nil, zero value otherwise.

### GetMspWebhookOk

`func (o *MspRoleVO) GetMspWebhookOk() (*int32, bool)`

GetMspWebhookOk returns a tuple with the MspWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspWebhook

`func (o *MspRoleVO) SetMspWebhook(v int32)`

SetMspWebhook sets MspWebhook field to given value.

### HasMspWebhook

`func (o *MspRoleVO) HasMspWebhook() bool`

HasMspWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


