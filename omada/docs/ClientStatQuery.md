# ClientStatQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndSec** | **int64** | End timestamp, unit: second. | 
**StartSec** | **int64** | Start timestamp, unit: second. | 

## Methods

### NewClientStatQuery

`func NewClientStatQuery(endSec int64, startSec int64, ) *ClientStatQuery`

NewClientStatQuery instantiates a new ClientStatQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientStatQueryWithDefaults

`func NewClientStatQueryWithDefaults() *ClientStatQuery`

NewClientStatQueryWithDefaults instantiates a new ClientStatQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndSec

`func (o *ClientStatQuery) GetEndSec() int64`

GetEndSec returns the EndSec field if non-nil, zero value otherwise.

### GetEndSecOk

`func (o *ClientStatQuery) GetEndSecOk() (*int64, bool)`

GetEndSecOk returns a tuple with the EndSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSec

`func (o *ClientStatQuery) SetEndSec(v int64)`

SetEndSec sets EndSec field to given value.


### GetStartSec

`func (o *ClientStatQuery) GetStartSec() int64`

GetStartSec returns the StartSec field if non-nil, zero value otherwise.

### GetStartSecOk

`func (o *ClientStatQuery) GetStartSecOk() (*int64, bool)`

GetStartSecOk returns a tuple with the StartSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSec

`func (o *ClientStatQuery) SetStartSec(v int64)`

SetStartSec sets StartSec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


