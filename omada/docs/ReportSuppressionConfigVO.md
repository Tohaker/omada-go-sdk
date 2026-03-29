# ReportSuppressionConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceList** | Pointer to [**[]MulticastExceptDeviceVO**](MulticastExceptDeviceVO.md) | report suppression except device list | [optional] 
**Enable** | Pointer to **bool** | enable | [optional] 

## Methods

### NewReportSuppressionConfigVO

`func NewReportSuppressionConfigVO() *ReportSuppressionConfigVO`

NewReportSuppressionConfigVO instantiates a new ReportSuppressionConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportSuppressionConfigVOWithDefaults

`func NewReportSuppressionConfigVOWithDefaults() *ReportSuppressionConfigVO`

NewReportSuppressionConfigVOWithDefaults instantiates a new ReportSuppressionConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceList

`func (o *ReportSuppressionConfigVO) GetDeviceList() []MulticastExceptDeviceVO`

GetDeviceList returns the DeviceList field if non-nil, zero value otherwise.

### GetDeviceListOk

`func (o *ReportSuppressionConfigVO) GetDeviceListOk() (*[]MulticastExceptDeviceVO, bool)`

GetDeviceListOk returns a tuple with the DeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceList

`func (o *ReportSuppressionConfigVO) SetDeviceList(v []MulticastExceptDeviceVO)`

SetDeviceList sets DeviceList field to given value.

### HasDeviceList

`func (o *ReportSuppressionConfigVO) HasDeviceList() bool`

HasDeviceList returns a boolean if a field has been set.

### GetEnable

`func (o *ReportSuppressionConfigVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ReportSuppressionConfigVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ReportSuppressionConfigVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ReportSuppressionConfigVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


