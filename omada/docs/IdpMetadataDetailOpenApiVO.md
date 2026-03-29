# IdpMetadataDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description. | [optional] 
**EntityId** | Pointer to **string** | The IdP entity ID which must be unique in same Omadac. | [optional] 
**EntityUrl** | Pointer to **string** | Entity url. | [optional] 
**IdpId** | Pointer to **string** | IdP ID(resource ID). | [optional] 
**LoginUrl** | Pointer to **string** | Login url. | [optional] 
**Name** | Pointer to **string** | IdP name. | [optional] 
**SignOnUrl** | Pointer to **string** | Sign On url. | [optional] 
**X509Certificate** | Pointer to **string** | BASE64 encoded string of x509 certificate. | [optional] 

## Methods

### NewIdpMetadataDetailOpenApiVO

`func NewIdpMetadataDetailOpenApiVO() *IdpMetadataDetailOpenApiVO`

NewIdpMetadataDetailOpenApiVO instantiates a new IdpMetadataDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpMetadataDetailOpenApiVOWithDefaults

`func NewIdpMetadataDetailOpenApiVOWithDefaults() *IdpMetadataDetailOpenApiVO`

NewIdpMetadataDetailOpenApiVOWithDefaults instantiates a new IdpMetadataDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IdpMetadataDetailOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdpMetadataDetailOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdpMetadataDetailOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdpMetadataDetailOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityId

`func (o *IdpMetadataDetailOpenApiVO) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IdpMetadataDetailOpenApiVO) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IdpMetadataDetailOpenApiVO) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *IdpMetadataDetailOpenApiVO) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityUrl

`func (o *IdpMetadataDetailOpenApiVO) GetEntityUrl() string`

GetEntityUrl returns the EntityUrl field if non-nil, zero value otherwise.

### GetEntityUrlOk

`func (o *IdpMetadataDetailOpenApiVO) GetEntityUrlOk() (*string, bool)`

GetEntityUrlOk returns a tuple with the EntityUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityUrl

`func (o *IdpMetadataDetailOpenApiVO) SetEntityUrl(v string)`

SetEntityUrl sets EntityUrl field to given value.

### HasEntityUrl

`func (o *IdpMetadataDetailOpenApiVO) HasEntityUrl() bool`

HasEntityUrl returns a boolean if a field has been set.

### GetIdpId

`func (o *IdpMetadataDetailOpenApiVO) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *IdpMetadataDetailOpenApiVO) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *IdpMetadataDetailOpenApiVO) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.

### HasIdpId

`func (o *IdpMetadataDetailOpenApiVO) HasIdpId() bool`

HasIdpId returns a boolean if a field has been set.

### GetLoginUrl

`func (o *IdpMetadataDetailOpenApiVO) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *IdpMetadataDetailOpenApiVO) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *IdpMetadataDetailOpenApiVO) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *IdpMetadataDetailOpenApiVO) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetName

`func (o *IdpMetadataDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpMetadataDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpMetadataDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdpMetadataDetailOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSignOnUrl

`func (o *IdpMetadataDetailOpenApiVO) GetSignOnUrl() string`

GetSignOnUrl returns the SignOnUrl field if non-nil, zero value otherwise.

### GetSignOnUrlOk

`func (o *IdpMetadataDetailOpenApiVO) GetSignOnUrlOk() (*string, bool)`

GetSignOnUrlOk returns a tuple with the SignOnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnUrl

`func (o *IdpMetadataDetailOpenApiVO) SetSignOnUrl(v string)`

SetSignOnUrl sets SignOnUrl field to given value.

### HasSignOnUrl

`func (o *IdpMetadataDetailOpenApiVO) HasSignOnUrl() bool`

HasSignOnUrl returns a boolean if a field has been set.

### GetX509Certificate

`func (o *IdpMetadataDetailOpenApiVO) GetX509Certificate() string`

GetX509Certificate returns the X509Certificate field if non-nil, zero value otherwise.

### GetX509CertificateOk

`func (o *IdpMetadataDetailOpenApiVO) GetX509CertificateOk() (*string, bool)`

GetX509CertificateOk returns a tuple with the X509Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Certificate

`func (o *IdpMetadataDetailOpenApiVO) SetX509Certificate(v string)`

SetX509Certificate sets X509Certificate field to given value.

### HasX509Certificate

`func (o *IdpMetadataDetailOpenApiVO) HasX509Certificate() bool`

HasX509Certificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


