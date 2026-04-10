# PortOnlineStatusOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualWanOnlineStatus** | Pointer to [**[]VirtualWanOnlineStatusOpenApiVO**](VirtualWanOnlineStatusOpenApiVO.md) | Virtual WAN online status list | [optional] 
**WanOnlineStatus** | Pointer to [**[]WanOnlineStatusOpenApiVO**](WanOnlineStatusOpenApiVO.md) | WAN online status list | [optional] 

## Methods

### NewPortOnlineStatusOpenApiVO

`func NewPortOnlineStatusOpenApiVO() *PortOnlineStatusOpenApiVO`

NewPortOnlineStatusOpenApiVO instantiates a new PortOnlineStatusOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortOnlineStatusOpenApiVOWithDefaults

`func NewPortOnlineStatusOpenApiVOWithDefaults() *PortOnlineStatusOpenApiVO`

NewPortOnlineStatusOpenApiVOWithDefaults instantiates a new PortOnlineStatusOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualWanOnlineStatus

`func (o *PortOnlineStatusOpenApiVO) GetVirtualWanOnlineStatus() []VirtualWanOnlineStatusOpenApiVO`

GetVirtualWanOnlineStatus returns the VirtualWanOnlineStatus field if non-nil, zero value otherwise.

### GetVirtualWanOnlineStatusOk

`func (o *PortOnlineStatusOpenApiVO) GetVirtualWanOnlineStatusOk() (*[]VirtualWanOnlineStatusOpenApiVO, bool)`

GetVirtualWanOnlineStatusOk returns a tuple with the VirtualWanOnlineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanOnlineStatus

`func (o *PortOnlineStatusOpenApiVO) SetVirtualWanOnlineStatus(v []VirtualWanOnlineStatusOpenApiVO)`

SetVirtualWanOnlineStatus sets VirtualWanOnlineStatus field to given value.

### HasVirtualWanOnlineStatus

`func (o *PortOnlineStatusOpenApiVO) HasVirtualWanOnlineStatus() bool`

HasVirtualWanOnlineStatus returns a boolean if a field has been set.

### GetWanOnlineStatus

`func (o *PortOnlineStatusOpenApiVO) GetWanOnlineStatus() []WanOnlineStatusOpenApiVO`

GetWanOnlineStatus returns the WanOnlineStatus field if non-nil, zero value otherwise.

### GetWanOnlineStatusOk

`func (o *PortOnlineStatusOpenApiVO) GetWanOnlineStatusOk() (*[]WanOnlineStatusOpenApiVO, bool)`

GetWanOnlineStatusOk returns a tuple with the WanOnlineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanOnlineStatus

`func (o *PortOnlineStatusOpenApiVO) SetWanOnlineStatus(v []WanOnlineStatusOpenApiVO)`

SetWanOnlineStatus sets WanOnlineStatus field to given value.

### HasWanOnlineStatus

`func (o *PortOnlineStatusOpenApiVO) HasWanOnlineStatus() bool`

HasWanOnlineStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


