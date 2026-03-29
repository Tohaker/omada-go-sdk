# OltDevCapVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagCount** | Pointer to **int32** | Number of LAGs(Link Aggregation Group) | [optional] 
**NetworkCheckSupport** | Pointer to **bool** | Whether this OLT supports network check function | [optional] 
**PingSupport** | Pointer to **bool** | Whether this OLT supports ping function | [optional] 
**PonPortCount** | Pointer to **int32** | Number of pon ports | [optional] 
**TraceRouteSupport** | Pointer to **bool** | Whether this OLT supports traceRoute function | [optional] 

## Methods

### NewOltDevCapVO

`func NewOltDevCapVO() *OltDevCapVO`

NewOltDevCapVO instantiates a new OltDevCapVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOltDevCapVOWithDefaults

`func NewOltDevCapVOWithDefaults() *OltDevCapVO`

NewOltDevCapVOWithDefaults instantiates a new OltDevCapVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagCount

`func (o *OltDevCapVO) GetLagCount() int32`

GetLagCount returns the LagCount field if non-nil, zero value otherwise.

### GetLagCountOk

`func (o *OltDevCapVO) GetLagCountOk() (*int32, bool)`

GetLagCountOk returns a tuple with the LagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagCount

`func (o *OltDevCapVO) SetLagCount(v int32)`

SetLagCount sets LagCount field to given value.

### HasLagCount

`func (o *OltDevCapVO) HasLagCount() bool`

HasLagCount returns a boolean if a field has been set.

### GetNetworkCheckSupport

`func (o *OltDevCapVO) GetNetworkCheckSupport() bool`

GetNetworkCheckSupport returns the NetworkCheckSupport field if non-nil, zero value otherwise.

### GetNetworkCheckSupportOk

`func (o *OltDevCapVO) GetNetworkCheckSupportOk() (*bool, bool)`

GetNetworkCheckSupportOk returns a tuple with the NetworkCheckSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCheckSupport

`func (o *OltDevCapVO) SetNetworkCheckSupport(v bool)`

SetNetworkCheckSupport sets NetworkCheckSupport field to given value.

### HasNetworkCheckSupport

`func (o *OltDevCapVO) HasNetworkCheckSupport() bool`

HasNetworkCheckSupport returns a boolean if a field has been set.

### GetPingSupport

`func (o *OltDevCapVO) GetPingSupport() bool`

GetPingSupport returns the PingSupport field if non-nil, zero value otherwise.

### GetPingSupportOk

`func (o *OltDevCapVO) GetPingSupportOk() (*bool, bool)`

GetPingSupportOk returns a tuple with the PingSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSupport

`func (o *OltDevCapVO) SetPingSupport(v bool)`

SetPingSupport sets PingSupport field to given value.

### HasPingSupport

`func (o *OltDevCapVO) HasPingSupport() bool`

HasPingSupport returns a boolean if a field has been set.

### GetPonPortCount

`func (o *OltDevCapVO) GetPonPortCount() int32`

GetPonPortCount returns the PonPortCount field if non-nil, zero value otherwise.

### GetPonPortCountOk

`func (o *OltDevCapVO) GetPonPortCountOk() (*int32, bool)`

GetPonPortCountOk returns a tuple with the PonPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPortCount

`func (o *OltDevCapVO) SetPonPortCount(v int32)`

SetPonPortCount sets PonPortCount field to given value.

### HasPonPortCount

`func (o *OltDevCapVO) HasPonPortCount() bool`

HasPonPortCount returns a boolean if a field has been set.

### GetTraceRouteSupport

`func (o *OltDevCapVO) GetTraceRouteSupport() bool`

GetTraceRouteSupport returns the TraceRouteSupport field if non-nil, zero value otherwise.

### GetTraceRouteSupportOk

`func (o *OltDevCapVO) GetTraceRouteSupportOk() (*bool, bool)`

GetTraceRouteSupportOk returns a tuple with the TraceRouteSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteSupport

`func (o *OltDevCapVO) SetTraceRouteSupport(v bool)`

SetTraceRouteSupport sets TraceRouteSupport field to given value.

### HasTraceRouteSupport

`func (o *OltDevCapVO) HasTraceRouteSupport() bool`

HasTraceRouteSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


