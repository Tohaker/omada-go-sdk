# PortBindingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lags** | Pointer to **[]int32** | Selected lags. | [optional] 
**Mac** | Pointer to **string** | Device mac for Site Setting or Device template ID for Site Template Setting | [optional] 
**Ports** | Pointer to **[]string** | Selected ports. | [optional] 
**StackId** | Pointer to **string** | Stack ID.Site Template Setting ignores this field | [optional] 
**Type** | Pointer to **int32** | Device type, 1: gateway  2: switch  3: ap | [optional] 

## Methods

### NewPortBindingVO

`func NewPortBindingVO() *PortBindingVO`

NewPortBindingVO instantiates a new PortBindingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortBindingVOWithDefaults

`func NewPortBindingVOWithDefaults() *PortBindingVO`

NewPortBindingVOWithDefaults instantiates a new PortBindingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLags

`func (o *PortBindingVO) GetLags() []int32`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *PortBindingVO) GetLagsOk() (*[]int32, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *PortBindingVO) SetLags(v []int32)`

SetLags sets Lags field to given value.

### HasLags

`func (o *PortBindingVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetMac

`func (o *PortBindingVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *PortBindingVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *PortBindingVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *PortBindingVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPorts

`func (o *PortBindingVO) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *PortBindingVO) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *PortBindingVO) SetPorts(v []string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *PortBindingVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetStackId

`func (o *PortBindingVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *PortBindingVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *PortBindingVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *PortBindingVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetType

`func (o *PortBindingVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortBindingVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortBindingVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *PortBindingVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


