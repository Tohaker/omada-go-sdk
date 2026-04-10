# VirtualWanOnlineStatusOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnlineDetection** | Pointer to **int32** | Port Online Detection | [optional] 
**VirtualWanId** | Pointer to **string** | This field represents VirtualWan ID. VirtualWan ID can be obtained from &#39;Query available virtual WAN list&#39; interface | [optional] 

## Methods

### NewVirtualWanOnlineStatusOpenApiVO

`func NewVirtualWanOnlineStatusOpenApiVO() *VirtualWanOnlineStatusOpenApiVO`

NewVirtualWanOnlineStatusOpenApiVO instantiates a new VirtualWanOnlineStatusOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanOnlineStatusOpenApiVOWithDefaults

`func NewVirtualWanOnlineStatusOpenApiVOWithDefaults() *VirtualWanOnlineStatusOpenApiVO`

NewVirtualWanOnlineStatusOpenApiVOWithDefaults instantiates a new VirtualWanOnlineStatusOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnlineDetection

`func (o *VirtualWanOnlineStatusOpenApiVO) GetOnlineDetection() int32`

GetOnlineDetection returns the OnlineDetection field if non-nil, zero value otherwise.

### GetOnlineDetectionOk

`func (o *VirtualWanOnlineStatusOpenApiVO) GetOnlineDetectionOk() (*int32, bool)`

GetOnlineDetectionOk returns a tuple with the OnlineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineDetection

`func (o *VirtualWanOnlineStatusOpenApiVO) SetOnlineDetection(v int32)`

SetOnlineDetection sets OnlineDetection field to given value.

### HasOnlineDetection

`func (o *VirtualWanOnlineStatusOpenApiVO) HasOnlineDetection() bool`

HasOnlineDetection returns a boolean if a field has been set.

### GetVirtualWanId

`func (o *VirtualWanOnlineStatusOpenApiVO) GetVirtualWanId() string`

GetVirtualWanId returns the VirtualWanId field if non-nil, zero value otherwise.

### GetVirtualWanIdOk

`func (o *VirtualWanOnlineStatusOpenApiVO) GetVirtualWanIdOk() (*string, bool)`

GetVirtualWanIdOk returns a tuple with the VirtualWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanId

`func (o *VirtualWanOnlineStatusOpenApiVO) SetVirtualWanId(v string)`

SetVirtualWanId sets VirtualWanId field to given value.

### HasVirtualWanId

`func (o *VirtualWanOnlineStatusOpenApiVO) HasVirtualWanId() bool`

HasVirtualWanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


