# UsbLteSettingConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **int32** | 0:Auto; 1:PAP; 2:CHAP. | 
**AutoConfig** | Pointer to [**AutoConfigOpenApiVO**](AutoConfigOpenApiVO.md) |  | [optional] 
**ConfigType** | **int32** | 0: Auto; 1：Manually | 
**ConnectionMode** | **int32** | 1: Connect Automatically, 2: Connect Manually. | 
**DhcpOptions** | Pointer to [**[]WanDhcpOptionOpenApiVO**](WanDhcpOptionOpenApiVO.md) |  | [optional] 
**DnsConfig** | Pointer to [**DnsConfigOpenApiVO**](DnsConfigOpenApiVO.md) |  | [optional] 
**DnsEnable** | **bool** | 0:off,1:on. | 
**ManuallyConfig** | Pointer to [**ManuallyConfigOpenApiVO**](ManuallyConfigOpenApiVO.md) |  | [optional] 
**MtuSize** | **int32** | MTU ranges from 576 ~ 1500. | 
**Pin** | Pointer to **string** | It is required when [usbModemMsgId] is 1 or 3. | [optional] 
**PortDescription** | Pointer to **string** | Port description should contain 1 to 32 characters. | [optional] 
**PortId** | **string** | Port ID | 

## Methods

### NewUsbLteSettingConfigOpenApiVO

`func NewUsbLteSettingConfigOpenApiVO(authType int32, configType int32, connectionMode int32, dnsEnable bool, mtuSize int32, portId string, ) *UsbLteSettingConfigOpenApiVO`

NewUsbLteSettingConfigOpenApiVO instantiates a new UsbLteSettingConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsbLteSettingConfigOpenApiVOWithDefaults

`func NewUsbLteSettingConfigOpenApiVOWithDefaults() *UsbLteSettingConfigOpenApiVO`

NewUsbLteSettingConfigOpenApiVOWithDefaults instantiates a new UsbLteSettingConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *UsbLteSettingConfigOpenApiVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *UsbLteSettingConfigOpenApiVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *UsbLteSettingConfigOpenApiVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.


### GetAutoConfig

`func (o *UsbLteSettingConfigOpenApiVO) GetAutoConfig() AutoConfigOpenApiVO`

GetAutoConfig returns the AutoConfig field if non-nil, zero value otherwise.

### GetAutoConfigOk

`func (o *UsbLteSettingConfigOpenApiVO) GetAutoConfigOk() (*AutoConfigOpenApiVO, bool)`

GetAutoConfigOk returns a tuple with the AutoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfig

`func (o *UsbLteSettingConfigOpenApiVO) SetAutoConfig(v AutoConfigOpenApiVO)`

SetAutoConfig sets AutoConfig field to given value.

### HasAutoConfig

`func (o *UsbLteSettingConfigOpenApiVO) HasAutoConfig() bool`

HasAutoConfig returns a boolean if a field has been set.

### GetConfigType

`func (o *UsbLteSettingConfigOpenApiVO) GetConfigType() int32`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *UsbLteSettingConfigOpenApiVO) GetConfigTypeOk() (*int32, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *UsbLteSettingConfigOpenApiVO) SetConfigType(v int32)`

SetConfigType sets ConfigType field to given value.


### GetConnectionMode

`func (o *UsbLteSettingConfigOpenApiVO) GetConnectionMode() int32`

GetConnectionMode returns the ConnectionMode field if non-nil, zero value otherwise.

### GetConnectionModeOk

`func (o *UsbLteSettingConfigOpenApiVO) GetConnectionModeOk() (*int32, bool)`

GetConnectionModeOk returns a tuple with the ConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionMode

`func (o *UsbLteSettingConfigOpenApiVO) SetConnectionMode(v int32)`

SetConnectionMode sets ConnectionMode field to given value.


### GetDhcpOptions

`func (o *UsbLteSettingConfigOpenApiVO) GetDhcpOptions() []WanDhcpOptionOpenApiVO`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *UsbLteSettingConfigOpenApiVO) GetDhcpOptionsOk() (*[]WanDhcpOptionOpenApiVO, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *UsbLteSettingConfigOpenApiVO) SetDhcpOptions(v []WanDhcpOptionOpenApiVO)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *UsbLteSettingConfigOpenApiVO) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDnsConfig

`func (o *UsbLteSettingConfigOpenApiVO) GetDnsConfig() DnsConfigOpenApiVO`

GetDnsConfig returns the DnsConfig field if non-nil, zero value otherwise.

### GetDnsConfigOk

`func (o *UsbLteSettingConfigOpenApiVO) GetDnsConfigOk() (*DnsConfigOpenApiVO, bool)`

GetDnsConfigOk returns a tuple with the DnsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfig

`func (o *UsbLteSettingConfigOpenApiVO) SetDnsConfig(v DnsConfigOpenApiVO)`

SetDnsConfig sets DnsConfig field to given value.

### HasDnsConfig

`func (o *UsbLteSettingConfigOpenApiVO) HasDnsConfig() bool`

HasDnsConfig returns a boolean if a field has been set.

### GetDnsEnable

`func (o *UsbLteSettingConfigOpenApiVO) GetDnsEnable() bool`

GetDnsEnable returns the DnsEnable field if non-nil, zero value otherwise.

### GetDnsEnableOk

`func (o *UsbLteSettingConfigOpenApiVO) GetDnsEnableOk() (*bool, bool)`

GetDnsEnableOk returns a tuple with the DnsEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsEnable

`func (o *UsbLteSettingConfigOpenApiVO) SetDnsEnable(v bool)`

SetDnsEnable sets DnsEnable field to given value.


### GetManuallyConfig

`func (o *UsbLteSettingConfigOpenApiVO) GetManuallyConfig() ManuallyConfigOpenApiVO`

GetManuallyConfig returns the ManuallyConfig field if non-nil, zero value otherwise.

### GetManuallyConfigOk

`func (o *UsbLteSettingConfigOpenApiVO) GetManuallyConfigOk() (*ManuallyConfigOpenApiVO, bool)`

GetManuallyConfigOk returns a tuple with the ManuallyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyConfig

`func (o *UsbLteSettingConfigOpenApiVO) SetManuallyConfig(v ManuallyConfigOpenApiVO)`

SetManuallyConfig sets ManuallyConfig field to given value.

### HasManuallyConfig

`func (o *UsbLteSettingConfigOpenApiVO) HasManuallyConfig() bool`

HasManuallyConfig returns a boolean if a field has been set.

### GetMtuSize

`func (o *UsbLteSettingConfigOpenApiVO) GetMtuSize() int32`

GetMtuSize returns the MtuSize field if non-nil, zero value otherwise.

### GetMtuSizeOk

`func (o *UsbLteSettingConfigOpenApiVO) GetMtuSizeOk() (*int32, bool)`

GetMtuSizeOk returns a tuple with the MtuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtuSize

`func (o *UsbLteSettingConfigOpenApiVO) SetMtuSize(v int32)`

SetMtuSize sets MtuSize field to given value.


### GetPin

`func (o *UsbLteSettingConfigOpenApiVO) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *UsbLteSettingConfigOpenApiVO) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *UsbLteSettingConfigOpenApiVO) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *UsbLteSettingConfigOpenApiVO) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetPortDescription

`func (o *UsbLteSettingConfigOpenApiVO) GetPortDescription() string`

GetPortDescription returns the PortDescription field if non-nil, zero value otherwise.

### GetPortDescriptionOk

`func (o *UsbLteSettingConfigOpenApiVO) GetPortDescriptionOk() (*string, bool)`

GetPortDescriptionOk returns a tuple with the PortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDescription

`func (o *UsbLteSettingConfigOpenApiVO) SetPortDescription(v string)`

SetPortDescription sets PortDescription field to given value.

### HasPortDescription

`func (o *UsbLteSettingConfigOpenApiVO) HasPortDescription() bool`

HasPortDescription returns a boolean if a field has been set.

### GetPortId

`func (o *UsbLteSettingConfigOpenApiVO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *UsbLteSettingConfigOpenApiVO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *UsbLteSettingConfigOpenApiVO) SetPortId(v string)`

SetPortId sets PortId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


