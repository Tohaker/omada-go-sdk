# VoipDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedInAdvanced** | **bool** | Whether the device is added in advanced. | 
**Id** | Pointer to **string** | The ID of voip device. | [optional] 
**Ip** | Pointer to **string** | The IP of voip device. | [optional] 
**Mac** | Pointer to **string** | The mac of voip device. | [optional] 
**Model** | Pointer to **string** | The model of voip device. | [optional] 
**ModelVersion** | Pointer to **string** | The model version of voip device. | [optional] 
**Name** | Pointer to **string** | The name of voip device. | [optional] 
**Status** | Pointer to **int32** | The status of voip device. | [optional] 
**StatusCategory** | Pointer to **int32** | The status category of voip device. | [optional] 
**TelephoneNumberAdvancedSetting** | Pointer to **map[string]interface{}** | The telephone number advanced setting of voip device. | [optional] 
**Type** | Pointer to **string** | The type of voip device. | [optional] 
**VoipDeviceApConfiguration** | Pointer to [**VoipDeviceApConfigurationOpenApiVO**](VoipDeviceApConfigurationOpenApiVO.md) |  | [optional] 
**VoipDeviceOsgConfiguration** | Pointer to [**VoipDeviceOsgConfigurationOpenApiVO**](VoipDeviceOsgConfigurationOpenApiVO.md) |  | [optional] 
**VoipVlanIp** | Pointer to **string** | IP address of VoIP VLAN. | [optional] 
**WirelessLinked** | **bool** | Whether to enable wireless linked. | 

## Methods

### NewVoipDevice

`func NewVoipDevice(addedInAdvanced bool, wirelessLinked bool, ) *VoipDevice`

NewVoipDevice instantiates a new VoipDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoipDeviceWithDefaults

`func NewVoipDeviceWithDefaults() *VoipDevice`

NewVoipDeviceWithDefaults instantiates a new VoipDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedInAdvanced

`func (o *VoipDevice) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *VoipDevice) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *VoipDevice) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.


### GetId

`func (o *VoipDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoipDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoipDevice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoipDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *VoipDevice) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *VoipDevice) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *VoipDevice) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *VoipDevice) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *VoipDevice) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *VoipDevice) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *VoipDevice) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *VoipDevice) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *VoipDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VoipDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VoipDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VoipDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *VoipDevice) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *VoipDevice) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *VoipDevice) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *VoipDevice) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *VoipDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoipDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoipDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VoipDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *VoipDevice) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VoipDevice) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VoipDevice) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VoipDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *VoipDevice) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *VoipDevice) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *VoipDevice) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *VoipDevice) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetTelephoneNumberAdvancedSetting

`func (o *VoipDevice) GetTelephoneNumberAdvancedSetting() map[string]interface{}`

GetTelephoneNumberAdvancedSetting returns the TelephoneNumberAdvancedSetting field if non-nil, zero value otherwise.

### GetTelephoneNumberAdvancedSettingOk

`func (o *VoipDevice) GetTelephoneNumberAdvancedSettingOk() (*map[string]interface{}, bool)`

GetTelephoneNumberAdvancedSettingOk returns a tuple with the TelephoneNumberAdvancedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumberAdvancedSetting

`func (o *VoipDevice) SetTelephoneNumberAdvancedSetting(v map[string]interface{})`

SetTelephoneNumberAdvancedSetting sets TelephoneNumberAdvancedSetting field to given value.

### HasTelephoneNumberAdvancedSetting

`func (o *VoipDevice) HasTelephoneNumberAdvancedSetting() bool`

HasTelephoneNumberAdvancedSetting returns a boolean if a field has been set.

### GetType

`func (o *VoipDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VoipDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VoipDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VoipDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVoipDeviceApConfiguration

`func (o *VoipDevice) GetVoipDeviceApConfiguration() VoipDeviceApConfigurationOpenApiVO`

GetVoipDeviceApConfiguration returns the VoipDeviceApConfiguration field if non-nil, zero value otherwise.

### GetVoipDeviceApConfigurationOk

`func (o *VoipDevice) GetVoipDeviceApConfigurationOk() (*VoipDeviceApConfigurationOpenApiVO, bool)`

GetVoipDeviceApConfigurationOk returns a tuple with the VoipDeviceApConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipDeviceApConfiguration

`func (o *VoipDevice) SetVoipDeviceApConfiguration(v VoipDeviceApConfigurationOpenApiVO)`

SetVoipDeviceApConfiguration sets VoipDeviceApConfiguration field to given value.

### HasVoipDeviceApConfiguration

`func (o *VoipDevice) HasVoipDeviceApConfiguration() bool`

HasVoipDeviceApConfiguration returns a boolean if a field has been set.

### GetVoipDeviceOsgConfiguration

`func (o *VoipDevice) GetVoipDeviceOsgConfiguration() VoipDeviceOsgConfigurationOpenApiVO`

GetVoipDeviceOsgConfiguration returns the VoipDeviceOsgConfiguration field if non-nil, zero value otherwise.

### GetVoipDeviceOsgConfigurationOk

`func (o *VoipDevice) GetVoipDeviceOsgConfigurationOk() (*VoipDeviceOsgConfigurationOpenApiVO, bool)`

GetVoipDeviceOsgConfigurationOk returns a tuple with the VoipDeviceOsgConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipDeviceOsgConfiguration

`func (o *VoipDevice) SetVoipDeviceOsgConfiguration(v VoipDeviceOsgConfigurationOpenApiVO)`

SetVoipDeviceOsgConfiguration sets VoipDeviceOsgConfiguration field to given value.

### HasVoipDeviceOsgConfiguration

`func (o *VoipDevice) HasVoipDeviceOsgConfiguration() bool`

HasVoipDeviceOsgConfiguration returns a boolean if a field has been set.

### GetVoipVlanIp

`func (o *VoipDevice) GetVoipVlanIp() string`

GetVoipVlanIp returns the VoipVlanIp field if non-nil, zero value otherwise.

### GetVoipVlanIpOk

`func (o *VoipDevice) GetVoipVlanIpOk() (*string, bool)`

GetVoipVlanIpOk returns a tuple with the VoipVlanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipVlanIp

`func (o *VoipDevice) SetVoipVlanIp(v string)`

SetVoipVlanIp sets VoipVlanIp field to given value.

### HasVoipVlanIp

`func (o *VoipDevice) HasVoipVlanIp() bool`

HasVoipVlanIp returns a boolean if a field has been set.

### GetWirelessLinked

`func (o *VoipDevice) GetWirelessLinked() bool`

GetWirelessLinked returns the WirelessLinked field if non-nil, zero value otherwise.

### GetWirelessLinkedOk

`func (o *VoipDevice) GetWirelessLinkedOk() (*bool, bool)`

GetWirelessLinkedOk returns a tuple with the WirelessLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessLinked

`func (o *VoipDevice) SetWirelessLinked(v bool)`

SetWirelessLinked sets WirelessLinked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


