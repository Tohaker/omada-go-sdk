# TcontDeleteResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TcontsInUse** | Pointer to **[]int32** |  List of T-cont IDs that failed to delete because they are in use. | [optional] 

## Methods

### NewTcontDeleteResultDTO

`func NewTcontDeleteResultDTO() *TcontDeleteResultDTO`

NewTcontDeleteResultDTO instantiates a new TcontDeleteResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTcontDeleteResultDTOWithDefaults

`func NewTcontDeleteResultDTOWithDefaults() *TcontDeleteResultDTO`

NewTcontDeleteResultDTOWithDefaults instantiates a new TcontDeleteResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTcontsInUse

`func (o *TcontDeleteResultDTO) GetTcontsInUse() []int32`

GetTcontsInUse returns the TcontsInUse field if non-nil, zero value otherwise.

### GetTcontsInUseOk

`func (o *TcontDeleteResultDTO) GetTcontsInUseOk() (*[]int32, bool)`

GetTcontsInUseOk returns a tuple with the TcontsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcontsInUse

`func (o *TcontDeleteResultDTO) SetTcontsInUse(v []int32)`

SetTcontsInUse sets TcontsInUse field to given value.

### HasTcontsInUse

`func (o *TcontDeleteResultDTO) HasTcontsInUse() bool`

HasTcontsInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


