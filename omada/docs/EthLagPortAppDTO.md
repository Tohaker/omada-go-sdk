# EthLagPortAppDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of port | [optional] 
**Duplex** | Pointer to **int32** | Port duplex negotiation mode.Duplex should be a value as follows: 2:FULL,0:AUTO | [optional] 
**FlowControl** | Pointer to **int32** | Port flow control function switch.FlowControl should be a value as follows: 0:DISABLE;1:ENABLE | [optional] 
**Lag** | **string** | ID of LAG | 
**Speed** | Pointer to **int32** | Port speed negotiation mode. speed should be a value as follows: 0;10;100;1000;2500;10000.0 represents Auto, and all other values are in Mbps. | [optional] 
**Status** | Pointer to **int32** | Port switch status.Status should be a value as follows:0:DISABLE;1:ENABLE | [optional] 

## Methods

### NewEthLagPortAppDTO

`func NewEthLagPortAppDTO(lag string, ) *EthLagPortAppDTO`

NewEthLagPortAppDTO instantiates a new EthLagPortAppDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthLagPortAppDTOWithDefaults

`func NewEthLagPortAppDTOWithDefaults() *EthLagPortAppDTO`

NewEthLagPortAppDTOWithDefaults instantiates a new EthLagPortAppDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EthLagPortAppDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EthLagPortAppDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EthLagPortAppDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EthLagPortAppDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuplex

`func (o *EthLagPortAppDTO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *EthLagPortAppDTO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *EthLagPortAppDTO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *EthLagPortAppDTO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetFlowControl

`func (o *EthLagPortAppDTO) GetFlowControl() int32`

GetFlowControl returns the FlowControl field if non-nil, zero value otherwise.

### GetFlowControlOk

`func (o *EthLagPortAppDTO) GetFlowControlOk() (*int32, bool)`

GetFlowControlOk returns a tuple with the FlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControl

`func (o *EthLagPortAppDTO) SetFlowControl(v int32)`

SetFlowControl sets FlowControl field to given value.

### HasFlowControl

`func (o *EthLagPortAppDTO) HasFlowControl() bool`

HasFlowControl returns a boolean if a field has been set.

### GetLag

`func (o *EthLagPortAppDTO) GetLag() string`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *EthLagPortAppDTO) GetLagOk() (*string, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *EthLagPortAppDTO) SetLag(v string)`

SetLag sets Lag field to given value.


### GetSpeed

`func (o *EthLagPortAppDTO) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *EthLagPortAppDTO) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *EthLagPortAppDTO) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *EthLagPortAppDTO) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *EthLagPortAppDTO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EthLagPortAppDTO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EthLagPortAppDTO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EthLagPortAppDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


