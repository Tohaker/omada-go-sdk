# OswPoeRecoverOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | **[]int32** | Please select the ports to perform PoE recovery. For example, [1, 2, 3]. | 

## Methods

### NewOswPoeRecoverOpenApiVO

`func NewOswPoeRecoverOpenApiVO(ports []int32, ) *OswPoeRecoverOpenApiVO`

NewOswPoeRecoverOpenApiVO instantiates a new OswPoeRecoverOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPoeRecoverOpenApiVOWithDefaults

`func NewOswPoeRecoverOpenApiVOWithDefaults() *OswPoeRecoverOpenApiVO`

NewOswPoeRecoverOpenApiVOWithDefaults instantiates a new OswPoeRecoverOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *OswPoeRecoverOpenApiVO) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswPoeRecoverOpenApiVO) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswPoeRecoverOpenApiVO) SetPorts(v []int32)`

SetPorts sets Ports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


