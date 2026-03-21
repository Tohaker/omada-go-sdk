# OswRoutingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestIpVrf** | Pointer to **string** | Destination Vrf. | [optional] 
**DestinationIp** | Pointer to **string** | Destination IP. | [optional] 
**Distance** | Pointer to **int32** | Distance. | [optional] 
**InterfaceName** | Pointer to **[]string** | Interface Name | [optional] 
**Mac** | Pointer to **string** | Switch Mac. | [optional] 
**Metric** | Pointer to **int32** | Metric | [optional] 
**Name** | Pointer to **string** | Switch Name. | [optional] 
**NextHop** | Pointer to **string** | Next Ip. | [optional] 
**NextHopVrf** | Pointer to **string** | Next Vrf. | [optional] 
**NextHops** | Pointer to **[]string** | Next IP list. | [optional] 
**NextHopsVrf** | Pointer to **[]string** | Next vrfs | [optional] 
**OswDestinationIps** | Pointer to **[]string** | Destination IP list. | [optional] 
**StackMsg** | Pointer to [**StackMsgOpenApiVO**](StackMsgOpenApiVO.md) |  | [optional] 
**Type** | Pointer to **string** | Type | [optional] 

## Methods

### NewOswRoutingOpenApiVO

`func NewOswRoutingOpenApiVO() *OswRoutingOpenApiVO`

NewOswRoutingOpenApiVO instantiates a new OswRoutingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswRoutingOpenApiVOWithDefaults

`func NewOswRoutingOpenApiVOWithDefaults() *OswRoutingOpenApiVO`

NewOswRoutingOpenApiVOWithDefaults instantiates a new OswRoutingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestIpVrf

`func (o *OswRoutingOpenApiVO) GetDestIpVrf() string`

GetDestIpVrf returns the DestIpVrf field if non-nil, zero value otherwise.

### GetDestIpVrfOk

`func (o *OswRoutingOpenApiVO) GetDestIpVrfOk() (*string, bool)`

GetDestIpVrfOk returns a tuple with the DestIpVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestIpVrf

`func (o *OswRoutingOpenApiVO) SetDestIpVrf(v string)`

SetDestIpVrf sets DestIpVrf field to given value.

### HasDestIpVrf

`func (o *OswRoutingOpenApiVO) HasDestIpVrf() bool`

HasDestIpVrf returns a boolean if a field has been set.

### GetDestinationIp

`func (o *OswRoutingOpenApiVO) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *OswRoutingOpenApiVO) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *OswRoutingOpenApiVO) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.

### HasDestinationIp

`func (o *OswRoutingOpenApiVO) HasDestinationIp() bool`

HasDestinationIp returns a boolean if a field has been set.

### GetDistance

`func (o *OswRoutingOpenApiVO) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *OswRoutingOpenApiVO) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *OswRoutingOpenApiVO) SetDistance(v int32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *OswRoutingOpenApiVO) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetInterfaceName

`func (o *OswRoutingOpenApiVO) GetInterfaceName() []string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *OswRoutingOpenApiVO) GetInterfaceNameOk() (*[]string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *OswRoutingOpenApiVO) SetInterfaceName(v []string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *OswRoutingOpenApiVO) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetMac

`func (o *OswRoutingOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswRoutingOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswRoutingOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswRoutingOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMetric

`func (o *OswRoutingOpenApiVO) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *OswRoutingOpenApiVO) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *OswRoutingOpenApiVO) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *OswRoutingOpenApiVO) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetName

`func (o *OswRoutingOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswRoutingOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswRoutingOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswRoutingOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextHop

`func (o *OswRoutingOpenApiVO) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *OswRoutingOpenApiVO) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *OswRoutingOpenApiVO) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *OswRoutingOpenApiVO) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetNextHopVrf

`func (o *OswRoutingOpenApiVO) GetNextHopVrf() string`

GetNextHopVrf returns the NextHopVrf field if non-nil, zero value otherwise.

### GetNextHopVrfOk

`func (o *OswRoutingOpenApiVO) GetNextHopVrfOk() (*string, bool)`

GetNextHopVrfOk returns a tuple with the NextHopVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopVrf

`func (o *OswRoutingOpenApiVO) SetNextHopVrf(v string)`

SetNextHopVrf sets NextHopVrf field to given value.

### HasNextHopVrf

`func (o *OswRoutingOpenApiVO) HasNextHopVrf() bool`

HasNextHopVrf returns a boolean if a field has been set.

### GetNextHops

`func (o *OswRoutingOpenApiVO) GetNextHops() []string`

GetNextHops returns the NextHops field if non-nil, zero value otherwise.

### GetNextHopsOk

`func (o *OswRoutingOpenApiVO) GetNextHopsOk() (*[]string, bool)`

GetNextHopsOk returns a tuple with the NextHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHops

`func (o *OswRoutingOpenApiVO) SetNextHops(v []string)`

SetNextHops sets NextHops field to given value.

### HasNextHops

`func (o *OswRoutingOpenApiVO) HasNextHops() bool`

HasNextHops returns a boolean if a field has been set.

### GetNextHopsVrf

`func (o *OswRoutingOpenApiVO) GetNextHopsVrf() []string`

GetNextHopsVrf returns the NextHopsVrf field if non-nil, zero value otherwise.

### GetNextHopsVrfOk

`func (o *OswRoutingOpenApiVO) GetNextHopsVrfOk() (*[]string, bool)`

GetNextHopsVrfOk returns a tuple with the NextHopsVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopsVrf

`func (o *OswRoutingOpenApiVO) SetNextHopsVrf(v []string)`

SetNextHopsVrf sets NextHopsVrf field to given value.

### HasNextHopsVrf

`func (o *OswRoutingOpenApiVO) HasNextHopsVrf() bool`

HasNextHopsVrf returns a boolean if a field has been set.

### GetOswDestinationIps

`func (o *OswRoutingOpenApiVO) GetOswDestinationIps() []string`

GetOswDestinationIps returns the OswDestinationIps field if non-nil, zero value otherwise.

### GetOswDestinationIpsOk

`func (o *OswRoutingOpenApiVO) GetOswDestinationIpsOk() (*[]string, bool)`

GetOswDestinationIpsOk returns a tuple with the OswDestinationIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswDestinationIps

`func (o *OswRoutingOpenApiVO) SetOswDestinationIps(v []string)`

SetOswDestinationIps sets OswDestinationIps field to given value.

### HasOswDestinationIps

`func (o *OswRoutingOpenApiVO) HasOswDestinationIps() bool`

HasOswDestinationIps returns a boolean if a field has been set.

### GetStackMsg

`func (o *OswRoutingOpenApiVO) GetStackMsg() StackMsgOpenApiVO`

GetStackMsg returns the StackMsg field if non-nil, zero value otherwise.

### GetStackMsgOk

`func (o *OswRoutingOpenApiVO) GetStackMsgOk() (*StackMsgOpenApiVO, bool)`

GetStackMsgOk returns a tuple with the StackMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackMsg

`func (o *OswRoutingOpenApiVO) SetStackMsg(v StackMsgOpenApiVO)`

SetStackMsg sets StackMsg field to given value.

### HasStackMsg

`func (o *OswRoutingOpenApiVO) HasStackMsg() bool`

HasStackMsg returns a boolean if a field has been set.

### GetType

`func (o *OswRoutingOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswRoutingOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswRoutingOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OswRoutingOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


