# UnbindDeviceSnOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sns** | **[]string** | Unbind device serial number list, serial number should contains 13 characters. | 

## Methods

### NewUnbindDeviceSnOpenApiVO

`func NewUnbindDeviceSnOpenApiVO(sns []string, ) *UnbindDeviceSnOpenApiVO`

NewUnbindDeviceSnOpenApiVO instantiates a new UnbindDeviceSnOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnbindDeviceSnOpenApiVOWithDefaults

`func NewUnbindDeviceSnOpenApiVOWithDefaults() *UnbindDeviceSnOpenApiVO`

NewUnbindDeviceSnOpenApiVOWithDefaults instantiates a new UnbindDeviceSnOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSns

`func (o *UnbindDeviceSnOpenApiVO) GetSns() []string`

GetSns returns the Sns field if non-nil, zero value otherwise.

### GetSnsOk

`func (o *UnbindDeviceSnOpenApiVO) GetSnsOk() (*[]string, bool)`

GetSnsOk returns a tuple with the Sns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSns

`func (o *UnbindDeviceSnOpenApiVO) SetSns(v []string)`

SetSns sets Sns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


