# StatusContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **bool** | Whether to query device MAC address. | [optional] 
**Name** | Pointer to **bool** | Whether to query device model name. | [optional] 
**Ports** | Pointer to [**[]ContentPortInfo**](ContentPortInfo.md) |  | [optional] 
**Uptime** | Pointer to **bool** | Whether to query device boot time. | [optional] 
**Version** | Pointer to **bool** | Whether to query device firmware version. | [optional] 

## Methods

### NewStatusContent

`func NewStatusContent() *StatusContent`

NewStatusContent instantiates a new StatusContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusContentWithDefaults

`func NewStatusContentWithDefaults() *StatusContent`

NewStatusContentWithDefaults instantiates a new StatusContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *StatusContent) GetMac() bool`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *StatusContent) GetMacOk() (*bool, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *StatusContent) SetMac(v bool)`

SetMac sets Mac field to given value.

### HasMac

`func (o *StatusContent) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *StatusContent) GetName() bool`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatusContent) GetNameOk() (*bool, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatusContent) SetName(v bool)`

SetName sets Name field to given value.

### HasName

`func (o *StatusContent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *StatusContent) GetPorts() []ContentPortInfo`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *StatusContent) GetPortsOk() (*[]ContentPortInfo, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *StatusContent) SetPorts(v []ContentPortInfo)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *StatusContent) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetUptime

`func (o *StatusContent) GetUptime() bool`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *StatusContent) GetUptimeOk() (*bool, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *StatusContent) SetUptime(v bool)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *StatusContent) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVersion

`func (o *StatusContent) GetVersion() bool`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StatusContent) GetVersionOk() (*bool, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StatusContent) SetVersion(v bool)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StatusContent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


