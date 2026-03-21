# OswLagStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedType** | Pointer to **int32** | Blocked Type | [optional] 
**BlockedVlans** | Pointer to **string** | Blocked Vlans | [optional] 
**Duplex** | Pointer to **int32** | Duplex should be a value as follows: 0: Auto; 1: Half; 2: Full | [optional] 
**LagId** | Pointer to **int32** | Lag ID | [optional] 
**LagType** | Pointer to **int32** | LagType should be a value as follows: 1: static; 2: LACP | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed should be a value as follows: 0: auto; 1: 10M; 2: 100M; 3: 1000M; 4: 2.5G; 5: 10G | [optional] 
**LinkStatus** | Pointer to **int32** | Link Status should be a value as follows: 1: link up; 0: link down | [optional] 
**Name** | Pointer to **string** | Lag name | [optional] 
**Ports** | Pointer to **[]int32** | Ports member | [optional] 
**Rx** | Pointer to **int64** | Rx | [optional] 
**RxRate** | Pointer to **int64** | Rx Rate | [optional] 
**StandardPorts** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | Standard Ports | [optional] 
**StpDiscarding** | Pointer to **bool** | STP Discarding | [optional] 
**Tx** | Pointer to **int64** | Tx | [optional] 
**TxRate** | Pointer to **int64** | Tx Rate | [optional] 

## Methods

### NewOswLagStatusVO

`func NewOswLagStatusVO() *OswLagStatusVO`

NewOswLagStatusVO instantiates a new OswLagStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswLagStatusVOWithDefaults

`func NewOswLagStatusVOWithDefaults() *OswLagStatusVO`

NewOswLagStatusVOWithDefaults instantiates a new OswLagStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedType

`func (o *OswLagStatusVO) GetBlockedType() int32`

GetBlockedType returns the BlockedType field if non-nil, zero value otherwise.

### GetBlockedTypeOk

`func (o *OswLagStatusVO) GetBlockedTypeOk() (*int32, bool)`

GetBlockedTypeOk returns a tuple with the BlockedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedType

`func (o *OswLagStatusVO) SetBlockedType(v int32)`

SetBlockedType sets BlockedType field to given value.

### HasBlockedType

`func (o *OswLagStatusVO) HasBlockedType() bool`

HasBlockedType returns a boolean if a field has been set.

### GetBlockedVlans

`func (o *OswLagStatusVO) GetBlockedVlans() string`

GetBlockedVlans returns the BlockedVlans field if non-nil, zero value otherwise.

### GetBlockedVlansOk

`func (o *OswLagStatusVO) GetBlockedVlansOk() (*string, bool)`

GetBlockedVlansOk returns a tuple with the BlockedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedVlans

`func (o *OswLagStatusVO) SetBlockedVlans(v string)`

SetBlockedVlans sets BlockedVlans field to given value.

### HasBlockedVlans

`func (o *OswLagStatusVO) HasBlockedVlans() bool`

HasBlockedVlans returns a boolean if a field has been set.

### GetDuplex

`func (o *OswLagStatusVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswLagStatusVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswLagStatusVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswLagStatusVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLagId

`func (o *OswLagStatusVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *OswLagStatusVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *OswLagStatusVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *OswLagStatusVO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLagType

`func (o *OswLagStatusVO) GetLagType() int32`

GetLagType returns the LagType field if non-nil, zero value otherwise.

### GetLagTypeOk

`func (o *OswLagStatusVO) GetLagTypeOk() (*int32, bool)`

GetLagTypeOk returns a tuple with the LagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagType

`func (o *OswLagStatusVO) SetLagType(v int32)`

SetLagType sets LagType field to given value.

### HasLagType

`func (o *OswLagStatusVO) HasLagType() bool`

HasLagType returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswLagStatusVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswLagStatusVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswLagStatusVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswLagStatusVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkStatus

`func (o *OswLagStatusVO) GetLinkStatus() int32`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *OswLagStatusVO) GetLinkStatusOk() (*int32, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *OswLagStatusVO) SetLinkStatus(v int32)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *OswLagStatusVO) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetName

`func (o *OswLagStatusVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswLagStatusVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswLagStatusVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswLagStatusVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *OswLagStatusVO) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswLagStatusVO) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswLagStatusVO) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswLagStatusVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetRx

`func (o *OswLagStatusVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *OswLagStatusVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *OswLagStatusVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *OswLagStatusVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxRate

`func (o *OswLagStatusVO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OswLagStatusVO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OswLagStatusVO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OswLagStatusVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetStandardPorts

`func (o *OswLagStatusVO) GetStandardPorts() []OswStandPortVO`

GetStandardPorts returns the StandardPorts field if non-nil, zero value otherwise.

### GetStandardPortsOk

`func (o *OswLagStatusVO) GetStandardPortsOk() (*[]OswStandPortVO, bool)`

GetStandardPortsOk returns a tuple with the StandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPorts

`func (o *OswLagStatusVO) SetStandardPorts(v []OswStandPortVO)`

SetStandardPorts sets StandardPorts field to given value.

### HasStandardPorts

`func (o *OswLagStatusVO) HasStandardPorts() bool`

HasStandardPorts returns a boolean if a field has been set.

### GetStpDiscarding

`func (o *OswLagStatusVO) GetStpDiscarding() bool`

GetStpDiscarding returns the StpDiscarding field if non-nil, zero value otherwise.

### GetStpDiscardingOk

`func (o *OswLagStatusVO) GetStpDiscardingOk() (*bool, bool)`

GetStpDiscardingOk returns a tuple with the StpDiscarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpDiscarding

`func (o *OswLagStatusVO) SetStpDiscarding(v bool)`

SetStpDiscarding sets StpDiscarding field to given value.

### HasStpDiscarding

`func (o *OswLagStatusVO) HasStpDiscarding() bool`

HasStpDiscarding returns a boolean if a field has been set.

### GetTx

`func (o *OswLagStatusVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *OswLagStatusVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *OswLagStatusVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *OswLagStatusVO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxRate

`func (o *OswLagStatusVO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OswLagStatusVO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OswLagStatusVO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OswLagStatusVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


