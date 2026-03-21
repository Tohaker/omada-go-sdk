# OswPortStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedType** | Pointer to **int32** | Blocked Type | [optional] 
**BlockedVlans** | Pointer to **string** | Blocked Vlans | [optional] 
**ChassisId** | Pointer to **string** | Chassis ID | [optional] 
**ChassisIdSubtype** | Pointer to **string** | Chassis ID Subtype | [optional] 
**Duplex** | Pointer to **int32** | Duplex should be a value as follows: 1: Half; 2: Full | [optional] 
**FecMode** | Pointer to **int32** | The configured FEC Mode | [optional] 
**FecRealMode** | Pointer to **int32** | The actual effective FEC Mode | [optional] 
**LinkSpeed** | Pointer to **int32** | LinkSpeed should be a value as follows: 1: 10Mbps; 2: 100Mbps; 3: 1000Mbps; 4: 10Gbps | [optional] 
**LinkStatus** | Pointer to **int32** | LinkStatus should be a value as follows: 0: link down; 1: link up | [optional] 
**Poe** | Pointer to **bool** | Indicates whether PoE power supply is in use | [optional] 
**PoePower** | Pointer to **float64** | PoE Power | [optional] 
**Port** | Pointer to **int32** | Port Number | [optional] 
**PortId** | Pointer to **string** | Port ID | [optional] 
**PortIdSubtype** | Pointer to **string** | Port ID Subtype | [optional] 
**Rx** | Pointer to **int64** | Rx | [optional] 
**RxRate** | Pointer to **int64** | Rx Rate | [optional] 
**StkStatus** | Pointer to **int32** | Stk Status | [optional] 
**Stp** | Pointer to **string** | STP Status | [optional] 
**StpDiscarding** | Pointer to **bool** | STP Discarding | [optional] 
**Tx** | Pointer to **int64** | Tx | [optional] 
**TxRate** | Pointer to **int64** | Tx Rate | [optional] 

## Methods

### NewOswPortStatusVO

`func NewOswPortStatusVO() *OswPortStatusVO`

NewOswPortStatusVO instantiates a new OswPortStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPortStatusVOWithDefaults

`func NewOswPortStatusVOWithDefaults() *OswPortStatusVO`

NewOswPortStatusVOWithDefaults instantiates a new OswPortStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedType

`func (o *OswPortStatusVO) GetBlockedType() int32`

GetBlockedType returns the BlockedType field if non-nil, zero value otherwise.

### GetBlockedTypeOk

`func (o *OswPortStatusVO) GetBlockedTypeOk() (*int32, bool)`

GetBlockedTypeOk returns a tuple with the BlockedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedType

`func (o *OswPortStatusVO) SetBlockedType(v int32)`

SetBlockedType sets BlockedType field to given value.

### HasBlockedType

`func (o *OswPortStatusVO) HasBlockedType() bool`

HasBlockedType returns a boolean if a field has been set.

### GetBlockedVlans

`func (o *OswPortStatusVO) GetBlockedVlans() string`

GetBlockedVlans returns the BlockedVlans field if non-nil, zero value otherwise.

### GetBlockedVlansOk

`func (o *OswPortStatusVO) GetBlockedVlansOk() (*string, bool)`

GetBlockedVlansOk returns a tuple with the BlockedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedVlans

`func (o *OswPortStatusVO) SetBlockedVlans(v string)`

SetBlockedVlans sets BlockedVlans field to given value.

### HasBlockedVlans

`func (o *OswPortStatusVO) HasBlockedVlans() bool`

HasBlockedVlans returns a boolean if a field has been set.

### GetChassisId

`func (o *OswPortStatusVO) GetChassisId() string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *OswPortStatusVO) GetChassisIdOk() (*string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *OswPortStatusVO) SetChassisId(v string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *OswPortStatusVO) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetChassisIdSubtype

`func (o *OswPortStatusVO) GetChassisIdSubtype() string`

GetChassisIdSubtype returns the ChassisIdSubtype field if non-nil, zero value otherwise.

### GetChassisIdSubtypeOk

`func (o *OswPortStatusVO) GetChassisIdSubtypeOk() (*string, bool)`

GetChassisIdSubtypeOk returns a tuple with the ChassisIdSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisIdSubtype

`func (o *OswPortStatusVO) SetChassisIdSubtype(v string)`

SetChassisIdSubtype sets ChassisIdSubtype field to given value.

### HasChassisIdSubtype

`func (o *OswPortStatusVO) HasChassisIdSubtype() bool`

HasChassisIdSubtype returns a boolean if a field has been set.

### GetDuplex

`func (o *OswPortStatusVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswPortStatusVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswPortStatusVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswPortStatusVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetFecMode

`func (o *OswPortStatusVO) GetFecMode() int32`

GetFecMode returns the FecMode field if non-nil, zero value otherwise.

### GetFecModeOk

`func (o *OswPortStatusVO) GetFecModeOk() (*int32, bool)`

GetFecModeOk returns a tuple with the FecMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecMode

`func (o *OswPortStatusVO) SetFecMode(v int32)`

SetFecMode sets FecMode field to given value.

### HasFecMode

`func (o *OswPortStatusVO) HasFecMode() bool`

HasFecMode returns a boolean if a field has been set.

### GetFecRealMode

`func (o *OswPortStatusVO) GetFecRealMode() int32`

GetFecRealMode returns the FecRealMode field if non-nil, zero value otherwise.

### GetFecRealModeOk

`func (o *OswPortStatusVO) GetFecRealModeOk() (*int32, bool)`

GetFecRealModeOk returns a tuple with the FecRealMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecRealMode

`func (o *OswPortStatusVO) SetFecRealMode(v int32)`

SetFecRealMode sets FecRealMode field to given value.

### HasFecRealMode

`func (o *OswPortStatusVO) HasFecRealMode() bool`

HasFecRealMode returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswPortStatusVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswPortStatusVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswPortStatusVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswPortStatusVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkStatus

`func (o *OswPortStatusVO) GetLinkStatus() int32`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *OswPortStatusVO) GetLinkStatusOk() (*int32, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *OswPortStatusVO) SetLinkStatus(v int32)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *OswPortStatusVO) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetPoe

`func (o *OswPortStatusVO) GetPoe() bool`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *OswPortStatusVO) GetPoeOk() (*bool, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *OswPortStatusVO) SetPoe(v bool)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *OswPortStatusVO) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetPoePower

`func (o *OswPortStatusVO) GetPoePower() float64`

GetPoePower returns the PoePower field if non-nil, zero value otherwise.

### GetPoePowerOk

`func (o *OswPortStatusVO) GetPoePowerOk() (*float64, bool)`

GetPoePowerOk returns a tuple with the PoePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePower

`func (o *OswPortStatusVO) SetPoePower(v float64)`

SetPoePower sets PoePower field to given value.

### HasPoePower

`func (o *OswPortStatusVO) HasPoePower() bool`

HasPoePower returns a boolean if a field has been set.

### GetPort

`func (o *OswPortStatusVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswPortStatusVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswPortStatusVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswPortStatusVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortId

`func (o *OswPortStatusVO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *OswPortStatusVO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *OswPortStatusVO) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *OswPortStatusVO) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortIdSubtype

`func (o *OswPortStatusVO) GetPortIdSubtype() string`

GetPortIdSubtype returns the PortIdSubtype field if non-nil, zero value otherwise.

### GetPortIdSubtypeOk

`func (o *OswPortStatusVO) GetPortIdSubtypeOk() (*string, bool)`

GetPortIdSubtypeOk returns a tuple with the PortIdSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIdSubtype

`func (o *OswPortStatusVO) SetPortIdSubtype(v string)`

SetPortIdSubtype sets PortIdSubtype field to given value.

### HasPortIdSubtype

`func (o *OswPortStatusVO) HasPortIdSubtype() bool`

HasPortIdSubtype returns a boolean if a field has been set.

### GetRx

`func (o *OswPortStatusVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *OswPortStatusVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *OswPortStatusVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *OswPortStatusVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxRate

`func (o *OswPortStatusVO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OswPortStatusVO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OswPortStatusVO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OswPortStatusVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetStkStatus

`func (o *OswPortStatusVO) GetStkStatus() int32`

GetStkStatus returns the StkStatus field if non-nil, zero value otherwise.

### GetStkStatusOk

`func (o *OswPortStatusVO) GetStkStatusOk() (*int32, bool)`

GetStkStatusOk returns a tuple with the StkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkStatus

`func (o *OswPortStatusVO) SetStkStatus(v int32)`

SetStkStatus sets StkStatus field to given value.

### HasStkStatus

`func (o *OswPortStatusVO) HasStkStatus() bool`

HasStkStatus returns a boolean if a field has been set.

### GetStp

`func (o *OswPortStatusVO) GetStp() string`

GetStp returns the Stp field if non-nil, zero value otherwise.

### GetStpOk

`func (o *OswPortStatusVO) GetStpOk() (*string, bool)`

GetStpOk returns a tuple with the Stp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStp

`func (o *OswPortStatusVO) SetStp(v string)`

SetStp sets Stp field to given value.

### HasStp

`func (o *OswPortStatusVO) HasStp() bool`

HasStp returns a boolean if a field has been set.

### GetStpDiscarding

`func (o *OswPortStatusVO) GetStpDiscarding() bool`

GetStpDiscarding returns the StpDiscarding field if non-nil, zero value otherwise.

### GetStpDiscardingOk

`func (o *OswPortStatusVO) GetStpDiscardingOk() (*bool, bool)`

GetStpDiscardingOk returns a tuple with the StpDiscarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpDiscarding

`func (o *OswPortStatusVO) SetStpDiscarding(v bool)`

SetStpDiscarding sets StpDiscarding field to given value.

### HasStpDiscarding

`func (o *OswPortStatusVO) HasStpDiscarding() bool`

HasStpDiscarding returns a boolean if a field has been set.

### GetTx

`func (o *OswPortStatusVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *OswPortStatusVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *OswPortStatusVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *OswPortStatusVO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxRate

`func (o *OswPortStatusVO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OswPortStatusVO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OswPortStatusVO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OswPortStatusVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


