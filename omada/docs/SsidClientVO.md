# SsidClientVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientMac** | Pointer to **string** | Client Mac. | [optional] 
**ClientModel** | Pointer to **string** | Model of client device. | [optional] 
**ClientType** | Pointer to **string** | Type of client device: iphone, ipod, android, pc, printer, tv... | [optional] 
**DeviceMac** | Pointer to **string** | Device Mac. | [optional] 
**DeviceModel** | Pointer to **string** | Model of the device that client is connected to. | [optional] 
**DeviceModelVersion** | Pointer to **string** | Model version of the device that client is connected to. | [optional] 
**DeviceName** | Pointer to **string** | Device Name. | [optional] 
**DeviceType** | Pointer to **string** | Type of the device that client is connected to. | [optional] 
**DownTraffic** | Pointer to **int64** | Downstream traffic (Byte). | [optional] 
**Ip** | Pointer to **string** | IP Address. | [optional] 
**Name** | Pointer to **string** | Client Name, alias. | [optional] 
**Traffic** | Pointer to **int64** | Total traffic (Byte). | [optional] 
**UpTraffic** | Pointer to **int64** | Upstream traffic (Byte). | [optional] 

## Methods

### NewSsidClientVO

`func NewSsidClientVO() *SsidClientVO`

NewSsidClientVO instantiates a new SsidClientVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidClientVOWithDefaults

`func NewSsidClientVOWithDefaults() *SsidClientVO`

NewSsidClientVOWithDefaults instantiates a new SsidClientVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientMac

`func (o *SsidClientVO) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *SsidClientVO) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *SsidClientVO) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *SsidClientVO) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetClientModel

`func (o *SsidClientVO) GetClientModel() string`

GetClientModel returns the ClientModel field if non-nil, zero value otherwise.

### GetClientModelOk

`func (o *SsidClientVO) GetClientModelOk() (*string, bool)`

GetClientModelOk returns a tuple with the ClientModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientModel

`func (o *SsidClientVO) SetClientModel(v string)`

SetClientModel sets ClientModel field to given value.

### HasClientModel

`func (o *SsidClientVO) HasClientModel() bool`

HasClientModel returns a boolean if a field has been set.

### GetClientType

`func (o *SsidClientVO) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *SsidClientVO) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *SsidClientVO) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *SsidClientVO) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetDeviceMac

`func (o *SsidClientVO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *SsidClientVO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *SsidClientVO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *SsidClientVO) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDeviceModel

`func (o *SsidClientVO) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *SsidClientVO) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *SsidClientVO) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *SsidClientVO) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceModelVersion

`func (o *SsidClientVO) GetDeviceModelVersion() string`

GetDeviceModelVersion returns the DeviceModelVersion field if non-nil, zero value otherwise.

### GetDeviceModelVersionOk

`func (o *SsidClientVO) GetDeviceModelVersionOk() (*string, bool)`

GetDeviceModelVersionOk returns a tuple with the DeviceModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModelVersion

`func (o *SsidClientVO) SetDeviceModelVersion(v string)`

SetDeviceModelVersion sets DeviceModelVersion field to given value.

### HasDeviceModelVersion

`func (o *SsidClientVO) HasDeviceModelVersion() bool`

HasDeviceModelVersion returns a boolean if a field has been set.

### GetDeviceName

`func (o *SsidClientVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SsidClientVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SsidClientVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *SsidClientVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *SsidClientVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *SsidClientVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *SsidClientVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *SsidClientVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDownTraffic

`func (o *SsidClientVO) GetDownTraffic() int64`

GetDownTraffic returns the DownTraffic field if non-nil, zero value otherwise.

### GetDownTrafficOk

`func (o *SsidClientVO) GetDownTrafficOk() (*int64, bool)`

GetDownTrafficOk returns a tuple with the DownTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownTraffic

`func (o *SsidClientVO) SetDownTraffic(v int64)`

SetDownTraffic sets DownTraffic field to given value.

### HasDownTraffic

`func (o *SsidClientVO) HasDownTraffic() bool`

HasDownTraffic returns a boolean if a field has been set.

### GetIp

`func (o *SsidClientVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SsidClientVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SsidClientVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SsidClientVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetName

`func (o *SsidClientVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SsidClientVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SsidClientVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SsidClientVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTraffic

`func (o *SsidClientVO) GetTraffic() int64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *SsidClientVO) GetTrafficOk() (*int64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *SsidClientVO) SetTraffic(v int64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *SsidClientVO) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetUpTraffic

`func (o *SsidClientVO) GetUpTraffic() int64`

GetUpTraffic returns the UpTraffic field if non-nil, zero value otherwise.

### GetUpTrafficOk

`func (o *SsidClientVO) GetUpTrafficOk() (*int64, bool)`

GetUpTrafficOk returns a tuple with the UpTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTraffic

`func (o *SsidClientVO) SetUpTraffic(v int64)`

SetUpTraffic sets UpTraffic field to given value.

### HasUpTraffic

`func (o *SsidClientVO) HasUpTraffic() bool`

HasUpTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


