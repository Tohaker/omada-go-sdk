# TopologyClientNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedNetwork** | Pointer to [**TopologyClientConnectedNetwork**](TopologyClientConnectedNetwork.md) |  | [optional] 
**ConnectedSsid** | Pointer to [**TopologyClientConnectedSsid**](TopologyClientConnectedSsid.md) |  | [optional] 
**DevRxRate** | Pointer to **int64** | Realtime rxRate | [optional] 
**DevTxRate** | Pointer to **int64** | Realtime txRate | [optional] 
**Guest** | Pointer to **bool** | Whether the client is a guest. | [optional] 
**HealthScore** | Pointer to **int32** | Client health score. | [optional] 
**Ip** | Pointer to **string** | Client Ip. | [optional] 
**Mac** | Pointer to **string** | Client MAC address, like AA-BB-CC-DD-EE-FF. | [optional] 
**Model** | Pointer to **string** | Client Model. | [optional] 
**Name** | Pointer to **string** | Client Name. | [optional] 
**Type** | Pointer to **string** | Client Type. | [optional] 
**Vendor** | Pointer to **string** | Client Vendor. | [optional] 
**WiredUpInfo** | Pointer to [**TopologyClientWiredUpInfo**](TopologyClientWiredUpInfo.md) |  | [optional] 
**Wireless** | Pointer to **bool** | Whether the client is wireless. | [optional] 
**WirelessUpInfo** | Pointer to [**TopologyClientWirelessUpInfo**](TopologyClientWirelessUpInfo.md) |  | [optional] 

## Methods

### NewTopologyClientNode

`func NewTopologyClientNode() *TopologyClientNode`

NewTopologyClientNode instantiates a new TopologyClientNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyClientNodeWithDefaults

`func NewTopologyClientNodeWithDefaults() *TopologyClientNode`

NewTopologyClientNodeWithDefaults instantiates a new TopologyClientNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedNetwork

`func (o *TopologyClientNode) GetConnectedNetwork() TopologyClientConnectedNetwork`

GetConnectedNetwork returns the ConnectedNetwork field if non-nil, zero value otherwise.

### GetConnectedNetworkOk

`func (o *TopologyClientNode) GetConnectedNetworkOk() (*TopologyClientConnectedNetwork, bool)`

GetConnectedNetworkOk returns a tuple with the ConnectedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedNetwork

`func (o *TopologyClientNode) SetConnectedNetwork(v TopologyClientConnectedNetwork)`

SetConnectedNetwork sets ConnectedNetwork field to given value.

### HasConnectedNetwork

`func (o *TopologyClientNode) HasConnectedNetwork() bool`

HasConnectedNetwork returns a boolean if a field has been set.

### GetConnectedSsid

`func (o *TopologyClientNode) GetConnectedSsid() TopologyClientConnectedSsid`

GetConnectedSsid returns the ConnectedSsid field if non-nil, zero value otherwise.

### GetConnectedSsidOk

`func (o *TopologyClientNode) GetConnectedSsidOk() (*TopologyClientConnectedSsid, bool)`

GetConnectedSsidOk returns a tuple with the ConnectedSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSsid

`func (o *TopologyClientNode) SetConnectedSsid(v TopologyClientConnectedSsid)`

SetConnectedSsid sets ConnectedSsid field to given value.

### HasConnectedSsid

`func (o *TopologyClientNode) HasConnectedSsid() bool`

HasConnectedSsid returns a boolean if a field has been set.

### GetDevRxRate

`func (o *TopologyClientNode) GetDevRxRate() int64`

GetDevRxRate returns the DevRxRate field if non-nil, zero value otherwise.

### GetDevRxRateOk

`func (o *TopologyClientNode) GetDevRxRateOk() (*int64, bool)`

GetDevRxRateOk returns a tuple with the DevRxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevRxRate

`func (o *TopologyClientNode) SetDevRxRate(v int64)`

SetDevRxRate sets DevRxRate field to given value.

### HasDevRxRate

`func (o *TopologyClientNode) HasDevRxRate() bool`

HasDevRxRate returns a boolean if a field has been set.

### GetDevTxRate

`func (o *TopologyClientNode) GetDevTxRate() int64`

GetDevTxRate returns the DevTxRate field if non-nil, zero value otherwise.

### GetDevTxRateOk

`func (o *TopologyClientNode) GetDevTxRateOk() (*int64, bool)`

GetDevTxRateOk returns a tuple with the DevTxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevTxRate

`func (o *TopologyClientNode) SetDevTxRate(v int64)`

SetDevTxRate sets DevTxRate field to given value.

### HasDevTxRate

`func (o *TopologyClientNode) HasDevTxRate() bool`

HasDevTxRate returns a boolean if a field has been set.

### GetGuest

`func (o *TopologyClientNode) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *TopologyClientNode) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *TopologyClientNode) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *TopologyClientNode) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetHealthScore

`func (o *TopologyClientNode) GetHealthScore() int32`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *TopologyClientNode) GetHealthScoreOk() (*int32, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *TopologyClientNode) SetHealthScore(v int32)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *TopologyClientNode) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetIp

`func (o *TopologyClientNode) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *TopologyClientNode) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *TopologyClientNode) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *TopologyClientNode) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *TopologyClientNode) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *TopologyClientNode) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *TopologyClientNode) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *TopologyClientNode) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *TopologyClientNode) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TopologyClientNode) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TopologyClientNode) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TopologyClientNode) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *TopologyClientNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyClientNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyClientNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopologyClientNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *TopologyClientNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopologyClientNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopologyClientNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopologyClientNode) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *TopologyClientNode) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *TopologyClientNode) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *TopologyClientNode) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *TopologyClientNode) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetWiredUpInfo

`func (o *TopologyClientNode) GetWiredUpInfo() TopologyClientWiredUpInfo`

GetWiredUpInfo returns the WiredUpInfo field if non-nil, zero value otherwise.

### GetWiredUpInfoOk

`func (o *TopologyClientNode) GetWiredUpInfoOk() (*TopologyClientWiredUpInfo, bool)`

GetWiredUpInfoOk returns a tuple with the WiredUpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredUpInfo

`func (o *TopologyClientNode) SetWiredUpInfo(v TopologyClientWiredUpInfo)`

SetWiredUpInfo sets WiredUpInfo field to given value.

### HasWiredUpInfo

`func (o *TopologyClientNode) HasWiredUpInfo() bool`

HasWiredUpInfo returns a boolean if a field has been set.

### GetWireless

`func (o *TopologyClientNode) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *TopologyClientNode) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *TopologyClientNode) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *TopologyClientNode) HasWireless() bool`

HasWireless returns a boolean if a field has been set.

### GetWirelessUpInfo

`func (o *TopologyClientNode) GetWirelessUpInfo() TopologyClientWirelessUpInfo`

GetWirelessUpInfo returns the WirelessUpInfo field if non-nil, zero value otherwise.

### GetWirelessUpInfoOk

`func (o *TopologyClientNode) GetWirelessUpInfoOk() (*TopologyClientWirelessUpInfo, bool)`

GetWirelessUpInfoOk returns a tuple with the WirelessUpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessUpInfo

`func (o *TopologyClientNode) SetWirelessUpInfo(v TopologyClientWirelessUpInfo)`

SetWirelessUpInfo sets WirelessUpInfo field to given value.

### HasWirelessUpInfo

`func (o *TopologyClientNode) HasWirelessUpInfo() bool`

HasWirelessUpInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


