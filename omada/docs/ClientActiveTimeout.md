# ClientActiveTimeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientActiveTimeoutMin** | **int32** | Client Active Timeout should be within the range of 3-10 minutes | 

## Methods

### NewClientActiveTimeout

`func NewClientActiveTimeout(clientActiveTimeoutMin int32, ) *ClientActiveTimeout`

NewClientActiveTimeout instantiates a new ClientActiveTimeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientActiveTimeoutWithDefaults

`func NewClientActiveTimeoutWithDefaults() *ClientActiveTimeout`

NewClientActiveTimeoutWithDefaults instantiates a new ClientActiveTimeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientActiveTimeoutMin

`func (o *ClientActiveTimeout) GetClientActiveTimeoutMin() int32`

GetClientActiveTimeoutMin returns the ClientActiveTimeoutMin field if non-nil, zero value otherwise.

### GetClientActiveTimeoutMinOk

`func (o *ClientActiveTimeout) GetClientActiveTimeoutMinOk() (*int32, bool)`

GetClientActiveTimeoutMinOk returns a tuple with the ClientActiveTimeoutMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientActiveTimeoutMin

`func (o *ClientActiveTimeout) SetClientActiveTimeoutMin(v int32)`

SetClientActiveTimeoutMin sets ClientActiveTimeoutMin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


