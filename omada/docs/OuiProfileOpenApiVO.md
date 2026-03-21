# OuiProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | OUI Profile name should contain 1 to 64 characters. | 
**OuiCombine** | Pointer to [**[]OUIAndDescription**](OUIAndDescription.md) | OUI and description | [optional] 

## Methods

### NewOuiProfileOpenApiVO

`func NewOuiProfileOpenApiVO(name string, ) *OuiProfileOpenApiVO`

NewOuiProfileOpenApiVO instantiates a new OuiProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiProfileOpenApiVOWithDefaults

`func NewOuiProfileOpenApiVOWithDefaults() *OuiProfileOpenApiVO`

NewOuiProfileOpenApiVOWithDefaults instantiates a new OuiProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OuiProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OuiProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OuiProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetOuiCombine

`func (o *OuiProfileOpenApiVO) GetOuiCombine() []OUIAndDescription`

GetOuiCombine returns the OuiCombine field if non-nil, zero value otherwise.

### GetOuiCombineOk

`func (o *OuiProfileOpenApiVO) GetOuiCombineOk() (*[]OUIAndDescription, bool)`

GetOuiCombineOk returns a tuple with the OuiCombine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuiCombine

`func (o *OuiProfileOpenApiVO) SetOuiCombine(v []OUIAndDescription)`

SetOuiCombine sets OuiCombine field to given value.

### HasOuiCombine

`func (o *OuiProfileOpenApiVO) HasOuiCombine() bool`

HasOuiCombine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


