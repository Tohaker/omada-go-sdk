# ApPlanningHistoryDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailMsg** | Pointer to **string** | Fail message. | [optional] 
**FailMsgType** | Pointer to **int32** | Fail message type. -1: Success. 0: Failed to optimize device because of no scan result. 1: Failed to apply deploy config because the device is not connected or the configuration is invalid. | [optional] 
**Ip** | Pointer to **string** | IP Address. | [optional] 
**Model** | Pointer to **string** | Device model. | [optional] 
**ModelVersion** | Pointer to **string** | Device model version. | [optional] 
**Name** | Pointer to **string** | Device name. | [optional] 
**OptimizeSuccess** | Pointer to **bool** | Parameter [optimizeSuccess] indicates whether the optimization has been executed successfully. | [optional] 
**Radio2g** | Pointer to [**ApPlanningRadioVO**](ApPlanningRadioVO.md) |  | [optional] 
**Radio5g** | Pointer to [**ApPlanningRadioVO**](ApPlanningRadioVO.md) |  | [optional] 
**Radio5g2** | Pointer to [**ApPlanningRadioVO**](ApPlanningRadioVO.md) |  | [optional] 
**Radio6g** | Pointer to [**ApPlanningRadioVO**](ApPlanningRadioVO.md) |  | [optional] 
**Support5g2** | Pointer to **bool** | Parameter [support5g2] indicates whether the device supports 5 GHz-2. | [optional] 
**Type** | Pointer to **string** | Device type. | [optional] 

## Methods

### NewApPlanningHistoryDetailVO

`func NewApPlanningHistoryDetailVO() *ApPlanningHistoryDetailVO`

NewApPlanningHistoryDetailVO instantiates a new ApPlanningHistoryDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApPlanningHistoryDetailVOWithDefaults

`func NewApPlanningHistoryDetailVOWithDefaults() *ApPlanningHistoryDetailVO`

NewApPlanningHistoryDetailVOWithDefaults instantiates a new ApPlanningHistoryDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailMsg

`func (o *ApPlanningHistoryDetailVO) GetFailMsg() string`

GetFailMsg returns the FailMsg field if non-nil, zero value otherwise.

### GetFailMsgOk

`func (o *ApPlanningHistoryDetailVO) GetFailMsgOk() (*string, bool)`

GetFailMsgOk returns a tuple with the FailMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailMsg

`func (o *ApPlanningHistoryDetailVO) SetFailMsg(v string)`

SetFailMsg sets FailMsg field to given value.

### HasFailMsg

`func (o *ApPlanningHistoryDetailVO) HasFailMsg() bool`

HasFailMsg returns a boolean if a field has been set.

### GetFailMsgType

`func (o *ApPlanningHistoryDetailVO) GetFailMsgType() int32`

GetFailMsgType returns the FailMsgType field if non-nil, zero value otherwise.

### GetFailMsgTypeOk

`func (o *ApPlanningHistoryDetailVO) GetFailMsgTypeOk() (*int32, bool)`

GetFailMsgTypeOk returns a tuple with the FailMsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailMsgType

`func (o *ApPlanningHistoryDetailVO) SetFailMsgType(v int32)`

SetFailMsgType sets FailMsgType field to given value.

### HasFailMsgType

`func (o *ApPlanningHistoryDetailVO) HasFailMsgType() bool`

HasFailMsgType returns a boolean if a field has been set.

### GetIp

`func (o *ApPlanningHistoryDetailVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ApPlanningHistoryDetailVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ApPlanningHistoryDetailVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ApPlanningHistoryDetailVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetModel

`func (o *ApPlanningHistoryDetailVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApPlanningHistoryDetailVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApPlanningHistoryDetailVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApPlanningHistoryDetailVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ApPlanningHistoryDetailVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ApPlanningHistoryDetailVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ApPlanningHistoryDetailVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ApPlanningHistoryDetailVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *ApPlanningHistoryDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApPlanningHistoryDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApPlanningHistoryDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApPlanningHistoryDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptimizeSuccess

`func (o *ApPlanningHistoryDetailVO) GetOptimizeSuccess() bool`

GetOptimizeSuccess returns the OptimizeSuccess field if non-nil, zero value otherwise.

### GetOptimizeSuccessOk

`func (o *ApPlanningHistoryDetailVO) GetOptimizeSuccessOk() (*bool, bool)`

GetOptimizeSuccessOk returns a tuple with the OptimizeSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeSuccess

`func (o *ApPlanningHistoryDetailVO) SetOptimizeSuccess(v bool)`

SetOptimizeSuccess sets OptimizeSuccess field to given value.

### HasOptimizeSuccess

`func (o *ApPlanningHistoryDetailVO) HasOptimizeSuccess() bool`

HasOptimizeSuccess returns a boolean if a field has been set.

### GetRadio2g

`func (o *ApPlanningHistoryDetailVO) GetRadio2g() ApPlanningRadioVO`

GetRadio2g returns the Radio2g field if non-nil, zero value otherwise.

### GetRadio2gOk

`func (o *ApPlanningHistoryDetailVO) GetRadio2gOk() (*ApPlanningRadioVO, bool)`

GetRadio2gOk returns a tuple with the Radio2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio2g

`func (o *ApPlanningHistoryDetailVO) SetRadio2g(v ApPlanningRadioVO)`

SetRadio2g sets Radio2g field to given value.

### HasRadio2g

`func (o *ApPlanningHistoryDetailVO) HasRadio2g() bool`

HasRadio2g returns a boolean if a field has been set.

### GetRadio5g

`func (o *ApPlanningHistoryDetailVO) GetRadio5g() ApPlanningRadioVO`

GetRadio5g returns the Radio5g field if non-nil, zero value otherwise.

### GetRadio5gOk

`func (o *ApPlanningHistoryDetailVO) GetRadio5gOk() (*ApPlanningRadioVO, bool)`

GetRadio5gOk returns a tuple with the Radio5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio5g

`func (o *ApPlanningHistoryDetailVO) SetRadio5g(v ApPlanningRadioVO)`

SetRadio5g sets Radio5g field to given value.

### HasRadio5g

`func (o *ApPlanningHistoryDetailVO) HasRadio5g() bool`

HasRadio5g returns a boolean if a field has been set.

### GetRadio5g2

`func (o *ApPlanningHistoryDetailVO) GetRadio5g2() ApPlanningRadioVO`

GetRadio5g2 returns the Radio5g2 field if non-nil, zero value otherwise.

### GetRadio5g2Ok

`func (o *ApPlanningHistoryDetailVO) GetRadio5g2Ok() (*ApPlanningRadioVO, bool)`

GetRadio5g2Ok returns a tuple with the Radio5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio5g2

`func (o *ApPlanningHistoryDetailVO) SetRadio5g2(v ApPlanningRadioVO)`

SetRadio5g2 sets Radio5g2 field to given value.

### HasRadio5g2

`func (o *ApPlanningHistoryDetailVO) HasRadio5g2() bool`

HasRadio5g2 returns a boolean if a field has been set.

### GetRadio6g

`func (o *ApPlanningHistoryDetailVO) GetRadio6g() ApPlanningRadioVO`

GetRadio6g returns the Radio6g field if non-nil, zero value otherwise.

### GetRadio6gOk

`func (o *ApPlanningHistoryDetailVO) GetRadio6gOk() (*ApPlanningRadioVO, bool)`

GetRadio6gOk returns a tuple with the Radio6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio6g

`func (o *ApPlanningHistoryDetailVO) SetRadio6g(v ApPlanningRadioVO)`

SetRadio6g sets Radio6g field to given value.

### HasRadio6g

`func (o *ApPlanningHistoryDetailVO) HasRadio6g() bool`

HasRadio6g returns a boolean if a field has been set.

### GetSupport5g2

`func (o *ApPlanningHistoryDetailVO) GetSupport5g2() bool`

GetSupport5g2 returns the Support5g2 field if non-nil, zero value otherwise.

### GetSupport5g2Ok

`func (o *ApPlanningHistoryDetailVO) GetSupport5g2Ok() (*bool, bool)`

GetSupport5g2Ok returns a tuple with the Support5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g2

`func (o *ApPlanningHistoryDetailVO) SetSupport5g2(v bool)`

SetSupport5g2 sets Support5g2 field to given value.

### HasSupport5g2

`func (o *ApPlanningHistoryDetailVO) HasSupport5g2() bool`

HasSupport5g2 returns a boolean if a field has been set.

### GetType

`func (o *ApPlanningHistoryDetailVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApPlanningHistoryDetailVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApPlanningHistoryDetailVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApPlanningHistoryDetailVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


