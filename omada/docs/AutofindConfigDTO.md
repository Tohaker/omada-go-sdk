# AutofindConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgingTime** | Pointer to **int32** | Configure the actions corresponding to received DDM messages on the optical port.If aging time status is &#39;No-Aging&#39;,AgingTime should be null;If aging time status is &#39;Timeout&#39;,AgingTime should not be null and should be within the range of 100 to 300s. | [optional] 
**AgingTimeStatus** | **string** | Configure the DDM (Digital Diagnostic Monitoring) feature enable status for the optical port.AgingTimeStatus should be a value as follows:TIMEOUT,NO_AGING | 
**AutofindInterval** | **int32** | Auto find interval should be within the range of 1 to 10s | 
**PonPort** | Pointer to **string** | Pon port collection.e.g.,&#39;GPON 1/1/1&#39;、&#39;GPON 1/1/1-3&#39;、&#39;GPON 1/1/1-3,GPON 1/1/5,GPON 1/1/7-8&#39; | [optional] 

## Methods

### NewAutofindConfigDTO

`func NewAutofindConfigDTO(agingTimeStatus string, autofindInterval int32, ) *AutofindConfigDTO`

NewAutofindConfigDTO instantiates a new AutofindConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutofindConfigDTOWithDefaults

`func NewAutofindConfigDTOWithDefaults() *AutofindConfigDTO`

NewAutofindConfigDTOWithDefaults instantiates a new AutofindConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgingTime

`func (o *AutofindConfigDTO) GetAgingTime() int32`

GetAgingTime returns the AgingTime field if non-nil, zero value otherwise.

### GetAgingTimeOk

`func (o *AutofindConfigDTO) GetAgingTimeOk() (*int32, bool)`

GetAgingTimeOk returns a tuple with the AgingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgingTime

`func (o *AutofindConfigDTO) SetAgingTime(v int32)`

SetAgingTime sets AgingTime field to given value.

### HasAgingTime

`func (o *AutofindConfigDTO) HasAgingTime() bool`

HasAgingTime returns a boolean if a field has been set.

### GetAgingTimeStatus

`func (o *AutofindConfigDTO) GetAgingTimeStatus() string`

GetAgingTimeStatus returns the AgingTimeStatus field if non-nil, zero value otherwise.

### GetAgingTimeStatusOk

`func (o *AutofindConfigDTO) GetAgingTimeStatusOk() (*string, bool)`

GetAgingTimeStatusOk returns a tuple with the AgingTimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgingTimeStatus

`func (o *AutofindConfigDTO) SetAgingTimeStatus(v string)`

SetAgingTimeStatus sets AgingTimeStatus field to given value.


### GetAutofindInterval

`func (o *AutofindConfigDTO) GetAutofindInterval() int32`

GetAutofindInterval returns the AutofindInterval field if non-nil, zero value otherwise.

### GetAutofindIntervalOk

`func (o *AutofindConfigDTO) GetAutofindIntervalOk() (*int32, bool)`

GetAutofindIntervalOk returns a tuple with the AutofindInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofindInterval

`func (o *AutofindConfigDTO) SetAutofindInterval(v int32)`

SetAutofindInterval sets AutofindInterval field to given value.


### GetPonPort

`func (o *AutofindConfigDTO) GetPonPort() string`

GetPonPort returns the PonPort field if non-nil, zero value otherwise.

### GetPonPortOk

`func (o *AutofindConfigDTO) GetPonPortOk() (*string, bool)`

GetPonPortOk returns a tuple with the PonPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPort

`func (o *AutofindConfigDTO) SetPonPort(v string)`

SetPonPort sets PonPort field to given value.

### HasPonPort

`func (o *AutofindConfigDTO) HasPonPort() bool`

HasPonPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


