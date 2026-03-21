# ControllerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaInfo** | Pointer to **string** | Storage area information | [optional] 
**Category** | Pointer to **string** | Category should be a value as follows: localPro; local; advanced; pro. | [optional] 
**ControllerType** | Pointer to **int32** | Controller type should be a value as follows: 0: Windows; 1: Linux; 10: OC200; 11: OC300; 12: ER7212PC; 13: OC400; 20: CBC. | [optional] 
**ControllerVersion** | Pointer to **string** | Controller version | [optional] 
**DeviceCapacity** | Pointer to [**DeviceCapacity**](DeviceCapacity.md) |  | [optional] 
**FirmwareVersion** | Pointer to **string** | Firmware Version | [optional] 
**HwcStorage** | Pointer to [**[]DiskCondition**](DiskCondition.md) | Hardware Storage | [optional] 
**InformUrl** | Pointer to **string** | Inform URL | [optional] 
**MacAddress** | Pointer to **string** | MAC address, should be a valid MAC address format, e.g. AA-BB-CC-DD-11-22 | [optional] 
**Model** | Pointer to **string** | Controller Model | [optional] 
**Name** | Pointer to **string** | Controller Name | [optional] 
**Sn** | Pointer to **string** | Device serial number | [optional] 
**SystemTime** | Pointer to **int64** | System time | [optional] 
**TimeZone** | Pointer to **string** | For the values of timeZone, refer to section 5.1 of the Open API Access Guide. | [optional] 
**UpTime** | Pointer to **int64** | Up time (unit: s) | [optional] 

## Methods

### NewControllerStatus

`func NewControllerStatus() *ControllerStatus`

NewControllerStatus instantiates a new ControllerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerStatusWithDefaults

`func NewControllerStatusWithDefaults() *ControllerStatus`

NewControllerStatusWithDefaults instantiates a new ControllerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaInfo

`func (o *ControllerStatus) GetAreaInfo() string`

GetAreaInfo returns the AreaInfo field if non-nil, zero value otherwise.

### GetAreaInfoOk

`func (o *ControllerStatus) GetAreaInfoOk() (*string, bool)`

GetAreaInfoOk returns a tuple with the AreaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaInfo

`func (o *ControllerStatus) SetAreaInfo(v string)`

SetAreaInfo sets AreaInfo field to given value.

### HasAreaInfo

`func (o *ControllerStatus) HasAreaInfo() bool`

HasAreaInfo returns a boolean if a field has been set.

### GetCategory

`func (o *ControllerStatus) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ControllerStatus) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ControllerStatus) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ControllerStatus) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetControllerType

`func (o *ControllerStatus) GetControllerType() int32`

GetControllerType returns the ControllerType field if non-nil, zero value otherwise.

### GetControllerTypeOk

`func (o *ControllerStatus) GetControllerTypeOk() (*int32, bool)`

GetControllerTypeOk returns a tuple with the ControllerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerType

`func (o *ControllerStatus) SetControllerType(v int32)`

SetControllerType sets ControllerType field to given value.

### HasControllerType

`func (o *ControllerStatus) HasControllerType() bool`

HasControllerType returns a boolean if a field has been set.

### GetControllerVersion

`func (o *ControllerStatus) GetControllerVersion() string`

GetControllerVersion returns the ControllerVersion field if non-nil, zero value otherwise.

### GetControllerVersionOk

`func (o *ControllerStatus) GetControllerVersionOk() (*string, bool)`

GetControllerVersionOk returns a tuple with the ControllerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVersion

`func (o *ControllerStatus) SetControllerVersion(v string)`

SetControllerVersion sets ControllerVersion field to given value.

### HasControllerVersion

`func (o *ControllerStatus) HasControllerVersion() bool`

HasControllerVersion returns a boolean if a field has been set.

### GetDeviceCapacity

`func (o *ControllerStatus) GetDeviceCapacity() DeviceCapacity`

GetDeviceCapacity returns the DeviceCapacity field if non-nil, zero value otherwise.

### GetDeviceCapacityOk

`func (o *ControllerStatus) GetDeviceCapacityOk() (*DeviceCapacity, bool)`

GetDeviceCapacityOk returns a tuple with the DeviceCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCapacity

`func (o *ControllerStatus) SetDeviceCapacity(v DeviceCapacity)`

SetDeviceCapacity sets DeviceCapacity field to given value.

### HasDeviceCapacity

`func (o *ControllerStatus) HasDeviceCapacity() bool`

HasDeviceCapacity returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *ControllerStatus) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *ControllerStatus) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *ControllerStatus) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *ControllerStatus) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetHwcStorage

`func (o *ControllerStatus) GetHwcStorage() []DiskCondition`

GetHwcStorage returns the HwcStorage field if non-nil, zero value otherwise.

### GetHwcStorageOk

`func (o *ControllerStatus) GetHwcStorageOk() (*[]DiskCondition, bool)`

GetHwcStorageOk returns a tuple with the HwcStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwcStorage

`func (o *ControllerStatus) SetHwcStorage(v []DiskCondition)`

SetHwcStorage sets HwcStorage field to given value.

### HasHwcStorage

`func (o *ControllerStatus) HasHwcStorage() bool`

HasHwcStorage returns a boolean if a field has been set.

### GetInformUrl

`func (o *ControllerStatus) GetInformUrl() string`

GetInformUrl returns the InformUrl field if non-nil, zero value otherwise.

### GetInformUrlOk

`func (o *ControllerStatus) GetInformUrlOk() (*string, bool)`

GetInformUrlOk returns a tuple with the InformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformUrl

`func (o *ControllerStatus) SetInformUrl(v string)`

SetInformUrl sets InformUrl field to given value.

### HasInformUrl

`func (o *ControllerStatus) HasInformUrl() bool`

HasInformUrl returns a boolean if a field has been set.

### GetMacAddress

`func (o *ControllerStatus) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ControllerStatus) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ControllerStatus) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ControllerStatus) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetModel

`func (o *ControllerStatus) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ControllerStatus) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ControllerStatus) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ControllerStatus) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *ControllerStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllerStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllerStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllerStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSn

`func (o *ControllerStatus) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *ControllerStatus) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *ControllerStatus) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *ControllerStatus) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetSystemTime

`func (o *ControllerStatus) GetSystemTime() int64`

GetSystemTime returns the SystemTime field if non-nil, zero value otherwise.

### GetSystemTimeOk

`func (o *ControllerStatus) GetSystemTimeOk() (*int64, bool)`

GetSystemTimeOk returns a tuple with the SystemTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTime

`func (o *ControllerStatus) SetSystemTime(v int64)`

SetSystemTime sets SystemTime field to given value.

### HasSystemTime

`func (o *ControllerStatus) HasSystemTime() bool`

HasSystemTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *ControllerStatus) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ControllerStatus) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ControllerStatus) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *ControllerStatus) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUpTime

`func (o *ControllerStatus) GetUpTime() int64`

GetUpTime returns the UpTime field if non-nil, zero value otherwise.

### GetUpTimeOk

`func (o *ControllerStatus) GetUpTimeOk() (*int64, bool)`

GetUpTimeOk returns a tuple with the UpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTime

`func (o *ControllerStatus) SetUpTime(v int64)`

SetUpTime sets UpTime field to given value.

### HasUpTime

`func (o *ControllerStatus) HasUpTime() bool`

HasUpTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


