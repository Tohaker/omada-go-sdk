# ReportTab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelCards** | Pointer to **[]string** | cards be cancelled | [optional] 
**CardGroupList** | Pointer to [**[]CardGroupVO**](CardGroupVO.md) | list of card group | [optional] 
**Cards** | Pointer to **[]string** | list of cards in tab | [optional] 
**DefaultKey** | Pointer to **string** | defaultKey | [optional] 
**DefaultTab** | Pointer to **string** | defaultTab | [optional] 
**Name** | Pointer to **string** | tab name | [optional] 
**Rank** | Pointer to **int32** | tab rank | [optional] 
**Status** | Pointer to **int32** | tab status, 0: disable, 1: enable | [optional] 
**TabId** | Pointer to **string** | tab id | [optional] 
**Type** | Pointer to **int32** | tab type, default 0, customize 1 | [optional] 

## Methods

### NewReportTab

`func NewReportTab() *ReportTab`

NewReportTab instantiates a new ReportTab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportTabWithDefaults

`func NewReportTabWithDefaults() *ReportTab`

NewReportTabWithDefaults instantiates a new ReportTab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelCards

`func (o *ReportTab) GetCancelCards() []string`

GetCancelCards returns the CancelCards field if non-nil, zero value otherwise.

### GetCancelCardsOk

`func (o *ReportTab) GetCancelCardsOk() (*[]string, bool)`

GetCancelCardsOk returns a tuple with the CancelCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelCards

`func (o *ReportTab) SetCancelCards(v []string)`

SetCancelCards sets CancelCards field to given value.

### HasCancelCards

`func (o *ReportTab) HasCancelCards() bool`

HasCancelCards returns a boolean if a field has been set.

### GetCardGroupList

`func (o *ReportTab) GetCardGroupList() []CardGroupVO`

GetCardGroupList returns the CardGroupList field if non-nil, zero value otherwise.

### GetCardGroupListOk

`func (o *ReportTab) GetCardGroupListOk() (*[]CardGroupVO, bool)`

GetCardGroupListOk returns a tuple with the CardGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardGroupList

`func (o *ReportTab) SetCardGroupList(v []CardGroupVO)`

SetCardGroupList sets CardGroupList field to given value.

### HasCardGroupList

`func (o *ReportTab) HasCardGroupList() bool`

HasCardGroupList returns a boolean if a field has been set.

### GetCards

`func (o *ReportTab) GetCards() []string`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *ReportTab) GetCardsOk() (*[]string, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *ReportTab) SetCards(v []string)`

SetCards sets Cards field to given value.

### HasCards

`func (o *ReportTab) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetDefaultKey

`func (o *ReportTab) GetDefaultKey() string`

GetDefaultKey returns the DefaultKey field if non-nil, zero value otherwise.

### GetDefaultKeyOk

`func (o *ReportTab) GetDefaultKeyOk() (*string, bool)`

GetDefaultKeyOk returns a tuple with the DefaultKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKey

`func (o *ReportTab) SetDefaultKey(v string)`

SetDefaultKey sets DefaultKey field to given value.

### HasDefaultKey

`func (o *ReportTab) HasDefaultKey() bool`

HasDefaultKey returns a boolean if a field has been set.

### GetDefaultTab

`func (o *ReportTab) GetDefaultTab() string`

GetDefaultTab returns the DefaultTab field if non-nil, zero value otherwise.

### GetDefaultTabOk

`func (o *ReportTab) GetDefaultTabOk() (*string, bool)`

GetDefaultTabOk returns a tuple with the DefaultTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTab

`func (o *ReportTab) SetDefaultTab(v string)`

SetDefaultTab sets DefaultTab field to given value.

### HasDefaultTab

`func (o *ReportTab) HasDefaultTab() bool`

HasDefaultTab returns a boolean if a field has been set.

### GetName

`func (o *ReportTab) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportTab) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportTab) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportTab) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRank

`func (o *ReportTab) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ReportTab) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ReportTab) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *ReportTab) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetStatus

`func (o *ReportTab) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportTab) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportTab) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportTab) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTabId

`func (o *ReportTab) GetTabId() string`

GetTabId returns the TabId field if non-nil, zero value otherwise.

### GetTabIdOk

`func (o *ReportTab) GetTabIdOk() (*string, bool)`

GetTabIdOk returns a tuple with the TabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabId

`func (o *ReportTab) SetTabId(v string)`

SetTabId sets TabId field to given value.

### HasTabId

`func (o *ReportTab) HasTabId() bool`

HasTabId returns a boolean if a field has been set.

### GetType

`func (o *ReportTab) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportTab) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportTab) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ReportTab) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


