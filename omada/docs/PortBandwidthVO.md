# PortBandwidthVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortUuid** | Pointer to **string** | Port uuid. | [optional] 
**RxBandwidth** | Pointer to **int64** | Set down link bandwidth. | [optional] 
**TxBandwidth** | Pointer to **int64** | Set up link bandwidth. | [optional] 

## Methods

### NewPortBandwidthVO

`func NewPortBandwidthVO() *PortBandwidthVO`

NewPortBandwidthVO instantiates a new PortBandwidthVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortBandwidthVOWithDefaults

`func NewPortBandwidthVOWithDefaults() *PortBandwidthVO`

NewPortBandwidthVOWithDefaults instantiates a new PortBandwidthVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortUuid

`func (o *PortBandwidthVO) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *PortBandwidthVO) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *PortBandwidthVO) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.

### HasPortUuid

`func (o *PortBandwidthVO) HasPortUuid() bool`

HasPortUuid returns a boolean if a field has been set.

### GetRxBandwidth

`func (o *PortBandwidthVO) GetRxBandwidth() int64`

GetRxBandwidth returns the RxBandwidth field if non-nil, zero value otherwise.

### GetRxBandwidthOk

`func (o *PortBandwidthVO) GetRxBandwidthOk() (*int64, bool)`

GetRxBandwidthOk returns a tuple with the RxBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBandwidth

`func (o *PortBandwidthVO) SetRxBandwidth(v int64)`

SetRxBandwidth sets RxBandwidth field to given value.

### HasRxBandwidth

`func (o *PortBandwidthVO) HasRxBandwidth() bool`

HasRxBandwidth returns a boolean if a field has been set.

### GetTxBandwidth

`func (o *PortBandwidthVO) GetTxBandwidth() int64`

GetTxBandwidth returns the TxBandwidth field if non-nil, zero value otherwise.

### GetTxBandwidthOk

`func (o *PortBandwidthVO) GetTxBandwidthOk() (*int64, bool)`

GetTxBandwidthOk returns a tuple with the TxBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBandwidth

`func (o *PortBandwidthVO) SetTxBandwidth(v int64)`

SetTxBandwidth sets TxBandwidth field to given value.

### HasTxBandwidth

`func (o *PortBandwidthVO) HasTxBandwidth() bool`

HasTxBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


