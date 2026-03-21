# AnomalyEventSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyCode** | Pointer to **string** | For the values of Anomaly event code, refer to section 5.7.2.1 of the Open API Access | [optional] 
**Category** | Pointer to **int32** | Anomaly event category, it should be a value as follows: 0:Networking, 1:Mesh, 2:Access, 3:Roaming, 4:Network Service, 5:Software/Configuration, 6:Hardware, 7:Security, 8:Throughput, 9:Coverage | [optional] 
**Enable** | Pointer to **bool** | Whether to detect anomaly events | [optional] 
**Level** | Pointer to **int32** | Anomaly event level, it should be a value as follows: 0:Critical, 1:Error, 2:Warning, 3: Info | [optional] 
**Params** | Pointer to **map[string]int32** | For the values of Anomaly event params, refer to section 5.7.2.1 of the Open API Access | [optional] 

## Methods

### NewAnomalyEventSettingOpenApiVO

`func NewAnomalyEventSettingOpenApiVO() *AnomalyEventSettingOpenApiVO`

NewAnomalyEventSettingOpenApiVO instantiates a new AnomalyEventSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyEventSettingOpenApiVOWithDefaults

`func NewAnomalyEventSettingOpenApiVOWithDefaults() *AnomalyEventSettingOpenApiVO`

NewAnomalyEventSettingOpenApiVOWithDefaults instantiates a new AnomalyEventSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyCode

`func (o *AnomalyEventSettingOpenApiVO) GetAnomalyCode() string`

GetAnomalyCode returns the AnomalyCode field if non-nil, zero value otherwise.

### GetAnomalyCodeOk

`func (o *AnomalyEventSettingOpenApiVO) GetAnomalyCodeOk() (*string, bool)`

GetAnomalyCodeOk returns a tuple with the AnomalyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyCode

`func (o *AnomalyEventSettingOpenApiVO) SetAnomalyCode(v string)`

SetAnomalyCode sets AnomalyCode field to given value.

### HasAnomalyCode

`func (o *AnomalyEventSettingOpenApiVO) HasAnomalyCode() bool`

HasAnomalyCode returns a boolean if a field has been set.

### GetCategory

`func (o *AnomalyEventSettingOpenApiVO) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AnomalyEventSettingOpenApiVO) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AnomalyEventSettingOpenApiVO) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AnomalyEventSettingOpenApiVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnable

`func (o *AnomalyEventSettingOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AnomalyEventSettingOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AnomalyEventSettingOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AnomalyEventSettingOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetLevel

`func (o *AnomalyEventSettingOpenApiVO) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AnomalyEventSettingOpenApiVO) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AnomalyEventSettingOpenApiVO) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AnomalyEventSettingOpenApiVO) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetParams

`func (o *AnomalyEventSettingOpenApiVO) GetParams() map[string]int32`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *AnomalyEventSettingOpenApiVO) GetParamsOk() (*map[string]int32, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *AnomalyEventSettingOpenApiVO) SetParams(v map[string]int32)`

SetParams sets Params field to given value.

### HasParams

`func (o *AnomalyEventSettingOpenApiVO) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


