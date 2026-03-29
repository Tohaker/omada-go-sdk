# RadiusAuthServerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCert** | Pointer to **string** | CA certification profile id | [optional] 
**ClientCert** | Pointer to **string** | Client certification profile id | [optional] 
**RadSecEnable** | Pointer to **bool** | Radius RadSec enable status | [optional] 
**RadiusPort** | **int32** | Radius authentication server port, radiusPort should be within the range of 1-65535 | 
**RadiusPwd** | **string** | Radius authentication server password, radiusPwd should contain 1 to 64 characters,The question mark (?), double quote (\&quot;), percent sign (%), and backslash (\\) may cause the RADIUS function to fail and are not recommended. | 
**RadiusServerIp** | **string** | Radius authentication server IP. In Pro Site of the Omada Pro Controller, [radiusServerIp] should be a valid IP or domain address. In Omada Controller and Basic Site of the Omada Pro Controller, [radiusServerIp] should be a valid IP address. | 

## Methods

### NewRadiusAuthServerOpenApiVO

`func NewRadiusAuthServerOpenApiVO(radiusPort int32, radiusPwd string, radiusServerIp string, ) *RadiusAuthServerOpenApiVO`

NewRadiusAuthServerOpenApiVO instantiates a new RadiusAuthServerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusAuthServerOpenApiVOWithDefaults

`func NewRadiusAuthServerOpenApiVOWithDefaults() *RadiusAuthServerOpenApiVO`

NewRadiusAuthServerOpenApiVOWithDefaults instantiates a new RadiusAuthServerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCert

`func (o *RadiusAuthServerOpenApiVO) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *RadiusAuthServerOpenApiVO) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *RadiusAuthServerOpenApiVO) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *RadiusAuthServerOpenApiVO) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetClientCert

`func (o *RadiusAuthServerOpenApiVO) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *RadiusAuthServerOpenApiVO) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *RadiusAuthServerOpenApiVO) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *RadiusAuthServerOpenApiVO) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetRadSecEnable

`func (o *RadiusAuthServerOpenApiVO) GetRadSecEnable() bool`

GetRadSecEnable returns the RadSecEnable field if non-nil, zero value otherwise.

### GetRadSecEnableOk

`func (o *RadiusAuthServerOpenApiVO) GetRadSecEnableOk() (*bool, bool)`

GetRadSecEnableOk returns a tuple with the RadSecEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadSecEnable

`func (o *RadiusAuthServerOpenApiVO) SetRadSecEnable(v bool)`

SetRadSecEnable sets RadSecEnable field to given value.

### HasRadSecEnable

`func (o *RadiusAuthServerOpenApiVO) HasRadSecEnable() bool`

HasRadSecEnable returns a boolean if a field has been set.

### GetRadiusPort

`func (o *RadiusAuthServerOpenApiVO) GetRadiusPort() int32`

GetRadiusPort returns the RadiusPort field if non-nil, zero value otherwise.

### GetRadiusPortOk

`func (o *RadiusAuthServerOpenApiVO) GetRadiusPortOk() (*int32, bool)`

GetRadiusPortOk returns a tuple with the RadiusPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusPort

`func (o *RadiusAuthServerOpenApiVO) SetRadiusPort(v int32)`

SetRadiusPort sets RadiusPort field to given value.


### GetRadiusPwd

`func (o *RadiusAuthServerOpenApiVO) GetRadiusPwd() string`

GetRadiusPwd returns the RadiusPwd field if non-nil, zero value otherwise.

### GetRadiusPwdOk

`func (o *RadiusAuthServerOpenApiVO) GetRadiusPwdOk() (*string, bool)`

GetRadiusPwdOk returns a tuple with the RadiusPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusPwd

`func (o *RadiusAuthServerOpenApiVO) SetRadiusPwd(v string)`

SetRadiusPwd sets RadiusPwd field to given value.


### GetRadiusServerIp

`func (o *RadiusAuthServerOpenApiVO) GetRadiusServerIp() string`

GetRadiusServerIp returns the RadiusServerIp field if non-nil, zero value otherwise.

### GetRadiusServerIpOk

`func (o *RadiusAuthServerOpenApiVO) GetRadiusServerIpOk() (*string, bool)`

GetRadiusServerIpOk returns a tuple with the RadiusServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServerIp

`func (o *RadiusAuthServerOpenApiVO) SetRadiusServerIp(v string)`

SetRadiusServerIp sets RadiusServerIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


