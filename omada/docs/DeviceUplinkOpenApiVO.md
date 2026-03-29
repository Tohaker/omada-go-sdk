# DeviceUplinkOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplex** | Pointer to **int32** | Device uplink port duplex mode, duplex should be a value as follows: 0: Auto; 1: Half; 2: Full. | [optional] 
**LinkSpeed** | Pointer to **int32** | Device uplink port linkSpeed, linkSpeed should be a value as follows: 0: Auto; 1: 10M; 2: 100M; 3: 1000M; 4: 2500M; 5: 10G; 6: 5G; 7: 25G, 8: 100G. | [optional] 
**Mac** | Pointer to **string** | Current Device MAC | [optional] 
**UplinkDeviceMac** | Pointer to **string** | Uplink device mac | [optional] 
**UplinkDeviceName** | Pointer to **string** | Uplink device name | [optional] 
**UplinkDevicePort** | Pointer to **string** | Uplink device port | [optional] 

## Methods

### NewDeviceUplinkOpenApiVO

`func NewDeviceUplinkOpenApiVO() *DeviceUplinkOpenApiVO`

NewDeviceUplinkOpenApiVO instantiates a new DeviceUplinkOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceUplinkOpenApiVOWithDefaults

`func NewDeviceUplinkOpenApiVOWithDefaults() *DeviceUplinkOpenApiVO`

NewDeviceUplinkOpenApiVOWithDefaults instantiates a new DeviceUplinkOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplex

`func (o *DeviceUplinkOpenApiVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *DeviceUplinkOpenApiVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *DeviceUplinkOpenApiVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *DeviceUplinkOpenApiVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *DeviceUplinkOpenApiVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *DeviceUplinkOpenApiVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *DeviceUplinkOpenApiVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *DeviceUplinkOpenApiVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetMac

`func (o *DeviceUplinkOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceUplinkOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceUplinkOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceUplinkOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetUplinkDeviceMac

`func (o *DeviceUplinkOpenApiVO) GetUplinkDeviceMac() string`

GetUplinkDeviceMac returns the UplinkDeviceMac field if non-nil, zero value otherwise.

### GetUplinkDeviceMacOk

`func (o *DeviceUplinkOpenApiVO) GetUplinkDeviceMacOk() (*string, bool)`

GetUplinkDeviceMacOk returns a tuple with the UplinkDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkDeviceMac

`func (o *DeviceUplinkOpenApiVO) SetUplinkDeviceMac(v string)`

SetUplinkDeviceMac sets UplinkDeviceMac field to given value.

### HasUplinkDeviceMac

`func (o *DeviceUplinkOpenApiVO) HasUplinkDeviceMac() bool`

HasUplinkDeviceMac returns a boolean if a field has been set.

### GetUplinkDeviceName

`func (o *DeviceUplinkOpenApiVO) GetUplinkDeviceName() string`

GetUplinkDeviceName returns the UplinkDeviceName field if non-nil, zero value otherwise.

### GetUplinkDeviceNameOk

`func (o *DeviceUplinkOpenApiVO) GetUplinkDeviceNameOk() (*string, bool)`

GetUplinkDeviceNameOk returns a tuple with the UplinkDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkDeviceName

`func (o *DeviceUplinkOpenApiVO) SetUplinkDeviceName(v string)`

SetUplinkDeviceName sets UplinkDeviceName field to given value.

### HasUplinkDeviceName

`func (o *DeviceUplinkOpenApiVO) HasUplinkDeviceName() bool`

HasUplinkDeviceName returns a boolean if a field has been set.

### GetUplinkDevicePort

`func (o *DeviceUplinkOpenApiVO) GetUplinkDevicePort() string`

GetUplinkDevicePort returns the UplinkDevicePort field if non-nil, zero value otherwise.

### GetUplinkDevicePortOk

`func (o *DeviceUplinkOpenApiVO) GetUplinkDevicePortOk() (*string, bool)`

GetUplinkDevicePortOk returns a tuple with the UplinkDevicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkDevicePort

`func (o *DeviceUplinkOpenApiVO) SetUplinkDevicePort(v string)`

SetUplinkDevicePort sets UplinkDevicePort field to given value.

### HasUplinkDevicePort

`func (o *DeviceUplinkOpenApiVO) HasUplinkDevicePort() bool`

HasUplinkDevicePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


