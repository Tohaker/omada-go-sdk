# SdWanIpPoolRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The SD-WAN group ID. | [optional] 
**IpPoolEnd** | **string** | The end of the IP pool of the sdWan group. | 
**IpPoolStart** | **string** | The start of the IP pool of the sdWan group. | 
**TunnelLimit** | **int32** | Enter the number of tunnels required by this group to determine if the number of IPs in the IP pool is sufficient. | 

## Methods

### NewSdWanIpPoolRange

`func NewSdWanIpPoolRange(ipPoolEnd string, ipPoolStart string, tunnelLimit int32, ) *SdWanIpPoolRange`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


