# DBAProfileDeleteResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfilesInUse** | Pointer to **[]int32** | List of DBA Profile IDs that failed to delete due to being in use. | [optional] 

## Methods

### NewDBAProfileDeleteResultDTO

`func NewDBAProfileDeleteResultDTO() *DBAProfileDeleteResultDTO`

NewDBAProfileDeleteResultDTO instantiates a new DBAProfileDeleteResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDBAProfileDeleteResultDTOWithDefaults

`func NewDBAProfileDeleteResultDTOWithDefaults() *DBAProfileDeleteResultDTO`

NewDBAProfileDeleteResultDTOWithDefaults instantiates a new DBAProfileDeleteResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfilesInUse

`func (o *DBAProfileDeleteResultDTO) GetProfilesInUse() []int32`

GetProfilesInUse returns the ProfilesInUse field if non-nil, zero value otherwise.

### GetProfilesInUseOk

`func (o *DBAProfileDeleteResultDTO) GetProfilesInUseOk() (*[]int32, bool)`

GetProfilesInUseOk returns a tuple with the ProfilesInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilesInUse

`func (o *DBAProfileDeleteResultDTO) SetProfilesInUse(v []int32)`

SetProfilesInUse sets ProfilesInUse field to given value.

### HasProfilesInUse

`func (o *DBAProfileDeleteResultDTO) HasProfilesInUse() bool`

HasProfilesInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


