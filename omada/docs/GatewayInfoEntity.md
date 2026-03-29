# GatewayInfoEntity

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

### NewGatewayInfoEntity

`func NewGatewayInfoEntity() *GatewayInfoEntity`

NewGatewayInfoEntity instantiates a new GatewayInfoEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayInfoEntityWithDefaults

`func NewGatewayInfoEntityWithDefaults() *GatewayInfoEntity`

NewGatewayInfoEntityWithDefaults instantiates a new GatewayInfoEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplex

`func (o *GatewayInfoEntity) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *GatewayInfoEntity) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *GatewayInfoEntity) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *GatewayInfoEntity) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *GatewayInfoEntity) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *GatewayInfoEntity) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *GatewayInfoEntity) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *GatewayInfoEntity) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetRxDropPkts

`func (o *GatewayInfoEntity) GetRxDropPkts() int64`

GetRxDropPkts returns the RxDropPkts field if non-nil, zero value otherwise.

### GetRxDropPktsOk

`func (o *GatewayInfoEntity) GetRxDropPktsOk() (*int64, bool)`

GetRxDropPktsOk returns a tuple with the RxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDropPkts

`func (o *GatewayInfoEntity) SetRxDropPkts(v int64)`

SetRxDropPkts sets RxDropPkts field to given value.

### HasRxDropPkts

`func (o *GatewayInfoEntity) HasRxDropPkts() bool`

HasRxDropPkts returns a boolean if a field has been set.

### GetRxErrPkts

`func (o *GatewayInfoEntity) GetRxErrPkts() int64`

GetRxErrPkts returns the RxErrPkts field if non-nil, zero value otherwise.

### GetRxErrPktsOk

`func (o *GatewayInfoEntity) GetRxErrPktsOk() (*int64, bool)`

GetRxErrPktsOk returns a tuple with the RxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrPkts

`func (o *GatewayInfoEntity) SetRxErrPkts(v int64)`

SetRxErrPkts sets RxErrPkts field to given value.

### HasRxErrPkts

`func (o *GatewayInfoEntity) HasRxErrPkts() bool`

HasRxErrPkts returns a boolean if a field has been set.

### GetTxDropPkts

`func (o *GatewayInfoEntity) GetTxDropPkts() int64`

GetTxDropPkts returns the TxDropPkts field if non-nil, zero value otherwise.

### GetTxDropPktsOk

`func (o *GatewayInfoEntity) GetTxDropPktsOk() (*int64, bool)`

GetTxDropPktsOk returns a tuple with the TxDropPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDropPkts

`func (o *GatewayInfoEntity) SetTxDropPkts(v int64)`

SetTxDropPkts sets TxDropPkts field to given value.

### HasTxDropPkts

`func (o *GatewayInfoEntity) HasTxDropPkts() bool`

HasTxDropPkts returns a boolean if a field has been set.

### GetTxErrPkts

`func (o *GatewayInfoEntity) GetTxErrPkts() int64`

GetTxErrPkts returns the TxErrPkts field if non-nil, zero value otherwise.

### GetTxErrPktsOk

`func (o *GatewayInfoEntity) GetTxErrPktsOk() (*int64, bool)`

GetTxErrPktsOk returns a tuple with the TxErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrPkts

`func (o *GatewayInfoEntity) SetTxErrPkts(v int64)`

SetTxErrPkts sets TxErrPkts field to given value.

### HasTxErrPkts

`func (o *GatewayInfoEntity) HasTxErrPkts() bool`

HasTxErrPkts returns a boolean if a field has been set.

### GetUpLinkPort

`func (o *GatewayInfoEntity) GetUpLinkPort() WiredPortV3DTO`

GetUpLinkPort returns the UpLinkPort field if non-nil, zero value otherwise.

### GetUpLinkPortOk

`func (o *GatewayInfoEntity) GetUpLinkPortOk() (*WiredPortV3DTO, bool)`

GetUpLinkPortOk returns a tuple with the UpLinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLinkPort

`func (o *GatewayInfoEntity) SetUpLinkPort(v WiredPortV3DTO)`

SetUpLinkPort sets UpLinkPort field to given value.

### HasUpLinkPort

`func (o *GatewayInfoEntity) HasUpLinkPort() bool`

HasUpLinkPort returns a boolean if a field has been set.

### GetUpPort

`func (o *GatewayInfoEntity) GetUpPort() WiredPortV3DTO`

GetUpPort returns the UpPort field if non-nil, zero value otherwise.

### GetUpPortOk

`func (o *GatewayInfoEntity) GetUpPortOk() (*WiredPortV3DTO, bool)`

GetUpPortOk returns a tuple with the UpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPort

`func (o *GatewayInfoEntity) SetUpPort(v WiredPortV3DTO)`

SetUpPort sets UpPort field to given value.

### HasUpPort

`func (o *GatewayInfoEntity) HasUpPort() bool`

HasUpPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


