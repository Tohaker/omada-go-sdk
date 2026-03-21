# QueryUrlFilteringOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **map[string][]string** | categories of the URL filtering, TreeMap&lt;String, List&lt;String&gt;&gt; categories | [optional] 
**ExistCategory** | Pointer to **bool** | Whether this URL filtering is category mode. | [optional] 
**ExistKeyword** | Pointer to **bool** | Whether this URL filtering is keyword mode. | [optional] 
**ExistTimeSchedule** | Pointer to **bool** | Whether this URL filtering contains time schedule setting. | [optional] 
**FilterMode** | **int32** | filterMode should be a value as follows: 0: URL; 1: category. | 
**Id** | Pointer to **string** | ID of the URL filtering. | [optional] 
**Index** | Pointer to **int32** | Index of the URL filtering. | [optional] 
**Keywords** | Pointer to **[]string** | Keywords of the URL filtering. | [optional] 
**Mode** | Pointer to **int32** | Mode should be a value as follows: 0: URL; 1: keyword. | [optional] 
**Name** | **string** | Name should contain 1 to 64 characters. | 
**Policy** | **int32** | Policy should be a value as follows: 0: drop; 1: allow. | 
**ScenarioMode** | Pointer to **int32** | scenarioMode should be a value as follows:0: Custom  1: Family 2：Work 3：Education 4：Guest. | [optional] 
**SourceIds** | **[]string** | Source IDs of the URL filtering. Network can be created using &#39;Create LAN network&#39; interface, and network ID can be obtained from &#39;Get LAN network list&#39; interface. IP group can be created using &#39;Create a new group profile&#39; interface, and IP group ID can be obtained from &#39;Get group profile list&#39; interface. SSID can be created using &#39;Create new SSID&#39; interface, and SSID ID can be obtained from &#39;Get SSID list&#39; interface. | 
**SourceType** | **int32** | Source type should be a value as follows: 0: network; 1: IP group; 2: SSID. | 
**Status** | **bool** | Status of the URL filtering. | 
**TimeRange** | Pointer to **string** | timeRange of the URL filtering, TimeSchedule | [optional] 
**Type** | **string** | Type should be a value as follows: \&quot;gateway\&quot;; \&quot;ap\&quot;. | 
**Urls** | Pointer to **[]string** | URLs of the URL filtering, eg: www.google.com. | [optional] 

## Methods

### NewQueryUrlFilteringOpenApiVO

`func NewQueryUrlFilteringOpenApiVO(filterMode int32, name string, policy int32, sourceIds []string, sourceType int32, status bool, type_ string, ) *QueryUrlFilteringOpenApiVO`

NewQueryUrlFilteringOpenApiVO instantiates a new QueryUrlFilteringOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryUrlFilteringOpenApiVOWithDefaults

`func NewQueryUrlFilteringOpenApiVOWithDefaults() *QueryUrlFilteringOpenApiVO`

NewQueryUrlFilteringOpenApiVOWithDefaults instantiates a new QueryUrlFilteringOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *QueryUrlFilteringOpenApiVO) GetCategories() map[string][]string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *QueryUrlFilteringOpenApiVO) GetCategoriesOk() (*map[string][]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *QueryUrlFilteringOpenApiVO) SetCategories(v map[string][]string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *QueryUrlFilteringOpenApiVO) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetExistCategory

`func (o *QueryUrlFilteringOpenApiVO) GetExistCategory() bool`

GetExistCategory returns the ExistCategory field if non-nil, zero value otherwise.

### GetExistCategoryOk

`func (o *QueryUrlFilteringOpenApiVO) GetExistCategoryOk() (*bool, bool)`

GetExistCategoryOk returns a tuple with the ExistCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistCategory

`func (o *QueryUrlFilteringOpenApiVO) SetExistCategory(v bool)`

SetExistCategory sets ExistCategory field to given value.

### HasExistCategory

`func (o *QueryUrlFilteringOpenApiVO) HasExistCategory() bool`

HasExistCategory returns a boolean if a field has been set.

### GetExistKeyword

`func (o *QueryUrlFilteringOpenApiVO) GetExistKeyword() bool`

GetExistKeyword returns the ExistKeyword field if non-nil, zero value otherwise.

### GetExistKeywordOk

`func (o *QueryUrlFilteringOpenApiVO) GetExistKeywordOk() (*bool, bool)`

GetExistKeywordOk returns a tuple with the ExistKeyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistKeyword

`func (o *QueryUrlFilteringOpenApiVO) SetExistKeyword(v bool)`

SetExistKeyword sets ExistKeyword field to given value.

### HasExistKeyword

`func (o *QueryUrlFilteringOpenApiVO) HasExistKeyword() bool`

HasExistKeyword returns a boolean if a field has been set.

### GetExistTimeSchedule

`func (o *QueryUrlFilteringOpenApiVO) GetExistTimeSchedule() bool`

GetExistTimeSchedule returns the ExistTimeSchedule field if non-nil, zero value otherwise.

### GetExistTimeScheduleOk

`func (o *QueryUrlFilteringOpenApiVO) GetExistTimeScheduleOk() (*bool, bool)`

GetExistTimeScheduleOk returns a tuple with the ExistTimeSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistTimeSchedule

`func (o *QueryUrlFilteringOpenApiVO) SetExistTimeSchedule(v bool)`

SetExistTimeSchedule sets ExistTimeSchedule field to given value.

### HasExistTimeSchedule

`func (o *QueryUrlFilteringOpenApiVO) HasExistTimeSchedule() bool`

HasExistTimeSchedule returns a boolean if a field has been set.

### GetFilterMode

`func (o *QueryUrlFilteringOpenApiVO) GetFilterMode() int32`

GetFilterMode returns the FilterMode field if non-nil, zero value otherwise.

### GetFilterModeOk

`func (o *QueryUrlFilteringOpenApiVO) GetFilterModeOk() (*int32, bool)`

GetFilterModeOk returns a tuple with the FilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMode

`func (o *QueryUrlFilteringOpenApiVO) SetFilterMode(v int32)`

SetFilterMode sets FilterMode field to given value.


### GetId

`func (o *QueryUrlFilteringOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueryUrlFilteringOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueryUrlFilteringOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueryUrlFilteringOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *QueryUrlFilteringOpenApiVO) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *QueryUrlFilteringOpenApiVO) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *QueryUrlFilteringOpenApiVO) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *QueryUrlFilteringOpenApiVO) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetKeywords

`func (o *QueryUrlFilteringOpenApiVO) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *QueryUrlFilteringOpenApiVO) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *QueryUrlFilteringOpenApiVO) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *QueryUrlFilteringOpenApiVO) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetMode

`func (o *QueryUrlFilteringOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *QueryUrlFilteringOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *QueryUrlFilteringOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *QueryUrlFilteringOpenApiVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *QueryUrlFilteringOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryUrlFilteringOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryUrlFilteringOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPolicy

`func (o *QueryUrlFilteringOpenApiVO) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *QueryUrlFilteringOpenApiVO) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *QueryUrlFilteringOpenApiVO) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetScenarioMode

`func (o *QueryUrlFilteringOpenApiVO) GetScenarioMode() int32`

GetScenarioMode returns the ScenarioMode field if non-nil, zero value otherwise.

### GetScenarioModeOk

`func (o *QueryUrlFilteringOpenApiVO) GetScenarioModeOk() (*int32, bool)`

GetScenarioModeOk returns a tuple with the ScenarioMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioMode

`func (o *QueryUrlFilteringOpenApiVO) SetScenarioMode(v int32)`

SetScenarioMode sets ScenarioMode field to given value.

### HasScenarioMode

`func (o *QueryUrlFilteringOpenApiVO) HasScenarioMode() bool`

HasScenarioMode returns a boolean if a field has been set.

### GetSourceIds

`func (o *QueryUrlFilteringOpenApiVO) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *QueryUrlFilteringOpenApiVO) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *QueryUrlFilteringOpenApiVO) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.


### GetSourceType

`func (o *QueryUrlFilteringOpenApiVO) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *QueryUrlFilteringOpenApiVO) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *QueryUrlFilteringOpenApiVO) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.


### GetStatus

`func (o *QueryUrlFilteringOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryUrlFilteringOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryUrlFilteringOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTimeRange

`func (o *QueryUrlFilteringOpenApiVO) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *QueryUrlFilteringOpenApiVO) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *QueryUrlFilteringOpenApiVO) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *QueryUrlFilteringOpenApiVO) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetType

`func (o *QueryUrlFilteringOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryUrlFilteringOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryUrlFilteringOpenApiVO) SetType(v string)`

SetType sets Type field to given value.


### GetUrls

`func (o *QueryUrlFilteringOpenApiVO) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *QueryUrlFilteringOpenApiVO) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *QueryUrlFilteringOpenApiVO) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *QueryUrlFilteringOpenApiVO) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


