# PortVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagId** | Pointer to **int32** | (Wired) LAG ID. Exists only when the client is connected to the LAG | [optional] 
**Port** | Pointer to **int32** | The port id of the device port. | [optional] 
**StandardPort** | Pointer to **string** | The stand port id of the device port. | [optional] 

## Methods

### NewPortVO

`func NewPortVO() *PortVO`

NewPortVO instantiates a new PortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortVOWithDefaults

`func NewPortVOWithDefaults() *PortVO`

NewPortVOWithDefaults instantiates a new PortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagId

`func (o *PortVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *PortVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *PortVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *PortVO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetPort

`func (o *PortVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PortVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PortVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PortVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStandardPort

`func (o *PortVO) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *PortVO) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *PortVO) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *PortVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


