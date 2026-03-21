# GeoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int64** | Attempts. | [optional] 
**Country** | Pointer to **string** | Country code. | [optional] 
**Source** | Pointer to **string** | Source Ip, shown as Multiple when there are more than one ip. | [optional] 

## Methods

### NewGeoOpenApiVO

`func NewGeoOpenApiVO() *GeoOpenApiVO`

NewGeoOpenApiVO instantiates a new GeoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoOpenApiVOWithDefaults

`func NewGeoOpenApiVOWithDefaults() *GeoOpenApiVO`

NewGeoOpenApiVOWithDefaults instantiates a new GeoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *GeoOpenApiVO) GetAttempts() int64`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *GeoOpenApiVO) GetAttemptsOk() (*int64, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *GeoOpenApiVO) SetAttempts(v int64)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *GeoOpenApiVO) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetCountry

`func (o *GeoOpenApiVO) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GeoOpenApiVO) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GeoOpenApiVO) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GeoOpenApiVO) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetSource

`func (o *GeoOpenApiVO) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GeoOpenApiVO) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GeoOpenApiVO) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GeoOpenApiVO) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


