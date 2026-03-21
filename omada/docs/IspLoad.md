# IspLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]IspLoadStat**](IspLoadStat.md) | WAN port data list | [optional] 
**PortId** | Pointer to **int32** | Port ID | [optional] 
**PortName** | Pointer to **string** | Port name | [optional] 

## Methods

### NewIspLoad

`func NewIspLoad() *IspLoad`

NewIspLoad instantiates a new IspLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIspLoadWithDefaults

`func NewIspLoadWithDefaults() *IspLoad`

NewIspLoadWithDefaults instantiates a new IspLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IspLoad) GetData() []IspLoadStat`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IspLoad) GetDataOk() (*[]IspLoadStat, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IspLoad) SetData(v []IspLoadStat)`

SetData sets Data field to given value.

### HasData

`func (o *IspLoad) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPortId

`func (o *IspLoad) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *IspLoad) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *IspLoad) SetPortId(v int32)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *IspLoad) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortName

`func (o *IspLoad) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *IspLoad) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *IspLoad) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *IspLoad) HasPortName() bool`

HasPortName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


