# ApBridgeTdmaConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to [**[]BridgeClientApConfigVO**](BridgeClientApConfigVO.md) | Bridge TDMA Client config. | [optional] 
**Status** | **int32** | Bridge TDMA config status. 0: disable, 1: enable. | 

## Methods

### NewApBridgeTdmaConfigVO

`func NewApBridgeTdmaConfigVO(status int32, ) *ApBridgeTdmaConfigVO`

NewApBridgeTdmaConfigVO instantiates a new ApBridgeTdmaConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApBridgeTdmaConfigVOWithDefaults

`func NewApBridgeTdmaConfigVOWithDefaults() *ApBridgeTdmaConfigVO`

NewApBridgeTdmaConfigVOWithDefaults instantiates a new ApBridgeTdmaConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *ApBridgeTdmaConfigVO) GetClients() []BridgeClientApConfigVO`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ApBridgeTdmaConfigVO) GetClientsOk() (*[]BridgeClientApConfigVO, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ApBridgeTdmaConfigVO) SetClients(v []BridgeClientApConfigVO)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ApBridgeTdmaConfigVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetStatus

`func (o *ApBridgeTdmaConfigVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApBridgeTdmaConfigVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApBridgeTdmaConfigVO) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


