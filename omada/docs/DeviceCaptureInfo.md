# DeviceCaptureInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelList2g** | Pointer to [**[]ChannelInfo**](ChannelInfo.md) | Channel list that device supports in 2.4 GHz. | [optional] 
**ChannelList5g** | Pointer to [**[]ChannelInfo**](ChannelInfo.md) | Channel list that device supports in 5 GHz. | [optional] 
**ChannelList5g2** | Pointer to [**[]ChannelInfo**](ChannelInfo.md) | Channel list that device supports in 5 GHz(2). | [optional] 
**ChannelList6g** | Pointer to [**[]ChannelInfo**](ChannelInfo.md) | Channel list that device supports in 6 GHz. | [optional] 
**CurrentChannel2g** | Pointer to **int32** | Current channel of 2.4 GHz. | [optional] 
**CurrentChannel5g** | Pointer to **int32** | Current channel of 5 GHz. | [optional] 
**CurrentChannel5g2** | Pointer to **int32** | Current channel of 5 GHz(2). | [optional] 
**CurrentChannel6g** | Pointer to **int32** | Current channel of 6 GHz. | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device type: 0: Advanced; 1: Pro. | [optional] 
**FileSize** | Pointer to **int32** | Size limit of the captured file, in MB. | [optional] 
**FirmwareVersion** | Pointer to **string** | Device firmware version. | [optional] 
**Interfaces** | Pointer to [**[]InterfaceInfo**](InterfaceInfo.md) | Interface info. | [optional] 
**Ip** | Pointer to **string** | IP Address. | [optional] 
**Mac** | Pointer to **string** | Device MAC. | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225. | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Device name. | [optional] 
**ShowModel** | Pointer to **string** | Model complex shown in the front end.Ap：model+(country)+modelVersion,EAP225(EU) v3.0  Gateway/Switch：model+modelVersion,Osg v3.0 | [optional] 
**Stack** | Pointer to **bool** | Parameter [stack] indicates whether the device supports stacking. | [optional] 
**StackId** | Pointer to **string** | Stack ID. | [optional] 
**Support2g** | Pointer to **bool** | Whether the device supports 2.4GHz. | [optional] 
**Support5g** | Pointer to **bool** | Whether the device supports 5GHz. | [optional] 
**Support5g2** | Pointer to **bool** | Whether the device supports 5GHz frequency splitting into two parts. | [optional] 
**Support6g** | Pointer to **bool** | Whether the device supports 6GHz. | [optional] 
**SupportCapture** | Pointer to **bool** | Parameter [supportCapture] indicates whether the device supports packet capture. | [optional] 
**SupportLanCapture** | Pointer to **bool** | Parameter [supportLanCapture] indicates whether the device supports LAN port packet capture. | [optional] 
**SupportOTACapture** | Pointer to **bool** | Parameter [supportOTACapture] indicates whether the device supports air interface packet capture and flow-mode packet capture. | [optional] 
**SupportScanRadio** | Pointer to **bool** | Parameter [supportScanRadio] indicates whether the device has independent radio frequency capabilities. | [optional] 
**SupportStreamCapture** | Pointer to **bool** | Parameter [supportStreamCapture] indicates whether the device supports flow-mode packet capture. | [optional] 
**Type** | Pointer to **string** | Device Type. | [optional] 

## Methods

### NewDeviceCaptureInfo

`func NewDeviceCaptureInfo() *DeviceCaptureInfo`

NewDeviceCaptureInfo instantiates a new DeviceCaptureInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCaptureInfoWithDefaults

`func NewDeviceCaptureInfoWithDefaults() *DeviceCaptureInfo`

NewDeviceCaptureInfoWithDefaults instantiates a new DeviceCaptureInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelList2g

`func (o *DeviceCaptureInfo) GetChannelList2g() []ChannelInfo`

GetChannelList2g returns the ChannelList2g field if non-nil, zero value otherwise.

### GetChannelList2gOk

`func (o *DeviceCaptureInfo) GetChannelList2gOk() (*[]ChannelInfo, bool)`

GetChannelList2gOk returns a tuple with the ChannelList2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelList2g

`func (o *DeviceCaptureInfo) SetChannelList2g(v []ChannelInfo)`

SetChannelList2g sets ChannelList2g field to given value.

### HasChannelList2g

`func (o *DeviceCaptureInfo) HasChannelList2g() bool`

HasChannelList2g returns a boolean if a field has been set.

### GetChannelList5g

`func (o *DeviceCaptureInfo) GetChannelList5g() []ChannelInfo`

GetChannelList5g returns the ChannelList5g field if non-nil, zero value otherwise.

### GetChannelList5gOk

`func (o *DeviceCaptureInfo) GetChannelList5gOk() (*[]ChannelInfo, bool)`

GetChannelList5gOk returns a tuple with the ChannelList5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelList5g

`func (o *DeviceCaptureInfo) SetChannelList5g(v []ChannelInfo)`

SetChannelList5g sets ChannelList5g field to given value.

### HasChannelList5g

`func (o *DeviceCaptureInfo) HasChannelList5g() bool`

HasChannelList5g returns a boolean if a field has been set.

### GetChannelList5g2

`func (o *DeviceCaptureInfo) GetChannelList5g2() []ChannelInfo`

GetChannelList5g2 returns the ChannelList5g2 field if non-nil, zero value otherwise.

### GetChannelList5g2Ok

`func (o *DeviceCaptureInfo) GetChannelList5g2Ok() (*[]ChannelInfo, bool)`

GetChannelList5g2Ok returns a tuple with the ChannelList5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelList5g2

`func (o *DeviceCaptureInfo) SetChannelList5g2(v []ChannelInfo)`

SetChannelList5g2 sets ChannelList5g2 field to given value.

### HasChannelList5g2

`func (o *DeviceCaptureInfo) HasChannelList5g2() bool`

HasChannelList5g2 returns a boolean if a field has been set.

### GetChannelList6g

`func (o *DeviceCaptureInfo) GetChannelList6g() []ChannelInfo`

GetChannelList6g returns the ChannelList6g field if non-nil, zero value otherwise.

### GetChannelList6gOk

`func (o *DeviceCaptureInfo) GetChannelList6gOk() (*[]ChannelInfo, bool)`

GetChannelList6gOk returns a tuple with the ChannelList6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelList6g

`func (o *DeviceCaptureInfo) SetChannelList6g(v []ChannelInfo)`

SetChannelList6g sets ChannelList6g field to given value.

### HasChannelList6g

`func (o *DeviceCaptureInfo) HasChannelList6g() bool`

HasChannelList6g returns a boolean if a field has been set.

### GetCurrentChannel2g

`func (o *DeviceCaptureInfo) GetCurrentChannel2g() int32`

GetCurrentChannel2g returns the CurrentChannel2g field if non-nil, zero value otherwise.

### GetCurrentChannel2gOk

`func (o *DeviceCaptureInfo) GetCurrentChannel2gOk() (*int32, bool)`

GetCurrentChannel2gOk returns a tuple with the CurrentChannel2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChannel2g

`func (o *DeviceCaptureInfo) SetCurrentChannel2g(v int32)`

SetCurrentChannel2g sets CurrentChannel2g field to given value.

### HasCurrentChannel2g

`func (o *DeviceCaptureInfo) HasCurrentChannel2g() bool`

HasCurrentChannel2g returns a boolean if a field has been set.

### GetCurrentChannel5g

`func (o *DeviceCaptureInfo) GetCurrentChannel5g() int32`

GetCurrentChannel5g returns the CurrentChannel5g field if non-nil, zero value otherwise.

### GetCurrentChannel5gOk

`func (o *DeviceCaptureInfo) GetCurrentChannel5gOk() (*int32, bool)`

GetCurrentChannel5gOk returns a tuple with the CurrentChannel5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChannel5g

`func (o *DeviceCaptureInfo) SetCurrentChannel5g(v int32)`

SetCurrentChannel5g sets CurrentChannel5g field to given value.

### HasCurrentChannel5g

`func (o *DeviceCaptureInfo) HasCurrentChannel5g() bool`

HasCurrentChannel5g returns a boolean if a field has been set.

### GetCurrentChannel5g2

`func (o *DeviceCaptureInfo) GetCurrentChannel5g2() int32`

GetCurrentChannel5g2 returns the CurrentChannel5g2 field if non-nil, zero value otherwise.

### GetCurrentChannel5g2Ok

`func (o *DeviceCaptureInfo) GetCurrentChannel5g2Ok() (*int32, bool)`

GetCurrentChannel5g2Ok returns a tuple with the CurrentChannel5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChannel5g2

`func (o *DeviceCaptureInfo) SetCurrentChannel5g2(v int32)`

SetCurrentChannel5g2 sets CurrentChannel5g2 field to given value.

### HasCurrentChannel5g2

`func (o *DeviceCaptureInfo) HasCurrentChannel5g2() bool`

HasCurrentChannel5g2 returns a boolean if a field has been set.

### GetCurrentChannel6g

`func (o *DeviceCaptureInfo) GetCurrentChannel6g() int32`

GetCurrentChannel6g returns the CurrentChannel6g field if non-nil, zero value otherwise.

### GetCurrentChannel6gOk

`func (o *DeviceCaptureInfo) GetCurrentChannel6gOk() (*int32, bool)`

GetCurrentChannel6gOk returns a tuple with the CurrentChannel6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChannel6g

`func (o *DeviceCaptureInfo) SetCurrentChannel6g(v int32)`

SetCurrentChannel6g sets CurrentChannel6g field to given value.

### HasCurrentChannel6g

`func (o *DeviceCaptureInfo) HasCurrentChannel6g() bool`

HasCurrentChannel6g returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *DeviceCaptureInfo) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *DeviceCaptureInfo) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *DeviceCaptureInfo) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *DeviceCaptureInfo) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetFileSize

`func (o *DeviceCaptureInfo) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *DeviceCaptureInfo) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *DeviceCaptureInfo) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *DeviceCaptureInfo) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *DeviceCaptureInfo) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *DeviceCaptureInfo) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *DeviceCaptureInfo) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *DeviceCaptureInfo) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetInterfaces

`func (o *DeviceCaptureInfo) GetInterfaces() []InterfaceInfo`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *DeviceCaptureInfo) GetInterfacesOk() (*[]InterfaceInfo, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *DeviceCaptureInfo) SetInterfaces(v []InterfaceInfo)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *DeviceCaptureInfo) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetIp

`func (o *DeviceCaptureInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DeviceCaptureInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DeviceCaptureInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DeviceCaptureInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *DeviceCaptureInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceCaptureInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceCaptureInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceCaptureInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *DeviceCaptureInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceCaptureInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceCaptureInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceCaptureInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *DeviceCaptureInfo) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *DeviceCaptureInfo) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *DeviceCaptureInfo) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *DeviceCaptureInfo) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *DeviceCaptureInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceCaptureInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceCaptureInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceCaptureInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShowModel

`func (o *DeviceCaptureInfo) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *DeviceCaptureInfo) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *DeviceCaptureInfo) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *DeviceCaptureInfo) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStack

`func (o *DeviceCaptureInfo) GetStack() bool`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *DeviceCaptureInfo) GetStackOk() (*bool, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *DeviceCaptureInfo) SetStack(v bool)`

SetStack sets Stack field to given value.

### HasStack

`func (o *DeviceCaptureInfo) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetStackId

`func (o *DeviceCaptureInfo) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *DeviceCaptureInfo) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *DeviceCaptureInfo) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *DeviceCaptureInfo) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetSupport2g

`func (o *DeviceCaptureInfo) GetSupport2g() bool`

GetSupport2g returns the Support2g field if non-nil, zero value otherwise.

### GetSupport2gOk

`func (o *DeviceCaptureInfo) GetSupport2gOk() (*bool, bool)`

GetSupport2gOk returns a tuple with the Support2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport2g

`func (o *DeviceCaptureInfo) SetSupport2g(v bool)`

SetSupport2g sets Support2g field to given value.

### HasSupport2g

`func (o *DeviceCaptureInfo) HasSupport2g() bool`

HasSupport2g returns a boolean if a field has been set.

### GetSupport5g

`func (o *DeviceCaptureInfo) GetSupport5g() bool`

GetSupport5g returns the Support5g field if non-nil, zero value otherwise.

### GetSupport5gOk

`func (o *DeviceCaptureInfo) GetSupport5gOk() (*bool, bool)`

GetSupport5gOk returns a tuple with the Support5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g

`func (o *DeviceCaptureInfo) SetSupport5g(v bool)`

SetSupport5g sets Support5g field to given value.

### HasSupport5g

`func (o *DeviceCaptureInfo) HasSupport5g() bool`

HasSupport5g returns a boolean if a field has been set.

### GetSupport5g2

`func (o *DeviceCaptureInfo) GetSupport5g2() bool`

GetSupport5g2 returns the Support5g2 field if non-nil, zero value otherwise.

### GetSupport5g2Ok

`func (o *DeviceCaptureInfo) GetSupport5g2Ok() (*bool, bool)`

GetSupport5g2Ok returns a tuple with the Support5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g2

`func (o *DeviceCaptureInfo) SetSupport5g2(v bool)`

SetSupport5g2 sets Support5g2 field to given value.

### HasSupport5g2

`func (o *DeviceCaptureInfo) HasSupport5g2() bool`

HasSupport5g2 returns a boolean if a field has been set.

### GetSupport6g

`func (o *DeviceCaptureInfo) GetSupport6g() bool`

GetSupport6g returns the Support6g field if non-nil, zero value otherwise.

### GetSupport6gOk

`func (o *DeviceCaptureInfo) GetSupport6gOk() (*bool, bool)`

GetSupport6gOk returns a tuple with the Support6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport6g

`func (o *DeviceCaptureInfo) SetSupport6g(v bool)`

SetSupport6g sets Support6g field to given value.

### HasSupport6g

`func (o *DeviceCaptureInfo) HasSupport6g() bool`

HasSupport6g returns a boolean if a field has been set.

### GetSupportCapture

`func (o *DeviceCaptureInfo) GetSupportCapture() bool`

GetSupportCapture returns the SupportCapture field if non-nil, zero value otherwise.

### GetSupportCaptureOk

`func (o *DeviceCaptureInfo) GetSupportCaptureOk() (*bool, bool)`

GetSupportCaptureOk returns a tuple with the SupportCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCapture

`func (o *DeviceCaptureInfo) SetSupportCapture(v bool)`

SetSupportCapture sets SupportCapture field to given value.

### HasSupportCapture

`func (o *DeviceCaptureInfo) HasSupportCapture() bool`

HasSupportCapture returns a boolean if a field has been set.

### GetSupportLanCapture

`func (o *DeviceCaptureInfo) GetSupportLanCapture() bool`

GetSupportLanCapture returns the SupportLanCapture field if non-nil, zero value otherwise.

### GetSupportLanCaptureOk

`func (o *DeviceCaptureInfo) GetSupportLanCaptureOk() (*bool, bool)`

GetSupportLanCaptureOk returns a tuple with the SupportLanCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLanCapture

`func (o *DeviceCaptureInfo) SetSupportLanCapture(v bool)`

SetSupportLanCapture sets SupportLanCapture field to given value.

### HasSupportLanCapture

`func (o *DeviceCaptureInfo) HasSupportLanCapture() bool`

HasSupportLanCapture returns a boolean if a field has been set.

### GetSupportOTACapture

`func (o *DeviceCaptureInfo) GetSupportOTACapture() bool`

GetSupportOTACapture returns the SupportOTACapture field if non-nil, zero value otherwise.

### GetSupportOTACaptureOk

`func (o *DeviceCaptureInfo) GetSupportOTACaptureOk() (*bool, bool)`

GetSupportOTACaptureOk returns a tuple with the SupportOTACapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOTACapture

`func (o *DeviceCaptureInfo) SetSupportOTACapture(v bool)`

SetSupportOTACapture sets SupportOTACapture field to given value.

### HasSupportOTACapture

`func (o *DeviceCaptureInfo) HasSupportOTACapture() bool`

HasSupportOTACapture returns a boolean if a field has been set.

### GetSupportScanRadio

`func (o *DeviceCaptureInfo) GetSupportScanRadio() bool`

GetSupportScanRadio returns the SupportScanRadio field if non-nil, zero value otherwise.

### GetSupportScanRadioOk

`func (o *DeviceCaptureInfo) GetSupportScanRadioOk() (*bool, bool)`

GetSupportScanRadioOk returns a tuple with the SupportScanRadio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportScanRadio

`func (o *DeviceCaptureInfo) SetSupportScanRadio(v bool)`

SetSupportScanRadio sets SupportScanRadio field to given value.

### HasSupportScanRadio

`func (o *DeviceCaptureInfo) HasSupportScanRadio() bool`

HasSupportScanRadio returns a boolean if a field has been set.

### GetSupportStreamCapture

`func (o *DeviceCaptureInfo) GetSupportStreamCapture() bool`

GetSupportStreamCapture returns the SupportStreamCapture field if non-nil, zero value otherwise.

### GetSupportStreamCaptureOk

`func (o *DeviceCaptureInfo) GetSupportStreamCaptureOk() (*bool, bool)`

GetSupportStreamCaptureOk returns a tuple with the SupportStreamCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStreamCapture

`func (o *DeviceCaptureInfo) SetSupportStreamCapture(v bool)`

SetSupportStreamCapture sets SupportStreamCapture field to given value.

### HasSupportStreamCapture

`func (o *DeviceCaptureInfo) HasSupportStreamCapture() bool`

HasSupportStreamCapture returns a boolean if a field has been set.

### GetType

`func (o *DeviceCaptureInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceCaptureInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceCaptureInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceCaptureInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


