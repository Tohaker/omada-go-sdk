# OswStaticRoutingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | **[]string** | IP address/SubNet, up to 16 entries are allowed for the destinations list. | 
**Distance** | **int32** | Distance should be within the range of 1–255. | 
**EntryId** | Pointer to **int32** | The entry ID of StaticRouting | [optional] 
**Id** | Pointer to **string** | StaticRouting ID | [optional] 
**IpVersion** | Pointer to **int32** | The IP Version of class rule should be a value as follows: 0: IPv4; 1: IPv6. | [optional] 
**Mac** | Pointer to **string** | Switch Mac | [optional] 
**Name** | Pointer to **string** | Switch name | [optional] 
**NextHopIp** | Pointer to **string** | NextHopIp | [optional] 
**NextHopVrfId** | Pointer to **string** |  | [optional] 
**OmadacId** | Pointer to **string** | Omada ID | [optional] 
**Resource** | Pointer to **int32** | Resource is a value as follows: 0: new created; 1: from template; 2: override | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**Status** | **bool** | StaticRouting status | 
**VrfId** | Pointer to **string** |  | [optional] 

## Methods

### NewOswStaticRoutingVO

`func NewOswStaticRoutingVO(destinations []string, distance int32, status bool, ) *OswStaticRoutingVO`

NewOswStaticRoutingVO instantiates a new OswStaticRoutingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStaticRoutingVOWithDefaults

`func NewOswStaticRoutingVOWithDefaults() *OswStaticRoutingVO`

NewOswStaticRoutingVOWithDefaults instantiates a new OswStaticRoutingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *OswStaticRoutingVO) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *OswStaticRoutingVO) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *OswStaticRoutingVO) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.


### GetDistance

`func (o *OswStaticRoutingVO) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *OswStaticRoutingVO) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *OswStaticRoutingVO) SetDistance(v int32)`

SetDistance sets Distance field to given value.


### GetEntryId

`func (o *OswStaticRoutingVO) GetEntryId() int32`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *OswStaticRoutingVO) GetEntryIdOk() (*int32, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *OswStaticRoutingVO) SetEntryId(v int32)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *OswStaticRoutingVO) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetId

`func (o *OswStaticRoutingVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswStaticRoutingVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswStaticRoutingVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswStaticRoutingVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpVersion

`func (o *OswStaticRoutingVO) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *OswStaticRoutingVO) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *OswStaticRoutingVO) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *OswStaticRoutingVO) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetMac

`func (o *OswStaticRoutingVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswStaticRoutingVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswStaticRoutingVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswStaticRoutingVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *OswStaticRoutingVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStaticRoutingVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStaticRoutingVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswStaticRoutingVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextHopIp

`func (o *OswStaticRoutingVO) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *OswStaticRoutingVO) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *OswStaticRoutingVO) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.

### HasNextHopIp

`func (o *OswStaticRoutingVO) HasNextHopIp() bool`

HasNextHopIp returns a boolean if a field has been set.

### GetNextHopVrfId

`func (o *OswStaticRoutingVO) GetNextHopVrfId() string`

GetNextHopVrfId returns the NextHopVrfId field if non-nil, zero value otherwise.

### GetNextHopVrfIdOk

`func (o *OswStaticRoutingVO) GetNextHopVrfIdOk() (*string, bool)`

GetNextHopVrfIdOk returns a tuple with the NextHopVrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopVrfId

`func (o *OswStaticRoutingVO) SetNextHopVrfId(v string)`

SetNextHopVrfId sets NextHopVrfId field to given value.

### HasNextHopVrfId

`func (o *OswStaticRoutingVO) HasNextHopVrfId() bool`

HasNextHopVrfId returns a boolean if a field has been set.

### GetOmadacId

`func (o *OswStaticRoutingVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *OswStaticRoutingVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *OswStaticRoutingVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *OswStaticRoutingVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetResource

`func (o *OswStaticRoutingVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OswStaticRoutingVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OswStaticRoutingVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OswStaticRoutingVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSiteId

`func (o *OswStaticRoutingVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *OswStaticRoutingVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *OswStaticRoutingVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *OswStaticRoutingVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStatus

`func (o *OswStaticRoutingVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswStaticRoutingVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswStaticRoutingVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetVrfId

`func (o *OswStaticRoutingVO) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *OswStaticRoutingVO) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *OswStaticRoutingVO) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *OswStaticRoutingVO) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


