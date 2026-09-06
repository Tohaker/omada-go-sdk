# ClientIncidentCountResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incidents** | Pointer to [**[]Incident**](Incident.md) | Incident counts for the requested clients. | [optional] 

## Methods

### NewClientIncidentCountResultOpenApiVO

`func NewClientIncidentCountResultOpenApiVO() *ClientIncidentCountResultOpenApiVO`

NewClientIncidentCountResultOpenApiVO instantiates a new ClientIncidentCountResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientIncidentCountResultOpenApiVOWithDefaults

`func NewClientIncidentCountResultOpenApiVOWithDefaults() *ClientIncidentCountResultOpenApiVO`

NewClientIncidentCountResultOpenApiVOWithDefaults instantiates a new ClientIncidentCountResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidents

`func (o *ClientIncidentCountResultOpenApiVO) GetIncidents() []Incident`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *ClientIncidentCountResultOpenApiVO) GetIncidentsOk() (*[]Incident, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *ClientIncidentCountResultOpenApiVO) SetIncidents(v []Incident)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *ClientIncidentCountResultOpenApiVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


