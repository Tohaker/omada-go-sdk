# VlanNetworkDeliverBriefDataVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveringDevices** | Pointer to **[]string** | Delivering devices | [optional] 
**DeliveringStacks** | Pointer to **[]string** | Delivering stacks | [optional] 
**DoneDevicesNum** | Pointer to **int32** | Delivered device num | [optional] 
**FailDevices** | Pointer to **[]string** | Failed devices | [optional] 
**FailStacks** | Pointer to **[]string** | Failed stacks | [optional] 
**State** | Pointer to **int32** | Delivering state. 0: free ( it means the network has been already delivered)  1: delivering 2. deliver done. | [optional] 
**SuccessDevices** | Pointer to **[]string** | Successfully delivered devices | [optional] 
**SuccessStacks** | Pointer to **[]string** | Successfully delivered stacks | [optional] 
**TotalDevicesNum** | Pointer to **int32** | Total device num | [optional] 

## Methods

### NewVlanNetworkDeliverBriefDataVO

`func NewVlanNetworkDeliverBriefDataVO() *VlanNetworkDeliverBriefDataVO`

NewVlanNetworkDeliverBriefDataVO instantiates a new VlanNetworkDeliverBriefDataVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanNetworkDeliverBriefDataVOWithDefaults

`func NewVlanNetworkDeliverBriefDataVOWithDefaults() *VlanNetworkDeliverBriefDataVO`

NewVlanNetworkDeliverBriefDataVOWithDefaults instantiates a new VlanNetworkDeliverBriefDataVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveringDevices

`func (o *VlanNetworkDeliverBriefDataVO) GetDeliveringDevices() []string`

GetDeliveringDevices returns the DeliveringDevices field if non-nil, zero value otherwise.

### GetDeliveringDevicesOk

`func (o *VlanNetworkDeliverBriefDataVO) GetDeliveringDevicesOk() (*[]string, bool)`

GetDeliveringDevicesOk returns a tuple with the DeliveringDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveringDevices

`func (o *VlanNetworkDeliverBriefDataVO) SetDeliveringDevices(v []string)`

SetDeliveringDevices sets DeliveringDevices field to given value.

### HasDeliveringDevices

`func (o *VlanNetworkDeliverBriefDataVO) HasDeliveringDevices() bool`

HasDeliveringDevices returns a boolean if a field has been set.

### GetDeliveringStacks

`func (o *VlanNetworkDeliverBriefDataVO) GetDeliveringStacks() []string`

GetDeliveringStacks returns the DeliveringStacks field if non-nil, zero value otherwise.

### GetDeliveringStacksOk

`func (o *VlanNetworkDeliverBriefDataVO) GetDeliveringStacksOk() (*[]string, bool)`

GetDeliveringStacksOk returns a tuple with the DeliveringStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveringStacks

`func (o *VlanNetworkDeliverBriefDataVO) SetDeliveringStacks(v []string)`

SetDeliveringStacks sets DeliveringStacks field to given value.

### HasDeliveringStacks

`func (o *VlanNetworkDeliverBriefDataVO) HasDeliveringStacks() bool`

HasDeliveringStacks returns a boolean if a field has been set.

### GetDoneDevicesNum

`func (o *VlanNetworkDeliverBriefDataVO) GetDoneDevicesNum() int32`

GetDoneDevicesNum returns the DoneDevicesNum field if non-nil, zero value otherwise.

### GetDoneDevicesNumOk

`func (o *VlanNetworkDeliverBriefDataVO) GetDoneDevicesNumOk() (*int32, bool)`

GetDoneDevicesNumOk returns a tuple with the DoneDevicesNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneDevicesNum

`func (o *VlanNetworkDeliverBriefDataVO) SetDoneDevicesNum(v int32)`

SetDoneDevicesNum sets DoneDevicesNum field to given value.

### HasDoneDevicesNum

`func (o *VlanNetworkDeliverBriefDataVO) HasDoneDevicesNum() bool`

HasDoneDevicesNum returns a boolean if a field has been set.

### GetFailDevices

`func (o *VlanNetworkDeliverBriefDataVO) GetFailDevices() []string`

GetFailDevices returns the FailDevices field if non-nil, zero value otherwise.

### GetFailDevicesOk

`func (o *VlanNetworkDeliverBriefDataVO) GetFailDevicesOk() (*[]string, bool)`

GetFailDevicesOk returns a tuple with the FailDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDevices

`func (o *VlanNetworkDeliverBriefDataVO) SetFailDevices(v []string)`

SetFailDevices sets FailDevices field to given value.

### HasFailDevices

`func (o *VlanNetworkDeliverBriefDataVO) HasFailDevices() bool`

HasFailDevices returns a boolean if a field has been set.

### GetFailStacks

`func (o *VlanNetworkDeliverBriefDataVO) GetFailStacks() []string`

GetFailStacks returns the FailStacks field if non-nil, zero value otherwise.

### GetFailStacksOk

`func (o *VlanNetworkDeliverBriefDataVO) GetFailStacksOk() (*[]string, bool)`

GetFailStacksOk returns a tuple with the FailStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailStacks

`func (o *VlanNetworkDeliverBriefDataVO) SetFailStacks(v []string)`

SetFailStacks sets FailStacks field to given value.

### HasFailStacks

`func (o *VlanNetworkDeliverBriefDataVO) HasFailStacks() bool`

HasFailStacks returns a boolean if a field has been set.

### GetState

`func (o *VlanNetworkDeliverBriefDataVO) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VlanNetworkDeliverBriefDataVO) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VlanNetworkDeliverBriefDataVO) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *VlanNetworkDeliverBriefDataVO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSuccessDevices

`func (o *VlanNetworkDeliverBriefDataVO) GetSuccessDevices() []string`

GetSuccessDevices returns the SuccessDevices field if non-nil, zero value otherwise.

### GetSuccessDevicesOk

`func (o *VlanNetworkDeliverBriefDataVO) GetSuccessDevicesOk() (*[]string, bool)`

GetSuccessDevicesOk returns a tuple with the SuccessDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessDevices

`func (o *VlanNetworkDeliverBriefDataVO) SetSuccessDevices(v []string)`

SetSuccessDevices sets SuccessDevices field to given value.

### HasSuccessDevices

`func (o *VlanNetworkDeliverBriefDataVO) HasSuccessDevices() bool`

HasSuccessDevices returns a boolean if a field has been set.

### GetSuccessStacks

`func (o *VlanNetworkDeliverBriefDataVO) GetSuccessStacks() []string`

GetSuccessStacks returns the SuccessStacks field if non-nil, zero value otherwise.

### GetSuccessStacksOk

`func (o *VlanNetworkDeliverBriefDataVO) GetSuccessStacksOk() (*[]string, bool)`

GetSuccessStacksOk returns a tuple with the SuccessStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessStacks

`func (o *VlanNetworkDeliverBriefDataVO) SetSuccessStacks(v []string)`

SetSuccessStacks sets SuccessStacks field to given value.

### HasSuccessStacks

`func (o *VlanNetworkDeliverBriefDataVO) HasSuccessStacks() bool`

HasSuccessStacks returns a boolean if a field has been set.

### GetTotalDevicesNum

`func (o *VlanNetworkDeliverBriefDataVO) GetTotalDevicesNum() int32`

GetTotalDevicesNum returns the TotalDevicesNum field if non-nil, zero value otherwise.

### GetTotalDevicesNumOk

`func (o *VlanNetworkDeliverBriefDataVO) GetTotalDevicesNumOk() (*int32, bool)`

GetTotalDevicesNumOk returns a tuple with the TotalDevicesNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDevicesNum

`func (o *VlanNetworkDeliverBriefDataVO) SetTotalDevicesNum(v int32)`

SetTotalDevicesNum sets TotalDevicesNum field to given value.

### HasTotalDevicesNum

`func (o *VlanNetworkDeliverBriefDataVO) HasTotalDevicesNum() bool`

HasTotalDevicesNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


