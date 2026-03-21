# MspDevicesQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**[]MspDeviceItem**](MspDeviceItem.md) | Devices to query | 

## Methods

### NewMspDevicesQueryOpenApiVO

`func NewMspDevicesQueryOpenApiVO(devices []MspDeviceItem, ) *MspDevicesQueryOpenApiVO`

NewMspDevicesQueryOpenApiVO instantiates a new MspDevicesQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspDevicesQueryOpenApiVOWithDefaults

`func NewMspDevicesQueryOpenApiVOWithDefaults() *MspDevicesQueryOpenApiVO`

NewMspDevicesQueryOpenApiVOWithDefaults instantiates a new MspDevicesQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *MspDevicesQueryOpenApiVO) GetDevices() []MspDeviceItem`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *MspDevicesQueryOpenApiVO) GetDevicesOk() (*[]MspDeviceItem, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *MspDevicesQueryOpenApiVO) SetDevices(v []MspDeviceItem)`

SetDevices sets Devices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


