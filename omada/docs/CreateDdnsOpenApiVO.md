# CreateDdnsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** | Dynamic DNS domainName, valid when parameter [service] is 0, 1 or 4 | [optional] 
**InterfacePortId** | **string** | This field represents interface Port ID | 
**Interval** | Pointer to [**DdnsIntervalOpenApiVO**](DdnsIntervalOpenApiVO.md) |  | [optional] 
**Password** | **string** | Dynamic DNS password. Password should contain 1 to 128 characters | 
**Service** | **int32** | Dynamic DNS service type. Service should be a value as follows: 0: DynDNS, 1: NO-IP, 2: Peanuthull, 3: Comexe, 4: Custom, 5: TP-Link DDNS | 
**Status** | **bool** | Dynamic DNS enable status | 
**UpdateUrl** | Pointer to **string** | Dynamic DNS updateUrl, valid when parameter [service] is 4. Update URL must contain [USERNAME], [PASSWORD] and [DOMAIN], and Update-URL will be applied to all custom entries | [optional] 
**Username** | **string** | Dynamic DNS username. Username should contain 1 to 128 characters | 

## Methods

### NewCreateDdnsOpenApiVO

`func NewCreateDdnsOpenApiVO(interfacePortId string, password string, service int32, status bool, username string, ) *CreateDdnsOpenApiVO`

NewCreateDdnsOpenApiVO instantiates a new CreateDdnsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDdnsOpenApiVOWithDefaults

`func NewCreateDdnsOpenApiVOWithDefaults() *CreateDdnsOpenApiVO`

NewCreateDdnsOpenApiVOWithDefaults instantiates a new CreateDdnsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *CreateDdnsOpenApiVO) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CreateDdnsOpenApiVO) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CreateDdnsOpenApiVO) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CreateDdnsOpenApiVO) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetInterfacePortId

`func (o *CreateDdnsOpenApiVO) GetInterfacePortId() string`

GetInterfacePortId returns the InterfacePortId field if non-nil, zero value otherwise.

### GetInterfacePortIdOk

`func (o *CreateDdnsOpenApiVO) GetInterfacePortIdOk() (*string, bool)`

GetInterfacePortIdOk returns a tuple with the InterfacePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfacePortId

`func (o *CreateDdnsOpenApiVO) SetInterfacePortId(v string)`

SetInterfacePortId sets InterfacePortId field to given value.


### GetInterval

`func (o *CreateDdnsOpenApiVO) GetInterval() DdnsIntervalOpenApiVO`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CreateDdnsOpenApiVO) GetIntervalOk() (*DdnsIntervalOpenApiVO, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CreateDdnsOpenApiVO) SetInterval(v DdnsIntervalOpenApiVO)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CreateDdnsOpenApiVO) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetPassword

`func (o *CreateDdnsOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDdnsOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDdnsOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetService

`func (o *CreateDdnsOpenApiVO) GetService() int32`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateDdnsOpenApiVO) GetServiceOk() (*int32, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateDdnsOpenApiVO) SetService(v int32)`

SetService sets Service field to given value.


### GetStatus

`func (o *CreateDdnsOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateDdnsOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateDdnsOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetUpdateUrl

`func (o *CreateDdnsOpenApiVO) GetUpdateUrl() string`

GetUpdateUrl returns the UpdateUrl field if non-nil, zero value otherwise.

### GetUpdateUrlOk

`func (o *CreateDdnsOpenApiVO) GetUpdateUrlOk() (*string, bool)`

GetUpdateUrlOk returns a tuple with the UpdateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUrl

`func (o *CreateDdnsOpenApiVO) SetUpdateUrl(v string)`

SetUpdateUrl sets UpdateUrl field to given value.

### HasUpdateUrl

`func (o *CreateDdnsOpenApiVO) HasUpdateUrl() bool`

HasUpdateUrl returns a boolean if a field has been set.

### GetUsername

`func (o *CreateDdnsOpenApiVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateDdnsOpenApiVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateDdnsOpenApiVO) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


