# ChangeInternetStateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **int32** | operation, 1 for open, and 0 for close | [optional] 
**PortId** | Pointer to **int32** | port ID | [optional] 
**VirtualWanId** | Pointer to **string** | If operating virtualWan, should give virtualWanId | [optional] 

## Methods

### NewChangeInternetStateOpenApiVO

`func NewChangeInternetStateOpenApiVO() *ChangeInternetStateOpenApiVO`

NewChangeInternetStateOpenApiVO instantiates a new ChangeInternetStateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeInternetStateOpenApiVOWithDefaults

`func NewChangeInternetStateOpenApiVOWithDefaults() *ChangeInternetStateOpenApiVO`

NewChangeInternetStateOpenApiVOWithDefaults instantiates a new ChangeInternetStateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ChangeInternetStateOpenApiVO) GetOperation() int32`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ChangeInternetStateOpenApiVO) GetOperationOk() (*int32, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ChangeInternetStateOpenApiVO) SetOperation(v int32)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ChangeInternetStateOpenApiVO) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPortId

`func (o *ChangeInternetStateOpenApiVO) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *ChangeInternetStateOpenApiVO) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *ChangeInternetStateOpenApiVO) SetPortId(v int32)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *ChangeInternetStateOpenApiVO) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetVirtualWanId

`func (o *ChangeInternetStateOpenApiVO) GetVirtualWanId() string`

GetVirtualWanId returns the VirtualWanId field if non-nil, zero value otherwise.

### GetVirtualWanIdOk

`func (o *ChangeInternetStateOpenApiVO) GetVirtualWanIdOk() (*string, bool)`

GetVirtualWanIdOk returns a tuple with the VirtualWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanId

`func (o *ChangeInternetStateOpenApiVO) SetVirtualWanId(v string)`

SetVirtualWanId sets VirtualWanId field to given value.

### HasVirtualWanId

`func (o *ChangeInternetStateOpenApiVO) HasVirtualWanId() bool`

HasVirtualWanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


