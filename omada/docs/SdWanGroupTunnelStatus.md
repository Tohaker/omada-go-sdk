# SdWanGroupTunnelStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstCheck** | Pointer to **bool** | Whether the first check has been completed. | [optional] 
**LinkedSpokes** | Pointer to [**[]SdWanLinksToHub**](SdWanLinksToHub.md) | A list of hub-spokes of the sdWan group | [optional] 

## Methods

### NewSdWanGroupTunnelStatus

`func NewSdWanGroupTunnelStatus() *SdWanGroupTunnelStatus`

NewSdWanGroupTunnelStatus instantiates a new SdWanGroupTunnelStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanGroupTunnelStatusWithDefaults

`func NewSdWanGroupTunnelStatusWithDefaults() *SdWanGroupTunnelStatus`

NewSdWanGroupTunnelStatusWithDefaults instantiates a new SdWanGroupTunnelStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstCheck

`func (o *SdWanGroupTunnelStatus) GetFirstCheck() bool`

GetFirstCheck returns the FirstCheck field if non-nil, zero value otherwise.

### GetFirstCheckOk

`func (o *SdWanGroupTunnelStatus) GetFirstCheckOk() (*bool, bool)`

GetFirstCheckOk returns a tuple with the FirstCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCheck

`func (o *SdWanGroupTunnelStatus) SetFirstCheck(v bool)`

SetFirstCheck sets FirstCheck field to given value.

### HasFirstCheck

`func (o *SdWanGroupTunnelStatus) HasFirstCheck() bool`

HasFirstCheck returns a boolean if a field has been set.

### GetLinkedSpokes

`func (o *SdWanGroupTunnelStatus) GetLinkedSpokes() []SdWanLinksToHub`

GetLinkedSpokes returns the LinkedSpokes field if non-nil, zero value otherwise.

### GetLinkedSpokesOk

`func (o *SdWanGroupTunnelStatus) GetLinkedSpokesOk() (*[]SdWanLinksToHub, bool)`

GetLinkedSpokesOk returns a tuple with the LinkedSpokes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedSpokes

`func (o *SdWanGroupTunnelStatus) SetLinkedSpokes(v []SdWanLinksToHub)`

SetLinkedSpokes sets LinkedSpokes field to given value.

### HasLinkedSpokes

`func (o *SdWanGroupTunnelStatus) HasLinkedSpokes() bool`

HasLinkedSpokes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


