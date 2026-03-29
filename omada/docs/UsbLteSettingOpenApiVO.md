# UsbLteSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **int32** | 0:Auto, 1:PAP, 2:CHAP | 
**AutoConfig** | Pointer to [**AutoConfigOpenApiVO**](AutoConfigOpenApiVO.md) |  | [optional] 
**ConfigType** | **int32** | 0: Auto; 1：Manually. | 
**ConnectionMode** | **int32** | 1: Connect Automatically, 2: Connect Manually | 
**DhcpOptions** | Pointer to [**[]WanDhcpOptionOpenApiVO**](WanDhcpOptionOpenApiVO.md) |  | [optional] 
**DnsConfig** | Pointer to [**DnsConfigOpenApiVO**](DnsConfigOpenApiVO.md) |  | [optional] 
**DnsEnable** | **bool** |  | 
**ManuallyConfig** | Pointer to [**ManuallyConfigOpenApiVO**](ManuallyConfigOpenApiVO.md) |  | [optional] 
**MtuSize** | **int32** | MTU range from 576 ~ 1500 | 
**Pin** | Pointer to **string** | It is required when [usbModemMsgId] is 1 or 3 | [optional] 
**PortDescription** | Pointer to **string** | Port description | [optional] 
**PortId** | **string** | Port ID | 
**PortName** | Pointer to **string** | Port Name | [optional] 
**ReduceUsbRif** | Pointer to **bool** | whether to use usb2.0, false means close, default is open | [optional] 
**UsbModem** | Pointer to **string** | USB modem name | [optional] 
**UsbModemMsgId** | Pointer to **int32** |  | [optional] 

## Methods

### NewUsbLteSettingOpenApiVO

`func NewUsbLteSettingOpenApiVO(authType int32, configType int32, connectionMode int32, dnsEnable bool, mtuSize int32, portId string, ) *UsbLteSettingOpenApiVO`

NewUsbLteSettingOpenApiVO instantiates a new UsbLteSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsbLteSettingOpenApiVOWithDefaults

`func NewUsbLteSettingOpenApiVOWithDefaults() *UsbLteSettingOpenApiVO`

NewUsbLteSettingOpenApiVOWithDefaults instantiates a new UsbLteSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *UsbLteSettingOpenApiVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *UsbLteSettingOpenApiVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *UsbLteSettingOpenApiVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.


### GetAutoConfig

`func (o *UsbLteSettingOpenApiVO) GetAutoConfig() AutoConfigOpenApiVO`

GetAutoConfig returns the AutoConfig field if non-nil, zero value otherwise.

### GetAutoConfigOk

`func (o *UsbLteSettingOpenApiVO) GetAutoConfigOk() (*AutoConfigOpenApiVO, bool)`

GetAutoConfigOk returns a tuple with the AutoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfig

`func (o *UsbLteSettingOpenApiVO) SetAutoConfig(v AutoConfigOpenApiVO)`

SetAutoConfig sets AutoConfig field to given value.

### HasAutoConfig

`func (o *UsbLteSettingOpenApiVO) HasAutoConfig() bool`

HasAutoConfig returns a boolean if a field has been set.

### GetConfigType

`func (o *UsbLteSettingOpenApiVO) GetConfigType() int32`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *UsbLteSettingOpenApiVO) GetConfigTypeOk() (*int32, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *UsbLteSettingOpenApiVO) SetConfigType(v int32)`

SetConfigType sets ConfigType field to given value.


### GetConnectionMode

`func (o *UsbLteSettingOpenApiVO) GetConnectionMode() int32`

GetConnectionMode returns the ConnectionMode field if non-nil, zero value otherwise.

### GetConnectionModeOk

`func (o *UsbLteSettingOpenApiVO) GetConnectionModeOk() (*int32, bool)`

GetConnectionModeOk returns a tuple with the ConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionMode

`func (o *UsbLteSettingOpenApiVO) SetConnectionMode(v int32)`

SetConnectionMode sets ConnectionMode field to given value.


### GetDhcpOptions

`func (o *UsbLteSettingOpenApiVO) GetDhcpOptions() []WanDhcpOptionOpenApiVO`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *UsbLteSettingOpenApiVO) GetDhcpOptionsOk() (*[]WanDhcpOptionOpenApiVO, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *UsbLteSettingOpenApiVO) SetDhcpOptions(v []WanDhcpOptionOpenApiVO)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *UsbLteSettingOpenApiVO) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDnsConfig

`func (o *UsbLteSettingOpenApiVO) GetDnsConfig() DnsConfigOpenApiVO`

GetDnsConfig returns the DnsConfig field if non-nil, zero value otherwise.

### GetDnsConfigOk

`func (o *UsbLteSettingOpenApiVO) GetDnsConfigOk() (*DnsConfigOpenApiVO, bool)`

GetDnsConfigOk returns a tuple with the DnsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfig

`func (o *UsbLteSettingOpenApiVO) SetDnsConfig(v DnsConfigOpenApiVO)`

SetDnsConfig sets DnsConfig field to given value.

### HasDnsConfig

`func (o *UsbLteSettingOpenApiVO) HasDnsConfig() bool`

HasDnsConfig returns a boolean if a field has been set.

### GetDnsEnable

`func (o *UsbLteSettingOpenApiVO) GetDnsEnable() bool`

GetDnsEnable returns the DnsEnable field if non-nil, zero value otherwise.

### GetDnsEnableOk

`func (o *UsbLteSettingOpenApiVO) GetDnsEnableOk() (*bool, bool)`

GetDnsEnableOk returns a tuple with the DnsEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsEnable

`func (o *UsbLteSettingOpenApiVO) SetDnsEnable(v bool)`

SetDnsEnable sets DnsEnable field to given value.


### GetManuallyConfig

`func (o *UsbLteSettingOpenApiVO) GetManuallyConfig() ManuallyConfigOpenApiVO`

GetManuallyConfig returns the ManuallyConfig field if non-nil, zero value otherwise.

### GetManuallyConfigOk

`func (o *UsbLteSettingOpenApiVO) GetManuallyConfigOk() (*ManuallyConfigOpenApiVO, bool)`

GetManuallyConfigOk returns a tuple with the ManuallyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyConfig

`func (o *UsbLteSettingOpenApiVO) SetManuallyConfig(v ManuallyConfigOpenApiVO)`

SetManuallyConfig sets ManuallyConfig field to given value.

### HasManuallyConfig

`func (o *UsbLteSettingOpenApiVO) HasManuallyConfig() bool`

HasManuallyConfig returns a boolean if a field has been set.

### GetMtuSize

`func (o *UsbLteSettingOpenApiVO) GetMtuSize() int32`

GetMtuSize returns the MtuSize field if non-nil, zero value otherwise.

### GetMtuSizeOk

`func (o *UsbLteSettingOpenApiVO) GetMtuSizeOk() (*int32, bool)`

GetMtuSizeOk returns a tuple with the MtuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtuSize

`func (o *UsbLteSettingOpenApiVO) SetMtuSize(v int32)`

SetMtuSize sets MtuSize field to given value.


### GetPin

`func (o *UsbLteSettingOpenApiVO) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *UsbLteSettingOpenApiVO) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *UsbLteSettingOpenApiVO) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *UsbLteSettingOpenApiVO) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetPortDescription

`func (o *UsbLteSettingOpenApiVO) GetPortDescription() string`

GetPortDescription returns the PortDescription field if non-nil, zero value otherwise.

### GetPortDescriptionOk

`func (o *UsbLteSettingOpenApiVO) GetPortDescriptionOk() (*string, bool)`

GetPortDescriptionOk returns a tuple with the PortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDescription

`func (o *UsbLteSettingOpenApiVO) SetPortDescription(v string)`

SetPortDescription sets PortDescription field to given value.

### HasPortDescription

`func (o *UsbLteSettingOpenApiVO) HasPortDescription() bool`

HasPortDescription returns a boolean if a field has been set.

### GetPortId

`func (o *UsbLteSettingOpenApiVO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *UsbLteSettingOpenApiVO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *UsbLteSettingOpenApiVO) SetPortId(v string)`

SetPortId sets PortId field to given value.


### GetPortName

`func (o *UsbLteSettingOpenApiVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *UsbLteSettingOpenApiVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *UsbLteSettingOpenApiVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *UsbLteSettingOpenApiVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetReduceUsbRif

`func (o *UsbLteSettingOpenApiVO) GetReduceUsbRif() bool`

GetReduceUsbRif returns the ReduceUsbRif field if non-nil, zero value otherwise.

### GetReduceUsbRifOk

`func (o *UsbLteSettingOpenApiVO) GetReduceUsbRifOk() (*bool, bool)`

GetReduceUsbRifOk returns a tuple with the ReduceUsbRif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceUsbRif

`func (o *UsbLteSettingOpenApiVO) SetReduceUsbRif(v bool)`

SetReduceUsbRif sets ReduceUsbRif field to given value.

### HasReduceUsbRif

`func (o *UsbLteSettingOpenApiVO) HasReduceUsbRif() bool`

HasReduceUsbRif returns a boolean if a field has been set.

### GetUsbModem

`func (o *UsbLteSettingOpenApiVO) GetUsbModem() string`

GetUsbModem returns the UsbModem field if non-nil, zero value otherwise.

### GetUsbModemOk

`func (o *UsbLteSettingOpenApiVO) GetUsbModemOk() (*string, bool)`

GetUsbModemOk returns a tuple with the UsbModem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbModem

`func (o *UsbLteSettingOpenApiVO) SetUsbModem(v string)`

SetUsbModem sets UsbModem field to given value.

### HasUsbModem

`func (o *UsbLteSettingOpenApiVO) HasUsbModem() bool`

HasUsbModem returns a boolean if a field has been set.

### GetUsbModemMsgId

`func (o *UsbLteSettingOpenApiVO) GetUsbModemMsgId() int32`

GetUsbModemMsgId returns the UsbModemMsgId field if non-nil, zero value otherwise.

### GetUsbModemMsgIdOk

`func (o *UsbLteSettingOpenApiVO) GetUsbModemMsgIdOk() (*int32, bool)`

GetUsbModemMsgIdOk returns a tuple with the UsbModemMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbModemMsgId

`func (o *UsbLteSettingOpenApiVO) SetUsbModemMsgId(v int32)`

SetUsbModemMsgId sets UsbModemMsgId field to given value.

### HasUsbModemMsgId

`func (o *UsbLteSettingOpenApiVO) HasUsbModemMsgId() bool`

HasUsbModemMsgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


