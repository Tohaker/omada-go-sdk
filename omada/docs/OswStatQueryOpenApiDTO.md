# OswStatQueryOpenApiDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attrs** | **[]string** | Attributes to be queried. Attributes not included in attrs will return a value of 0. Item of attrs should be a value as follows: mem, cpu, tx, rx, txRate, rxRate, txPkts, rxPkts, txBroadPkts, rxBroadPkts, txMultiPkts, rxMultiPkts, dropPkts, txErrPkts, rxErrPkts | 
**End** | **int64** | End time, number of seconds from UTC0 1970/01/01 | 
**Ports** | Pointer to **[]int32** | Statistics of the selected ports of Switch would be queried. | [optional] 
**Start** | **int64** | Start time, number of seconds from UTC0 1970/01/01 | 

## Methods

### NewOswStatQueryOpenApiDTO

`func NewOswStatQueryOpenApiDTO(attrs []string, end int64, start int64, ) *OswStatQueryOpenApiDTO`

NewOswStatQueryOpenApiDTO instantiates a new OswStatQueryOpenApiDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStatQueryOpenApiDTOWithDefaults

`func NewOswStatQueryOpenApiDTOWithDefaults() *OswStatQueryOpenApiDTO`

NewOswStatQueryOpenApiDTOWithDefaults instantiates a new OswStatQueryOpenApiDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttrs

`func (o *OswStatQueryOpenApiDTO) GetAttrs() []string`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *OswStatQueryOpenApiDTO) GetAttrsOk() (*[]string, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *OswStatQueryOpenApiDTO) SetAttrs(v []string)`

SetAttrs sets Attrs field to given value.


### GetEnd

`func (o *OswStatQueryOpenApiDTO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *OswStatQueryOpenApiDTO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *OswStatQueryOpenApiDTO) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetPorts

`func (o *OswStatQueryOpenApiDTO) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswStatQueryOpenApiDTO) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswStatQueryOpenApiDTO) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswStatQueryOpenApiDTO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetStart

`func (o *OswStatQueryOpenApiDTO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *OswStatQueryOpenApiDTO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *OswStatQueryOpenApiDTO) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


