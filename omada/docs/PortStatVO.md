# PortStatVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **int64** | Total traffic of the port, in bytes | [optional] 
**Rx** | Pointer to **int64** | Total receive traffic of the port, in bytes | [optional] 
**Tx** | Pointer to **int64** | Total transmit traffic of the port, in bytes | [optional] 

## Methods

### NewPortStatVO

`func NewPortStatVO() *PortStatVO`

NewPortStatVO instantiates a new PortStatVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortStatVOWithDefaults

`func NewPortStatVOWithDefaults() *PortStatVO`

NewPortStatVOWithDefaults instantiates a new PortStatVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *PortStatVO) GetAll() int64`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *PortStatVO) GetAllOk() (*int64, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *PortStatVO) SetAll(v int64)`

SetAll sets All field to given value.

### HasAll

`func (o *PortStatVO) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetRx

`func (o *PortStatVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *PortStatVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *PortStatVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *PortStatVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetTx

`func (o *PortStatVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *PortStatVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *PortStatVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *PortStatVO) HasTx() bool`

HasTx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


