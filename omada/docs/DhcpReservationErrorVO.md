# DhcpReservationErrorVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMap** | Pointer to **map[string]int32** | Error details when importing some entries fails: 0: The imported MAC entries conflict; 1: The server device&#39;s user list has reached its limit. | [optional] 

## Methods

### NewDhcpReservationErrorVO

`func NewDhcpReservationErrorVO() *DhcpReservationErrorVO`

NewDhcpReservationErrorVO instantiates a new DhcpReservationErrorVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpReservationErrorVOWithDefaults

`func NewDhcpReservationErrorVOWithDefaults() *DhcpReservationErrorVO`

NewDhcpReservationErrorVOWithDefaults instantiates a new DhcpReservationErrorVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMap

`func (o *DhcpReservationErrorVO) GetErrorMap() map[string]int32`

GetErrorMap returns the ErrorMap field if non-nil, zero value otherwise.

### GetErrorMapOk

`func (o *DhcpReservationErrorVO) GetErrorMapOk() (*map[string]int32, bool)`

GetErrorMapOk returns a tuple with the ErrorMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMap

`func (o *DhcpReservationErrorVO) SetErrorMap(v map[string]int32)`

SetErrorMap sets ErrorMap field to given value.

### HasErrorMap

`func (o *DhcpReservationErrorVO) HasErrorMap() bool`

HasErrorMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


