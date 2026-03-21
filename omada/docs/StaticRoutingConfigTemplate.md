# StaticRoutingConfigTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | **[]string** | IP address/SubNet | 
**InterfaceId** | Pointer to **string** | Interface ID, for example: if interfaceType is network, interfaceId should be LAN network ID. LAN Network can be created using &#39;Create LAN network template&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list template&#39; interface. | [optional] 
**InterfaceType** | Pointer to **int32** | Only for routeType:1, interfaceType should be a value as follows: 0: Internet(WAN); 1: Network(LAN); 2: L2TP, 3: PPTP, 5: Virtual WAN. | [optional] 
**Metric** | **int32** | Metric should be within the range of 0–15. | 
**Name** | **string** | Name, name should contain 1 to 64 characters. | 
**NextHopIp** | Pointer to **string** | Only for routeType:0 or routeType:1 and selected WAN is Static IP/Dynamic IP mode | [optional] 
**RouteType** | **int32** | RouteType should be a value as follows: 0: NextHop; 1: Interface | 
**Status** | **bool** | Status | 

## Methods

### NewStaticRoutingConfigTemplate

`func NewStaticRoutingConfigTemplate(destinations []string, metric int32, name string, routeType int32, status bool, ) *StaticRoutingConfigTemplate`

NewStaticRoutingConfigTemplate instantiates a new StaticRoutingConfigTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticRoutingConfigTemplateWithDefaults

`func NewStaticRoutingConfigTemplateWithDefaults() *StaticRoutingConfigTemplate`

NewStaticRoutingConfigTemplateWithDefaults instantiates a new StaticRoutingConfigTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *StaticRoutingConfigTemplate) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *StaticRoutingConfigTemplate) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *StaticRoutingConfigTemplate) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.


### GetInterfaceId

`func (o *StaticRoutingConfigTemplate) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *StaticRoutingConfigTemplate) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *StaticRoutingConfigTemplate) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *StaticRoutingConfigTemplate) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *StaticRoutingConfigTemplate) GetInterfaceType() int32`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *StaticRoutingConfigTemplate) GetInterfaceTypeOk() (*int32, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *StaticRoutingConfigTemplate) SetInterfaceType(v int32)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *StaticRoutingConfigTemplate) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetMetric

`func (o *StaticRoutingConfigTemplate) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *StaticRoutingConfigTemplate) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *StaticRoutingConfigTemplate) SetMetric(v int32)`

SetMetric sets Metric field to given value.


### GetName

`func (o *StaticRoutingConfigTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StaticRoutingConfigTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StaticRoutingConfigTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetNextHopIp

`func (o *StaticRoutingConfigTemplate) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *StaticRoutingConfigTemplate) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *StaticRoutingConfigTemplate) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.

### HasNextHopIp

`func (o *StaticRoutingConfigTemplate) HasNextHopIp() bool`

HasNextHopIp returns a boolean if a field has been set.

### GetRouteType

`func (o *StaticRoutingConfigTemplate) GetRouteType() int32`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *StaticRoutingConfigTemplate) GetRouteTypeOk() (*int32, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *StaticRoutingConfigTemplate) SetRouteType(v int32)`

SetRouteType sets RouteType field to given value.


### GetStatus

`func (o *StaticRoutingConfigTemplate) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StaticRoutingConfigTemplate) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StaticRoutingConfigTemplate) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


