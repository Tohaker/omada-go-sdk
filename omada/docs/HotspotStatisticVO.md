# HotspotStatisticVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthTypeDistribution** | Pointer to [**AuthTypeDistributionVO**](AuthTypeDistributionVO.md) |  | [optional] 
**HotspotDistribution** | Pointer to [**HotspotTypeDistributionVO**](HotspotTypeDistributionVO.md) |  | [optional] 
**LastConnection** | Pointer to [**[]AuthedClientNumVO**](AuthedClientNumVO.md) |  | [optional] 

## Methods

### NewHotspotStatisticVO

`func NewHotspotStatisticVO() *HotspotStatisticVO`

NewHotspotStatisticVO instantiates a new HotspotStatisticVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHotspotStatisticVOWithDefaults

`func NewHotspotStatisticVOWithDefaults() *HotspotStatisticVO`

NewHotspotStatisticVOWithDefaults instantiates a new HotspotStatisticVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTypeDistribution

`func (o *HotspotStatisticVO) GetAuthTypeDistribution() AuthTypeDistributionVO`

GetAuthTypeDistribution returns the AuthTypeDistribution field if non-nil, zero value otherwise.

### GetAuthTypeDistributionOk

`func (o *HotspotStatisticVO) GetAuthTypeDistributionOk() (*AuthTypeDistributionVO, bool)`

GetAuthTypeDistributionOk returns a tuple with the AuthTypeDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTypeDistribution

`func (o *HotspotStatisticVO) SetAuthTypeDistribution(v AuthTypeDistributionVO)`

SetAuthTypeDistribution sets AuthTypeDistribution field to given value.

### HasAuthTypeDistribution

`func (o *HotspotStatisticVO) HasAuthTypeDistribution() bool`

HasAuthTypeDistribution returns a boolean if a field has been set.

### GetHotspotDistribution

`func (o *HotspotStatisticVO) GetHotspotDistribution() HotspotTypeDistributionVO`

GetHotspotDistribution returns the HotspotDistribution field if non-nil, zero value otherwise.

### GetHotspotDistributionOk

`func (o *HotspotStatisticVO) GetHotspotDistributionOk() (*HotspotTypeDistributionVO, bool)`

GetHotspotDistributionOk returns a tuple with the HotspotDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotspotDistribution

`func (o *HotspotStatisticVO) SetHotspotDistribution(v HotspotTypeDistributionVO)`

SetHotspotDistribution sets HotspotDistribution field to given value.

### HasHotspotDistribution

`func (o *HotspotStatisticVO) HasHotspotDistribution() bool`

HasHotspotDistribution returns a boolean if a field has been set.

### GetLastConnection

`func (o *HotspotStatisticVO) GetLastConnection() []AuthedClientNumVO`

GetLastConnection returns the LastConnection field if non-nil, zero value otherwise.

### GetLastConnectionOk

`func (o *HotspotStatisticVO) GetLastConnectionOk() (*[]AuthedClientNumVO, bool)`

GetLastConnectionOk returns a tuple with the LastConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnection

`func (o *HotspotStatisticVO) SetLastConnection(v []AuthedClientNumVO)`

SetLastConnection sets LastConnection field to given value.

### HasLastConnection

`func (o *HotspotStatisticVO) HasLastConnection() bool`

HasLastConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


