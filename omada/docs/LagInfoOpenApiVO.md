# LagInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lag** | Pointer to **int32** | Lag number. | [optional] 
**Name** | Pointer to **string** | Lag name. | [optional] 
**Ports** | Pointer to **[]int32** | Ports in lag | [optional] 

## Methods

### NewLagInfoOpenApiVO

`func NewLagInfoOpenApiVO() *LagInfoOpenApiVO`

NewLagInfoOpenApiVO instantiates a new LagInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLagInfoOpenApiVOWithDefaults

`func NewLagInfoOpenApiVOWithDefaults() *LagInfoOpenApiVO`

NewLagInfoOpenApiVOWithDefaults instantiates a new LagInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLag

`func (o *LagInfoOpenApiVO) GetLag() int32`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *LagInfoOpenApiVO) GetLagOk() (*int32, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *LagInfoOpenApiVO) SetLag(v int32)`

SetLag sets Lag field to given value.

### HasLag

`func (o *LagInfoOpenApiVO) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetName

`func (o *LagInfoOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LagInfoOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LagInfoOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LagInfoOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *LagInfoOpenApiVO) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *LagInfoOpenApiVO) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *LagInfoOpenApiVO) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *LagInfoOpenApiVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


