# SiteSubHealthScoreVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientHealthScoreDetail** | Pointer to [**HealthStatisticsScoreVO**](HealthStatisticsScoreVO.md) |  | [optional] 
**DeviceHealthScoreDetail** | Pointer to [**SiteDeviceSubHealthScoreVO**](SiteDeviceSubHealthScoreVO.md) |  | [optional] 
**WanScoreDetail** | Pointer to [**WanHealthListVO**](WanHealthListVO.md) |  | [optional] 

## Methods

### NewSiteSubHealthScoreVO

`func NewSiteSubHealthScoreVO() *SiteSubHealthScoreVO`

NewSiteSubHealthScoreVO instantiates a new SiteSubHealthScoreVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSubHealthScoreVOWithDefaults

`func NewSiteSubHealthScoreVOWithDefaults() *SiteSubHealthScoreVO`

NewSiteSubHealthScoreVOWithDefaults instantiates a new SiteSubHealthScoreVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientHealthScoreDetail

`func (o *SiteSubHealthScoreVO) GetClientHealthScoreDetail() HealthStatisticsScoreVO`

GetClientHealthScoreDetail returns the ClientHealthScoreDetail field if non-nil, zero value otherwise.

### GetClientHealthScoreDetailOk

`func (o *SiteSubHealthScoreVO) GetClientHealthScoreDetailOk() (*HealthStatisticsScoreVO, bool)`

GetClientHealthScoreDetailOk returns a tuple with the ClientHealthScoreDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHealthScoreDetail

`func (o *SiteSubHealthScoreVO) SetClientHealthScoreDetail(v HealthStatisticsScoreVO)`

SetClientHealthScoreDetail sets ClientHealthScoreDetail field to given value.

### HasClientHealthScoreDetail

`func (o *SiteSubHealthScoreVO) HasClientHealthScoreDetail() bool`

HasClientHealthScoreDetail returns a boolean if a field has been set.

### GetDeviceHealthScoreDetail

`func (o *SiteSubHealthScoreVO) GetDeviceHealthScoreDetail() SiteDeviceSubHealthScoreVO`

GetDeviceHealthScoreDetail returns the DeviceHealthScoreDetail field if non-nil, zero value otherwise.

### GetDeviceHealthScoreDetailOk

`func (o *SiteSubHealthScoreVO) GetDeviceHealthScoreDetailOk() (*SiteDeviceSubHealthScoreVO, bool)`

GetDeviceHealthScoreDetailOk returns a tuple with the DeviceHealthScoreDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHealthScoreDetail

`func (o *SiteSubHealthScoreVO) SetDeviceHealthScoreDetail(v SiteDeviceSubHealthScoreVO)`

SetDeviceHealthScoreDetail sets DeviceHealthScoreDetail field to given value.

### HasDeviceHealthScoreDetail

`func (o *SiteSubHealthScoreVO) HasDeviceHealthScoreDetail() bool`

HasDeviceHealthScoreDetail returns a boolean if a field has been set.

### GetWanScoreDetail

`func (o *SiteSubHealthScoreVO) GetWanScoreDetail() WanHealthListVO`

GetWanScoreDetail returns the WanScoreDetail field if non-nil, zero value otherwise.

### GetWanScoreDetailOk

`func (o *SiteSubHealthScoreVO) GetWanScoreDetailOk() (*WanHealthListVO, bool)`

GetWanScoreDetailOk returns a tuple with the WanScoreDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanScoreDetail

`func (o *SiteSubHealthScoreVO) SetWanScoreDetail(v WanHealthListVO)`

SetWanScoreDetail sets WanScoreDetail field to given value.

### HasWanScoreDetail

`func (o *SiteSubHealthScoreVO) HasWanScoreDetail() bool`

HasWanScoreDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


