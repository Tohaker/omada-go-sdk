# OltStatQueryOpenApiDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attrs** | **[]string** | Attributes to be queried. Attributes not included in attrs will return a value of 0.Item of attrs should be a value as follows: mem, cpu, up, down, onuCount, tx, rx, txRate, rxRate, txPackets, rxPackets, txBroadcastPackets, rxBroadcastPackets, txMulticastPackets, rxMulticastPackets | 
**End** | **int64** | End time, number of seconds from UTC0 1970/01/01 | 
**OltPorts** | Pointer to **[]string** | Statistics of the selected ports of Olt would be queried. | [optional] 
**Start** | **int64** | Start time, number of seconds from UTC0 1970/01/01 | 

## Methods

### NewOltStatQueryOpenApiDTO

`func NewOltStatQueryOpenApiDTO(attrs []string, end int64, start int64, ) *OltStatQueryOpenApiDTO`

NewOltStatQueryOpenApiDTO instantiates a new OltStatQueryOpenApiDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOltStatQueryOpenApiDTOWithDefaults

`func NewOltStatQueryOpenApiDTOWithDefaults() *OltStatQueryOpenApiDTO`

NewOltStatQueryOpenApiDTOWithDefaults instantiates a new OltStatQueryOpenApiDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttrs

`func (o *OltStatQueryOpenApiDTO) GetAttrs() []string`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *OltStatQueryOpenApiDTO) GetAttrsOk() (*[]string, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *OltStatQueryOpenApiDTO) SetAttrs(v []string)`

SetAttrs sets Attrs field to given value.


### GetEnd

`func (o *OltStatQueryOpenApiDTO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *OltStatQueryOpenApiDTO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *OltStatQueryOpenApiDTO) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetOltPorts

`func (o *OltStatQueryOpenApiDTO) GetOltPorts() []string`

GetOltPorts returns the OltPorts field if non-nil, zero value otherwise.

### GetOltPortsOk

`func (o *OltStatQueryOpenApiDTO) GetOltPortsOk() (*[]string, bool)`

GetOltPortsOk returns a tuple with the OltPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOltPorts

`func (o *OltStatQueryOpenApiDTO) SetOltPorts(v []string)`

SetOltPorts sets OltPorts field to given value.

### HasOltPorts

`func (o *OltStatQueryOpenApiDTO) HasOltPorts() bool`

HasOltPorts returns a boolean if a field has been set.

### GetStart

`func (o *OltStatQueryOpenApiDTO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *OltStatQueryOpenApiDTO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *OltStatQueryOpenApiDTO) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


