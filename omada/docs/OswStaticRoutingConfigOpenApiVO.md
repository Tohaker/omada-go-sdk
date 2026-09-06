# OswStaticRoutingConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of static routing. It may contain 0 to 128 characters, including digits (0–9), uppercase and lowercase letters (A–Z, a–z), spaces, and -_@:/.+# . | [optional] 
**Destinations** | **[]string** | IP address/SubNet, up to 16 entries are allowed for the destinations list. | 
**Distance** | **int32** | Distance should be within the range of 1–255. | 
**IpVersion** | **int32** | The IP Version of class rule should be a value as follows: 0: IPv4; 1: IPv6. | 
**NextHopIp** | **string** | NextHopIp | 
**NextHopVrfId** | Pointer to **string** | NextHopVrfId | [optional] 
**Status** | **bool** | StaticRouting status | 
**VrfId** | Pointer to **string** | VrfId | [optional] 

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

### GetDescription

`func (o *OswStaticRoutingConfigOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OswStaticRoutingConfigOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OswStaticRoutingConfigOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OswStaticRoutingConfigOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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


### GetNextHopVrfId

`func (o *OswStaticRoutingConfigOpenApiVO) GetNextHopVrfId() string`

GetNextHopVrfId returns the NextHopVrfId field if non-nil, zero value otherwise.

### GetNextHopVrfIdOk

`func (o *OswStaticRoutingConfigOpenApiVO) GetNextHopVrfIdOk() (*string, bool)`

GetNextHopVrfIdOk returns a tuple with the NextHopVrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopVrfId

`func (o *OswStaticRoutingConfigOpenApiVO) SetNextHopVrfId(v string)`

SetNextHopVrfId sets NextHopVrfId field to given value.

### HasNextHopVrfId

`func (o *OswStaticRoutingConfigOpenApiVO) HasNextHopVrfId() bool`

HasNextHopVrfId returns a boolean if a field has been set.

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


### GetVrfId

`func (o *OswStaticRoutingConfigOpenApiVO) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *OswStaticRoutingConfigOpenApiVO) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *OswStaticRoutingConfigOpenApiVO) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *OswStaticRoutingConfigOpenApiVO) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


