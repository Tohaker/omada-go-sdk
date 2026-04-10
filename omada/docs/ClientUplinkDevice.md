# ClientUplinkDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **int32** | The client UplinkDevice deviceType | [optional] 
**Mac** | Pointer to **string** | The client UplinkDevice mac | [optional] 
**Model** | Pointer to **string** | The client UplinkDevice model | [optional] 
**ModelVersion** | Pointer to **string** | The client UplinkDevice modelVersion | [optional] 
**Name** | Pointer to **string** | The client UplinkDevice name | [optional] 
**ShowModel** | Pointer to **string** | The client UplinkDevice showModel | [optional] 
**Type** | Pointer to **string** | The client UplinkDevice type.Such as: ap, switch, gateway | [optional] 
**UplinkPort** | Pointer to **string** | The client UplinkDevice port | [optional] 
**UplinkType** | Pointer to **int32** | The client link type | [optional] 

## Methods

### NewClientUplinkDevice

`func NewClientUplinkDevice() *ClientUplinkDevice`

NewClientUplinkDevice instantiates a new ClientUplinkDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientUplinkDeviceWithDefaults

`func NewClientUplinkDeviceWithDefaults() *ClientUplinkDevice`

NewClientUplinkDeviceWithDefaults instantiates a new ClientUplinkDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *ClientUplinkDevice) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ClientUplinkDevice) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ClientUplinkDevice) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ClientUplinkDevice) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetMac

`func (o *ClientUplinkDevice) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientUplinkDevice) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientUplinkDevice) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientUplinkDevice) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ClientUplinkDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClientUplinkDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClientUplinkDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ClientUplinkDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ClientUplinkDevice) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ClientUplinkDevice) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ClientUplinkDevice) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ClientUplinkDevice) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *ClientUplinkDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientUplinkDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientUplinkDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientUplinkDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShowModel

`func (o *ClientUplinkDevice) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *ClientUplinkDevice) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *ClientUplinkDevice) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *ClientUplinkDevice) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetType

`func (o *ClientUplinkDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientUplinkDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientUplinkDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientUplinkDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUplinkPort

`func (o *ClientUplinkDevice) GetUplinkPort() string`

GetUplinkPort returns the UplinkPort field if non-nil, zero value otherwise.

### GetUplinkPortOk

`func (o *ClientUplinkDevice) GetUplinkPortOk() (*string, bool)`

GetUplinkPortOk returns a tuple with the UplinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPort

`func (o *ClientUplinkDevice) SetUplinkPort(v string)`

SetUplinkPort sets UplinkPort field to given value.

### HasUplinkPort

`func (o *ClientUplinkDevice) HasUplinkPort() bool`

HasUplinkPort returns a boolean if a field has been set.

### GetUplinkType

`func (o *ClientUplinkDevice) GetUplinkType() int32`

GetUplinkType returns the UplinkType field if non-nil, zero value otherwise.

### GetUplinkTypeOk

`func (o *ClientUplinkDevice) GetUplinkTypeOk() (*int32, bool)`

GetUplinkTypeOk returns a tuple with the UplinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkType

`func (o *ClientUplinkDevice) SetUplinkType(v int32)`

SetUplinkType sets UplinkType field to given value.

### HasUplinkType

`func (o *ClientUplinkDevice) HasUplinkType() bool`

HasUplinkType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


