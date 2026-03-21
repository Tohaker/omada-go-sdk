# CertProfileDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertName** | Pointer to **string** | Certificate authority profile file name, such as : server_certificate | [optional] 
**CertType** | Pointer to **string** | Cert profile file type, such as: X.509. | [optional] 
**ExpiresOn** | Pointer to **string** | Certificate authority profile file expired on, such as : Mon Oct 30 16:54:19 CST 2034 | [optional] 
**FileName** | Pointer to **string** | Certificate authority profile file name | [optional] 
**IssuedOn** | Pointer to **string** | Certificate authority profile file issued on, such as : Fri Nov 01 16:54:19 CST 2024 | [optional] 
**Issuer** | Pointer to **string** | Certificate authority profile file issuer, such as : L&#x3D;$$$$, CN&#x3D;TLSGenSelfSignedRootCA 2024-11-01T16:54:19.242005 | [optional] 
**RsaKeySize** | Pointer to **string** | Certificate authority profile file rsa Key Size, such as : 3 | [optional] 
**SerialNumber** | Pointer to **string** | Cert profile serial number, such as 1. | [optional] 
**SignedUsing** | Pointer to **string** | Certificate authority profile file signed using such as : SHA256withRSA | [optional] 
**Subject** | Pointer to **string** | Certificate authority profile file subject, such as : O&#x3D;server, CN&#x3D;192.168.0.20 | [optional] 
**Type** | Pointer to **int32** | Certificate authority profile type, such as:  0: CA Cert; 1: Client Cert. | [optional] 
**Version** | Pointer to **int32** | Cert profile file version, such as: 3. | [optional] 

## Methods

### NewCertProfileDetailOpenApiVO

`func NewCertProfileDetailOpenApiVO() *CertProfileDetailOpenApiVO`

NewCertProfileDetailOpenApiVO instantiates a new CertProfileDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertProfileDetailOpenApiVOWithDefaults

`func NewCertProfileDetailOpenApiVOWithDefaults() *CertProfileDetailOpenApiVO`

NewCertProfileDetailOpenApiVOWithDefaults instantiates a new CertProfileDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertName

`func (o *CertProfileDetailOpenApiVO) GetCertName() string`

GetCertName returns the CertName field if non-nil, zero value otherwise.

### GetCertNameOk

`func (o *CertProfileDetailOpenApiVO) GetCertNameOk() (*string, bool)`

GetCertNameOk returns a tuple with the CertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertName

`func (o *CertProfileDetailOpenApiVO) SetCertName(v string)`

SetCertName sets CertName field to given value.

### HasCertName

`func (o *CertProfileDetailOpenApiVO) HasCertName() bool`

HasCertName returns a boolean if a field has been set.

### GetCertType

`func (o *CertProfileDetailOpenApiVO) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *CertProfileDetailOpenApiVO) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *CertProfileDetailOpenApiVO) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *CertProfileDetailOpenApiVO) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### GetExpiresOn

`func (o *CertProfileDetailOpenApiVO) GetExpiresOn() string`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *CertProfileDetailOpenApiVO) GetExpiresOnOk() (*string, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *CertProfileDetailOpenApiVO) SetExpiresOn(v string)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *CertProfileDetailOpenApiVO) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.

### GetFileName

`func (o *CertProfileDetailOpenApiVO) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CertProfileDetailOpenApiVO) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CertProfileDetailOpenApiVO) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CertProfileDetailOpenApiVO) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetIssuedOn

`func (o *CertProfileDetailOpenApiVO) GetIssuedOn() string`

GetIssuedOn returns the IssuedOn field if non-nil, zero value otherwise.

### GetIssuedOnOk

`func (o *CertProfileDetailOpenApiVO) GetIssuedOnOk() (*string, bool)`

GetIssuedOnOk returns a tuple with the IssuedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedOn

`func (o *CertProfileDetailOpenApiVO) SetIssuedOn(v string)`

SetIssuedOn sets IssuedOn field to given value.

### HasIssuedOn

`func (o *CertProfileDetailOpenApiVO) HasIssuedOn() bool`

HasIssuedOn returns a boolean if a field has been set.

### GetIssuer

`func (o *CertProfileDetailOpenApiVO) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertProfileDetailOpenApiVO) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertProfileDetailOpenApiVO) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CertProfileDetailOpenApiVO) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetRsaKeySize

`func (o *CertProfileDetailOpenApiVO) GetRsaKeySize() string`

GetRsaKeySize returns the RsaKeySize field if non-nil, zero value otherwise.

### GetRsaKeySizeOk

`func (o *CertProfileDetailOpenApiVO) GetRsaKeySizeOk() (*string, bool)`

GetRsaKeySizeOk returns a tuple with the RsaKeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaKeySize

`func (o *CertProfileDetailOpenApiVO) SetRsaKeySize(v string)`

SetRsaKeySize sets RsaKeySize field to given value.

### HasRsaKeySize

`func (o *CertProfileDetailOpenApiVO) HasRsaKeySize() bool`

HasRsaKeySize returns a boolean if a field has been set.

### GetSerialNumber

`func (o *CertProfileDetailOpenApiVO) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertProfileDetailOpenApiVO) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertProfileDetailOpenApiVO) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertProfileDetailOpenApiVO) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSignedUsing

`func (o *CertProfileDetailOpenApiVO) GetSignedUsing() string`

GetSignedUsing returns the SignedUsing field if non-nil, zero value otherwise.

### GetSignedUsingOk

`func (o *CertProfileDetailOpenApiVO) GetSignedUsingOk() (*string, bool)`

GetSignedUsingOk returns a tuple with the SignedUsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedUsing

`func (o *CertProfileDetailOpenApiVO) SetSignedUsing(v string)`

SetSignedUsing sets SignedUsing field to given value.

### HasSignedUsing

`func (o *CertProfileDetailOpenApiVO) HasSignedUsing() bool`

HasSignedUsing returns a boolean if a field has been set.

### GetSubject

`func (o *CertProfileDetailOpenApiVO) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CertProfileDetailOpenApiVO) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CertProfileDetailOpenApiVO) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CertProfileDetailOpenApiVO) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetType

`func (o *CertProfileDetailOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertProfileDetailOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertProfileDetailOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CertProfileDetailOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *CertProfileDetailOpenApiVO) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CertProfileDetailOpenApiVO) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CertProfileDetailOpenApiVO) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CertProfileDetailOpenApiVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


