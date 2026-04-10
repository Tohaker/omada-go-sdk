# WanList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **int32** | IPv6 enable. 0:\&quot;Disable\&quot;, 1:\&quot;Enable\&quot; | [optional] 
**NetworkId** | Pointer to **string** | If networkId is not null, indicates this wan is used by specific \&quot;Lan Network\&quot;. | [optional] 
**PdEnable** | Pointer to **int32** | Prefix Delegation. 0:\&quot;Disable\&quot;, 1:\&quot;Enable\&quot;(Default) | [optional] 
**PdSize** | Pointer to **int32** | Prefix size | [optional] 
**PortId** | Pointer to **string** | Wan port ID. | [optional] 
**PortName** | Pointer to **string** | Wan port name. | [optional] 
**Prefix** | Pointer to **string** | True prefix of wan port. | [optional] 
**Proto** | Pointer to **int32** | IPv6 connection type of wan port. Includes: 0:\&quot;static\&quot;, 1:\&quot;dynamic\&quot;, 2:\&quot;pppoe\&quot;, 3:\&quot;6to4Tunnel\&quot;, 4:\&quot;bridge\&quot;, 7:\&quot;pppoa\&quot;, 8:\&quot;ipoa\&quot;. Only type is \&quot;bridge\&quot;, can be selected by passthrough(Lan Network mode) | [optional] 

## Methods

### NewWanList

`func NewWanList() *WanList`

NewWanList instantiates a new WanList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanListWithDefaults

`func NewWanListWithDefaults() *WanList`

NewWanListWithDefaults instantiates a new WanList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *WanList) GetEnable() int32`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *WanList) GetEnableOk() (*int32, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *WanList) SetEnable(v int32)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *WanList) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetNetworkId

`func (o *WanList) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *WanList) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *WanList) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *WanList) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetPdEnable

`func (o *WanList) GetPdEnable() int32`

GetPdEnable returns the PdEnable field if non-nil, zero value otherwise.

### GetPdEnableOk

`func (o *WanList) GetPdEnableOk() (*int32, bool)`

GetPdEnableOk returns a tuple with the PdEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdEnable

`func (o *WanList) SetPdEnable(v int32)`

SetPdEnable sets PdEnable field to given value.

### HasPdEnable

`func (o *WanList) HasPdEnable() bool`

HasPdEnable returns a boolean if a field has been set.

### GetPdSize

`func (o *WanList) GetPdSize() int32`

GetPdSize returns the PdSize field if non-nil, zero value otherwise.

### GetPdSizeOk

`func (o *WanList) GetPdSizeOk() (*int32, bool)`

GetPdSizeOk returns a tuple with the PdSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdSize

`func (o *WanList) SetPdSize(v int32)`

SetPdSize sets PdSize field to given value.

### HasPdSize

`func (o *WanList) HasPdSize() bool`

HasPdSize returns a boolean if a field has been set.

### GetPortId

`func (o *WanList) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *WanList) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *WanList) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *WanList) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortName

`func (o *WanList) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *WanList) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *WanList) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *WanList) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPrefix

`func (o *WanList) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *WanList) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *WanList) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *WanList) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetProto

`func (o *WanList) GetProto() int32`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *WanList) GetProtoOk() (*int32, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *WanList) SetProto(v int32)`

SetProto sets Proto field to given value.

### HasProto

`func (o *WanList) HasProto() bool`

HasProto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


