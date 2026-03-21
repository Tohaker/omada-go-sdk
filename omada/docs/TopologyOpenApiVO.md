# TopologyOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateTime** | Pointer to **int64** | Last Update Time | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**TopologyEdges** | Pointer to [**[]TopologyOpenApiEdgeVO**](TopologyOpenApiEdgeVO.md) | Topology Edges | [optional] 
**TopologyNodes** | Pointer to [**[]TopologyOpenApiNodeVO**](TopologyOpenApiNodeVO.md) | Topology Nodes | [optional] 

## Methods

### NewTopologyOpenApiVO

`func NewTopologyOpenApiVO() *TopologyOpenApiVO`

NewTopologyOpenApiVO instantiates a new TopologyOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyOpenApiVOWithDefaults

`func NewTopologyOpenApiVOWithDefaults() *TopologyOpenApiVO`

NewTopologyOpenApiVOWithDefaults instantiates a new TopologyOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateTime

`func (o *TopologyOpenApiVO) GetLastUpdateTime() int64`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *TopologyOpenApiVO) GetLastUpdateTimeOk() (*int64, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *TopologyOpenApiVO) SetLastUpdateTime(v int64)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *TopologyOpenApiVO) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetSiteId

`func (o *TopologyOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *TopologyOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *TopologyOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *TopologyOpenApiVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTopologyEdges

`func (o *TopologyOpenApiVO) GetTopologyEdges() []TopologyOpenApiEdgeVO`

GetTopologyEdges returns the TopologyEdges field if non-nil, zero value otherwise.

### GetTopologyEdgesOk

`func (o *TopologyOpenApiVO) GetTopologyEdgesOk() (*[]TopologyOpenApiEdgeVO, bool)`

GetTopologyEdgesOk returns a tuple with the TopologyEdges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyEdges

`func (o *TopologyOpenApiVO) SetTopologyEdges(v []TopologyOpenApiEdgeVO)`

SetTopologyEdges sets TopologyEdges field to given value.

### HasTopologyEdges

`func (o *TopologyOpenApiVO) HasTopologyEdges() bool`

HasTopologyEdges returns a boolean if a field has been set.

### GetTopologyNodes

`func (o *TopologyOpenApiVO) GetTopologyNodes() []TopologyOpenApiNodeVO`

GetTopologyNodes returns the TopologyNodes field if non-nil, zero value otherwise.

### GetTopologyNodesOk

`func (o *TopologyOpenApiVO) GetTopologyNodesOk() (*[]TopologyOpenApiNodeVO, bool)`

GetTopologyNodesOk returns a tuple with the TopologyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyNodes

`func (o *TopologyOpenApiVO) SetTopologyNodes(v []TopologyOpenApiNodeVO)`

SetTopologyNodes sets TopologyNodes field to given value.

### HasTopologyNodes

`func (o *TopologyOpenApiVO) HasTopologyNodes() bool`

HasTopologyNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


