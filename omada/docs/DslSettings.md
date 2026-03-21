# DslSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnexType** | Pointer to **int32** | Enter a value as follows: 0-Annex Auto, 1-Annex A, 2-Annex B, 3-Annex I, 4-Annex J, 5-Annex M, 6-Annex A/L, 7-Annex B/J, 8-Annex A/I/J/L/M | [optional] 
**BitSwap** | Pointer to **int32** | Enter a value as follows: 0-off, 1-on. | [optional] 
**ModulationType** | Pointer to **int32** | Enter a value as follows: 0-Auto Sync-up, 2-ADSL2, 3-ADSL2+, 4-G.dmt, 5-T1.423, 6-VDSL2. | [optional] 
**Sra** | Pointer to **int32** | Enter a value as follows: 0-off, 1-on. | [optional] 

## Methods

### NewDslSettings

`func NewDslSettings() *DslSettings`

NewDslSettings instantiates a new DslSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDslSettingsWithDefaults

`func NewDslSettingsWithDefaults() *DslSettings`

NewDslSettingsWithDefaults instantiates a new DslSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnexType

`func (o *DslSettings) GetAnnexType() int32`

GetAnnexType returns the AnnexType field if non-nil, zero value otherwise.

### GetAnnexTypeOk

`func (o *DslSettings) GetAnnexTypeOk() (*int32, bool)`

GetAnnexTypeOk returns a tuple with the AnnexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnexType

`func (o *DslSettings) SetAnnexType(v int32)`

SetAnnexType sets AnnexType field to given value.

### HasAnnexType

`func (o *DslSettings) HasAnnexType() bool`

HasAnnexType returns a boolean if a field has been set.

### GetBitSwap

`func (o *DslSettings) GetBitSwap() int32`

GetBitSwap returns the BitSwap field if non-nil, zero value otherwise.

### GetBitSwapOk

`func (o *DslSettings) GetBitSwapOk() (*int32, bool)`

GetBitSwapOk returns a tuple with the BitSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitSwap

`func (o *DslSettings) SetBitSwap(v int32)`

SetBitSwap sets BitSwap field to given value.

### HasBitSwap

`func (o *DslSettings) HasBitSwap() bool`

HasBitSwap returns a boolean if a field has been set.

### GetModulationType

`func (o *DslSettings) GetModulationType() int32`

GetModulationType returns the ModulationType field if non-nil, zero value otherwise.

### GetModulationTypeOk

`func (o *DslSettings) GetModulationTypeOk() (*int32, bool)`

GetModulationTypeOk returns a tuple with the ModulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulationType

`func (o *DslSettings) SetModulationType(v int32)`

SetModulationType sets ModulationType field to given value.

### HasModulationType

`func (o *DslSettings) HasModulationType() bool`

HasModulationType returns a boolean if a field has been set.

### GetSra

`func (o *DslSettings) GetSra() int32`

GetSra returns the Sra field if non-nil, zero value otherwise.

### GetSraOk

`func (o *DslSettings) GetSraOk() (*int32, bool)`

GetSraOk returns a tuple with the Sra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSra

`func (o *DslSettings) SetSra(v int32)`

SetSra sets Sra field to given value.

### HasSra

`func (o *DslSettings) HasSra() bool`

HasSra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


