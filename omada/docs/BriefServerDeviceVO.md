# BriefServerDeviceVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpServerEnable** | Pointer to **bool** | Whether DHCP Server is enabled | [optional] 
**DhcpSettingsAuto** | Pointer to **bool** | Whether DHCP Setting is auto | [optional] 
**Mac** | Pointer to **string** | Dhcp Server Mac | [optional] 
**Model** | Pointer to **string** | Device Model | [optional] 
**ModelVersion** | Pointer to **string** | Device Model Version | [optional] 
**Name** | Pointer to **string** | Dhcp Server Name | [optional] 
**Ranges** | Pointer to [**[]DhcpRangeOpenApiVO**](DhcpRangeOpenApiVO.md) | Dhcp Server Ranges | [optional] 
**StackId** | Pointer to **string** | Dhcp Server Stack ID | [optional] 
**Type** | Pointer to **string** | Device Type | [optional] 

## Methods

### NewBriefServerDeviceVO

`func NewBriefServerDeviceVO() *BriefServerDeviceVO`

NewBriefServerDeviceVO instantiates a new BriefServerDeviceVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefServerDeviceVOWithDefaults

`func NewBriefServerDeviceVOWithDefaults() *BriefServerDeviceVO`

NewBriefServerDeviceVOWithDefaults instantiates a new BriefServerDeviceVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpServerEnable

`func (o *BriefServerDeviceVO) GetDhcpServerEnable() bool`

GetDhcpServerEnable returns the DhcpServerEnable field if non-nil, zero value otherwise.

### GetDhcpServerEnableOk

`func (o *BriefServerDeviceVO) GetDhcpServerEnableOk() (*bool, bool)`

GetDhcpServerEnableOk returns a tuple with the DhcpServerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerEnable

`func (o *BriefServerDeviceVO) SetDhcpServerEnable(v bool)`

SetDhcpServerEnable sets DhcpServerEnable field to given value.

### HasDhcpServerEnable

`func (o *BriefServerDeviceVO) HasDhcpServerEnable() bool`

HasDhcpServerEnable returns a boolean if a field has been set.

### GetDhcpSettingsAuto

`func (o *BriefServerDeviceVO) GetDhcpSettingsAuto() bool`

GetDhcpSettingsAuto returns the DhcpSettingsAuto field if non-nil, zero value otherwise.

### GetDhcpSettingsAutoOk

`func (o *BriefServerDeviceVO) GetDhcpSettingsAutoOk() (*bool, bool)`

GetDhcpSettingsAutoOk returns a tuple with the DhcpSettingsAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSettingsAuto

`func (o *BriefServerDeviceVO) SetDhcpSettingsAuto(v bool)`

SetDhcpSettingsAuto sets DhcpSettingsAuto field to given value.

### HasDhcpSettingsAuto

`func (o *BriefServerDeviceVO) HasDhcpSettingsAuto() bool`

HasDhcpSettingsAuto returns a boolean if a field has been set.

### GetMac

`func (o *BriefServerDeviceVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *BriefServerDeviceVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *BriefServerDeviceVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *BriefServerDeviceVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *BriefServerDeviceVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BriefServerDeviceVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BriefServerDeviceVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BriefServerDeviceVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *BriefServerDeviceVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *BriefServerDeviceVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *BriefServerDeviceVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *BriefServerDeviceVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *BriefServerDeviceVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefServerDeviceVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefServerDeviceVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BriefServerDeviceVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRanges

`func (o *BriefServerDeviceVO) GetRanges() []DhcpRangeOpenApiVO`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *BriefServerDeviceVO) GetRangesOk() (*[]DhcpRangeOpenApiVO, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *BriefServerDeviceVO) SetRanges(v []DhcpRangeOpenApiVO)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *BriefServerDeviceVO) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetStackId

`func (o *BriefServerDeviceVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *BriefServerDeviceVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *BriefServerDeviceVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *BriefServerDeviceVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetType

`func (o *BriefServerDeviceVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BriefServerDeviceVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BriefServerDeviceVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BriefServerDeviceVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


