# OnuAutofindConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentId** | Pointer to **string** | ONT device ID should contain 1 to 20 characters in ASCII code from \\x21 to \\x7e. | [optional] 
**HardwareVersion** | Pointer to **string** | Hardware version of the automatically discovered ONT | [optional] 
**Key** | Pointer to **string** | Unique identifier for entry | [optional] 
**Loid** | Pointer to **string** | Loid set by automatic discovery of the ONT, loid should be 1-32 characters, including letters, numbers, and symbols (-@_:/.). | [optional] 
**LoidPassword** | Pointer to **string** | Loid password set by automatic discovery of the ONT,loidPassword should contain ASCII characters from \\x21 to \\x7e | [optional] 
**MacAddress** | Pointer to **string** | Mac address of ONU | [optional] 
**Password** | Pointer to **string** | Authentication password set by automatic discovery of the ONT, password should contain ASCII characters from \\x21 to \\x7e and 1 to 20 hexadecimal digits | [optional] 
**PasswordType** | Pointer to **string** | Indicates the password type reported by the automatic discovery of the ONT.PasswordType should be a value as follows：ASCII,HEX | [optional] 
**PortId** | Pointer to **string** | Automatic discovery of the PON port ID corresponding to the ONT | [optional] 
**SerialNumber** | Pointer to **string** | SerialNumber of ONU, serialNumber should contain 12, 13, or 16 characters, in the format XXXXXXXXXXXX,XXXX-XXXXXXXX,XXXXXXXXXXXXXXXX | [optional] 
**SoftwareVersion** | Pointer to **string** | Software version of the automatically discovered ONT | [optional] 
**VendorId** | Pointer to **string** | Matching ONT vendor ID, vendorId should contain 1 to 4 ASCII characters from \\x21 to \\x7e | [optional] 

## Methods

### NewOnuAutofindConfigDTO

`func NewOnuAutofindConfigDTO() *OnuAutofindConfigDTO`

NewOnuAutofindConfigDTO instantiates a new OnuAutofindConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuAutofindConfigDTOWithDefaults

`func NewOnuAutofindConfigDTOWithDefaults() *OnuAutofindConfigDTO`

NewOnuAutofindConfigDTOWithDefaults instantiates a new OnuAutofindConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentId

`func (o *OnuAutofindConfigDTO) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *OnuAutofindConfigDTO) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *OnuAutofindConfigDTO) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.

### HasEquipmentId

`func (o *OnuAutofindConfigDTO) HasEquipmentId() bool`

HasEquipmentId returns a boolean if a field has been set.

### GetHardwareVersion

`func (o *OnuAutofindConfigDTO) GetHardwareVersion() string`

GetHardwareVersion returns the HardwareVersion field if non-nil, zero value otherwise.

### GetHardwareVersionOk

`func (o *OnuAutofindConfigDTO) GetHardwareVersionOk() (*string, bool)`

GetHardwareVersionOk returns a tuple with the HardwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareVersion

`func (o *OnuAutofindConfigDTO) SetHardwareVersion(v string)`

SetHardwareVersion sets HardwareVersion field to given value.

### HasHardwareVersion

`func (o *OnuAutofindConfigDTO) HasHardwareVersion() bool`

HasHardwareVersion returns a boolean if a field has been set.

### GetKey

`func (o *OnuAutofindConfigDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OnuAutofindConfigDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OnuAutofindConfigDTO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OnuAutofindConfigDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLoid

`func (o *OnuAutofindConfigDTO) GetLoid() string`

GetLoid returns the Loid field if non-nil, zero value otherwise.

### GetLoidOk

`func (o *OnuAutofindConfigDTO) GetLoidOk() (*string, bool)`

GetLoidOk returns a tuple with the Loid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoid

`func (o *OnuAutofindConfigDTO) SetLoid(v string)`

SetLoid sets Loid field to given value.

### HasLoid

`func (o *OnuAutofindConfigDTO) HasLoid() bool`

HasLoid returns a boolean if a field has been set.

### GetLoidPassword

`func (o *OnuAutofindConfigDTO) GetLoidPassword() string`

GetLoidPassword returns the LoidPassword field if non-nil, zero value otherwise.

### GetLoidPasswordOk

`func (o *OnuAutofindConfigDTO) GetLoidPasswordOk() (*string, bool)`

GetLoidPasswordOk returns a tuple with the LoidPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoidPassword

`func (o *OnuAutofindConfigDTO) SetLoidPassword(v string)`

SetLoidPassword sets LoidPassword field to given value.

### HasLoidPassword

`func (o *OnuAutofindConfigDTO) HasLoidPassword() bool`

HasLoidPassword returns a boolean if a field has been set.

### GetMacAddress

`func (o *OnuAutofindConfigDTO) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *OnuAutofindConfigDTO) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *OnuAutofindConfigDTO) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *OnuAutofindConfigDTO) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPassword

`func (o *OnuAutofindConfigDTO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OnuAutofindConfigDTO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OnuAutofindConfigDTO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OnuAutofindConfigDTO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordType

`func (o *OnuAutofindConfigDTO) GetPasswordType() string`

GetPasswordType returns the PasswordType field if non-nil, zero value otherwise.

### GetPasswordTypeOk

`func (o *OnuAutofindConfigDTO) GetPasswordTypeOk() (*string, bool)`

GetPasswordTypeOk returns a tuple with the PasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordType

`func (o *OnuAutofindConfigDTO) SetPasswordType(v string)`

SetPasswordType sets PasswordType field to given value.

### HasPasswordType

`func (o *OnuAutofindConfigDTO) HasPasswordType() bool`

HasPasswordType returns a boolean if a field has been set.

### GetPortId

`func (o *OnuAutofindConfigDTO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *OnuAutofindConfigDTO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *OnuAutofindConfigDTO) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *OnuAutofindConfigDTO) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *OnuAutofindConfigDTO) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *OnuAutofindConfigDTO) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *OnuAutofindConfigDTO) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *OnuAutofindConfigDTO) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *OnuAutofindConfigDTO) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *OnuAutofindConfigDTO) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *OnuAutofindConfigDTO) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *OnuAutofindConfigDTO) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetVendorId

`func (o *OnuAutofindConfigDTO) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *OnuAutofindConfigDTO) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *OnuAutofindConfigDTO) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *OnuAutofindConfigDTO) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


