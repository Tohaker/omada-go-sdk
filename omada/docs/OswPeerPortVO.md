# OswPeerPortVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedType** | Pointer to **int32** | Blocked Type | [optional] 
**BlockedVlans** | Pointer to **string** | Blocked Vlans | [optional] 
**Duplex** | Pointer to **int32** | Duplex | [optional] 
**LagId** | Pointer to **int32** | LagId. If it is not null, it indicates that the port is a LAG port | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed | [optional] 
**Port** | Pointer to **int32** | Port | [optional] 
**Rx** | Pointer to **int64** | Port total rx bytes | [optional] 
**RxRate** | Pointer to **int64** | Port rx Rate | [optional] 
**StPort** | Pointer to **string** | Standard Port, unit/slot/port | [optional] 
**StpDiscarding** | Pointer to **bool** | STP Discarding | [optional] 
**Tx** | Pointer to **int64** | Port total tx bytes | [optional] 
**TxRate** | Pointer to **int64** | Port tx Rate | [optional] 

## Methods

### NewOswPeerPortVO

`func NewOswPeerPortVO() *OswPeerPortVO`

NewOswPeerPortVO instantiates a new OswPeerPortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPeerPortVOWithDefaults

`func NewOswPeerPortVOWithDefaults() *OswPeerPortVO`

NewOswPeerPortVOWithDefaults instantiates a new OswPeerPortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedType

`func (o *OswPeerPortVO) GetBlockedType() int32`

GetBlockedType returns the BlockedType field if non-nil, zero value otherwise.

### GetBlockedTypeOk

`func (o *OswPeerPortVO) GetBlockedTypeOk() (*int32, bool)`

GetBlockedTypeOk returns a tuple with the BlockedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedType

`func (o *OswPeerPortVO) SetBlockedType(v int32)`

SetBlockedType sets BlockedType field to given value.

### HasBlockedType

`func (o *OswPeerPortVO) HasBlockedType() bool`

HasBlockedType returns a boolean if a field has been set.

### GetBlockedVlans

`func (o *OswPeerPortVO) GetBlockedVlans() string`

GetBlockedVlans returns the BlockedVlans field if non-nil, zero value otherwise.

### GetBlockedVlansOk

`func (o *OswPeerPortVO) GetBlockedVlansOk() (*string, bool)`

GetBlockedVlansOk returns a tuple with the BlockedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedVlans

`func (o *OswPeerPortVO) SetBlockedVlans(v string)`

SetBlockedVlans sets BlockedVlans field to given value.

### HasBlockedVlans

`func (o *OswPeerPortVO) HasBlockedVlans() bool`

HasBlockedVlans returns a boolean if a field has been set.

### GetDuplex

`func (o *OswPeerPortVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswPeerPortVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswPeerPortVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswPeerPortVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLagId

`func (o *OswPeerPortVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *OswPeerPortVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *OswPeerPortVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *OswPeerPortVO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswPeerPortVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswPeerPortVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswPeerPortVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswPeerPortVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetPort

`func (o *OswPeerPortVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswPeerPortVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswPeerPortVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswPeerPortVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRx

`func (o *OswPeerPortVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *OswPeerPortVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *OswPeerPortVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *OswPeerPortVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxRate

`func (o *OswPeerPortVO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OswPeerPortVO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OswPeerPortVO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OswPeerPortVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetStPort

`func (o *OswPeerPortVO) GetStPort() string`

GetStPort returns the StPort field if non-nil, zero value otherwise.

### GetStPortOk

`func (o *OswPeerPortVO) GetStPortOk() (*string, bool)`

GetStPortOk returns a tuple with the StPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStPort

`func (o *OswPeerPortVO) SetStPort(v string)`

SetStPort sets StPort field to given value.

### HasStPort

`func (o *OswPeerPortVO) HasStPort() bool`

HasStPort returns a boolean if a field has been set.

### GetStpDiscarding

`func (o *OswPeerPortVO) GetStpDiscarding() bool`

GetStpDiscarding returns the StpDiscarding field if non-nil, zero value otherwise.

### GetStpDiscardingOk

`func (o *OswPeerPortVO) GetStpDiscardingOk() (*bool, bool)`

GetStpDiscardingOk returns a tuple with the StpDiscarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpDiscarding

`func (o *OswPeerPortVO) SetStpDiscarding(v bool)`

SetStpDiscarding sets StpDiscarding field to given value.

### HasStpDiscarding

`func (o *OswPeerPortVO) HasStpDiscarding() bool`

HasStpDiscarding returns a boolean if a field has been set.

### GetTx

`func (o *OswPeerPortVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *OswPeerPortVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *OswPeerPortVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *OswPeerPortVO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxRate

`func (o *OswPeerPortVO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OswPeerPortVO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OswPeerPortVO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OswPeerPortVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


