# OswDownlinkVO

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
**Mac** | Pointer to **string** | The mac of downlink device | [optional] 
**Model** | Pointer to **string** | The model of uplink device | [optional] 
**ModelVersion** | Pointer to **string** | Model Version | [optional] 
**Name** | Pointer to **string** | The name of downlink device | [optional] 
**Port** | Pointer to **int32** | Port | [optional] 
**Rx** | Pointer to **int64** | Port total rx bytes | [optional] 
**RxRate** | Pointer to **int64** | Port rx Rate | [optional] 
**StPort** | Pointer to **string** | Standard Port, unit/slot/port | [optional] 
**StackId** | Pointer to **string** | The ID of uplink stack device | [optional] 
**StackName** | Pointer to **string** | The name of uplink stack device | [optional] 
**StpDiscarding** | Pointer to **bool** | STP Discarding | [optional] 
**Tx** | Pointer to **int64** | Port total tx bytes | [optional] 
**TxRate** | Pointer to **int64** | Port tx Rate | [optional] 
**Type** | Pointer to **string** | The type of downlink device, it should be a value as follows: ap, gateway, switch(stack) | [optional] 

## Methods

### NewOswDownlinkVO

`func NewOswDownlinkVO() *OswDownlinkVO`

NewOswDownlinkVO instantiates a new OswDownlinkVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswDownlinkVOWithDefaults

`func NewOswDownlinkVOWithDefaults() *OswDownlinkVO`

NewOswDownlinkVOWithDefaults instantiates a new OswDownlinkVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedType

`func (o *OswDownlinkVO) GetBlockedType() int32`

GetBlockedType returns the BlockedType field if non-nil, zero value otherwise.

### GetBlockedTypeOk

`func (o *OswDownlinkVO) GetBlockedTypeOk() (*int32, bool)`

GetBlockedTypeOk returns a tuple with the BlockedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedType

`func (o *OswDownlinkVO) SetBlockedType(v int32)`

SetBlockedType sets BlockedType field to given value.

### HasBlockedType

`func (o *OswDownlinkVO) HasBlockedType() bool`

HasBlockedType returns a boolean if a field has been set.

### GetBlockedVlans

`func (o *OswDownlinkVO) GetBlockedVlans() string`

GetBlockedVlans returns the BlockedVlans field if non-nil, zero value otherwise.

### GetBlockedVlansOk

`func (o *OswDownlinkVO) GetBlockedVlansOk() (*string, bool)`

GetBlockedVlansOk returns a tuple with the BlockedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedVlans

`func (o *OswDownlinkVO) SetBlockedVlans(v string)`

SetBlockedVlans sets BlockedVlans field to given value.

### HasBlockedVlans

`func (o *OswDownlinkVO) HasBlockedVlans() bool`

HasBlockedVlans returns a boolean if a field has been set.

### GetDuplex

`func (o *OswDownlinkVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswDownlinkVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswDownlinkVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswDownlinkVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetHwVersion

`func (o *OswDownlinkVO) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *OswDownlinkVO) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *OswDownlinkVO) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *OswDownlinkVO) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetIp

`func (o *OswDownlinkVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswDownlinkVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswDownlinkVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswDownlinkVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLagId

`func (o *OswDownlinkVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *OswDownlinkVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *OswDownlinkVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *OswDownlinkVO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLinkPeerPortInfo

`func (o *OswDownlinkVO) GetLinkPeerPortInfo() OswPeerPortVO`

GetLinkPeerPortInfo returns the LinkPeerPortInfo field if non-nil, zero value otherwise.

### GetLinkPeerPortInfoOk

`func (o *OswDownlinkVO) GetLinkPeerPortInfoOk() (*OswPeerPortVO, bool)`

GetLinkPeerPortInfoOk returns a tuple with the LinkPeerPortInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeerPortInfo

`func (o *OswDownlinkVO) SetLinkPeerPortInfo(v OswPeerPortVO)`

SetLinkPeerPortInfo sets LinkPeerPortInfo field to given value.

### HasLinkPeerPortInfo

`func (o *OswDownlinkVO) HasLinkPeerPortInfo() bool`

HasLinkPeerPortInfo returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswDownlinkVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswDownlinkVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswDownlinkVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswDownlinkVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetMac

`func (o *OswDownlinkVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswDownlinkVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswDownlinkVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswDownlinkVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *OswDownlinkVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswDownlinkVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswDownlinkVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswDownlinkVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswDownlinkVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswDownlinkVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswDownlinkVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswDownlinkVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *OswDownlinkVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswDownlinkVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswDownlinkVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswDownlinkVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *OswDownlinkVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswDownlinkVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswDownlinkVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswDownlinkVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRx

`func (o *OswDownlinkVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *OswDownlinkVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *OswDownlinkVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *OswDownlinkVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetRxRate

`func (o *OswDownlinkVO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OswDownlinkVO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OswDownlinkVO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OswDownlinkVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetStPort

`func (o *OswDownlinkVO) GetStPort() string`

GetStPort returns the StPort field if non-nil, zero value otherwise.

### GetStPortOk

`func (o *OswDownlinkVO) GetStPortOk() (*string, bool)`

GetStPortOk returns a tuple with the StPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStPort

`func (o *OswDownlinkVO) SetStPort(v string)`

SetStPort sets StPort field to given value.

### HasStPort

`func (o *OswDownlinkVO) HasStPort() bool`

HasStPort returns a boolean if a field has been set.

### GetStackId

`func (o *OswDownlinkVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswDownlinkVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswDownlinkVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswDownlinkVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *OswDownlinkVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *OswDownlinkVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *OswDownlinkVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *OswDownlinkVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStpDiscarding

`func (o *OswDownlinkVO) GetStpDiscarding() bool`

GetStpDiscarding returns the StpDiscarding field if non-nil, zero value otherwise.

### GetStpDiscardingOk

`func (o *OswDownlinkVO) GetStpDiscardingOk() (*bool, bool)`

GetStpDiscardingOk returns a tuple with the StpDiscarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpDiscarding

`func (o *OswDownlinkVO) SetStpDiscarding(v bool)`

SetStpDiscarding sets StpDiscarding field to given value.

### HasStpDiscarding

`func (o *OswDownlinkVO) HasStpDiscarding() bool`

HasStpDiscarding returns a boolean if a field has been set.

### GetTx

`func (o *OswDownlinkVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *OswDownlinkVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *OswDownlinkVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *OswDownlinkVO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetTxRate

`func (o *OswDownlinkVO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OswDownlinkVO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OswDownlinkVO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OswDownlinkVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetType

`func (o *OswDownlinkVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswDownlinkVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswDownlinkVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OswDownlinkVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


