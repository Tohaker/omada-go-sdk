# RetryEap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **string** | AP MAC | [optional] 
**Avg** | Pointer to **float64** | Average current AP retransmission packet rate | [optional] 
**Model** | Pointer to **string** | AP model | [optional] 
**ModelVersion** | Pointer to **string** | AP model Version | [optional] 
**Name** | Pointer to **string** | AP name | [optional] 
**Retries** | Pointer to [**[]Retry**](Retry.md) | AP retries timing list | [optional] 
**Status** | Pointer to **int32** | AP status, 0: connected, 1: disconnected | [optional] 

## Methods

### NewRetryEap

`func NewRetryEap() *RetryEap`

NewRetryEap instantiates a new RetryEap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryEapWithDefaults

`func NewRetryEapWithDefaults() *RetryEap`

NewRetryEapWithDefaults instantiates a new RetryEap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *RetryEap) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *RetryEap) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *RetryEap) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *RetryEap) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### GetAvg

`func (o *RetryEap) GetAvg() float64`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *RetryEap) GetAvgOk() (*float64, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *RetryEap) SetAvg(v float64)`

SetAvg sets Avg field to given value.

### HasAvg

`func (o *RetryEap) HasAvg() bool`

HasAvg returns a boolean if a field has been set.

### GetModel

`func (o *RetryEap) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RetryEap) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RetryEap) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *RetryEap) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *RetryEap) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *RetryEap) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *RetryEap) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *RetryEap) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *RetryEap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RetryEap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RetryEap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RetryEap) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetries

`func (o *RetryEap) GetRetries() []Retry`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *RetryEap) GetRetriesOk() (*[]Retry, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *RetryEap) SetRetries(v []Retry)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *RetryEap) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetStatus

`func (o *RetryEap) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RetryEap) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RetryEap) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RetryEap) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


