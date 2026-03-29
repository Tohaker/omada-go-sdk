# ClientAssociationActivities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From1To3sNum** | Pointer to **int32** | Number of clients with association time less than 3 seconds and greater than or equal to 1 second. | [optional] 
**From3To12sNum** | Pointer to **int32** | Number of clients with association time less than 12 seconds and greater than or equal to 3 second. | [optional] 
**LessThan1sNum** | Pointer to **int32** | Number of clients with association time less than 1 second. | [optional] 
**MoreThan12sNum** | Pointer to **int32** | Number of clients with association time greater than or equal to 12 second. | [optional] 
**Time** | Pointer to **int64** | Timestamp, unit is second. | [optional] 

## Methods

### NewClientAssociationActivities

`func NewClientAssociationActivities() *ClientAssociationActivities`

NewClientAssociationActivities instantiates a new ClientAssociationActivities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAssociationActivitiesWithDefaults

`func NewClientAssociationActivitiesWithDefaults() *ClientAssociationActivities`

NewClientAssociationActivitiesWithDefaults instantiates a new ClientAssociationActivities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom1To3sNum

`func (o *ClientAssociationActivities) GetFrom1To3sNum() int32`

GetFrom1To3sNum returns the From1To3sNum field if non-nil, zero value otherwise.

### GetFrom1To3sNumOk

`func (o *ClientAssociationActivities) GetFrom1To3sNumOk() (*int32, bool)`

GetFrom1To3sNumOk returns a tuple with the From1To3sNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom1To3sNum

`func (o *ClientAssociationActivities) SetFrom1To3sNum(v int32)`

SetFrom1To3sNum sets From1To3sNum field to given value.

### HasFrom1To3sNum

`func (o *ClientAssociationActivities) HasFrom1To3sNum() bool`

HasFrom1To3sNum returns a boolean if a field has been set.

### GetFrom3To12sNum

`func (o *ClientAssociationActivities) GetFrom3To12sNum() int32`

GetFrom3To12sNum returns the From3To12sNum field if non-nil, zero value otherwise.

### GetFrom3To12sNumOk

`func (o *ClientAssociationActivities) GetFrom3To12sNumOk() (*int32, bool)`

GetFrom3To12sNumOk returns a tuple with the From3To12sNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom3To12sNum

`func (o *ClientAssociationActivities) SetFrom3To12sNum(v int32)`

SetFrom3To12sNum sets From3To12sNum field to given value.

### HasFrom3To12sNum

`func (o *ClientAssociationActivities) HasFrom3To12sNum() bool`

HasFrom3To12sNum returns a boolean if a field has been set.

### GetLessThan1sNum

`func (o *ClientAssociationActivities) GetLessThan1sNum() int32`

GetLessThan1sNum returns the LessThan1sNum field if non-nil, zero value otherwise.

### GetLessThan1sNumOk

`func (o *ClientAssociationActivities) GetLessThan1sNumOk() (*int32, bool)`

GetLessThan1sNumOk returns a tuple with the LessThan1sNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLessThan1sNum

`func (o *ClientAssociationActivities) SetLessThan1sNum(v int32)`

SetLessThan1sNum sets LessThan1sNum field to given value.

### HasLessThan1sNum

`func (o *ClientAssociationActivities) HasLessThan1sNum() bool`

HasLessThan1sNum returns a boolean if a field has been set.

### GetMoreThan12sNum

`func (o *ClientAssociationActivities) GetMoreThan12sNum() int32`

GetMoreThan12sNum returns the MoreThan12sNum field if non-nil, zero value otherwise.

### GetMoreThan12sNumOk

`func (o *ClientAssociationActivities) GetMoreThan12sNumOk() (*int32, bool)`

GetMoreThan12sNumOk returns a tuple with the MoreThan12sNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreThan12sNum

`func (o *ClientAssociationActivities) SetMoreThan12sNum(v int32)`

SetMoreThan12sNum sets MoreThan12sNum field to given value.

### HasMoreThan12sNum

`func (o *ClientAssociationActivities) HasMoreThan12sNum() bool`

HasMoreThan12sNum returns a boolean if a field has been set.

### GetTime

`func (o *ClientAssociationActivities) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ClientAssociationActivities) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ClientAssociationActivities) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ClientAssociationActivities) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


