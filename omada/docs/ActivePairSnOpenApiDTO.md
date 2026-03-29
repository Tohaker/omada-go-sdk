# ActivePairSnOpenApiDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseId** | Pointer to **string** | License ID | [optional] 
**LicenseType** | **string** | License type should be a value as follows: 1year; 2years; 3years; 4years; 5years; others; trial | 
**Sn** | **string** | Device serial number. It should contains 13 characters. | 

## Methods

### NewActivePairSnOpenApiDTO

`func NewActivePairSnOpenApiDTO(licenseType string, sn string, ) *ActivePairSnOpenApiDTO`

NewActivePairSnOpenApiDTO instantiates a new ActivePairSnOpenApiDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivePairSnOpenApiDTOWithDefaults

`func NewActivePairSnOpenApiDTOWithDefaults() *ActivePairSnOpenApiDTO`

NewActivePairSnOpenApiDTOWithDefaults instantiates a new ActivePairSnOpenApiDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseId

`func (o *ActivePairSnOpenApiDTO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *ActivePairSnOpenApiDTO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *ActivePairSnOpenApiDTO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *ActivePairSnOpenApiDTO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseType

`func (o *ActivePairSnOpenApiDTO) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *ActivePairSnOpenApiDTO) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *ActivePairSnOpenApiDTO) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.


### GetSn

`func (o *ActivePairSnOpenApiDTO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *ActivePairSnOpenApiDTO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *ActivePairSnOpenApiDTO) SetSn(v string)`

SetSn sets Sn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


