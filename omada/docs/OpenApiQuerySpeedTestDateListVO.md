# OpenApiQuerySpeedTestDateListVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **int32** | Start from 1. | 
**CurrentPageSize** | **int32** | It should be within the range of 1–100. | 
**PortUuid** | Pointer to **string** | Port Uuid. Only one of the parameters portUuid and virtualWanId can exist | [optional] 
**VirtualWanId** | Pointer to **string** | Virtual Wan Id. Only one of the parameters portUuid and virtualWanId can exist | [optional] 

## Methods

### NewOpenApiQuerySpeedTestDateListVO

`func NewOpenApiQuerySpeedTestDateListVO(currentPage int32, currentPageSize int32, ) *OpenApiQuerySpeedTestDateListVO`

NewOpenApiQuerySpeedTestDateListVO instantiates a new OpenApiQuerySpeedTestDateListVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiQuerySpeedTestDateListVOWithDefaults

`func NewOpenApiQuerySpeedTestDateListVOWithDefaults() *OpenApiQuerySpeedTestDateListVO`

NewOpenApiQuerySpeedTestDateListVOWithDefaults instantiates a new OpenApiQuerySpeedTestDateListVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *OpenApiQuerySpeedTestDateListVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *OpenApiQuerySpeedTestDateListVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *OpenApiQuerySpeedTestDateListVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetCurrentPageSize

`func (o *OpenApiQuerySpeedTestDateListVO) GetCurrentPageSize() int32`

GetCurrentPageSize returns the CurrentPageSize field if non-nil, zero value otherwise.

### GetCurrentPageSizeOk

`func (o *OpenApiQuerySpeedTestDateListVO) GetCurrentPageSizeOk() (*int32, bool)`

GetCurrentPageSizeOk returns a tuple with the CurrentPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageSize

`func (o *OpenApiQuerySpeedTestDateListVO) SetCurrentPageSize(v int32)`

SetCurrentPageSize sets CurrentPageSize field to given value.


### GetPortUuid

`func (o *OpenApiQuerySpeedTestDateListVO) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *OpenApiQuerySpeedTestDateListVO) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *OpenApiQuerySpeedTestDateListVO) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.

### HasPortUuid

`func (o *OpenApiQuerySpeedTestDateListVO) HasPortUuid() bool`

HasPortUuid returns a boolean if a field has been set.

### GetVirtualWanId

`func (o *OpenApiQuerySpeedTestDateListVO) GetVirtualWanId() string`

GetVirtualWanId returns the VirtualWanId field if non-nil, zero value otherwise.

### GetVirtualWanIdOk

`func (o *OpenApiQuerySpeedTestDateListVO) GetVirtualWanIdOk() (*string, bool)`

GetVirtualWanIdOk returns a tuple with the VirtualWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanId

`func (o *OpenApiQuerySpeedTestDateListVO) SetVirtualWanId(v string)`

SetVirtualWanId sets VirtualWanId field to given value.

### HasVirtualWanId

`func (o *OpenApiQuerySpeedTestDateListVO) HasVirtualWanId() bool`

HasVirtualWanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


