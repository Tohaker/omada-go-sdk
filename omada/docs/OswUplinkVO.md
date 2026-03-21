# OswUplinkVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedType** | Pointer to **int32** | Blocked Type | [optional] 
**BlockedVlans** | Pointer to **string** | Blocked Vlans | [optional] 
**Duplex** | Pointer to **int32** | Duplex | [optional] 
**HwVersion** | Pointer to **string** | HwVersion | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**LagId** | Pointer to **int32** | LagId. If it is not null, it indicates that the port is a LAG port | [optional] 
**LinkPeerPortInfo** | Pointer to [**OswPeerPortVO**](OswPeerPortVO.md) |  | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed | [optional] 
**Mac** | Pointer to **string** | The mac of uplink device | [optional] 
**Model** | Pointer to **string** | The model of uplink device | [optional] 
**ModelVersion** | Pointer to **string** | Model Version | [optional] 
**Name** | Pointer to **string** | The name of uplink device | [optional] 
**Port** | Pointer to **int32** | Port | [optional] 
**Rx** | Pointer to **int64** | Port total rx bytes | [optional] 
**RxRate** | Pointer to **int64** | Rx Rate | [optional] 
**StPort** | Pointer to **string** | Standard Port, unit/slot/port | [optional] 
**StackId** | Pointer to **string** | The ID of uplink stack device | [optional] 
**StackName** | Pointer to **string** | The name of uplink stack device | [optional] 
**StpDiscarding** | Pointer to **bool** | STP Discarding | [optional] 
**Tx** | Pointer to **int64** | Port total tx bytes | [optional] 
**TxRate** | Pointer to **int64** | Tx Rate | [optional] 
**Type** | Pointer to **string** | The type of uplink device, it should be a value as follows: ap, gateway, switch(stack) | [optional] 

## Methods

### NewOswUplinkVO

`func NewOswUplinkVO() *OswUplinkVO`

NewOswUplinkVO instantiates a new OswUplinkVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswUplinkVOWithDefaults

`func NewOswUplinkVOWithDefaults() *OswUplinkVO`

NewOswUplinkVOWithDefaults instantiates a new OswUplinkVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedType

`func (o *OswUplinkVO) GetBlockedType() int32`

GetBlockedType returns the BlockedType field if non-nil, zero value otherwise.

### GetBlockedTypeOk

`func (o *OswUplinkVO) GetBlockedTypeOk() (*int32, bool)`

GetBlockedTypeOk returns a tuple with the BlockedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedType

`func (o *OswUplinkVO) SetBlockedType(v int32)`

SetBlockedType sets BlockedType field to given value.

### HasBlockedType

`func (o *OswUplinkVO) HasBlockedType() bool`

HasBlockedType returns a boolean if a field has been set.

### GetBlockedVlans

`func (o *OswUplinkVO) GetBlockedVlans() string`

GetBlockedVlans returns the BlockedVlans field if non-nil, zero value otherwise.

### GetBlockedVlansOk

`func (o *OswUplinkVO) GetBlockedVlansOk() (*string, bool)`

GetBlockedVlansOk returns a tuple with the BlockedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedVlans

`func (o *OswUplinkVO) SetBlockedVlans(v string)`

SetBlockedVlans sets BlockedVlans field to given value.

### HasBlockedVlans

`func (o *OswUplinkVO) HasBlockedVlans() bool`

HasBlockedVlans returns a boolean if a field has been set.

### GetDuplex

`func (o *OswUplinkVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswUplinkVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswUplinkVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswUplinkVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetHwVersion

`func (o *OswUplinkVO) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *OswUplinkVO) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *OswUplinkVO) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *OswUplinkVO) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetIp

`func (o *OswUplinkVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswUplinkVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswUplinkVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswUplinkVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLagId

`func (o *OswUplinkVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *OswUplinkVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *OswUplinkVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *OswUplinkVO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLinkPeerPortInfo

`func (o *OswUplinkVO) GetLinkPeerPortInfo() OswPeerPortVO`

GetLinkPeerPortInfo returns the LinkPeerPortInfo field if non-nil, zero value otherwise.

### GetLinkPeerPortInfoOk

`func (o *OswUplinkVO) GetLinkPeerPortInfoOk() (*OswPeerPortVO, bool)`

GetLinkPeerPortInfoOk returns a tuple with the LinkPeerPortInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeerPortInfo

`func (o *OswUplinkVO) SetLinkPeerPortInfo(v OswPeerPortVO)`

SetLinkPeerPortInfo sets LinkPeerPortInfo field to given value.

### HasLinkPeerPortInfo

`func (o *OswUplinkVO) HasLinkPeerPortInfo() bool`

HasLinkPeerPortInfo returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswUplinkVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswUplinkVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswUplinkVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswUplinkVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetMac

`func (o *OswUplinkVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswUplinkVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswUplinkVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswUplinkVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *OswUplinkVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswUplinkVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswUplinkVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswUplinkVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswUplinkVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswUplinkVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswUplinkVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswUplinkVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *OswUplinkVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswUplinkVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswUplinkVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswUplinkVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *OswUplinkVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswUplinkVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswUplinkVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswUplinkVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRx

`func (o *OswUplinkVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *OswUplinkVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *OswUplinkVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *OswUplinkVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxRate

`func (o *OswUplinkVO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OswUplinkVO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OswUplinkVO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OswUplinkVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetStPort

`func (o *OswUplinkVO) GetStPort() string`

GetStPort returns the StPort field if non-nil, zero value otherwise.

### GetStPortOk

`func (o *OswUplinkVO) GetStPortOk() (*string, bool)`

GetStPortOk returns a tuple with the StPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStPort

`func (o *OswUplinkVO) SetStPort(v string)`

SetStPort sets StPort field to given value.

### HasStPort

`func (o *OswUplinkVO) HasStPort() bool`

HasStPort returns a boolean if a field has been set.

### GetStackId

`func (o *OswUplinkVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswUplinkVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswUplinkVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswUplinkVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *OswUplinkVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *OswUplinkVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *OswUplinkVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *OswUplinkVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStpDiscarding

`func (o *OswUplinkVO) GetStpDiscarding() bool`

GetStpDiscarding returns the StpDiscarding field if non-nil, zero value otherwise.

### GetStpDiscardingOk

`func (o *OswUplinkVO) GetStpDiscardingOk() (*bool, bool)`

GetStpDiscardingOk returns a tuple with the StpDiscarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpDiscarding

`func (o *OswUplinkVO) SetStpDiscarding(v bool)`

SetStpDiscarding sets StpDiscarding field to given value.

### HasStpDiscarding

`func (o *OswUplinkVO) HasStpDiscarding() bool`

HasStpDiscarding returns a boolean if a field has been set.

### GetTx

`func (o *OswUplinkVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *OswUplinkVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *OswUplinkVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *OswUplinkVO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxRate

`func (o *OswUplinkVO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OswUplinkVO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OswUplinkVO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OswUplinkVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetType

`func (o *OswUplinkVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswUplinkVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswUplinkVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OswUplinkVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


