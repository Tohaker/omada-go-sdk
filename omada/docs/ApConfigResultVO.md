# ApConfigResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApConfigResultSettings** | Pointer to [**ApConfigResultSettingsVO**](ApConfigResultSettingsVO.md) |  | [optional] 
**ErrorCode** | Pointer to **int32** | error code. | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Msg** | Pointer to **string** | error msg | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 

## Methods

### NewApConfigResultVO

`func NewApConfigResultVO() *ApConfigResultVO`

NewApConfigResultVO instantiates a new ApConfigResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApConfigResultVOWithDefaults

`func NewApConfigResultVOWithDefaults() *ApConfigResultVO`

NewApConfigResultVOWithDefaults instantiates a new ApConfigResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApConfigResultSettings

`func (o *ApConfigResultVO) GetApConfigResultSettings() ApConfigResultSettingsVO`

GetApConfigResultSettings returns the ApConfigResultSettings field if non-nil, zero value otherwise.

### GetApConfigResultSettingsOk

`func (o *ApConfigResultVO) GetApConfigResultSettingsOk() (*ApConfigResultSettingsVO, bool)`

GetApConfigResultSettingsOk returns a tuple with the ApConfigResultSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApConfigResultSettings

`func (o *ApConfigResultVO) SetApConfigResultSettings(v ApConfigResultSettingsVO)`

SetApConfigResultSettings sets ApConfigResultSettings field to given value.

### HasApConfigResultSettings

`func (o *ApConfigResultVO) HasApConfigResultSettings() bool`

HasApConfigResultSettings returns a boolean if a field has been set.

### GetErrorCode

`func (o *ApConfigResultVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ApConfigResultVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ApConfigResultVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ApConfigResultVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMac

`func (o *ApConfigResultVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ApConfigResultVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ApConfigResultVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ApConfigResultVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ApConfigResultVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApConfigResultVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApConfigResultVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApConfigResultVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ApConfigResultVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ApConfigResultVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ApConfigResultVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ApConfigResultVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMsg

`func (o *ApConfigResultVO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ApConfigResultVO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ApConfigResultVO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ApConfigResultVO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetName

`func (o *ApConfigResultVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApConfigResultVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApConfigResultVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApConfigResultVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ApConfigResultVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApConfigResultVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApConfigResultVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApConfigResultVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *ApConfigResultVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *ApConfigResultVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *ApConfigResultVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *ApConfigResultVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetType

`func (o *ApConfigResultVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApConfigResultVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApConfigResultVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApConfigResultVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


