# IncidentRankingBandItemVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band** | Pointer to **int32** | Frequency band identifier (e.g. 0&#x3D;2.4GHz, 1&#x3D;5GHz, 2&#x3D;5GHz-2, 3&#x3D;6GHz) | [optional] 
**Incidents** | Pointer to **int32** | Number of incidents associated with this frequency band | [optional] 

## Methods

### NewIncidentRankingBandItemVO

`func NewIncidentRankingBandItemVO() *IncidentRankingBandItemVO`

NewIncidentRankingBandItemVO instantiates a new IncidentRankingBandItemVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentRankingBandItemVOWithDefaults

`func NewIncidentRankingBandItemVOWithDefaults() *IncidentRankingBandItemVO`

NewIncidentRankingBandItemVOWithDefaults instantiates a new IncidentRankingBandItemVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand

`func (o *IncidentRankingBandItemVO) GetBand() int32`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *IncidentRankingBandItemVO) GetBandOk() (*int32, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *IncidentRankingBandItemVO) SetBand(v int32)`

SetBand sets Band field to given value.

### HasBand

`func (o *IncidentRankingBandItemVO) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetIncidents

`func (o *IncidentRankingBandItemVO) GetIncidents() int32`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *IncidentRankingBandItemVO) GetIncidentsOk() (*int32, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *IncidentRankingBandItemVO) SetIncidents(v int32)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *IncidentRankingBandItemVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


