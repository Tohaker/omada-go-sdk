# BatchOspfProcessOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessIdList** | **[]string** | List of Process ID | 
**SearchKey** | Pointer to **string** | SearchKey | [optional] 
**SelectType** | **int32** | Select Type, it should be a value as follows: 1: SELECT ALL, 2: SELECT INCLUDE, 3: SELECT EXCLUDE. | 

## Methods

### NewBatchOspfProcessOpenApiVO

`func NewBatchOspfProcessOpenApiVO(processIdList []string, selectType int32, ) *BatchOspfProcessOpenApiVO`

NewBatchOspfProcessOpenApiVO instantiates a new BatchOspfProcessOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchOspfProcessOpenApiVOWithDefaults

`func NewBatchOspfProcessOpenApiVOWithDefaults() *BatchOspfProcessOpenApiVO`

NewBatchOspfProcessOpenApiVOWithDefaults instantiates a new BatchOspfProcessOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessIdList

`func (o *BatchOspfProcessOpenApiVO) GetProcessIdList() []string`

GetProcessIdList returns the ProcessIdList field if non-nil, zero value otherwise.

### GetProcessIdListOk

`func (o *BatchOspfProcessOpenApiVO) GetProcessIdListOk() (*[]string, bool)`

GetProcessIdListOk returns a tuple with the ProcessIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessIdList

`func (o *BatchOspfProcessOpenApiVO) SetProcessIdList(v []string)`

SetProcessIdList sets ProcessIdList field to given value.


### GetSearchKey

`func (o *BatchOspfProcessOpenApiVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *BatchOspfProcessOpenApiVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *BatchOspfProcessOpenApiVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *BatchOspfProcessOpenApiVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectType

`func (o *BatchOspfProcessOpenApiVO) GetSelectType() int32`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *BatchOspfProcessOpenApiVO) GetSelectTypeOk() (*int32, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *BatchOspfProcessOpenApiVO) SetSelectType(v int32)`

SetSelectType sets SelectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


