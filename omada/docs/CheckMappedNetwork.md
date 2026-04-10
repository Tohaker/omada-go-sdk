# CheckMappedNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MappedNetworks** | Pointer to **[]string** | A list of current mapped network of the SD-WAN group | [optional] 
**MemberList** | Pointer to [**[]SdWanMemberBriefInfo**](SdWanMemberBriefInfo.md) | A list of members of the SD-WAN group | [optional] 
**ModifiedNetwork** | Pointer to **string** | The IP Subnet of the modified network | [optional] 

## Methods

### NewCheckMappedNetwork

`func NewCheckMappedNetwork() *CheckMappedNetwork`

NewCheckMappedNetwork instantiates a new CheckMappedNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckMappedNetworkWithDefaults

`func NewCheckMappedNetworkWithDefaults() *CheckMappedNetwork`

NewCheckMappedNetworkWithDefaults instantiates a new CheckMappedNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappedNetworks

`func (o *CheckMappedNetwork) GetMappedNetworks() []string`

GetMappedNetworks returns the MappedNetworks field if non-nil, zero value otherwise.

### GetMappedNetworksOk

`func (o *CheckMappedNetwork) GetMappedNetworksOk() (*[]string, bool)`

GetMappedNetworksOk returns a tuple with the MappedNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedNetworks

`func (o *CheckMappedNetwork) SetMappedNetworks(v []string)`

SetMappedNetworks sets MappedNetworks field to given value.

### HasMappedNetworks

`func (o *CheckMappedNetwork) HasMappedNetworks() bool`

HasMappedNetworks returns a boolean if a field has been set.

### GetMemberList

`func (o *CheckMappedNetwork) GetMemberList() []SdWanMemberBriefInfo`

GetMemberList returns the MemberList field if non-nil, zero value otherwise.

### GetMemberListOk

`func (o *CheckMappedNetwork) GetMemberListOk() (*[]SdWanMemberBriefInfo, bool)`

GetMemberListOk returns a tuple with the MemberList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberList

`func (o *CheckMappedNetwork) SetMemberList(v []SdWanMemberBriefInfo)`

SetMemberList sets MemberList field to given value.

### HasMemberList

`func (o *CheckMappedNetwork) HasMemberList() bool`

HasMemberList returns a boolean if a field has been set.

### GetModifiedNetwork

`func (o *CheckMappedNetwork) GetModifiedNetwork() string`

GetModifiedNetwork returns the ModifiedNetwork field if non-nil, zero value otherwise.

### GetModifiedNetworkOk

`func (o *CheckMappedNetwork) GetModifiedNetworkOk() (*string, bool)`

GetModifiedNetworkOk returns a tuple with the ModifiedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedNetwork

`func (o *CheckMappedNetwork) SetModifiedNetwork(v string)`

SetModifiedNetwork sets ModifiedNetwork field to given value.

### HasModifiedNetwork

`func (o *CheckMappedNetwork) HasModifiedNetwork() bool`

HasModifiedNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


