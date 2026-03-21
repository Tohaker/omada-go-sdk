# IspLoadVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]IspLoadStatVO**](IspLoadStatVO.md) | WAN port ISP load stat data list | [optional] 
**DefaultWan** | Pointer to **bool** | Whether it is the default WAN port | [optional] 
**PortId** | Pointer to **int32** | Port ID | [optional] 
**PortName** | Pointer to **string** | Port Name | [optional] 

## Methods

### NewIspLoadVO

`func NewIspLoadVO() *IspLoadVO`

NewIspLoadVO instantiates a new IspLoadVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIspLoadVOWithDefaults

`func NewIspLoadVOWithDefaults() *IspLoadVO`

NewIspLoadVOWithDefaults instantiates a new IspLoadVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IspLoadVO) GetData() []IspLoadStatVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IspLoadVO) GetDataOk() (*[]IspLoadStatVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IspLoadVO) SetData(v []IspLoadStatVO)`

SetData sets Data field to given value.

### HasData

`func (o *IspLoadVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDefaultWan

`func (o *IspLoadVO) GetDefaultWan() bool`

GetDefaultWan returns the DefaultWan field if non-nil, zero value otherwise.

### GetDefaultWanOk

`func (o *IspLoadVO) GetDefaultWanOk() (*bool, bool)`

GetDefaultWanOk returns a tuple with the DefaultWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWan

`func (o *IspLoadVO) SetDefaultWan(v bool)`

SetDefaultWan sets DefaultWan field to given value.

### HasDefaultWan

`func (o *IspLoadVO) HasDefaultWan() bool`

HasDefaultWan returns a boolean if a field has been set.

### GetPortId

`func (o *IspLoadVO) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *IspLoadVO) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *IspLoadVO) SetPortId(v int32)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *IspLoadVO) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortName

`func (o *IspLoadVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *IspLoadVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *IspLoadVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *IspLoadVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


