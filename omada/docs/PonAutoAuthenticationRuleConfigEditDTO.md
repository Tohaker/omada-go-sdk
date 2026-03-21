# PonAutoAuthenticationRuleConfigEditDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipId** | Pointer to **string** | Matched ONT ID.EquipId should contain 1 to 20 characters in ASCII (from \\x21 to \\x7e). | [optional] 
**Key** | **string** | Identifier of entry | 
**LineProfile** | **string** | Line profile id | 
**LineProfileName** | Pointer to **string** | Line profile name | [optional] 
**PonPort** | **string** | Pon port.e.g.GPON 1/0/1 | 
**RuleId** | **int32** | Rule ID should be within the range of 1 to 128 | 
**ServicePortProfile** | Pointer to **string** | Service port profile id bound to the ONU | [optional] 
**ServiceProfile** | **string** | Service profile id | 
**ServiceProfileName** | Pointer to **string** | Service profile name | [optional] 
**SoftwareVersion** | Pointer to **string** | Software version should contain 1 to 14 characters, consisting of:Uppercase and lowercase letters,Digits,Special characters: -, @, _, :, /, . | [optional] 
**VendorId** | Pointer to **string** | Matched ONU vendor ID.VendorId should contain 1 to 4 characters in ASCII (from \\x21 to \\x7e). | [optional] 

## Methods

### NewPonAutoAuthenticationRuleConfigEditDTO

`func NewPonAutoAuthenticationRuleConfigEditDTO(key string, lineProfile string, ponPort string, ruleId int32, serviceProfile string, ) *PonAutoAuthenticationRuleConfigEditDTO`

NewPonAutoAuthenticationRuleConfigEditDTO instantiates a new PonAutoAuthenticationRuleConfigEditDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPonAutoAuthenticationRuleConfigEditDTOWithDefaults

`func NewPonAutoAuthenticationRuleConfigEditDTOWithDefaults() *PonAutoAuthenticationRuleConfigEditDTO`

NewPonAutoAuthenticationRuleConfigEditDTOWithDefaults instantiates a new PonAutoAuthenticationRuleConfigEditDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipId

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetEquipId() string`

GetEquipId returns the EquipId field if non-nil, zero value otherwise.

### GetEquipIdOk

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetEquipIdOk() (*string, bool)`

GetEquipIdOk returns a tuple with the EquipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipId

`func (o *PonAutoAuthenticationRuleConfigEditDTO) SetEquipId(v string)`

SetEquipId sets EquipId field to given value.

### HasEquipId

`func (o *PonAutoAuthenticationRuleConfigEditDTO) HasEquipId() bool`

HasEquipId returns a boolean if a field has been set.

### GetKey

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PonAutoAuthenticationRuleConfigEditDTO) SetKey(v string)`

SetKey sets Key field to given value.


### GetLineProfile

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetLineProfile() string`

GetLineProfile returns the LineProfile field if non-nil, zero value otherwise.

### GetLineProfileOk

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetLineProfileOk() (*string, bool)`

GetLineProfileOk returns a tuple with the LineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfile

`func (o *PonAutoAuthenticationRuleConfigEditDTO) SetLineProfile(v string)`

SetLineProfile sets LineProfile field to given value.


### GetLineProfileName

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetLineProfileName() string`

GetLineProfileName returns the LineProfileName field if non-nil, zero value otherwise.

### GetLineProfileNameOk

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetLineProfileNameOk() (*string, bool)`

GetLineProfileNameOk returns a tuple with the LineProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfileName

`func (o *PonAutoAuthenticationRuleConfigEditDTO) SetLineProfileName(v string)`

SetLineProfileName sets LineProfileName field to given value.

### HasLineProfileName

`func (o *PonAutoAuthenticationRuleConfigEditDTO) HasLineProfileName() bool`

HasLineProfileName returns a boolean if a field has been set.

### GetPonPort

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetPonPort() string`

GetPonPort returns the PonPort field if non-nil, zero value otherwise.

### GetPonPortOk

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetPonPortOk() (*string, bool)`

GetPonPortOk returns a tuple with the PonPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPort

`func (o *PonAutoAuthenticationRuleConfigEditDTO) SetPonPort(v string)`

SetPonPort sets PonPort field to given value.


### GetRuleId

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *PonAutoAuthenticationRuleConfigEditDTO) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.


### GetServicePortProfile

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetServicePortProfile() string`

GetServicePortProfile returns the ServicePortProfile field if non-nil, zero value otherwise.

### GetServicePortProfileOk

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetServicePortProfileOk() (*string, bool)`

GetServicePortProfileOk returns a tuple with the ServicePortProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortProfile

`func (o *PonAutoAuthenticationRuleConfigEditDTO) SetServicePortProfile(v string)`

SetServicePortProfile sets ServicePortProfile field to given value.

### HasServicePortProfile

`func (o *PonAutoAuthenticationRuleConfigEditDTO) HasServicePortProfile() bool`

HasServicePortProfile returns a boolean if a field has been set.

### GetServiceProfile

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *PonAutoAuthenticationRuleConfigEditDTO) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.


### GetServiceProfileName

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetServiceProfileName() string`

GetServiceProfileName returns the ServiceProfileName field if non-nil, zero value otherwise.

### GetServiceProfileNameOk

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetServiceProfileNameOk() (*string, bool)`

GetServiceProfileNameOk returns a tuple with the ServiceProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfileName

`func (o *PonAutoAuthenticationRuleConfigEditDTO) SetServiceProfileName(v string)`

SetServiceProfileName sets ServiceProfileName field to given value.

### HasServiceProfileName

`func (o *PonAutoAuthenticationRuleConfigEditDTO) HasServiceProfileName() bool`

HasServiceProfileName returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *PonAutoAuthenticationRuleConfigEditDTO) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *PonAutoAuthenticationRuleConfigEditDTO) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetVendorId

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *PonAutoAuthenticationRuleConfigEditDTO) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *PonAutoAuthenticationRuleConfigEditDTO) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *PonAutoAuthenticationRuleConfigEditDTO) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


