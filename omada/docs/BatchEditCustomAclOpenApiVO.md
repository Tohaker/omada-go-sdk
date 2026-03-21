# BatchEditCustomAclOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | When selectType is set to all, the ids do not need to be passed and all entries are processed, when selectType is set to include, the id entries contained in the ids are processed, when selectType is set to exclude, the id entries that are not contained in the ids are processed | 
**Log** | **int32** | log status of acl. | 
**SearchKey** | Pointer to **string** | Fuzzy query parameters, support field: voucher code | [optional] 
**SelectType** | Pointer to **string** | SelectType all, include or exclude | [optional] 
**Status** | **int32** | enable status of acl. | 

## Methods

### NewBatchEditCustomAclOpenApiVO

`func NewBatchEditCustomAclOpenApiVO(ids []string, log int32, status int32, ) *BatchEditCustomAclOpenApiVO`

NewBatchEditCustomAclOpenApiVO instantiates a new BatchEditCustomAclOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchEditCustomAclOpenApiVOWithDefaults

`func NewBatchEditCustomAclOpenApiVOWithDefaults() *BatchEditCustomAclOpenApiVO`

NewBatchEditCustomAclOpenApiVOWithDefaults instantiates a new BatchEditCustomAclOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *BatchEditCustomAclOpenApiVO) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BatchEditCustomAclOpenApiVO) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BatchEditCustomAclOpenApiVO) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetLog

`func (o *BatchEditCustomAclOpenApiVO) GetLog() int32`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *BatchEditCustomAclOpenApiVO) GetLogOk() (*int32, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *BatchEditCustomAclOpenApiVO) SetLog(v int32)`

SetLog sets Log field to given value.


### GetSearchKey

`func (o *BatchEditCustomAclOpenApiVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *BatchEditCustomAclOpenApiVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *BatchEditCustomAclOpenApiVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *BatchEditCustomAclOpenApiVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectType

`func (o *BatchEditCustomAclOpenApiVO) GetSelectType() string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *BatchEditCustomAclOpenApiVO) GetSelectTypeOk() (*string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *BatchEditCustomAclOpenApiVO) SetSelectType(v string)`

SetSelectType sets SelectType field to given value.

### HasSelectType

`func (o *BatchEditCustomAclOpenApiVO) HasSelectType() bool`

HasSelectType returns a boolean if a field has been set.

### GetStatus

`func (o *BatchEditCustomAclOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchEditCustomAclOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchEditCustomAclOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


