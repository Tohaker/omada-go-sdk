# ReportExportV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailList** | Pointer to **[]string** | The email list to send report | [optional] 
**End** | **int64** | The report end time | 
**ReportName** | **string** | The report name | 
**ReportType** | **int32** | The report type. 0 : pdf, 1 : csv. | 
**Start** | **int64** | The report start time | 
**TabIdList** | **[]string** | The tab need to export | 

## Methods

### NewReportExportV2

`func NewReportExportV2(end int64, reportName string, reportType int32, start int64, tabIdList []string, ) *ReportExportV2`

NewReportExportV2 instantiates a new ReportExportV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportExportV2WithDefaults

`func NewReportExportV2WithDefaults() *ReportExportV2`

NewReportExportV2WithDefaults instantiates a new ReportExportV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailList

`func (o *ReportExportV2) GetEmailList() []string`

GetEmailList returns the EmailList field if non-nil, zero value otherwise.

### GetEmailListOk

`func (o *ReportExportV2) GetEmailListOk() (*[]string, bool)`

GetEmailListOk returns a tuple with the EmailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailList

`func (o *ReportExportV2) SetEmailList(v []string)`

SetEmailList sets EmailList field to given value.

### HasEmailList

`func (o *ReportExportV2) HasEmailList() bool`

HasEmailList returns a boolean if a field has been set.

### GetEnd

`func (o *ReportExportV2) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ReportExportV2) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ReportExportV2) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetReportName

`func (o *ReportExportV2) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *ReportExportV2) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *ReportExportV2) SetReportName(v string)`

SetReportName sets ReportName field to given value.


### GetReportType

`func (o *ReportExportV2) GetReportType() int32`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportExportV2) GetReportTypeOk() (*int32, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportExportV2) SetReportType(v int32)`

SetReportType sets ReportType field to given value.


### GetStart

`func (o *ReportExportV2) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ReportExportV2) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ReportExportV2) SetStart(v int64)`

SetStart sets Start field to given value.


### GetTabIdList

`func (o *ReportExportV2) GetTabIdList() []string`

GetTabIdList returns the TabIdList field if non-nil, zero value otherwise.

### GetTabIdListOk

`func (o *ReportExportV2) GetTabIdListOk() (*[]string, bool)`

GetTabIdListOk returns a tuple with the TabIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabIdList

`func (o *ReportExportV2) SetTabIdList(v []string)`

SetTabIdList sets TabIdList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


