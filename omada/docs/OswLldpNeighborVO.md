# OswLldpNeighborVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **string** | Capabilities | [optional] 
**DeviceId** | Pointer to **string** | Device ID | [optional] 
**NeighborPortId** | Pointer to **string** | The port on which the neighboring device is connected to the current switch | [optional] 
**PortId** | Pointer to **int32** | Port ID | [optional] 
**StandardPort** | Pointer to **string** | Standard Port | [optional] 
**SystemName** | Pointer to **string** | System Name | [optional] 
**Ttl** | Pointer to **int32** | Time that neighbor LLDP packets can survive | [optional] 

## Methods

### NewOswLldpNeighborVO

`func NewOswLldpNeighborVO() *OswLldpNeighborVO`

NewOswLldpNeighborVO instantiates a new OswLldpNeighborVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswLldpNeighborVOWithDefaults

`func NewOswLldpNeighborVOWithDefaults() *OswLldpNeighborVO`

NewOswLldpNeighborVOWithDefaults instantiates a new OswLldpNeighborVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *OswLldpNeighborVO) GetCapabilities() string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *OswLldpNeighborVO) GetCapabilitiesOk() (*string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *OswLldpNeighborVO) SetCapabilities(v string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *OswLldpNeighborVO) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetDeviceId

`func (o *OswLldpNeighborVO) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *OswLldpNeighborVO) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *OswLldpNeighborVO) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *OswLldpNeighborVO) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetNeighborPortId

`func (o *OswLldpNeighborVO) GetNeighborPortId() string`

GetNeighborPortId returns the NeighborPortId field if non-nil, zero value otherwise.

### GetNeighborPortIdOk

`func (o *OswLldpNeighborVO) GetNeighborPortIdOk() (*string, bool)`

GetNeighborPortIdOk returns a tuple with the NeighborPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborPortId

`func (o *OswLldpNeighborVO) SetNeighborPortId(v string)`

SetNeighborPortId sets NeighborPortId field to given value.

### HasNeighborPortId

`func (o *OswLldpNeighborVO) HasNeighborPortId() bool`

HasNeighborPortId returns a boolean if a field has been set.

### GetPortId

`func (o *OswLldpNeighborVO) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *OswLldpNeighborVO) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *OswLldpNeighborVO) SetPortId(v int32)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *OswLldpNeighborVO) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswLldpNeighborVO) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswLldpNeighborVO) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswLldpNeighborVO) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswLldpNeighborVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetSystemName

`func (o *OswLldpNeighborVO) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *OswLldpNeighborVO) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *OswLldpNeighborVO) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *OswLldpNeighborVO) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTtl

`func (o *OswLldpNeighborVO) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *OswLldpNeighborVO) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *OswLldpNeighborVO) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *OswLldpNeighborVO) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


