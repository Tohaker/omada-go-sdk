# TopologyClientWirelessLinkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Channel. | [optional] 
**RadioId** | Pointer to **int32** | RadioId, it should be a value as follows: 0:2G, 1:5G, 2:5G2, 3:6G. | [optional] 

## Methods

### NewTopologyClientWirelessLinkInfo

`func NewTopologyClientWirelessLinkInfo() *TopologyClientWirelessLinkInfo`

NewTopologyClientWirelessLinkInfo instantiates a new TopologyClientWirelessLinkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyClientWirelessLinkInfoWithDefaults

`func NewTopologyClientWirelessLinkInfoWithDefaults() *TopologyClientWirelessLinkInfo`

NewTopologyClientWirelessLinkInfoWithDefaults instantiates a new TopologyClientWirelessLinkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *TopologyClientWirelessLinkInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *TopologyClientWirelessLinkInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *TopologyClientWirelessLinkInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *TopologyClientWirelessLinkInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetRadioId

`func (o *TopologyClientWirelessLinkInfo) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *TopologyClientWirelessLinkInfo) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *TopologyClientWirelessLinkInfo) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *TopologyClientWirelessLinkInfo) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


