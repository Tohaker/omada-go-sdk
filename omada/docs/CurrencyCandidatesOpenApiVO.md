# CurrencyCandidatesOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyList** | Pointer to **[]string** | All currency Short Code list. For the values of Currency Short Code, refer to section 5.4.2 of the Open API Access Guide. | [optional] 
**SelectedCurrency** | Pointer to **string** | The currency selected for the site | [optional] 

## Methods

### NewCurrencyCandidatesOpenApiVO

`func NewCurrencyCandidatesOpenApiVO() *CurrencyCandidatesOpenApiVO`

NewCurrencyCandidatesOpenApiVO instantiates a new CurrencyCandidatesOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyCandidatesOpenApiVOWithDefaults

`func NewCurrencyCandidatesOpenApiVOWithDefaults() *CurrencyCandidatesOpenApiVO`

NewCurrencyCandidatesOpenApiVOWithDefaults instantiates a new CurrencyCandidatesOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyList

`func (o *CurrencyCandidatesOpenApiVO) GetCurrencyList() []string`

GetCurrencyList returns the CurrencyList field if non-nil, zero value otherwise.

### GetCurrencyListOk

`func (o *CurrencyCandidatesOpenApiVO) GetCurrencyListOk() (*[]string, bool)`

GetCurrencyListOk returns a tuple with the CurrencyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyList

`func (o *CurrencyCandidatesOpenApiVO) SetCurrencyList(v []string)`

SetCurrencyList sets CurrencyList field to given value.

### HasCurrencyList

`func (o *CurrencyCandidatesOpenApiVO) HasCurrencyList() bool`

HasCurrencyList returns a boolean if a field has been set.

### GetSelectedCurrency

`func (o *CurrencyCandidatesOpenApiVO) GetSelectedCurrency() string`

GetSelectedCurrency returns the SelectedCurrency field if non-nil, zero value otherwise.

### GetSelectedCurrencyOk

`func (o *CurrencyCandidatesOpenApiVO) GetSelectedCurrencyOk() (*string, bool)`

GetSelectedCurrencyOk returns a tuple with the SelectedCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedCurrency

`func (o *CurrencyCandidatesOpenApiVO) SetSelectedCurrency(v string)`

SetSelectedCurrency sets SelectedCurrency field to given value.

### HasSelectedCurrency

`func (o *CurrencyCandidatesOpenApiVO) HasSelectedCurrency() bool`

HasSelectedCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


