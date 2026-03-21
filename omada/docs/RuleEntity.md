# RuleEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]ApplicationEntity**](ApplicationEntity.md) | Application list | [optional] 
**Qos** | Pointer to **bool** | Enable qos. true:enable / false:disable | [optional] 
**QosClass** | Pointer to **int32** | Qos class category | [optional] 
**RuleId** | Pointer to **int32** | Rule ID | [optional] 
**RuleName** | Pointer to **string** | Rule name | [optional] 
**Schedule** | Pointer to **string** | Schedule profile ID | [optional] 

## Methods

### NewRuleEntity

`func NewRuleEntity() *RuleEntity`

NewRuleEntity instantiates a new RuleEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleEntityWithDefaults

`func NewRuleEntityWithDefaults() *RuleEntity`

NewRuleEntityWithDefaults instantiates a new RuleEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *RuleEntity) GetApplications() []ApplicationEntity`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *RuleEntity) GetApplicationsOk() (*[]ApplicationEntity, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *RuleEntity) SetApplications(v []ApplicationEntity)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *RuleEntity) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetQos

`func (o *RuleEntity) GetQos() bool`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *RuleEntity) GetQosOk() (*bool, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *RuleEntity) SetQos(v bool)`

SetQos sets Qos field to given value.

### HasQos

`func (o *RuleEntity) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetQosClass

`func (o *RuleEntity) GetQosClass() int32`

GetQosClass returns the QosClass field if non-nil, zero value otherwise.

### GetQosClassOk

`func (o *RuleEntity) GetQosClassOk() (*int32, bool)`

GetQosClassOk returns a tuple with the QosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosClass

`func (o *RuleEntity) SetQosClass(v int32)`

SetQosClass sets QosClass field to given value.

### HasQosClass

`func (o *RuleEntity) HasQosClass() bool`

HasQosClass returns a boolean if a field has been set.

### GetRuleId

`func (o *RuleEntity) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RuleEntity) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RuleEntity) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *RuleEntity) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleName

`func (o *RuleEntity) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *RuleEntity) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *RuleEntity) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *RuleEntity) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetSchedule

`func (o *RuleEntity) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *RuleEntity) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *RuleEntity) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *RuleEntity) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


