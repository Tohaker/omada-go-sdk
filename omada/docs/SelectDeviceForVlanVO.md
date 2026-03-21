# SelectDeviceForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedInAdvanced** | Pointer to **bool** | Whether the device is added in advanced. | [optional] 
**Compatible** | Pointer to **int32** | Device firmware and controller compatibility type.Compatible should be a value as follows: 0:COMPATIBLE;1:HIGH_MAJOR_VER;2:LOW_MAJOR_VER;3:HIGH_MINOR_VER;4:LOW_MINOR_VER;7:HIGH_COMPONENT_VER;10:DEVICE_NOT_COMPATIBLE;11:HIGH_ADOPT_COMMPONENT;12:DEVICE_CATEGORY_NOT_COMPATIBLE;14:DEVICE_NOT_COMPATIBLE_IN_CLUSTER | [optional] 
**Es** | Pointer to **bool** | Whether the switch is Agile Series Switch | [optional] 
**EsAllowAllPortExceed** | Pointer to **bool** | It indicates whether there exists allow all port of Agile Series Switch that can only use partly of currently creating vlans | [optional] 
**Lags** | Pointer to [**[]SelectLagForVlanVO**](SelectLagForVlanVO.md) | Switch Lags, only valid when type is 2. | [optional] 
**Mac** | Pointer to **string** | Device Mac. | [optional] 
**ManuallySelect** | Pointer to **bool** | Whether the device is manually selected. | [optional] 
**MlagPeerDeviceMac** | Pointer to **string** | Mlag peer device mac, only valid when the device is in a mlag group. | [optional] 
**Mode** | Pointer to **string** | Device Model. | [optional] 
**ModelVersion** | Pointer to **string** | Device Model Version. | [optional] 
**Name** | Pointer to **string** | Device Name. | [optional] 
**Ports** | Pointer to [**[]SelectPortForVlanVO**](SelectPortForVlanVO.md) | Device Ports | [optional] 
**ReplacedDevice** | Pointer to **bool** | It indicates whether the switch is the dhcp server device that is replaced in the first step | [optional] 
**StatusCategory** | Pointer to **int32** | Device status category, 0: Disconnected, 1: Connected, 2: Pending,3: Heartbeat Missed, 4: Isolated | [optional] 
**Type** | Pointer to **string** | Device type, 1: gateway  2: switch | [optional] 

## Methods

### NewSelectDeviceForVlanVO

`func NewSelectDeviceForVlanVO() *SelectDeviceForVlanVO`

NewSelectDeviceForVlanVO instantiates a new SelectDeviceForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectDeviceForVlanVOWithDefaults

`func NewSelectDeviceForVlanVOWithDefaults() *SelectDeviceForVlanVO`

NewSelectDeviceForVlanVOWithDefaults instantiates a new SelectDeviceForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedInAdvanced

`func (o *SelectDeviceForVlanVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *SelectDeviceForVlanVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *SelectDeviceForVlanVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *SelectDeviceForVlanVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetCompatible

`func (o *SelectDeviceForVlanVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *SelectDeviceForVlanVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *SelectDeviceForVlanVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *SelectDeviceForVlanVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetEs

`func (o *SelectDeviceForVlanVO) GetEs() bool`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *SelectDeviceForVlanVO) GetEsOk() (*bool, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *SelectDeviceForVlanVO) SetEs(v bool)`

SetEs sets Es field to given value.

### HasEs

`func (o *SelectDeviceForVlanVO) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetEsAllowAllPortExceed

`func (o *SelectDeviceForVlanVO) GetEsAllowAllPortExceed() bool`

GetEsAllowAllPortExceed returns the EsAllowAllPortExceed field if non-nil, zero value otherwise.

### GetEsAllowAllPortExceedOk

`func (o *SelectDeviceForVlanVO) GetEsAllowAllPortExceedOk() (*bool, bool)`

GetEsAllowAllPortExceedOk returns a tuple with the EsAllowAllPortExceed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsAllowAllPortExceed

`func (o *SelectDeviceForVlanVO) SetEsAllowAllPortExceed(v bool)`

SetEsAllowAllPortExceed sets EsAllowAllPortExceed field to given value.

### HasEsAllowAllPortExceed

`func (o *SelectDeviceForVlanVO) HasEsAllowAllPortExceed() bool`

HasEsAllowAllPortExceed returns a boolean if a field has been set.

### GetLags

`func (o *SelectDeviceForVlanVO) GetLags() []SelectLagForVlanVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *SelectDeviceForVlanVO) GetLagsOk() (*[]SelectLagForVlanVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *SelectDeviceForVlanVO) SetLags(v []SelectLagForVlanVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *SelectDeviceForVlanVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetMac

`func (o *SelectDeviceForVlanVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SelectDeviceForVlanVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SelectDeviceForVlanVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SelectDeviceForVlanVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManuallySelect

`func (o *SelectDeviceForVlanVO) GetManuallySelect() bool`

GetManuallySelect returns the ManuallySelect field if non-nil, zero value otherwise.

### GetManuallySelectOk

`func (o *SelectDeviceForVlanVO) GetManuallySelectOk() (*bool, bool)`

GetManuallySelectOk returns a tuple with the ManuallySelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallySelect

`func (o *SelectDeviceForVlanVO) SetManuallySelect(v bool)`

SetManuallySelect sets ManuallySelect field to given value.

### HasManuallySelect

`func (o *SelectDeviceForVlanVO) HasManuallySelect() bool`

HasManuallySelect returns a boolean if a field has been set.

### GetMlagPeerDeviceMac

`func (o *SelectDeviceForVlanVO) GetMlagPeerDeviceMac() string`

GetMlagPeerDeviceMac returns the MlagPeerDeviceMac field if non-nil, zero value otherwise.

### GetMlagPeerDeviceMacOk

`func (o *SelectDeviceForVlanVO) GetMlagPeerDeviceMacOk() (*string, bool)`

GetMlagPeerDeviceMacOk returns a tuple with the MlagPeerDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerDeviceMac

`func (o *SelectDeviceForVlanVO) SetMlagPeerDeviceMac(v string)`

SetMlagPeerDeviceMac sets MlagPeerDeviceMac field to given value.

### HasMlagPeerDeviceMac

`func (o *SelectDeviceForVlanVO) HasMlagPeerDeviceMac() bool`

HasMlagPeerDeviceMac returns a boolean if a field has been set.

### GetMode

`func (o *SelectDeviceForVlanVO) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SelectDeviceForVlanVO) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SelectDeviceForVlanVO) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SelectDeviceForVlanVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetModelVersion

`func (o *SelectDeviceForVlanVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *SelectDeviceForVlanVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *SelectDeviceForVlanVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *SelectDeviceForVlanVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *SelectDeviceForVlanVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelectDeviceForVlanVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelectDeviceForVlanVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelectDeviceForVlanVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *SelectDeviceForVlanVO) GetPorts() []SelectPortForVlanVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *SelectDeviceForVlanVO) GetPortsOk() (*[]SelectPortForVlanVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *SelectDeviceForVlanVO) SetPorts(v []SelectPortForVlanVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *SelectDeviceForVlanVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetReplacedDevice

`func (o *SelectDeviceForVlanVO) GetReplacedDevice() bool`

GetReplacedDevice returns the ReplacedDevice field if non-nil, zero value otherwise.

### GetReplacedDeviceOk

`func (o *SelectDeviceForVlanVO) GetReplacedDeviceOk() (*bool, bool)`

GetReplacedDeviceOk returns a tuple with the ReplacedDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedDevice

`func (o *SelectDeviceForVlanVO) SetReplacedDevice(v bool)`

SetReplacedDevice sets ReplacedDevice field to given value.

### HasReplacedDevice

`func (o *SelectDeviceForVlanVO) HasReplacedDevice() bool`

HasReplacedDevice returns a boolean if a field has been set.

### GetStatusCategory

`func (o *SelectDeviceForVlanVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *SelectDeviceForVlanVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *SelectDeviceForVlanVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *SelectDeviceForVlanVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetType

`func (o *SelectDeviceForVlanVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectDeviceForVlanVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectDeviceForVlanVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SelectDeviceForVlanVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


