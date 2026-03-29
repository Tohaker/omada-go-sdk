# StackMsgVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbnormalReason** | Pointer to **int32** | Causes of abnormal | [optional] 
**AccountVrfId** | Pointer to **string** |  | [optional] 
**AuthVrfId** | Pointer to **string** |  | [optional] 
**MasterMac** | Pointer to **string** | Stack masterMac | [optional] 
**Priority** | Pointer to **int32** | Stack priority | [optional] 
**StackId** | Pointer to **string** | Stack ID | [optional] 
**StackName** | Pointer to **string** | Stack name | [optional] 
**StackPorts** | Pointer to [**[]OswStackPortGroupVO**](OswStackPortGroupVO.md) | The stack port that has been configured on the current unit | [optional] 
**StackStatus** | Pointer to **int32** | StackStatus should be a value as follows: 0: normal; 1: abnormal; 2: stack not ready | [optional] 
**Unit** | Pointer to **int32** | Stack unit ID | [optional] 

## Methods

### NewStackMsgVO

`func NewStackMsgVO() *StackMsgVO`

NewStackMsgVO instantiates a new StackMsgVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackMsgVOWithDefaults

`func NewStackMsgVOWithDefaults() *StackMsgVO`

NewStackMsgVOWithDefaults instantiates a new StackMsgVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbnormalReason

`func (o *StackMsgVO) GetAbnormalReason() int32`

GetAbnormalReason returns the AbnormalReason field if non-nil, zero value otherwise.

### GetAbnormalReasonOk

`func (o *StackMsgVO) GetAbnormalReasonOk() (*int32, bool)`

GetAbnormalReasonOk returns a tuple with the AbnormalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalReason

`func (o *StackMsgVO) SetAbnormalReason(v int32)`

SetAbnormalReason sets AbnormalReason field to given value.

### HasAbnormalReason

`func (o *StackMsgVO) HasAbnormalReason() bool`

HasAbnormalReason returns a boolean if a field has been set.

### GetAccountVrfId

`func (o *StackMsgVO) GetAccountVrfId() string`

GetAccountVrfId returns the AccountVrfId field if non-nil, zero value otherwise.

### GetAccountVrfIdOk

`func (o *StackMsgVO) GetAccountVrfIdOk() (*string, bool)`

GetAccountVrfIdOk returns a tuple with the AccountVrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountVrfId

`func (o *StackMsgVO) SetAccountVrfId(v string)`

SetAccountVrfId sets AccountVrfId field to given value.

### HasAccountVrfId

`func (o *StackMsgVO) HasAccountVrfId() bool`

HasAccountVrfId returns a boolean if a field has been set.

### GetAuthVrfId

`func (o *StackMsgVO) GetAuthVrfId() string`

GetAuthVrfId returns the AuthVrfId field if non-nil, zero value otherwise.

### GetAuthVrfIdOk

`func (o *StackMsgVO) GetAuthVrfIdOk() (*string, bool)`

GetAuthVrfIdOk returns a tuple with the AuthVrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthVrfId

`func (o *StackMsgVO) SetAuthVrfId(v string)`

SetAuthVrfId sets AuthVrfId field to given value.

### HasAuthVrfId

`func (o *StackMsgVO) HasAuthVrfId() bool`

HasAuthVrfId returns a boolean if a field has been set.

### GetMasterMac

`func (o *StackMsgVO) GetMasterMac() string`

GetMasterMac returns the MasterMac field if non-nil, zero value otherwise.

### GetMasterMacOk

`func (o *StackMsgVO) GetMasterMacOk() (*string, bool)`

GetMasterMacOk returns a tuple with the MasterMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterMac

`func (o *StackMsgVO) SetMasterMac(v string)`

SetMasterMac sets MasterMac field to given value.

### HasMasterMac

`func (o *StackMsgVO) HasMasterMac() bool`

HasMasterMac returns a boolean if a field has been set.

### GetPriority

`func (o *StackMsgVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *StackMsgVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *StackMsgVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *StackMsgVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStackId

`func (o *StackMsgVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *StackMsgVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *StackMsgVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *StackMsgVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *StackMsgVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *StackMsgVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *StackMsgVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *StackMsgVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStackPorts

`func (o *StackMsgVO) GetStackPorts() []OswStackPortGroupVO`

GetStackPorts returns the StackPorts field if non-nil, zero value otherwise.

### GetStackPortsOk

`func (o *StackMsgVO) GetStackPortsOk() (*[]OswStackPortGroupVO, bool)`

GetStackPortsOk returns a tuple with the StackPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPorts

`func (o *StackMsgVO) SetStackPorts(v []OswStackPortGroupVO)`

SetStackPorts sets StackPorts field to given value.

### HasStackPorts

`func (o *StackMsgVO) HasStackPorts() bool`

HasStackPorts returns a boolean if a field has been set.

### GetStackStatus

`func (o *StackMsgVO) GetStackStatus() int32`

GetStackStatus returns the StackStatus field if non-nil, zero value otherwise.

### GetStackStatusOk

`func (o *StackMsgVO) GetStackStatusOk() (*int32, bool)`

GetStackStatusOk returns a tuple with the StackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackStatus

`func (o *StackMsgVO) SetStackStatus(v int32)`

SetStackStatus sets StackStatus field to given value.

### HasStackStatus

`func (o *StackMsgVO) HasStackStatus() bool`

HasStackStatus returns a boolean if a field has been set.

### GetUnit

`func (o *StackMsgVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *StackMsgVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *StackMsgVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *StackMsgVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


