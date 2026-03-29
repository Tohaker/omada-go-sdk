# WiredUpInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to **bool** | Blocked Connection Or Not | [optional] 
**BlockedVlans** | Pointer to **string** | Blocked Vlans | [optional] 
**Duplex** | Pointer to **int32** | Duplex | [optional] 
**FiberOptic** | Pointer to **bool** | Whether The Device Is FiberOptic Or Not | [optional] 
**LinkSpeed** | Pointer to **int32** | LinkSpeed | [optional] 
**Port** | Pointer to [**WiredPortV3DTO**](WiredPortV3DTO.md) |  | [optional] 
**RxDropPkts** | Pointer to **int64** | Rx Dropped Packets | [optional] 
**RxErrPkts** | Pointer to **int64** | Rx Error Packets | [optional] 
**TxDropPkts** | Pointer to **int64** | Tx Dropped Packets | [optional] 
**TxErrPkts** | Pointer to **int64** | Tx Error Packets | [optional] 
**UpLinkPort** | Pointer to [**WiredPortV3DTO**](WiredPortV3DTO.md) |  | [optional] 

## Methods

### NewWiredUpInfoDTO

`func NewWiredUpInfoDTO() *WiredUpInfoDTO`

NewWiredUpInfoDTO instantiates a new WiredUpInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWiredUpInfoDTOWithDefaults

`func NewWiredUpInfoDTOWithDefaults() *WiredUpInfoDTO`

NewWiredUpInfoDTOWithDefaults instantiates a new WiredUpInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *WiredUpInfoDTO) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *WiredUpInfoDTO) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *WiredUpInfoDTO) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *WiredUpInfoDTO) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBlockedVlans

`func (o *WiredUpInfoDTO) GetBlockedVlans() string`

GetBlockedVlans returns the BlockedVlans field if non-nil, zero value otherwise.

### GetBlockedVlansOk

`func (o *WiredUpInfoDTO) GetBlockedVlansOk() (*string, bool)`

GetBlockedVlansOk returns a tuple with the BlockedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedVlans

`func (o *WiredUpInfoDTO) SetBlockedVlans(v string)`

SetBlockedVlans sets BlockedVlans field to given value.

### HasBlockedVlans

`func (o *WiredUpInfoDTO) HasBlockedVlans() bool`

HasBlockedVlans returns a boolean if a field has been set.

### GetDuplex

`func (o *WiredUpInfoDTO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *WiredUpInfoDTO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *WiredUpInfoDTO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *WiredUpInfoDTO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetFiberOptic

`func (o *WiredUpInfoDTO) GetFiberOptic() bool`

GetFiberOptic returns the FiberOptic field if non-nil, zero value otherwise.

### GetFiberOpticOk

`func (o *WiredUpInfoDTO) GetFiberOpticOk() (*bool, bool)`

GetFiberOpticOk returns a tuple with the FiberOptic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiberOptic

`func (o *WiredUpInfoDTO) SetFiberOptic(v bool)`

SetFiberOptic sets FiberOptic field to given value.

### HasFiberOptic

`func (o *WiredUpInfoDTO) HasFiberOptic() bool`

HasFiberOptic returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *WiredUpInfoDTO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *WiredUpInfoDTO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *WiredUpInfoDTO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *WiredUpInfoDTO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetPort

`func (o *WiredUpInfoDTO) GetPort() WiredPortV3DTO`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WiredUpInfoDTO) GetPortOk() (*WiredPortV3DTO, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WiredUpInfoDTO) SetPort(v WiredPortV3DTO)`

SetPort sets Port field to given value.

### HasPort

`func (o *WiredUpInfoDTO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRxDropPkts

`func (o *WiredUpInfoDTO) GetRxDropPkts() int64`

GetRxDropPkts returns the RxDropPkts field if non-nil, zero value otherwise.

### GetRxDropPktsOk

`func (o *WiredUpInfoDTO) GetRxDropPktsOk() (*int64, bool)`

GetRxDropPktsOk returns a tuple with the RxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDropPkts

`func (o *WiredUpInfoDTO) SetRxDropPkts(v int64)`

SetRxDropPkts sets RxDropPkts field to given value.

### HasRxDropPkts

`func (o *WiredUpInfoDTO) HasRxDropPkts() bool`

HasRxDropPkts returns a boolean if a field has been set.

### GetRxErrPkts

`func (o *WiredUpInfoDTO) GetRxErrPkts() int64`

GetRxErrPkts returns the RxErrPkts field if non-nil, zero value otherwise.

### GetRxErrPktsOk

`func (o *WiredUpInfoDTO) GetRxErrPktsOk() (*int64, bool)`

GetRxErrPktsOk returns a tuple with the RxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrPkts

`func (o *WiredUpInfoDTO) SetRxErrPkts(v int64)`

SetRxErrPkts sets RxErrPkts field to given value.

### HasRxErrPkts

`func (o *WiredUpInfoDTO) HasRxErrPkts() bool`

HasRxErrPkts returns a boolean if a field has been set.

### GetTxDropPkts

`func (o *WiredUpInfoDTO) GetTxDropPkts() int64`

GetTxDropPkts returns the TxDropPkts field if non-nil, zero value otherwise.

### GetTxDropPktsOk

`func (o *WiredUpInfoDTO) GetTxDropPktsOk() (*int64, bool)`

GetTxDropPktsOk returns a tuple with the TxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDropPkts

`func (o *WiredUpInfoDTO) SetTxDropPkts(v int64)`

SetTxDropPkts sets TxDropPkts field to given value.

### HasTxDropPkts

`func (o *WiredUpInfoDTO) HasTxDropPkts() bool`

HasTxDropPkts returns a boolean if a field has been set.

### GetTxErrPkts

`func (o *WiredUpInfoDTO) GetTxErrPkts() int64`

GetTxErrPkts returns the TxErrPkts field if non-nil, zero value otherwise.

### GetTxErrPktsOk

`func (o *WiredUpInfoDTO) GetTxErrPktsOk() (*int64, bool)`

GetTxErrPktsOk returns a tuple with the TxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrPkts

`func (o *WiredUpInfoDTO) SetTxErrPkts(v int64)`

SetTxErrPkts sets TxErrPkts field to given value.

### HasTxErrPkts

`func (o *WiredUpInfoDTO) HasTxErrPkts() bool`

HasTxErrPkts returns a boolean if a field has been set.

### GetUpLinkPort

`func (o *WiredUpInfoDTO) GetUpLinkPort() WiredPortV3DTO`

GetUpLinkPort returns the UpLinkPort field if non-nil, zero value otherwise.

### GetUpLinkPortOk

`func (o *WiredUpInfoDTO) GetUpLinkPortOk() (*WiredPortV3DTO, bool)`

GetUpLinkPortOk returns a tuple with the UpLinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLinkPort

`func (o *WiredUpInfoDTO) SetUpLinkPort(v WiredPortV3DTO)`

SetUpLinkPort sets UpLinkPort field to given value.

### HasUpLinkPort

`func (o *WiredUpInfoDTO) HasUpLinkPort() bool`

HasUpLinkPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


