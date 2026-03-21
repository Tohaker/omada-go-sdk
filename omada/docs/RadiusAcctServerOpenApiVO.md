# RadiusAcctServerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountingServerIp** | **string** | Radius Accounting Server IP.In Pro Site of the Omada Pro Controller, [accountingServerIp] should be a valid IP or domain address. In Omada Controller and Basic Site of the Omada Pro Controller, [accountingServerIp] should be a valid IP address. | 
**AccountingServerPort** | **int32** | Radius Accounting port. AccountingServerPort should be within the range of 1-65535 | 
**AccountingServerPwd** | **string** | Radius Accounting password,The question mark (?), double quote (\&quot;), percent sign (%), and backslash (\\) may cause the RADIUS function to fail and are not recommended. | 
**CaCert** | Pointer to **string** | CA certification profile id | [optional] 
**ClientCert** | Pointer to **string** | Client certification profile id | [optional] 
**RadSecEnable** | Pointer to **bool** | Radius RadSec enable status | [optional] 

## Methods

### NewRadiusAcctServerOpenApiVO

`func NewRadiusAcctServerOpenApiVO(accountingServerIp string, accountingServerPort int32, accountingServerPwd string, ) *RadiusAcctServerOpenApiVO`

NewRadiusAcctServerOpenApiVO instantiates a new RadiusAcctServerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusAcctServerOpenApiVOWithDefaults

`func NewRadiusAcctServerOpenApiVOWithDefaults() *RadiusAcctServerOpenApiVO`

NewRadiusAcctServerOpenApiVOWithDefaults instantiates a new RadiusAcctServerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountingServerIp

`func (o *RadiusAcctServerOpenApiVO) GetAccountingServerIp() string`

GetAccountingServerIp returns the AccountingServerIp field if non-nil, zero value otherwise.

### GetAccountingServerIpOk

`func (o *RadiusAcctServerOpenApiVO) GetAccountingServerIpOk() (*string, bool)`

GetAccountingServerIpOk returns a tuple with the AccountingServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingServerIp

`func (o *RadiusAcctServerOpenApiVO) SetAccountingServerIp(v string)`

SetAccountingServerIp sets AccountingServerIp field to given value.


### GetAccountingServerPort

`func (o *RadiusAcctServerOpenApiVO) GetAccountingServerPort() int32`

GetAccountingServerPort returns the AccountingServerPort field if non-nil, zero value otherwise.

### GetAccountingServerPortOk

`func (o *RadiusAcctServerOpenApiVO) GetAccountingServerPortOk() (*int32, bool)`

GetAccountingServerPortOk returns a tuple with the AccountingServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingServerPort

`func (o *RadiusAcctServerOpenApiVO) SetAccountingServerPort(v int32)`

SetAccountingServerPort sets AccountingServerPort field to given value.


### GetAccountingServerPwd

`func (o *RadiusAcctServerOpenApiVO) GetAccountingServerPwd() string`

GetAccountingServerPwd returns the AccountingServerPwd field if non-nil, zero value otherwise.

### GetAccountingServerPwdOk

`func (o *RadiusAcctServerOpenApiVO) GetAccountingServerPwdOk() (*string, bool)`

GetAccountingServerPwdOk returns a tuple with the AccountingServerPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingServerPwd

`func (o *RadiusAcctServerOpenApiVO) SetAccountingServerPwd(v string)`

SetAccountingServerPwd sets AccountingServerPwd field to given value.


### GetCaCert

`func (o *RadiusAcctServerOpenApiVO) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *RadiusAcctServerOpenApiVO) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *RadiusAcctServerOpenApiVO) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *RadiusAcctServerOpenApiVO) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetClientCert

`func (o *RadiusAcctServerOpenApiVO) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *RadiusAcctServerOpenApiVO) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *RadiusAcctServerOpenApiVO) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *RadiusAcctServerOpenApiVO) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetRadSecEnable

`func (o *RadiusAcctServerOpenApiVO) GetRadSecEnable() bool`

GetRadSecEnable returns the RadSecEnable field if non-nil, zero value otherwise.

### GetRadSecEnableOk

`func (o *RadiusAcctServerOpenApiVO) GetRadSecEnableOk() (*bool, bool)`

GetRadSecEnableOk returns a tuple with the RadSecEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadSecEnable

`func (o *RadiusAcctServerOpenApiVO) SetRadSecEnable(v bool)`

SetRadSecEnable sets RadSecEnable field to given value.

### HasRadSecEnable

`func (o *RadiusAcctServerOpenApiVO) HasRadSecEnable() bool`

HasRadSecEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


