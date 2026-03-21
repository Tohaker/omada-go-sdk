# LicenseRecycleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | Customer ID | 
**LicenseNum** | [**[]LicenseNumOpenApiVO**](LicenseNumOpenApiVO.md) | License num | 
**Others** | [**LicenseOthersIdsOpenApiVO**](LicenseOthersIdsOpenApiVO.md) |  | 

## Methods

### NewLicenseRecycleOpenApiVO

`func NewLicenseRecycleOpenApiVO(customerId string, licenseNum []LicenseNumOpenApiVO, others LicenseOthersIdsOpenApiVO, ) *LicenseRecycleOpenApiVO`

NewLicenseRecycleOpenApiVO instantiates a new LicenseRecycleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseRecycleOpenApiVOWithDefaults

`func NewLicenseRecycleOpenApiVOWithDefaults() *LicenseRecycleOpenApiVO`

NewLicenseRecycleOpenApiVOWithDefaults instantiates a new LicenseRecycleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *LicenseRecycleOpenApiVO) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *LicenseRecycleOpenApiVO) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *LicenseRecycleOpenApiVO) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetLicenseNum

`func (o *LicenseRecycleOpenApiVO) GetLicenseNum() []LicenseNumOpenApiVO`

GetLicenseNum returns the LicenseNum field if non-nil, zero value otherwise.

### GetLicenseNumOk

`func (o *LicenseRecycleOpenApiVO) GetLicenseNumOk() (*[]LicenseNumOpenApiVO, bool)`

GetLicenseNumOk returns a tuple with the LicenseNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNum

`func (o *LicenseRecycleOpenApiVO) SetLicenseNum(v []LicenseNumOpenApiVO)`

SetLicenseNum sets LicenseNum field to given value.


### GetOthers

`func (o *LicenseRecycleOpenApiVO) GetOthers() LicenseOthersIdsOpenApiVO`

GetOthers returns the Others field if non-nil, zero value otherwise.

### GetOthersOk

`func (o *LicenseRecycleOpenApiVO) GetOthersOk() (*LicenseOthersIdsOpenApiVO, bool)`

GetOthersOk returns a tuple with the Others field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthers

`func (o *LicenseRecycleOpenApiVO) SetOthers(v LicenseOthersIdsOpenApiVO)`

SetOthers sets Others field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


