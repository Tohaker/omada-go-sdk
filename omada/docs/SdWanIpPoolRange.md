# SdWanIpPoolRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The SD-WAN group ID. | [optional] 
**IpPoolEnd** | Pointer to **string** | The end of the IP pool of the sdWan group. | [optional] 
**IpPoolStart** | Pointer to **string** | The start of the IP pool of the sdWan group. | [optional] 
**TunnelLimit** | Pointer to **int32** | Maximum number of tunnels that each SD-WAN device. | [optional] 

## Methods

### NewSdWanIpPoolRange

`func NewSdWanIpPoolRange() *SdWanIpPoolRange`

NewSdWanIpPoolRange instantiates a new SdWanIpPoolRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanIpPoolRangeWithDefaults

`func NewSdWanIpPoolRangeWithDefaults() *SdWanIpPoolRange`

NewSdWanIpPoolRangeWithDefaults instantiates a new SdWanIpPoolRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SdWanIpPoolRange) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SdWanIpPoolRange) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SdWanIpPoolRange) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SdWanIpPoolRange) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetIpPoolEnd

`func (o *SdWanIpPoolRange) GetIpPoolEnd() string`

GetIpPoolEnd returns the IpPoolEnd field if non-nil, zero value otherwise.

### GetIpPoolEndOk

`func (o *SdWanIpPoolRange) GetIpPoolEndOk() (*string, bool)`

GetIpPoolEndOk returns a tuple with the IpPoolEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolEnd

`func (o *SdWanIpPoolRange) SetIpPoolEnd(v string)`

SetIpPoolEnd sets IpPoolEnd field to given value.

### HasIpPoolEnd

`func (o *SdWanIpPoolRange) HasIpPoolEnd() bool`

HasIpPoolEnd returns a boolean if a field has been set.

### GetIpPoolStart

`func (o *SdWanIpPoolRange) GetIpPoolStart() string`

GetIpPoolStart returns the IpPoolStart field if non-nil, zero value otherwise.

### GetIpPoolStartOk

`func (o *SdWanIpPoolRange) GetIpPoolStartOk() (*string, bool)`

GetIpPoolStartOk returns a tuple with the IpPoolStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolStart

`func (o *SdWanIpPoolRange) SetIpPoolStart(v string)`

SetIpPoolStart sets IpPoolStart field to given value.

### HasIpPoolStart

`func (o *SdWanIpPoolRange) HasIpPoolStart() bool`

HasIpPoolStart returns a boolean if a field has been set.

### GetTunnelLimit

`func (o *SdWanIpPoolRange) GetTunnelLimit() int32`

GetTunnelLimit returns the TunnelLimit field if non-nil, zero value otherwise.

### GetTunnelLimitOk

`func (o *SdWanIpPoolRange) GetTunnelLimitOk() (*int32, bool)`

GetTunnelLimitOk returns a tuple with the TunnelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelLimit

`func (o *SdWanIpPoolRange) SetTunnelLimit(v int32)`

SetTunnelLimit sets TunnelLimit field to given value.

### HasTunnelLimit

`func (o *SdWanIpPoolRange) HasTunnelLimit() bool`

HasTunnelLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


