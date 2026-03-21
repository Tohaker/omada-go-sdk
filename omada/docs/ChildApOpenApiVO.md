# ChildApOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceSeriesType** | Pointer to **int32** | 0-advanced; 1-pro | [optional] 
**Ip** | Pointer to **string** | ip | [optional] 
**Mac** | Pointer to **string** | parent ap mac | [optional] 
**Model** | Pointer to **string** | model | [optional] 
**ModelVersion** | Pointer to **string** | model version | [optional] 
**Name** | Pointer to **string** | parent ap name | [optional] 
**Rssi** | Pointer to **int32** | rssi | [optional] 
**SupportSpeedTest** | Pointer to **bool** | bridge device support speed test | [optional] 
**Type** | Pointer to **string** | device type | [optional] 

## Methods

### NewChildApOpenApiVO

`func NewChildApOpenApiVO() *ChildApOpenApiVO`

NewChildApOpenApiVO instantiates a new ChildApOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildApOpenApiVOWithDefaults

`func NewChildApOpenApiVOWithDefaults() *ChildApOpenApiVO`

NewChildApOpenApiVOWithDefaults instantiates a new ChildApOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceSeriesType

`func (o *ChildApOpenApiVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *ChildApOpenApiVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *ChildApOpenApiVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *ChildApOpenApiVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetIp

`func (o *ChildApOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ChildApOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ChildApOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ChildApOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *ChildApOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ChildApOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ChildApOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ChildApOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ChildApOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChildApOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChildApOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChildApOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ChildApOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ChildApOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ChildApOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ChildApOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *ChildApOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChildApOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChildApOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChildApOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRssi

`func (o *ChildApOpenApiVO) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *ChildApOpenApiVO) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *ChildApOpenApiVO) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *ChildApOpenApiVO) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetSupportSpeedTest

`func (o *ChildApOpenApiVO) GetSupportSpeedTest() bool`

GetSupportSpeedTest returns the SupportSpeedTest field if non-nil, zero value otherwise.

### GetSupportSpeedTestOk

`func (o *ChildApOpenApiVO) GetSupportSpeedTestOk() (*bool, bool)`

GetSupportSpeedTestOk returns a tuple with the SupportSpeedTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSpeedTest

`func (o *ChildApOpenApiVO) SetSupportSpeedTest(v bool)`

SetSupportSpeedTest sets SupportSpeedTest field to given value.

### HasSupportSpeedTest

`func (o *ChildApOpenApiVO) HasSupportSpeedTest() bool`

HasSupportSpeedTest returns a boolean if a field has been set.

### GetType

`func (o *ChildApOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChildApOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChildApOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChildApOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


