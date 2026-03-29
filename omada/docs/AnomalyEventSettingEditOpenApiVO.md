# AnomalyEventSettingEditOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyCode** | Pointer to **string** | For the values of Anomaly event code, refer to section 5.7.2.1 of the Open API Access | [optional] 
**Enable** | Pointer to **bool** | Whether to detect anomaly events | [optional] 
**Level** | Pointer to **int32** | Anomaly event level, it should be a value as follows: 0:Critical, 1:Error, 2:Warning, 3: Info | [optional] 
**Params** | Pointer to **map[string]int32** | For the values of Anomaly event params, refer to section 5.7.2.1 of the Open API Access | [optional] 

## Methods

### NewAnomalyEventSettingEditOpenApiVO

`func NewAnomalyEventSettingEditOpenApiVO() *AnomalyEventSettingEditOpenApiVO`

NewAnomalyEventSettingEditOpenApiVO instantiates a new AnomalyEventSettingEditOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyEventSettingEditOpenApiVOWithDefaults

`func NewAnomalyEventSettingEditOpenApiVOWithDefaults() *AnomalyEventSettingEditOpenApiVO`

NewAnomalyEventSettingEditOpenApiVOWithDefaults instantiates a new AnomalyEventSettingEditOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyCode

`func (o *AnomalyEventSettingEditOpenApiVO) GetAnomalyCode() string`

GetAnomalyCode returns the AnomalyCode field if non-nil, zero value otherwise.

### GetAnomalyCodeOk

`func (o *AnomalyEventSettingEditOpenApiVO) GetAnomalyCodeOk() (*string, bool)`

GetAnomalyCodeOk returns a tuple with the AnomalyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyCode

`func (o *AnomalyEventSettingEditOpenApiVO) SetAnomalyCode(v string)`

SetAnomalyCode sets AnomalyCode field to given value.

### HasAnomalyCode

`func (o *AnomalyEventSettingEditOpenApiVO) HasAnomalyCode() bool`

HasAnomalyCode returns a boolean if a field has been set.

### GetEnable

`func (o *AnomalyEventSettingEditOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AnomalyEventSettingEditOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AnomalyEventSettingEditOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AnomalyEventSettingEditOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetLevel

`func (o *AnomalyEventSettingEditOpenApiVO) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AnomalyEventSettingEditOpenApiVO) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AnomalyEventSettingEditOpenApiVO) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AnomalyEventSettingEditOpenApiVO) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetParams

`func (o *AnomalyEventSettingEditOpenApiVO) GetParams() map[string]int32`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *AnomalyEventSettingEditOpenApiVO) GetParamsOk() (*map[string]int32, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *AnomalyEventSettingEditOpenApiVO) SetParams(v map[string]int32)`

SetParams sets Params field to given value.

### HasParams

`func (o *AnomalyEventSettingEditOpenApiVO) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


