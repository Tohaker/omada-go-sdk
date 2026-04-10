# MonitorClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientUplinkDevice** | Pointer to [**ClientUplinkDevice**](ClientUplinkDevice.md) |  | [optional] 
**ClientFlag** | Pointer to **bool** | Whether the device is a client | [optional] 
**CustomId** | Pointer to **string** | The client Id | [optional] 
**DeviceType** | Pointer to **int32** | The client type | [optional] 
**EligibleFlag** | Pointer to **bool** | Whether the client can be monitored | [optional] 
**Id** | Pointer to **string** | The monitor client Id | [optional] 
**Ip** | Pointer to **string** | The client ip | [optional] 
**Mac** | **string** | The client mac | 
**Model** | Pointer to **string** | The client model | [optional] 
**ModelModelVersion** | Pointer to **string** | The client modelModelVersion | [optional] 
**ModelVersion** | Pointer to **string** | The client modelVersion | [optional] 
**MonitoringStatus** | Pointer to **int32** | The client monitoringStatus | [optional] 
**Name** | Pointer to **string** | The client name | [optional] 
**NotEligibleReason** | Pointer to **string** | The reason of the client can&#39;t be monitored | [optional] 
**RecoveryCycles** | Pointer to **int32** | The client recoveryCycles | [optional] 
**ShowModel** | Pointer to **string** | The client showModel | [optional] 
**SpecialModel** | Pointer to **string** | The client specialModel | [optional] 
**StatusCategory** | Pointer to **int32** | Whether the client is online | [optional] 
**Type** | Pointer to **string** | The client type.Such as: ap, switch, gateway | [optional] 

## Methods

### NewMonitorClient

`func NewMonitorClient(mac string, ) *MonitorClient`

NewMonitorClient instantiates a new MonitorClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorClientWithDefaults

`func NewMonitorClientWithDefaults() *MonitorClient`

NewMonitorClientWithDefaults instantiates a new MonitorClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientUplinkDevice

`func (o *MonitorClient) GetClientUplinkDevice() ClientUplinkDevice`

GetClientUplinkDevice returns the ClientUplinkDevice field if non-nil, zero value otherwise.

### GetClientUplinkDeviceOk

`func (o *MonitorClient) GetClientUplinkDeviceOk() (*ClientUplinkDevice, bool)`

GetClientUplinkDeviceOk returns a tuple with the ClientUplinkDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUplinkDevice

`func (o *MonitorClient) SetClientUplinkDevice(v ClientUplinkDevice)`

SetClientUplinkDevice sets ClientUplinkDevice field to given value.

### HasClientUplinkDevice

`func (o *MonitorClient) HasClientUplinkDevice() bool`

HasClientUplinkDevice returns a boolean if a field has been set.

### GetClientFlag

`func (o *MonitorClient) GetClientFlag() bool`

GetClientFlag returns the ClientFlag field if non-nil, zero value otherwise.

### GetClientFlagOk

`func (o *MonitorClient) GetClientFlagOk() (*bool, bool)`

GetClientFlagOk returns a tuple with the ClientFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFlag

`func (o *MonitorClient) SetClientFlag(v bool)`

SetClientFlag sets ClientFlag field to given value.

### HasClientFlag

`func (o *MonitorClient) HasClientFlag() bool`

HasClientFlag returns a boolean if a field has been set.

### GetCustomId

`func (o *MonitorClient) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *MonitorClient) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *MonitorClient) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *MonitorClient) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetDeviceType

`func (o *MonitorClient) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MonitorClient) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *MonitorClient) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *MonitorClient) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetEligibleFlag

`func (o *MonitorClient) GetEligibleFlag() bool`

GetEligibleFlag returns the EligibleFlag field if non-nil, zero value otherwise.

### GetEligibleFlagOk

`func (o *MonitorClient) GetEligibleFlagOk() (*bool, bool)`

GetEligibleFlagOk returns a tuple with the EligibleFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleFlag

`func (o *MonitorClient) SetEligibleFlag(v bool)`

SetEligibleFlag sets EligibleFlag field to given value.

### HasEligibleFlag

`func (o *MonitorClient) HasEligibleFlag() bool`

HasEligibleFlag returns a boolean if a field has been set.

### GetId

`func (o *MonitorClient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitorClient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitorClient) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MonitorClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *MonitorClient) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *MonitorClient) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *MonitorClient) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *MonitorClient) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *MonitorClient) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MonitorClient) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MonitorClient) SetMac(v string)`

SetMac sets Mac field to given value.


### GetModel

`func (o *MonitorClient) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MonitorClient) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MonitorClient) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MonitorClient) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelModelVersion

`func (o *MonitorClient) GetModelModelVersion() string`

GetModelModelVersion returns the ModelModelVersion field if non-nil, zero value otherwise.

### GetModelModelVersionOk

`func (o *MonitorClient) GetModelModelVersionOk() (*string, bool)`

GetModelModelVersionOk returns a tuple with the ModelModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelModelVersion

`func (o *MonitorClient) SetModelModelVersion(v string)`

SetModelModelVersion sets ModelModelVersion field to given value.

### HasModelModelVersion

`func (o *MonitorClient) HasModelModelVersion() bool`

HasModelModelVersion returns a boolean if a field has been set.

### GetModelVersion

`func (o *MonitorClient) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *MonitorClient) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *MonitorClient) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *MonitorClient) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMonitoringStatus

`func (o *MonitorClient) GetMonitoringStatus() int32`

GetMonitoringStatus returns the MonitoringStatus field if non-nil, zero value otherwise.

### GetMonitoringStatusOk

`func (o *MonitorClient) GetMonitoringStatusOk() (*int32, bool)`

GetMonitoringStatusOk returns a tuple with the MonitoringStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringStatus

`func (o *MonitorClient) SetMonitoringStatus(v int32)`

SetMonitoringStatus sets MonitoringStatus field to given value.

### HasMonitoringStatus

`func (o *MonitorClient) HasMonitoringStatus() bool`

HasMonitoringStatus returns a boolean if a field has been set.

### GetName

`func (o *MonitorClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MonitorClient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotEligibleReason

`func (o *MonitorClient) GetNotEligibleReason() string`

GetNotEligibleReason returns the NotEligibleReason field if non-nil, zero value otherwise.

### GetNotEligibleReasonOk

`func (o *MonitorClient) GetNotEligibleReasonOk() (*string, bool)`

GetNotEligibleReasonOk returns a tuple with the NotEligibleReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotEligibleReason

`func (o *MonitorClient) SetNotEligibleReason(v string)`

SetNotEligibleReason sets NotEligibleReason field to given value.

### HasNotEligibleReason

`func (o *MonitorClient) HasNotEligibleReason() bool`

HasNotEligibleReason returns a boolean if a field has been set.

### GetRecoveryCycles

`func (o *MonitorClient) GetRecoveryCycles() int32`

GetRecoveryCycles returns the RecoveryCycles field if non-nil, zero value otherwise.

### GetRecoveryCyclesOk

`func (o *MonitorClient) GetRecoveryCyclesOk() (*int32, bool)`

GetRecoveryCyclesOk returns a tuple with the RecoveryCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCycles

`func (o *MonitorClient) SetRecoveryCycles(v int32)`

SetRecoveryCycles sets RecoveryCycles field to given value.

### HasRecoveryCycles

`func (o *MonitorClient) HasRecoveryCycles() bool`

HasRecoveryCycles returns a boolean if a field has been set.

### GetShowModel

`func (o *MonitorClient) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *MonitorClient) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *MonitorClient) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *MonitorClient) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSpecialModel

`func (o *MonitorClient) GetSpecialModel() string`

GetSpecialModel returns the SpecialModel field if non-nil, zero value otherwise.

### GetSpecialModelOk

`func (o *MonitorClient) GetSpecialModelOk() (*string, bool)`

GetSpecialModelOk returns a tuple with the SpecialModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialModel

`func (o *MonitorClient) SetSpecialModel(v string)`

SetSpecialModel sets SpecialModel field to given value.

### HasSpecialModel

`func (o *MonitorClient) HasSpecialModel() bool`

HasSpecialModel returns a boolean if a field has been set.

### GetStatusCategory

`func (o *MonitorClient) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *MonitorClient) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *MonitorClient) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *MonitorClient) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetType

`func (o *MonitorClient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorClient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorClient) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MonitorClient) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


