# UrlFilteringOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **map[string][]string** | categories of the URL filtering, TreeMap&lt;String, List&lt;String&gt;&gt; categories | [optional] 
**FilterMode** | **int32** | filterMode should be a value as follows: 0: URL; 1: category. | 
**Keywords** | Pointer to **[]string** | Keywords of the URL filtering. | [optional] 
**Mode** | **int32** | Mode should be a value as follows: 0: URL Path; 1: keywords(Only for gateway). | 
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

### NewUrlFilteringOpenApiVO

`func NewUrlFilteringOpenApiVO(filterMode int32, mode int32, name string, policy int32, sourceIds []string, sourceType int32, status bool, type_ string, ) *UrlFilteringOpenApiVO`

NewUrlFilteringOpenApiVO instantiates a new UrlFilteringOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlFilteringOpenApiVOWithDefaults

`func NewUrlFilteringOpenApiVOWithDefaults() *UrlFilteringOpenApiVO`

NewUrlFilteringOpenApiVOWithDefaults instantiates a new UrlFilteringOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *UrlFilteringOpenApiVO) GetCategories() map[string][]string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *UrlFilteringOpenApiVO) GetCategoriesOk() (*map[string][]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *UrlFilteringOpenApiVO) SetCategories(v map[string][]string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *UrlFilteringOpenApiVO) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetFilterMode

`func (o *UrlFilteringOpenApiVO) GetFilterMode() int32`

GetFilterMode returns the FilterMode field if non-nil, zero value otherwise.

### GetFilterModeOk

`func (o *UrlFilteringOpenApiVO) GetFilterModeOk() (*int32, bool)`

GetFilterModeOk returns a tuple with the FilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMode

`func (o *UrlFilteringOpenApiVO) SetFilterMode(v int32)`

SetFilterMode sets FilterMode field to given value.


### GetKeywords

`func (o *UrlFilteringOpenApiVO) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *UrlFilteringOpenApiVO) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *UrlFilteringOpenApiVO) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *UrlFilteringOpenApiVO) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetMode

`func (o *UrlFilteringOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UrlFilteringOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UrlFilteringOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetName

`func (o *UrlFilteringOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UrlFilteringOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UrlFilteringOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPolicy

`func (o *UrlFilteringOpenApiVO) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UrlFilteringOpenApiVO) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UrlFilteringOpenApiVO) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetScenarioMode

`func (o *UrlFilteringOpenApiVO) GetScenarioMode() int32`

GetScenarioMode returns the ScenarioMode field if non-nil, zero value otherwise.

### GetScenarioModeOk

`func (o *UrlFilteringOpenApiVO) GetScenarioModeOk() (*int32, bool)`

GetScenarioModeOk returns a tuple with the ScenarioMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioMode

`func (o *UrlFilteringOpenApiVO) SetScenarioMode(v int32)`

SetScenarioMode sets ScenarioMode field to given value.

### HasScenarioMode

`func (o *UrlFilteringOpenApiVO) HasScenarioMode() bool`

HasScenarioMode returns a boolean if a field has been set.

### GetSourceIds

`func (o *UrlFilteringOpenApiVO) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *UrlFilteringOpenApiVO) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *UrlFilteringOpenApiVO) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.


### GetSourceType

`func (o *UrlFilteringOpenApiVO) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *UrlFilteringOpenApiVO) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *UrlFilteringOpenApiVO) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.


### GetStatus

`func (o *UrlFilteringOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UrlFilteringOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UrlFilteringOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTimeRange

`func (o *UrlFilteringOpenApiVO) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *UrlFilteringOpenApiVO) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *UrlFilteringOpenApiVO) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *UrlFilteringOpenApiVO) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetType

`func (o *UrlFilteringOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UrlFilteringOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UrlFilteringOpenApiVO) SetType(v string)`

SetType sets Type field to given value.


### GetUrls

`func (o *UrlFilteringOpenApiVO) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *UrlFilteringOpenApiVO) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *UrlFilteringOpenApiVO) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *UrlFilteringOpenApiVO) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


