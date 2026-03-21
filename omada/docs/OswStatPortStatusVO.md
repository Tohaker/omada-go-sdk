# OswStatPortStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkSpeed** | Pointer to **int32** | LinkSpeed should be a value as follows: 1:10Mbps; 2:100Mbps; 3:1000Mbps; 4:10Gbps | [optional] 
**LinkStatus** | Pointer to **int32** | LinkStatus should be a value as follows: 0:link down; 1:link up | [optional] 
**Poe** | Pointer to **bool** | Indicates whether PoE power supply is in use | [optional] 
**Rx** | Pointer to **int64** | Receive traffic of the port, in bytes | [optional] 
**StkStatus** | Pointer to **int32** | Stack Status should be a value as follows: 0:DOWN; 1:AUTHFAIL; 2:IEEE; 3:OK | [optional] 
**StpDiscarding** | Pointer to **bool** | STP Discarding | [optional] 
**Total** | Pointer to **int64** | Transmit traffic + Receive traffic of the port, in bytes | [optional] 
**Tx** | Pointer to **int64** | Transmit traffic of the port, in bytes | [optional] 

## Methods

### NewOswStatPortStatusVO

`func NewOswStatPortStatusVO() *OswStatPortStatusVO`

NewOswStatPortStatusVO instantiates a new OswStatPortStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStatPortStatusVOWithDefaults

`func NewOswStatPortStatusVOWithDefaults() *OswStatPortStatusVO`

NewOswStatPortStatusVOWithDefaults instantiates a new OswStatPortStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkSpeed

`func (o *OswStatPortStatusVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswStatPortStatusVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswStatPortStatusVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswStatPortStatusVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkStatus

`func (o *OswStatPortStatusVO) GetLinkStatus() int32`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *OswStatPortStatusVO) GetLinkStatusOk() (*int32, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *OswStatPortStatusVO) SetLinkStatus(v int32)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *OswStatPortStatusVO) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetPoe

`func (o *OswStatPortStatusVO) GetPoe() bool`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *OswStatPortStatusVO) GetPoeOk() (*bool, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *OswStatPortStatusVO) SetPoe(v bool)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *OswStatPortStatusVO) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetRx

`func (o *OswStatPortStatusVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *OswStatPortStatusVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *OswStatPortStatusVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *OswStatPortStatusVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetStkStatus

`func (o *OswStatPortStatusVO) GetStkStatus() int32`

GetStkStatus returns the StkStatus field if non-nil, zero value otherwise.

### GetStkStatusOk

`func (o *OswStatPortStatusVO) GetStkStatusOk() (*int32, bool)`

GetStkStatusOk returns a tuple with the StkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkStatus

`func (o *OswStatPortStatusVO) SetStkStatus(v int32)`

SetStkStatus sets StkStatus field to given value.

### HasStkStatus

`func (o *OswStatPortStatusVO) HasStkStatus() bool`

HasStkStatus returns a boolean if a field has been set.

### GetStpDiscarding

`func (o *OswStatPortStatusVO) GetStpDiscarding() bool`

GetStpDiscarding returns the StpDiscarding field if non-nil, zero value otherwise.

### GetStpDiscardingOk

`func (o *OswStatPortStatusVO) GetStpDiscardingOk() (*bool, bool)`

GetStpDiscardingOk returns a tuple with the StpDiscarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpDiscarding

`func (o *OswStatPortStatusVO) SetStpDiscarding(v bool)`

SetStpDiscarding sets StpDiscarding field to given value.

### HasStpDiscarding

`func (o *OswStatPortStatusVO) HasStpDiscarding() bool`

HasStpDiscarding returns a boolean if a field has been set.

### GetTotal

`func (o *OswStatPortStatusVO) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OswStatPortStatusVO) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OswStatPortStatusVO) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OswStatPortStatusVO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTx

`func (o *OswStatPortStatusVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *OswStatPortStatusVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *OswStatPortStatusVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *OswStatPortStatusVO) HasTx() bool`

HasTx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


