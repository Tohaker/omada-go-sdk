# MonitorDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientFlag** | Pointer to **bool** | Whether the device is a client | [optional] 
**ClientType** | Pointer to **string** | Client type | [optional] 
**CustomId** | Pointer to **string** | The device Id | [optional] 
**DeviceType** | Pointer to **int32** | The device type | [optional] 
**EligibleFlag** | Pointer to **bool** | Whether the device can be monitored | [optional] 
**Id** | Pointer to **string** | The monitor device Id | [optional] 
**Ip** | Pointer to **string** | The device ip | [optional] 
**Mac** | **string** | The device mac | 
**Model** | Pointer to **string** | The device model | [optional] 
**ModelModelVersion** | Pointer to **string** | The device modelModelVersion | [optional] 
**ModelVersion** | Pointer to **string** | The device modelVersion | [optional] 
**MonitoringStatus** | Pointer to **int32** | The device monitoringStatus | [optional] 
**Name** | Pointer to **string** | The device name | [optional] 
**NotEligibleReason** | Pointer to **string** | The reason of the device can&#39;t be monitored | [optional] 
**RecoveryCycles** | Pointer to **int32** | The device recoveryCycles | [optional] 
**ShowModel** | Pointer to **string** | The device showModel | [optional] 
**SpecialModel** | Pointer to **string** | The device specialModel | [optional] 
**StatusCategory** | Pointer to **int32** | Whether the device is online | [optional] 
**Type** | Pointer to **string** | The device type.Such as: ap, switch, gateway | [optional] 
**UplinkDevice** | Pointer to [**UplinkDevice**](UplinkDevice.md) |  | [optional] 

## Methods

### NewMonitorDevice

`func NewMonitorDevice(mac string, ) *MonitorDevice`

NewMonitorDevice instantiates a new MonitorDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorDeviceWithDefaults

`func NewMonitorDeviceWithDefaults() *MonitorDevice`

NewMonitorDeviceWithDefaults instantiates a new MonitorDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientFlag

`func (o *MonitorDevice) GetClientFlag() bool`

GetClientFlag returns the ClientFlag field if non-nil, zero value otherwise.

### GetClientFlagOk

`func (o *MonitorDevice) GetClientFlagOk() (*bool, bool)`

GetClientFlagOk returns a tuple with the ClientFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFlag

`func (o *MonitorDevice) SetClientFlag(v bool)`

SetClientFlag sets ClientFlag field to given value.

### HasClientFlag

`func (o *MonitorDevice) HasClientFlag() bool`

HasClientFlag returns a boolean if a field has been set.

### GetClientType

`func (o *MonitorDevice) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *MonitorDevice) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *MonitorDevice) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *MonitorDevice) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetCustomId

`func (o *MonitorDevice) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *MonitorDevice) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *MonitorDevice) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *MonitorDevice) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetDeviceType

`func (o *MonitorDevice) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MonitorDevice) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *MonitorDevice) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *MonitorDevice) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetEligibleFlag

`func (o *MonitorDevice) GetEligibleFlag() bool`

GetEligibleFlag returns the EligibleFlag field if non-nil, zero value otherwise.

### GetEligibleFlagOk

`func (o *MonitorDevice) GetEligibleFlagOk() (*bool, bool)`

GetEligibleFlagOk returns a tuple with the EligibleFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleFlag

`func (o *MonitorDevice) SetEligibleFlag(v bool)`

SetEligibleFlag sets EligibleFlag field to given value.

### HasEligibleFlag

`func (o *MonitorDevice) HasEligibleFlag() bool`

HasEligibleFlag returns a boolean if a field has been set.

### GetId

`func (o *MonitorDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitorDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitorDevice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MonitorDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *MonitorDevice) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *MonitorDevice) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *MonitorDevice) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *MonitorDevice) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *MonitorDevice) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MonitorDevice) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MonitorDevice) SetMac(v string)`

SetMac sets Mac field to given value.


### GetModel

`func (o *MonitorDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MonitorDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MonitorDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MonitorDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelModelVersion

`func (o *MonitorDevice) GetModelModelVersion() string`

GetModelModelVersion returns the ModelModelVersion field if non-nil, zero value otherwise.

### GetModelModelVersionOk

`func (o *MonitorDevice) GetModelModelVersionOk() (*string, bool)`

GetModelModelVersionOk returns a tuple with the ModelModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelModelVersion

`func (o *MonitorDevice) SetModelModelVersion(v string)`

SetModelModelVersion sets ModelModelVersion field to given value.

### HasModelModelVersion

`func (o *MonitorDevice) HasModelModelVersion() bool`

HasModelModelVersion returns a boolean if a field has been set.

### GetModelVersion

`func (o *MonitorDevice) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *MonitorDevice) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *MonitorDevice) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *MonitorDevice) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMonitoringStatus

`func (o *MonitorDevice) GetMonitoringStatus() int32`

GetMonitoringStatus returns the MonitoringStatus field if non-nil, zero value otherwise.

### GetMonitoringStatusOk

`func (o *MonitorDevice) GetMonitoringStatusOk() (*int32, bool)`

GetMonitoringStatusOk returns a tuple with the MonitoringStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringStatus

`func (o *MonitorDevice) SetMonitoringStatus(v int32)`

SetMonitoringStatus sets MonitoringStatus field to given value.

### HasMonitoringStatus

`func (o *MonitorDevice) HasMonitoringStatus() bool`

HasMonitoringStatus returns a boolean if a field has been set.

### GetName

`func (o *MonitorDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MonitorDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotEligibleReason

`func (o *MonitorDevice) GetNotEligibleReason() string`

GetNotEligibleReason returns the NotEligibleReason field if non-nil, zero value otherwise.

### GetNotEligibleReasonOk

`func (o *MonitorDevice) GetNotEligibleReasonOk() (*string, bool)`

GetNotEligibleReasonOk returns a tuple with the NotEligibleReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotEligibleReason

`func (o *MonitorDevice) SetNotEligibleReason(v string)`

SetNotEligibleReason sets NotEligibleReason field to given value.

### HasNotEligibleReason

`func (o *MonitorDevice) HasNotEligibleReason() bool`

HasNotEligibleReason returns a boolean if a field has been set.

### GetRecoveryCycles

`func (o *MonitorDevice) GetRecoveryCycles() int32`

GetRecoveryCycles returns the RecoveryCycles field if non-nil, zero value otherwise.

### GetRecoveryCyclesOk

`func (o *MonitorDevice) GetRecoveryCyclesOk() (*int32, bool)`

GetRecoveryCyclesOk returns a tuple with the RecoveryCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCycles

`func (o *MonitorDevice) SetRecoveryCycles(v int32)`

SetRecoveryCycles sets RecoveryCycles field to given value.

### HasRecoveryCycles

`func (o *MonitorDevice) HasRecoveryCycles() bool`

HasRecoveryCycles returns a boolean if a field has been set.

### GetShowModel

`func (o *MonitorDevice) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *MonitorDevice) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *MonitorDevice) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *MonitorDevice) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSpecialModel

`func (o *MonitorDevice) GetSpecialModel() string`

GetSpecialModel returns the SpecialModel field if non-nil, zero value otherwise.

### GetSpecialModelOk

`func (o *MonitorDevice) GetSpecialModelOk() (*string, bool)`

GetSpecialModelOk returns a tuple with the SpecialModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialModel

`func (o *MonitorDevice) SetSpecialModel(v string)`

SetSpecialModel sets SpecialModel field to given value.

### HasSpecialModel

`func (o *MonitorDevice) HasSpecialModel() bool`

HasSpecialModel returns a boolean if a field has been set.

### GetStatusCategory

`func (o *MonitorDevice) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *MonitorDevice) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *MonitorDevice) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *MonitorDevice) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetType

`func (o *MonitorDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MonitorDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUplinkDevice

`func (o *MonitorDevice) GetUplinkDevice() UplinkDevice`

GetUplinkDevice returns the UplinkDevice field if non-nil, zero value otherwise.

### GetUplinkDeviceOk

`func (o *MonitorDevice) GetUplinkDeviceOk() (*UplinkDevice, bool)`

GetUplinkDeviceOk returns a tuple with the UplinkDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkDevice

`func (o *MonitorDevice) SetUplinkDevice(v UplinkDevice)`

SetUplinkDevice sets UplinkDevice field to given value.

### HasUplinkDevice

`func (o *MonitorDevice) HasUplinkDevice() bool`

HasUplinkDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


