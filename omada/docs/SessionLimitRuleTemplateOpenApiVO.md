# SessionLimitRuleTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | IP of the session limit rule. | [optional] 
**MaxSession** | **int32** | Max sessions should be within the range of 1–999999. | 
**Name** | **string** | Name should contain 1 to 64 characters. | 
**SourceIds** | Pointer to **[]string** | Source IDs of the session limit rule, only for network and IP group type.Network can be created using &#39;Create LAN network template&#39; interface, and network ID can be obtained from &#39;Get LAN network template list&#39; interface. IP group can be created using &#39;Create a new group profile template&#39; interface, and IP group ID can be obtained from &#39;Get group profile template list&#39; interface. | [optional] 
**SourceType** | **int32** | Source type should be a value as follows: 0: network; 1: IP group; 2: IP. | 
**Status** | **bool** | Status of the session limit rule. | 

## Methods

### NewSessionLimitRuleTemplateOpenApiVO

`func NewSessionLimitRuleTemplateOpenApiVO(maxSession int32, name string, sourceType int32, status bool, ) *SessionLimitRuleTemplateOpenApiVO`

NewSessionLimitRuleTemplateOpenApiVO instantiates a new SessionLimitRuleTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionLimitRuleTemplateOpenApiVOWithDefaults

`func NewSessionLimitRuleTemplateOpenApiVOWithDefaults() *SessionLimitRuleTemplateOpenApiVO`

NewSessionLimitRuleTemplateOpenApiVOWithDefaults instantiates a new SessionLimitRuleTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *SessionLimitRuleTemplateOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SessionLimitRuleTemplateOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SessionLimitRuleTemplateOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SessionLimitRuleTemplateOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMaxSession

`func (o *SessionLimitRuleTemplateOpenApiVO) GetMaxSession() int32`

GetMaxSession returns the MaxSession field if non-nil, zero value otherwise.

### GetMaxSessionOk

`func (o *SessionLimitRuleTemplateOpenApiVO) GetMaxSessionOk() (*int32, bool)`

GetMaxSessionOk returns a tuple with the MaxSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSession

`func (o *SessionLimitRuleTemplateOpenApiVO) SetMaxSession(v int32)`

SetMaxSession sets MaxSession field to given value.


### GetName

`func (o *SessionLimitRuleTemplateOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionLimitRuleTemplateOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionLimitRuleTemplateOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetSourceIds

`func (o *SessionLimitRuleTemplateOpenApiVO) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *SessionLimitRuleTemplateOpenApiVO) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *SessionLimitRuleTemplateOpenApiVO) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *SessionLimitRuleTemplateOpenApiVO) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### GetSourceType

`func (o *SessionLimitRuleTemplateOpenApiVO) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *SessionLimitRuleTemplateOpenApiVO) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *SessionLimitRuleTemplateOpenApiVO) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.


### GetStatus

`func (o *SessionLimitRuleTemplateOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SessionLimitRuleTemplateOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SessionLimitRuleTemplateOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


