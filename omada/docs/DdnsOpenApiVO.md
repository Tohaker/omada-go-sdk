# DdnsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomInterval** | Pointer to **int32** | Dynamic DNS custom interval, valid when parameter [server] is 0, 1 or 4, unit: minute | [optional] 
**DomainName** | Pointer to **string** | Dynamic DNS domainName | [optional] 
**ExistCustomDdns** | Pointer to **bool** | Whether Custom Service Provider has been configured in the current Dynamic DNS. | [optional] 
**ExistCustomInterval** | Pointer to **bool** | Whether Custom Update Interval has been configured in the current Dynamic DNS. | [optional] 
**ExistTpLinkddns** | Pointer to **bool** | Whether TP-Link has been configured as Service Provider in the current Dynamic DNS. | [optional] 
**Id** | Pointer to **string** | Dynamic DNS ID | [optional] 
**InterfacePortId** | Pointer to **string** | Port ID of interface | [optional] 
**Password** | Pointer to **string** | Dynamic DNS password | [optional] 
**Service** | Pointer to **int32** | Dynamic DNS service type, 0: DynDNS, 1: NO-IP, 2: Peanuthull, 3: Comexe, 4: Custom, 5: TP-Link DDNS | [optional] 
**Status** | Pointer to **bool** | Dynamic DNS enable status | [optional] 
**UpdateInterval** | Pointer to **int32** | Dynamic DNS update interval, unit: hour | [optional] 
**UpdateUrl** | Pointer to **string** | Dynamic DNS updateUrl, valid when parameter [server] is 4 | [optional] 
**Username** | Pointer to **string** | Dynamic DNS username | [optional] 

## Methods

### NewDdnsOpenApiVO

`func NewDdnsOpenApiVO() *DdnsOpenApiVO`

NewDdnsOpenApiVO instantiates a new DdnsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdnsOpenApiVOWithDefaults

`func NewDdnsOpenApiVOWithDefaults() *DdnsOpenApiVO`

NewDdnsOpenApiVOWithDefaults instantiates a new DdnsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomInterval

`func (o *DdnsOpenApiVO) GetCustomInterval() int32`

GetCustomInterval returns the CustomInterval field if non-nil, zero value otherwise.

### GetCustomIntervalOk

`func (o *DdnsOpenApiVO) GetCustomIntervalOk() (*int32, bool)`

GetCustomIntervalOk returns a tuple with the CustomInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInterval

`func (o *DdnsOpenApiVO) SetCustomInterval(v int32)`

SetCustomInterval sets CustomInterval field to given value.

### HasCustomInterval

`func (o *DdnsOpenApiVO) HasCustomInterval() bool`

HasCustomInterval returns a boolean if a field has been set.

### GetDomainName

`func (o *DdnsOpenApiVO) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DdnsOpenApiVO) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DdnsOpenApiVO) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DdnsOpenApiVO) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetExistCustomDdns

`func (o *DdnsOpenApiVO) GetExistCustomDdns() bool`

GetExistCustomDdns returns the ExistCustomDdns field if non-nil, zero value otherwise.

### GetExistCustomDdnsOk

`func (o *DdnsOpenApiVO) GetExistCustomDdnsOk() (*bool, bool)`

GetExistCustomDdnsOk returns a tuple with the ExistCustomDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistCustomDdns

`func (o *DdnsOpenApiVO) SetExistCustomDdns(v bool)`

SetExistCustomDdns sets ExistCustomDdns field to given value.

### HasExistCustomDdns

`func (o *DdnsOpenApiVO) HasExistCustomDdns() bool`

HasExistCustomDdns returns a boolean if a field has been set.

### GetExistCustomInterval

`func (o *DdnsOpenApiVO) GetExistCustomInterval() bool`

GetExistCustomInterval returns the ExistCustomInterval field if non-nil, zero value otherwise.

### GetExistCustomIntervalOk

`func (o *DdnsOpenApiVO) GetExistCustomIntervalOk() (*bool, bool)`

GetExistCustomIntervalOk returns a tuple with the ExistCustomInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistCustomInterval

`func (o *DdnsOpenApiVO) SetExistCustomInterval(v bool)`

SetExistCustomInterval sets ExistCustomInterval field to given value.

### HasExistCustomInterval

`func (o *DdnsOpenApiVO) HasExistCustomInterval() bool`

HasExistCustomInterval returns a boolean if a field has been set.

### GetExistTpLinkddns

`func (o *DdnsOpenApiVO) GetExistTpLinkddns() bool`

GetExistTpLinkddns returns the ExistTpLinkddns field if non-nil, zero value otherwise.

### GetExistTpLinkddnsOk

`func (o *DdnsOpenApiVO) GetExistTpLinkddnsOk() (*bool, bool)`

GetExistTpLinkddnsOk returns a tuple with the ExistTpLinkddns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistTpLinkddns

`func (o *DdnsOpenApiVO) SetExistTpLinkddns(v bool)`

SetExistTpLinkddns sets ExistTpLinkddns field to given value.

### HasExistTpLinkddns

`func (o *DdnsOpenApiVO) HasExistTpLinkddns() bool`

HasExistTpLinkddns returns a boolean if a field has been set.

### GetId

`func (o *DdnsOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DdnsOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DdnsOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DdnsOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfacePortId

`func (o *DdnsOpenApiVO) GetInterfacePortId() string`

GetInterfacePortId returns the InterfacePortId field if non-nil, zero value otherwise.

### GetInterfacePortIdOk

`func (o *DdnsOpenApiVO) GetInterfacePortIdOk() (*string, bool)`

GetInterfacePortIdOk returns a tuple with the InterfacePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfacePortId

`func (o *DdnsOpenApiVO) SetInterfacePortId(v string)`

SetInterfacePortId sets InterfacePortId field to given value.

### HasInterfacePortId

`func (o *DdnsOpenApiVO) HasInterfacePortId() bool`

HasInterfacePortId returns a boolean if a field has been set.

### GetPassword

`func (o *DdnsOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DdnsOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DdnsOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DdnsOpenApiVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetService

`func (o *DdnsOpenApiVO) GetService() int32`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DdnsOpenApiVO) GetServiceOk() (*int32, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DdnsOpenApiVO) SetService(v int32)`

SetService sets Service field to given value.

### HasService

`func (o *DdnsOpenApiVO) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *DdnsOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DdnsOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DdnsOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DdnsOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdateInterval

`func (o *DdnsOpenApiVO) GetUpdateInterval() int32`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *DdnsOpenApiVO) GetUpdateIntervalOk() (*int32, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *DdnsOpenApiVO) SetUpdateInterval(v int32)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *DdnsOpenApiVO) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.

### GetUpdateUrl

`func (o *DdnsOpenApiVO) GetUpdateUrl() string`

GetUpdateUrl returns the UpdateUrl field if non-nil, zero value otherwise.

### GetUpdateUrlOk

`func (o *DdnsOpenApiVO) GetUpdateUrlOk() (*string, bool)`

GetUpdateUrlOk returns a tuple with the UpdateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUrl

`func (o *DdnsOpenApiVO) SetUpdateUrl(v string)`

SetUpdateUrl sets UpdateUrl field to given value.

### HasUpdateUrl

`func (o *DdnsOpenApiVO) HasUpdateUrl() bool`

HasUpdateUrl returns a boolean if a field has been set.

### GetUsername

`func (o *DdnsOpenApiVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DdnsOpenApiVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DdnsOpenApiVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DdnsOpenApiVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


