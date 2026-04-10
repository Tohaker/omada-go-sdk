# ApPortStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rx** | Pointer to **int64** | Rx Bytes | [optional] 
**Tx** | Pointer to **int64** | Tx Bytes | [optional] 

## Methods

### NewApPortStatusVO

`func NewApPortStatusVO() *ApPortStatusVO`

NewApPortStatusVO instantiates a new ApPortStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApPortStatusVOWithDefaults

`func NewApPortStatusVOWithDefaults() *ApPortStatusVO`

NewApPortStatusVOWithDefaults instantiates a new ApPortStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRx

`func (o *ApPortStatusVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *ApPortStatusVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *ApPortStatusVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *ApPortStatusVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetTx

`func (o *ApPortStatusVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *ApPortStatusVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *ApPortStatusVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *ApPortStatusVO) HasTx() bool`

HasTx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


