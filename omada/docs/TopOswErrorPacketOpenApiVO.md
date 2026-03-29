# TopOswErrorPacketOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceSeriesType** | Pointer to **int32** | Device series type should be a value as follows: 0:advanced, 1:pro | [optional] 
**ErrPkts** | Pointer to **int64** | Total number of error packets | [optional] 
**ErrPortCnt** | Pointer to **int32** | Number of Ports with error packets | [optional] 
**HealthScore** | Pointer to **int32** | Health score | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Device model | [optional] 
**ModelVersion** | Pointer to **string** | Device model version | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**PortCnt** | Pointer to **int32** | Total number of ports | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 

## Methods

### NewTopOswErrorPacketOpenApiVO

`func NewTopOswErrorPacketOpenApiVO() *TopOswErrorPacketOpenApiVO`

NewTopOswErrorPacketOpenApiVO instantiates a new TopOswErrorPacketOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopOswErrorPacketOpenApiVOWithDefaults

`func NewTopOswErrorPacketOpenApiVOWithDefaults() *TopOswErrorPacketOpenApiVO`

NewTopOswErrorPacketOpenApiVOWithDefaults instantiates a new TopOswErrorPacketOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceSeriesType

`func (o *TopOswErrorPacketOpenApiVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *TopOswErrorPacketOpenApiVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *TopOswErrorPacketOpenApiVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *TopOswErrorPacketOpenApiVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetErrPkts

`func (o *TopOswErrorPacketOpenApiVO) GetErrPkts() int64`

GetErrPkts returns the ErrPkts field if non-nil, zero value otherwise.

### GetErrPktsOk

`func (o *TopOswErrorPacketOpenApiVO) GetErrPktsOk() (*int64, bool)`

GetErrPktsOk returns a tuple with the ErrPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrPkts

`func (o *TopOswErrorPacketOpenApiVO) SetErrPkts(v int64)`

SetErrPkts sets ErrPkts field to given value.

### HasErrPkts

`func (o *TopOswErrorPacketOpenApiVO) HasErrPkts() bool`

HasErrPkts returns a boolean if a field has been set.

### GetErrPortCnt

`func (o *TopOswErrorPacketOpenApiVO) GetErrPortCnt() int32`

GetErrPortCnt returns the ErrPortCnt field if non-nil, zero value otherwise.

### GetErrPortCntOk

`func (o *TopOswErrorPacketOpenApiVO) GetErrPortCntOk() (*int32, bool)`

GetErrPortCntOk returns a tuple with the ErrPortCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrPortCnt

`func (o *TopOswErrorPacketOpenApiVO) SetErrPortCnt(v int32)`

SetErrPortCnt sets ErrPortCnt field to given value.

### HasErrPortCnt

`func (o *TopOswErrorPacketOpenApiVO) HasErrPortCnt() bool`

HasErrPortCnt returns a boolean if a field has been set.

### GetHealthScore

`func (o *TopOswErrorPacketOpenApiVO) GetHealthScore() int32`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *TopOswErrorPacketOpenApiVO) GetHealthScoreOk() (*int32, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *TopOswErrorPacketOpenApiVO) SetHealthScore(v int32)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *TopOswErrorPacketOpenApiVO) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetMac

`func (o *TopOswErrorPacketOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *TopOswErrorPacketOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *TopOswErrorPacketOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *TopOswErrorPacketOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *TopOswErrorPacketOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TopOswErrorPacketOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TopOswErrorPacketOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TopOswErrorPacketOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *TopOswErrorPacketOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *TopOswErrorPacketOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *TopOswErrorPacketOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *TopOswErrorPacketOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *TopOswErrorPacketOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopOswErrorPacketOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopOswErrorPacketOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopOswErrorPacketOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortCnt

`func (o *TopOswErrorPacketOpenApiVO) GetPortCnt() int32`

GetPortCnt returns the PortCnt field if non-nil, zero value otherwise.

### GetPortCntOk

`func (o *TopOswErrorPacketOpenApiVO) GetPortCntOk() (*int32, bool)`

GetPortCntOk returns a tuple with the PortCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCnt

`func (o *TopOswErrorPacketOpenApiVO) SetPortCnt(v int32)`

SetPortCnt sets PortCnt field to given value.

### HasPortCnt

`func (o *TopOswErrorPacketOpenApiVO) HasPortCnt() bool`

HasPortCnt returns a boolean if a field has been set.

### GetType

`func (o *TopOswErrorPacketOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopOswErrorPacketOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopOswErrorPacketOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopOswErrorPacketOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


