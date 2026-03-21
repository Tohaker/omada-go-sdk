# ActivePairOpenApiDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**LicenseId** | Pointer to **string** | License ID | [optional] 
**LicenseType** | **string** | License type should be a value as follows: Cloud Based Controller(1year, 2years, 3years, 4years, 5years, others, trial); Local Controller(trial, permanent); | 

## Methods

### NewActivePairOpenApiDTO

`func NewActivePairOpenApiDTO(deviceMac string, licenseType string, ) *ActivePairOpenApiDTO`

NewActivePairOpenApiDTO instantiates a new ActivePairOpenApiDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivePairOpenApiDTOWithDefaults

`func NewActivePairOpenApiDTOWithDefaults() *ActivePairOpenApiDTO`

NewActivePairOpenApiDTOWithDefaults instantiates a new ActivePairOpenApiDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *ActivePairOpenApiDTO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *ActivePairOpenApiDTO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *ActivePairOpenApiDTO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.


### GetLicenseId

`func (o *ActivePairOpenApiDTO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *ActivePairOpenApiDTO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *ActivePairOpenApiDTO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *ActivePairOpenApiDTO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseType

`func (o *ActivePairOpenApiDTO) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *ActivePairOpenApiDTO) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *ActivePairOpenApiDTO) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


