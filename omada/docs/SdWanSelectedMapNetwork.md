# SdWanSelectedMapNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The SD-WAN group ID. | [optional] 
**LanNetworks** | Pointer to [**[]SdWanLanNetworkNatReq**](SdWanLanNetworkNatReq.md) | A list of original lan network | [optional] 
**MapNetworkList** | Pointer to **[]string** | A list of map network range | [optional] 
**MemberList** | Pointer to [**[]SdWanMemberBriefInfo**](SdWanMemberBriefInfo.md) | A list of members of the SD-WAN group | [optional] 

## Methods

### NewSdWanSelectedMapNetwork

`func NewSdWanSelectedMapNetwork() *SdWanSelectedMapNetwork`

NewSdWanSelectedMapNetwork instantiates a new SdWanSelectedMapNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanSelectedMapNetworkWithDefaults

`func NewSdWanSelectedMapNetworkWithDefaults() *SdWanSelectedMapNetwork`

NewSdWanSelectedMapNetworkWithDefaults instantiates a new SdWanSelectedMapNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SdWanSelectedMapNetwork) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SdWanSelectedMapNetwork) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SdWanSelectedMapNetwork) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SdWanSelectedMapNetwork) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetLanNetworks

`func (o *SdWanSelectedMapNetwork) GetLanNetworks() []SdWanLanNetworkNatReq`

GetLanNetworks returns the LanNetworks field if non-nil, zero value otherwise.

### GetLanNetworksOk

`func (o *SdWanSelectedMapNetwork) GetLanNetworksOk() (*[]SdWanLanNetworkNatReq, bool)`

GetLanNetworksOk returns a tuple with the LanNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworks

`func (o *SdWanSelectedMapNetwork) SetLanNetworks(v []SdWanLanNetworkNatReq)`

SetLanNetworks sets LanNetworks field to given value.

### HasLanNetworks

`func (o *SdWanSelectedMapNetwork) HasLanNetworks() bool`

HasLanNetworks returns a boolean if a field has been set.

### GetMapNetworkList

`func (o *SdWanSelectedMapNetwork) GetMapNetworkList() []string`

GetMapNetworkList returns the MapNetworkList field if non-nil, zero value otherwise.

### GetMapNetworkListOk

`func (o *SdWanSelectedMapNetwork) GetMapNetworkListOk() (*[]string, bool)`

GetMapNetworkListOk returns a tuple with the MapNetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapNetworkList

`func (o *SdWanSelectedMapNetwork) SetMapNetworkList(v []string)`

SetMapNetworkList sets MapNetworkList field to given value.

### HasMapNetworkList

`func (o *SdWanSelectedMapNetwork) HasMapNetworkList() bool`

HasMapNetworkList returns a boolean if a field has been set.

### GetMemberList

`func (o *SdWanSelectedMapNetwork) GetMemberList() []SdWanMemberBriefInfo`

GetMemberList returns the MemberList field if non-nil, zero value otherwise.

### GetMemberListOk

`func (o *SdWanSelectedMapNetwork) GetMemberListOk() (*[]SdWanMemberBriefInfo, bool)`

GetMemberListOk returns a tuple with the MemberList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberList

`func (o *SdWanSelectedMapNetwork) SetMemberList(v []SdWanMemberBriefInfo)`

SetMemberList sets MemberList field to given value.

### HasMemberList

`func (o *SdWanSelectedMapNetwork) HasMemberList() bool`

HasMemberList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


