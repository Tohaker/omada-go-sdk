# StaticRoutingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | **[]string** | IP address/SubNet, up to 16 entries are allowed for the destinations list. | 
**ExistVirtualWan** | Pointer to **bool** | Whether the interface option exists virtual WAN | [optional] 
**ExistVpnClient** | Pointer to **bool** | Whether the interface option exists VPN client | [optional] 
**Id** | Pointer to **string** | ID | [optional] 
**InterfaceId** | Pointer to **string** | Interface ID, for example: if interfaceType is network, interfaceId should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**InterfaceType** | Pointer to **int32** | Only for routeType:1, interfaceType should be a value as follows: 0: Internet(WAN); 1: Network(LAN); 2: L2TP, 3: PPTP, 5: Virtual WAN. | [optional] 
**Metric** | **int32** | Metric should be within the range of 0–15. | 
**Name** | **string** | Name, name should contain 1 to 64 characters. | 
**NextHopIp** | Pointer to **string** | Only for routeType:0 or routeType:1 and selected WAN is Static IP/Dynamic IP mode | [optional] 
**RouteType** | **int32** | RouteType should be a value as follows: 0: NextHop; 1: Interface | 
**Status** | **bool** | Status | 

## Methods

### NewStaticRoutingInfo

`func NewStaticRoutingInfo(destinations []string, metric int32, name string, routeType int32, status bool, ) *StaticRoutingInfo`

NewStaticRoutingInfo instantiates a new StaticRoutingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticRoutingInfoWithDefaults

`func NewStaticRoutingInfoWithDefaults() *StaticRoutingInfo`

NewStaticRoutingInfoWithDefaults instantiates a new StaticRoutingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *StaticRoutingInfo) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *StaticRoutingInfo) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *StaticRoutingInfo) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.


### GetExistVirtualWan

`func (o *StaticRoutingInfo) GetExistVirtualWan() bool`

GetExistVirtualWan returns the ExistVirtualWan field if non-nil, zero value otherwise.

### GetExistVirtualWanOk

`func (o *StaticRoutingInfo) GetExistVirtualWanOk() (*bool, bool)`

GetExistVirtualWanOk returns a tuple with the ExistVirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistVirtualWan

`func (o *StaticRoutingInfo) SetExistVirtualWan(v bool)`

SetExistVirtualWan sets ExistVirtualWan field to given value.

### HasExistVirtualWan

`func (o *StaticRoutingInfo) HasExistVirtualWan() bool`

HasExistVirtualWan returns a boolean if a field has been set.

### GetExistVpnClient

`func (o *StaticRoutingInfo) GetExistVpnClient() bool`

GetExistVpnClient returns the ExistVpnClient field if non-nil, zero value otherwise.

### GetExistVpnClientOk

`func (o *StaticRoutingInfo) GetExistVpnClientOk() (*bool, bool)`

GetExistVpnClientOk returns a tuple with the ExistVpnClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistVpnClient

`func (o *StaticRoutingInfo) SetExistVpnClient(v bool)`

SetExistVpnClient sets ExistVpnClient field to given value.

### HasExistVpnClient

`func (o *StaticRoutingInfo) HasExistVpnClient() bool`

HasExistVpnClient returns a boolean if a field has been set.

### GetId

`func (o *StaticRoutingInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StaticRoutingInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StaticRoutingInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StaticRoutingInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *StaticRoutingInfo) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *StaticRoutingInfo) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *StaticRoutingInfo) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *StaticRoutingInfo) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *StaticRoutingInfo) GetInterfaceType() int32`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *StaticRoutingInfo) GetInterfaceTypeOk() (*int32, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *StaticRoutingInfo) SetInterfaceType(v int32)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *StaticRoutingInfo) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetMetric

`func (o *StaticRoutingInfo) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *StaticRoutingInfo) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *StaticRoutingInfo) SetMetric(v int32)`

SetMetric sets Metric field to given value.


### GetName

`func (o *StaticRoutingInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StaticRoutingInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StaticRoutingInfo) SetName(v string)`

SetName sets Name field to given value.


### GetNextHopIp

`func (o *StaticRoutingInfo) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *StaticRoutingInfo) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *StaticRoutingInfo) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.

### HasNextHopIp

`func (o *StaticRoutingInfo) HasNextHopIp() bool`

HasNextHopIp returns a boolean if a field has been set.

### GetRouteType

`func (o *StaticRoutingInfo) GetRouteType() int32`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *StaticRoutingInfo) GetRouteTypeOk() (*int32, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *StaticRoutingInfo) SetRouteType(v int32)`

SetRouteType sets RouteType field to given value.


### GetStatus

`func (o *StaticRoutingInfo) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StaticRoutingInfo) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StaticRoutingInfo) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


