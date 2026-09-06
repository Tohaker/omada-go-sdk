# Incident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncidentCounts** | Pointer to **int32** | Incident counts. | [optional] 
**Mac** | Pointer to **string** | Device MAC address. | [optional] 

## Methods

### NewIncident

`func NewIncident() *Incident`

NewIncident instantiates a new Incident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentWithDefaults

`func NewIncidentWithDefaults() *Incident`

NewIncidentWithDefaults instantiates a new Incident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidentCounts

`func (o *Incident) GetIncidentCounts() int32`

GetIncidentCounts returns the IncidentCounts field if non-nil, zero value otherwise.

### GetIncidentCountsOk

`func (o *Incident) GetIncidentCountsOk() (*int32, bool)`

GetIncidentCountsOk returns a tuple with the IncidentCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentCounts

`func (o *Incident) SetIncidentCounts(v int32)`

SetIncidentCounts sets IncidentCounts field to given value.

### HasIncidentCounts

`func (o *Incident) HasIncidentCounts() bool`

HasIncidentCounts returns a boolean if a field has been set.

### GetMac

`func (o *Incident) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Incident) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Incident) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Incident) HasMac() bool`

HasMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


