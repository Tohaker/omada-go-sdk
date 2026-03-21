# IdpMetadataOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | IdP description should contain 0 to 128 characters. | [optional] 
**EntityId** | **string** | The IdP entity id which must be unique in same Omadac. | 
**LoginUrl** | **string** | Login url | 
**Name** | **string** | IdP name should contain 1 to 32 characters. | 
**X509Certificate** | **string** | BASE64 encoded string of x509 certificate. | 

## Methods

### NewIdpMetadataOpenApiVO

`func NewIdpMetadataOpenApiVO(entityId string, loginUrl string, name string, x509Certificate string, ) *IdpMetadataOpenApiVO`

NewIdpMetadataOpenApiVO instantiates a new IdpMetadataOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpMetadataOpenApiVOWithDefaults

`func NewIdpMetadataOpenApiVOWithDefaults() *IdpMetadataOpenApiVO`

NewIdpMetadataOpenApiVOWithDefaults instantiates a new IdpMetadataOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IdpMetadataOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdpMetadataOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdpMetadataOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdpMetadataOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityId

`func (o *IdpMetadataOpenApiVO) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IdpMetadataOpenApiVO) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IdpMetadataOpenApiVO) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetLoginUrl

`func (o *IdpMetadataOpenApiVO) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *IdpMetadataOpenApiVO) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *IdpMetadataOpenApiVO) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.


### GetName

`func (o *IdpMetadataOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpMetadataOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpMetadataOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetX509Certificate

`func (o *IdpMetadataOpenApiVO) GetX509Certificate() string`

GetX509Certificate returns the X509Certificate field if non-nil, zero value otherwise.

### GetX509CertificateOk

`func (o *IdpMetadataOpenApiVO) GetX509CertificateOk() (*string, bool)`

GetX509CertificateOk returns a tuple with the X509Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Certificate

`func (o *IdpMetadataOpenApiVO) SetX509Certificate(v string)`

SetX509Certificate sets X509Certificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


