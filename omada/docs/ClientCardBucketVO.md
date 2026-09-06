# ClientCardBucketVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients2g** | Pointer to **int32** | Number of 2.4GHz clients in this bucket | [optional] 
**Clients5g** | Pointer to **int32** | Number of 5GHz clients in this bucket | [optional] 
**Clients6g** | Pointer to **int32** | Number of 6GHz clients in this bucket | [optional] 
**Count** | Pointer to **int32** | Total number of clients in this bucket | [optional] 
**Label** | Pointer to **string** | Bucket label, e.g. &#39;&gt;&#x3D; -50 dBm&#39; or &#39;[10, 20) dB&#39; | [optional] 
**LowerBound** | Pointer to **int32** | Lower bound of the bucket, null means negative infinity | [optional] 
**LowerInclusive** | Pointer to **bool** | Whether the lower bound is inclusive | [optional] 
**Percent** | Pointer to **int32** | Percentage of clients in this bucket (0-100) | [optional] 
**UpperBound** | Pointer to **int32** | Upper bound of the bucket, null means positive infinity | [optional] 
**UpperInclusive** | Pointer to **bool** | Whether the upper bound is inclusive | [optional] 

## Methods

### NewClientCardBucketVO

`func NewClientCardBucketVO() *ClientCardBucketVO`

NewClientCardBucketVO instantiates a new ClientCardBucketVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCardBucketVOWithDefaults

`func NewClientCardBucketVOWithDefaults() *ClientCardBucketVO`

NewClientCardBucketVOWithDefaults instantiates a new ClientCardBucketVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients2g

`func (o *ClientCardBucketVO) GetClients2g() int32`

GetClients2g returns the Clients2g field if non-nil, zero value otherwise.

### GetClients2gOk

`func (o *ClientCardBucketVO) GetClients2gOk() (*int32, bool)`

GetClients2gOk returns a tuple with the Clients2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients2g

`func (o *ClientCardBucketVO) SetClients2g(v int32)`

SetClients2g sets Clients2g field to given value.

### HasClients2g

`func (o *ClientCardBucketVO) HasClients2g() bool`

HasClients2g returns a boolean if a field has been set.

### GetClients5g

`func (o *ClientCardBucketVO) GetClients5g() int32`

GetClients5g returns the Clients5g field if non-nil, zero value otherwise.

### GetClients5gOk

`func (o *ClientCardBucketVO) GetClients5gOk() (*int32, bool)`

GetClients5gOk returns a tuple with the Clients5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients5g

`func (o *ClientCardBucketVO) SetClients5g(v int32)`

SetClients5g sets Clients5g field to given value.

### HasClients5g

`func (o *ClientCardBucketVO) HasClients5g() bool`

HasClients5g returns a boolean if a field has been set.

### GetClients6g

`func (o *ClientCardBucketVO) GetClients6g() int32`

GetClients6g returns the Clients6g field if non-nil, zero value otherwise.

### GetClients6gOk

`func (o *ClientCardBucketVO) GetClients6gOk() (*int32, bool)`

GetClients6gOk returns a tuple with the Clients6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients6g

`func (o *ClientCardBucketVO) SetClients6g(v int32)`

SetClients6g sets Clients6g field to given value.

### HasClients6g

`func (o *ClientCardBucketVO) HasClients6g() bool`

HasClients6g returns a boolean if a field has been set.

### GetCount

`func (o *ClientCardBucketVO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ClientCardBucketVO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ClientCardBucketVO) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ClientCardBucketVO) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLabel

`func (o *ClientCardBucketVO) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ClientCardBucketVO) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ClientCardBucketVO) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ClientCardBucketVO) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLowerBound

`func (o *ClientCardBucketVO) GetLowerBound() int32`

GetLowerBound returns the LowerBound field if non-nil, zero value otherwise.

### GetLowerBoundOk

`func (o *ClientCardBucketVO) GetLowerBoundOk() (*int32, bool)`

GetLowerBoundOk returns a tuple with the LowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBound

`func (o *ClientCardBucketVO) SetLowerBound(v int32)`

SetLowerBound sets LowerBound field to given value.

### HasLowerBound

`func (o *ClientCardBucketVO) HasLowerBound() bool`

HasLowerBound returns a boolean if a field has been set.

### GetLowerInclusive

`func (o *ClientCardBucketVO) GetLowerInclusive() bool`

GetLowerInclusive returns the LowerInclusive field if non-nil, zero value otherwise.

### GetLowerInclusiveOk

`func (o *ClientCardBucketVO) GetLowerInclusiveOk() (*bool, bool)`

GetLowerInclusiveOk returns a tuple with the LowerInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerInclusive

`func (o *ClientCardBucketVO) SetLowerInclusive(v bool)`

SetLowerInclusive sets LowerInclusive field to given value.

### HasLowerInclusive

`func (o *ClientCardBucketVO) HasLowerInclusive() bool`

HasLowerInclusive returns a boolean if a field has been set.

### GetPercent

`func (o *ClientCardBucketVO) GetPercent() int32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *ClientCardBucketVO) GetPercentOk() (*int32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *ClientCardBucketVO) SetPercent(v int32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *ClientCardBucketVO) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetUpperBound

`func (o *ClientCardBucketVO) GetUpperBound() int32`

GetUpperBound returns the UpperBound field if non-nil, zero value otherwise.

### GetUpperBoundOk

`func (o *ClientCardBucketVO) GetUpperBoundOk() (*int32, bool)`

GetUpperBoundOk returns a tuple with the UpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBound

`func (o *ClientCardBucketVO) SetUpperBound(v int32)`

SetUpperBound sets UpperBound field to given value.

### HasUpperBound

`func (o *ClientCardBucketVO) HasUpperBound() bool`

HasUpperBound returns a boolean if a field has been set.

### GetUpperInclusive

`func (o *ClientCardBucketVO) GetUpperInclusive() bool`

GetUpperInclusive returns the UpperInclusive field if non-nil, zero value otherwise.

### GetUpperInclusiveOk

`func (o *ClientCardBucketVO) GetUpperInclusiveOk() (*bool, bool)`

GetUpperInclusiveOk returns a tuple with the UpperInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperInclusive

`func (o *ClientCardBucketVO) SetUpperInclusive(v bool)`

SetUpperInclusive sets UpperInclusive field to given value.

### HasUpperInclusive

`func (o *ClientCardBucketVO) HasUpperInclusive() bool`

HasUpperInclusive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


