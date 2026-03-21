# TelephoneNumberAdvancedSettingOsgOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | Pointer to **string** | Interface ID, for example: if interfaceType is LAN network, interfaceId should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. When parameter [interfaceType] is 0 or 1, parameter [interfaceId] should not be empty. | [optional] 
**InterfaceType** | **int32** | InterfaceType should be a value as follows: 0: WAN; 1: LAN; 2: Virtual WAN. | 
**Locale** | Pointer to **string** | The country code of telephone number. | [optional] 
**NoAnswerTime** | Pointer to **int32** | The no answer time of telephone number. | [optional] 
**T38Support** | Pointer to **bool** | Whether to enable t38 Support. | [optional] 
**VirtualWanId** | Pointer to **string** | Virtual WAN ID, can be obtained from &#39;Query virtual WAN list&#39; interface. When parameter [interfaceType] is 2, parameter [virtualWanId] should not be empty. | [optional] 

## Methods

### NewTelephoneNumberAdvancedSettingOsgOpenApiVO

`func NewTelephoneNumberAdvancedSettingOsgOpenApiVO(interfaceType int32, ) *TelephoneNumberAdvancedSettingOsgOpenApiVO`

NewTelephoneNumberAdvancedSettingOsgOpenApiVO instantiates a new TelephoneNumberAdvancedSettingOsgOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephoneNumberAdvancedSettingOsgOpenApiVOWithDefaults

`func NewTelephoneNumberAdvancedSettingOsgOpenApiVOWithDefaults() *TelephoneNumberAdvancedSettingOsgOpenApiVO`

NewTelephoneNumberAdvancedSettingOsgOpenApiVOWithDefaults instantiates a new TelephoneNumberAdvancedSettingOsgOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) GetInterfaceType() int32`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) GetInterfaceTypeOk() (*int32, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) SetInterfaceType(v int32)`

SetInterfaceType sets InterfaceType field to given value.


### GetLocale

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetNoAnswerTime

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) GetNoAnswerTime() int32`

GetNoAnswerTime returns the NoAnswerTime field if non-nil, zero value otherwise.

### GetNoAnswerTimeOk

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) GetNoAnswerTimeOk() (*int32, bool)`

GetNoAnswerTimeOk returns a tuple with the NoAnswerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAnswerTime

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) SetNoAnswerTime(v int32)`

SetNoAnswerTime sets NoAnswerTime field to given value.

### HasNoAnswerTime

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) HasNoAnswerTime() bool`

HasNoAnswerTime returns a boolean if a field has been set.

### GetT38Support

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) GetT38Support() bool`

GetT38Support returns the T38Support field if non-nil, zero value otherwise.

### GetT38SupportOk

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) GetT38SupportOk() (*bool, bool)`

GetT38SupportOk returns a tuple with the T38Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT38Support

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) SetT38Support(v bool)`

SetT38Support sets T38Support field to given value.

### HasT38Support

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) HasT38Support() bool`

HasT38Support returns a boolean if a field has been set.

### GetVirtualWanId

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) GetVirtualWanId() string`

GetVirtualWanId returns the VirtualWanId field if non-nil, zero value otherwise.

### GetVirtualWanIdOk

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) GetVirtualWanIdOk() (*string, bool)`

GetVirtualWanIdOk returns a tuple with the VirtualWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanId

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) SetVirtualWanId(v string)`

SetVirtualWanId sets VirtualWanId field to given value.

### HasVirtualWanId

`func (o *TelephoneNumberAdvancedSettingOsgOpenApiVO) HasVirtualWanId() bool`

HasVirtualWanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


