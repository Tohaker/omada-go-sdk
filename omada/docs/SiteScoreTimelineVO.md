# SiteScoreTimelineVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientHealthScore** | Pointer to **int32** | Client health score | [optional] 
**DeviceHealthScore** | Pointer to **int32** | Device health score | [optional] 
**ScoreDetail** | Pointer to [**SiteSubHealthScoreVO**](SiteSubHealthScoreVO.md) |  | [optional] 
**SiteHealthScore** | Pointer to **int32** | Site health score | [optional] 
**Time** | Pointer to **int64** | Time(unit:ms) | [optional] 
**WanHealthScore** | Pointer to **int32** | WAN health score | [optional] 
**WifiHealthScore** | Pointer to **int32** | Wifi health score | [optional] 

## Methods

### NewSiteScoreTimelineVO

`func NewSiteScoreTimelineVO() *SiteScoreTimelineVO`

NewSiteScoreTimelineVO instantiates a new SiteScoreTimelineVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteScoreTimelineVOWithDefaults

`func NewSiteScoreTimelineVOWithDefaults() *SiteScoreTimelineVO`

NewSiteScoreTimelineVOWithDefaults instantiates a new SiteScoreTimelineVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientHealthScore

`func (o *SiteScoreTimelineVO) GetClientHealthScore() int32`

GetClientHealthScore returns the ClientHealthScore field if non-nil, zero value otherwise.

### GetClientHealthScoreOk

`func (o *SiteScoreTimelineVO) GetClientHealthScoreOk() (*int32, bool)`

GetClientHealthScoreOk returns a tuple with the ClientHealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHealthScore

`func (o *SiteScoreTimelineVO) SetClientHealthScore(v int32)`

SetClientHealthScore sets ClientHealthScore field to given value.

### HasClientHealthScore

`func (o *SiteScoreTimelineVO) HasClientHealthScore() bool`

HasClientHealthScore returns a boolean if a field has been set.

### GetDeviceHealthScore

`func (o *SiteScoreTimelineVO) GetDeviceHealthScore() int32`

GetDeviceHealthScore returns the DeviceHealthScore field if non-nil, zero value otherwise.

### GetDeviceHealthScoreOk

`func (o *SiteScoreTimelineVO) GetDeviceHealthScoreOk() (*int32, bool)`

GetDeviceHealthScoreOk returns a tuple with the DeviceHealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHealthScore

`func (o *SiteScoreTimelineVO) SetDeviceHealthScore(v int32)`

SetDeviceHealthScore sets DeviceHealthScore field to given value.

### HasDeviceHealthScore

`func (o *SiteScoreTimelineVO) HasDeviceHealthScore() bool`

HasDeviceHealthScore returns a boolean if a field has been set.

### GetScoreDetail

`func (o *SiteScoreTimelineVO) GetScoreDetail() SiteSubHealthScoreVO`

GetScoreDetail returns the ScoreDetail field if non-nil, zero value otherwise.

### GetScoreDetailOk

`func (o *SiteScoreTimelineVO) GetScoreDetailOk() (*SiteSubHealthScoreVO, bool)`

GetScoreDetailOk returns a tuple with the ScoreDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreDetail

`func (o *SiteScoreTimelineVO) SetScoreDetail(v SiteSubHealthScoreVO)`

SetScoreDetail sets ScoreDetail field to given value.

### HasScoreDetail

`func (o *SiteScoreTimelineVO) HasScoreDetail() bool`

HasScoreDetail returns a boolean if a field has been set.

### GetSiteHealthScore

`func (o *SiteScoreTimelineVO) GetSiteHealthScore() int32`

GetSiteHealthScore returns the SiteHealthScore field if non-nil, zero value otherwise.

### GetSiteHealthScoreOk

`func (o *SiteScoreTimelineVO) GetSiteHealthScoreOk() (*int32, bool)`

GetSiteHealthScoreOk returns a tuple with the SiteHealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteHealthScore

`func (o *SiteScoreTimelineVO) SetSiteHealthScore(v int32)`

SetSiteHealthScore sets SiteHealthScore field to given value.

### HasSiteHealthScore

`func (o *SiteScoreTimelineVO) HasSiteHealthScore() bool`

HasSiteHealthScore returns a boolean if a field has been set.

### GetTime

`func (o *SiteScoreTimelineVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SiteScoreTimelineVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SiteScoreTimelineVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *SiteScoreTimelineVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetWanHealthScore

`func (o *SiteScoreTimelineVO) GetWanHealthScore() int32`

GetWanHealthScore returns the WanHealthScore field if non-nil, zero value otherwise.

### GetWanHealthScoreOk

`func (o *SiteScoreTimelineVO) GetWanHealthScoreOk() (*int32, bool)`

GetWanHealthScoreOk returns a tuple with the WanHealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanHealthScore

`func (o *SiteScoreTimelineVO) SetWanHealthScore(v int32)`

SetWanHealthScore sets WanHealthScore field to given value.

### HasWanHealthScore

`func (o *SiteScoreTimelineVO) HasWanHealthScore() bool`

HasWanHealthScore returns a boolean if a field has been set.

### GetWifiHealthScore

`func (o *SiteScoreTimelineVO) GetWifiHealthScore() int32`

GetWifiHealthScore returns the WifiHealthScore field if non-nil, zero value otherwise.

### GetWifiHealthScoreOk

`func (o *SiteScoreTimelineVO) GetWifiHealthScoreOk() (*int32, bool)`

GetWifiHealthScoreOk returns a tuple with the WifiHealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiHealthScore

`func (o *SiteScoreTimelineVO) SetWifiHealthScore(v int32)`

SetWifiHealthScore sets WifiHealthScore field to given value.

### HasWifiHealthScore

`func (o *SiteScoreTimelineVO) HasWifiHealthScore() bool`

HasWifiHealthScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


