# PoeDeviceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site ID | [optional] 
**CombinedGateway** | Pointer to **bool** | Indicates whether it is a combined-gateway. | [optional] 
**FirmwareVersion** | Pointer to **string** | Device firmware version. | [optional] 
**HwVersion** | Pointer to **string** | Device hardware version. | [optional] 
**Mac** | Pointer to **string** | Device MAC. | [optional] 
**Model** | Pointer to **string** | Device model. | [optional] 
**ModelVersion** | Pointer to **string** | Device model version. | [optional] 
**Name** | Pointer to **string** | Device name. | [optional] 
**PoePort** | Pointer to [**DevicePoePorts**](DevicePoePorts.md) |  | [optional] 
**Sn** | Pointer to **string** | Device SN code. | [optional] 
**Status** | Pointer to **int32** | Device status. | [optional] 
**Type** | Pointer to **string** | Device type. \&quot;gateway\&quot;, \&quot;switch\&quot; | [optional] 

## Methods

### NewPoeDeviceDetail

`func NewPoeDeviceDetail() *PoeDeviceDetail`

NewPoeDeviceDetail instantiates a new PoeDeviceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoeDeviceDetailWithDefaults

`func NewPoeDeviceDetailWithDefaults() *PoeDeviceDetail`

NewPoeDeviceDetailWithDefaults instantiates a new PoeDeviceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *PoeDeviceDetail) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PoeDeviceDetail) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PoeDeviceDetail) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PoeDeviceDetail) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetCombinedGateway

`func (o *PoeDeviceDetail) GetCombinedGateway() bool`

GetCombinedGateway returns the CombinedGateway field if non-nil, zero value otherwise.

### GetCombinedGatewayOk

`func (o *PoeDeviceDetail) GetCombinedGatewayOk() (*bool, bool)`

GetCombinedGatewayOk returns a tuple with the CombinedGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedGateway

`func (o *PoeDeviceDetail) SetCombinedGateway(v bool)`

SetCombinedGateway sets CombinedGateway field to given value.

### HasCombinedGateway

`func (o *PoeDeviceDetail) HasCombinedGateway() bool`

HasCombinedGateway returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *PoeDeviceDetail) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *PoeDeviceDetail) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *PoeDeviceDetail) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *PoeDeviceDetail) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetHwVersion

`func (o *PoeDeviceDetail) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *PoeDeviceDetail) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *PoeDeviceDetail) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *PoeDeviceDetail) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetMac

`func (o *PoeDeviceDetail) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *PoeDeviceDetail) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *PoeDeviceDetail) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *PoeDeviceDetail) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *PoeDeviceDetail) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PoeDeviceDetail) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PoeDeviceDetail) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PoeDeviceDetail) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *PoeDeviceDetail) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *PoeDeviceDetail) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *PoeDeviceDetail) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *PoeDeviceDetail) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *PoeDeviceDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoeDeviceDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoeDeviceDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoeDeviceDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoePort

`func (o *PoeDeviceDetail) GetPoePort() DevicePoePorts`

GetPoePort returns the PoePort field if non-nil, zero value otherwise.

### GetPoePortOk

`func (o *PoeDeviceDetail) GetPoePortOk() (*DevicePoePorts, bool)`

GetPoePortOk returns a tuple with the PoePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePort

`func (o *PoeDeviceDetail) SetPoePort(v DevicePoePorts)`

SetPoePort sets PoePort field to given value.

### HasPoePort

`func (o *PoeDeviceDetail) HasPoePort() bool`

HasPoePort returns a boolean if a field has been set.

### GetSn

`func (o *PoeDeviceDetail) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *PoeDeviceDetail) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *PoeDeviceDetail) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *PoeDeviceDetail) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetStatus

`func (o *PoeDeviceDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoeDeviceDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoeDeviceDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PoeDeviceDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *PoeDeviceDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PoeDeviceDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PoeDeviceDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PoeDeviceDetail) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


