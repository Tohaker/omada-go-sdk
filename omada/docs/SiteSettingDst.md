# SiteSettingDst

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | DST config status; If false, other parameters are not required. | [optional] 
**End** | Pointer to [**DstTimeDTO**](DstTimeDTO.md) |  | [optional] 
**Mode** | Pointer to **int32** | DST config mode; If disable, other parameters are not required. 0: disable, 1: auto, 2: manually | [optional] 
**Offset** | Pointer to **int64** | DST offset config(Unit: ms); It should be a value as follows: [1800000, 3600000, 5400000, 7200000]. | [optional] 
**Start** | Pointer to [**DstTimeDTO**](DstTimeDTO.md) |  | [optional] 

## Methods

### NewSiteSettingDst

`func NewSiteSettingDst() *SiteSettingDst`

NewSiteSettingDst instantiates a new SiteSettingDst object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingDstWithDefaults

`func NewSiteSettingDstWithDefaults() *SiteSettingDst`

NewSiteSettingDstWithDefaults instantiates a new SiteSettingDst object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SiteSettingDst) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SiteSettingDst) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SiteSettingDst) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SiteSettingDst) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnd

`func (o *SiteSettingDst) GetEnd() DstTimeDTO`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SiteSettingDst) GetEndOk() (*DstTimeDTO, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SiteSettingDst) SetEnd(v DstTimeDTO)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SiteSettingDst) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetMode

`func (o *SiteSettingDst) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SiteSettingDst) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SiteSettingDst) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SiteSettingDst) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOffset

`func (o *SiteSettingDst) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SiteSettingDst) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SiteSettingDst) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SiteSettingDst) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStart

`func (o *SiteSettingDst) GetStart() DstTimeDTO`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SiteSettingDst) GetStartOk() (*DstTimeDTO, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SiteSettingDst) SetStart(v DstTimeDTO)`

SetStart sets Start field to given value.

### HasStart

`func (o *SiteSettingDst) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


