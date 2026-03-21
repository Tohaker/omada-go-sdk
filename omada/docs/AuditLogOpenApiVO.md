# AuditLogOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditType** | Pointer to **string** | Log Type | [optional] 
**Content** | Pointer to **string** | Log Content | [optional] 
**Ip** | Pointer to **string** | User Login IP address | [optional] 
**Label** | Pointer to **string** | Configuration card or request path, may be empty. | [optional] 
**Level** | Pointer to **string** | Log Level. It should be a value as follows: Error, Warning, Information. | [optional] 
**NewValue** | Pointer to **map[string]interface{}** | Configuration after modification, may be empty. | [optional] 
**OldValue** | Pointer to **map[string]interface{}** | Configuration before modification, may be empty. | [optional] 
**Operator** | Pointer to **string** | Operator | [optional] 
**Resource** | Pointer to **string** | Log Creation Resource. It should be a value as follows: WEB、Open API | [optional] 
**Result** | Pointer to **string** | Operation Result, it should be a value as follows: Succeed、Failed | [optional] 
**Time** | Pointer to **int64** | Log Creation TimeStamp, Unit:ms | [optional] 

## Methods

### NewAuditLogOpenApiVO

`func NewAuditLogOpenApiVO() *AuditLogOpenApiVO`

NewAuditLogOpenApiVO instantiates a new AuditLogOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogOpenApiVOWithDefaults

`func NewAuditLogOpenApiVOWithDefaults() *AuditLogOpenApiVO`

NewAuditLogOpenApiVOWithDefaults instantiates a new AuditLogOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditType

`func (o *AuditLogOpenApiVO) GetAuditType() string`

GetAuditType returns the AuditType field if non-nil, zero value otherwise.

### GetAuditTypeOk

`func (o *AuditLogOpenApiVO) GetAuditTypeOk() (*string, bool)`

GetAuditTypeOk returns a tuple with the AuditType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditType

`func (o *AuditLogOpenApiVO) SetAuditType(v string)`

SetAuditType sets AuditType field to given value.

### HasAuditType

`func (o *AuditLogOpenApiVO) HasAuditType() bool`

HasAuditType returns a boolean if a field has been set.

### GetContent

`func (o *AuditLogOpenApiVO) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AuditLogOpenApiVO) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AuditLogOpenApiVO) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AuditLogOpenApiVO) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetIp

`func (o *AuditLogOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AuditLogOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AuditLogOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AuditLogOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLabel

`func (o *AuditLogOpenApiVO) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AuditLogOpenApiVO) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AuditLogOpenApiVO) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AuditLogOpenApiVO) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLevel

`func (o *AuditLogOpenApiVO) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AuditLogOpenApiVO) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AuditLogOpenApiVO) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AuditLogOpenApiVO) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetNewValue

`func (o *AuditLogOpenApiVO) GetNewValue() map[string]interface{}`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *AuditLogOpenApiVO) GetNewValueOk() (*map[string]interface{}, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *AuditLogOpenApiVO) SetNewValue(v map[string]interface{})`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *AuditLogOpenApiVO) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetOldValue

`func (o *AuditLogOpenApiVO) GetOldValue() map[string]interface{}`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *AuditLogOpenApiVO) GetOldValueOk() (*map[string]interface{}, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *AuditLogOpenApiVO) SetOldValue(v map[string]interface{})`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *AuditLogOpenApiVO) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetOperator

`func (o *AuditLogOpenApiVO) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AuditLogOpenApiVO) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AuditLogOpenApiVO) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AuditLogOpenApiVO) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetResource

`func (o *AuditLogOpenApiVO) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AuditLogOpenApiVO) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AuditLogOpenApiVO) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AuditLogOpenApiVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResult

`func (o *AuditLogOpenApiVO) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AuditLogOpenApiVO) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AuditLogOpenApiVO) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *AuditLogOpenApiVO) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetTime

`func (o *AuditLogOpenApiVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AuditLogOpenApiVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AuditLogOpenApiVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *AuditLogOpenApiVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


