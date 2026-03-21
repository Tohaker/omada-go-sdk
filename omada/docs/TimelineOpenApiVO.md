# TimelineOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to [**[]ConnectPeriodVO**](ConnectPeriodVO.md) | The collection of device online intervals. | [optional] 

## Methods

### NewTimelineOpenApiVO

`func NewTimelineOpenApiVO() *TimelineOpenApiVO`

NewTimelineOpenApiVO instantiates a new TimelineOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineOpenApiVOWithDefaults

`func NewTimelineOpenApiVOWithDefaults() *TimelineOpenApiVO`

NewTimelineOpenApiVOWithDefaults instantiates a new TimelineOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *TimelineOpenApiVO) GetConnections() []ConnectPeriodVO`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *TimelineOpenApiVO) GetConnectionsOk() (*[]ConnectPeriodVO, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *TimelineOpenApiVO) SetConnections(v []ConnectPeriodVO)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *TimelineOpenApiVO) HasConnections() bool`

HasConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


