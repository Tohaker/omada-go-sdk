# PoeUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Device MAC | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**PoePorts** | Pointer to [**[]PortPoe**](PortPoe.md) | Device PoE Ports | [optional] 
**PortNum** | Pointer to **int32** | Device port Number | [optional] 
**TotalPercentUsed** | Pointer to **float64** | Percentage of total power used (integer reserved) | [optional] 
**TotalPower** | Pointer to **int32** | The total power of the switch port&#39;s POE (in watts) | [optional] 
**TotalPowerUsed** | Pointer to **int32** | Total power used (in watts) | [optional] 

## Methods

### NewPoeUsage

`func NewPoeUsage() *PoeUsage`

NewPoeUsage instantiates a new PoeUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoeUsageWithDefaults

`func NewPoeUsageWithDefaults() *PoeUsage`

NewPoeUsageWithDefaults instantiates a new PoeUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *PoeUsage) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *PoeUsage) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *PoeUsage) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *PoeUsage) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *PoeUsage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoeUsage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoeUsage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoeUsage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoePorts

`func (o *PoeUsage) GetPoePorts() []PortPoe`

GetPoePorts returns the PoePorts field if non-nil, zero value otherwise.

### GetPoePortsOk

`func (o *PoeUsage) GetPoePortsOk() (*[]PortPoe, bool)`

GetPoePortsOk returns a tuple with the PoePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePorts

`func (o *PoeUsage) SetPoePorts(v []PortPoe)`

SetPoePorts sets PoePorts field to given value.

### HasPoePorts

`func (o *PoeUsage) HasPoePorts() bool`

HasPoePorts returns a boolean if a field has been set.

### GetPortNum

`func (o *PoeUsage) GetPortNum() int32`

GetPortNum returns the PortNum field if non-nil, zero value otherwise.

### GetPortNumOk

`func (o *PoeUsage) GetPortNumOk() (*int32, bool)`

GetPortNumOk returns a tuple with the PortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNum

`func (o *PoeUsage) SetPortNum(v int32)`

SetPortNum sets PortNum field to given value.

### HasPortNum

`func (o *PoeUsage) HasPortNum() bool`

HasPortNum returns a boolean if a field has been set.

### GetTotalPercentUsed

`func (o *PoeUsage) GetTotalPercentUsed() float64`

GetTotalPercentUsed returns the TotalPercentUsed field if non-nil, zero value otherwise.

### GetTotalPercentUsedOk

`func (o *PoeUsage) GetTotalPercentUsedOk() (*float64, bool)`

GetTotalPercentUsedOk returns a tuple with the TotalPercentUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPercentUsed

`func (o *PoeUsage) SetTotalPercentUsed(v float64)`

SetTotalPercentUsed sets TotalPercentUsed field to given value.

### HasTotalPercentUsed

`func (o *PoeUsage) HasTotalPercentUsed() bool`

HasTotalPercentUsed returns a boolean if a field has been set.

### GetTotalPower

`func (o *PoeUsage) GetTotalPower() int32`

GetTotalPower returns the TotalPower field if non-nil, zero value otherwise.

### GetTotalPowerOk

`func (o *PoeUsage) GetTotalPowerOk() (*int32, bool)`

GetTotalPowerOk returns a tuple with the TotalPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPower

`func (o *PoeUsage) SetTotalPower(v int32)`

SetTotalPower sets TotalPower field to given value.

### HasTotalPower

`func (o *PoeUsage) HasTotalPower() bool`

HasTotalPower returns a boolean if a field has been set.

### GetTotalPowerUsed

`func (o *PoeUsage) GetTotalPowerUsed() int32`

GetTotalPowerUsed returns the TotalPowerUsed field if non-nil, zero value otherwise.

### GetTotalPowerUsedOk

`func (o *PoeUsage) GetTotalPowerUsedOk() (*int32, bool)`

GetTotalPowerUsedOk returns a tuple with the TotalPowerUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPowerUsed

`func (o *PoeUsage) SetTotalPowerUsed(v int32)`

SetTotalPowerUsed sets TotalPowerUsed field to given value.

### HasTotalPowerUsed

`func (o *PoeUsage) HasTotalPowerUsed() bool`

HasTotalPowerUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


