# BasicDetailConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveStatus** | Pointer to **string** | Active status should be a value as follows:ACTIVE: ONU is in an active state, capable of data communication and carrying services. INACTIVE: ONU is in an inactive state, possibly due to unactivated configuration, being offline, configuration failure, or mismatched configuration. | [optional] 
**AdminStatus** | Pointer to **string** | Admin status should be a value as follows:ACTIVATE,DEACTIVATE | [optional] 
**ConfigStatus** | Pointer to **string** | Config status should be a value as follows:SUCCESS: Configuration successfully delivered and recognized by the ONU.FAILED: Configuration not successfully delivered or not recognized by the ONU. | [optional] 
**EquipmentId** | **string** | ONT device ID should contain 1 to 20 characters in ASCII code from \\x21 to \\x7e. | 
**HardwareVersion** | Pointer to **string** | Hardware version of the ONU, hardwareVersion should contain 1 to 20 characters, including uppercase and lowercase letters, numbers, and the symbols -@_:/.  | [optional] 
**LineProfile** | Pointer to **string** | LineProfile is displayed in the format id(name) | [optional] 
**MacAddress** | Pointer to **string** | Mac address of ONU | [optional] 
**MatchStatus** | Pointer to **string** | Match status should be a value as follows:MATCH: ONU hardware capabilities are consistent with the bound service template.MISMATCH: ONU hardware capabilities are inconsistent with the bound service template. | [optional] 
**OnlineStatus** | Pointer to **string** | Online status should be a value as follows:ONLINE,OFFLINE,UPGRADING | [optional] 
**OnlineTime** | Pointer to **string** | Duration of ONU registration and online status, format: hh:mm:ss. | [optional] 
**OnuDescription** | Pointer to **string** | ONU description.onuDescription should contain 0 to 32 characters, including uppercase and lowercase letters, numbers, and the symbols -@_:/.  | [optional] 
**OnuDistance** | Pointer to **int32** | Distance between ONU and OLT should be within the range of  0 to 20,000 m | [optional] 
**SerialNumber** | Pointer to **string** | SerialNumber of ONU should contain 12, 13, or 16 characters, in the format XXXXXXXXXXXX ,XXXX-XXXXXXXX,XXXXXXXXXXXXXXXX | [optional] 
**ServiceProfile** | Pointer to **string** | ServiceProfile is displayed in the format id(name) | [optional] 
**VendorId** | Pointer to **string** | The matching ONT vendor ID,vendorId should contain 1 to 4 characters in ASCII, with character range from \\x21 to \\x7e. | [optional] 

## Methods

### NewBasicDetailConfigDTO

`func NewBasicDetailConfigDTO(equipmentId string, ) *BasicDetailConfigDTO`

NewBasicDetailConfigDTO instantiates a new BasicDetailConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicDetailConfigDTOWithDefaults

`func NewBasicDetailConfigDTOWithDefaults() *BasicDetailConfigDTO`

NewBasicDetailConfigDTOWithDefaults instantiates a new BasicDetailConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveStatus

`func (o *BasicDetailConfigDTO) GetActiveStatus() string`

GetActiveStatus returns the ActiveStatus field if non-nil, zero value otherwise.

### GetActiveStatusOk

`func (o *BasicDetailConfigDTO) GetActiveStatusOk() (*string, bool)`

GetActiveStatusOk returns a tuple with the ActiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStatus

`func (o *BasicDetailConfigDTO) SetActiveStatus(v string)`

SetActiveStatus sets ActiveStatus field to given value.

### HasActiveStatus

`func (o *BasicDetailConfigDTO) HasActiveStatus() bool`

HasActiveStatus returns a boolean if a field has been set.

### GetAdminStatus

`func (o *BasicDetailConfigDTO) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *BasicDetailConfigDTO) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *BasicDetailConfigDTO) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *BasicDetailConfigDTO) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetConfigStatus

`func (o *BasicDetailConfigDTO) GetConfigStatus() string`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *BasicDetailConfigDTO) GetConfigStatusOk() (*string, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *BasicDetailConfigDTO) SetConfigStatus(v string)`

SetConfigStatus sets ConfigStatus field to given value.

### HasConfigStatus

`func (o *BasicDetailConfigDTO) HasConfigStatus() bool`

HasConfigStatus returns a boolean if a field has been set.

### GetEquipmentId

`func (o *BasicDetailConfigDTO) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *BasicDetailConfigDTO) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *BasicDetailConfigDTO) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.


### GetHardwareVersion

`func (o *BasicDetailConfigDTO) GetHardwareVersion() string`

GetHardwareVersion returns the HardwareVersion field if non-nil, zero value otherwise.

### GetHardwareVersionOk

`func (o *BasicDetailConfigDTO) GetHardwareVersionOk() (*string, bool)`

GetHardwareVersionOk returns a tuple with the HardwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareVersion

`func (o *BasicDetailConfigDTO) SetHardwareVersion(v string)`

SetHardwareVersion sets HardwareVersion field to given value.

### HasHardwareVersion

`func (o *BasicDetailConfigDTO) HasHardwareVersion() bool`

HasHardwareVersion returns a boolean if a field has been set.

### GetLineProfile

`func (o *BasicDetailConfigDTO) GetLineProfile() string`

GetLineProfile returns the LineProfile field if non-nil, zero value otherwise.

### GetLineProfileOk

`func (o *BasicDetailConfigDTO) GetLineProfileOk() (*string, bool)`

GetLineProfileOk returns a tuple with the LineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfile

`func (o *BasicDetailConfigDTO) SetLineProfile(v string)`

SetLineProfile sets LineProfile field to given value.

### HasLineProfile

`func (o *BasicDetailConfigDTO) HasLineProfile() bool`

HasLineProfile returns a boolean if a field has been set.

### GetMacAddress

`func (o *BasicDetailConfigDTO) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BasicDetailConfigDTO) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BasicDetailConfigDTO) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *BasicDetailConfigDTO) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMatchStatus

`func (o *BasicDetailConfigDTO) GetMatchStatus() string`

GetMatchStatus returns the MatchStatus field if non-nil, zero value otherwise.

### GetMatchStatusOk

`func (o *BasicDetailConfigDTO) GetMatchStatusOk() (*string, bool)`

GetMatchStatusOk returns a tuple with the MatchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchStatus

`func (o *BasicDetailConfigDTO) SetMatchStatus(v string)`

SetMatchStatus sets MatchStatus field to given value.

### HasMatchStatus

`func (o *BasicDetailConfigDTO) HasMatchStatus() bool`

HasMatchStatus returns a boolean if a field has been set.

### GetOnlineStatus

`func (o *BasicDetailConfigDTO) GetOnlineStatus() string`

GetOnlineStatus returns the OnlineStatus field if non-nil, zero value otherwise.

### GetOnlineStatusOk

`func (o *BasicDetailConfigDTO) GetOnlineStatusOk() (*string, bool)`

GetOnlineStatusOk returns a tuple with the OnlineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineStatus

`func (o *BasicDetailConfigDTO) SetOnlineStatus(v string)`

SetOnlineStatus sets OnlineStatus field to given value.

### HasOnlineStatus

`func (o *BasicDetailConfigDTO) HasOnlineStatus() bool`

HasOnlineStatus returns a boolean if a field has been set.

### GetOnlineTime

`func (o *BasicDetailConfigDTO) GetOnlineTime() string`

GetOnlineTime returns the OnlineTime field if non-nil, zero value otherwise.

### GetOnlineTimeOk

`func (o *BasicDetailConfigDTO) GetOnlineTimeOk() (*string, bool)`

GetOnlineTimeOk returns a tuple with the OnlineTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineTime

`func (o *BasicDetailConfigDTO) SetOnlineTime(v string)`

SetOnlineTime sets OnlineTime field to given value.

### HasOnlineTime

`func (o *BasicDetailConfigDTO) HasOnlineTime() bool`

HasOnlineTime returns a boolean if a field has been set.

### GetOnuDescription

`func (o *BasicDetailConfigDTO) GetOnuDescription() string`

GetOnuDescription returns the OnuDescription field if non-nil, zero value otherwise.

### GetOnuDescriptionOk

`func (o *BasicDetailConfigDTO) GetOnuDescriptionOk() (*string, bool)`

GetOnuDescriptionOk returns a tuple with the OnuDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuDescription

`func (o *BasicDetailConfigDTO) SetOnuDescription(v string)`

SetOnuDescription sets OnuDescription field to given value.

### HasOnuDescription

`func (o *BasicDetailConfigDTO) HasOnuDescription() bool`

HasOnuDescription returns a boolean if a field has been set.

### GetOnuDistance

`func (o *BasicDetailConfigDTO) GetOnuDistance() int32`

GetOnuDistance returns the OnuDistance field if non-nil, zero value otherwise.

### GetOnuDistanceOk

`func (o *BasicDetailConfigDTO) GetOnuDistanceOk() (*int32, bool)`

GetOnuDistanceOk returns a tuple with the OnuDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuDistance

`func (o *BasicDetailConfigDTO) SetOnuDistance(v int32)`

SetOnuDistance sets OnuDistance field to given value.

### HasOnuDistance

`func (o *BasicDetailConfigDTO) HasOnuDistance() bool`

HasOnuDistance returns a boolean if a field has been set.

### GetSerialNumber

`func (o *BasicDetailConfigDTO) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *BasicDetailConfigDTO) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *BasicDetailConfigDTO) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *BasicDetailConfigDTO) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetServiceProfile

`func (o *BasicDetailConfigDTO) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *BasicDetailConfigDTO) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *BasicDetailConfigDTO) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *BasicDetailConfigDTO) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.

### GetVendorId

`func (o *BasicDetailConfigDTO) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *BasicDetailConfigDTO) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *BasicDetailConfigDTO) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *BasicDetailConfigDTO) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


