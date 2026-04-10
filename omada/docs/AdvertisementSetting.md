# AdvertisementSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Advertisement enable. | 
**PictureIds** | Pointer to **[]string** | Picture ID list, Up to 5 entries are allowed for the pictureIds list. | [optional] 
**PictureInterval** | Pointer to **int32** | Advertisement picture interval, should be within the range of 1–10, time unit is second. | [optional] 
**SkipDelay** | Pointer to **int32** | Skip delay, should be within the range of 1–10, time unit is second. | [optional] 
**SkipEnable** | Pointer to **bool** | Whether allow users to skip the advertisement. | [optional] 
**TotalDuration** | Pointer to **int32** | Advertisement totalDuration, should be within the range of 1–30, time unit is second. | [optional] 

## Methods

### NewAdvertisementSetting

`func NewAdvertisementSetting(enable bool, ) *AdvertisementSetting`

NewAdvertisementSetting instantiates a new AdvertisementSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvertisementSettingWithDefaults

`func NewAdvertisementSettingWithDefaults() *AdvertisementSetting`

NewAdvertisementSettingWithDefaults instantiates a new AdvertisementSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *AdvertisementSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AdvertisementSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AdvertisementSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetPictureIds

`func (o *AdvertisementSetting) GetPictureIds() []string`

GetPictureIds returns the PictureIds field if non-nil, zero value otherwise.

### GetPictureIdsOk

`func (o *AdvertisementSetting) GetPictureIdsOk() (*[]string, bool)`

GetPictureIdsOk returns a tuple with the PictureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureIds

`func (o *AdvertisementSetting) SetPictureIds(v []string)`

SetPictureIds sets PictureIds field to given value.

### HasPictureIds

`func (o *AdvertisementSetting) HasPictureIds() bool`

HasPictureIds returns a boolean if a field has been set.

### GetPictureInterval

`func (o *AdvertisementSetting) GetPictureInterval() int32`

GetPictureInterval returns the PictureInterval field if non-nil, zero value otherwise.

### GetPictureIntervalOk

`func (o *AdvertisementSetting) GetPictureIntervalOk() (*int32, bool)`

GetPictureIntervalOk returns a tuple with the PictureInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureInterval

`func (o *AdvertisementSetting) SetPictureInterval(v int32)`

SetPictureInterval sets PictureInterval field to given value.

### HasPictureInterval

`func (o *AdvertisementSetting) HasPictureInterval() bool`

HasPictureInterval returns a boolean if a field has been set.

### GetSkipDelay

`func (o *AdvertisementSetting) GetSkipDelay() int32`

GetSkipDelay returns the SkipDelay field if non-nil, zero value otherwise.

### GetSkipDelayOk

`func (o *AdvertisementSetting) GetSkipDelayOk() (*int32, bool)`

GetSkipDelayOk returns a tuple with the SkipDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDelay

`func (o *AdvertisementSetting) SetSkipDelay(v int32)`

SetSkipDelay sets SkipDelay field to given value.

### HasSkipDelay

`func (o *AdvertisementSetting) HasSkipDelay() bool`

HasSkipDelay returns a boolean if a field has been set.

### GetSkipEnable

`func (o *AdvertisementSetting) GetSkipEnable() bool`

GetSkipEnable returns the SkipEnable field if non-nil, zero value otherwise.

### GetSkipEnableOk

`func (o *AdvertisementSetting) GetSkipEnableOk() (*bool, bool)`

GetSkipEnableOk returns a tuple with the SkipEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEnable

`func (o *AdvertisementSetting) SetSkipEnable(v bool)`

SetSkipEnable sets SkipEnable field to given value.

### HasSkipEnable

`func (o *AdvertisementSetting) HasSkipEnable() bool`

HasSkipEnable returns a boolean if a field has been set.

### GetTotalDuration

`func (o *AdvertisementSetting) GetTotalDuration() int32`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *AdvertisementSetting) GetTotalDurationOk() (*int32, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *AdvertisementSetting) SetTotalDuration(v int32)`

SetTotalDuration sets TotalDuration field to given value.

### HasTotalDuration

`func (o *AdvertisementSetting) HasTotalDuration() bool`

HasTotalDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


