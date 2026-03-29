# OswStaticRoutingConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | **[]string** | IP address/SubNet, up to 16 entries are allowed for the destinations list. | 
**Distance** | **int32** | Distance should be within the range of 1–255. | 
**IpVersion** | **int32** | The IP Version of class rule should be a value as follows: 0: IPv4; 1: IPv6. | 
**NextHopIp** | **string** | NextHopIp | 
**Status** | **bool** | StaticRouting status | 

## Methods

### NewOswStaticRoutingConfigOpenApiVO

`func NewOswStaticRoutingConfigOpenApiVO(destinations []string, distance int32, ipVersion int32, nextHopIp string, status bool, ) *OswStaticRoutingConfigOpenApiVO`

NewOswStaticRoutingConfigOpenApiVO instantiates a new OswStaticRoutingConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStaticRoutingConfigOpenApiVOWithDefaults

`func NewOswStaticRoutingConfigOpenApiVOWithDefaults() *OswStaticRoutingConfigOpenApiVO`

NewOswStaticRoutingConfigOpenApiVOWithDefaults instantiates a new OswStaticRoutingConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *OswStaticRoutingConfigOpenApiVO) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *OswStaticRoutingConfigOpenApiVO) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *OswStaticRoutingConfigOpenApiVO) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.


### GetDistance

`func (o *OswStaticRoutingConfigOpenApiVO) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *OswStaticRoutingConfigOpenApiVO) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *OswStaticRoutingConfigOpenApiVO) SetDistance(v int32)`

SetDistance sets Distance field to given value.


### GetIpVersion

`func (o *OswStaticRoutingConfigOpenApiVO) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *OswStaticRoutingConfigOpenApiVO) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *OswStaticRoutingConfigOpenApiVO) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.


### GetNextHopIp

`func (o *OswStaticRoutingConfigOpenApiVO) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *OswStaticRoutingConfigOpenApiVO) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *OswStaticRoutingConfigOpenApiVO) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.


### GetStatus

`func (o *OswStaticRoutingConfigOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswStaticRoutingConfigOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswStaticRoutingConfigOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


