# ServiceProfileDeleteResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfilesInUse** | Pointer to **[]int32** | List of Service Profile IDs that failed to delete due to being in use. | [optional] 

## Methods

### NewServiceProfileDeleteResultDTO

`func NewServiceProfileDeleteResultDTO() *ServiceProfileDeleteResultDTO`

NewServiceProfileDeleteResultDTO instantiates a new ServiceProfileDeleteResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileDeleteResultDTOWithDefaults

`func NewServiceProfileDeleteResultDTOWithDefaults() *ServiceProfileDeleteResultDTO`

NewServiceProfileDeleteResultDTOWithDefaults instantiates a new ServiceProfileDeleteResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfilesInUse

`func (o *ServiceProfileDeleteResultDTO) GetProfilesInUse() []int32`

GetProfilesInUse returns the ProfilesInUse field if non-nil, zero value otherwise.

### GetProfilesInUseOk

`func (o *ServiceProfileDeleteResultDTO) GetProfilesInUseOk() (*[]int32, bool)`

GetProfilesInUseOk returns a tuple with the ProfilesInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilesInUse

`func (o *ServiceProfileDeleteResultDTO) SetProfilesInUse(v []int32)`

SetProfilesInUse sets ProfilesInUse field to given value.

### HasProfilesInUse

`func (o *ServiceProfileDeleteResultDTO) HasProfilesInUse() bool`

HasProfilesInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


