# SelectPortBindingBriefVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceList** | Pointer to [**[]PortBindingVO**](PortBindingVO.md) | Device List | [optional] 
**FlowControlEnable** | Pointer to **bool** | Enable Flow Control, only valid when creating network. | [optional] 
**InternetPorts** | Pointer to **[]string** | Internet Ports | [optional] 
**PortIsolationEnable** | Pointer to **bool** | Enable Port Isolation, only valid when creating network. | [optional] 
**TagIds** | Pointer to **[]string** | Tag ID Set | [optional] 

## Methods

### NewSelectPortBindingBriefVO

`func NewSelectPortBindingBriefVO() *SelectPortBindingBriefVO`

NewSelectPortBindingBriefVO instantiates a new SelectPortBindingBriefVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectPortBindingBriefVOWithDefaults

`func NewSelectPortBindingBriefVOWithDefaults() *SelectPortBindingBriefVO`

NewSelectPortBindingBriefVOWithDefaults instantiates a new SelectPortBindingBriefVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceList

`func (o *SelectPortBindingBriefVO) GetDeviceList() []PortBindingVO`

GetDeviceList returns the DeviceList field if non-nil, zero value otherwise.

### GetDeviceListOk

`func (o *SelectPortBindingBriefVO) GetDeviceListOk() (*[]PortBindingVO, bool)`

GetDeviceListOk returns a tuple with the DeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceList

`func (o *SelectPortBindingBriefVO) SetDeviceList(v []PortBindingVO)`

SetDeviceList sets DeviceList field to given value.

### HasDeviceList

`func (o *SelectPortBindingBriefVO) HasDeviceList() bool`

HasDeviceList returns a boolean if a field has been set.

### GetFlowControlEnable

`func (o *SelectPortBindingBriefVO) GetFlowControlEnable() bool`

GetFlowControlEnable returns the FlowControlEnable field if non-nil, zero value otherwise.

### GetFlowControlEnableOk

`func (o *SelectPortBindingBriefVO) GetFlowControlEnableOk() (*bool, bool)`

GetFlowControlEnableOk returns a tuple with the FlowControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlEnable

`func (o *SelectPortBindingBriefVO) SetFlowControlEnable(v bool)`

SetFlowControlEnable sets FlowControlEnable field to given value.

### HasFlowControlEnable

`func (o *SelectPortBindingBriefVO) HasFlowControlEnable() bool`

HasFlowControlEnable returns a boolean if a field has been set.

### GetInternetPorts

`func (o *SelectPortBindingBriefVO) GetInternetPorts() []string`

GetInternetPorts returns the InternetPorts field if non-nil, zero value otherwise.

### GetInternetPortsOk

`func (o *SelectPortBindingBriefVO) GetInternetPortsOk() (*[]string, bool)`

GetInternetPortsOk returns a tuple with the InternetPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetPorts

`func (o *SelectPortBindingBriefVO) SetInternetPorts(v []string)`

SetInternetPorts sets InternetPorts field to given value.

### HasInternetPorts

`func (o *SelectPortBindingBriefVO) HasInternetPorts() bool`

HasInternetPorts returns a boolean if a field has been set.

### GetPortIsolationEnable

`func (o *SelectPortBindingBriefVO) GetPortIsolationEnable() bool`

GetPortIsolationEnable returns the PortIsolationEnable field if non-nil, zero value otherwise.

### GetPortIsolationEnableOk

`func (o *SelectPortBindingBriefVO) GetPortIsolationEnableOk() (*bool, bool)`

GetPortIsolationEnableOk returns a tuple with the PortIsolationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolationEnable

`func (o *SelectPortBindingBriefVO) SetPortIsolationEnable(v bool)`

SetPortIsolationEnable sets PortIsolationEnable field to given value.

### HasPortIsolationEnable

`func (o *SelectPortBindingBriefVO) HasPortIsolationEnable() bool`

HasPortIsolationEnable returns a boolean if a field has been set.

### GetTagIds

`func (o *SelectPortBindingBriefVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *SelectPortBindingBriefVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *SelectPortBindingBriefVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *SelectPortBindingBriefVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


