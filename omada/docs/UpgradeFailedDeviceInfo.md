# UpgradeFailedDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to **string** | The current firmware version of the device which upgrade failed. | [optional] 
**DeviceName** | Pointer to **string** | Name of the device which failed to upgrade. | [optional] 
**Mac** | Pointer to **string** | Mac address of device which upgrade failed. | [optional] 
**ModelTypeInfo** | Pointer to [**ModelTypeInfoOpenApiVO**](ModelTypeInfoOpenApiVO.md) |  | [optional] 
**SiteName** | Pointer to **string** | The siteName of the device which upgrade failed. | [optional] 
**Type** | Pointer to **string** | Device type should be a value as follows: gateway, switch, ap. | [optional] 

## Methods

### NewUpgradeFailedDeviceInfo

`func NewUpgradeFailedDeviceInfo() *UpgradeFailedDeviceInfo`

NewUpgradeFailedDeviceInfo instantiates a new UpgradeFailedDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeFailedDeviceInfoWithDefaults

`func NewUpgradeFailedDeviceInfoWithDefaults() *UpgradeFailedDeviceInfo`

NewUpgradeFailedDeviceInfoWithDefaults instantiates a new UpgradeFailedDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *UpgradeFailedDeviceInfo) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *UpgradeFailedDeviceInfo) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *UpgradeFailedDeviceInfo) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *UpgradeFailedDeviceInfo) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetDeviceName

`func (o *UpgradeFailedDeviceInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *UpgradeFailedDeviceInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *UpgradeFailedDeviceInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *UpgradeFailedDeviceInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetMac

`func (o *UpgradeFailedDeviceInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *UpgradeFailedDeviceInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *UpgradeFailedDeviceInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *UpgradeFailedDeviceInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModelTypeInfo

`func (o *UpgradeFailedDeviceInfo) GetModelTypeInfo() ModelTypeInfoOpenApiVO`

GetModelTypeInfo returns the ModelTypeInfo field if non-nil, zero value otherwise.

### GetModelTypeInfoOk

`func (o *UpgradeFailedDeviceInfo) GetModelTypeInfoOk() (*ModelTypeInfoOpenApiVO, bool)`

GetModelTypeInfoOk returns a tuple with the ModelTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeInfo

`func (o *UpgradeFailedDeviceInfo) SetModelTypeInfo(v ModelTypeInfoOpenApiVO)`

SetModelTypeInfo sets ModelTypeInfo field to given value.

### HasModelTypeInfo

`func (o *UpgradeFailedDeviceInfo) HasModelTypeInfo() bool`

HasModelTypeInfo returns a boolean if a field has been set.

### GetSiteName

`func (o *UpgradeFailedDeviceInfo) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *UpgradeFailedDeviceInfo) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *UpgradeFailedDeviceInfo) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *UpgradeFailedDeviceInfo) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetType

`func (o *UpgradeFailedDeviceInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpgradeFailedDeviceInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpgradeFailedDeviceInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpgradeFailedDeviceInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


