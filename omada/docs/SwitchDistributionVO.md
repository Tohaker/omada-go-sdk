# SwitchDistributionVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Others** | Pointer to [**DeviceClientVO**](DeviceClientVO.md) |  | [optional] 
**Switches** | Pointer to [**[]DeviceClientVO**](DeviceClientVO.md) |  | [optional] 
**TotalSwitchClients** | Pointer to **int32** |  | [optional] 
**TotalSwitchDistribution** | Pointer to **float64** |  | [optional] 

## Methods

### NewSwitchDistributionVO

`func NewSwitchDistributionVO() *SwitchDistributionVO`

NewSwitchDistributionVO instantiates a new SwitchDistributionVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchDistributionVOWithDefaults

`func NewSwitchDistributionVOWithDefaults() *SwitchDistributionVO`

NewSwitchDistributionVOWithDefaults instantiates a new SwitchDistributionVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOthers

`func (o *SwitchDistributionVO) GetOthers() DeviceClientVO`

GetOthers returns the Others field if non-nil, zero value otherwise.

### GetOthersOk

`func (o *SwitchDistributionVO) GetOthersOk() (*DeviceClientVO, bool)`

GetOthersOk returns a tuple with the Others field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthers

`func (o *SwitchDistributionVO) SetOthers(v DeviceClientVO)`

SetOthers sets Others field to given value.

### HasOthers

`func (o *SwitchDistributionVO) HasOthers() bool`

HasOthers returns a boolean if a field has been set.

### GetSwitches

`func (o *SwitchDistributionVO) GetSwitches() []DeviceClientVO`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *SwitchDistributionVO) GetSwitchesOk() (*[]DeviceClientVO, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *SwitchDistributionVO) SetSwitches(v []DeviceClientVO)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *SwitchDistributionVO) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetTotalSwitchClients

`func (o *SwitchDistributionVO) GetTotalSwitchClients() int32`

GetTotalSwitchClients returns the TotalSwitchClients field if non-nil, zero value otherwise.

### GetTotalSwitchClientsOk

`func (o *SwitchDistributionVO) GetTotalSwitchClientsOk() (*int32, bool)`

GetTotalSwitchClientsOk returns a tuple with the TotalSwitchClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSwitchClients

`func (o *SwitchDistributionVO) SetTotalSwitchClients(v int32)`

SetTotalSwitchClients sets TotalSwitchClients field to given value.

### HasTotalSwitchClients

`func (o *SwitchDistributionVO) HasTotalSwitchClients() bool`

HasTotalSwitchClients returns a boolean if a field has been set.

### GetTotalSwitchDistribution

`func (o *SwitchDistributionVO) GetTotalSwitchDistribution() float64`

GetTotalSwitchDistribution returns the TotalSwitchDistribution field if non-nil, zero value otherwise.

### GetTotalSwitchDistributionOk

`func (o *SwitchDistributionVO) GetTotalSwitchDistributionOk() (*float64, bool)`

GetTotalSwitchDistributionOk returns a tuple with the TotalSwitchDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSwitchDistribution

`func (o *SwitchDistributionVO) SetTotalSwitchDistribution(v float64)`

SetTotalSwitchDistribution sets TotalSwitchDistribution field to given value.

### HasTotalSwitchDistribution

`func (o *SwitchDistributionVO) HasTotalSwitchDistribution() bool`

HasTotalSwitchDistribution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


