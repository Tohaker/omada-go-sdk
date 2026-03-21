# TopologyV3OpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthStat** | Pointer to [**TopologyOpenApiHealthStatusVO**](TopologyOpenApiHealthStatusVO.md) |  | [optional] 
**LastUpdateTime** | Pointer to **int64** | Last Update Time | [optional] 
**Status** | Pointer to [**TopologyOpenApiStatusVO**](TopologyOpenApiStatusVO.md) |  | [optional] 
**TopologyArray** | Pointer to [**[]TopologyV3OpenApiNodeVO**](TopologyV3OpenApiNodeVO.md) | Topology Nodes | [optional] 

## Methods

### NewTopologyV3OpenApiVO

`func NewTopologyV3OpenApiVO() *TopologyV3OpenApiVO`

NewTopologyV3OpenApiVO instantiates a new TopologyV3OpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyV3OpenApiVOWithDefaults

`func NewTopologyV3OpenApiVOWithDefaults() *TopologyV3OpenApiVO`

NewTopologyV3OpenApiVOWithDefaults instantiates a new TopologyV3OpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthStat

`func (o *TopologyV3OpenApiVO) GetHealthStat() TopologyOpenApiHealthStatusVO`

GetHealthStat returns the HealthStat field if non-nil, zero value otherwise.

### GetHealthStatOk

`func (o *TopologyV3OpenApiVO) GetHealthStatOk() (*TopologyOpenApiHealthStatusVO, bool)`

GetHealthStatOk returns a tuple with the HealthStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStat

`func (o *TopologyV3OpenApiVO) SetHealthStat(v TopologyOpenApiHealthStatusVO)`

SetHealthStat sets HealthStat field to given value.

### HasHealthStat

`func (o *TopologyV3OpenApiVO) HasHealthStat() bool`

HasHealthStat returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *TopologyV3OpenApiVO) GetLastUpdateTime() int64`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *TopologyV3OpenApiVO) GetLastUpdateTimeOk() (*int64, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *TopologyV3OpenApiVO) SetLastUpdateTime(v int64)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *TopologyV3OpenApiVO) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetStatus

`func (o *TopologyV3OpenApiVO) GetStatus() TopologyOpenApiStatusVO`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TopologyV3OpenApiVO) GetStatusOk() (*TopologyOpenApiStatusVO, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TopologyV3OpenApiVO) SetStatus(v TopologyOpenApiStatusVO)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TopologyV3OpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTopologyArray

`func (o *TopologyV3OpenApiVO) GetTopologyArray() []TopologyV3OpenApiNodeVO`

GetTopologyArray returns the TopologyArray field if non-nil, zero value otherwise.

### GetTopologyArrayOk

`func (o *TopologyV3OpenApiVO) GetTopologyArrayOk() (*[]TopologyV3OpenApiNodeVO, bool)`

GetTopologyArrayOk returns a tuple with the TopologyArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyArray

`func (o *TopologyV3OpenApiVO) SetTopologyArray(v []TopologyV3OpenApiNodeVO)`

SetTopologyArray sets TopologyArray field to given value.

### HasTopologyArray

`func (o *TopologyV3OpenApiVO) HasTopologyArray() bool`

HasTopologyArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


