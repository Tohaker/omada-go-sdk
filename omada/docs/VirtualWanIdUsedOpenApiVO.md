# VirtualWanIdUsedOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDelete** | Pointer to **bool** | Whether all Virtual WAN related functions are allowed to be deleted. | [optional] 
**AllowDeleteFunctions** | Pointer to **[]string** | The Virtual WAN related settings that are allowed to be deleted. | [optional] 
**Functions** | Pointer to **[]string** | The functions that have adopted the current Virtual WAN | [optional] 
**NotAllowDeleteFunctions** | Pointer to **[]string** | The Virtual WAN related settings that are not allowed to be deleted. | [optional] 
**VirtualWanId** | Pointer to **string** | The ID of the Virtual WAN. | [optional] 

## Methods

### NewVirtualWanIdUsedOpenApiVO

`func NewVirtualWanIdUsedOpenApiVO() *VirtualWanIdUsedOpenApiVO`

NewVirtualWanIdUsedOpenApiVO instantiates a new VirtualWanIdUsedOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanIdUsedOpenApiVOWithDefaults

`func NewVirtualWanIdUsedOpenApiVOWithDefaults() *VirtualWanIdUsedOpenApiVO`

NewVirtualWanIdUsedOpenApiVOWithDefaults instantiates a new VirtualWanIdUsedOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDelete

`func (o *VirtualWanIdUsedOpenApiVO) GetAllowDelete() bool`

GetAllowDelete returns the AllowDelete field if non-nil, zero value otherwise.

### GetAllowDeleteOk

`func (o *VirtualWanIdUsedOpenApiVO) GetAllowDeleteOk() (*bool, bool)`

GetAllowDeleteOk returns a tuple with the AllowDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDelete

`func (o *VirtualWanIdUsedOpenApiVO) SetAllowDelete(v bool)`

SetAllowDelete sets AllowDelete field to given value.

### HasAllowDelete

`func (o *VirtualWanIdUsedOpenApiVO) HasAllowDelete() bool`

HasAllowDelete returns a boolean if a field has been set.

### GetAllowDeleteFunctions

`func (o *VirtualWanIdUsedOpenApiVO) GetAllowDeleteFunctions() []string`

GetAllowDeleteFunctions returns the AllowDeleteFunctions field if non-nil, zero value otherwise.

### GetAllowDeleteFunctionsOk

`func (o *VirtualWanIdUsedOpenApiVO) GetAllowDeleteFunctionsOk() (*[]string, bool)`

GetAllowDeleteFunctionsOk returns a tuple with the AllowDeleteFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDeleteFunctions

`func (o *VirtualWanIdUsedOpenApiVO) SetAllowDeleteFunctions(v []string)`

SetAllowDeleteFunctions sets AllowDeleteFunctions field to given value.

### HasAllowDeleteFunctions

`func (o *VirtualWanIdUsedOpenApiVO) HasAllowDeleteFunctions() bool`

HasAllowDeleteFunctions returns a boolean if a field has been set.

### GetFunctions

`func (o *VirtualWanIdUsedOpenApiVO) GetFunctions() []string`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *VirtualWanIdUsedOpenApiVO) GetFunctionsOk() (*[]string, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *VirtualWanIdUsedOpenApiVO) SetFunctions(v []string)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *VirtualWanIdUsedOpenApiVO) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetNotAllowDeleteFunctions

`func (o *VirtualWanIdUsedOpenApiVO) GetNotAllowDeleteFunctions() []string`

GetNotAllowDeleteFunctions returns the NotAllowDeleteFunctions field if non-nil, zero value otherwise.

### GetNotAllowDeleteFunctionsOk

`func (o *VirtualWanIdUsedOpenApiVO) GetNotAllowDeleteFunctionsOk() (*[]string, bool)`

GetNotAllowDeleteFunctionsOk returns a tuple with the NotAllowDeleteFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllowDeleteFunctions

`func (o *VirtualWanIdUsedOpenApiVO) SetNotAllowDeleteFunctions(v []string)`

SetNotAllowDeleteFunctions sets NotAllowDeleteFunctions field to given value.

### HasNotAllowDeleteFunctions

`func (o *VirtualWanIdUsedOpenApiVO) HasNotAllowDeleteFunctions() bool`

HasNotAllowDeleteFunctions returns a boolean if a field has been set.

### GetVirtualWanId

`func (o *VirtualWanIdUsedOpenApiVO) GetVirtualWanId() string`

GetVirtualWanId returns the VirtualWanId field if non-nil, zero value otherwise.

### GetVirtualWanIdOk

`func (o *VirtualWanIdUsedOpenApiVO) GetVirtualWanIdOk() (*string, bool)`

GetVirtualWanIdOk returns a tuple with the VirtualWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanId

`func (o *VirtualWanIdUsedOpenApiVO) SetVirtualWanId(v string)`

SetVirtualWanId sets VirtualWanId field to given value.

### HasVirtualWanId

`func (o *VirtualWanIdUsedOpenApiVO) HasVirtualWanId() bool`

HasVirtualWanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


