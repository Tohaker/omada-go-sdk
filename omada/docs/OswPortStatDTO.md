# OswPortStatDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DropPkts** | Pointer to **int64** |  | [optional] 
**LinkDownCnt** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Rx** | Pointer to **int64** |  | [optional] 
**RxBroadPkts** | Pointer to **int64** |  | [optional] 
**RxErrPkts** | Pointer to **int64** |  | [optional] 
**RxMultiPkts** | Pointer to **int64** |  | [optional] 
**RxPkts** | Pointer to **int64** |  | [optional] 
**RxRate** | Pointer to **int64** |  | [optional] 
**StandardPort** | Pointer to **string** |  | [optional] 
**Tx** | Pointer to **int64** |  | [optional] 
**TxBroadPkts** | Pointer to **int64** |  | [optional] 
**TxErrPkts** | Pointer to **int64** |  | [optional] 
**TxMultiPkts** | Pointer to **int64** |  | [optional] 
**TxPkts** | Pointer to **int64** |  | [optional] 
**TxRate** | Pointer to **int64** |  | [optional] 

## Methods

### NewOswPortStatDTO

`func NewOswPortStatDTO() *OswPortStatDTO`

NewOswPortStatDTO instantiates a new OswPortStatDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPortStatDTOWithDefaults

`func NewOswPortStatDTOWithDefaults() *OswPortStatDTO`

NewOswPortStatDTOWithDefaults instantiates a new OswPortStatDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropPkts

`func (o *OswPortStatDTO) GetDropPkts() int64`

GetDropPkts returns the DropPkts field if non-nil, zero value otherwise.

### GetDropPktsOk

`func (o *OswPortStatDTO) GetDropPktsOk() (*int64, bool)`

GetDropPktsOk returns a tuple with the DropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropPkts

`func (o *OswPortStatDTO) SetDropPkts(v int64)`

SetDropPkts sets DropPkts field to given value.

### HasDropPkts

`func (o *OswPortStatDTO) HasDropPkts() bool`

HasDropPkts returns a boolean if a field has been set.

### GetLinkDownCnt

`func (o *OswPortStatDTO) GetLinkDownCnt() int32`

GetLinkDownCnt returns the LinkDownCnt field if non-nil, zero value otherwise.

### GetLinkDownCntOk

`func (o *OswPortStatDTO) GetLinkDownCntOk() (*int32, bool)`

GetLinkDownCntOk returns a tuple with the LinkDownCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDownCnt

`func (o *OswPortStatDTO) SetLinkDownCnt(v int32)`

SetLinkDownCnt sets LinkDownCnt field to given value.

### HasLinkDownCnt

`func (o *OswPortStatDTO) HasLinkDownCnt() bool`

HasLinkDownCnt returns a boolean if a field has been set.

### GetPort

`func (o *OswPortStatDTO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswPortStatDTO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswPortStatDTO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswPortStatDTO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRx

`func (o *OswPortStatDTO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *OswPortStatDTO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *OswPortStatDTO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *OswPortStatDTO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxBroadPkts

`func (o *OswPortStatDTO) GetRxBroadPkts() int64`

GetRxBroadPkts returns the RxBroadPkts field if non-nil, zero value otherwise.

### GetRxBroadPktsOk

`func (o *OswPortStatDTO) GetRxBroadPktsOk() (*int64, bool)`

GetRxBroadPktsOk returns a tuple with the RxBroadPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBroadPkts

`func (o *OswPortStatDTO) SetRxBroadPkts(v int64)`

SetRxBroadPkts sets RxBroadPkts field to given value.

### HasRxBroadPkts

`func (o *OswPortStatDTO) HasRxBroadPkts() bool`

HasRxBroadPkts returns a boolean if a field has been set.

### GetRxErrPkts

`func (o *OswPortStatDTO) GetRxErrPkts() int64`

GetRxErrPkts returns the RxErrPkts field if non-nil, zero value otherwise.

### GetRxErrPktsOk

`func (o *OswPortStatDTO) GetRxErrPktsOk() (*int64, bool)`

GetRxErrPktsOk returns a tuple with the RxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrPkts

`func (o *OswPortStatDTO) SetRxErrPkts(v int64)`

SetRxErrPkts sets RxErrPkts field to given value.

### HasRxErrPkts

`func (o *OswPortStatDTO) HasRxErrPkts() bool`

HasRxErrPkts returns a boolean if a field has been set.

### GetRxMultiPkts

`func (o *OswPortStatDTO) GetRxMultiPkts() int64`

GetRxMultiPkts returns the RxMultiPkts field if non-nil, zero value otherwise.

### GetRxMultiPktsOk

`func (o *OswPortStatDTO) GetRxMultiPktsOk() (*int64, bool)`

GetRxMultiPktsOk returns a tuple with the RxMultiPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMultiPkts

`func (o *OswPortStatDTO) SetRxMultiPkts(v int64)`

SetRxMultiPkts sets RxMultiPkts field to given value.

### HasRxMultiPkts

`func (o *OswPortStatDTO) HasRxMultiPkts() bool`

HasRxMultiPkts returns a boolean if a field has been set.

### GetRxPkts

`func (o *OswPortStatDTO) GetRxPkts() int64`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *OswPortStatDTO) GetRxPktsOk() (*int64, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *OswPortStatDTO) SetRxPkts(v int64)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *OswPortStatDTO) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetRxRate

`func (o *OswPortStatDTO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OswPortStatDTO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OswPortStatDTO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OswPortStatDTO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswPortStatDTO) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswPortStatDTO) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswPortStatDTO) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswPortStatDTO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetTx

`func (o *OswPortStatDTO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *OswPortStatDTO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *OswPortStatDTO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *OswPortStatDTO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxBroadPkts

`func (o *OswPortStatDTO) GetTxBroadPkts() int64`

GetTxBroadPkts returns the TxBroadPkts field if non-nil, zero value otherwise.

### GetTxBroadPktsOk

`func (o *OswPortStatDTO) GetTxBroadPktsOk() (*int64, bool)`

GetTxBroadPktsOk returns a tuple with the TxBroadPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBroadPkts

`func (o *OswPortStatDTO) SetTxBroadPkts(v int64)`

SetTxBroadPkts sets TxBroadPkts field to given value.

### HasTxBroadPkts

`func (o *OswPortStatDTO) HasTxBroadPkts() bool`

HasTxBroadPkts returns a boolean if a field has been set.

### GetTxErrPkts

`func (o *OswPortStatDTO) GetTxErrPkts() int64`

GetTxErrPkts returns the TxErrPkts field if non-nil, zero value otherwise.

### GetTxErrPktsOk

`func (o *OswPortStatDTO) GetTxErrPktsOk() (*int64, bool)`

GetTxErrPktsOk returns a tuple with the TxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrPkts

`func (o *OswPortStatDTO) SetTxErrPkts(v int64)`

SetTxErrPkts sets TxErrPkts field to given value.

### HasTxErrPkts

`func (o *OswPortStatDTO) HasTxErrPkts() bool`

HasTxErrPkts returns a boolean if a field has been set.

### GetTxMultiPkts

`func (o *OswPortStatDTO) GetTxMultiPkts() int64`

GetTxMultiPkts returns the TxMultiPkts field if non-nil, zero value otherwise.

### GetTxMultiPktsOk

`func (o *OswPortStatDTO) GetTxMultiPktsOk() (*int64, bool)`

GetTxMultiPktsOk returns a tuple with the TxMultiPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMultiPkts

`func (o *OswPortStatDTO) SetTxMultiPkts(v int64)`

SetTxMultiPkts sets TxMultiPkts field to given value.

### HasTxMultiPkts

`func (o *OswPortStatDTO) HasTxMultiPkts() bool`

HasTxMultiPkts returns a boolean if a field has been set.

### GetTxPkts

`func (o *OswPortStatDTO) GetTxPkts() int64`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *OswPortStatDTO) GetTxPktsOk() (*int64, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *OswPortStatDTO) SetTxPkts(v int64)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *OswPortStatDTO) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetTxRate

`func (o *OswPortStatDTO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OswPortStatDTO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OswPortStatDTO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OswPortStatDTO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


