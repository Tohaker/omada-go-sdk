# SwitchInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplex** | Pointer to **int32** | Duplex: 1:Half, 2:Full. | [optional] 
**LinkSpeed** | Pointer to **int32** | LinkSpeed: 1:10Mbps, 2:100Mbps, 3:1000Mbps, 4:2.5Gbps, 5:10Gbps. | [optional] 
**RxDropPkts** | Pointer to **int64** | Rx Dropped Packets | [optional] 
**RxErrPkts** | Pointer to **int64** | Rx Error Packets | [optional] 
**TxDropPkts** | Pointer to **int64** | Tx Dropped Packets | [optional] 
**TxErrPkts** | Pointer to **int64** | Tx Error Packets | [optional] 
**UpLinkPort** | Pointer to [**WiredPortV3DTO**](WiredPortV3DTO.md) |  | [optional] 
**UpPort** | Pointer to [**WiredPortV3DTO**](WiredPortV3DTO.md) |  | [optional] 

## Methods

### NewSwitchInfo

`func NewSwitchInfo() *SwitchInfo`

NewSwitchInfo instantiates a new SwitchInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchInfoWithDefaults

`func NewSwitchInfoWithDefaults() *SwitchInfo`

NewSwitchInfoWithDefaults instantiates a new SwitchInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplex

`func (o *SwitchInfo) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *SwitchInfo) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *SwitchInfo) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *SwitchInfo) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *SwitchInfo) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *SwitchInfo) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *SwitchInfo) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *SwitchInfo) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetRxDropPkts

`func (o *SwitchInfo) GetRxDropPkts() int64`

GetRxDropPkts returns the RxDropPkts field if non-nil, zero value otherwise.

### GetRxDropPktsOk

`func (o *SwitchInfo) GetRxDropPktsOk() (*int64, bool)`

GetRxDropPktsOk returns a tuple with the RxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDropPkts

`func (o *SwitchInfo) SetRxDropPkts(v int64)`

SetRxDropPkts sets RxDropPkts field to given value.

### HasRxDropPkts

`func (o *SwitchInfo) HasRxDropPkts() bool`

HasRxDropPkts returns a boolean if a field has been set.

### GetRxErrPkts

`func (o *SwitchInfo) GetRxErrPkts() int64`

GetRxErrPkts returns the RxErrPkts field if non-nil, zero value otherwise.

### GetRxErrPktsOk

`func (o *SwitchInfo) GetRxErrPktsOk() (*int64, bool)`

GetRxErrPktsOk returns a tuple with the RxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrPkts

`func (o *SwitchInfo) SetRxErrPkts(v int64)`

SetRxErrPkts sets RxErrPkts field to given value.

### HasRxErrPkts

`func (o *SwitchInfo) HasRxErrPkts() bool`

HasRxErrPkts returns a boolean if a field has been set.

### GetTxDropPkts

`func (o *SwitchInfo) GetTxDropPkts() int64`

GetTxDropPkts returns the TxDropPkts field if non-nil, zero value otherwise.

### GetTxDropPktsOk

`func (o *SwitchInfo) GetTxDropPktsOk() (*int64, bool)`

GetTxDropPktsOk returns a tuple with the TxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDropPkts

`func (o *SwitchInfo) SetTxDropPkts(v int64)`

SetTxDropPkts sets TxDropPkts field to given value.

### HasTxDropPkts

`func (o *SwitchInfo) HasTxDropPkts() bool`

HasTxDropPkts returns a boolean if a field has been set.

### GetTxErrPkts

`func (o *SwitchInfo) GetTxErrPkts() int64`

GetTxErrPkts returns the TxErrPkts field if non-nil, zero value otherwise.

### GetTxErrPktsOk

`func (o *SwitchInfo) GetTxErrPktsOk() (*int64, bool)`

GetTxErrPktsOk returns a tuple with the TxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrPkts

`func (o *SwitchInfo) SetTxErrPkts(v int64)`

SetTxErrPkts sets TxErrPkts field to given value.

### HasTxErrPkts

`func (o *SwitchInfo) HasTxErrPkts() bool`

HasTxErrPkts returns a boolean if a field has been set.

### GetUpLinkPort

`func (o *SwitchInfo) GetUpLinkPort() WiredPortV3DTO`

GetUpLinkPort returns the UpLinkPort field if non-nil, zero value otherwise.

### GetUpLinkPortOk

`func (o *SwitchInfo) GetUpLinkPortOk() (*WiredPortV3DTO, bool)`

GetUpLinkPortOk returns a tuple with the UpLinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLinkPort

`func (o *SwitchInfo) SetUpLinkPort(v WiredPortV3DTO)`

SetUpLinkPort sets UpLinkPort field to given value.

### HasUpLinkPort

`func (o *SwitchInfo) HasUpLinkPort() bool`

HasUpLinkPort returns a boolean if a field has been set.

### GetUpPort

`func (o *SwitchInfo) GetUpPort() WiredPortV3DTO`

GetUpPort returns the UpPort field if non-nil, zero value otherwise.

### GetUpPortOk

`func (o *SwitchInfo) GetUpPortOk() (*WiredPortV3DTO, bool)`

GetUpPortOk returns a tuple with the UpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPort

`func (o *SwitchInfo) SetUpPort(v WiredPortV3DTO)`

SetUpPort sets UpPort field to given value.

### HasUpPort

`func (o *SwitchInfo) HasUpPort() bool`

HasUpPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


