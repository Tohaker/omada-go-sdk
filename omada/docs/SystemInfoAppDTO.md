# SystemInfoAppDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootLoaderVersion** | Pointer to **string** | Boot loader version | [optional] 
**ContactInformation** | Pointer to **string** | Contact information should contain 1-32 bits numbers, Upper and lower letters, -@_:/. . | [optional] 
**DeviceLocation** | Pointer to **string** | Device location should contain 1-32 bits numbers, Upper and lower letters, -@_:/. . | [optional] 
**DeviceName** | Pointer to **string** | Device name should contain 1-32 bits numbers, Upper and lower letters, -@_:/. . | [optional] 
**DhcpRelayStatus** | Pointer to **int32** | Whether DHCP Relaying is enabled, dhcpRelayStatus should be a value as follows: 1:ENABLE;0:DISABLE. | [optional] 
**FirmwareVersion** | Pointer to **string** | Firmware version | [optional] 
**HardwareVersion** | Pointer to **string** | Hardware version | [optional] 
**HttpServerStatus** | Pointer to **int32** | Whether the HTTP Server function is enabled, httpServerStatus should be a value as follows: 1:ENABLE;0:DISABLE. | [optional] 
**IgmpSnoopingStatus** | Pointer to **int32** | Whether IGMP Snooping is enabled, igmpSnoopingStatus should be a value as follows: 1:ENABLE;0:DISABLE. | [optional] 
**JumboFrameStatus** | Pointer to **int32** | Whether Jumbo Frame is enabled, jumboFrameStatus should be a value as follows: 1:ENABLE;0:DISABLE. | [optional] 
**Mac** | Pointer to **string** | Mac address | [optional] 
**RunningTime** | Pointer to **int64** | Uptime, unit (s). | [optional] 
**SerialNumber** | Pointer to **string** | SerialNumber of ONU should contain 12, 13, or 16 characters, in the format &lt;XXXXXXXXXXXX&gt; ,&lt;XXXX-XXXXXXXX&gt;,&lt;XXXXXXXXXXXXXXXX&gt; | [optional] 
**SnmpStatus** | Pointer to **int32** | Whether SNMP is enabled, snmpStatus should be a value as follows: 1:ENABLE;0:DISABLE. | [optional] 
**SntpStatus** | Pointer to **int32** | Whether to obtain time from the NTP server, sntpStatus should be a value as follows:1:ENABLE;0:DISABLE. | [optional] 
**SpanningTreeStatus** | Pointer to **int32** | Whether Spanning Tree is enabled, spanningTreeStatus should be a value as follows: 1:ENABLE;0:DISABLE. | [optional] 
**SshStatus** | Pointer to **int32** | Whether SSH is enabled, sshStatus should be a value as follows: 1:ENABLE;0:DISABLE. | [optional] 
**SystemDescription** | Pointer to **string** | OLT device system description. | [optional] 
**SystemTime** | Pointer to **string** | System time | [optional] 
**TelnetStatus** | Pointer to **int32** | Whether Telnet is enabled, telnetStatus should be a value as follows: 1:ENABLE;0:DISABLE. | [optional] 

## Methods

### NewSystemInfoAppDTO

`func NewSystemInfoAppDTO() *SystemInfoAppDTO`

NewSystemInfoAppDTO instantiates a new SystemInfoAppDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInfoAppDTOWithDefaults

`func NewSystemInfoAppDTOWithDefaults() *SystemInfoAppDTO`

NewSystemInfoAppDTOWithDefaults instantiates a new SystemInfoAppDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootLoaderVersion

`func (o *SystemInfoAppDTO) GetBootLoaderVersion() string`

GetBootLoaderVersion returns the BootLoaderVersion field if non-nil, zero value otherwise.

### GetBootLoaderVersionOk

`func (o *SystemInfoAppDTO) GetBootLoaderVersionOk() (*string, bool)`

GetBootLoaderVersionOk returns a tuple with the BootLoaderVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootLoaderVersion

`func (o *SystemInfoAppDTO) SetBootLoaderVersion(v string)`

SetBootLoaderVersion sets BootLoaderVersion field to given value.

### HasBootLoaderVersion

`func (o *SystemInfoAppDTO) HasBootLoaderVersion() bool`

HasBootLoaderVersion returns a boolean if a field has been set.

### GetContactInformation

`func (o *SystemInfoAppDTO) GetContactInformation() string`

GetContactInformation returns the ContactInformation field if non-nil, zero value otherwise.

### GetContactInformationOk

`func (o *SystemInfoAppDTO) GetContactInformationOk() (*string, bool)`

GetContactInformationOk returns a tuple with the ContactInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInformation

`func (o *SystemInfoAppDTO) SetContactInformation(v string)`

SetContactInformation sets ContactInformation field to given value.

### HasContactInformation

`func (o *SystemInfoAppDTO) HasContactInformation() bool`

HasContactInformation returns a boolean if a field has been set.

### GetDeviceLocation

`func (o *SystemInfoAppDTO) GetDeviceLocation() string`

GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.

### GetDeviceLocationOk

`func (o *SystemInfoAppDTO) GetDeviceLocationOk() (*string, bool)`

GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocation

`func (o *SystemInfoAppDTO) SetDeviceLocation(v string)`

SetDeviceLocation sets DeviceLocation field to given value.

### HasDeviceLocation

`func (o *SystemInfoAppDTO) HasDeviceLocation() bool`

HasDeviceLocation returns a boolean if a field has been set.

### GetDeviceName

`func (o *SystemInfoAppDTO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SystemInfoAppDTO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SystemInfoAppDTO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *SystemInfoAppDTO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDhcpRelayStatus

`func (o *SystemInfoAppDTO) GetDhcpRelayStatus() int32`

GetDhcpRelayStatus returns the DhcpRelayStatus field if non-nil, zero value otherwise.

### GetDhcpRelayStatusOk

`func (o *SystemInfoAppDTO) GetDhcpRelayStatusOk() (*int32, bool)`

GetDhcpRelayStatusOk returns a tuple with the DhcpRelayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelayStatus

`func (o *SystemInfoAppDTO) SetDhcpRelayStatus(v int32)`

SetDhcpRelayStatus sets DhcpRelayStatus field to given value.

### HasDhcpRelayStatus

`func (o *SystemInfoAppDTO) HasDhcpRelayStatus() bool`

HasDhcpRelayStatus returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *SystemInfoAppDTO) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *SystemInfoAppDTO) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *SystemInfoAppDTO) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *SystemInfoAppDTO) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetHardwareVersion

`func (o *SystemInfoAppDTO) GetHardwareVersion() string`

GetHardwareVersion returns the HardwareVersion field if non-nil, zero value otherwise.

### GetHardwareVersionOk

`func (o *SystemInfoAppDTO) GetHardwareVersionOk() (*string, bool)`

GetHardwareVersionOk returns a tuple with the HardwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareVersion

`func (o *SystemInfoAppDTO) SetHardwareVersion(v string)`

SetHardwareVersion sets HardwareVersion field to given value.

### HasHardwareVersion

`func (o *SystemInfoAppDTO) HasHardwareVersion() bool`

HasHardwareVersion returns a boolean if a field has been set.

### GetHttpServerStatus

`func (o *SystemInfoAppDTO) GetHttpServerStatus() int32`

GetHttpServerStatus returns the HttpServerStatus field if non-nil, zero value otherwise.

### GetHttpServerStatusOk

`func (o *SystemInfoAppDTO) GetHttpServerStatusOk() (*int32, bool)`

GetHttpServerStatusOk returns a tuple with the HttpServerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerStatus

`func (o *SystemInfoAppDTO) SetHttpServerStatus(v int32)`

SetHttpServerStatus sets HttpServerStatus field to given value.

### HasHttpServerStatus

`func (o *SystemInfoAppDTO) HasHttpServerStatus() bool`

HasHttpServerStatus returns a boolean if a field has been set.

### GetIgmpSnoopingStatus

`func (o *SystemInfoAppDTO) GetIgmpSnoopingStatus() int32`

GetIgmpSnoopingStatus returns the IgmpSnoopingStatus field if non-nil, zero value otherwise.

### GetIgmpSnoopingStatusOk

`func (o *SystemInfoAppDTO) GetIgmpSnoopingStatusOk() (*int32, bool)`

GetIgmpSnoopingStatusOk returns a tuple with the IgmpSnoopingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopingStatus

`func (o *SystemInfoAppDTO) SetIgmpSnoopingStatus(v int32)`

SetIgmpSnoopingStatus sets IgmpSnoopingStatus field to given value.

### HasIgmpSnoopingStatus

`func (o *SystemInfoAppDTO) HasIgmpSnoopingStatus() bool`

HasIgmpSnoopingStatus returns a boolean if a field has been set.

### GetJumboFrameStatus

`func (o *SystemInfoAppDTO) GetJumboFrameStatus() int32`

GetJumboFrameStatus returns the JumboFrameStatus field if non-nil, zero value otherwise.

### GetJumboFrameStatusOk

`func (o *SystemInfoAppDTO) GetJumboFrameStatusOk() (*int32, bool)`

GetJumboFrameStatusOk returns a tuple with the JumboFrameStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumboFrameStatus

`func (o *SystemInfoAppDTO) SetJumboFrameStatus(v int32)`

SetJumboFrameStatus sets JumboFrameStatus field to given value.

### HasJumboFrameStatus

`func (o *SystemInfoAppDTO) HasJumboFrameStatus() bool`

HasJumboFrameStatus returns a boolean if a field has been set.

### GetMac

`func (o *SystemInfoAppDTO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SystemInfoAppDTO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SystemInfoAppDTO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SystemInfoAppDTO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetRunningTime

`func (o *SystemInfoAppDTO) GetRunningTime() int64`

GetRunningTime returns the RunningTime field if non-nil, zero value otherwise.

### GetRunningTimeOk

`func (o *SystemInfoAppDTO) GetRunningTimeOk() (*int64, bool)`

GetRunningTimeOk returns a tuple with the RunningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningTime

`func (o *SystemInfoAppDTO) SetRunningTime(v int64)`

SetRunningTime sets RunningTime field to given value.

### HasRunningTime

`func (o *SystemInfoAppDTO) HasRunningTime() bool`

HasRunningTime returns a boolean if a field has been set.

### GetSerialNumber

`func (o *SystemInfoAppDTO) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SystemInfoAppDTO) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SystemInfoAppDTO) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SystemInfoAppDTO) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSnmpStatus

`func (o *SystemInfoAppDTO) GetSnmpStatus() int32`

GetSnmpStatus returns the SnmpStatus field if non-nil, zero value otherwise.

### GetSnmpStatusOk

`func (o *SystemInfoAppDTO) GetSnmpStatusOk() (*int32, bool)`

GetSnmpStatusOk returns a tuple with the SnmpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpStatus

`func (o *SystemInfoAppDTO) SetSnmpStatus(v int32)`

SetSnmpStatus sets SnmpStatus field to given value.

### HasSnmpStatus

`func (o *SystemInfoAppDTO) HasSnmpStatus() bool`

HasSnmpStatus returns a boolean if a field has been set.

### GetSntpStatus

`func (o *SystemInfoAppDTO) GetSntpStatus() int32`

GetSntpStatus returns the SntpStatus field if non-nil, zero value otherwise.

### GetSntpStatusOk

`func (o *SystemInfoAppDTO) GetSntpStatusOk() (*int32, bool)`

GetSntpStatusOk returns a tuple with the SntpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSntpStatus

`func (o *SystemInfoAppDTO) SetSntpStatus(v int32)`

SetSntpStatus sets SntpStatus field to given value.

### HasSntpStatus

`func (o *SystemInfoAppDTO) HasSntpStatus() bool`

HasSntpStatus returns a boolean if a field has been set.

### GetSpanningTreeStatus

`func (o *SystemInfoAppDTO) GetSpanningTreeStatus() int32`

GetSpanningTreeStatus returns the SpanningTreeStatus field if non-nil, zero value otherwise.

### GetSpanningTreeStatusOk

`func (o *SystemInfoAppDTO) GetSpanningTreeStatusOk() (*int32, bool)`

GetSpanningTreeStatusOk returns a tuple with the SpanningTreeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTreeStatus

`func (o *SystemInfoAppDTO) SetSpanningTreeStatus(v int32)`

SetSpanningTreeStatus sets SpanningTreeStatus field to given value.

### HasSpanningTreeStatus

`func (o *SystemInfoAppDTO) HasSpanningTreeStatus() bool`

HasSpanningTreeStatus returns a boolean if a field has been set.

### GetSshStatus

`func (o *SystemInfoAppDTO) GetSshStatus() int32`

GetSshStatus returns the SshStatus field if non-nil, zero value otherwise.

### GetSshStatusOk

`func (o *SystemInfoAppDTO) GetSshStatusOk() (*int32, bool)`

GetSshStatusOk returns a tuple with the SshStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshStatus

`func (o *SystemInfoAppDTO) SetSshStatus(v int32)`

SetSshStatus sets SshStatus field to given value.

### HasSshStatus

`func (o *SystemInfoAppDTO) HasSshStatus() bool`

HasSshStatus returns a boolean if a field has been set.

### GetSystemDescription

`func (o *SystemInfoAppDTO) GetSystemDescription() string`

GetSystemDescription returns the SystemDescription field if non-nil, zero value otherwise.

### GetSystemDescriptionOk

`func (o *SystemInfoAppDTO) GetSystemDescriptionOk() (*string, bool)`

GetSystemDescriptionOk returns a tuple with the SystemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDescription

`func (o *SystemInfoAppDTO) SetSystemDescription(v string)`

SetSystemDescription sets SystemDescription field to given value.

### HasSystemDescription

`func (o *SystemInfoAppDTO) HasSystemDescription() bool`

HasSystemDescription returns a boolean if a field has been set.

### GetSystemTime

`func (o *SystemInfoAppDTO) GetSystemTime() string`

GetSystemTime returns the SystemTime field if non-nil, zero value otherwise.

### GetSystemTimeOk

`func (o *SystemInfoAppDTO) GetSystemTimeOk() (*string, bool)`

GetSystemTimeOk returns a tuple with the SystemTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTime

`func (o *SystemInfoAppDTO) SetSystemTime(v string)`

SetSystemTime sets SystemTime field to given value.

### HasSystemTime

`func (o *SystemInfoAppDTO) HasSystemTime() bool`

HasSystemTime returns a boolean if a field has been set.

### GetTelnetStatus

`func (o *SystemInfoAppDTO) GetTelnetStatus() int32`

GetTelnetStatus returns the TelnetStatus field if non-nil, zero value otherwise.

### GetTelnetStatusOk

`func (o *SystemInfoAppDTO) GetTelnetStatusOk() (*int32, bool)`

GetTelnetStatusOk returns a tuple with the TelnetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnetStatus

`func (o *SystemInfoAppDTO) SetTelnetStatus(v int32)`

SetTelnetStatus sets TelnetStatus field to given value.

### HasTelnetStatus

`func (o *SystemInfoAppDTO) HasTelnetStatus() bool`

HasTelnetStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


