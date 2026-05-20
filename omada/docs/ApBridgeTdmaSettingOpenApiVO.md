# ApBridgeTdmaSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to [**[]ApBridgeTdmaClientApOpenApiVO**](ApBridgeTdmaClientApOpenApiVO.md) | Bridge TDMA Client AP config. | [optional] 
**NotSupportTdmaClients** | Pointer to [**[]ApBridgeNotSupportTdmaClientApOpenApiVO**](ApBridgeNotSupportTdmaClientApOpenApiVO.md) | Bridge Not Support TDMA Client AP info. | [optional] 
**Status** | Pointer to **int32** | Bridge TDMA Switch config status. 0: disable, 1: enable. | [optional] 

## Methods

### NewApBridgeTdmaSettingOpenApiVO

`func NewApBridgeTdmaSettingOpenApiVO() *ApBridgeTdmaSettingOpenApiVO`

NewApBridgeTdmaSettingOpenApiVO instantiates a new ApBridgeTdmaSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApBridgeTdmaSettingOpenApiVOWithDefaults

`func NewApBridgeTdmaSettingOpenApiVOWithDefaults() *ApBridgeTdmaSettingOpenApiVO`

NewApBridgeTdmaSettingOpenApiVOWithDefaults instantiates a new ApBridgeTdmaSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *ApBridgeTdmaSettingOpenApiVO) GetClients() []ApBridgeTdmaClientApOpenApiVO`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ApBridgeTdmaSettingOpenApiVO) GetClientsOk() (*[]ApBridgeTdmaClientApOpenApiVO, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ApBridgeTdmaSettingOpenApiVO) SetClients(v []ApBridgeTdmaClientApOpenApiVO)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ApBridgeTdmaSettingOpenApiVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetNotSupportTdmaClients

`func (o *ApBridgeTdmaSettingOpenApiVO) GetNotSupportTdmaClients() []ApBridgeNotSupportTdmaClientApOpenApiVO`

GetNotSupportTdmaClients returns the NotSupportTdmaClients field if non-nil, zero value otherwise.

### GetNotSupportTdmaClientsOk

`func (o *ApBridgeTdmaSettingOpenApiVO) GetNotSupportTdmaClientsOk() (*[]ApBridgeNotSupportTdmaClientApOpenApiVO, bool)`

GetNotSupportTdmaClientsOk returns a tuple with the NotSupportTdmaClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotSupportTdmaClients

`func (o *ApBridgeTdmaSettingOpenApiVO) SetNotSupportTdmaClients(v []ApBridgeNotSupportTdmaClientApOpenApiVO)`

SetNotSupportTdmaClients sets NotSupportTdmaClients field to given value.

### HasNotSupportTdmaClients

`func (o *ApBridgeTdmaSettingOpenApiVO) HasNotSupportTdmaClients() bool`

HasNotSupportTdmaClients returns a boolean if a field has been set.

### GetStatus

`func (o *ApBridgeTdmaSettingOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApBridgeTdmaSettingOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApBridgeTdmaSettingOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApBridgeTdmaSettingOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


