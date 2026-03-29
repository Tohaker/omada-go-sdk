# ClientConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int64** | Duration, unit: ms | [optional] 
**End** | Pointer to **int64** | End timestamp, unit: ms | [optional] 
**Start** | Pointer to **int64** | Start timestamp, unit: ms | [optional] 

## Methods

### NewClientConnectionInfo

`func NewClientConnectionInfo() *ClientConnectionInfo`

NewClientConnectionInfo instantiates a new ClientConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConnectionInfoWithDefaults

`func NewClientConnectionInfoWithDefaults() *ClientConnectionInfo`

NewClientConnectionInfoWithDefaults instantiates a new ClientConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *ClientConnectionInfo) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ClientConnectionInfo) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ClientConnectionInfo) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ClientConnectionInfo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnd

`func (o *ClientConnectionInfo) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ClientConnectionInfo) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ClientConnectionInfo) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ClientConnectionInfo) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *ClientConnectionInfo) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ClientConnectionInfo) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ClientConnectionInfo) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *ClientConnectionInfo) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


