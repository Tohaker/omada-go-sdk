# UrlFilterGlobalOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockPage** | **bool** | blockPage should be a value as follows: false: close; true: open. | 
**BlockPageMessage** | Pointer to **string** | blockPageMessage of blockPage,when blockPage switch is open. | [optional] 
**SafeSearch** | **bool** | safeSearch should be a value as follows: false: close; true: open. | 

## Methods

### NewUrlFilterGlobalOpenApiVO

`func NewUrlFilterGlobalOpenApiVO(blockPage bool, safeSearch bool, ) *UrlFilterGlobalOpenApiVO`

NewUrlFilterGlobalOpenApiVO instantiates a new UrlFilterGlobalOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlFilterGlobalOpenApiVOWithDefaults

`func NewUrlFilterGlobalOpenApiVOWithDefaults() *UrlFilterGlobalOpenApiVO`

NewUrlFilterGlobalOpenApiVOWithDefaults instantiates a new UrlFilterGlobalOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockPage

`func (o *UrlFilterGlobalOpenApiVO) GetBlockPage() bool`

GetBlockPage returns the BlockPage field if non-nil, zero value otherwise.

### GetBlockPageOk

`func (o *UrlFilterGlobalOpenApiVO) GetBlockPageOk() (*bool, bool)`

GetBlockPageOk returns a tuple with the BlockPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockPage

`func (o *UrlFilterGlobalOpenApiVO) SetBlockPage(v bool)`

SetBlockPage sets BlockPage field to given value.


### GetBlockPageMessage

`func (o *UrlFilterGlobalOpenApiVO) GetBlockPageMessage() string`

GetBlockPageMessage returns the BlockPageMessage field if non-nil, zero value otherwise.

### GetBlockPageMessageOk

`func (o *UrlFilterGlobalOpenApiVO) GetBlockPageMessageOk() (*string, bool)`

GetBlockPageMessageOk returns a tuple with the BlockPageMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockPageMessage

`func (o *UrlFilterGlobalOpenApiVO) SetBlockPageMessage(v string)`

SetBlockPageMessage sets BlockPageMessage field to given value.

### HasBlockPageMessage

`func (o *UrlFilterGlobalOpenApiVO) HasBlockPageMessage() bool`

HasBlockPageMessage returns a boolean if a field has been set.

### GetSafeSearch

`func (o *UrlFilterGlobalOpenApiVO) GetSafeSearch() bool`

GetSafeSearch returns the SafeSearch field if non-nil, zero value otherwise.

### GetSafeSearchOk

`func (o *UrlFilterGlobalOpenApiVO) GetSafeSearchOk() (*bool, bool)`

GetSafeSearchOk returns a tuple with the SafeSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeSearch

`func (o *UrlFilterGlobalOpenApiVO) SetSafeSearch(v bool)`

SetSafeSearch sets SafeSearch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


