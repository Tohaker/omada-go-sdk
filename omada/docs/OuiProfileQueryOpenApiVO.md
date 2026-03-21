# OuiProfileQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | OUI Profile count | [optional] 
**Id** | Pointer to **string** | OUI Profile ID | [optional] 
**Name** | **string** | OUI Profile name should contain 1 to 64 characters. | 
**OuiCombine** | Pointer to [**[]OUIAndDescription**](OUIAndDescription.md) | OUI and description | [optional] 

## Methods

### NewOuiProfileQueryOpenApiVO

`func NewOuiProfileQueryOpenApiVO(name string, ) *OuiProfileQueryOpenApiVO`

NewOuiProfileQueryOpenApiVO instantiates a new OuiProfileQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiProfileQueryOpenApiVOWithDefaults

`func NewOuiProfileQueryOpenApiVOWithDefaults() *OuiProfileQueryOpenApiVO`

NewOuiProfileQueryOpenApiVOWithDefaults instantiates a new OuiProfileQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *OuiProfileQueryOpenApiVO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OuiProfileQueryOpenApiVO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OuiProfileQueryOpenApiVO) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *OuiProfileQueryOpenApiVO) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetId

`func (o *OuiProfileQueryOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OuiProfileQueryOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OuiProfileQueryOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OuiProfileQueryOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OuiProfileQueryOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OuiProfileQueryOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OuiProfileQueryOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetOuiCombine

`func (o *OuiProfileQueryOpenApiVO) GetOuiCombine() []OUIAndDescription`

GetOuiCombine returns the OuiCombine field if non-nil, zero value otherwise.

### GetOuiCombineOk

`func (o *OuiProfileQueryOpenApiVO) GetOuiCombineOk() (*[]OUIAndDescription, bool)`

GetOuiCombineOk returns a tuple with the OuiCombine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuiCombine

`func (o *OuiProfileQueryOpenApiVO) SetOuiCombine(v []OUIAndDescription)`

SetOuiCombine sets OuiCombine field to given value.

### HasOuiCombine

`func (o *OuiProfileQueryOpenApiVO) HasOuiCombine() bool`

HasOuiCombine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


