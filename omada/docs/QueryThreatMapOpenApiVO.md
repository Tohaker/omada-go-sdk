# QueryThreatMapOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | The country that was attacked. | [optional] 
**End** | Pointer to **int64** | The timestamp for the end time is in milliseconds. | [optional] 
**Severity** | Pointer to **string** | Severity. | [optional] 
**Sites** | Pointer to **string** | Comma separated sites should be explained. If it is not passed, select all. | [optional] 
**Start** | Pointer to **int64** | The timestamp for the start time is in milliseconds. | [optional] 

## Methods

### NewQueryThreatMapOpenApiVO

`func NewQueryThreatMapOpenApiVO() *QueryThreatMapOpenApiVO`

NewQueryThreatMapOpenApiVO instantiates a new QueryThreatMapOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryThreatMapOpenApiVOWithDefaults

`func NewQueryThreatMapOpenApiVOWithDefaults() *QueryThreatMapOpenApiVO`

NewQueryThreatMapOpenApiVOWithDefaults instantiates a new QueryThreatMapOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *QueryThreatMapOpenApiVO) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *QueryThreatMapOpenApiVO) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *QueryThreatMapOpenApiVO) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *QueryThreatMapOpenApiVO) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEnd

`func (o *QueryThreatMapOpenApiVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *QueryThreatMapOpenApiVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *QueryThreatMapOpenApiVO) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *QueryThreatMapOpenApiVO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetSeverity

`func (o *QueryThreatMapOpenApiVO) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *QueryThreatMapOpenApiVO) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *QueryThreatMapOpenApiVO) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *QueryThreatMapOpenApiVO) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSites

`func (o *QueryThreatMapOpenApiVO) GetSites() string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *QueryThreatMapOpenApiVO) GetSitesOk() (*string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *QueryThreatMapOpenApiVO) SetSites(v string)`

SetSites sets Sites field to given value.

### HasSites

`func (o *QueryThreatMapOpenApiVO) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetStart

`func (o *QueryThreatMapOpenApiVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *QueryThreatMapOpenApiVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *QueryThreatMapOpenApiVO) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *QueryThreatMapOpenApiVO) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


