# IspVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadPercent** | Pointer to **float64** | Download utilization. | [optional] 
**DownloadSpeed** | Pointer to **string** | The current download speed. | [optional] 
**DownloadSpeedSet** | Pointer to **int64** | The set download speed. | [optional] 
**Ip** | Pointer to **string** | Port ip. | [optional] 
**LoadBalance** | Pointer to **string** | Port load balance. | [optional] 
**MaxBandwidth** | Pointer to **int64** | Port max bandwidth. | [optional] 
**Name** | Pointer to **string** | Port name. | [optional] 
**Port** | Pointer to **int32** | Port id. | [optional] 
**PortUuid** | Pointer to **string** | Port uuid. | [optional] 
**Status** | Pointer to **int32** | Isp status, should be a value as follows:2 : normal ISP1 : primary ISP0 : backup ISP | [optional] 
**UploadPercent** | Pointer to **float64** | Upload utilization. | [optional] 
**UploadSpeed** | Pointer to **string** | The current upload speed. | [optional] 
**UploadSpeedSet** | Pointer to **int64** | The set upload speed. | [optional] 

## Methods

### NewIspVO

`func NewIspVO() *IspVO`

NewIspVO instantiates a new IspVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIspVOWithDefaults

`func NewIspVOWithDefaults() *IspVO`

NewIspVOWithDefaults instantiates a new IspVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadPercent

`func (o *IspVO) GetDownloadPercent() float64`

GetDownloadPercent returns the DownloadPercent field if non-nil, zero value otherwise.

### GetDownloadPercentOk

`func (o *IspVO) GetDownloadPercentOk() (*float64, bool)`

GetDownloadPercentOk returns a tuple with the DownloadPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPercent

`func (o *IspVO) SetDownloadPercent(v float64)`

SetDownloadPercent sets DownloadPercent field to given value.

### HasDownloadPercent

`func (o *IspVO) HasDownloadPercent() bool`

HasDownloadPercent returns a boolean if a field has been set.

### GetDownloadSpeed

`func (o *IspVO) GetDownloadSpeed() string`

GetDownloadSpeed returns the DownloadSpeed field if non-nil, zero value otherwise.

### GetDownloadSpeedOk

`func (o *IspVO) GetDownloadSpeedOk() (*string, bool)`

GetDownloadSpeedOk returns a tuple with the DownloadSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadSpeed

`func (o *IspVO) SetDownloadSpeed(v string)`

SetDownloadSpeed sets DownloadSpeed field to given value.

### HasDownloadSpeed

`func (o *IspVO) HasDownloadSpeed() bool`

HasDownloadSpeed returns a boolean if a field has been set.

### GetDownloadSpeedSet

`func (o *IspVO) GetDownloadSpeedSet() int64`

GetDownloadSpeedSet returns the DownloadSpeedSet field if non-nil, zero value otherwise.

### GetDownloadSpeedSetOk

`func (o *IspVO) GetDownloadSpeedSetOk() (*int64, bool)`

GetDownloadSpeedSetOk returns a tuple with the DownloadSpeedSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadSpeedSet

`func (o *IspVO) SetDownloadSpeedSet(v int64)`

SetDownloadSpeedSet sets DownloadSpeedSet field to given value.

### HasDownloadSpeedSet

`func (o *IspVO) HasDownloadSpeedSet() bool`

HasDownloadSpeedSet returns a boolean if a field has been set.

### GetIp

`func (o *IspVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IspVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IspVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *IspVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLoadBalance

`func (o *IspVO) GetLoadBalance() string`

GetLoadBalance returns the LoadBalance field if non-nil, zero value otherwise.

### GetLoadBalanceOk

`func (o *IspVO) GetLoadBalanceOk() (*string, bool)`

GetLoadBalanceOk returns a tuple with the LoadBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalance

`func (o *IspVO) SetLoadBalance(v string)`

SetLoadBalance sets LoadBalance field to given value.

### HasLoadBalance

`func (o *IspVO) HasLoadBalance() bool`

HasLoadBalance returns a boolean if a field has been set.

### GetMaxBandwidth

`func (o *IspVO) GetMaxBandwidth() int64`

GetMaxBandwidth returns the MaxBandwidth field if non-nil, zero value otherwise.

### GetMaxBandwidthOk

`func (o *IspVO) GetMaxBandwidthOk() (*int64, bool)`

GetMaxBandwidthOk returns a tuple with the MaxBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBandwidth

`func (o *IspVO) SetMaxBandwidth(v int64)`

SetMaxBandwidth sets MaxBandwidth field to given value.

### HasMaxBandwidth

`func (o *IspVO) HasMaxBandwidth() bool`

HasMaxBandwidth returns a boolean if a field has been set.

### GetName

`func (o *IspVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IspVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IspVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IspVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *IspVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IspVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IspVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *IspVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortUuid

`func (o *IspVO) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *IspVO) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *IspVO) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.

### HasPortUuid

`func (o *IspVO) HasPortUuid() bool`

HasPortUuid returns a boolean if a field has been set.

### GetStatus

`func (o *IspVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IspVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IspVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IspVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUploadPercent

`func (o *IspVO) GetUploadPercent() float64`

GetUploadPercent returns the UploadPercent field if non-nil, zero value otherwise.

### GetUploadPercentOk

`func (o *IspVO) GetUploadPercentOk() (*float64, bool)`

GetUploadPercentOk returns a tuple with the UploadPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadPercent

`func (o *IspVO) SetUploadPercent(v float64)`

SetUploadPercent sets UploadPercent field to given value.

### HasUploadPercent

`func (o *IspVO) HasUploadPercent() bool`

HasUploadPercent returns a boolean if a field has been set.

### GetUploadSpeed

`func (o *IspVO) GetUploadSpeed() string`

GetUploadSpeed returns the UploadSpeed field if non-nil, zero value otherwise.

### GetUploadSpeedOk

`func (o *IspVO) GetUploadSpeedOk() (*string, bool)`

GetUploadSpeedOk returns a tuple with the UploadSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSpeed

`func (o *IspVO) SetUploadSpeed(v string)`

SetUploadSpeed sets UploadSpeed field to given value.

### HasUploadSpeed

`func (o *IspVO) HasUploadSpeed() bool`

HasUploadSpeed returns a boolean if a field has been set.

### GetUploadSpeedSet

`func (o *IspVO) GetUploadSpeedSet() int64`

GetUploadSpeedSet returns the UploadSpeedSet field if non-nil, zero value otherwise.

### GetUploadSpeedSetOk

`func (o *IspVO) GetUploadSpeedSetOk() (*int64, bool)`

GetUploadSpeedSetOk returns a tuple with the UploadSpeedSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSpeedSet

`func (o *IspVO) SetUploadSpeedSet(v int64)`

SetUploadSpeedSet sets UploadSpeedSet field to given value.

### HasUploadSpeedSet

`func (o *IspVO) HasUploadSpeedSet() bool`

HasUploadSpeedSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


