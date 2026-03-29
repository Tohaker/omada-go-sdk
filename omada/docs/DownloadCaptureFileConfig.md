# DownloadCaptureFileConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | A GUID based on the timestamp. | 
**Stack** | Pointer to **bool** | Parameter [stack] indicates whether the device supports stacking. | [optional] 

## Methods

### NewDownloadCaptureFileConfig

`func NewDownloadCaptureFileConfig(requestId string, ) *DownloadCaptureFileConfig`

NewDownloadCaptureFileConfig instantiates a new DownloadCaptureFileConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadCaptureFileConfigWithDefaults

`func NewDownloadCaptureFileConfigWithDefaults() *DownloadCaptureFileConfig`

NewDownloadCaptureFileConfigWithDefaults instantiates a new DownloadCaptureFileConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *DownloadCaptureFileConfig) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DownloadCaptureFileConfig) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DownloadCaptureFileConfig) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetStack

`func (o *DownloadCaptureFileConfig) GetStack() bool`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *DownloadCaptureFileConfig) GetStackOk() (*bool, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *DownloadCaptureFileConfig) SetStack(v bool)`

SetStack sets Stack field to given value.

### HasStack

`func (o *DownloadCaptureFileConfig) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


