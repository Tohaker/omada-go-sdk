# PortMaskOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mask** | **string** | Port mask should be 4 hex number(0-9, A-F), e.g. 0000 or FFFF | 
**Port** | **int32** | Port should be within the range of 0-65535 | 

## Methods

### NewPortMaskOpenApiVO

`func NewPortMaskOpenApiVO(mask string, port int32, ) *PortMaskOpenApiVO`

NewPortMaskOpenApiVO instantiates a new PortMaskOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortMaskOpenApiVOWithDefaults

`func NewPortMaskOpenApiVOWithDefaults() *PortMaskOpenApiVO`

NewPortMaskOpenApiVOWithDefaults instantiates a new PortMaskOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMask

`func (o *PortMaskOpenApiVO) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *PortMaskOpenApiVO) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *PortMaskOpenApiVO) SetMask(v string)`

SetMask sets Mask field to given value.


### GetPort

`func (o *PortMaskOpenApiVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PortMaskOpenApiVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PortMaskOpenApiVO) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


