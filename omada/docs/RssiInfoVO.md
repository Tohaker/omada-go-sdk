# RssiInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientModel** | Pointer to **string** | Client model for icon display | [optional] 
**ClientName** | Pointer to **string** | Client name | [optional] 
**ClientType** | Pointer to **string** | Client type for icon display (e.g. Mobile, Laptop, IPC) | [optional] 
**ConnectDeviceMac** | Pointer to **string** | Client-connected device mac | [optional] 
**ConnectDeviceName** | Pointer to **string** | Client-connected device name | [optional] 
**DeviceType** | Pointer to **string** | Client-connected device type | [optional] 
**Ip** | Pointer to **string** | ip | [optional] 
**Mac** | Pointer to **string** | Mac address | [optional] 
**Rssi** | Pointer to **int32** | rssi | [optional] 
**Ssid** | Pointer to **string** | ssid | [optional] 

## Methods

### NewRssiInfoVO

`func NewRssiInfoVO() *RssiInfoVO`

NewRssiInfoVO instantiates a new RssiInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRssiInfoVOWithDefaults

`func NewRssiInfoVOWithDefaults() *RssiInfoVO`

NewRssiInfoVOWithDefaults instantiates a new RssiInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientModel

`func (o *RssiInfoVO) GetClientModel() string`

GetClientModel returns the ClientModel field if non-nil, zero value otherwise.

### GetClientModelOk

`func (o *RssiInfoVO) GetClientModelOk() (*string, bool)`

GetClientModelOk returns a tuple with the ClientModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientModel

`func (o *RssiInfoVO) SetClientModel(v string)`

SetClientModel sets ClientModel field to given value.

### HasClientModel

`func (o *RssiInfoVO) HasClientModel() bool`

HasClientModel returns a boolean if a field has been set.

### GetClientName

`func (o *RssiInfoVO) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *RssiInfoVO) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *RssiInfoVO) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *RssiInfoVO) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientType

`func (o *RssiInfoVO) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *RssiInfoVO) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *RssiInfoVO) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *RssiInfoVO) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetConnectDeviceMac

`func (o *RssiInfoVO) GetConnectDeviceMac() string`

GetConnectDeviceMac returns the ConnectDeviceMac field if non-nil, zero value otherwise.

### GetConnectDeviceMacOk

`func (o *RssiInfoVO) GetConnectDeviceMacOk() (*string, bool)`

GetConnectDeviceMacOk returns a tuple with the ConnectDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectDeviceMac

`func (o *RssiInfoVO) SetConnectDeviceMac(v string)`

SetConnectDeviceMac sets ConnectDeviceMac field to given value.

### HasConnectDeviceMac

`func (o *RssiInfoVO) HasConnectDeviceMac() bool`

HasConnectDeviceMac returns a boolean if a field has been set.

### GetConnectDeviceName

`func (o *RssiInfoVO) GetConnectDeviceName() string`

GetConnectDeviceName returns the ConnectDeviceName field if non-nil, zero value otherwise.

### GetConnectDeviceNameOk

`func (o *RssiInfoVO) GetConnectDeviceNameOk() (*string, bool)`

GetConnectDeviceNameOk returns a tuple with the ConnectDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectDeviceName

`func (o *RssiInfoVO) SetConnectDeviceName(v string)`

SetConnectDeviceName sets ConnectDeviceName field to given value.

### HasConnectDeviceName

`func (o *RssiInfoVO) HasConnectDeviceName() bool`

HasConnectDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *RssiInfoVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *RssiInfoVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *RssiInfoVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *RssiInfoVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetIp

`func (o *RssiInfoVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RssiInfoVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RssiInfoVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RssiInfoVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *RssiInfoVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *RssiInfoVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *RssiInfoVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *RssiInfoVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetRssi

`func (o *RssiInfoVO) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *RssiInfoVO) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *RssiInfoVO) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *RssiInfoVO) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetSsid

`func (o *RssiInfoVO) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *RssiInfoVO) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *RssiInfoVO) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *RssiInfoVO) HasSsid() bool`

HasSsid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


