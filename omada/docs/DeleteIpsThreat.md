# DeleteIpsThreat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreatId** | Pointer to [**[]IpsOperateThreatIdAndTime**](IpsOperateThreatIdAndTime.md) | IPS threatId and Time list. | [optional] 

## Methods

### NewDeleteIpsThreat

`func NewDeleteIpsThreat() *DeleteIpsThreat`

NewDeleteIpsThreat instantiates a new DeleteIpsThreat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteIpsThreatWithDefaults

`func NewDeleteIpsThreatWithDefaults() *DeleteIpsThreat`

NewDeleteIpsThreatWithDefaults instantiates a new DeleteIpsThreat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreatId

`func (o *DeleteIpsThreat) GetThreatId() []IpsOperateThreatIdAndTime`

GetThreatId returns the ThreatId field if non-nil, zero value otherwise.

### GetThreatIdOk

`func (o *DeleteIpsThreat) GetThreatIdOk() (*[]IpsOperateThreatIdAndTime, bool)`

GetThreatIdOk returns a tuple with the ThreatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatId

`func (o *DeleteIpsThreat) SetThreatId(v []IpsOperateThreatIdAndTime)`

SetThreatId sets ThreatId field to given value.

### HasThreatId

`func (o *DeleteIpsThreat) HasThreatId() bool`

HasThreatId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


