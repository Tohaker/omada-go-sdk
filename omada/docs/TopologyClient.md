# TopologyClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guest** | Pointer to **bool** | Whether the client is a guest. | [optional] 
**Mac** | Pointer to **string** | Client MAC address, like AA-BB-CC-DD-EE-FF. | [optional] 
**Name** | Pointer to **string** | Client name. | [optional] 
**Type** | Pointer to **string** | Client Type. | [optional] 
**UplinkApInfo** | Pointer to [**TopologyClientUplinkApInfo**](TopologyClientUplinkApInfo.md) |  | [optional] 
**UplinkGatewayInfo** | Pointer to [**TopologyClientUplinkGatewayInfo**](TopologyClientUplinkGatewayInfo.md) |  | [optional] 
**UplinkSwitchInfo** | Pointer to [**TopologyClientUplinkSwitchInfo**](TopologyClientUplinkSwitchInfo.md) |  | [optional] 
**Wireless** | Pointer to **bool** | Whether the client is wireless. | [optional] 

## Methods

### NewTopologyClient

`func NewTopologyClient() *TopologyClient`

NewTopologyClient instantiates a new TopologyClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyClientWithDefaults

`func NewTopologyClientWithDefaults() *TopologyClient`

NewTopologyClientWithDefaults instantiates a new TopologyClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuest

`func (o *TopologyClient) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *TopologyClient) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *TopologyClient) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *TopologyClient) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetMac

`func (o *TopologyClient) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *TopologyClient) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *TopologyClient) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *TopologyClient) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *TopologyClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopologyClient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *TopologyClient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopologyClient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopologyClient) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopologyClient) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUplinkApInfo

`func (o *TopologyClient) GetUplinkApInfo() TopologyClientUplinkApInfo`

GetUplinkApInfo returns the UplinkApInfo field if non-nil, zero value otherwise.

### GetUplinkApInfoOk

`func (o *TopologyClient) GetUplinkApInfoOk() (*TopologyClientUplinkApInfo, bool)`

GetUplinkApInfoOk returns a tuple with the UplinkApInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkApInfo

`func (o *TopologyClient) SetUplinkApInfo(v TopologyClientUplinkApInfo)`

SetUplinkApInfo sets UplinkApInfo field to given value.

### HasUplinkApInfo

`func (o *TopologyClient) HasUplinkApInfo() bool`

HasUplinkApInfo returns a boolean if a field has been set.

### GetUplinkGatewayInfo

`func (o *TopologyClient) GetUplinkGatewayInfo() TopologyClientUplinkGatewayInfo`

GetUplinkGatewayInfo returns the UplinkGatewayInfo field if non-nil, zero value otherwise.

### GetUplinkGatewayInfoOk

`func (o *TopologyClient) GetUplinkGatewayInfoOk() (*TopologyClientUplinkGatewayInfo, bool)`

GetUplinkGatewayInfoOk returns a tuple with the UplinkGatewayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkGatewayInfo

`func (o *TopologyClient) SetUplinkGatewayInfo(v TopologyClientUplinkGatewayInfo)`

SetUplinkGatewayInfo sets UplinkGatewayInfo field to given value.

### HasUplinkGatewayInfo

`func (o *TopologyClient) HasUplinkGatewayInfo() bool`

HasUplinkGatewayInfo returns a boolean if a field has been set.

### GetUplinkSwitchInfo

`func (o *TopologyClient) GetUplinkSwitchInfo() TopologyClientUplinkSwitchInfo`

GetUplinkSwitchInfo returns the UplinkSwitchInfo field if non-nil, zero value otherwise.

### GetUplinkSwitchInfoOk

`func (o *TopologyClient) GetUplinkSwitchInfoOk() (*TopologyClientUplinkSwitchInfo, bool)`

GetUplinkSwitchInfoOk returns a tuple with the UplinkSwitchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkSwitchInfo

`func (o *TopologyClient) SetUplinkSwitchInfo(v TopologyClientUplinkSwitchInfo)`

SetUplinkSwitchInfo sets UplinkSwitchInfo field to given value.

### HasUplinkSwitchInfo

`func (o *TopologyClient) HasUplinkSwitchInfo() bool`

HasUplinkSwitchInfo returns a boolean if a field has been set.

### GetWireless

`func (o *TopologyClient) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *TopologyClient) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *TopologyClient) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *TopologyClient) HasWireless() bool`

HasWireless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


