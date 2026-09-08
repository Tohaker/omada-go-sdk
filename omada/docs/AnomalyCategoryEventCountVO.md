# AnomalyCategoryEventCountVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyCode** | Pointer to **string** | For the values of Anomaly event code, refer to section 5.7.2.1 of the Open API Access | [optional] 
**Counts** | Pointer to **int32** | Total count of unresolved and ongoing events | [optional] 
**CurrentStatusCounts** | Pointer to **int32** | Count of events in the current status | [optional] 
**InfluencingClients** | Pointer to **int32** | Number of influenced clients | [optional] 
**InfluencingDevices** | Pointer to **int32** | Number of influenced devices | [optional] 
**Level** | Pointer to **int32** | Level of this incident | [optional] 
**TotalCounts** | Pointer to **int32** | Total count of events in all status | [optional] 

## Methods

### NewAnomalyCategoryEventCountVO

`func NewAnomalyCategoryEventCountVO() *AnomalyCategoryEventCountVO`

NewAnomalyCategoryEventCountVO instantiates a new AnomalyCategoryEventCountVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyCategoryEventCountVOWithDefaults

`func NewAnomalyCategoryEventCountVOWithDefaults() *AnomalyCategoryEventCountVO`

NewAnomalyCategoryEventCountVOWithDefaults instantiates a new AnomalyCategoryEventCountVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyCode

`func (o *AnomalyCategoryEventCountVO) GetAnomalyCode() string`

GetAnomalyCode returns the AnomalyCode field if non-nil, zero value otherwise.

### GetAnomalyCodeOk

`func (o *AnomalyCategoryEventCountVO) GetAnomalyCodeOk() (*string, bool)`

GetAnomalyCodeOk returns a tuple with the AnomalyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyCode

`func (o *AnomalyCategoryEventCountVO) SetAnomalyCode(v string)`

SetAnomalyCode sets AnomalyCode field to given value.

### HasAnomalyCode

`func (o *AnomalyCategoryEventCountVO) HasAnomalyCode() bool`

HasAnomalyCode returns a boolean if a field has been set.

### GetCounts

`func (o *AnomalyCategoryEventCountVO) GetCounts() int32`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *AnomalyCategoryEventCountVO) GetCountsOk() (*int32, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *AnomalyCategoryEventCountVO) SetCounts(v int32)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *AnomalyCategoryEventCountVO) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetCurrentStatusCounts

`func (o *AnomalyCategoryEventCountVO) GetCurrentStatusCounts() int32`

GetCurrentStatusCounts returns the CurrentStatusCounts field if non-nil, zero value otherwise.

### GetCurrentStatusCountsOk

`func (o *AnomalyCategoryEventCountVO) GetCurrentStatusCountsOk() (*int32, bool)`

GetCurrentStatusCountsOk returns a tuple with the CurrentStatusCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatusCounts

`func (o *AnomalyCategoryEventCountVO) SetCurrentStatusCounts(v int32)`

SetCurrentStatusCounts sets CurrentStatusCounts field to given value.

### HasCurrentStatusCounts

`func (o *AnomalyCategoryEventCountVO) HasCurrentStatusCounts() bool`

HasCurrentStatusCounts returns a boolean if a field has been set.

### GetInfluencingClients

`func (o *AnomalyCategoryEventCountVO) GetInfluencingClients() int32`

GetInfluencingClients returns the InfluencingClients field if non-nil, zero value otherwise.

### GetInfluencingClientsOk

`func (o *AnomalyCategoryEventCountVO) GetInfluencingClientsOk() (*int32, bool)`

GetInfluencingClientsOk returns a tuple with the InfluencingClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluencingClients

`func (o *AnomalyCategoryEventCountVO) SetInfluencingClients(v int32)`

SetInfluencingClients sets InfluencingClients field to given value.

### HasInfluencingClients

`func (o *AnomalyCategoryEventCountVO) HasInfluencingClients() bool`

HasInfluencingClients returns a boolean if a field has been set.

### GetInfluencingDevices

`func (o *AnomalyCategoryEventCountVO) GetInfluencingDevices() int32`

GetInfluencingDevices returns the InfluencingDevices field if non-nil, zero value otherwise.

### GetInfluencingDevicesOk

`func (o *AnomalyCategoryEventCountVO) GetInfluencingDevicesOk() (*int32, bool)`

GetInfluencingDevicesOk returns a tuple with the InfluencingDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluencingDevices

`func (o *AnomalyCategoryEventCountVO) SetInfluencingDevices(v int32)`

SetInfluencingDevices sets InfluencingDevices field to given value.

### HasInfluencingDevices

`func (o *AnomalyCategoryEventCountVO) HasInfluencingDevices() bool`

HasInfluencingDevices returns a boolean if a field has been set.

### GetLevel

`func (o *AnomalyCategoryEventCountVO) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AnomalyCategoryEventCountVO) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AnomalyCategoryEventCountVO) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AnomalyCategoryEventCountVO) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetTotalCounts

`func (o *AnomalyCategoryEventCountVO) GetTotalCounts() int32`

GetTotalCounts returns the TotalCounts field if non-nil, zero value otherwise.

### GetTotalCountsOk

`func (o *AnomalyCategoryEventCountVO) GetTotalCountsOk() (*int32, bool)`

GetTotalCountsOk returns a tuple with the TotalCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCounts

`func (o *AnomalyCategoryEventCountVO) SetTotalCounts(v int32)`

SetTotalCounts sets TotalCounts field to given value.

### HasTotalCounts

`func (o *AnomalyCategoryEventCountVO) HasTotalCounts() bool`

HasTotalCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


