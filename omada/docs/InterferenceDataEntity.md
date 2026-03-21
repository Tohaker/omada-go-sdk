# InterferenceDataEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cnt** | Pointer to **[]int32** | interference counts, range [0, 200] | [optional] 
**Str** | Pointer to **[]int32** | Interference strength, range [-91, -89, -87, …., -33, -31, -29] (up to 32 strength levels), for example: -91 means the range is -92-90. | [optional] 

## Methods

### NewInterferenceDataEntity

`func NewInterferenceDataEntity() *InterferenceDataEntity`

NewInterferenceDataEntity instantiates a new InterferenceDataEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterferenceDataEntityWithDefaults

`func NewInterferenceDataEntityWithDefaults() *InterferenceDataEntity`

NewInterferenceDataEntityWithDefaults instantiates a new InterferenceDataEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCnt

`func (o *InterferenceDataEntity) GetCnt() []int32`

GetCnt returns the Cnt field if non-nil, zero value otherwise.

### GetCntOk

`func (o *InterferenceDataEntity) GetCntOk() (*[]int32, bool)`

GetCntOk returns a tuple with the Cnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnt

`func (o *InterferenceDataEntity) SetCnt(v []int32)`

SetCnt sets Cnt field to given value.

### HasCnt

`func (o *InterferenceDataEntity) HasCnt() bool`

HasCnt returns a boolean if a field has been set.

### GetStr

`func (o *InterferenceDataEntity) GetStr() []int32`

GetStr returns the Str field if non-nil, zero value otherwise.

### GetStrOk

`func (o *InterferenceDataEntity) GetStrOk() (*[]int32, bool)`

GetStrOk returns a tuple with the Str field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStr

`func (o *InterferenceDataEntity) SetStr(v []int32)`

SetStr sets Str field to given value.

### HasStr

`func (o *InterferenceDataEntity) HasStr() bool`

HasStr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


