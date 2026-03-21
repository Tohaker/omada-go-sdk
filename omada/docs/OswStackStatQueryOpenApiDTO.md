# OswStackStatQueryOpenApiDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attrs** | **[]string** | Attributes to be queried. Attributes not included in attrs will return a value of 0. Item of attrs should be a value as follows: mem, cpu, tx, rx, txRate, rxRate, txPkts, rxPkts, txBroadPkts, rxBroadPkts, txMultiPkts, rxMultiPkts, dropPkts, txErrPkts, rxErrPkts | 
**End** | **int64** | End time, number of seconds from UTC0 1970/01/01 | 
**StandardPorts** | Pointer to **[]string** | Statistics of the selected ports of Stack would be queried. | [optional] 
**Start** | **int64** | Start time, number of seconds from UTC0 1970/01/01 | 

## Methods

### NewOswStackStatQueryOpenApiDTO

`func NewOswStackStatQueryOpenApiDTO(attrs []string, end int64, start int64, ) *OswStackStatQueryOpenApiDTO`

NewOswStackStatQueryOpenApiDTO instantiates a new OswStackStatQueryOpenApiDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackStatQueryOpenApiDTOWithDefaults

`func NewOswStackStatQueryOpenApiDTOWithDefaults() *OswStackStatQueryOpenApiDTO`

NewOswStackStatQueryOpenApiDTOWithDefaults instantiates a new OswStackStatQueryOpenApiDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttrs

`func (o *OswStackStatQueryOpenApiDTO) GetAttrs() []string`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *OswStackStatQueryOpenApiDTO) GetAttrsOk() (*[]string, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *OswStackStatQueryOpenApiDTO) SetAttrs(v []string)`

SetAttrs sets Attrs field to given value.


### GetEnd

`func (o *OswStackStatQueryOpenApiDTO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *OswStackStatQueryOpenApiDTO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *OswStackStatQueryOpenApiDTO) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetStandardPorts

`func (o *OswStackStatQueryOpenApiDTO) GetStandardPorts() []string`

GetStandardPorts returns the StandardPorts field if non-nil, zero value otherwise.

### GetStandardPortsOk

`func (o *OswStackStatQueryOpenApiDTO) GetStandardPortsOk() (*[]string, bool)`

GetStandardPortsOk returns a tuple with the StandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPorts

`func (o *OswStackStatQueryOpenApiDTO) SetStandardPorts(v []string)`

SetStandardPorts sets StandardPorts field to given value.

### HasStandardPorts

`func (o *OswStackStatQueryOpenApiDTO) HasStandardPorts() bool`

HasStandardPorts returns a boolean if a field has been set.

### GetStart

`func (o *OswStackStatQueryOpenApiDTO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *OswStackStatQueryOpenApiDTO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *OswStackStatQueryOpenApiDTO) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


