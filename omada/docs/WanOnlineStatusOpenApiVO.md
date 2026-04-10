# WanOnlineStatusOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnlineDetection** | Pointer to **int32** | Port Online Detection | [optional] 
**PortUuid** | Pointer to **string** | The port UUID of WAN | [optional] 

## Methods

### NewWanOnlineStatusOpenApiVO

`func NewWanOnlineStatusOpenApiVO() *WanOnlineStatusOpenApiVO`

NewWanOnlineStatusOpenApiVO instantiates a new WanOnlineStatusOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanOnlineStatusOpenApiVOWithDefaults

`func NewWanOnlineStatusOpenApiVOWithDefaults() *WanOnlineStatusOpenApiVO`

NewWanOnlineStatusOpenApiVOWithDefaults instantiates a new WanOnlineStatusOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnlineDetection

`func (o *WanOnlineStatusOpenApiVO) GetOnlineDetection() int32`

GetOnlineDetection returns the OnlineDetection field if non-nil, zero value otherwise.

### GetOnlineDetectionOk

`func (o *WanOnlineStatusOpenApiVO) GetOnlineDetectionOk() (*int32, bool)`

GetOnlineDetectionOk returns a tuple with the OnlineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineDetection

`func (o *WanOnlineStatusOpenApiVO) SetOnlineDetection(v int32)`

SetOnlineDetection sets OnlineDetection field to given value.

### HasOnlineDetection

`func (o *WanOnlineStatusOpenApiVO) HasOnlineDetection() bool`

HasOnlineDetection returns a boolean if a field has been set.

### GetPortUuid

`func (o *WanOnlineStatusOpenApiVO) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *WanOnlineStatusOpenApiVO) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *WanOnlineStatusOpenApiVO) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.

### HasPortUuid

`func (o *WanOnlineStatusOpenApiVO) HasPortUuid() bool`

HasPortUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


