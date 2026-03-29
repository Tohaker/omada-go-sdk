# PortDeviceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site ID | [optional] 
**HwVersion** | Pointer to **string** | Device hardware version. | [optional] 
**Mac** | Pointer to **string** | Device mac. | [optional] 
**Model** | Pointer to **string** | Device model. | [optional] 
**ModelVersion** | Pointer to **string** | Device model version. | [optional] 
**Name** | Pointer to **string** | Device name. | [optional] 
**Ports** | Pointer to [**DevicePorts**](DevicePorts.md) |  | [optional] 
**Sn** | Pointer to **string** | Device SN code. | [optional] 
**Status** | Pointer to **int32** | Device status. | [optional] 
**Type** | Pointer to **string** | Device type. \&quot;gateway\&quot;, \&quot;switch\&quot; | [optional] 

## Methods

### NewPortDeviceDetail

`func NewPortDeviceDetail() *PortDeviceDetail`

NewPortDeviceDetail instantiates a new PortDeviceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortDeviceDetailWithDefaults

`func NewPortDeviceDetailWithDefaults() *PortDeviceDetail`

NewPortDeviceDetailWithDefaults instantiates a new PortDeviceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *PortDeviceDetail) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PortDeviceDetail) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PortDeviceDetail) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PortDeviceDetail) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetHwVersion

`func (o *PortDeviceDetail) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *PortDeviceDetail) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *PortDeviceDetail) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *PortDeviceDetail) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetMac

`func (o *PortDeviceDetail) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *PortDeviceDetail) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *PortDeviceDetail) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *PortDeviceDetail) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *PortDeviceDetail) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PortDeviceDetail) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PortDeviceDetail) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PortDeviceDetail) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *PortDeviceDetail) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *PortDeviceDetail) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *PortDeviceDetail) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *PortDeviceDetail) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *PortDeviceDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortDeviceDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortDeviceDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PortDeviceDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *PortDeviceDetail) GetPorts() DevicePorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *PortDeviceDetail) GetPortsOk() (*DevicePorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *PortDeviceDetail) SetPorts(v DevicePorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *PortDeviceDetail) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetSn

`func (o *PortDeviceDetail) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *PortDeviceDetail) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *PortDeviceDetail) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *PortDeviceDetail) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetStatus

`func (o *PortDeviceDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortDeviceDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortDeviceDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PortDeviceDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *PortDeviceDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortDeviceDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortDeviceDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PortDeviceDetail) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


