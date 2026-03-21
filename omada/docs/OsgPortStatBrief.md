# OsgPortStatBrief

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkStatus** | Pointer to **int32** | The status of link connection. | [optional] 
**Mode** | **int32** | The mode whether the port works as a WAN or LAN. | 
**Name** | Pointer to **string** | The port name of the osg. | [optional] 
**Port** | Pointer to **int32** | The port sequence number of the device. | [optional] 
**PortUuid** | Pointer to **string** | The UUID of port. | [optional] 
**PublicIp** | Pointer to **bool** | Whether the device has a public IP. | [optional] 
**WanIp** | Pointer to **string** | The wan IP of the device. | [optional] 

## Methods

### NewOsgPortStatBrief

`func NewOsgPortStatBrief(mode int32, ) *OsgPortStatBrief`

NewOsgPortStatBrief instantiates a new OsgPortStatBrief object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgPortStatBriefWithDefaults

`func NewOsgPortStatBriefWithDefaults() *OsgPortStatBrief`

NewOsgPortStatBriefWithDefaults instantiates a new OsgPortStatBrief object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkStatus

`func (o *OsgPortStatBrief) GetLinkStatus() int32`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *OsgPortStatBrief) GetLinkStatusOk() (*int32, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *OsgPortStatBrief) SetLinkStatus(v int32)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *OsgPortStatBrief) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetMode

`func (o *OsgPortStatBrief) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OsgPortStatBrief) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OsgPortStatBrief) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetName

`func (o *OsgPortStatBrief) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsgPortStatBrief) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsgPortStatBrief) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsgPortStatBrief) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *OsgPortStatBrief) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OsgPortStatBrief) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OsgPortStatBrief) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OsgPortStatBrief) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortUuid

`func (o *OsgPortStatBrief) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *OsgPortStatBrief) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *OsgPortStatBrief) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.

### HasPortUuid

`func (o *OsgPortStatBrief) HasPortUuid() bool`

HasPortUuid returns a boolean if a field has been set.

### GetPublicIp

`func (o *OsgPortStatBrief) GetPublicIp() bool`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *OsgPortStatBrief) GetPublicIpOk() (*bool, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *OsgPortStatBrief) SetPublicIp(v bool)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *OsgPortStatBrief) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetWanIp

`func (o *OsgPortStatBrief) GetWanIp() string`

GetWanIp returns the WanIp field if non-nil, zero value otherwise.

### GetWanIpOk

`func (o *OsgPortStatBrief) GetWanIpOk() (*string, bool)`

GetWanIpOk returns a tuple with the WanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanIp

`func (o *OsgPortStatBrief) SetWanIp(v string)`

SetWanIp sets WanIp field to given value.

### HasWanIp

`func (o *OsgPortStatBrief) HasWanIp() bool`

HasWanIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


