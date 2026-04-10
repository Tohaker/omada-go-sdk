# IgmpOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** |  | 
**Version** | **int32** | Version should be one of the following values: 2:v2; 3:v3. | 
**VirtualWanId** | Pointer to **string** | Virtual WAN ID, can be obtained from &#39;Query virtual WAN list&#39; interface. At least one of the WAN Port IDs or Virtual WAN Port IDs should not be null. | [optional] 
**WanPortId** | Pointer to **string** | WAN port ID, can be obtained from &#39;Get internet basic info&#39; interface. At least one of the WAN Port IDs or Virtual WAN Port IDs should not be null. | [optional] 

## Methods

### NewIgmpOpenApiVO

`func NewIgmpOpenApiVO(enable bool, version int32, ) *IgmpOpenApiVO`

NewIgmpOpenApiVO instantiates a new IgmpOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIgmpOpenApiVOWithDefaults

`func NewIgmpOpenApiVOWithDefaults() *IgmpOpenApiVO`

NewIgmpOpenApiVOWithDefaults instantiates a new IgmpOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *IgmpOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IgmpOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IgmpOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetVersion

`func (o *IgmpOpenApiVO) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IgmpOpenApiVO) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IgmpOpenApiVO) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVirtualWanId

`func (o *IgmpOpenApiVO) GetVirtualWanId() string`

GetVirtualWanId returns the VirtualWanId field if non-nil, zero value otherwise.

### GetVirtualWanIdOk

`func (o *IgmpOpenApiVO) GetVirtualWanIdOk() (*string, bool)`

GetVirtualWanIdOk returns a tuple with the VirtualWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanId

`func (o *IgmpOpenApiVO) SetVirtualWanId(v string)`

SetVirtualWanId sets VirtualWanId field to given value.

### HasVirtualWanId

`func (o *IgmpOpenApiVO) HasVirtualWanId() bool`

HasVirtualWanId returns a boolean if a field has been set.

### GetWanPortId

`func (o *IgmpOpenApiVO) GetWanPortId() string`

GetWanPortId returns the WanPortId field if non-nil, zero value otherwise.

### GetWanPortIdOk

`func (o *IgmpOpenApiVO) GetWanPortIdOk() (*string, bool)`

GetWanPortIdOk returns a tuple with the WanPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortId

`func (o *IgmpOpenApiVO) SetWanPortId(v string)`

SetWanPortId sets WanPortId field to given value.

### HasWanPortId

`func (o *IgmpOpenApiVO) HasWanPortId() bool`

HasWanPortId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


