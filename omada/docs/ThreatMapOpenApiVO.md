# ThreatMapOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int64** | Number of threats. | [optional] 
**Coords** | Pointer to **[]float64** | Longitude, latitude. | [optional] 
**Country** | Pointer to **string** | Country. | [optional] 

## Methods

### NewThreatMapOpenApiVO

`func NewThreatMapOpenApiVO() *ThreatMapOpenApiVO`

NewThreatMapOpenApiVO instantiates a new ThreatMapOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatMapOpenApiVOWithDefaults

`func NewThreatMapOpenApiVOWithDefaults() *ThreatMapOpenApiVO`

NewThreatMapOpenApiVOWithDefaults instantiates a new ThreatMapOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *ThreatMapOpenApiVO) GetAttempts() int64`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *ThreatMapOpenApiVO) GetAttemptsOk() (*int64, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *ThreatMapOpenApiVO) SetAttempts(v int64)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *ThreatMapOpenApiVO) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetCoords

`func (o *ThreatMapOpenApiVO) GetCoords() []float64`

GetCoords returns the Coords field if non-nil, zero value otherwise.

### GetCoordsOk

`func (o *ThreatMapOpenApiVO) GetCoordsOk() (*[]float64, bool)`

GetCoordsOk returns a tuple with the Coords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoords

`func (o *ThreatMapOpenApiVO) SetCoords(v []float64)`

SetCoords sets Coords field to given value.

### HasCoords

`func (o *ThreatMapOpenApiVO) HasCoords() bool`

HasCoords returns a boolean if a field has been set.

### GetCountry

`func (o *ThreatMapOpenApiVO) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ThreatMapOpenApiVO) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ThreatMapOpenApiVO) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ThreatMapOpenApiVO) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


