# APBridgeParingWindowResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAps** | Pointer to [**[]ApBridgeClientApOpenApiVO**](ApBridgeClientApOpenApiVO.md) |  | [optional] 
**Countdown** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewAPBridgeParingWindowResult

`func NewAPBridgeParingWindowResult() *APBridgeParingWindowResult`

NewAPBridgeParingWindowResult instantiates a new APBridgeParingWindowResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPBridgeParingWindowResultWithDefaults

`func NewAPBridgeParingWindowResultWithDefaults() *APBridgeParingWindowResult`

NewAPBridgeParingWindowResultWithDefaults instantiates a new APBridgeParingWindowResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientAps

`func (o *APBridgeParingWindowResult) GetClientAps() []ApBridgeClientApOpenApiVO`

GetClientAps returns the ClientAps field if non-nil, zero value otherwise.

### GetClientApsOk

`func (o *APBridgeParingWindowResult) GetClientApsOk() (*[]ApBridgeClientApOpenApiVO, bool)`

GetClientApsOk returns a tuple with the ClientAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAps

`func (o *APBridgeParingWindowResult) SetClientAps(v []ApBridgeClientApOpenApiVO)`

SetClientAps sets ClientAps field to given value.

### HasClientAps

`func (o *APBridgeParingWindowResult) HasClientAps() bool`

HasClientAps returns a boolean if a field has been set.

### GetCountdown

`func (o *APBridgeParingWindowResult) GetCountdown() int64`

GetCountdown returns the Countdown field if non-nil, zero value otherwise.

### GetCountdownOk

`func (o *APBridgeParingWindowResult) GetCountdownOk() (*int64, bool)`

GetCountdownOk returns a tuple with the Countdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountdown

`func (o *APBridgeParingWindowResult) SetCountdown(v int64)`

SetCountdown sets Countdown field to given value.

### HasCountdown

`func (o *APBridgeParingWindowResult) HasCountdown() bool`

HasCountdown returns a boolean if a field has been set.

### GetStatus

`func (o *APBridgeParingWindowResult) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APBridgeParingWindowResult) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APBridgeParingWindowResult) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *APBridgeParingWindowResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


