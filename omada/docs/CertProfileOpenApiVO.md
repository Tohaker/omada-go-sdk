# CertProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCertFileId** | Pointer to **string** | CA Certificate file id, it&#39;s obtained from the interface &#39;uploadCaCertFile&#39;; When type &#x3D; 0, Parameter [caCertFileId] should not be null. | [optional] 
**CertFileMd5** | Pointer to **string** | Client certificate file MD5. | [optional] 
**ClientCertFileId** | Pointer to **string** | Client certificate file id, it&#39;s obtained from the interface &#39;uploadClientCertFile&#39;; When type &#x3D; 1, Parameter [clientCertFileId] should not be null. | [optional] 
**ExpiredOn** | Pointer to **int64** | Cert profile expired time, expired timestamp, in seconds, such as 1682000000. | [optional] 
**FileName** | Pointer to **string** | Certificate authority profile file name | [optional] 
**Format** | Pointer to **int32** | Cert profile format, format should be a value as follows: 0: X509; 1: DER. | [optional] 
**Id** | Pointer to **string** | Certificate authority profile entry id. | [optional] 
**Name** | Pointer to **string** | Certificate authority profile name | [optional] 
**PrivateKeyFileId** | Pointer to **string** | Client private key file id, it&#39;s obtained from the interface &#39;UploadClientPrivateKeyFile&#39;; When type &#x3D; 1, Parameter [privateKeyFileId] should not be null. | [optional] 
**PrivateKeyFileMd5** | Pointer to **string** | Client private key file MD5. | [optional] 
**PrivateKeyFileName** | Pointer to **string** | Client private key file name. | [optional] 
**PrivateKeyPassword** | Pointer to **string** | Certificate authority profile privateKey Password should contain 1 to 64 characters, spaces, comma, single quotation marks and double quotation marks are not allowed. | [optional] 
**Resource** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** | Cert profile status, status should be a value as follows: 0: Normal 1: Expired Soon 2: Expired. | [optional] 
**Type** | Pointer to **int32** | Certificate authority profile type, such as:  0: CA Cert; 1: Client Cert. | [optional] 

## Methods

### NewCertProfileOpenApiVO

`func NewCertProfileOpenApiVO() *CertProfileOpenApiVO`

NewCertProfileOpenApiVO instantiates a new CertProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertProfileOpenApiVOWithDefaults

`func NewCertProfileOpenApiVOWithDefaults() *CertProfileOpenApiVO`

NewCertProfileOpenApiVOWithDefaults instantiates a new CertProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCertFileId

`func (o *CertProfileOpenApiVO) GetCaCertFileId() string`

GetCaCertFileId returns the CaCertFileId field if non-nil, zero value otherwise.

### GetCaCertFileIdOk

`func (o *CertProfileOpenApiVO) GetCaCertFileIdOk() (*string, bool)`

GetCaCertFileIdOk returns a tuple with the CaCertFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertFileId

`func (o *CertProfileOpenApiVO) SetCaCertFileId(v string)`

SetCaCertFileId sets CaCertFileId field to given value.

### HasCaCertFileId

`func (o *CertProfileOpenApiVO) HasCaCertFileId() bool`

HasCaCertFileId returns a boolean if a field has been set.

### GetCertFileMd5

`func (o *CertProfileOpenApiVO) GetCertFileMd5() string`

GetCertFileMd5 returns the CertFileMd5 field if non-nil, zero value otherwise.

### GetCertFileMd5Ok

`func (o *CertProfileOpenApiVO) GetCertFileMd5Ok() (*string, bool)`

GetCertFileMd5Ok returns a tuple with the CertFileMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFileMd5

`func (o *CertProfileOpenApiVO) SetCertFileMd5(v string)`

SetCertFileMd5 sets CertFileMd5 field to given value.

### HasCertFileMd5

`func (o *CertProfileOpenApiVO) HasCertFileMd5() bool`

HasCertFileMd5 returns a boolean if a field has been set.

### GetClientCertFileId

`func (o *CertProfileOpenApiVO) GetClientCertFileId() string`

GetClientCertFileId returns the ClientCertFileId field if non-nil, zero value otherwise.

### GetClientCertFileIdOk

`func (o *CertProfileOpenApiVO) GetClientCertFileIdOk() (*string, bool)`

GetClientCertFileIdOk returns a tuple with the ClientCertFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertFileId

`func (o *CertProfileOpenApiVO) SetClientCertFileId(v string)`

SetClientCertFileId sets ClientCertFileId field to given value.

### HasClientCertFileId

`func (o *CertProfileOpenApiVO) HasClientCertFileId() bool`

HasClientCertFileId returns a boolean if a field has been set.

### GetExpiredOn

`func (o *CertProfileOpenApiVO) GetExpiredOn() int64`

GetExpiredOn returns the ExpiredOn field if non-nil, zero value otherwise.

### GetExpiredOnOk

`func (o *CertProfileOpenApiVO) GetExpiredOnOk() (*int64, bool)`

GetExpiredOnOk returns a tuple with the ExpiredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredOn

`func (o *CertProfileOpenApiVO) SetExpiredOn(v int64)`

SetExpiredOn sets ExpiredOn field to given value.

### HasExpiredOn

`func (o *CertProfileOpenApiVO) HasExpiredOn() bool`

HasExpiredOn returns a boolean if a field has been set.

### GetFileName

`func (o *CertProfileOpenApiVO) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CertProfileOpenApiVO) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CertProfileOpenApiVO) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CertProfileOpenApiVO) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFormat

`func (o *CertProfileOpenApiVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CertProfileOpenApiVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CertProfileOpenApiVO) SetFormat(v int32)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CertProfileOpenApiVO) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetId

`func (o *CertProfileOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertProfileOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertProfileOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertProfileOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CertProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertProfileOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivateKeyFileId

`func (o *CertProfileOpenApiVO) GetPrivateKeyFileId() string`

GetPrivateKeyFileId returns the PrivateKeyFileId field if non-nil, zero value otherwise.

### GetPrivateKeyFileIdOk

`func (o *CertProfileOpenApiVO) GetPrivateKeyFileIdOk() (*string, bool)`

GetPrivateKeyFileIdOk returns a tuple with the PrivateKeyFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFileId

`func (o *CertProfileOpenApiVO) SetPrivateKeyFileId(v string)`

SetPrivateKeyFileId sets PrivateKeyFileId field to given value.

### HasPrivateKeyFileId

`func (o *CertProfileOpenApiVO) HasPrivateKeyFileId() bool`

HasPrivateKeyFileId returns a boolean if a field has been set.

### GetPrivateKeyFileMd5

`func (o *CertProfileOpenApiVO) GetPrivateKeyFileMd5() string`

GetPrivateKeyFileMd5 returns the PrivateKeyFileMd5 field if non-nil, zero value otherwise.

### GetPrivateKeyFileMd5Ok

`func (o *CertProfileOpenApiVO) GetPrivateKeyFileMd5Ok() (*string, bool)`

GetPrivateKeyFileMd5Ok returns a tuple with the PrivateKeyFileMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFileMd5

`func (o *CertProfileOpenApiVO) SetPrivateKeyFileMd5(v string)`

SetPrivateKeyFileMd5 sets PrivateKeyFileMd5 field to given value.

### HasPrivateKeyFileMd5

`func (o *CertProfileOpenApiVO) HasPrivateKeyFileMd5() bool`

HasPrivateKeyFileMd5 returns a boolean if a field has been set.

### GetPrivateKeyFileName

`func (o *CertProfileOpenApiVO) GetPrivateKeyFileName() string`

GetPrivateKeyFileName returns the PrivateKeyFileName field if non-nil, zero value otherwise.

### GetPrivateKeyFileNameOk

`func (o *CertProfileOpenApiVO) GetPrivateKeyFileNameOk() (*string, bool)`

GetPrivateKeyFileNameOk returns a tuple with the PrivateKeyFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFileName

`func (o *CertProfileOpenApiVO) SetPrivateKeyFileName(v string)`

SetPrivateKeyFileName sets PrivateKeyFileName field to given value.

### HasPrivateKeyFileName

`func (o *CertProfileOpenApiVO) HasPrivateKeyFileName() bool`

HasPrivateKeyFileName returns a boolean if a field has been set.

### GetPrivateKeyPassword

`func (o *CertProfileOpenApiVO) GetPrivateKeyPassword() string`

GetPrivateKeyPassword returns the PrivateKeyPassword field if non-nil, zero value otherwise.

### GetPrivateKeyPasswordOk

`func (o *CertProfileOpenApiVO) GetPrivateKeyPasswordOk() (*string, bool)`

GetPrivateKeyPasswordOk returns a tuple with the PrivateKeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassword

`func (o *CertProfileOpenApiVO) SetPrivateKeyPassword(v string)`

SetPrivateKeyPassword sets PrivateKeyPassword field to given value.

### HasPrivateKeyPassword

`func (o *CertProfileOpenApiVO) HasPrivateKeyPassword() bool`

HasPrivateKeyPassword returns a boolean if a field has been set.

### GetResource

`func (o *CertProfileOpenApiVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CertProfileOpenApiVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CertProfileOpenApiVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *CertProfileOpenApiVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetStatus

`func (o *CertProfileOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertProfileOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertProfileOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertProfileOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *CertProfileOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertProfileOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertProfileOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CertProfileOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


