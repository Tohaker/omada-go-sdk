# TopologyIsolatedAndPreConfigDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isolated** | Pointer to [**[]TopologyBriefDevice**](TopologyBriefDevice.md) | Isolated devices. | [optional] 
**Preconfig** | Pointer to [**[]TopologyBriefDevice**](TopologyBriefDevice.md) | PreConfig devices. | [optional] 
**Total** | Pointer to **int32** | Total number of devices. | [optional] 

## Methods

### NewTopologyIsolatedAndPreConfigDevice

`func NewTopologyIsolatedAndPreConfigDevice() *TopologyIsolatedAndPreConfigDevice`

NewTopologyIsolatedAndPreConfigDevice instantiates a new TopologyIsolatedAndPreConfigDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyIsolatedAndPreConfigDeviceWithDefaults

`func NewTopologyIsolatedAndPreConfigDeviceWithDefaults() *TopologyIsolatedAndPreConfigDevice`

NewTopologyIsolatedAndPreConfigDeviceWithDefaults instantiates a new TopologyIsolatedAndPreConfigDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsolated

`func (o *TopologyIsolatedAndPreConfigDevice) GetIsolated() []TopologyBriefDevice`

GetIsolated returns the Isolated field if non-nil, zero value otherwise.

### GetIsolatedOk

`func (o *TopologyIsolatedAndPreConfigDevice) GetIsolatedOk() (*[]TopologyBriefDevice, bool)`

GetIsolatedOk returns a tuple with the Isolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolated

`func (o *TopologyIsolatedAndPreConfigDevice) SetIsolated(v []TopologyBriefDevice)`

SetIsolated sets Isolated field to given value.

### HasIsolated

`func (o *TopologyIsolatedAndPreConfigDevice) HasIsolated() bool`

HasIsolated returns a boolean if a field has been set.

### GetPreconfig

`func (o *TopologyIsolatedAndPreConfigDevice) GetPreconfig() []TopologyBriefDevice`

GetPreconfig returns the Preconfig field if non-nil, zero value otherwise.

### GetPreconfigOk

`func (o *TopologyIsolatedAndPreConfigDevice) GetPreconfigOk() (*[]TopologyBriefDevice, bool)`

GetPreconfigOk returns a tuple with the Preconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconfig

`func (o *TopologyIsolatedAndPreConfigDevice) SetPreconfig(v []TopologyBriefDevice)`

SetPreconfig sets Preconfig field to given value.

### HasPreconfig

`func (o *TopologyIsolatedAndPreConfigDevice) HasPreconfig() bool`

HasPreconfig returns a boolean if a field has been set.

### GetTotal

`func (o *TopologyIsolatedAndPreConfigDevice) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TopologyIsolatedAndPreConfigDevice) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TopologyIsolatedAndPreConfigDevice) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TopologyIsolatedAndPreConfigDevice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


