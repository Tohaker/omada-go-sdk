# LanPortMappingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DstPort** | Pointer to **string** |  | [optional] 
**SrcPorts** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLanPortMappingVO

`func NewLanPortMappingVO() *LanPortMappingVO`

NewLanPortMappingVO instantiates a new LanPortMappingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanPortMappingVOWithDefaults

`func NewLanPortMappingVOWithDefaults() *LanPortMappingVO`

NewLanPortMappingVOWithDefaults instantiates a new LanPortMappingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDstPort

`func (o *LanPortMappingVO) GetDstPort() string`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *LanPortMappingVO) GetDstPortOk() (*string, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *LanPortMappingVO) SetDstPort(v string)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *LanPortMappingVO) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### GetSrcPorts

`func (o *LanPortMappingVO) GetSrcPorts() []string`

GetSrcPorts returns the SrcPorts field if non-nil, zero value otherwise.

### GetSrcPortsOk

`func (o *LanPortMappingVO) GetSrcPortsOk() (*[]string, bool)`

GetSrcPortsOk returns a tuple with the SrcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPorts

`func (o *LanPortMappingVO) SetSrcPorts(v []string)`

SetSrcPorts sets SrcPorts field to given value.

### HasSrcPorts

`func (o *LanPortMappingVO) HasSrcPorts() bool`

HasSrcPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


