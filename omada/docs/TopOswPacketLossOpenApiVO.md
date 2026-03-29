# TopOswPacketLossOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceSeriesType** | Pointer to **int32** | Device series type should be a value as follows: 0:advanced, 1:pro | [optional] 
**HealthScore** | Pointer to **int32** | Health score | [optional] 
**LossPkts** | Pointer to **int64** | Total number of lost packets | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Device model | [optional] 
**ModelVersion** | Pointer to **string** | Device model version | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**PktsLossPortCnt** | Pointer to **int32** | Number of Ports with Packet Loss | [optional] 
**PortCnt** | Pointer to **int32** | Total number of ports | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 

## Methods

### NewTopOswPacketLossOpenApiVO

`func NewTopOswPacketLossOpenApiVO() *TopOswPacketLossOpenApiVO`

NewTopOswPacketLossOpenApiVO instantiates a new TopOswPacketLossOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopOswPacketLossOpenApiVOWithDefaults

`func NewTopOswPacketLossOpenApiVOWithDefaults() *TopOswPacketLossOpenApiVO`

NewTopOswPacketLossOpenApiVOWithDefaults instantiates a new TopOswPacketLossOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceSeriesType

`func (o *TopOswPacketLossOpenApiVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *TopOswPacketLossOpenApiVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *TopOswPacketLossOpenApiVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *TopOswPacketLossOpenApiVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetHealthScore

`func (o *TopOswPacketLossOpenApiVO) GetHealthScore() int32`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *TopOswPacketLossOpenApiVO) GetHealthScoreOk() (*int32, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *TopOswPacketLossOpenApiVO) SetHealthScore(v int32)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *TopOswPacketLossOpenApiVO) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetLossPkts

`func (o *TopOswPacketLossOpenApiVO) GetLossPkts() int64`

GetLossPkts returns the LossPkts field if non-nil, zero value otherwise.

### GetLossPktsOk

`func (o *TopOswPacketLossOpenApiVO) GetLossPktsOk() (*int64, bool)`

GetLossPktsOk returns a tuple with the LossPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossPkts

`func (o *TopOswPacketLossOpenApiVO) SetLossPkts(v int64)`

SetLossPkts sets LossPkts field to given value.

### HasLossPkts

`func (o *TopOswPacketLossOpenApiVO) HasLossPkts() bool`

HasLossPkts returns a boolean if a field has been set.

### GetMac

`func (o *TopOswPacketLossOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *TopOswPacketLossOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *TopOswPacketLossOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *TopOswPacketLossOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *TopOswPacketLossOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TopOswPacketLossOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TopOswPacketLossOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TopOswPacketLossOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *TopOswPacketLossOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *TopOswPacketLossOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *TopOswPacketLossOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *TopOswPacketLossOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *TopOswPacketLossOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopOswPacketLossOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopOswPacketLossOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopOswPacketLossOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPktsLossPortCnt

`func (o *TopOswPacketLossOpenApiVO) GetPktsLossPortCnt() int32`

GetPktsLossPortCnt returns the PktsLossPortCnt field if non-nil, zero value otherwise.

### GetPktsLossPortCntOk

`func (o *TopOswPacketLossOpenApiVO) GetPktsLossPortCntOk() (*int32, bool)`

GetPktsLossPortCntOk returns a tuple with the PktsLossPortCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPktsLossPortCnt

`func (o *TopOswPacketLossOpenApiVO) SetPktsLossPortCnt(v int32)`

SetPktsLossPortCnt sets PktsLossPortCnt field to given value.

### HasPktsLossPortCnt

`func (o *TopOswPacketLossOpenApiVO) HasPktsLossPortCnt() bool`

HasPktsLossPortCnt returns a boolean if a field has been set.

### GetPortCnt

`func (o *TopOswPacketLossOpenApiVO) GetPortCnt() int32`

GetPortCnt returns the PortCnt field if non-nil, zero value otherwise.

### GetPortCntOk

`func (o *TopOswPacketLossOpenApiVO) GetPortCntOk() (*int32, bool)`

GetPortCntOk returns a tuple with the PortCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCnt

`func (o *TopOswPacketLossOpenApiVO) SetPortCnt(v int32)`

SetPortCnt sets PortCnt field to given value.

### HasPortCnt

`func (o *TopOswPacketLossOpenApiVO) HasPortCnt() bool`

HasPortCnt returns a boolean if a field has been set.

### GetType

`func (o *TopOswPacketLossOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopOswPacketLossOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopOswPacketLossOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopOswPacketLossOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


