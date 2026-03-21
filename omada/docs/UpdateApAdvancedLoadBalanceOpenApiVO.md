# UpdateApAdvancedLoadBalanceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxClients2g** | Pointer to **int32** | The maximum number of clients connected to the 2G band. MaxClients2g should be within the range of 1–512. | [optional] 
**MaxClients5g** | Pointer to **int32** | The maximum number of clients connected to the 5G band. MaxClients5g should be within the range of 1–512. | [optional] 
**MaxClients5g1** | Pointer to **int32** | The maximum number of clients connected to the 5G1 band. MaxClients5g1 should be within the range of 1–512. | [optional] 
**MaxClients5g2** | Pointer to **int32** | The maximum number of clients connected to the 5G2 band. MaxClients5g2 should be within the range of 1–512. | [optional] 
**MaxClients6g** | Pointer to **int32** | The maximum number of clients connected to the 6G band. MaxClients6g should be within the range of 1–512. | [optional] 
**MaxClientsEnable2g** | Pointer to **bool** | Whether the device enable maximum associated clients 2g. | [optional] 
**MaxClientsEnable5g** | Pointer to **bool** | Whether the device enable maximum associated clients 5g. | [optional] 
**MaxClientsEnable5g1** | Pointer to **bool** | Whether the device enable load balance 5g1. | [optional] 
**MaxClientsEnable5g2** | Pointer to **bool** | Whether the device enable maximum associated clients 5g2. | [optional] 
**MaxClientsEnable6g** | Pointer to **bool** | Whether the device enable maximum associated clients 6g. | [optional] 
**RssiEnable2g** | Pointer to **bool** | Whether the device enable rssi 2g. | [optional] 
**RssiEnable5g** | Pointer to **bool** | Whether the device enable rssi 5g. | [optional] 
**RssiEnable5g1** | Pointer to **bool** | Whether the device enable rssi 5g1. | [optional] 
**RssiEnable5g2** | Pointer to **bool** | Whether the device enable rssi 5g2. | [optional] 
**RssiEnable6g** | Pointer to **bool** | Whether the device enable rssi 6g. | [optional] 
**Threshold2g** | Pointer to **int32** | The RSSI threshold for the 2G band. threshold2g should be within the range of -95–0. | [optional] 
**Threshold5g** | Pointer to **int32** | The RSSI threshold for the 5G band. threshold5g should be within the range of -95–0. | [optional] 
**Threshold5g1** | Pointer to **int32** | The RSSI threshold for the 5G1 band. threshold5g1 should be within the range of -95–0. | [optional] 
**Threshold5g2** | Pointer to **int32** | The RSSI threshold for the 5G2 band. threshold5g2 should be within the range of -95–0. | [optional] 
**Threshold6g** | Pointer to **int32** | The RSSI threshold for the 6G band. threshold6g should be within the range of -95–0. | [optional] 

## Methods

### NewUpdateApAdvancedLoadBalanceOpenApiVO

`func NewUpdateApAdvancedLoadBalanceOpenApiVO() *UpdateApAdvancedLoadBalanceOpenApiVO`

NewUpdateApAdvancedLoadBalanceOpenApiVO instantiates a new UpdateApAdvancedLoadBalanceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApAdvancedLoadBalanceOpenApiVOWithDefaults

`func NewUpdateApAdvancedLoadBalanceOpenApiVOWithDefaults() *UpdateApAdvancedLoadBalanceOpenApiVO`

NewUpdateApAdvancedLoadBalanceOpenApiVOWithDefaults instantiates a new UpdateApAdvancedLoadBalanceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxClients2g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClients2g() int32`

GetMaxClients2g returns the MaxClients2g field if non-nil, zero value otherwise.

### GetMaxClients2gOk

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClients2gOk() (*int32, bool)`

GetMaxClients2gOk returns a tuple with the MaxClients2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClients2g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetMaxClients2g(v int32)`

SetMaxClients2g sets MaxClients2g field to given value.

### HasMaxClients2g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasMaxClients2g() bool`

HasMaxClients2g returns a boolean if a field has been set.

### GetMaxClients5g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClients5g() int32`

GetMaxClients5g returns the MaxClients5g field if non-nil, zero value otherwise.

### GetMaxClients5gOk

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClients5gOk() (*int32, bool)`

GetMaxClients5gOk returns a tuple with the MaxClients5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClients5g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetMaxClients5g(v int32)`

SetMaxClients5g sets MaxClients5g field to given value.

### HasMaxClients5g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasMaxClients5g() bool`

HasMaxClients5g returns a boolean if a field has been set.

### GetMaxClients5g1

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClients5g1() int32`

GetMaxClients5g1 returns the MaxClients5g1 field if non-nil, zero value otherwise.

### GetMaxClients5g1Ok

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClients5g1Ok() (*int32, bool)`

GetMaxClients5g1Ok returns a tuple with the MaxClients5g1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClients5g1

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetMaxClients5g1(v int32)`

SetMaxClients5g1 sets MaxClients5g1 field to given value.

### HasMaxClients5g1

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasMaxClients5g1() bool`

HasMaxClients5g1 returns a boolean if a field has been set.

### GetMaxClients5g2

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClients5g2() int32`

GetMaxClients5g2 returns the MaxClients5g2 field if non-nil, zero value otherwise.

### GetMaxClients5g2Ok

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClients5g2Ok() (*int32, bool)`

GetMaxClients5g2Ok returns a tuple with the MaxClients5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClients5g2

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetMaxClients5g2(v int32)`

SetMaxClients5g2 sets MaxClients5g2 field to given value.

### HasMaxClients5g2

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasMaxClients5g2() bool`

HasMaxClients5g2 returns a boolean if a field has been set.

### GetMaxClients6g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClients6g() int32`

GetMaxClients6g returns the MaxClients6g field if non-nil, zero value otherwise.

### GetMaxClients6gOk

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClients6gOk() (*int32, bool)`

GetMaxClients6gOk returns a tuple with the MaxClients6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClients6g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetMaxClients6g(v int32)`

SetMaxClients6g sets MaxClients6g field to given value.

### HasMaxClients6g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasMaxClients6g() bool`

HasMaxClients6g returns a boolean if a field has been set.

### GetMaxClientsEnable2g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClientsEnable2g() bool`

GetMaxClientsEnable2g returns the MaxClientsEnable2g field if non-nil, zero value otherwise.

### GetMaxClientsEnable2gOk

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClientsEnable2gOk() (*bool, bool)`

GetMaxClientsEnable2gOk returns a tuple with the MaxClientsEnable2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClientsEnable2g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetMaxClientsEnable2g(v bool)`

SetMaxClientsEnable2g sets MaxClientsEnable2g field to given value.

### HasMaxClientsEnable2g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasMaxClientsEnable2g() bool`

HasMaxClientsEnable2g returns a boolean if a field has been set.

### GetMaxClientsEnable5g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClientsEnable5g() bool`

GetMaxClientsEnable5g returns the MaxClientsEnable5g field if non-nil, zero value otherwise.

### GetMaxClientsEnable5gOk

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClientsEnable5gOk() (*bool, bool)`

GetMaxClientsEnable5gOk returns a tuple with the MaxClientsEnable5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClientsEnable5g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetMaxClientsEnable5g(v bool)`

SetMaxClientsEnable5g sets MaxClientsEnable5g field to given value.

### HasMaxClientsEnable5g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasMaxClientsEnable5g() bool`

HasMaxClientsEnable5g returns a boolean if a field has been set.

### GetMaxClientsEnable5g1

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClientsEnable5g1() bool`

GetMaxClientsEnable5g1 returns the MaxClientsEnable5g1 field if non-nil, zero value otherwise.

### GetMaxClientsEnable5g1Ok

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClientsEnable5g1Ok() (*bool, bool)`

GetMaxClientsEnable5g1Ok returns a tuple with the MaxClientsEnable5g1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClientsEnable5g1

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetMaxClientsEnable5g1(v bool)`

SetMaxClientsEnable5g1 sets MaxClientsEnable5g1 field to given value.

### HasMaxClientsEnable5g1

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasMaxClientsEnable5g1() bool`

HasMaxClientsEnable5g1 returns a boolean if a field has been set.

### GetMaxClientsEnable5g2

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClientsEnable5g2() bool`

GetMaxClientsEnable5g2 returns the MaxClientsEnable5g2 field if non-nil, zero value otherwise.

### GetMaxClientsEnable5g2Ok

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClientsEnable5g2Ok() (*bool, bool)`

GetMaxClientsEnable5g2Ok returns a tuple with the MaxClientsEnable5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClientsEnable5g2

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetMaxClientsEnable5g2(v bool)`

SetMaxClientsEnable5g2 sets MaxClientsEnable5g2 field to given value.

### HasMaxClientsEnable5g2

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasMaxClientsEnable5g2() bool`

HasMaxClientsEnable5g2 returns a boolean if a field has been set.

### GetMaxClientsEnable6g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClientsEnable6g() bool`

GetMaxClientsEnable6g returns the MaxClientsEnable6g field if non-nil, zero value otherwise.

### GetMaxClientsEnable6gOk

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetMaxClientsEnable6gOk() (*bool, bool)`

GetMaxClientsEnable6gOk returns a tuple with the MaxClientsEnable6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClientsEnable6g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetMaxClientsEnable6g(v bool)`

SetMaxClientsEnable6g sets MaxClientsEnable6g field to given value.

### HasMaxClientsEnable6g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasMaxClientsEnable6g() bool`

HasMaxClientsEnable6g returns a boolean if a field has been set.

### GetRssiEnable2g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetRssiEnable2g() bool`

GetRssiEnable2g returns the RssiEnable2g field if non-nil, zero value otherwise.

### GetRssiEnable2gOk

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetRssiEnable2gOk() (*bool, bool)`

GetRssiEnable2gOk returns a tuple with the RssiEnable2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiEnable2g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetRssiEnable2g(v bool)`

SetRssiEnable2g sets RssiEnable2g field to given value.

### HasRssiEnable2g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasRssiEnable2g() bool`

HasRssiEnable2g returns a boolean if a field has been set.

### GetRssiEnable5g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetRssiEnable5g() bool`

GetRssiEnable5g returns the RssiEnable5g field if non-nil, zero value otherwise.

### GetRssiEnable5gOk

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetRssiEnable5gOk() (*bool, bool)`

GetRssiEnable5gOk returns a tuple with the RssiEnable5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiEnable5g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetRssiEnable5g(v bool)`

SetRssiEnable5g sets RssiEnable5g field to given value.

### HasRssiEnable5g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasRssiEnable5g() bool`

HasRssiEnable5g returns a boolean if a field has been set.

### GetRssiEnable5g1

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetRssiEnable5g1() bool`

GetRssiEnable5g1 returns the RssiEnable5g1 field if non-nil, zero value otherwise.

### GetRssiEnable5g1Ok

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetRssiEnable5g1Ok() (*bool, bool)`

GetRssiEnable5g1Ok returns a tuple with the RssiEnable5g1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiEnable5g1

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetRssiEnable5g1(v bool)`

SetRssiEnable5g1 sets RssiEnable5g1 field to given value.

### HasRssiEnable5g1

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasRssiEnable5g1() bool`

HasRssiEnable5g1 returns a boolean if a field has been set.

### GetRssiEnable5g2

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetRssiEnable5g2() bool`

GetRssiEnable5g2 returns the RssiEnable5g2 field if non-nil, zero value otherwise.

### GetRssiEnable5g2Ok

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetRssiEnable5g2Ok() (*bool, bool)`

GetRssiEnable5g2Ok returns a tuple with the RssiEnable5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiEnable5g2

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetRssiEnable5g2(v bool)`

SetRssiEnable5g2 sets RssiEnable5g2 field to given value.

### HasRssiEnable5g2

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasRssiEnable5g2() bool`

HasRssiEnable5g2 returns a boolean if a field has been set.

### GetRssiEnable6g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetRssiEnable6g() bool`

GetRssiEnable6g returns the RssiEnable6g field if non-nil, zero value otherwise.

### GetRssiEnable6gOk

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetRssiEnable6gOk() (*bool, bool)`

GetRssiEnable6gOk returns a tuple with the RssiEnable6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiEnable6g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetRssiEnable6g(v bool)`

SetRssiEnable6g sets RssiEnable6g field to given value.

### HasRssiEnable6g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasRssiEnable6g() bool`

HasRssiEnable6g returns a boolean if a field has been set.

### GetThreshold2g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetThreshold2g() int32`

GetThreshold2g returns the Threshold2g field if non-nil, zero value otherwise.

### GetThreshold2gOk

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetThreshold2gOk() (*int32, bool)`

GetThreshold2gOk returns a tuple with the Threshold2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold2g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetThreshold2g(v int32)`

SetThreshold2g sets Threshold2g field to given value.

### HasThreshold2g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasThreshold2g() bool`

HasThreshold2g returns a boolean if a field has been set.

### GetThreshold5g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetThreshold5g() int32`

GetThreshold5g returns the Threshold5g field if non-nil, zero value otherwise.

### GetThreshold5gOk

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetThreshold5gOk() (*int32, bool)`

GetThreshold5gOk returns a tuple with the Threshold5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold5g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetThreshold5g(v int32)`

SetThreshold5g sets Threshold5g field to given value.

### HasThreshold5g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasThreshold5g() bool`

HasThreshold5g returns a boolean if a field has been set.

### GetThreshold5g1

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetThreshold5g1() int32`

GetThreshold5g1 returns the Threshold5g1 field if non-nil, zero value otherwise.

### GetThreshold5g1Ok

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetThreshold5g1Ok() (*int32, bool)`

GetThreshold5g1Ok returns a tuple with the Threshold5g1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold5g1

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetThreshold5g1(v int32)`

SetThreshold5g1 sets Threshold5g1 field to given value.

### HasThreshold5g1

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasThreshold5g1() bool`

HasThreshold5g1 returns a boolean if a field has been set.

### GetThreshold5g2

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetThreshold5g2() int32`

GetThreshold5g2 returns the Threshold5g2 field if non-nil, zero value otherwise.

### GetThreshold5g2Ok

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetThreshold5g2Ok() (*int32, bool)`

GetThreshold5g2Ok returns a tuple with the Threshold5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold5g2

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetThreshold5g2(v int32)`

SetThreshold5g2 sets Threshold5g2 field to given value.

### HasThreshold5g2

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasThreshold5g2() bool`

HasThreshold5g2 returns a boolean if a field has been set.

### GetThreshold6g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetThreshold6g() int32`

GetThreshold6g returns the Threshold6g field if non-nil, zero value otherwise.

### GetThreshold6gOk

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) GetThreshold6gOk() (*int32, bool)`

GetThreshold6gOk returns a tuple with the Threshold6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold6g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) SetThreshold6g(v int32)`

SetThreshold6g sets Threshold6g field to given value.

### HasThreshold6g

`func (o *UpdateApAdvancedLoadBalanceOpenApiVO) HasThreshold6g() bool`

HasThreshold6g returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


