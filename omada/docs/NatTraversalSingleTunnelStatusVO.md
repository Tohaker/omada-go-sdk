# NatTraversalSingleTunnelStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the remote access tunnel. | [optional] 
**OpenStatus** | Pointer to **bool** | The open and closed status of the remote access tunnel. | [optional] 
**Status** | Pointer to **int32** | Status of the remote access tunnel. 0: Disconnected, 1: Connected, 2: Opening, -1: Heartbeat Missed, -2: Expired. | [optional] 
**TunnelId** | Pointer to **int32** |  | [optional] 

## Methods

### NewNatTraversalSingleTunnelStatusVO

`func NewNatTraversalSingleTunnelStatusVO() *NatTraversalSingleTunnelStatusVO`

NewNatTraversalSingleTunnelStatusVO instantiates a new NatTraversalSingleTunnelStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatTraversalSingleTunnelStatusVOWithDefaults

`func NewNatTraversalSingleTunnelStatusVOWithDefaults() *NatTraversalSingleTunnelStatusVO`

NewNatTraversalSingleTunnelStatusVOWithDefaults instantiates a new NatTraversalSingleTunnelStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NatTraversalSingleTunnelStatusVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NatTraversalSingleTunnelStatusVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NatTraversalSingleTunnelStatusVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NatTraversalSingleTunnelStatusVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpenStatus

`func (o *NatTraversalSingleTunnelStatusVO) GetOpenStatus() bool`

GetOpenStatus returns the OpenStatus field if non-nil, zero value otherwise.

### GetOpenStatusOk

`func (o *NatTraversalSingleTunnelStatusVO) GetOpenStatusOk() (*bool, bool)`

GetOpenStatusOk returns a tuple with the OpenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStatus

`func (o *NatTraversalSingleTunnelStatusVO) SetOpenStatus(v bool)`

SetOpenStatus sets OpenStatus field to given value.

### HasOpenStatus

`func (o *NatTraversalSingleTunnelStatusVO) HasOpenStatus() bool`

HasOpenStatus returns a boolean if a field has been set.

### GetStatus

`func (o *NatTraversalSingleTunnelStatusVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NatTraversalSingleTunnelStatusVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NatTraversalSingleTunnelStatusVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NatTraversalSingleTunnelStatusVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTunnelId

`func (o *NatTraversalSingleTunnelStatusVO) GetTunnelId() int32`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *NatTraversalSingleTunnelStatusVO) GetTunnelIdOk() (*int32, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *NatTraversalSingleTunnelStatusVO) SetTunnelId(v int32)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *NatTraversalSingleTunnelStatusVO) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


