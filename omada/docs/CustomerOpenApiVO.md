# CustomerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerName** | Pointer to **string** | Customer name should contain 1 to 31 characters. | [optional] 
**Description** | Pointer to **string** | Customer description should contain 1 to 128 characters. | [optional] 

## Methods

### NewCustomerOpenApiVO

`func NewCustomerOpenApiVO() *CustomerOpenApiVO`

NewCustomerOpenApiVO instantiates a new CustomerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerOpenApiVOWithDefaults

`func NewCustomerOpenApiVOWithDefaults() *CustomerOpenApiVO`

NewCustomerOpenApiVOWithDefaults instantiates a new CustomerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerName

`func (o *CustomerOpenApiVO) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *CustomerOpenApiVO) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *CustomerOpenApiVO) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *CustomerOpenApiVO) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetDescription

`func (o *CustomerOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomerOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomerOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomerOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


