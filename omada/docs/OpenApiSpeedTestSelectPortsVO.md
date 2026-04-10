# OpenApiSpeedTestSelectPortsVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortUuidList** | Pointer to **[]string** | The uuid of selected wan ports for speed test. | [optional] 
**VirtualWanIdList** | Pointer to **[]string** | The id of selected virtual wan ports for speed test. | [optional] 

## Methods

### NewOpenApiSpeedTestSelectPortsVO

`func NewOpenApiSpeedTestSelectPortsVO() *OpenApiSpeedTestSelectPortsVO`

NewOpenApiSpeedTestSelectPortsVO instantiates a new OpenApiSpeedTestSelectPortsVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiSpeedTestSelectPortsVOWithDefaults

`func NewOpenApiSpeedTestSelectPortsVOWithDefaults() *OpenApiSpeedTestSelectPortsVO`

NewOpenApiSpeedTestSelectPortsVOWithDefaults instantiates a new OpenApiSpeedTestSelectPortsVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortUuidList

`func (o *OpenApiSpeedTestSelectPortsVO) GetPortUuidList() []string`

GetPortUuidList returns the PortUuidList field if non-nil, zero value otherwise.

### GetPortUuidListOk

`func (o *OpenApiSpeedTestSelectPortsVO) GetPortUuidListOk() (*[]string, bool)`

GetPortUuidListOk returns a tuple with the PortUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuidList

`func (o *OpenApiSpeedTestSelectPortsVO) SetPortUuidList(v []string)`

SetPortUuidList sets PortUuidList field to given value.

### HasPortUuidList

`func (o *OpenApiSpeedTestSelectPortsVO) HasPortUuidList() bool`

HasPortUuidList returns a boolean if a field has been set.

### GetVirtualWanIdList

`func (o *OpenApiSpeedTestSelectPortsVO) GetVirtualWanIdList() []string`

GetVirtualWanIdList returns the VirtualWanIdList field if non-nil, zero value otherwise.

### GetVirtualWanIdListOk

`func (o *OpenApiSpeedTestSelectPortsVO) GetVirtualWanIdListOk() (*[]string, bool)`

GetVirtualWanIdListOk returns a tuple with the VirtualWanIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanIdList

`func (o *OpenApiSpeedTestSelectPortsVO) SetVirtualWanIdList(v []string)`

SetVirtualWanIdList sets VirtualWanIdList field to given value.

### HasVirtualWanIdList

`func (o *OpenApiSpeedTestSelectPortsVO) HasVirtualWanIdList() bool`

HasVirtualWanIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


