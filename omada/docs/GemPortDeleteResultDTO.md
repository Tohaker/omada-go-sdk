# GemPortDeleteResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GemPortsInUse** | Pointer to **[]int32** | The list of Gem Port ID that failed to delete due to being in use. | [optional] 

## Methods

### NewGemPortDeleteResultDTO

`func NewGemPortDeleteResultDTO() *GemPortDeleteResultDTO`

NewGemPortDeleteResultDTO instantiates a new GemPortDeleteResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGemPortDeleteResultDTOWithDefaults

`func NewGemPortDeleteResultDTOWithDefaults() *GemPortDeleteResultDTO`

NewGemPortDeleteResultDTOWithDefaults instantiates a new GemPortDeleteResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGemPortsInUse

`func (o *GemPortDeleteResultDTO) GetGemPortsInUse() []int32`

GetGemPortsInUse returns the GemPortsInUse field if non-nil, zero value otherwise.

### GetGemPortsInUseOk

`func (o *GemPortDeleteResultDTO) GetGemPortsInUseOk() (*[]int32, bool)`

GetGemPortsInUseOk returns a tuple with the GemPortsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemPortsInUse

`func (o *GemPortDeleteResultDTO) SetGemPortsInUse(v []int32)`

SetGemPortsInUse sets GemPortsInUse field to given value.

### HasGemPortsInUse

`func (o *GemPortDeleteResultDTO) HasGemPortsInUse() bool`

HasGemPortsInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


