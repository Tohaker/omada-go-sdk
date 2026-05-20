# ClientMultiLinkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RadioId** | Pointer to **int32** | (Wireless) Radio ID should be a value as follows: 0: 2.4GHz; 1: 5GHz; 2:5GHz-2; 3: 6GHz | [optional] 
**Signal** | Pointer to **int32** | (Wireless) Signal strength, unit: dBm. | [optional] 

## Methods

### NewClientMultiLinkInfo

`func NewClientMultiLinkInfo() *ClientMultiLinkInfo`

NewClientMultiLinkInfo instantiates a new ClientMultiLinkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientMultiLinkInfoWithDefaults

`func NewClientMultiLinkInfoWithDefaults() *ClientMultiLinkInfo`

NewClientMultiLinkInfoWithDefaults instantiates a new ClientMultiLinkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRadioId

`func (o *ClientMultiLinkInfo) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *ClientMultiLinkInfo) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *ClientMultiLinkInfo) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *ClientMultiLinkInfo) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.

### GetSignal

`func (o *ClientMultiLinkInfo) GetSignal() int32`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *ClientMultiLinkInfo) GetSignalOk() (*int32, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *ClientMultiLinkInfo) SetSignal(v int32)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *ClientMultiLinkInfo) HasSignal() bool`

HasSignal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


