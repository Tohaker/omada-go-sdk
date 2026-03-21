# LineProfileDeleteResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfilesInUse** | Pointer to **[]int32** | A list of Line Profile IDs that failed to delete because they are in use. | [optional] 

## Methods

### NewLineProfileDeleteResultDTO

`func NewLineProfileDeleteResultDTO() *LineProfileDeleteResultDTO`

NewLineProfileDeleteResultDTO instantiates a new LineProfileDeleteResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineProfileDeleteResultDTOWithDefaults

`func NewLineProfileDeleteResultDTOWithDefaults() *LineProfileDeleteResultDTO`

NewLineProfileDeleteResultDTOWithDefaults instantiates a new LineProfileDeleteResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfilesInUse

`func (o *LineProfileDeleteResultDTO) GetProfilesInUse() []int32`

GetProfilesInUse returns the ProfilesInUse field if non-nil, zero value otherwise.

### GetProfilesInUseOk

`func (o *LineProfileDeleteResultDTO) GetProfilesInUseOk() (*[]int32, bool)`

GetProfilesInUseOk returns a tuple with the ProfilesInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilesInUse

`func (o *LineProfileDeleteResultDTO) SetProfilesInUse(v []int32)`

SetProfilesInUse sets ProfilesInUse field to given value.

### HasProfilesInUse

`func (o *LineProfileDeleteResultDTO) HasProfilesInUse() bool`

HasProfilesInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


