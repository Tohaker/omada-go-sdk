# ThreatCategoryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **[]int64** | Number of threats,List structure : [Low, Minor, Moderate, major, critical]. | [optional] 
**Coords** | Pointer to **[]float64** | Longitude, latitude. | [optional] 
**Country** | Pointer to **string** | Country. | [optional] 

## Methods

### NewThreatCategoryOpenApiVO

`func NewThreatCategoryOpenApiVO() *ThreatCategoryOpenApiVO`

NewThreatCategoryOpenApiVO instantiates a new ThreatCategoryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatCategoryOpenApiVOWithDefaults

`func NewThreatCategoryOpenApiVOWithDefaults() *ThreatCategoryOpenApiVO`

NewThreatCategoryOpenApiVOWithDefaults instantiates a new ThreatCategoryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *ThreatCategoryOpenApiVO) GetAttempts() []int64`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *ThreatCategoryOpenApiVO) GetAttemptsOk() (*[]int64, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *ThreatCategoryOpenApiVO) SetAttempts(v []int64)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *ThreatCategoryOpenApiVO) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetCoords

`func (o *ThreatCategoryOpenApiVO) GetCoords() []float64`

GetCoords returns the Coords field if non-nil, zero value otherwise.

### GetCoordsOk

`func (o *ThreatCategoryOpenApiVO) GetCoordsOk() (*[]float64, bool)`

GetCoordsOk returns a tuple with the Coords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoords

`func (o *ThreatCategoryOpenApiVO) SetCoords(v []float64)`

SetCoords sets Coords field to given value.

### HasCoords

`func (o *ThreatCategoryOpenApiVO) HasCoords() bool`

HasCoords returns a boolean if a field has been set.

### GetCountry

`func (o *ThreatCategoryOpenApiVO) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ThreatCategoryOpenApiVO) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ThreatCategoryOpenApiVO) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ThreatCategoryOpenApiVO) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


