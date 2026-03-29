# ActiveDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Device MAC address, linked to the device details page | [optional] 
**Model** | Pointer to **string** | Device model, such as EAP620 HD | [optional] 
**ModelVersion** | Pointer to **string** | Model version, such as 3.0 | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**Status** | Pointer to **int32** | Status should be a value as follows: 0: Connected; 1: Disconnected. | [optional] 
**Traffic** | Pointer to **float64** | Traffic measured in GB | [optional] 

## Methods

### NewActiveDevice

`func NewActiveDevice() *ActiveDevice`

NewActiveDevice instantiates a new ActiveDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDeviceWithDefaults

`func NewActiveDeviceWithDefaults() *ActiveDevice`

NewActiveDeviceWithDefaults instantiates a new ActiveDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *ActiveDevice) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ActiveDevice) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ActiveDevice) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ActiveDevice) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ActiveDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ActiveDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ActiveDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ActiveDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ActiveDevice) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ActiveDevice) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ActiveDevice) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ActiveDevice) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *ActiveDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActiveDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActiveDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActiveDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ActiveDevice) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActiveDevice) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActiveDevice) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActiveDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTraffic

`func (o *ActiveDevice) GetTraffic() float64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *ActiveDevice) GetTrafficOk() (*float64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *ActiveDevice) SetTraffic(v float64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *ActiveDevice) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


