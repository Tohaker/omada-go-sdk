# AutoSelectWanPortReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | Pointer to **string** | The MAC of a SD-WAN candidate device. | [optional] 
**LinkedSpokes** | Pointer to **[]string** | A list MAC of linked-spokes of the sdWan group. | [optional] 
**Role** | **int32** | The role of SD-WAN member, hub: 0 or spoke: 1. | 
**SiteId** | Pointer to **string** | Site ID. | [optional] 

## Methods

### NewAutoSelectWanPortReq

`func NewAutoSelectWanPortReq(role int32, ) *AutoSelectWanPortReq`

NewAutoSelectWanPortReq instantiates a new AutoSelectWanPortReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoSelectWanPortReqWithDefaults

`func NewAutoSelectWanPortReqWithDefaults() *AutoSelectWanPortReq`

NewAutoSelectWanPortReqWithDefaults instantiates a new AutoSelectWanPortReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *AutoSelectWanPortReq) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *AutoSelectWanPortReq) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *AutoSelectWanPortReq) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *AutoSelectWanPortReq) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetLinkedSpokes

`func (o *AutoSelectWanPortReq) GetLinkedSpokes() []string`

GetLinkedSpokes returns the LinkedSpokes field if non-nil, zero value otherwise.

### GetLinkedSpokesOk

`func (o *AutoSelectWanPortReq) GetLinkedSpokesOk() (*[]string, bool)`

GetLinkedSpokesOk returns a tuple with the LinkedSpokes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedSpokes

`func (o *AutoSelectWanPortReq) SetLinkedSpokes(v []string)`

SetLinkedSpokes sets LinkedSpokes field to given value.

### HasLinkedSpokes

`func (o *AutoSelectWanPortReq) HasLinkedSpokes() bool`

HasLinkedSpokes returns a boolean if a field has been set.

### GetRole

`func (o *AutoSelectWanPortReq) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AutoSelectWanPortReq) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AutoSelectWanPortReq) SetRole(v int32)`

SetRole sets Role field to given value.


### GetSiteId

`func (o *AutoSelectWanPortReq) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AutoSelectWanPortReq) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AutoSelectWanPortReq) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AutoSelectWanPortReq) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


