# ReactivePonPortDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PonPorts** | Pointer to **[]string** | Port list to be activated | [optional] 
**ReactiveStatus** | Pointer to **string** | Whether there are ports that need to be reactivated.ReactiveStatus should be a value as follows:DISABLE:Indicates that there are no ports that need to be reactivated;ENABLE:Reading the port data indicates which ports need to be activated. | [optional] 

## Methods

### NewReactivePonPortDTO

`func NewReactivePonPortDTO() *ReactivePonPortDTO`

NewReactivePonPortDTO instantiates a new ReactivePonPortDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactivePonPortDTOWithDefaults

`func NewReactivePonPortDTOWithDefaults() *ReactivePonPortDTO`

NewReactivePonPortDTOWithDefaults instantiates a new ReactivePonPortDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPonPorts

`func (o *ReactivePonPortDTO) GetPonPorts() []string`

GetPonPorts returns the PonPorts field if non-nil, zero value otherwise.

### GetPonPortsOk

`func (o *ReactivePonPortDTO) GetPonPortsOk() (*[]string, bool)`

GetPonPortsOk returns a tuple with the PonPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPorts

`func (o *ReactivePonPortDTO) SetPonPorts(v []string)`

SetPonPorts sets PonPorts field to given value.

### HasPonPorts

`func (o *ReactivePonPortDTO) HasPonPorts() bool`

HasPonPorts returns a boolean if a field has been set.

### GetReactiveStatus

`func (o *ReactivePonPortDTO) GetReactiveStatus() string`

GetReactiveStatus returns the ReactiveStatus field if non-nil, zero value otherwise.

### GetReactiveStatusOk

`func (o *ReactivePonPortDTO) GetReactiveStatusOk() (*string, bool)`

GetReactiveStatusOk returns a tuple with the ReactiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactiveStatus

`func (o *ReactivePonPortDTO) SetReactiveStatus(v string)`

SetReactiveStatus sets ReactiveStatus field to given value.

### HasReactiveStatus

`func (o *ReactivePonPortDTO) HasReactiveStatus() bool`

HasReactiveStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


