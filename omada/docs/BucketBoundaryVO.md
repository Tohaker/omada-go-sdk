# BucketBoundaryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LowerInclusive** | Pointer to **bool** | Whether the boundary value is included in the upper bucket (&gt;&#x3D; value). Default: true | [optional] 
**Value** | Pointer to **int32** | Boundary value | [optional] 

## Methods

### NewBucketBoundaryVO

`func NewBucketBoundaryVO() *BucketBoundaryVO`

NewBucketBoundaryVO instantiates a new BucketBoundaryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketBoundaryVOWithDefaults

`func NewBucketBoundaryVOWithDefaults() *BucketBoundaryVO`

NewBucketBoundaryVOWithDefaults instantiates a new BucketBoundaryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLowerInclusive

`func (o *BucketBoundaryVO) GetLowerInclusive() bool`

GetLowerInclusive returns the LowerInclusive field if non-nil, zero value otherwise.

### GetLowerInclusiveOk

`func (o *BucketBoundaryVO) GetLowerInclusiveOk() (*bool, bool)`

GetLowerInclusiveOk returns a tuple with the LowerInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerInclusive

`func (o *BucketBoundaryVO) SetLowerInclusive(v bool)`

SetLowerInclusive sets LowerInclusive field to given value.

### HasLowerInclusive

`func (o *BucketBoundaryVO) HasLowerInclusive() bool`

HasLowerInclusive returns a boolean if a field has been set.

### GetValue

`func (o *BucketBoundaryVO) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BucketBoundaryVO) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BucketBoundaryVO) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *BucketBoundaryVO) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


