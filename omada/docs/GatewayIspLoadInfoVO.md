# GatewayIspLoadInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IspInfo** | Pointer to [**IspInfoVO**](IspInfoVO.md) |  | [optional] 
**Mac** | Pointer to **string** | Gateway MAC. | [optional] 
**Name** | Pointer to **string** | Gateway name. | [optional] 
**SpeedTestLimit** | Pointer to **int64** | The limited bandwidth of the gateway speed test. | [optional] 
**Status** | Pointer to **int32** | Gateway status, should be a value as follows:-1 : disconnected1 : active0 : backup gateway under multiple gateways | [optional] 
**SupportSpeedTest** | Pointer to **bool** | Whether speed measurement is supported. | [optional] 

## Methods

### NewGatewayIspLoadInfoVO

`func NewGatewayIspLoadInfoVO() *GatewayIspLoadInfoVO`

NewGatewayIspLoadInfoVO instantiates a new GatewayIspLoadInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayIspLoadInfoVOWithDefaults

`func NewGatewayIspLoadInfoVOWithDefaults() *GatewayIspLoadInfoVO`

NewGatewayIspLoadInfoVOWithDefaults instantiates a new GatewayIspLoadInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIspInfo

`func (o *GatewayIspLoadInfoVO) GetIspInfo() IspInfoVO`

GetIspInfo returns the IspInfo field if non-nil, zero value otherwise.

### GetIspInfoOk

`func (o *GatewayIspLoadInfoVO) GetIspInfoOk() (*IspInfoVO, bool)`

GetIspInfoOk returns a tuple with the IspInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIspInfo

`func (o *GatewayIspLoadInfoVO) SetIspInfo(v IspInfoVO)`

SetIspInfo sets IspInfo field to given value.

### HasIspInfo

`func (o *GatewayIspLoadInfoVO) HasIspInfo() bool`

HasIspInfo returns a boolean if a field has been set.

### GetMac

`func (o *GatewayIspLoadInfoVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GatewayIspLoadInfoVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GatewayIspLoadInfoVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GatewayIspLoadInfoVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *GatewayIspLoadInfoVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayIspLoadInfoVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayIspLoadInfoVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayIspLoadInfoVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpeedTestLimit

`func (o *GatewayIspLoadInfoVO) GetSpeedTestLimit() int64`

GetSpeedTestLimit returns the SpeedTestLimit field if non-nil, zero value otherwise.

### GetSpeedTestLimitOk

`func (o *GatewayIspLoadInfoVO) GetSpeedTestLimitOk() (*int64, bool)`

GetSpeedTestLimitOk returns a tuple with the SpeedTestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedTestLimit

`func (o *GatewayIspLoadInfoVO) SetSpeedTestLimit(v int64)`

SetSpeedTestLimit sets SpeedTestLimit field to given value.

### HasSpeedTestLimit

`func (o *GatewayIspLoadInfoVO) HasSpeedTestLimit() bool`

HasSpeedTestLimit returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayIspLoadInfoVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayIspLoadInfoVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayIspLoadInfoVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GatewayIspLoadInfoVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportSpeedTest

`func (o *GatewayIspLoadInfoVO) GetSupportSpeedTest() bool`

GetSupportSpeedTest returns the SupportSpeedTest field if non-nil, zero value otherwise.

### GetSupportSpeedTestOk

`func (o *GatewayIspLoadInfoVO) GetSupportSpeedTestOk() (*bool, bool)`

GetSupportSpeedTestOk returns a tuple with the SupportSpeedTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSpeedTest

`func (o *GatewayIspLoadInfoVO) SetSupportSpeedTest(v bool)`

SetSupportSpeedTest sets SupportSpeedTest field to given value.

### HasSupportSpeedTest

`func (o *GatewayIspLoadInfoVO) HasSupportSpeedTest() bool`

HasSupportSpeedTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


