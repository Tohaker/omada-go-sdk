# CustomerCopyOpenapiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Customer description should contain 1 to 128 characters. | [optional] 
**NewCustomerName** | **string** | New Customer name should contain 1 to 31 characters. | 
**SourceCustomerId** | **string** | Source customer ID to be copied. | 

## Methods

### NewCustomerCopyOpenapiVO

`func NewCustomerCopyOpenapiVO(newCustomerName string, sourceCustomerId string, ) *CustomerCopyOpenapiVO`

NewCustomerCopyOpenapiVO instantiates a new CustomerCopyOpenapiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerCopyOpenapiVOWithDefaults

`func NewCustomerCopyOpenapiVOWithDefaults() *CustomerCopyOpenapiVO`

NewCustomerCopyOpenapiVOWithDefaults instantiates a new CustomerCopyOpenapiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CustomerCopyOpenapiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomerCopyOpenapiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomerCopyOpenapiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomerCopyOpenapiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNewCustomerName

`func (o *CustomerCopyOpenapiVO) GetNewCustomerName() string`

GetNewCustomerName returns the NewCustomerName field if non-nil, zero value otherwise.

### GetNewCustomerNameOk

`func (o *CustomerCopyOpenapiVO) GetNewCustomerNameOk() (*string, bool)`

GetNewCustomerNameOk returns a tuple with the NewCustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCustomerName

`func (o *CustomerCopyOpenapiVO) SetNewCustomerName(v string)`

SetNewCustomerName sets NewCustomerName field to given value.


### GetSourceCustomerId

`func (o *CustomerCopyOpenapiVO) GetSourceCustomerId() string`

GetSourceCustomerId returns the SourceCustomerId field if non-nil, zero value otherwise.

### GetSourceCustomerIdOk

`func (o *CustomerCopyOpenapiVO) GetSourceCustomerIdOk() (*string, bool)`

GetSourceCustomerIdOk returns a tuple with the SourceCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCustomerId

`func (o *CustomerCopyOpenapiVO) SetSourceCustomerId(v string)`

SetSourceCustomerId sets SourceCustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


