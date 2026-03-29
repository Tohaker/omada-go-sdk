# ActivityRecordsOfAClientSSingleConnections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** | End timestamp, unit: millisecond. | [optional] 
**Events** | Pointer to [**[]ClientConnectionEvents**](ClientConnectionEvents.md) | Client connection events. | [optional] 
**Mac** | Pointer to **string** | Client MAC | [optional] 
**Start** | Pointer to **int64** | Start timestamp, unit: millisecond. | [optional] 

## Methods

### NewActivityRecordsOfAClientSSingleConnections

`func NewActivityRecordsOfAClientSSingleConnections() *ActivityRecordsOfAClientSSingleConnections`

NewActivityRecordsOfAClientSSingleConnections instantiates a new ActivityRecordsOfAClientSSingleConnections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityRecordsOfAClientSSingleConnectionsWithDefaults

`func NewActivityRecordsOfAClientSSingleConnectionsWithDefaults() *ActivityRecordsOfAClientSSingleConnections`

NewActivityRecordsOfAClientSSingleConnectionsWithDefaults instantiates a new ActivityRecordsOfAClientSSingleConnections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ActivityRecordsOfAClientSSingleConnections) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ActivityRecordsOfAClientSSingleConnections) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ActivityRecordsOfAClientSSingleConnections) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ActivityRecordsOfAClientSSingleConnections) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetEvents

`func (o *ActivityRecordsOfAClientSSingleConnections) GetEvents() []ClientConnectionEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ActivityRecordsOfAClientSSingleConnections) GetEventsOk() (*[]ClientConnectionEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ActivityRecordsOfAClientSSingleConnections) SetEvents(v []ClientConnectionEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ActivityRecordsOfAClientSSingleConnections) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetMac

`func (o *ActivityRecordsOfAClientSSingleConnections) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ActivityRecordsOfAClientSSingleConnections) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ActivityRecordsOfAClientSSingleConnections) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ActivityRecordsOfAClientSSingleConnections) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStart

`func (o *ActivityRecordsOfAClientSSingleConnections) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ActivityRecordsOfAClientSSingleConnections) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ActivityRecordsOfAClientSSingleConnections) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *ActivityRecordsOfAClientSSingleConnections) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


