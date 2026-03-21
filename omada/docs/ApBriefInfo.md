# ApBriefInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Mark whether the device is activated: When license (specific to cloud base) is false, the status column shows the pre-bound status. When it is true or null, the status column displays as it originally did. | [optional] 
**Ip** | Pointer to **string** | Device IP | [optional] 
**Mac** | Pointer to **string** | Device MAC, e.g. 00-00-FF-FF-0C-E9 | [optional] 
**Model** | Pointer to **string** | Device model name, such as EAP225. | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**ShowModel** | Pointer to **string** | Device model name with version, such as EAP225(EU) v3.0. | [optional] 
**SpecialModel** | Pointer to **string** | Special device model,for example:EAP225-Outdoor-1a20a950b8d950e8 | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Device status should be a value as follows: 0: Disconnected; 1: Connected; 2: Pending; 3: Heartbeat Missed; 4: Isolated | [optional] 
**SupportLockToAp** | Pointer to **bool** | Whether the device support function [Lock to ap]. | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 

## Methods

### NewApBriefInfo

`func NewApBriefInfo() *ApBriefInfo`

NewApBriefInfo instantiates a new ApBriefInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApBriefInfoWithDefaults

`func NewApBriefInfoWithDefaults() *ApBriefInfo`

NewApBriefInfoWithDefaults instantiates a new ApBriefInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ApBriefInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApBriefInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApBriefInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApBriefInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetIp

`func (o *ApBriefInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ApBriefInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ApBriefInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ApBriefInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *ApBriefInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ApBriefInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ApBriefInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ApBriefInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ApBriefInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApBriefInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApBriefInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApBriefInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ApBriefInfo) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ApBriefInfo) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ApBriefInfo) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ApBriefInfo) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *ApBriefInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApBriefInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApBriefInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApBriefInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShowModel

`func (o *ApBriefInfo) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *ApBriefInfo) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *ApBriefInfo) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *ApBriefInfo) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSpecialModel

`func (o *ApBriefInfo) GetSpecialModel() string`

GetSpecialModel returns the SpecialModel field if non-nil, zero value otherwise.

### GetSpecialModelOk

`func (o *ApBriefInfo) GetSpecialModelOk() (*string, bool)`

GetSpecialModelOk returns a tuple with the SpecialModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialModel

`func (o *ApBriefInfo) SetSpecialModel(v string)`

SetSpecialModel sets SpecialModel field to given value.

### HasSpecialModel

`func (o *ApBriefInfo) HasSpecialModel() bool`

HasSpecialModel returns a boolean if a field has been set.

### GetStatus

`func (o *ApBriefInfo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApBriefInfo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApBriefInfo) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApBriefInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *ApBriefInfo) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *ApBriefInfo) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *ApBriefInfo) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *ApBriefInfo) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetSupportLockToAp

`func (o *ApBriefInfo) GetSupportLockToAp() bool`

GetSupportLockToAp returns the SupportLockToAp field if non-nil, zero value otherwise.

### GetSupportLockToApOk

`func (o *ApBriefInfo) GetSupportLockToApOk() (*bool, bool)`

GetSupportLockToApOk returns a tuple with the SupportLockToAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLockToAp

`func (o *ApBriefInfo) SetSupportLockToAp(v bool)`

SetSupportLockToAp sets SupportLockToAp field to given value.

### HasSupportLockToAp

`func (o *ApBriefInfo) HasSupportLockToAp() bool`

HasSupportLockToAp returns a boolean if a field has been set.

### GetType

`func (o *ApBriefInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApBriefInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApBriefInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApBriefInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


