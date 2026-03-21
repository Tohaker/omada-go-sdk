# OspfProcessOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaList** | [**[]OspfProcessAreaOpenApiVO**](OspfProcessAreaOpenApiVO.md) |  Up to 16 entries are allowed for the areaList. | 
**ConnectedEnable** | **bool** | Direct connection routing protocol switch | 
**ConnectedMetric** | Pointer to **int32** | Set the metric value to be used as the metric of redistributed routes. It should be within the range of 1-16777214 and the default is equal to Default Metric configured on Basic page. | [optional] 
**ConnectedMetricType** | Pointer to **int32** | Set the OSPF metric type of redistributed routes. The default is External Type 2. | [optional] 
**DeviceName** | **string** | The name of the device | 
**Id** | Pointer to **string** | OSPF Process ID | [optional] 
**IsStack** | Pointer to **bool** | Indicates whether the device is a stack member device. | [optional] 
**Mac** | **string** | Device Mac | 
**ProcessId** | **int32** | Process ID should be within the range of 1-65535. | 
**RouterId** | Pointer to **string** | Router ID | [optional] 
**RouterMode** | **int32** | RouterMode, its value should be a value as follows: 0: AUTO, 1: MANUAL. | 
**StackId** | Pointer to **string** | Used for backend verification, the front-end does not need to pass a value, and there are values when the MAC is master | [optional] 
**StaticEnable** | **bool** | Static routing protocol switch | 
**StaticMetric** | Pointer to **int32** | Set the metric value to be used as the metric of redistributed routes. It should be within the range of 1-16777214 and the default is equal to Default Metric configured on Basic page. | [optional] 
**StaticMetricType** | Pointer to **int32** | Set the OSPF metric type of redistributed routes. The default is External Type 2. | [optional] 

## Methods

### NewOspfProcessOpenApiVO

`func NewOspfProcessOpenApiVO(areaList []OspfProcessAreaOpenApiVO, connectedEnable bool, deviceName string, mac string, processId int32, routerMode int32, staticEnable bool, ) *OspfProcessOpenApiVO`

NewOspfProcessOpenApiVO instantiates a new OspfProcessOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfProcessOpenApiVOWithDefaults

`func NewOspfProcessOpenApiVOWithDefaults() *OspfProcessOpenApiVO`

NewOspfProcessOpenApiVOWithDefaults instantiates a new OspfProcessOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaList

`func (o *OspfProcessOpenApiVO) GetAreaList() []OspfProcessAreaOpenApiVO`

GetAreaList returns the AreaList field if non-nil, zero value otherwise.

### GetAreaListOk

`func (o *OspfProcessOpenApiVO) GetAreaListOk() (*[]OspfProcessAreaOpenApiVO, bool)`

GetAreaListOk returns a tuple with the AreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaList

`func (o *OspfProcessOpenApiVO) SetAreaList(v []OspfProcessAreaOpenApiVO)`

SetAreaList sets AreaList field to given value.


### GetConnectedEnable

`func (o *OspfProcessOpenApiVO) GetConnectedEnable() bool`

GetConnectedEnable returns the ConnectedEnable field if non-nil, zero value otherwise.

### GetConnectedEnableOk

`func (o *OspfProcessOpenApiVO) GetConnectedEnableOk() (*bool, bool)`

GetConnectedEnableOk returns a tuple with the ConnectedEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEnable

`func (o *OspfProcessOpenApiVO) SetConnectedEnable(v bool)`

SetConnectedEnable sets ConnectedEnable field to given value.


### GetConnectedMetric

`func (o *OspfProcessOpenApiVO) GetConnectedMetric() int32`

GetConnectedMetric returns the ConnectedMetric field if non-nil, zero value otherwise.

### GetConnectedMetricOk

`func (o *OspfProcessOpenApiVO) GetConnectedMetricOk() (*int32, bool)`

GetConnectedMetricOk returns a tuple with the ConnectedMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedMetric

`func (o *OspfProcessOpenApiVO) SetConnectedMetric(v int32)`

SetConnectedMetric sets ConnectedMetric field to given value.

### HasConnectedMetric

`func (o *OspfProcessOpenApiVO) HasConnectedMetric() bool`

HasConnectedMetric returns a boolean if a field has been set.

### GetConnectedMetricType

`func (o *OspfProcessOpenApiVO) GetConnectedMetricType() int32`

GetConnectedMetricType returns the ConnectedMetricType field if non-nil, zero value otherwise.

### GetConnectedMetricTypeOk

`func (o *OspfProcessOpenApiVO) GetConnectedMetricTypeOk() (*int32, bool)`

GetConnectedMetricTypeOk returns a tuple with the ConnectedMetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedMetricType

`func (o *OspfProcessOpenApiVO) SetConnectedMetricType(v int32)`

SetConnectedMetricType sets ConnectedMetricType field to given value.

### HasConnectedMetricType

`func (o *OspfProcessOpenApiVO) HasConnectedMetricType() bool`

HasConnectedMetricType returns a boolean if a field has been set.

### GetDeviceName

`func (o *OspfProcessOpenApiVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *OspfProcessOpenApiVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *OspfProcessOpenApiVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetId

`func (o *OspfProcessOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OspfProcessOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OspfProcessOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OspfProcessOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsStack

`func (o *OspfProcessOpenApiVO) GetIsStack() bool`

GetIsStack returns the IsStack field if non-nil, zero value otherwise.

### GetIsStackOk

`func (o *OspfProcessOpenApiVO) GetIsStackOk() (*bool, bool)`

GetIsStackOk returns a tuple with the IsStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStack

`func (o *OspfProcessOpenApiVO) SetIsStack(v bool)`

SetIsStack sets IsStack field to given value.

### HasIsStack

`func (o *OspfProcessOpenApiVO) HasIsStack() bool`

HasIsStack returns a boolean if a field has been set.

### GetMac

`func (o *OspfProcessOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OspfProcessOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OspfProcessOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetProcessId

`func (o *OspfProcessOpenApiVO) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *OspfProcessOpenApiVO) GetProcessIdOk() (*int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *OspfProcessOpenApiVO) SetProcessId(v int32)`

SetProcessId sets ProcessId field to given value.


### GetRouterId

`func (o *OspfProcessOpenApiVO) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *OspfProcessOpenApiVO) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *OspfProcessOpenApiVO) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *OspfProcessOpenApiVO) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetRouterMode

`func (o *OspfProcessOpenApiVO) GetRouterMode() int32`

GetRouterMode returns the RouterMode field if non-nil, zero value otherwise.

### GetRouterModeOk

`func (o *OspfProcessOpenApiVO) GetRouterModeOk() (*int32, bool)`

GetRouterModeOk returns a tuple with the RouterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterMode

`func (o *OspfProcessOpenApiVO) SetRouterMode(v int32)`

SetRouterMode sets RouterMode field to given value.


### GetStackId

`func (o *OspfProcessOpenApiVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OspfProcessOpenApiVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OspfProcessOpenApiVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OspfProcessOpenApiVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStaticEnable

`func (o *OspfProcessOpenApiVO) GetStaticEnable() bool`

GetStaticEnable returns the StaticEnable field if non-nil, zero value otherwise.

### GetStaticEnableOk

`func (o *OspfProcessOpenApiVO) GetStaticEnableOk() (*bool, bool)`

GetStaticEnableOk returns a tuple with the StaticEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticEnable

`func (o *OspfProcessOpenApiVO) SetStaticEnable(v bool)`

SetStaticEnable sets StaticEnable field to given value.


### GetStaticMetric

`func (o *OspfProcessOpenApiVO) GetStaticMetric() int32`

GetStaticMetric returns the StaticMetric field if non-nil, zero value otherwise.

### GetStaticMetricOk

`func (o *OspfProcessOpenApiVO) GetStaticMetricOk() (*int32, bool)`

GetStaticMetricOk returns a tuple with the StaticMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMetric

`func (o *OspfProcessOpenApiVO) SetStaticMetric(v int32)`

SetStaticMetric sets StaticMetric field to given value.

### HasStaticMetric

`func (o *OspfProcessOpenApiVO) HasStaticMetric() bool`

HasStaticMetric returns a boolean if a field has been set.

### GetStaticMetricType

`func (o *OspfProcessOpenApiVO) GetStaticMetricType() int32`

GetStaticMetricType returns the StaticMetricType field if non-nil, zero value otherwise.

### GetStaticMetricTypeOk

`func (o *OspfProcessOpenApiVO) GetStaticMetricTypeOk() (*int32, bool)`

GetStaticMetricTypeOk returns a tuple with the StaticMetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMetricType

`func (o *OspfProcessOpenApiVO) SetStaticMetricType(v int32)`

SetStaticMetricType sets StaticMetricType field to given value.

### HasStaticMetricType

`func (o *OspfProcessOpenApiVO) HasStaticMetricType() bool`

HasStaticMetricType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


