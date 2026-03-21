# CertProfileRequestOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCertFileId** | Pointer to **string** | CA Certificate file id, it&#39;s obtained from the interface &#39;uploadCaCertFile&#39;; When type &#x3D; 0, Parameter [caCertFileId] should not be null. | [optional] 
**ClientCertFileId** | Pointer to **string** | Client certificate file id, it&#39;s obtained from the interface &#39;uploadClientCertFile&#39;; When type &#x3D; 1, Parameter [clientCertFileId] should not be null. | [optional] 
**Format** | **int32** | Cert profile format, format should be a value as follows: 0: X509; 1: DER. | 
**Name** | **string** | Cert profile name, name should contain 1 to 64 characters. | 
**PrivateKeyFileId** | Pointer to **string** | Client private key file id, it&#39;s obtained from the interface &#39;UploadClientPrivateKeyFile&#39;; When type &#x3D; 1, Parameter [privateKeyFileId] should not be null. | [optional] 
**PrivateKeyPassword** | Pointer to **string** | Client private key password; When type &#x3D; 1, Parameter [privateKeyPassword] is an optional input. | [optional] 
**Type** | **int32** | Cert profile type, type should be a value as follows: 0: CA Cert; 1: Client Cert.. | 

## Methods

### NewCertProfileRequestOpenApiVO

`func NewCertProfileRequestOpenApiVO(format int32, name string, type_ int32, ) *CertProfileRequestOpenApiVO`

NewCertProfileRequestOpenApiVO instantiates a new CertProfileRequestOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertProfileRequestOpenApiVOWithDefaults

`func NewCertProfileRequestOpenApiVOWithDefaults() *CertProfileRequestOpenApiVO`

NewCertProfileRequestOpenApiVOWithDefaults instantiates a new CertProfileRequestOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCertFileId

`func (o *CertProfileRequestOpenApiVO) GetCaCertFileId() string`

GetCaCertFileId returns the CaCertFileId field if non-nil, zero value otherwise.

### GetCaCertFileIdOk

`func (o *CertProfileRequestOpenApiVO) GetCaCertFileIdOk() (*string, bool)`

GetCaCertFileIdOk returns a tuple with the CaCertFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertFileId

`func (o *CertProfileRequestOpenApiVO) SetCaCertFileId(v string)`

SetCaCertFileId sets CaCertFileId field to given value.

### HasCaCertFileId

`func (o *CertProfileRequestOpenApiVO) HasCaCertFileId() bool`

HasCaCertFileId returns a boolean if a field has been set.

### GetClientCertFileId

`func (o *CertProfileRequestOpenApiVO) GetClientCertFileId() string`

GetClientCertFileId returns the ClientCertFileId field if non-nil, zero value otherwise.

### GetClientCertFileIdOk

`func (o *CertProfileRequestOpenApiVO) GetClientCertFileIdOk() (*string, bool)`

GetClientCertFileIdOk returns a tuple with the ClientCertFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertFileId

`func (o *CertProfileRequestOpenApiVO) SetClientCertFileId(v string)`

SetClientCertFileId sets ClientCertFileId field to given value.

### HasClientCertFileId

`func (o *CertProfileRequestOpenApiVO) HasClientCertFileId() bool`

HasClientCertFileId returns a boolean if a field has been set.

### GetFormat

`func (o *CertProfileRequestOpenApiVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CertProfileRequestOpenApiVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CertProfileRequestOpenApiVO) SetFormat(v int32)`

SetFormat sets Format field to given value.


### GetName

`func (o *CertProfileRequestOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertProfileRequestOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertProfileRequestOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPrivateKeyFileId

`func (o *CertProfileRequestOpenApiVO) GetPrivateKeyFileId() string`

GetPrivateKeyFileId returns the PrivateKeyFileId field if non-nil, zero value otherwise.

### GetPrivateKeyFileIdOk

`func (o *CertProfileRequestOpenApiVO) GetPrivateKeyFileIdOk() (*string, bool)`

GetPrivateKeyFileIdOk returns a tuple with the PrivateKeyFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFileId

`func (o *CertProfileRequestOpenApiVO) SetPrivateKeyFileId(v string)`

SetPrivateKeyFileId sets PrivateKeyFileId field to given value.

### HasPrivateKeyFileId

`func (o *CertProfileRequestOpenApiVO) HasPrivateKeyFileId() bool`

HasPrivateKeyFileId returns a boolean if a field has been set.

### GetPrivateKeyPassword

`func (o *CertProfileRequestOpenApiVO) GetPrivateKeyPassword() string`

GetPrivateKeyPassword returns the PrivateKeyPassword field if non-nil, zero value otherwise.

### GetPrivateKeyPasswordOk

`func (o *CertProfileRequestOpenApiVO) GetPrivateKeyPasswordOk() (*string, bool)`

GetPrivateKeyPasswordOk returns a tuple with the PrivateKeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassword

`func (o *CertProfileRequestOpenApiVO) SetPrivateKeyPassword(v string)`

SetPrivateKeyPassword sets PrivateKeyPassword field to given value.

### HasPrivateKeyPassword

`func (o *CertProfileRequestOpenApiVO) HasPrivateKeyPassword() bool`

HasPrivateKeyPassword returns a boolean if a field has been set.

### GetType

`func (o *CertProfileRequestOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertProfileRequestOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertProfileRequestOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


