# TopCpuUsageHealthOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtil** | Pointer to **int32** | CPU utilization | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device series type should be a value as follows: 0:advanced, 1:pro | [optional] 
**HealthScore** | Pointer to **int32** | Health score | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Device model | [optional] 
**ModelVersion** | Pointer to **string** | Device model version | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 

## Methods

### NewTopCpuUsageHealthOpenApiVO

`func NewTopCpuUsageHealthOpenApiVO() *TopCpuUsageHealthOpenApiVO`

NewTopCpuUsageHealthOpenApiVO instantiates a new TopCpuUsageHealthOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopCpuUsageHealthOpenApiVOWithDefaults

`func NewTopCpuUsageHealthOpenApiVOWithDefaults() *TopCpuUsageHealthOpenApiVO`

NewTopCpuUsageHealthOpenApiVOWithDefaults instantiates a new TopCpuUsageHealthOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtil

`func (o *TopCpuUsageHealthOpenApiVO) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *TopCpuUsageHealthOpenApiVO) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *TopCpuUsageHealthOpenApiVO) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *TopCpuUsageHealthOpenApiVO) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *TopCpuUsageHealthOpenApiVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *TopCpuUsageHealthOpenApiVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *TopCpuUsageHealthOpenApiVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *TopCpuUsageHealthOpenApiVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetHealthScore

`func (o *TopCpuUsageHealthOpenApiVO) GetHealthScore() int32`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *TopCpuUsageHealthOpenApiVO) GetHealthScoreOk() (*int32, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *TopCpuUsageHealthOpenApiVO) SetHealthScore(v int32)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *TopCpuUsageHealthOpenApiVO) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetMac

`func (o *TopCpuUsageHealthOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *TopCpuUsageHealthOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *TopCpuUsageHealthOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *TopCpuUsageHealthOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *TopCpuUsageHealthOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TopCpuUsageHealthOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TopCpuUsageHealthOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TopCpuUsageHealthOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *TopCpuUsageHealthOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *TopCpuUsageHealthOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *TopCpuUsageHealthOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *TopCpuUsageHealthOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *TopCpuUsageHealthOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopCpuUsageHealthOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopCpuUsageHealthOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopCpuUsageHealthOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *TopCpuUsageHealthOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopCpuUsageHealthOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopCpuUsageHealthOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopCpuUsageHealthOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


