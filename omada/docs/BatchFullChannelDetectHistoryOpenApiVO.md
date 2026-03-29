# BatchFullChannelDetectHistoryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoryId** | Pointer to **string** | Id of batch full channel detect history. | [optional] 
**Status** | Pointer to **int32** | Status of full channel detect, status should be a value as follows: 0: finish; 1: scanning. | [optional] 

## Methods

### NewBatchFullChannelDetectHistoryOpenApiVO

`func NewBatchFullChannelDetectHistoryOpenApiVO() *BatchFullChannelDetectHistoryOpenApiVO`

NewBatchFullChannelDetectHistoryOpenApiVO instantiates a new BatchFullChannelDetectHistoryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchFullChannelDetectHistoryOpenApiVOWithDefaults

`func NewBatchFullChannelDetectHistoryOpenApiVOWithDefaults() *BatchFullChannelDetectHistoryOpenApiVO`

NewBatchFullChannelDetectHistoryOpenApiVOWithDefaults instantiates a new BatchFullChannelDetectHistoryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoryId

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetHistoryId() string`

GetHistoryId returns the HistoryId field if non-nil, zero value otherwise.

### GetHistoryIdOk

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetHistoryIdOk() (*string, bool)`

GetHistoryIdOk returns a tuple with the HistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryId

`func (o *BatchFullChannelDetectHistoryOpenApiVO) SetHistoryId(v string)`

SetHistoryId sets HistoryId field to given value.

### HasHistoryId

`func (o *BatchFullChannelDetectHistoryOpenApiVO) HasHistoryId() bool`

HasHistoryId returns a boolean if a field has been set.

### GetStatus

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchFullChannelDetectHistoryOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchFullChannelDetectHistoryOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchFullChannelDetectHistoryOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


