# CaptureResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | The Status of package capturing. 0: Ongoing, 1: Some error occurs, 2: Local mode capture done, which can download capture file, 3: Default status, which can not download capture file. | [optional] 

## Methods

### NewCaptureResult

`func NewCaptureResult() *CaptureResult`

NewCaptureResult instantiates a new CaptureResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureResultWithDefaults

`func NewCaptureResultWithDefaults() *CaptureResult`

NewCaptureResultWithDefaults instantiates a new CaptureResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CaptureResult) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaptureResult) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaptureResult) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CaptureResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


