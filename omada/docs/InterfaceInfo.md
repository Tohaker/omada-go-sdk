# InterfaceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Parameter [channel] should be a value as follows: 0: 2.4GHz  1: 5GHz-1  2:5GHz-2 3: 6GHz | [optional] 
**DownlinkDevices** | Pointer to [**[]DeviceBriefVO**](DeviceBriefVO.md) | Downlink Devices | [optional] 
**InterfaceId** | Pointer to **string** | Interface ID, for example: if interfaceType is network, interfaceId should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**InterfaceName** | Pointer to **string** | Interface name. | [optional] 
**InterfaceType** | Pointer to **int32** | Parameter [interfaceType] should be a value as follows: 0: Wired, 1: Wireless. | [optional] 
**LinkStatus** | Pointer to **int32** | Link status should be a value as follows: 0:LINK_DOWN;1:LINK_UP | [optional] 

## Methods

### NewInterfaceInfo

`func NewInterfaceInfo() *InterfaceInfo`

NewInterfaceInfo instantiates a new InterfaceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceInfoWithDefaults

`func NewInterfaceInfoWithDefaults() *InterfaceInfo`

NewInterfaceInfoWithDefaults instantiates a new InterfaceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *InterfaceInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InterfaceInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InterfaceInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InterfaceInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDownlinkDevices

`func (o *InterfaceInfo) GetDownlinkDevices() []DeviceBriefVO`

GetDownlinkDevices returns the DownlinkDevices field if non-nil, zero value otherwise.

### GetDownlinkDevicesOk

`func (o *InterfaceInfo) GetDownlinkDevicesOk() (*[]DeviceBriefVO, bool)`

GetDownlinkDevicesOk returns a tuple with the DownlinkDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkDevices

`func (o *InterfaceInfo) SetDownlinkDevices(v []DeviceBriefVO)`

SetDownlinkDevices sets DownlinkDevices field to given value.

### HasDownlinkDevices

`func (o *InterfaceInfo) HasDownlinkDevices() bool`

HasDownlinkDevices returns a boolean if a field has been set.

### GetInterfaceId

`func (o *InterfaceInfo) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *InterfaceInfo) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *InterfaceInfo) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *InterfaceInfo) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *InterfaceInfo) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *InterfaceInfo) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *InterfaceInfo) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *InterfaceInfo) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *InterfaceInfo) GetInterfaceType() int32`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *InterfaceInfo) GetInterfaceTypeOk() (*int32, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *InterfaceInfo) SetInterfaceType(v int32)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *InterfaceInfo) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetLinkStatus

`func (o *InterfaceInfo) GetLinkStatus() int32`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *InterfaceInfo) GetLinkStatusOk() (*int32, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *InterfaceInfo) SetLinkStatus(v int32)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *InterfaceInfo) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


