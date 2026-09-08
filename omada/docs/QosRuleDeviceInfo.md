# QosRuleDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrCode** | Pointer to **int32** | The code of the device&#39;s response after the QoS rules are issued. | [optional] 
**Mac** | Pointer to **string** | The device mac. | [optional] 
**Model** | Pointer to **string** | The device model. | [optional] 
**ModelVersion** | Pointer to **string** | The model version of device, for example:3.0 | [optional] 
**Name** | Pointer to **string** | The device name. | [optional] 
**ShowModel** | Pointer to **string** | The device showModel(model + modelVersion). | [optional] 
**StackDevice** | Pointer to **bool** | Stack device identifier, true: stack device, false: normal device. | [optional] 
**StatusCategory** | Pointer to **int32** | Device status should be a value as follows: 0: Disconnected; 1: Connected; 2: Pending; 3: Heartbeat Missed; 4: Isolated | [optional] 

## Methods

### NewQosRuleDeviceInfo

`func NewQosRuleDeviceInfo() *QosRuleDeviceInfo`

NewQosRuleDeviceInfo instantiates a new QosRuleDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosRuleDeviceInfoWithDefaults

`func NewQosRuleDeviceInfoWithDefaults() *QosRuleDeviceInfo`

NewQosRuleDeviceInfoWithDefaults instantiates a new QosRuleDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrCode

`func (o *QosRuleDeviceInfo) GetErrCode() int32`

GetErrCode returns the ErrCode field if non-nil, zero value otherwise.

### GetErrCodeOk

`func (o *QosRuleDeviceInfo) GetErrCodeOk() (*int32, bool)`

GetErrCodeOk returns a tuple with the ErrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrCode

`func (o *QosRuleDeviceInfo) SetErrCode(v int32)`

SetErrCode sets ErrCode field to given value.

### HasErrCode

`func (o *QosRuleDeviceInfo) HasErrCode() bool`

HasErrCode returns a boolean if a field has been set.

### GetMac

`func (o *QosRuleDeviceInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *QosRuleDeviceInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *QosRuleDeviceInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *QosRuleDeviceInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *QosRuleDeviceInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *QosRuleDeviceInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *QosRuleDeviceInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *QosRuleDeviceInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *QosRuleDeviceInfo) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *QosRuleDeviceInfo) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *QosRuleDeviceInfo) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *QosRuleDeviceInfo) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *QosRuleDeviceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QosRuleDeviceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QosRuleDeviceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QosRuleDeviceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShowModel

`func (o *QosRuleDeviceInfo) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *QosRuleDeviceInfo) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *QosRuleDeviceInfo) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *QosRuleDeviceInfo) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStackDevice

`func (o *QosRuleDeviceInfo) GetStackDevice() bool`

GetStackDevice returns the StackDevice field if non-nil, zero value otherwise.

### GetStackDeviceOk

`func (o *QosRuleDeviceInfo) GetStackDeviceOk() (*bool, bool)`

GetStackDeviceOk returns a tuple with the StackDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackDevice

`func (o *QosRuleDeviceInfo) SetStackDevice(v bool)`

SetStackDevice sets StackDevice field to given value.

### HasStackDevice

`func (o *QosRuleDeviceInfo) HasStackDevice() bool`

HasStackDevice returns a boolean if a field has been set.

### GetStatusCategory

`func (o *QosRuleDeviceInfo) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *QosRuleDeviceInfo) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *QosRuleDeviceInfo) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *QosRuleDeviceInfo) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


