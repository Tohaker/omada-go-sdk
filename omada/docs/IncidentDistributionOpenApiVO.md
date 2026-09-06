# IncidentDistributionOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncidentCounts** | Pointer to [**[]IncidentCountOpenApiVO**](IncidentCountOpenApiVO.md) | List of incident count items grouped by severity level (critical/error/warning/info) and by function type (category) | [optional] 

## Methods

### NewIncidentDistributionOpenApiVO

`func NewIncidentDistributionOpenApiVO() *IncidentDistributionOpenApiVO`

NewIncidentDistributionOpenApiVO instantiates a new IncidentDistributionOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentDistributionOpenApiVOWithDefaults

`func NewIncidentDistributionOpenApiVOWithDefaults() *IncidentDistributionOpenApiVO`

NewIncidentDistributionOpenApiVOWithDefaults instantiates a new IncidentDistributionOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidentCounts

`func (o *IncidentDistributionOpenApiVO) GetIncidentCounts() []IncidentCountOpenApiVO`

GetIncidentCounts returns the IncidentCounts field if non-nil, zero value otherwise.

### GetIncidentCountsOk

`func (o *IncidentDistributionOpenApiVO) GetIncidentCountsOk() (*[]IncidentCountOpenApiVO, bool)`

GetIncidentCountsOk returns a tuple with the IncidentCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentCounts

`func (o *IncidentDistributionOpenApiVO) SetIncidentCounts(v []IncidentCountOpenApiVO)`

SetIncidentCounts sets IncidentCounts field to given value.

### HasIncidentCounts

`func (o *IncidentDistributionOpenApiVO) HasIncidentCounts() bool`

HasIncidentCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


