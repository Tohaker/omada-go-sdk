# SdWanSelectedLanNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConflictLanNetworks** | Pointer to [**[]LanNetworkBrief**](LanNetworkBrief.md) | A list of the conflicted Lan Network. | [optional] 
**LanNetworks** | Pointer to [**[]LanNetworkBrief**](LanNetworkBrief.md) | A list of the SD-WAN selected Lan Network. | [optional] 

## Methods

### NewSdWanSelectedLanNetwork

`func NewSdWanSelectedLanNetwork() *SdWanSelectedLanNetwork`

NewSdWanSelectedLanNetwork instantiates a new SdWanSelectedLanNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanSelectedLanNetworkWithDefaults

`func NewSdWanSelectedLanNetworkWithDefaults() *SdWanSelectedLanNetwork`

NewSdWanSelectedLanNetworkWithDefaults instantiates a new SdWanSelectedLanNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConflictLanNetworks

`func (o *SdWanSelectedLanNetwork) GetConflictLanNetworks() []LanNetworkBrief`

GetConflictLanNetworks returns the ConflictLanNetworks field if non-nil, zero value otherwise.

### GetConflictLanNetworksOk

`func (o *SdWanSelectedLanNetwork) GetConflictLanNetworksOk() (*[]LanNetworkBrief, bool)`

GetConflictLanNetworksOk returns a tuple with the ConflictLanNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictLanNetworks

`func (o *SdWanSelectedLanNetwork) SetConflictLanNetworks(v []LanNetworkBrief)`

SetConflictLanNetworks sets ConflictLanNetworks field to given value.

### HasConflictLanNetworks

`func (o *SdWanSelectedLanNetwork) HasConflictLanNetworks() bool`

HasConflictLanNetworks returns a boolean if a field has been set.

### GetLanNetworks

`func (o *SdWanSelectedLanNetwork) GetLanNetworks() []LanNetworkBrief`

GetLanNetworks returns the LanNetworks field if non-nil, zero value otherwise.

### GetLanNetworksOk

`func (o *SdWanSelectedLanNetwork) GetLanNetworksOk() (*[]LanNetworkBrief, bool)`

GetLanNetworksOk returns a tuple with the LanNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworks

`func (o *SdWanSelectedLanNetwork) SetLanNetworks(v []LanNetworkBrief)`

SetLanNetworks sets LanNetworks field to given value.

### HasLanNetworks

`func (o *SdWanSelectedLanNetwork) HasLanNetworks() bool`

HasLanNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


