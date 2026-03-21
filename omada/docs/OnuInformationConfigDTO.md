# OnuInformationConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveStatus** | Pointer to **string** | Active status should be a value as follows:ACTIVE: ONU is in an active state, capable of data communication and carrying services. INACTIVE: ONU is in an inactive state, possibly due to unactivated configuration, being offline, configuration failure, or mismatched configuration. | [optional] 
**AdminStatus** | Pointer to **string** | Admin status should be a value as follows:ACTIVATE,DEACTIVATE | [optional] 
**ConfigStatus** | Pointer to **string** | Config status should be a value as follows:SUCCESS: Configuration successfully delivered and recognized by the ONU.FAILED: Configuration not successfully delivered or not recognized by the ONU. | [optional] 
**EquipmentId** | Pointer to **string** | Matched ONU device ID, equipmentId should contain 1 to 20 characters in ASCII code from \\x21 to \\x7e. | [optional] 
**HardwareVersion** | Pointer to **string** | Hardware version of the ONU, hardwareVersion should contain 1 to 20 characters, including uppercase and lowercase letters, numbers, and the symbols -@_:/.  | [optional] 
**Key** | Pointer to **string** | Identifier of ONU | [optional] 
**LineProfile** | Pointer to **string** | Line profile.Display in the format: Line_Profile_Name + Line_Profile_ID. | [optional] 
**MacAddress** | Pointer to **string** | Mac address of ONU | [optional] 
**MatchStatus** | Pointer to **string** | Match status should be a value as follows:MATCH: ONU hardware capabilities are consistent with the bound service template.MISMATCH: ONU hardware capabilities are inconsistent with the bound service template. | [optional] 
**OnlineStatus** | Pointer to **string** | Online status should be a value as follows:ONLINE,OFFLINE | [optional] 
**OnuDescription** | Pointer to **string** | Description of ONU.OnuDescription should be 1-32 characters, including letters, numbers, and symbols (-@_:/.). | [optional] 
**OnuId** | Pointer to **int32** | Onu Id should be within the range of 0 to 127 | [optional] 
**PonPortId** | Pointer to **int32** | PON port ID of the registered and online ONT | [optional] 
**PonPortStr** | Pointer to **string** | String form of pon port.e.g.,PON 1/2/1 | [optional] 
**ReceivedOpticalPower** | Pointer to **string** | ONU&#39;s received power, in dBm. | [optional] 
**SerialNumber** | Pointer to **string** | SerialNumber should contain 12, 13, or 16 characters, formatted as &#x60;XXXXXXXXXXXX&#x60;,&#39;XXXX-XXXXXXXX&#39;,&#39;XXXXXXXXXXXXXXXX&#39;. | [optional] 
**ServicePortProfile** | Pointer to **string** | Service port profile.Display in the format: Service_Port_Profile_Name + Service_Port_Profile_ID. | [optional] 
**ServiceProfile** | Pointer to **string** | Service profile.Display in the format: Service_Profile_Name + Service_Profile_ID. | [optional] 
**SoftWareVersion** | Pointer to **string** |  | [optional] 
**SoftwareVersion** | Pointer to [**OnuInformationConfigDTO**](OnuInformationConfigDTO.md) |  | [optional] 
**TransmittedOpticalPower** | Pointer to **string** | ONU&#39;s transmission power, in dBm. | [optional] 

## Methods

### NewOnuInformationConfigDTO

`func NewOnuInformationConfigDTO() *OnuInformationConfigDTO`

NewOnuInformationConfigDTO instantiates a new OnuInformationConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuInformationConfigDTOWithDefaults

`func NewOnuInformationConfigDTOWithDefaults() *OnuInformationConfigDTO`

NewOnuInformationConfigDTOWithDefaults instantiates a new OnuInformationConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveStatus

`func (o *OnuInformationConfigDTO) GetActiveStatus() string`

GetActiveStatus returns the ActiveStatus field if non-nil, zero value otherwise.

### GetActiveStatusOk

`func (o *OnuInformationConfigDTO) GetActiveStatusOk() (*string, bool)`

GetActiveStatusOk returns a tuple with the ActiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStatus

`func (o *OnuInformationConfigDTO) SetActiveStatus(v string)`

SetActiveStatus sets ActiveStatus field to given value.

### HasActiveStatus

`func (o *OnuInformationConfigDTO) HasActiveStatus() bool`

HasActiveStatus returns a boolean if a field has been set.

### GetAdminStatus

`func (o *OnuInformationConfigDTO) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *OnuInformationConfigDTO) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *OnuInformationConfigDTO) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *OnuInformationConfigDTO) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetConfigStatus

`func (o *OnuInformationConfigDTO) GetConfigStatus() string`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *OnuInformationConfigDTO) GetConfigStatusOk() (*string, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *OnuInformationConfigDTO) SetConfigStatus(v string)`

SetConfigStatus sets ConfigStatus field to given value.

### HasConfigStatus

`func (o *OnuInformationConfigDTO) HasConfigStatus() bool`

HasConfigStatus returns a boolean if a field has been set.

### GetEquipmentId

`func (o *OnuInformationConfigDTO) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *OnuInformationConfigDTO) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *OnuInformationConfigDTO) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.

### HasEquipmentId

`func (o *OnuInformationConfigDTO) HasEquipmentId() bool`

HasEquipmentId returns a boolean if a field has been set.

### GetHardwareVersion

`func (o *OnuInformationConfigDTO) GetHardwareVersion() string`

GetHardwareVersion returns the HardwareVersion field if non-nil, zero value otherwise.

### GetHardwareVersionOk

`func (o *OnuInformationConfigDTO) GetHardwareVersionOk() (*string, bool)`

GetHardwareVersionOk returns a tuple with the HardwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareVersion

`func (o *OnuInformationConfigDTO) SetHardwareVersion(v string)`

SetHardwareVersion sets HardwareVersion field to given value.

### HasHardwareVersion

`func (o *OnuInformationConfigDTO) HasHardwareVersion() bool`

HasHardwareVersion returns a boolean if a field has been set.

### GetKey

`func (o *OnuInformationConfigDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OnuInformationConfigDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OnuInformationConfigDTO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OnuInformationConfigDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLineProfile

`func (o *OnuInformationConfigDTO) GetLineProfile() string`

GetLineProfile returns the LineProfile field if non-nil, zero value otherwise.

### GetLineProfileOk

`func (o *OnuInformationConfigDTO) GetLineProfileOk() (*string, bool)`

GetLineProfileOk returns a tuple with the LineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfile

`func (o *OnuInformationConfigDTO) SetLineProfile(v string)`

SetLineProfile sets LineProfile field to given value.

### HasLineProfile

`func (o *OnuInformationConfigDTO) HasLineProfile() bool`

HasLineProfile returns a boolean if a field has been set.

### GetMacAddress

`func (o *OnuInformationConfigDTO) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *OnuInformationConfigDTO) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *OnuInformationConfigDTO) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *OnuInformationConfigDTO) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMatchStatus

`func (o *OnuInformationConfigDTO) GetMatchStatus() string`

GetMatchStatus returns the MatchStatus field if non-nil, zero value otherwise.

### GetMatchStatusOk

`func (o *OnuInformationConfigDTO) GetMatchStatusOk() (*string, bool)`

GetMatchStatusOk returns a tuple with the MatchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchStatus

`func (o *OnuInformationConfigDTO) SetMatchStatus(v string)`

SetMatchStatus sets MatchStatus field to given value.

### HasMatchStatus

`func (o *OnuInformationConfigDTO) HasMatchStatus() bool`

HasMatchStatus returns a boolean if a field has been set.

### GetOnlineStatus

`func (o *OnuInformationConfigDTO) GetOnlineStatus() string`

GetOnlineStatus returns the OnlineStatus field if non-nil, zero value otherwise.

### GetOnlineStatusOk

`func (o *OnuInformationConfigDTO) GetOnlineStatusOk() (*string, bool)`

GetOnlineStatusOk returns a tuple with the OnlineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineStatus

`func (o *OnuInformationConfigDTO) SetOnlineStatus(v string)`

SetOnlineStatus sets OnlineStatus field to given value.

### HasOnlineStatus

`func (o *OnuInformationConfigDTO) HasOnlineStatus() bool`

HasOnlineStatus returns a boolean if a field has been set.

### GetOnuDescription

`func (o *OnuInformationConfigDTO) GetOnuDescription() string`

GetOnuDescription returns the OnuDescription field if non-nil, zero value otherwise.

### GetOnuDescriptionOk

`func (o *OnuInformationConfigDTO) GetOnuDescriptionOk() (*string, bool)`

GetOnuDescriptionOk returns a tuple with the OnuDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuDescription

`func (o *OnuInformationConfigDTO) SetOnuDescription(v string)`

SetOnuDescription sets OnuDescription field to given value.

### HasOnuDescription

`func (o *OnuInformationConfigDTO) HasOnuDescription() bool`

HasOnuDescription returns a boolean if a field has been set.

### GetOnuId

`func (o *OnuInformationConfigDTO) GetOnuId() int32`

GetOnuId returns the OnuId field if non-nil, zero value otherwise.

### GetOnuIdOk

`func (o *OnuInformationConfigDTO) GetOnuIdOk() (*int32, bool)`

GetOnuIdOk returns a tuple with the OnuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuId

`func (o *OnuInformationConfigDTO) SetOnuId(v int32)`

SetOnuId sets OnuId field to given value.

### HasOnuId

`func (o *OnuInformationConfigDTO) HasOnuId() bool`

HasOnuId returns a boolean if a field has been set.

### GetPonPortId

`func (o *OnuInformationConfigDTO) GetPonPortId() int32`

GetPonPortId returns the PonPortId field if non-nil, zero value otherwise.

### GetPonPortIdOk

`func (o *OnuInformationConfigDTO) GetPonPortIdOk() (*int32, bool)`

GetPonPortIdOk returns a tuple with the PonPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPortId

`func (o *OnuInformationConfigDTO) SetPonPortId(v int32)`

SetPonPortId sets PonPortId field to given value.

### HasPonPortId

`func (o *OnuInformationConfigDTO) HasPonPortId() bool`

HasPonPortId returns a boolean if a field has been set.

### GetPonPortStr

`func (o *OnuInformationConfigDTO) GetPonPortStr() string`

GetPonPortStr returns the PonPortStr field if non-nil, zero value otherwise.

### GetPonPortStrOk

`func (o *OnuInformationConfigDTO) GetPonPortStrOk() (*string, bool)`

GetPonPortStrOk returns a tuple with the PonPortStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPortStr

`func (o *OnuInformationConfigDTO) SetPonPortStr(v string)`

SetPonPortStr sets PonPortStr field to given value.

### HasPonPortStr

`func (o *OnuInformationConfigDTO) HasPonPortStr() bool`

HasPonPortStr returns a boolean if a field has been set.

### GetReceivedOpticalPower

`func (o *OnuInformationConfigDTO) GetReceivedOpticalPower() string`

GetReceivedOpticalPower returns the ReceivedOpticalPower field if non-nil, zero value otherwise.

### GetReceivedOpticalPowerOk

`func (o *OnuInformationConfigDTO) GetReceivedOpticalPowerOk() (*string, bool)`

GetReceivedOpticalPowerOk returns a tuple with the ReceivedOpticalPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedOpticalPower

`func (o *OnuInformationConfigDTO) SetReceivedOpticalPower(v string)`

SetReceivedOpticalPower sets ReceivedOpticalPower field to given value.

### HasReceivedOpticalPower

`func (o *OnuInformationConfigDTO) HasReceivedOpticalPower() bool`

HasReceivedOpticalPower returns a boolean if a field has been set.

### GetSerialNumber

`func (o *OnuInformationConfigDTO) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *OnuInformationConfigDTO) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *OnuInformationConfigDTO) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *OnuInformationConfigDTO) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetServicePortProfile

`func (o *OnuInformationConfigDTO) GetServicePortProfile() string`

GetServicePortProfile returns the ServicePortProfile field if non-nil, zero value otherwise.

### GetServicePortProfileOk

`func (o *OnuInformationConfigDTO) GetServicePortProfileOk() (*string, bool)`

GetServicePortProfileOk returns a tuple with the ServicePortProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortProfile

`func (o *OnuInformationConfigDTO) SetServicePortProfile(v string)`

SetServicePortProfile sets ServicePortProfile field to given value.

### HasServicePortProfile

`func (o *OnuInformationConfigDTO) HasServicePortProfile() bool`

HasServicePortProfile returns a boolean if a field has been set.

### GetServiceProfile

`func (o *OnuInformationConfigDTO) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *OnuInformationConfigDTO) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *OnuInformationConfigDTO) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *OnuInformationConfigDTO) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.

### GetSoftWareVersion

`func (o *OnuInformationConfigDTO) GetSoftWareVersion() string`

GetSoftWareVersion returns the SoftWareVersion field if non-nil, zero value otherwise.

### GetSoftWareVersionOk

`func (o *OnuInformationConfigDTO) GetSoftWareVersionOk() (*string, bool)`

GetSoftWareVersionOk returns a tuple with the SoftWareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftWareVersion

`func (o *OnuInformationConfigDTO) SetSoftWareVersion(v string)`

SetSoftWareVersion sets SoftWareVersion field to given value.

### HasSoftWareVersion

`func (o *OnuInformationConfigDTO) HasSoftWareVersion() bool`

HasSoftWareVersion returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *OnuInformationConfigDTO) GetSoftwareVersion() OnuInformationConfigDTO`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *OnuInformationConfigDTO) GetSoftwareVersionOk() (*OnuInformationConfigDTO, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *OnuInformationConfigDTO) SetSoftwareVersion(v OnuInformationConfigDTO)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *OnuInformationConfigDTO) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetTransmittedOpticalPower

`func (o *OnuInformationConfigDTO) GetTransmittedOpticalPower() string`

GetTransmittedOpticalPower returns the TransmittedOpticalPower field if non-nil, zero value otherwise.

### GetTransmittedOpticalPowerOk

`func (o *OnuInformationConfigDTO) GetTransmittedOpticalPowerOk() (*string, bool)`

GetTransmittedOpticalPowerOk returns a tuple with the TransmittedOpticalPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmittedOpticalPower

`func (o *OnuInformationConfigDTO) SetTransmittedOpticalPower(v string)`

SetTransmittedOpticalPower sets TransmittedOpticalPower field to given value.

### HasTransmittedOpticalPower

`func (o *OnuInformationConfigDTO) HasTransmittedOpticalPower() bool`

HasTransmittedOpticalPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


