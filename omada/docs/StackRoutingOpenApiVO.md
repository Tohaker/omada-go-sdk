# StackRoutingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestIpVrfName** | Pointer to **string** | Destination ip vrf name | [optional] 
**DestinationIp** | Pointer to **string** | Destination IP. | [optional] 
**Distance** | Pointer to **int32** | Distance. | [optional] 
**InterfaceName** | Pointer to **[]string** | Interface Name | [optional] 
**Metric** | Pointer to **int32** | Metric | [optional] 
**Name** | Pointer to **string** | Stack Name. | [optional] 
**NextHop** | Pointer to **string** | Next Ip. | [optional] 
**NextHopVrfName** | Pointer to **string** | Next ip vrf name | [optional] 
**NextHops** | Pointer to **[]string** | Next IP list. | [optional] 
**NextHopsVrf** | Pointer to **[]string** | Next ip vrf | [optional] 
**StackId** | Pointer to **string** | Stack ID. | [optional] 
**Type** | Pointer to **string** | Type | [optional] 

## Methods

### NewStackRoutingOpenApiVO

`func NewStackRoutingOpenApiVO() *StackRoutingOpenApiVO`

NewStackRoutingOpenApiVO instantiates a new StackRoutingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackRoutingOpenApiVOWithDefaults

`func NewStackRoutingOpenApiVOWithDefaults() *StackRoutingOpenApiVO`

NewStackRoutingOpenApiVOWithDefaults instantiates a new StackRoutingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestIpVrfName

`func (o *StackRoutingOpenApiVO) GetDestIpVrfName() string`

GetDestIpVrfName returns the DestIpVrfName field if non-nil, zero value otherwise.

### GetDestIpVrfNameOk

`func (o *StackRoutingOpenApiVO) GetDestIpVrfNameOk() (*string, bool)`

GetDestIpVrfNameOk returns a tuple with the DestIpVrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestIpVrfName

`func (o *StackRoutingOpenApiVO) SetDestIpVrfName(v string)`

SetDestIpVrfName sets DestIpVrfName field to given value.

### HasDestIpVrfName

`func (o *StackRoutingOpenApiVO) HasDestIpVrfName() bool`

HasDestIpVrfName returns a boolean if a field has been set.

### GetDestinationIp

`func (o *StackRoutingOpenApiVO) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *StackRoutingOpenApiVO) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *StackRoutingOpenApiVO) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.

### HasDestinationIp

`func (o *StackRoutingOpenApiVO) HasDestinationIp() bool`

HasDestinationIp returns a boolean if a field has been set.

### GetDistance

`func (o *StackRoutingOpenApiVO) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *StackRoutingOpenApiVO) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *StackRoutingOpenApiVO) SetDistance(v int32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *StackRoutingOpenApiVO) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetInterfaceName

`func (o *StackRoutingOpenApiVO) GetInterfaceName() []string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *StackRoutingOpenApiVO) GetInterfaceNameOk() (*[]string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *StackRoutingOpenApiVO) SetInterfaceName(v []string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *StackRoutingOpenApiVO) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetMetric

`func (o *StackRoutingOpenApiVO) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *StackRoutingOpenApiVO) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *StackRoutingOpenApiVO) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *StackRoutingOpenApiVO) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetName

`func (o *StackRoutingOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackRoutingOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackRoutingOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StackRoutingOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextHop

`func (o *StackRoutingOpenApiVO) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *StackRoutingOpenApiVO) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *StackRoutingOpenApiVO) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *StackRoutingOpenApiVO) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetNextHopVrfName

`func (o *StackRoutingOpenApiVO) GetNextHopVrfName() string`

GetNextHopVrfName returns the NextHopVrfName field if non-nil, zero value otherwise.

### GetNextHopVrfNameOk

`func (o *StackRoutingOpenApiVO) GetNextHopVrfNameOk() (*string, bool)`

GetNextHopVrfNameOk returns a tuple with the NextHopVrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopVrfName

`func (o *StackRoutingOpenApiVO) SetNextHopVrfName(v string)`

SetNextHopVrfName sets NextHopVrfName field to given value.

### HasNextHopVrfName

`func (o *StackRoutingOpenApiVO) HasNextHopVrfName() bool`

HasNextHopVrfName returns a boolean if a field has been set.

### GetNextHops

`func (o *StackRoutingOpenApiVO) GetNextHops() []string`

GetNextHops returns the NextHops field if non-nil, zero value otherwise.

### GetNextHopsOk

`func (o *StackRoutingOpenApiVO) GetNextHopsOk() (*[]string, bool)`

GetNextHopsOk returns a tuple with the NextHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHops

`func (o *StackRoutingOpenApiVO) SetNextHops(v []string)`

SetNextHops sets NextHops field to given value.

### HasNextHops

`func (o *StackRoutingOpenApiVO) HasNextHops() bool`

HasNextHops returns a boolean if a field has been set.

### GetNextHopsVrf

`func (o *StackRoutingOpenApiVO) GetNextHopsVrf() []string`

GetNextHopsVrf returns the NextHopsVrf field if non-nil, zero value otherwise.

### GetNextHopsVrfOk

`func (o *StackRoutingOpenApiVO) GetNextHopsVrfOk() (*[]string, bool)`

GetNextHopsVrfOk returns a tuple with the NextHopsVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopsVrf

`func (o *StackRoutingOpenApiVO) SetNextHopsVrf(v []string)`

SetNextHopsVrf sets NextHopsVrf field to given value.

### HasNextHopsVrf

`func (o *StackRoutingOpenApiVO) HasNextHopsVrf() bool`

HasNextHopsVrf returns a boolean if a field has been set.

### GetStackId

`func (o *StackRoutingOpenApiVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *StackRoutingOpenApiVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *StackRoutingOpenApiVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *StackRoutingOpenApiVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetType

`func (o *StackRoutingOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StackRoutingOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StackRoutingOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StackRoutingOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


