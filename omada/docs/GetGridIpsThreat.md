# GetGridIpsThreat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **int32** | IPS threat category | [optional] 
**Classification** | Pointer to **string** | Ips threat classification | [optional] 
**Id** | Pointer to **string** | IPS threat ID | [optional] 
**Service** | Pointer to **string** | IPS threat description | [optional] 
**Severity** | Pointer to **int32** | IPS threat severity should be a value as follows: 0: Critical; 1: Major; 2: Moderate; 3: Minor; 4: Low | [optional] 
**SrcCountry** | Pointer to **string** | IPS threat source Country | [optional] 
**Time** | Pointer to **int64** | Timestamp, in seconds, such as 1682000000 | [optional] 

## Methods

### NewGetGridIpsThreat

`func NewGetGridIpsThreat() *GetGridIpsThreat`

NewGetGridIpsThreat instantiates a new GetGridIpsThreat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridIpsThreatWithDefaults

`func NewGetGridIpsThreatWithDefaults() *GetGridIpsThreat`

NewGetGridIpsThreatWithDefaults instantiates a new GetGridIpsThreat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *GetGridIpsThreat) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetGridIpsThreat) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetGridIpsThreat) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetGridIpsThreat) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClassification

`func (o *GetGridIpsThreat) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *GetGridIpsThreat) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *GetGridIpsThreat) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *GetGridIpsThreat) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetId

`func (o *GetGridIpsThreat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGridIpsThreat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGridIpsThreat) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetGridIpsThreat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetService

`func (o *GetGridIpsThreat) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetGridIpsThreat) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetGridIpsThreat) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetGridIpsThreat) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSeverity

`func (o *GetGridIpsThreat) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetGridIpsThreat) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetGridIpsThreat) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetGridIpsThreat) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSrcCountry

`func (o *GetGridIpsThreat) GetSrcCountry() string`

GetSrcCountry returns the SrcCountry field if non-nil, zero value otherwise.

### GetSrcCountryOk

`func (o *GetGridIpsThreat) GetSrcCountryOk() (*string, bool)`

GetSrcCountryOk returns a tuple with the SrcCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCountry

`func (o *GetGridIpsThreat) SetSrcCountry(v string)`

SetSrcCountry sets SrcCountry field to given value.

### HasSrcCountry

`func (o *GetGridIpsThreat) HasSrcCountry() bool`

HasSrcCountry returns a boolean if a field has been set.

### GetTime

`func (o *GetGridIpsThreat) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetGridIpsThreat) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetGridIpsThreat) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetGridIpsThreat) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


