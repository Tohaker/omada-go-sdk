# ReportCardQueryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | Pointer to [**[]TabCardVO**](TabCardVO.md) | card info | [optional] 
**End** | Pointer to **int64** | end time, unit: seconds | [optional] 
**OmadacId** | Pointer to **string** | omadacId | [optional] 
**RequestToken** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** | site id | [optional] 
**Start** | Pointer to **int64** | start time, unit: seconds | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewReportCardQueryVO

`func NewReportCardQueryVO() *ReportCardQueryVO`

NewReportCardQueryVO instantiates a new ReportCardQueryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCardQueryVOWithDefaults

`func NewReportCardQueryVOWithDefaults() *ReportCardQueryVO`

NewReportCardQueryVOWithDefaults instantiates a new ReportCardQueryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *ReportCardQueryVO) GetCards() []TabCardVO`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *ReportCardQueryVO) GetCardsOk() (*[]TabCardVO, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *ReportCardQueryVO) SetCards(v []TabCardVO)`

SetCards sets Cards field to given value.

### HasCards

`func (o *ReportCardQueryVO) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetEnd

`func (o *ReportCardQueryVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ReportCardQueryVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ReportCardQueryVO) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ReportCardQueryVO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetOmadacId

`func (o *ReportCardQueryVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *ReportCardQueryVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *ReportCardQueryVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *ReportCardQueryVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetRequestToken

`func (o *ReportCardQueryVO) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *ReportCardQueryVO) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *ReportCardQueryVO) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *ReportCardQueryVO) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### GetSiteId

`func (o *ReportCardQueryVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ReportCardQueryVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ReportCardQueryVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ReportCardQueryVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStart

`func (o *ReportCardQueryVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ReportCardQueryVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ReportCardQueryVO) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *ReportCardQueryVO) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetToken

`func (o *ReportCardQueryVO) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ReportCardQueryVO) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ReportCardQueryVO) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ReportCardQueryVO) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


