# QuerySessionLimitRuleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistIpAddress** | Pointer to **bool** | Whether Source Type of Current Session Limit rule is IP Address. | [optional] 
**Id** | Pointer to **string** | ID of the session limit rule. | [optional] 
**Index** | Pointer to **int32** | Index of the session limit rule. When the [sourceType] is 0 or 1, the index is counted in order, when the sourceType is 2, the index is always -1 | [optional] 
**Ip** | Pointer to **string** | IP of the session limit rule. | [optional] 
**MaxSession** | **int32** | Max sessions should be within the range of 1–999999. | 
**Name** | **string** | Name should contain 1 to 64 characters. | 
**SourceIds** | Pointer to **[]string** | Source IDs of the session limit rule, only for network and IP group type.Network can be created using &#39;Create LAN network&#39; interface, and network ID can be obtained from &#39;Get LAN network list&#39; interface. IP group can be created using &#39;Create a new group profile&#39; interface, and IP group ID can be obtained from &#39;Get group profile list&#39; interface. | [optional] 
**SourceType** | **int32** | Source type should be a value as follows: 0: network; 1: IP group; 2: IP. | 
**Status** | **bool** | Status of the session limit rule. | 

## Methods

### NewQuerySessionLimitRuleOpenApiVO

`func NewQuerySessionLimitRuleOpenApiVO(maxSession int32, name string, sourceType int32, status bool, ) *QuerySessionLimitRuleOpenApiVO`

NewQuerySessionLimitRuleOpenApiVO instantiates a new QuerySessionLimitRuleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySessionLimitRuleOpenApiVOWithDefaults

`func NewQuerySessionLimitRuleOpenApiVOWithDefaults() *QuerySessionLimitRuleOpenApiVO`

NewQuerySessionLimitRuleOpenApiVOWithDefaults instantiates a new QuerySessionLimitRuleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistIpAddress

`func (o *QuerySessionLimitRuleOpenApiVO) GetExistIpAddress() bool`

GetExistIpAddress returns the ExistIpAddress field if non-nil, zero value otherwise.

### GetExistIpAddressOk

`func (o *QuerySessionLimitRuleOpenApiVO) GetExistIpAddressOk() (*bool, bool)`

GetExistIpAddressOk returns a tuple with the ExistIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistIpAddress

`func (o *QuerySessionLimitRuleOpenApiVO) SetExistIpAddress(v bool)`

SetExistIpAddress sets ExistIpAddress field to given value.

### HasExistIpAddress

`func (o *QuerySessionLimitRuleOpenApiVO) HasExistIpAddress() bool`

HasExistIpAddress returns a boolean if a field has been set.

### GetId

`func (o *QuerySessionLimitRuleOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuerySessionLimitRuleOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuerySessionLimitRuleOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QuerySessionLimitRuleOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *QuerySessionLimitRuleOpenApiVO) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *QuerySessionLimitRuleOpenApiVO) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *QuerySessionLimitRuleOpenApiVO) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *QuerySessionLimitRuleOpenApiVO) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIp

`func (o *QuerySessionLimitRuleOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *QuerySessionLimitRuleOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *QuerySessionLimitRuleOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *QuerySessionLimitRuleOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMaxSession

`func (o *QuerySessionLimitRuleOpenApiVO) GetMaxSession() int32`

GetMaxSession returns the MaxSession field if non-nil, zero value otherwise.

### GetMaxSessionOk

`func (o *QuerySessionLimitRuleOpenApiVO) GetMaxSessionOk() (*int32, bool)`

GetMaxSessionOk returns a tuple with the MaxSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSession

`func (o *QuerySessionLimitRuleOpenApiVO) SetMaxSession(v int32)`

SetMaxSession sets MaxSession field to given value.


### GetName

`func (o *QuerySessionLimitRuleOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuerySessionLimitRuleOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuerySessionLimitRuleOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetSourceIds

`func (o *QuerySessionLimitRuleOpenApiVO) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *QuerySessionLimitRuleOpenApiVO) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *QuerySessionLimitRuleOpenApiVO) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *QuerySessionLimitRuleOpenApiVO) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### GetSourceType

`func (o *QuerySessionLimitRuleOpenApiVO) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *QuerySessionLimitRuleOpenApiVO) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *QuerySessionLimitRuleOpenApiVO) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.


### GetStatus

`func (o *QuerySessionLimitRuleOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QuerySessionLimitRuleOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QuerySessionLimitRuleOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


