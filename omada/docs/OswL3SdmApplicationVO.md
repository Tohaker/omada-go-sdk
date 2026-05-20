# OswL3SdmApplicationVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurTmplName** | Pointer to **string** | Current used sdm template. | [optional] 
**SdmResourceUsageList** | Pointer to [**[]SdmResourceUsage**](SdmResourceUsage.md) | The list contains usage details of SDM resources on the device, including used and available resources. | [optional] 
**TcamUtilization** | Pointer to **int32** | The overall TCAM utilization, calculated as usedTcam / totalTcam. | [optional] 

## Methods

### NewOswL3SdmApplicationVO

`func NewOswL3SdmApplicationVO() *OswL3SdmApplicationVO`

NewOswL3SdmApplicationVO instantiates a new OswL3SdmApplicationVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswL3SdmApplicationVOWithDefaults

`func NewOswL3SdmApplicationVOWithDefaults() *OswL3SdmApplicationVO`

NewOswL3SdmApplicationVOWithDefaults instantiates a new OswL3SdmApplicationVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurTmplName

`func (o *OswL3SdmApplicationVO) GetCurTmplName() string`

GetCurTmplName returns the CurTmplName field if non-nil, zero value otherwise.

### GetCurTmplNameOk

`func (o *OswL3SdmApplicationVO) GetCurTmplNameOk() (*string, bool)`

GetCurTmplNameOk returns a tuple with the CurTmplName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurTmplName

`func (o *OswL3SdmApplicationVO) SetCurTmplName(v string)`

SetCurTmplName sets CurTmplName field to given value.

### HasCurTmplName

`func (o *OswL3SdmApplicationVO) HasCurTmplName() bool`

HasCurTmplName returns a boolean if a field has been set.

### GetSdmResourceUsageList

`func (o *OswL3SdmApplicationVO) GetSdmResourceUsageList() []SdmResourceUsage`

GetSdmResourceUsageList returns the SdmResourceUsageList field if non-nil, zero value otherwise.

### GetSdmResourceUsageListOk

`func (o *OswL3SdmApplicationVO) GetSdmResourceUsageListOk() (*[]SdmResourceUsage, bool)`

GetSdmResourceUsageListOk returns a tuple with the SdmResourceUsageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdmResourceUsageList

`func (o *OswL3SdmApplicationVO) SetSdmResourceUsageList(v []SdmResourceUsage)`

SetSdmResourceUsageList sets SdmResourceUsageList field to given value.

### HasSdmResourceUsageList

`func (o *OswL3SdmApplicationVO) HasSdmResourceUsageList() bool`

HasSdmResourceUsageList returns a boolean if a field has been set.

### GetTcamUtilization

`func (o *OswL3SdmApplicationVO) GetTcamUtilization() int32`

GetTcamUtilization returns the TcamUtilization field if non-nil, zero value otherwise.

### GetTcamUtilizationOk

`func (o *OswL3SdmApplicationVO) GetTcamUtilizationOk() (*int32, bool)`

GetTcamUtilizationOk returns a tuple with the TcamUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcamUtilization

`func (o *OswL3SdmApplicationVO) SetTcamUtilization(v int32)`

SetTcamUtilization sets TcamUtilization field to given value.

### HasTcamUtilization

`func (o *OswL3SdmApplicationVO) HasTcamUtilization() bool`

HasTcamUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


