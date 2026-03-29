# EnableFullChannelDetectApInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedInAdvanced** | Pointer to **bool** | Whether the device is added in advanced. | [optional] 
**Ip** | Pointer to **string** | Ip address,such as 192.168.0.105 | [optional] 
**Mac** | Pointer to **string** | Mac address | [optional] 
**Model** | Pointer to **string** | Model, such as EAP225. | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Default uses the MAC address as the name. | [optional] 
**PowerModeList** | Pointer to **[]int32** | Indicating the power supply modes for AP devices across frequency bands, with array indices corresponding to the frequency bands: 0: 2.4G 1: 5G/5G1 2: 5G2 3: 6G. The corresponding values for power supply modes are: 0: Normal operation 1: Power limited 2: Frequency band disabled | [optional] 
**ScanStatus** | Pointer to **int32** | Scan Status of device,status should be a value as follows:  0:Not Scanned, 1:Spectrum Scanning, 2:RFScanning, 3:packet capturing, 4:RFPlanning; | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**UptimeLong** | Pointer to **int64** | Runtime duration, in seconds (s). | [optional] 
**Version** | Pointer to **string** | Software version, such as \&quot;2.5.0,\&quot; extracted from DeviceDO.firmwareVersion - \&quot;2.5.0 Build 20190118 Rel. 64821.\&quot; | [optional] 
**WirelessLinked** | Pointer to **bool** | Whether the device is wireless linked. | [optional] 

## Methods

### NewEnableFullChannelDetectApInfoOpenApiVO

`func NewEnableFullChannelDetectApInfoOpenApiVO() *EnableFullChannelDetectApInfoOpenApiVO`

NewEnableFullChannelDetectApInfoOpenApiVO instantiates a new EnableFullChannelDetectApInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableFullChannelDetectApInfoOpenApiVOWithDefaults

`func NewEnableFullChannelDetectApInfoOpenApiVOWithDefaults() *EnableFullChannelDetectApInfoOpenApiVO`

NewEnableFullChannelDetectApInfoOpenApiVOWithDefaults instantiates a new EnableFullChannelDetectApInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedInAdvanced

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetIp

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPowerModeList

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetPowerModeList() []int32`

GetPowerModeList returns the PowerModeList field if non-nil, zero value otherwise.

### GetPowerModeListOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetPowerModeListOk() (*[]int32, bool)`

GetPowerModeListOk returns a tuple with the PowerModeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerModeList

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetPowerModeList(v []int32)`

SetPowerModeList sets PowerModeList field to given value.

### HasPowerModeList

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasPowerModeList() bool`

HasPowerModeList returns a boolean if a field has been set.

### GetScanStatus

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetScanStatus() int32`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetScanStatusOk() (*int32, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetScanStatus(v int32)`

SetScanStatus sets ScanStatus field to given value.

### HasScanStatus

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasScanStatus() bool`

HasScanStatus returns a boolean if a field has been set.

### GetStatus

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetType

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptimeLong

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetUptimeLong() int64`

GetUptimeLong returns the UptimeLong field if non-nil, zero value otherwise.

### GetUptimeLongOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetUptimeLongOk() (*int64, bool)`

GetUptimeLongOk returns a tuple with the UptimeLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeLong

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetUptimeLong(v int64)`

SetUptimeLong sets UptimeLong field to given value.

### HasUptimeLong

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasUptimeLong() bool`

HasUptimeLong returns a boolean if a field has been set.

### GetVersion

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWirelessLinked

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetWirelessLinked() bool`

GetWirelessLinked returns the WirelessLinked field if non-nil, zero value otherwise.

### GetWirelessLinkedOk

`func (o *EnableFullChannelDetectApInfoOpenApiVO) GetWirelessLinkedOk() (*bool, bool)`

GetWirelessLinkedOk returns a tuple with the WirelessLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessLinked

`func (o *EnableFullChannelDetectApInfoOpenApiVO) SetWirelessLinked(v bool)`

SetWirelessLinked sets WirelessLinked field to given value.

### HasWirelessLinked

`func (o *EnableFullChannelDetectApInfoOpenApiVO) HasWirelessLinked() bool`

HasWirelessLinked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


