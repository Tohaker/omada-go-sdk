# SelectDeviceForVlanTemplateVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedInAdvanced** | Pointer to **bool** | Whether the device is added in advanced. | [optional] 
**Compatible** | Pointer to **int32** | Device firmware and controller compatibility type.Compatible should be a value as follows: 0:COMPATIBLE; 1:HIGH_MAJOR_VER; 2:LOW_MAJOR_VER; 3:HIGH_MINOR_VER; 4:LOW_MINOR_VER; 7:HIGH_COMPONENT_VER; 10:DEVICE_NOT_COMPATIBLE; 11:HIGH_ADOPT_COMMPONENT; 12:DEVICE_CATEGORY_NOT_COMPATIBLE; 14:DEVICE_NOT_COMPATIBLE_IN_CLUSTER | [optional] 
**Lags** | Pointer to [**[]SelectLagForVlanVO**](SelectLagForVlanVO.md) | Switch Lags, only valid when type is 2. | [optional] 
**Mac** | Pointer to **string** | Device mac for Site Setting or Device template ID for Site Template Setting | [optional] 
**Mode** | Pointer to **string** | Device Model. | [optional] 
**ModelVersion** | Pointer to **string** | Device Model Version. | [optional] 
**Name** | Pointer to **string** | Device Name. | [optional] 
**Ports** | Pointer to [**[]SelectPortForVlanVO**](SelectPortForVlanVO.md) | Device Ports | [optional] 
**Type** | Pointer to **string** | Device type should be a value as follows: 1:gateway;  2:switch | [optional] 

## Methods

### NewSelectDeviceForVlanTemplateVO

`func NewSelectDeviceForVlanTemplateVO() *SelectDeviceForVlanTemplateVO`

NewSelectDeviceForVlanTemplateVO instantiates a new SelectDeviceForVlanTemplateVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectDeviceForVlanTemplateVOWithDefaults

`func NewSelectDeviceForVlanTemplateVOWithDefaults() *SelectDeviceForVlanTemplateVO`

NewSelectDeviceForVlanTemplateVOWithDefaults instantiates a new SelectDeviceForVlanTemplateVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedInAdvanced

`func (o *SelectDeviceForVlanTemplateVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *SelectDeviceForVlanTemplateVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *SelectDeviceForVlanTemplateVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *SelectDeviceForVlanTemplateVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetCompatible

`func (o *SelectDeviceForVlanTemplateVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *SelectDeviceForVlanTemplateVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *SelectDeviceForVlanTemplateVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *SelectDeviceForVlanTemplateVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetLags

`func (o *SelectDeviceForVlanTemplateVO) GetLags() []SelectLagForVlanVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *SelectDeviceForVlanTemplateVO) GetLagsOk() (*[]SelectLagForVlanVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *SelectDeviceForVlanTemplateVO) SetLags(v []SelectLagForVlanVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *SelectDeviceForVlanTemplateVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetMac

`func (o *SelectDeviceForVlanTemplateVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SelectDeviceForVlanTemplateVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SelectDeviceForVlanTemplateVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SelectDeviceForVlanTemplateVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMode

`func (o *SelectDeviceForVlanTemplateVO) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SelectDeviceForVlanTemplateVO) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SelectDeviceForVlanTemplateVO) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SelectDeviceForVlanTemplateVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetModelVersion

`func (o *SelectDeviceForVlanTemplateVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *SelectDeviceForVlanTemplateVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *SelectDeviceForVlanTemplateVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *SelectDeviceForVlanTemplateVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *SelectDeviceForVlanTemplateVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelectDeviceForVlanTemplateVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelectDeviceForVlanTemplateVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelectDeviceForVlanTemplateVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *SelectDeviceForVlanTemplateVO) GetPorts() []SelectPortForVlanVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *SelectDeviceForVlanTemplateVO) GetPortsOk() (*[]SelectPortForVlanVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *SelectDeviceForVlanTemplateVO) SetPorts(v []SelectPortForVlanVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *SelectDeviceForVlanTemplateVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetType

`func (o *SelectDeviceForVlanTemplateVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectDeviceForVlanTemplateVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectDeviceForVlanTemplateVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SelectDeviceForVlanTemplateVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


