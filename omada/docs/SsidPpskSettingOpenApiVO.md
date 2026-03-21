# SsidPpskSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacFormat** | Pointer to **int32** | MAC address format. This is necessary when the value of security is 5(PPSK with RADIUS); It should be a value as follows: 0: aabbccddeeff; 1: aa-bb-cc-dd-ee-ff; 2: aa:bb:cc:dd:ee:ff; 3: AABBCCDDEEFF; 4: AA-BB-CC-DD-EE-FF; 5: AA:BB:CC:DD:EE:FF. | [optional] 
**NasId** | Pointer to **string** | NAS ID. This is necessary when the value of security is 5(PPSK with RADIUS); It should contain 1 to 64 characters. | [optional] 
**PpskProfileId** | Pointer to **string** | This field represents PPSK Profile ID; This is necessary when the value of security is 4(PPSK without RADIUS); PPSK Profile(PPSK Profile Template) can be created using Create PPSK profile interface(Create PPSK profile template interface), and PPSK Profile ID(PPSK Profile Template ID) can be obtained from Get PPSK profiles list(Get PPSK profile templates list) interface. | [optional] 
**RadiusProfileId** | Pointer to **string** | This field represents RADIUS Profile ID; This is necessary when the value of security is 5(PPSK with RADIUS); RADIUS Profile(RADIUS Profile Template) can be created using Create a new Radius profile(Create a new Radius profile template) interface, and RADIUS Profile ID(RADIUS Profile Template ID) can be obtained from Get Radius profile list(Get Radius profile template list) interface. | [optional] 
**Type** | Pointer to **int32** | Authentication type. This is necessary when the value of security is 5(PPSK with RADIUS); It should be a value as follows: 0: Mac Auth(Generic Radius with bound MAC); 1: EKMS(This configuration applies to the Pro Site of the Omada Pro Controller only); 2: Generic Radius with unbound MAC(This configuration applies to the Pro Site of the Omada Pro Controller only). | [optional] 

## Methods

### NewSsidPpskSettingOpenApiVO

`func NewSsidPpskSettingOpenApiVO() *SsidPpskSettingOpenApiVO`

NewSsidPpskSettingOpenApiVO instantiates a new SsidPpskSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidPpskSettingOpenApiVOWithDefaults

`func NewSsidPpskSettingOpenApiVOWithDefaults() *SsidPpskSettingOpenApiVO`

NewSsidPpskSettingOpenApiVOWithDefaults instantiates a new SsidPpskSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacFormat

`func (o *SsidPpskSettingOpenApiVO) GetMacFormat() int32`

GetMacFormat returns the MacFormat field if non-nil, zero value otherwise.

### GetMacFormatOk

`func (o *SsidPpskSettingOpenApiVO) GetMacFormatOk() (*int32, bool)`

GetMacFormatOk returns a tuple with the MacFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFormat

`func (o *SsidPpskSettingOpenApiVO) SetMacFormat(v int32)`

SetMacFormat sets MacFormat field to given value.

### HasMacFormat

`func (o *SsidPpskSettingOpenApiVO) HasMacFormat() bool`

HasMacFormat returns a boolean if a field has been set.

### GetNasId

`func (o *SsidPpskSettingOpenApiVO) GetNasId() string`

GetNasId returns the NasId field if non-nil, zero value otherwise.

### GetNasIdOk

`func (o *SsidPpskSettingOpenApiVO) GetNasIdOk() (*string, bool)`

GetNasIdOk returns a tuple with the NasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasId

`func (o *SsidPpskSettingOpenApiVO) SetNasId(v string)`

SetNasId sets NasId field to given value.

### HasNasId

`func (o *SsidPpskSettingOpenApiVO) HasNasId() bool`

HasNasId returns a boolean if a field has been set.

### GetPpskProfileId

`func (o *SsidPpskSettingOpenApiVO) GetPpskProfileId() string`

GetPpskProfileId returns the PpskProfileId field if non-nil, zero value otherwise.

### GetPpskProfileIdOk

`func (o *SsidPpskSettingOpenApiVO) GetPpskProfileIdOk() (*string, bool)`

GetPpskProfileIdOk returns a tuple with the PpskProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpskProfileId

`func (o *SsidPpskSettingOpenApiVO) SetPpskProfileId(v string)`

SetPpskProfileId sets PpskProfileId field to given value.

### HasPpskProfileId

`func (o *SsidPpskSettingOpenApiVO) HasPpskProfileId() bool`

HasPpskProfileId returns a boolean if a field has been set.

### GetRadiusProfileId

`func (o *SsidPpskSettingOpenApiVO) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *SsidPpskSettingOpenApiVO) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *SsidPpskSettingOpenApiVO) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.

### HasRadiusProfileId

`func (o *SsidPpskSettingOpenApiVO) HasRadiusProfileId() bool`

HasRadiusProfileId returns a boolean if a field has been set.

### GetType

`func (o *SsidPpskSettingOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SsidPpskSettingOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SsidPpskSettingOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SsidPpskSettingOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


