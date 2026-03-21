# SdWanGroupDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the SD-WAN group | [optional] 
**Id** | Pointer to **string** | The ID of the SD-WAN group | [optional] 
**IpPoolEnd** | Pointer to **string** | The end of the IP pool of the SD-WAN group | [optional] 
**IpPoolStart** | Pointer to **string** | The start of the IP pool of the SD-WAN group | [optional] 
**LinkedSpokes** | Pointer to [**[]SdWanLinkedSpoke**](SdWanLinkedSpoke.md) | A list of linked-spokes of the SD-WAN group | [optional] 
**MemberList** | Pointer to [**[]SdWanMemberInfo**](SdWanMemberInfo.md) | A list of members of the SD-WAN group | [optional] 
**Name** | Pointer to **string** | The name of the SD-WAN group | [optional] 

## Methods

### NewSdWanGroupDetail

`func NewSdWanGroupDetail() *SdWanGroupDetail`

NewSdWanGroupDetail instantiates a new SdWanGroupDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanGroupDetailWithDefaults

`func NewSdWanGroupDetailWithDefaults() *SdWanGroupDetail`

NewSdWanGroupDetailWithDefaults instantiates a new SdWanGroupDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SdWanGroupDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdWanGroupDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdWanGroupDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdWanGroupDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SdWanGroupDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SdWanGroupDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SdWanGroupDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SdWanGroupDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpPoolEnd

`func (o *SdWanGroupDetail) GetIpPoolEnd() string`

GetIpPoolEnd returns the IpPoolEnd field if non-nil, zero value otherwise.

### GetIpPoolEndOk

`func (o *SdWanGroupDetail) GetIpPoolEndOk() (*string, bool)`

GetIpPoolEndOk returns a tuple with the IpPoolEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolEnd

`func (o *SdWanGroupDetail) SetIpPoolEnd(v string)`

SetIpPoolEnd sets IpPoolEnd field to given value.

### HasIpPoolEnd

`func (o *SdWanGroupDetail) HasIpPoolEnd() bool`

HasIpPoolEnd returns a boolean if a field has been set.

### GetIpPoolStart

`func (o *SdWanGroupDetail) GetIpPoolStart() string`

GetIpPoolStart returns the IpPoolStart field if non-nil, zero value otherwise.

### GetIpPoolStartOk

`func (o *SdWanGroupDetail) GetIpPoolStartOk() (*string, bool)`

GetIpPoolStartOk returns a tuple with the IpPoolStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolStart

`func (o *SdWanGroupDetail) SetIpPoolStart(v string)`

SetIpPoolStart sets IpPoolStart field to given value.

### HasIpPoolStart

`func (o *SdWanGroupDetail) HasIpPoolStart() bool`

HasIpPoolStart returns a boolean if a field has been set.

### GetLinkedSpokes

`func (o *SdWanGroupDetail) GetLinkedSpokes() []SdWanLinkedSpoke`

GetLinkedSpokes returns the LinkedSpokes field if non-nil, zero value otherwise.

### GetLinkedSpokesOk

`func (o *SdWanGroupDetail) GetLinkedSpokesOk() (*[]SdWanLinkedSpoke, bool)`

GetLinkedSpokesOk returns a tuple with the LinkedSpokes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedSpokes

`func (o *SdWanGroupDetail) SetLinkedSpokes(v []SdWanLinkedSpoke)`

SetLinkedSpokes sets LinkedSpokes field to given value.

### HasLinkedSpokes

`func (o *SdWanGroupDetail) HasLinkedSpokes() bool`

HasLinkedSpokes returns a boolean if a field has been set.

### GetMemberList

`func (o *SdWanGroupDetail) GetMemberList() []SdWanMemberInfo`

GetMemberList returns the MemberList field if non-nil, zero value otherwise.

### GetMemberListOk

`func (o *SdWanGroupDetail) GetMemberListOk() (*[]SdWanMemberInfo, bool)`

GetMemberListOk returns a tuple with the MemberList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberList

`func (o *SdWanGroupDetail) SetMemberList(v []SdWanMemberInfo)`

SetMemberList sets MemberList field to given value.

### HasMemberList

`func (o *SdWanGroupDetail) HasMemberList() bool`

HasMemberList returns a boolean if a field has been set.

### GetName

`func (o *SdWanGroupDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdWanGroupDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdWanGroupDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdWanGroupDetail) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


