# PonPortRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PonPort** | **string** | Pon port.e.g.GPON 1/0/1 | 

## Methods

### NewPonPortRequestDTO

`func NewPonPortRequestDTO(ponPort string, ) *PonPortRequestDTO`

NewPonPortRequestDTO instantiates a new PonPortRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPonPortRequestDTOWithDefaults

`func NewPonPortRequestDTOWithDefaults() *PonPortRequestDTO`

NewPonPortRequestDTOWithDefaults instantiates a new PonPortRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPonPort

`func (o *PonPortRequestDTO) GetPonPort() string`

GetPonPort returns the PonPort field if non-nil, zero value otherwise.

### GetPonPortOk

`func (o *PonPortRequestDTO) GetPonPortOk() (*string, bool)`

GetPonPortOk returns a tuple with the PonPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPort

`func (o *PonPortRequestDTO) SetPonPort(v string)`

SetPonPort sets PonPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


