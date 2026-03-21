# DdnsIntervalOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomInterval** | Pointer to **int32** | Dynamic DNS custom interval, valid when parameter [server] is 0, 1 or 4, unit: minute. CustomInterval should be within the range of 1-60 | [optional] 
**UpdateInterval** | Pointer to **int32** | Dynamic DNS update interval, unit: hour. UpdateInterval should be a value as follows: 0, 1, 6, 12, 24, 48 or 72 | [optional] 

## Methods

### NewDdnsIntervalOpenApiVO

`func NewDdnsIntervalOpenApiVO() *DdnsIntervalOpenApiVO`

NewDdnsIntervalOpenApiVO instantiates a new DdnsIntervalOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdnsIntervalOpenApiVOWithDefaults

`func NewDdnsIntervalOpenApiVOWithDefaults() *DdnsIntervalOpenApiVO`

NewDdnsIntervalOpenApiVOWithDefaults instantiates a new DdnsIntervalOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomInterval

`func (o *DdnsIntervalOpenApiVO) GetCustomInterval() int32`

GetCustomInterval returns the CustomInterval field if non-nil, zero value otherwise.

### GetCustomIntervalOk

`func (o *DdnsIntervalOpenApiVO) GetCustomIntervalOk() (*int32, bool)`

GetCustomIntervalOk returns a tuple with the CustomInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInterval

`func (o *DdnsIntervalOpenApiVO) SetCustomInterval(v int32)`

SetCustomInterval sets CustomInterval field to given value.

### HasCustomInterval

`func (o *DdnsIntervalOpenApiVO) HasCustomInterval() bool`

HasCustomInterval returns a boolean if a field has been set.

### GetUpdateInterval

`func (o *DdnsIntervalOpenApiVO) GetUpdateInterval() int32`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *DdnsIntervalOpenApiVO) GetUpdateIntervalOk() (*int32, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *DdnsIntervalOpenApiVO) SetUpdateInterval(v int32)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *DdnsIntervalOpenApiVO) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


