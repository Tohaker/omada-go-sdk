# OspfProcessConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaList** | [**[]OspfProcessAreaOpenApiVO**](OspfProcessAreaOpenApiVO.md) |  Up to 16 entries are allowed for the areaList. | 
**ConnectedEnable** | **bool** | Direct connection routing protocol switch | 
**ConnectedMetric** | Pointer to **int32** | Set the metric value to be used as the metric of redistributed routes. It should be within the range of 1-16777214 and the default is equal to Default Metric configured on Basic page. | [optional] 
**ConnectedMetricType** | Pointer to **int32** | Set the OSPF metric type of redistributed routes. It should be a value as follows: 1: External Type 1, 2: External Type 2. The default is External Type 2. | [optional] 
**DeviceName** | **string** | The name of the device | 
**Mac** | **string** | Device Mac | 
**ProcessId** | **int32** | Process ID should be within the range of 1-65535. | 
**RouterId** | Pointer to **string** | Router ID | [optional] 
**RouterMode** | **int32** | RouterMode, its value should be a value as follows: 0: AUTO, 1: MANUAL. | 
**StaticEnable** | **bool** | Static routing protocol switch | 
**StaticMetric** | Pointer to **int32** | Set the metric value to be used as the metric of redistributed routes. It should be within the range of 1-16777214 and the default is equal to Default Metric configured on Basic page. | [optional] 
**StaticMetricType** | Pointer to **int32** | Set the OSPF metric type of redistributed routes. It should be a value as follows: 1: External Type 1, 2: External Type 2. The default is External Type 2. | [optional] 

## Methods

### NewOspfProcessConfigOpenApiVO

`func NewOspfProcessConfigOpenApiVO(areaList []OspfProcessAreaOpenApiVO, connectedEnable bool, deviceName string, mac string, processId int32, routerMode int32, staticEnable bool, ) *OspfProcessConfigOpenApiVO`

NewOspfProcessConfigOpenApiVO instantiates a new OspfProcessConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfProcessConfigOpenApiVOWithDefaults

`func NewOspfProcessConfigOpenApiVOWithDefaults() *OspfProcessConfigOpenApiVO`

NewOspfProcessConfigOpenApiVOWithDefaults instantiates a new OspfProcessConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaList

`func (o *OspfProcessConfigOpenApiVO) GetAreaList() []OspfProcessAreaOpenApiVO`

GetAreaList returns the AreaList field if non-nil, zero value otherwise.

### GetAreaListOk

`func (o *OspfProcessConfigOpenApiVO) GetAreaListOk() (*[]OspfProcessAreaOpenApiVO, bool)`

GetAreaListOk returns a tuple with the AreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaList

`func (o *OspfProcessConfigOpenApiVO) SetAreaList(v []OspfProcessAreaOpenApiVO)`

SetAreaList sets AreaList field to given value.


### GetConnectedEnable

`func (o *OspfProcessConfigOpenApiVO) GetConnectedEnable() bool`

GetConnectedEnable returns the ConnectedEnable field if non-nil, zero value otherwise.

### GetConnectedEnableOk

`func (o *OspfProcessConfigOpenApiVO) GetConnectedEnableOk() (*bool, bool)`

GetConnectedEnableOk returns a tuple with the ConnectedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEnable

`func (o *OspfProcessConfigOpenApiVO) SetConnectedEnable(v bool)`

SetConnectedEnable sets ConnectedEnable field to given value.


### GetConnectedMetric

`func (o *OspfProcessConfigOpenApiVO) GetConnectedMetric() int32`

GetConnectedMetric returns the ConnectedMetric field if non-nil, zero value otherwise.

### GetConnectedMetricOk

`func (o *OspfProcessConfigOpenApiVO) GetConnectedMetricOk() (*int32, bool)`

GetConnectedMetricOk returns a tuple with the ConnectedMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedMetric

`func (o *OspfProcessConfigOpenApiVO) SetConnectedMetric(v int32)`

SetConnectedMetric sets ConnectedMetric field to given value.

### HasConnectedMetric

`func (o *OspfProcessConfigOpenApiVO) HasConnectedMetric() bool`

HasConnectedMetric returns a boolean if a field has been set.

### GetConnectedMetricType

`func (o *OspfProcessConfigOpenApiVO) GetConnectedMetricType() int32`

GetConnectedMetricType returns the ConnectedMetricType field if non-nil, zero value otherwise.

### GetConnectedMetricTypeOk

`func (o *OspfProcessConfigOpenApiVO) GetConnectedMetricTypeOk() (*int32, bool)`

GetConnectedMetricTypeOk returns a tuple with the ConnectedMetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedMetricType

`func (o *OspfProcessConfigOpenApiVO) SetConnectedMetricType(v int32)`

SetConnectedMetricType sets ConnectedMetricType field to given value.

### HasConnectedMetricType

`func (o *OspfProcessConfigOpenApiVO) HasConnectedMetricType() bool`

HasConnectedMetricType returns a boolean if a field has been set.

### GetDeviceName

`func (o *OspfProcessConfigOpenApiVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *OspfProcessConfigOpenApiVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *OspfProcessConfigOpenApiVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetMac

`func (o *OspfProcessConfigOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OspfProcessConfigOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OspfProcessConfigOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetProcessId

`func (o *OspfProcessConfigOpenApiVO) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *OspfProcessConfigOpenApiVO) GetProcessIdOk() (*int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *OspfProcessConfigOpenApiVO) SetProcessId(v int32)`

SetProcessId sets ProcessId field to given value.


### GetRouterId

`func (o *OspfProcessConfigOpenApiVO) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *OspfProcessConfigOpenApiVO) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *OspfProcessConfigOpenApiVO) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *OspfProcessConfigOpenApiVO) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetRouterMode

`func (o *OspfProcessConfigOpenApiVO) GetRouterMode() int32`

GetRouterMode returns the RouterMode field if non-nil, zero value otherwise.

### GetRouterModeOk

`func (o *OspfProcessConfigOpenApiVO) GetRouterModeOk() (*int32, bool)`

GetRouterModeOk returns a tuple with the RouterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterMode

`func (o *OspfProcessConfigOpenApiVO) SetRouterMode(v int32)`

SetRouterMode sets RouterMode field to given value.


### GetStaticEnable

`func (o *OspfProcessConfigOpenApiVO) GetStaticEnable() bool`

GetStaticEnable returns the StaticEnable field if non-nil, zero value otherwise.

### GetStaticEnableOk

`func (o *OspfProcessConfigOpenApiVO) GetStaticEnableOk() (*bool, bool)`

GetStaticEnableOk returns a tuple with the StaticEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticEnable

`func (o *OspfProcessConfigOpenApiVO) SetStaticEnable(v bool)`

SetStaticEnable sets StaticEnable field to given value.


### GetStaticMetric

`func (o *OspfProcessConfigOpenApiVO) GetStaticMetric() int32`

GetStaticMetric returns the StaticMetric field if non-nil, zero value otherwise.

### GetStaticMetricOk

`func (o *OspfProcessConfigOpenApiVO) GetStaticMetricOk() (*int32, bool)`

GetStaticMetricOk returns a tuple with the StaticMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMetric

`func (o *OspfProcessConfigOpenApiVO) SetStaticMetric(v int32)`

SetStaticMetric sets StaticMetric field to given value.

### HasStaticMetric

`func (o *OspfProcessConfigOpenApiVO) HasStaticMetric() bool`

HasStaticMetric returns a boolean if a field has been set.

### GetStaticMetricType

`func (o *OspfProcessConfigOpenApiVO) GetStaticMetricType() int32`

GetStaticMetricType returns the StaticMetricType field if non-nil, zero value otherwise.

### GetStaticMetricTypeOk

`func (o *OspfProcessConfigOpenApiVO) GetStaticMetricTypeOk() (*int32, bool)`

GetStaticMetricTypeOk returns a tuple with the StaticMetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMetricType

`func (o *OspfProcessConfigOpenApiVO) SetStaticMetricType(v int32)`

SetStaticMetricType sets StaticMetricType field to given value.

### HasStaticMetricType

`func (o *OspfProcessConfigOpenApiVO) HasStaticMetricType() bool`

HasStaticMetricType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


