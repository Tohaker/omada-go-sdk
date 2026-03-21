# TrafficProfileDeleteResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfilesInUse** | Pointer to **[]int32** | A list of Traffic Profile IDs that failed to delete due to being in use. | [optional] 

## Methods

### NewTrafficProfileDeleteResultDTO

`func NewTrafficProfileDeleteResultDTO() *TrafficProfileDeleteResultDTO`

NewTrafficProfileDeleteResultDTO instantiates a new TrafficProfileDeleteResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficProfileDeleteResultDTOWithDefaults

`func NewTrafficProfileDeleteResultDTOWithDefaults() *TrafficProfileDeleteResultDTO`

NewTrafficProfileDeleteResultDTOWithDefaults instantiates a new TrafficProfileDeleteResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfilesInUse

`func (o *TrafficProfileDeleteResultDTO) GetProfilesInUse() []int32`

GetProfilesInUse returns the ProfilesInUse field if non-nil, zero value otherwise.

### GetProfilesInUseOk

`func (o *TrafficProfileDeleteResultDTO) GetProfilesInUseOk() (*[]int32, bool)`

GetProfilesInUseOk returns a tuple with the ProfilesInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilesInUse

`func (o *TrafficProfileDeleteResultDTO) SetProfilesInUse(v []int32)`

SetProfilesInUse sets ProfilesInUse field to given value.

### HasProfilesInUse

`func (o *TrafficProfileDeleteResultDTO) HasProfilesInUse() bool`

HasProfilesInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


