# FailedDeviceUpgradeFirmwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirmwareStatus** | Pointer to **int32** | Target firmware status. 0: normal. 1: The target firmware has been removed/deleted. | [optional] 

## Methods

### NewFailedDeviceUpgradeFirmwareInfo

`func NewFailedDeviceUpgradeFirmwareInfo() *FailedDeviceUpgradeFirmwareInfo`

NewFailedDeviceUpgradeFirmwareInfo instantiates a new FailedDeviceUpgradeFirmwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedDeviceUpgradeFirmwareInfoWithDefaults

`func NewFailedDeviceUpgradeFirmwareInfoWithDefaults() *FailedDeviceUpgradeFirmwareInfo`

NewFailedDeviceUpgradeFirmwareInfoWithDefaults instantiates a new FailedDeviceUpgradeFirmwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirmwareStatus

`func (o *FailedDeviceUpgradeFirmwareInfo) GetFirmwareStatus() int32`

GetFirmwareStatus returns the FirmwareStatus field if non-nil, zero value otherwise.

### GetFirmwareStatusOk

`func (o *FailedDeviceUpgradeFirmwareInfo) GetFirmwareStatusOk() (*int32, bool)`

GetFirmwareStatusOk returns a tuple with the FirmwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareStatus

`func (o *FailedDeviceUpgradeFirmwareInfo) SetFirmwareStatus(v int32)`

SetFirmwareStatus sets FirmwareStatus field to given value.

### HasFirmwareStatus

`func (o *FailedDeviceUpgradeFirmwareInfo) HasFirmwareStatus() bool`

HasFirmwareStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


