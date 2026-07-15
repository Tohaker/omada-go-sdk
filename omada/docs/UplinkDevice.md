# UplinkDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **int32** | The device UplinkDevice deviceType | [optional] 
**IsStack** | Pointer to **bool** | Whether The device Uplink StackDevice is Stack or not | [optional] 
**Mac** | Pointer to **string** | The device UplinkDevice mac | [optional] 
**Model** | Pointer to **string** | The device UplinkDevice model | [optional] 
**ModelVersion** | Pointer to **string** | The device UplinkDevice modelVersion | [optional] 
**Name** | Pointer to **string** | The device UplinkDevice name | [optional] 
**Port** | Pointer to **string** | The device Uplink Device port number | [optional] 
**ShowModel** | Pointer to **string** | The device UplinkDevice showModel | [optional] 
**StackId** | Pointer to **string** | The device Uplink StackDevice StackId | [optional] 
**StackName** | Pointer to **string** | The device Uplink StackDevice StackName | [optional] 
**StandardPort** | Pointer to **string** | The device Uplink Device standardPort | [optional] 
**Type** | Pointer to **string** | The device UplinkDevice type.Such as: ap, switch, gateway | [optional] 
**Unit** | Pointer to **int32** | The device Uplink StackDevice unit | [optional] 
**UplinkPort** | Pointer to **string** | The device UplinkDevice port | [optional] 
**UplinkType** | Pointer to **int32** | The device link type | [optional] 

## Methods

### NewUplinkDevice

`func NewUplinkDevice() *UplinkDevice`

NewUplinkDevice instantiates a new UplinkDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUplinkDeviceWithDefaults

`func NewUplinkDeviceWithDefaults() *UplinkDevice`

NewUplinkDeviceWithDefaults instantiates a new UplinkDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *UplinkDevice) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *UplinkDevice) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *UplinkDevice) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *UplinkDevice) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetIsStack

`func (o *UplinkDevice) GetIsStack() bool`

GetIsStack returns the IsStack field if non-nil, zero value otherwise.

### GetIsStackOk

`func (o *UplinkDevice) GetIsStackOk() (*bool, bool)`

GetIsStackOk returns a tuple with the IsStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStack

`func (o *UplinkDevice) SetIsStack(v bool)`

SetIsStack sets IsStack field to given value.

### HasIsStack

`func (o *UplinkDevice) HasIsStack() bool`

HasIsStack returns a boolean if a field has been set.

### GetMac

`func (o *UplinkDevice) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *UplinkDevice) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *UplinkDevice) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *UplinkDevice) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *UplinkDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UplinkDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UplinkDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *UplinkDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *UplinkDevice) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *UplinkDevice) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *UplinkDevice) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *UplinkDevice) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *UplinkDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UplinkDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UplinkDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UplinkDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *UplinkDevice) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UplinkDevice) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UplinkDevice) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UplinkDevice) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetShowModel

`func (o *UplinkDevice) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *UplinkDevice) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *UplinkDevice) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *UplinkDevice) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStackId

`func (o *UplinkDevice) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *UplinkDevice) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *UplinkDevice) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *UplinkDevice) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *UplinkDevice) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *UplinkDevice) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *UplinkDevice) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *UplinkDevice) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStandardPort

`func (o *UplinkDevice) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *UplinkDevice) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *UplinkDevice) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *UplinkDevice) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetType

`func (o *UplinkDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UplinkDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UplinkDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UplinkDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *UplinkDevice) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *UplinkDevice) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *UplinkDevice) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *UplinkDevice) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUplinkPort

`func (o *UplinkDevice) GetUplinkPort() string`

GetUplinkPort returns the UplinkPort field if non-nil, zero value otherwise.

### GetUplinkPortOk

`func (o *UplinkDevice) GetUplinkPortOk() (*string, bool)`

GetUplinkPortOk returns a tuple with the UplinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPort

`func (o *UplinkDevice) SetUplinkPort(v string)`

SetUplinkPort sets UplinkPort field to given value.

### HasUplinkPort

`func (o *UplinkDevice) HasUplinkPort() bool`

HasUplinkPort returns a boolean if a field has been set.

### GetUplinkType

`func (o *UplinkDevice) GetUplinkType() int32`

GetUplinkType returns the UplinkType field if non-nil, zero value otherwise.

### GetUplinkTypeOk

`func (o *UplinkDevice) GetUplinkTypeOk() (*int32, bool)`

GetUplinkTypeOk returns a tuple with the UplinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkType

`func (o *UplinkDevice) SetUplinkType(v int32)`

SetUplinkType sets UplinkType field to given value.

### HasUplinkType

`func (o *UplinkDevice) HasUplinkType() bool`

HasUplinkType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


