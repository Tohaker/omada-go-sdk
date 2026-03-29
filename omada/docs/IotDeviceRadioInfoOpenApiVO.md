# IotDeviceRadioInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | Device IP | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**Override** | Pointer to **bool** | Whether to override the configuration. | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**TransmitPower** | Pointer to **int32** | Broadcast transmission power.&lt;br /&gt;The parameter [transmitPower] should be a value as follows:[-20, -18, -15, -12, -10, -9, -6, -5, -3, 0, 1, 2, 3, 4, 5, 14, 15, 16, 17, 18, 19, 20].(0 by default) | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 

## Methods

### NewIotDeviceRadioInfoOpenApiVO

`func NewIotDeviceRadioInfoOpenApiVO() *IotDeviceRadioInfoOpenApiVO`

NewIotDeviceRadioInfoOpenApiVO instantiates a new IotDeviceRadioInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIotDeviceRadioInfoOpenApiVOWithDefaults

`func NewIotDeviceRadioInfoOpenApiVOWithDefaults() *IotDeviceRadioInfoOpenApiVO`

NewIotDeviceRadioInfoOpenApiVOWithDefaults instantiates a new IotDeviceRadioInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IotDeviceRadioInfoOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IotDeviceRadioInfoOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IotDeviceRadioInfoOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *IotDeviceRadioInfoOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *IotDeviceRadioInfoOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *IotDeviceRadioInfoOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *IotDeviceRadioInfoOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *IotDeviceRadioInfoOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *IotDeviceRadioInfoOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *IotDeviceRadioInfoOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *IotDeviceRadioInfoOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *IotDeviceRadioInfoOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *IotDeviceRadioInfoOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *IotDeviceRadioInfoOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *IotDeviceRadioInfoOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *IotDeviceRadioInfoOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *IotDeviceRadioInfoOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IotDeviceRadioInfoOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IotDeviceRadioInfoOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IotDeviceRadioInfoOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverride

`func (o *IotDeviceRadioInfoOpenApiVO) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *IotDeviceRadioInfoOpenApiVO) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *IotDeviceRadioInfoOpenApiVO) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *IotDeviceRadioInfoOpenApiVO) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetStatus

`func (o *IotDeviceRadioInfoOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IotDeviceRadioInfoOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IotDeviceRadioInfoOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IotDeviceRadioInfoOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *IotDeviceRadioInfoOpenApiVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *IotDeviceRadioInfoOpenApiVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *IotDeviceRadioInfoOpenApiVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *IotDeviceRadioInfoOpenApiVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetTransmitPower

`func (o *IotDeviceRadioInfoOpenApiVO) GetTransmitPower() int32`

GetTransmitPower returns the TransmitPower field if non-nil, zero value otherwise.

### GetTransmitPowerOk

`func (o *IotDeviceRadioInfoOpenApiVO) GetTransmitPowerOk() (*int32, bool)`

GetTransmitPowerOk returns a tuple with the TransmitPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitPower

`func (o *IotDeviceRadioInfoOpenApiVO) SetTransmitPower(v int32)`

SetTransmitPower sets TransmitPower field to given value.

### HasTransmitPower

`func (o *IotDeviceRadioInfoOpenApiVO) HasTransmitPower() bool`

HasTransmitPower returns a boolean if a field has been set.

### GetType

`func (o *IotDeviceRadioInfoOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IotDeviceRadioInfoOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IotDeviceRadioInfoOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IotDeviceRadioInfoOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


