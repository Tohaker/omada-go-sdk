# AddRuleEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | **[]int32** | Application ID list can be obtained from &#39;Get application list&#39; interface. | 
**Qos** | **bool** | Enable qos. true:enable / false:disable | 
**QosClass** | Pointer to **int32** | The Class value selected in the Qos Class configuration, required when qos is enable. Valid values is 0: Others, 1: Class 1, 2: Class 2, 3: Class 3. | [optional] 
**RuleName** | **string** | Rule name. It should be 1 - 128 characters | 
**Schedule** | **string** | Schedule profile ID, which can be queried by request： Get time range profile list. | 
**SelectType** | **string** | Select type of applications. include: include selected applications, exclude: all but exclude selected applications, all: include all applications. | 

## Methods

### NewAddRuleEntity

`func NewAddRuleEntity(applications []int32, qos bool, ruleName string, schedule string, selectType string, ) *AddRuleEntity`

NewAddRuleEntity instantiates a new AddRuleEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRuleEntityWithDefaults

`func NewAddRuleEntityWithDefaults() *AddRuleEntity`

NewAddRuleEntityWithDefaults instantiates a new AddRuleEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *AddRuleEntity) GetApplications() []int32`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *AddRuleEntity) GetApplicationsOk() (*[]int32, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *AddRuleEntity) SetApplications(v []int32)`

SetApplications sets Applications field to given value.


### GetQos

`func (o *AddRuleEntity) GetQos() bool`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *AddRuleEntity) GetQosOk() (*bool, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *AddRuleEntity) SetQos(v bool)`

SetQos sets Qos field to given value.


### GetQosClass

`func (o *AddRuleEntity) GetQosClass() int32`

GetQosClass returns the QosClass field if non-nil, zero value otherwise.

### GetQosClassOk

`func (o *AddRuleEntity) GetQosClassOk() (*int32, bool)`

GetQosClassOk returns a tuple with the QosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosClass

`func (o *AddRuleEntity) SetQosClass(v int32)`

SetQosClass sets QosClass field to given value.

### HasQosClass

`func (o *AddRuleEntity) HasQosClass() bool`

HasQosClass returns a boolean if a field has been set.

### GetRuleName

`func (o *AddRuleEntity) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *AddRuleEntity) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *AddRuleEntity) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.


### GetSchedule

`func (o *AddRuleEntity) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AddRuleEntity) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AddRuleEntity) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetSelectType

`func (o *AddRuleEntity) GetSelectType() string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *AddRuleEntity) GetSelectTypeOk() (*string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *AddRuleEntity) SetSelectType(v string)`

SetSelectType sets SelectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


