# DeviceBriefVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **string** | Only valid when the device is client. It indicates the device type. | [optional] 
**Ip** | Pointer to **string** | IP | [optional] 
**IsClient** | Pointer to **bool** | It indicates whether the device is client. | [optional] 
**Mac** | Pointer to **string** | Mac | [optional] 
**Model** | Pointer to **string** | Model | [optional] 
**ModelVersion** | Pointer to **string** | Model version | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**StatusCategory** | Pointer to **int32** | Device status category, 0: Disconnected, 1: Connected, 2: Pending,3: Heartbeat Missed, 4: Isolated | [optional] 
**Type** | Pointer to **string** | Type | [optional] 

## Methods

### NewDeviceBriefVO

`func NewDeviceBriefVO() *DeviceBriefVO`

NewDeviceBriefVO instantiates a new DeviceBriefVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceBriefVOWithDefaults

`func NewDeviceBriefVOWithDefaults() *DeviceBriefVO`

NewDeviceBriefVOWithDefaults instantiates a new DeviceBriefVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *DeviceBriefVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceBriefVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceBriefVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DeviceBriefVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetIp

`func (o *DeviceBriefVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DeviceBriefVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DeviceBriefVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DeviceBriefVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsClient

`func (o *DeviceBriefVO) GetIsClient() bool`

GetIsClient returns the IsClient field if non-nil, zero value otherwise.

### GetIsClientOk

`func (o *DeviceBriefVO) GetIsClientOk() (*bool, bool)`

GetIsClientOk returns a tuple with the IsClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClient

`func (o *DeviceBriefVO) SetIsClient(v bool)`

SetIsClient sets IsClient field to given value.

### HasIsClient

`func (o *DeviceBriefVO) HasIsClient() bool`

HasIsClient returns a boolean if a field has been set.

### GetMac

`func (o *DeviceBriefVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceBriefVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceBriefVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceBriefVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *DeviceBriefVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceBriefVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceBriefVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceBriefVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *DeviceBriefVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *DeviceBriefVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *DeviceBriefVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *DeviceBriefVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *DeviceBriefVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceBriefVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceBriefVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceBriefVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatusCategory

`func (o *DeviceBriefVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *DeviceBriefVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *DeviceBriefVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *DeviceBriefVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetType

`func (o *DeviceBriefVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceBriefVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceBriefVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceBriefVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


