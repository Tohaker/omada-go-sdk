# ClientRoamingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int64** | Duration, unit: ms. | [optional] 
**Start** | Pointer to **int64** | Roaming start timestamp, unit: ms. | [optional] 

## Methods

### NewClientRoamingInfo

`func NewClientRoamingInfo() *ClientRoamingInfo`

NewClientRoamingInfo instantiates a new ClientRoamingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientRoamingInfoWithDefaults

`func NewClientRoamingInfoWithDefaults() *ClientRoamingInfo`

NewClientRoamingInfoWithDefaults instantiates a new ClientRoamingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *ClientRoamingInfo) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ClientRoamingInfo) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ClientRoamingInfo) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ClientRoamingInfo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStart

`func (o *ClientRoamingInfo) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ClientRoamingInfo) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ClientRoamingInfo) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *ClientRoamingInfo) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


