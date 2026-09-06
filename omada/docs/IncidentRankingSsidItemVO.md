# IncidentRankingSsidItemVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incidents** | Pointer to **int32** | Number of incidents associated with this SSID | [optional] 
**SsidName** | Pointer to **string** | SSID name | [optional] 

## Methods

### NewIncidentRankingSsidItemVO

`func NewIncidentRankingSsidItemVO() *IncidentRankingSsidItemVO`

NewIncidentRankingSsidItemVO instantiates a new IncidentRankingSsidItemVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentRankingSsidItemVOWithDefaults

`func NewIncidentRankingSsidItemVOWithDefaults() *IncidentRankingSsidItemVO`

NewIncidentRankingSsidItemVOWithDefaults instantiates a new IncidentRankingSsidItemVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidents

`func (o *IncidentRankingSsidItemVO) GetIncidents() int32`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *IncidentRankingSsidItemVO) GetIncidentsOk() (*int32, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *IncidentRankingSsidItemVO) SetIncidents(v int32)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *IncidentRankingSsidItemVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetSsidName

`func (o *IncidentRankingSsidItemVO) GetSsidName() string`

GetSsidName returns the SsidName field if non-nil, zero value otherwise.

### GetSsidNameOk

`func (o *IncidentRankingSsidItemVO) GetSsidNameOk() (*string, bool)`

GetSsidNameOk returns a tuple with the SsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidName

`func (o *IncidentRankingSsidItemVO) SetSsidName(v string)`

SetSsidName sets SsidName field to given value.

### HasSsidName

`func (o *IncidentRankingSsidItemVO) HasSsidName() bool`

HasSsidName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


