# VlanNetworkAffectingDeviceDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Es** | Pointer to **bool** | It indicates whether the device is easy managed switch. | [optional] 
**EsInfo** | Pointer to [**VlanNetworkAffectingEsDetailVO**](VlanNetworkAffectingEsDetailVO.md) |  | [optional] 
**GatewayInfo** | Pointer to [**VlanNetworkAffectingOsgDetailVO**](VlanNetworkAffectingOsgDetailVO.md) |  | [optional] 
**InternetInfo** | Pointer to [**VlanNetworkAffectingInternetDetailVO**](VlanNetworkAffectingInternetDetailVO.md) |  | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**StackId** | Pointer to **string** | Only valid when the device is stack. | [optional] 
**StackInfo** | Pointer to [**VlanNetworkAffectingStackDetailVO**](VlanNetworkAffectingStackDetailVO.md) |  | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**SwitchInfo** | Pointer to [**VlanNetworkAffectingSwitchDetailVO**](VlanNetworkAffectingSwitchDetailVO.md) |  | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 

## Methods

### NewVlanNetworkAffectingDeviceDetailVO

`func NewVlanNetworkAffectingDeviceDetailVO() *VlanNetworkAffectingDeviceDetailVO`

NewVlanNetworkAffectingDeviceDetailVO instantiates a new VlanNetworkAffectingDeviceDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanNetworkAffectingDeviceDetailVOWithDefaults

`func NewVlanNetworkAffectingDeviceDetailVOWithDefaults() *VlanNetworkAffectingDeviceDetailVO`

NewVlanNetworkAffectingDeviceDetailVOWithDefaults instantiates a new VlanNetworkAffectingDeviceDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEs

`func (o *VlanNetworkAffectingDeviceDetailVO) GetEs() bool`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetEsOk() (*bool, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *VlanNetworkAffectingDeviceDetailVO) SetEs(v bool)`

SetEs sets Es field to given value.

### HasEs

`func (o *VlanNetworkAffectingDeviceDetailVO) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetEsInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) GetEsInfo() VlanNetworkAffectingEsDetailVO`

GetEsInfo returns the EsInfo field if non-nil, zero value otherwise.

### GetEsInfoOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetEsInfoOk() (*VlanNetworkAffectingEsDetailVO, bool)`

GetEsInfoOk returns a tuple with the EsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) SetEsInfo(v VlanNetworkAffectingEsDetailVO)`

SetEsInfo sets EsInfo field to given value.

### HasEsInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) HasEsInfo() bool`

HasEsInfo returns a boolean if a field has been set.

### GetGatewayInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) GetGatewayInfo() VlanNetworkAffectingOsgDetailVO`

GetGatewayInfo returns the GatewayInfo field if non-nil, zero value otherwise.

### GetGatewayInfoOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetGatewayInfoOk() (*VlanNetworkAffectingOsgDetailVO, bool)`

GetGatewayInfoOk returns a tuple with the GatewayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) SetGatewayInfo(v VlanNetworkAffectingOsgDetailVO)`

SetGatewayInfo sets GatewayInfo field to given value.

### HasGatewayInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) HasGatewayInfo() bool`

HasGatewayInfo returns a boolean if a field has been set.

### GetInternetInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) GetInternetInfo() VlanNetworkAffectingInternetDetailVO`

GetInternetInfo returns the InternetInfo field if non-nil, zero value otherwise.

### GetInternetInfoOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetInternetInfoOk() (*VlanNetworkAffectingInternetDetailVO, bool)`

GetInternetInfoOk returns a tuple with the InternetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) SetInternetInfo(v VlanNetworkAffectingInternetDetailVO)`

SetInternetInfo sets InternetInfo field to given value.

### HasInternetInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) HasInternetInfo() bool`

HasInternetInfo returns a boolean if a field has been set.

### GetMac

`func (o *VlanNetworkAffectingDeviceDetailVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *VlanNetworkAffectingDeviceDetailVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *VlanNetworkAffectingDeviceDetailVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *VlanNetworkAffectingDeviceDetailVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VlanNetworkAffectingDeviceDetailVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VlanNetworkAffectingDeviceDetailVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *VlanNetworkAffectingDeviceDetailVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *VlanNetworkAffectingDeviceDetailVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *VlanNetworkAffectingDeviceDetailVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *VlanNetworkAffectingDeviceDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VlanNetworkAffectingDeviceDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VlanNetworkAffectingDeviceDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStackId

`func (o *VlanNetworkAffectingDeviceDetailVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *VlanNetworkAffectingDeviceDetailVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *VlanNetworkAffectingDeviceDetailVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) GetStackInfo() VlanNetworkAffectingStackDetailVO`

GetStackInfo returns the StackInfo field if non-nil, zero value otherwise.

### GetStackInfoOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetStackInfoOk() (*VlanNetworkAffectingStackDetailVO, bool)`

GetStackInfoOk returns a tuple with the StackInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) SetStackInfo(v VlanNetworkAffectingStackDetailVO)`

SetStackInfo sets StackInfo field to given value.

### HasStackInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) HasStackInfo() bool`

HasStackInfo returns a boolean if a field has been set.

### GetStatusCategory

`func (o *VlanNetworkAffectingDeviceDetailVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *VlanNetworkAffectingDeviceDetailVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *VlanNetworkAffectingDeviceDetailVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetSwitchInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) GetSwitchInfo() VlanNetworkAffectingSwitchDetailVO`

GetSwitchInfo returns the SwitchInfo field if non-nil, zero value otherwise.

### GetSwitchInfoOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetSwitchInfoOk() (*VlanNetworkAffectingSwitchDetailVO, bool)`

GetSwitchInfoOk returns a tuple with the SwitchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) SetSwitchInfo(v VlanNetworkAffectingSwitchDetailVO)`

SetSwitchInfo sets SwitchInfo field to given value.

### HasSwitchInfo

`func (o *VlanNetworkAffectingDeviceDetailVO) HasSwitchInfo() bool`

HasSwitchInfo returns a boolean if a field has been set.

### GetType

`func (o *VlanNetworkAffectingDeviceDetailVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VlanNetworkAffectingDeviceDetailVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VlanNetworkAffectingDeviceDetailVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VlanNetworkAffectingDeviceDetailVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


