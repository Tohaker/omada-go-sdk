# ActivePairMultiSiteOpenApiDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**LicenseId** | Pointer to **string** | License ID | [optional] 
**LicenseType** | **string** | License type should be a value as follows: Cloud Based Controller(1year, 2years, 3years, 4years, 5years, others, trial); Local Controller(trial, permanent); | 
**SiteId** | Pointer to **string** | Site ID | [optional] 

## Methods

### NewActivePairMultiSiteOpenApiDTO

`func NewActivePairMultiSiteOpenApiDTO(deviceMac string, licenseType string, ) *ActivePairMultiSiteOpenApiDTO`

NewActivePairMultiSiteOpenApiDTO instantiates a new ActivePairMultiSiteOpenApiDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivePairMultiSiteOpenApiDTOWithDefaults

`func NewActivePairMultiSiteOpenApiDTOWithDefaults() *ActivePairMultiSiteOpenApiDTO`

NewActivePairMultiSiteOpenApiDTOWithDefaults instantiates a new ActivePairMultiSiteOpenApiDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *ActivePairMultiSiteOpenApiDTO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *ActivePairMultiSiteOpenApiDTO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *ActivePairMultiSiteOpenApiDTO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.


### GetLicenseId

`func (o *ActivePairMultiSiteOpenApiDTO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *ActivePairMultiSiteOpenApiDTO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *ActivePairMultiSiteOpenApiDTO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *ActivePairMultiSiteOpenApiDTO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseType

`func (o *ActivePairMultiSiteOpenApiDTO) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *ActivePairMultiSiteOpenApiDTO) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *ActivePairMultiSiteOpenApiDTO) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.


### GetSiteId

`func (o *ActivePairMultiSiteOpenApiDTO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ActivePairMultiSiteOpenApiDTO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ActivePairMultiSiteOpenApiDTO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ActivePairMultiSiteOpenApiDTO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


