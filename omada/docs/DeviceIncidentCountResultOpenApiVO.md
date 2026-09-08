# DeviceIncidentCountResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incidents** | Pointer to [**[]Incident**](Incident.md) | Incident counts for the requested devices. | [optional] 

## Methods

### NewDeviceIncidentCountResultOpenApiVO

`func NewDeviceIncidentCountResultOpenApiVO() *DeviceIncidentCountResultOpenApiVO`

NewDeviceIncidentCountResultOpenApiVO instantiates a new DeviceIncidentCountResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceIncidentCountResultOpenApiVOWithDefaults

`func NewDeviceIncidentCountResultOpenApiVOWithDefaults() *DeviceIncidentCountResultOpenApiVO`

NewDeviceIncidentCountResultOpenApiVOWithDefaults instantiates a new DeviceIncidentCountResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidents

`func (o *DeviceIncidentCountResultOpenApiVO) GetIncidents() []Incident`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *DeviceIncidentCountResultOpenApiVO) GetIncidentsOk() (*[]Incident, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *DeviceIncidentCountResultOpenApiVO) SetIncidents(v []Incident)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *DeviceIncidentCountResultOpenApiVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


