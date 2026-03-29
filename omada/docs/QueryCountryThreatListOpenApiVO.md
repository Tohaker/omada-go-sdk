# QueryCountryThreatListOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Country. | [optional] 
**CurrentPage** | Pointer to **int32** | Current Page. | [optional] 
**CurrentPageSize** | Pointer to **int32** | Current Page Size. | [optional] 
**EndTime** | Pointer to **int64** | End Time. | [optional] 
**SearchField** | Pointer to **string** | Specify the search domain, separated by commas, such as \&quot;clientMac,clientName\&quot;.Can be null, in which case the default search domain should be used. | [optional] 
**SearchKey** | Pointer to **string** | Search Key. | [optional] 
**Severity** | Pointer to **string** | Displays the number of threats at different levels,0 - very high，1- high, 2- medium，3 - low. | [optional] 
**Sites** | Pointer to **string** | List of siteIds, separated by commas (,). All sites are returned if they are not passed. | [optional] 
**Sorts** | Pointer to **string** | sorts. | [optional] 
**StartTime** | Pointer to **int64** | The timestamp for the start time is in milliseconds. | [optional] 

## Methods

### NewQueryCountryThreatListOpenApiVO

`func NewQueryCountryThreatListOpenApiVO() *QueryCountryThreatListOpenApiVO`

NewQueryCountryThreatListOpenApiVO instantiates a new QueryCountryThreatListOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCountryThreatListOpenApiVOWithDefaults

`func NewQueryCountryThreatListOpenApiVOWithDefaults() *QueryCountryThreatListOpenApiVO`

NewQueryCountryThreatListOpenApiVOWithDefaults instantiates a new QueryCountryThreatListOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *QueryCountryThreatListOpenApiVO) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *QueryCountryThreatListOpenApiVO) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *QueryCountryThreatListOpenApiVO) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *QueryCountryThreatListOpenApiVO) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrentPage

`func (o *QueryCountryThreatListOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *QueryCountryThreatListOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *QueryCountryThreatListOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *QueryCountryThreatListOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentPageSize

`func (o *QueryCountryThreatListOpenApiVO) GetCurrentPageSize() int32`

GetCurrentPageSize returns the CurrentPageSize field if non-nil, zero value otherwise.

### GetCurrentPageSizeOk

`func (o *QueryCountryThreatListOpenApiVO) GetCurrentPageSizeOk() (*int32, bool)`

GetCurrentPageSizeOk returns a tuple with the CurrentPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageSize

`func (o *QueryCountryThreatListOpenApiVO) SetCurrentPageSize(v int32)`

SetCurrentPageSize sets CurrentPageSize field to given value.

### HasCurrentPageSize

`func (o *QueryCountryThreatListOpenApiVO) HasCurrentPageSize() bool`

HasCurrentPageSize returns a boolean if a field has been set.

### GetEndTime

`func (o *QueryCountryThreatListOpenApiVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *QueryCountryThreatListOpenApiVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *QueryCountryThreatListOpenApiVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *QueryCountryThreatListOpenApiVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetSearchField

`func (o *QueryCountryThreatListOpenApiVO) GetSearchField() string`

GetSearchField returns the SearchField field if non-nil, zero value otherwise.

### GetSearchFieldOk

`func (o *QueryCountryThreatListOpenApiVO) GetSearchFieldOk() (*string, bool)`

GetSearchFieldOk returns a tuple with the SearchField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchField

`func (o *QueryCountryThreatListOpenApiVO) SetSearchField(v string)`

SetSearchField sets SearchField field to given value.

### HasSearchField

`func (o *QueryCountryThreatListOpenApiVO) HasSearchField() bool`

HasSearchField returns a boolean if a field has been set.

### GetSearchKey

`func (o *QueryCountryThreatListOpenApiVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *QueryCountryThreatListOpenApiVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *QueryCountryThreatListOpenApiVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *QueryCountryThreatListOpenApiVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSeverity

`func (o *QueryCountryThreatListOpenApiVO) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *QueryCountryThreatListOpenApiVO) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *QueryCountryThreatListOpenApiVO) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *QueryCountryThreatListOpenApiVO) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSites

`func (o *QueryCountryThreatListOpenApiVO) GetSites() string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *QueryCountryThreatListOpenApiVO) GetSitesOk() (*string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *QueryCountryThreatListOpenApiVO) SetSites(v string)`

SetSites sets Sites field to given value.

### HasSites

`func (o *QueryCountryThreatListOpenApiVO) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetSorts

`func (o *QueryCountryThreatListOpenApiVO) GetSorts() string`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *QueryCountryThreatListOpenApiVO) GetSortsOk() (*string, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *QueryCountryThreatListOpenApiVO) SetSorts(v string)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *QueryCountryThreatListOpenApiVO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.

### GetStartTime

`func (o *QueryCountryThreatListOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *QueryCountryThreatListOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *QueryCountryThreatListOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *QueryCountryThreatListOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


