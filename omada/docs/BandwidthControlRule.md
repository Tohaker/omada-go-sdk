# BandwidthControlRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownstreamBandwidth** | **int32** | Downstream bandwidth should be within the range of 1–9999999. | 
**DownstreamBandwidthUnit** | **int32** | Downstream bandwidth unit should be a value as follows: 1: Kbps, 2: Mbps. | 
**Id** | Pointer to **string** | ID of the bandwidth control rule. | [optional] 
**Index** | Pointer to **int32** | Index of the bandwidth control rule. | [optional] 
**Mode** | **int32** | Mode should be a value as follows: 0: share; 1: individual. | 
**Name** | **string** | Name should contain 1 to 64 characters. | 
**SourceIds** | **[]string** | Source IDs of the bandwidth control rule. Network can be created using &#39;Create LAN network&#39; interface, and network ID can be obtained from &#39;Get LAN network list&#39; interface. IP group can be created using &#39;Create a new group profile&#39; interface, and IP group ID can be obtained from &#39;Get group profile list&#39; interface. | 
**SourceType** | Pointer to **int32** | Source type should be a value as follows: 0: network; 1: IP group. | [optional] 
**Status** | **bool** | Status of the bandwidth control rule. | 
**UpstreamBandwidth** | **int32** | Upstream bandwidth should be within the range of 1–9999999. | 
**UpstreamBandwidthUnit** | **int32** | Upstream bandwidth unit should be a value as follows: 1: Kbps, 2: Mbps. | 
**WanPortIds** | **[]string** | WAN port IDs of the bandwidth control rule.WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | 

## Methods

### NewBandwidthControlRule

`func NewBandwidthControlRule(downstreamBandwidth int32, downstreamBandwidthUnit int32, mode int32, name string, sourceIds []string, status bool, upstreamBandwidth int32, upstreamBandwidthUnit int32, wanPortIds []string, ) *BandwidthControlRule`

NewBandwidthControlRule instantiates a new BandwidthControlRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthControlRuleWithDefaults

`func NewBandwidthControlRuleWithDefaults() *BandwidthControlRule`

NewBandwidthControlRuleWithDefaults instantiates a new BandwidthControlRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownstreamBandwidth

`func (o *BandwidthControlRule) GetDownstreamBandwidth() int32`

GetDownstreamBandwidth returns the DownstreamBandwidth field if non-nil, zero value otherwise.

### GetDownstreamBandwidthOk

`func (o *BandwidthControlRule) GetDownstreamBandwidthOk() (*int32, bool)`

GetDownstreamBandwidthOk returns a tuple with the DownstreamBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamBandwidth

`func (o *BandwidthControlRule) SetDownstreamBandwidth(v int32)`

SetDownstreamBandwidth sets DownstreamBandwidth field to given value.


### GetDownstreamBandwidthUnit

`func (o *BandwidthControlRule) GetDownstreamBandwidthUnit() int32`

GetDownstreamBandwidthUnit returns the DownstreamBandwidthUnit field if non-nil, zero value otherwise.

### GetDownstreamBandwidthUnitOk

`func (o *BandwidthControlRule) GetDownstreamBandwidthUnitOk() (*int32, bool)`

GetDownstreamBandwidthUnitOk returns a tuple with the DownstreamBandwidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamBandwidthUnit

`func (o *BandwidthControlRule) SetDownstreamBandwidthUnit(v int32)`

SetDownstreamBandwidthUnit sets DownstreamBandwidthUnit field to given value.


### GetId

`func (o *BandwidthControlRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BandwidthControlRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BandwidthControlRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BandwidthControlRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *BandwidthControlRule) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BandwidthControlRule) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BandwidthControlRule) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *BandwidthControlRule) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetMode

`func (o *BandwidthControlRule) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BandwidthControlRule) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BandwidthControlRule) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetName

`func (o *BandwidthControlRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BandwidthControlRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BandwidthControlRule) SetName(v string)`

SetName sets Name field to given value.


### GetSourceIds

`func (o *BandwidthControlRule) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *BandwidthControlRule) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *BandwidthControlRule) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.


### GetSourceType

`func (o *BandwidthControlRule) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *BandwidthControlRule) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *BandwidthControlRule) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *BandwidthControlRule) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetStatus

`func (o *BandwidthControlRule) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BandwidthControlRule) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BandwidthControlRule) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetUpstreamBandwidth

`func (o *BandwidthControlRule) GetUpstreamBandwidth() int32`

GetUpstreamBandwidth returns the UpstreamBandwidth field if non-nil, zero value otherwise.

### GetUpstreamBandwidthOk

`func (o *BandwidthControlRule) GetUpstreamBandwidthOk() (*int32, bool)`

GetUpstreamBandwidthOk returns a tuple with the UpstreamBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamBandwidth

`func (o *BandwidthControlRule) SetUpstreamBandwidth(v int32)`

SetUpstreamBandwidth sets UpstreamBandwidth field to given value.


### GetUpstreamBandwidthUnit

`func (o *BandwidthControlRule) GetUpstreamBandwidthUnit() int32`

GetUpstreamBandwidthUnit returns the UpstreamBandwidthUnit field if non-nil, zero value otherwise.

### GetUpstreamBandwidthUnitOk

`func (o *BandwidthControlRule) GetUpstreamBandwidthUnitOk() (*int32, bool)`

GetUpstreamBandwidthUnitOk returns a tuple with the UpstreamBandwidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamBandwidthUnit

`func (o *BandwidthControlRule) SetUpstreamBandwidthUnit(v int32)`

SetUpstreamBandwidthUnit sets UpstreamBandwidthUnit field to given value.


### GetWanPortIds

`func (o *BandwidthControlRule) GetWanPortIds() []string`

GetWanPortIds returns the WanPortIds field if non-nil, zero value otherwise.

### GetWanPortIdsOk

`func (o *BandwidthControlRule) GetWanPortIdsOk() (*[]string, bool)`

GetWanPortIdsOk returns a tuple with the WanPortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortIds

`func (o *BandwidthControlRule) SetWanPortIds(v []string)`

SetWanPortIds sets WanPortIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


