# BriefDeviceInfoDetailForBatchAdoptOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compatible** | Pointer to **int32** | device compatible | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device type: 0: advanced 1: pro | [optional] 
**Es** | Pointer to **bool** | Whether it is an ES device or not | [optional] 
**HwVersion** | Pointer to **string** | device hardware version | [optional] 
**InWhitelist** | Pointer to **bool** | device inWhitelist | [optional] 
**LicenseStatus** | Pointer to **int32** | license status. 0 means unactive; 1 means unbind; 2 means expired; 3 means active; 4 means incompatible | [optional] 
**Mac** | Pointer to **string** | device mac | [optional] 
**Model** | Pointer to **string** | device model | [optional] 
**ModelVersion** | Pointer to **string** | device model version | [optional] 
**Name** | Pointer to **string** | device name | [optional] 
**OmadacId** | Pointer to **string** | the omadacId of device | [optional] 
**Status** | Pointer to **int32** | device status subcategory: connected， pending， wireless pending | [optional] 
**StatusCategory** | Pointer to **int32** | device status: connected, pending, and disconnected | [optional] 
**Type** | Pointer to **string** | device type | [optional] 
**Wireless** | Pointer to **bool** | Whether the device is wireless | [optional] 

## Methods

### NewBriefDeviceInfoDetailForBatchAdoptOpenApiVO

`func NewBriefDeviceInfoDetailForBatchAdoptOpenApiVO() *BriefDeviceInfoDetailForBatchAdoptOpenApiVO`

NewBriefDeviceInfoDetailForBatchAdoptOpenApiVO instantiates a new BriefDeviceInfoDetailForBatchAdoptOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefDeviceInfoDetailForBatchAdoptOpenApiVOWithDefaults

`func NewBriefDeviceInfoDetailForBatchAdoptOpenApiVOWithDefaults() *BriefDeviceInfoDetailForBatchAdoptOpenApiVO`

NewBriefDeviceInfoDetailForBatchAdoptOpenApiVOWithDefaults instantiates a new BriefDeviceInfoDetailForBatchAdoptOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatible

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetEs

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetEs() bool`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetEsOk() (*bool, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetEs(v bool)`

SetEs sets Es field to given value.

### HasEs

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetHwVersion

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetInWhitelist

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetInWhitelist() bool`

GetInWhitelist returns the InWhitelist field if non-nil, zero value otherwise.

### GetInWhitelistOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetInWhitelistOk() (*bool, bool)`

GetInWhitelistOk returns a tuple with the InWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWhitelist

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetInWhitelist(v bool)`

SetInWhitelist sets InWhitelist field to given value.

### HasInWhitelist

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasInWhitelist() bool`

HasInWhitelist returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetMac

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetStatus

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetType

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWireless

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *BriefDeviceInfoDetailForBatchAdoptOpenApiVO) HasWireless() bool`

HasWireless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


