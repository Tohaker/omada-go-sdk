# SdWanGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the SD-WAN group | [optional] 
**EnableNat** | **bool** | Whether the group enable SD-WAN virtual network Map | 
**IpPoolEnd** | Pointer to **string** | The end of the IP pool of the SD-WAN group, it is recommended to ignore it as it will be generated automatically | [optional] 
**IpPoolStart** | Pointer to **string** | The start of the IP pool of the SD-WAN group, it is recommended to ignore it as it will be generated automatically | [optional] 
**LinkedSpokes** | Pointer to [**[]SdWanLinkedSpokeConfig**](SdWanLinkedSpokeConfig.md) | A list of linked-spokes of the SD-WAN group | [optional] 
**MemberList** | [**[]SdWanMemberConfig**](SdWanMemberConfig.md) | A list of members of the SD-WAN group | 
**Name** | **string** | The name of the SD-WAN group | 
**NatInfo** | Pointer to [**SdWanNatInfoConfig**](SdWanNatInfoConfig.md) |  | [optional] 

## Methods

### NewSdWanGroup

`func NewSdWanGroup(enableNat bool, memberList []SdWanMemberConfig, name string, ) *SdWanGroup`

NewSdWanGroup instantiates a new SdWanGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanGroupWithDefaults

`func NewSdWanGroupWithDefaults() *SdWanGroup`

NewSdWanGroupWithDefaults instantiates a new SdWanGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SdWanGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdWanGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdWanGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdWanGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableNat

`func (o *SdWanGroup) GetEnableNat() bool`

GetEnableNat returns the EnableNat field if non-nil, zero value otherwise.

### GetEnableNatOk

`func (o *SdWanGroup) GetEnableNatOk() (*bool, bool)`

GetEnableNatOk returns a tuple with the EnableNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNat

`func (o *SdWanGroup) SetEnableNat(v bool)`

SetEnableNat sets EnableNat field to given value.


### GetIpPoolEnd

`func (o *SdWanGroup) GetIpPoolEnd() string`

GetIpPoolEnd returns the IpPoolEnd field if non-nil, zero value otherwise.

### GetIpPoolEndOk

`func (o *SdWanGroup) GetIpPoolEndOk() (*string, bool)`

GetIpPoolEndOk returns a tuple with the IpPoolEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolEnd

`func (o *SdWanGroup) SetIpPoolEnd(v string)`

SetIpPoolEnd sets IpPoolEnd field to given value.

### HasIpPoolEnd

`func (o *SdWanGroup) HasIpPoolEnd() bool`

HasIpPoolEnd returns a boolean if a field has been set.

### GetIpPoolStart

`func (o *SdWanGroup) GetIpPoolStart() string`

GetIpPoolStart returns the IpPoolStart field if non-nil, zero value otherwise.

### GetIpPoolStartOk

`func (o *SdWanGroup) GetIpPoolStartOk() (*string, bool)`

GetIpPoolStartOk returns a tuple with the IpPoolStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolStart

`func (o *SdWanGroup) SetIpPoolStart(v string)`

SetIpPoolStart sets IpPoolStart field to given value.

### HasIpPoolStart

`func (o *SdWanGroup) HasIpPoolStart() bool`

HasIpPoolStart returns a boolean if a field has been set.

### GetLinkedSpokes

`func (o *SdWanGroup) GetLinkedSpokes() []SdWanLinkedSpokeConfig`

GetLinkedSpokes returns the LinkedSpokes field if non-nil, zero value otherwise.

### GetLinkedSpokesOk

`func (o *SdWanGroup) GetLinkedSpokesOk() (*[]SdWanLinkedSpokeConfig, bool)`

GetLinkedSpokesOk returns a tuple with the LinkedSpokes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedSpokes

`func (o *SdWanGroup) SetLinkedSpokes(v []SdWanLinkedSpokeConfig)`

SetLinkedSpokes sets LinkedSpokes field to given value.

### HasLinkedSpokes

`func (o *SdWanGroup) HasLinkedSpokes() bool`

HasLinkedSpokes returns a boolean if a field has been set.

### GetMemberList

`func (o *SdWanGroup) GetMemberList() []SdWanMemberConfig`

GetMemberList returns the MemberList field if non-nil, zero value otherwise.

### GetMemberListOk

`func (o *SdWanGroup) GetMemberListOk() (*[]SdWanMemberConfig, bool)`

GetMemberListOk returns a tuple with the MemberList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberList

`func (o *SdWanGroup) SetMemberList(v []SdWanMemberConfig)`

SetMemberList sets MemberList field to given value.


### GetName

`func (o *SdWanGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdWanGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdWanGroup) SetName(v string)`

SetName sets Name field to given value.


### GetNatInfo

`func (o *SdWanGroup) GetNatInfo() SdWanNatInfoConfig`

GetNatInfo returns the NatInfo field if non-nil, zero value otherwise.

### GetNatInfoOk

`func (o *SdWanGroup) GetNatInfoOk() (*SdWanNatInfoConfig, bool)`

GetNatInfoOk returns a tuple with the NatInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInfo

`func (o *SdWanGroup) SetNatInfo(v SdWanNatInfoConfig)`

SetNatInfo sets NatInfo field to given value.

### HasNatInfo

`func (o *SdWanGroup) HasNatInfo() bool`

HasNatInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


