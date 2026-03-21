# AdvertisementSettingResOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Whether to enable Advertisement setting | [optional] 
**PictureInterval** | Pointer to **int32** | Advertisement picture interval, should be within the range of 1–10 seconds | [optional] 
**Pictures** | Pointer to [**[]PortalPictureInfo**](PortalPictureInfo.md) | Advertisement picture list. Up to 5 entries are allowed for the pictures list | [optional] 
**SkipEnable** | Pointer to **bool** | Whether to allow users to skip the advertisement | [optional] 
**TotalDuration** | Pointer to **int32** | Total duration of advertisement, should be within the range of 1–30 seconds | [optional] 

## Methods

### NewAdvertisementSettingResOpenApiVO

`func NewAdvertisementSettingResOpenApiVO() *AdvertisementSettingResOpenApiVO`

NewAdvertisementSettingResOpenApiVO instantiates a new AdvertisementSettingResOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvertisementSettingResOpenApiVOWithDefaults

`func NewAdvertisementSettingResOpenApiVOWithDefaults() *AdvertisementSettingResOpenApiVO`

NewAdvertisementSettingResOpenApiVOWithDefaults instantiates a new AdvertisementSettingResOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *AdvertisementSettingResOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AdvertisementSettingResOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AdvertisementSettingResOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AdvertisementSettingResOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPictureInterval

`func (o *AdvertisementSettingResOpenApiVO) GetPictureInterval() int32`

GetPictureInterval returns the PictureInterval field if non-nil, zero value otherwise.

### GetPictureIntervalOk

`func (o *AdvertisementSettingResOpenApiVO) GetPictureIntervalOk() (*int32, bool)`

GetPictureIntervalOk returns a tuple with the PictureInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureInterval

`func (o *AdvertisementSettingResOpenApiVO) SetPictureInterval(v int32)`

SetPictureInterval sets PictureInterval field to given value.

### HasPictureInterval

`func (o *AdvertisementSettingResOpenApiVO) HasPictureInterval() bool`

HasPictureInterval returns a boolean if a field has been set.

### GetPictures

`func (o *AdvertisementSettingResOpenApiVO) GetPictures() []PortalPictureInfo`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *AdvertisementSettingResOpenApiVO) GetPicturesOk() (*[]PortalPictureInfo, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *AdvertisementSettingResOpenApiVO) SetPictures(v []PortalPictureInfo)`

SetPictures sets Pictures field to given value.

### HasPictures

`func (o *AdvertisementSettingResOpenApiVO) HasPictures() bool`

HasPictures returns a boolean if a field has been set.

### GetSkipEnable

`func (o *AdvertisementSettingResOpenApiVO) GetSkipEnable() bool`

GetSkipEnable returns the SkipEnable field if non-nil, zero value otherwise.

### GetSkipEnableOk

`func (o *AdvertisementSettingResOpenApiVO) GetSkipEnableOk() (*bool, bool)`

GetSkipEnableOk returns a tuple with the SkipEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEnable

`func (o *AdvertisementSettingResOpenApiVO) SetSkipEnable(v bool)`

SetSkipEnable sets SkipEnable field to given value.

### HasSkipEnable

`func (o *AdvertisementSettingResOpenApiVO) HasSkipEnable() bool`

HasSkipEnable returns a boolean if a field has been set.

### GetTotalDuration

`func (o *AdvertisementSettingResOpenApiVO) GetTotalDuration() int32`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *AdvertisementSettingResOpenApiVO) GetTotalDurationOk() (*int32, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *AdvertisementSettingResOpenApiVO) SetTotalDuration(v int32)`

SetTotalDuration sets TotalDuration field to given value.

### HasTotalDuration

`func (o *AdvertisementSettingResOpenApiVO) HasTotalDuration() bool`

HasTotalDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


