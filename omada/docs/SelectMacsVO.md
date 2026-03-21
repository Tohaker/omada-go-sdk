# SelectMacsVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Macs** | Pointer to **[]string** | MAC list. When parameter [selectType] is &#39;all&#39;, it should be null.When parameter [selectType] is &#39;include&#39;, it means the MAC list that needs to be included.When parameter [selectType] is &#39;exclude&#39;, it means the MAC list that needs to be excluded. | [optional] 
**SelectType** | Pointer to **string** | The value of parameter [selectType] must be in [all, include, exclude]. | [optional] 

## Methods

### NewSelectMacsVO

`func NewSelectMacsVO() *SelectMacsVO`

NewSelectMacsVO instantiates a new SelectMacsVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectMacsVOWithDefaults

`func NewSelectMacsVOWithDefaults() *SelectMacsVO`

NewSelectMacsVOWithDefaults instantiates a new SelectMacsVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacs

`func (o *SelectMacsVO) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *SelectMacsVO) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *SelectMacsVO) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *SelectMacsVO) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetSelectType

`func (o *SelectMacsVO) GetSelectType() string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *SelectMacsVO) GetSelectTypeOk() (*string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *SelectMacsVO) SetSelectType(v string)`

SetSelectType sets SelectType field to given value.

### HasSelectType

`func (o *SelectMacsVO) HasSelectType() bool`

HasSelectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


