# ClientCardDataVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]ClientCardBucketVO**](ClientCardBucketVO.md) | Distribution buckets for this card | [optional] 
**TotalCount** | Pointer to **int32** | Total number of clients counted for this card | [optional] 

## Methods

### NewClientCardDataVO

`func NewClientCardDataVO() *ClientCardDataVO`

NewClientCardDataVO instantiates a new ClientCardDataVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCardDataVOWithDefaults

`func NewClientCardDataVOWithDefaults() *ClientCardDataVO`

NewClientCardDataVOWithDefaults instantiates a new ClientCardDataVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *ClientCardDataVO) GetBuckets() []ClientCardBucketVO`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *ClientCardDataVO) GetBucketsOk() (*[]ClientCardBucketVO, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *ClientCardDataVO) SetBuckets(v []ClientCardBucketVO)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *ClientCardDataVO) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetTotalCount

`func (o *ClientCardDataVO) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ClientCardDataVO) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ClientCardDataVO) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ClientCardDataVO) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


