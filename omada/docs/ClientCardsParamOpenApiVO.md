# ClientCardsParamOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]BucketBoundaryVO**](BucketBoundaryVO.md) | Custom bucket boundaries (ascending by value). If not provided, default boundaries for the card type will be used. | [optional] 
**TopK** | Pointer to **string** | TopK parameter of the card. | [optional] 
**Type** | Pointer to **string** | Type of the card, e.g. &#39;rssi&#39;, &#39;snr&#39; | [optional] 

## Methods

### NewClientCardsParamOpenApiVO

`func NewClientCardsParamOpenApiVO() *ClientCardsParamOpenApiVO`

NewClientCardsParamOpenApiVO instantiates a new ClientCardsParamOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCardsParamOpenApiVOWithDefaults

`func NewClientCardsParamOpenApiVOWithDefaults() *ClientCardsParamOpenApiVO`

NewClientCardsParamOpenApiVOWithDefaults instantiates a new ClientCardsParamOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *ClientCardsParamOpenApiVO) GetBuckets() []BucketBoundaryVO`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *ClientCardsParamOpenApiVO) GetBucketsOk() (*[]BucketBoundaryVO, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *ClientCardsParamOpenApiVO) SetBuckets(v []BucketBoundaryVO)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *ClientCardsParamOpenApiVO) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetTopK

`func (o *ClientCardsParamOpenApiVO) GetTopK() string`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *ClientCardsParamOpenApiVO) GetTopKOk() (*string, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *ClientCardsParamOpenApiVO) SetTopK(v string)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *ClientCardsParamOpenApiVO) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetType

`func (o *ClientCardsParamOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientCardsParamOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientCardsParamOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientCardsParamOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


