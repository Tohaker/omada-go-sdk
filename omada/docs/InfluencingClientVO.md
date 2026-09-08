# InfluencingClientVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Client status (true &#x3D; connected, false &#x3D; disconnected) | [optional] 
**AuthStatus** | Pointer to **int32** | Authentication status.0: CONNECTED // Access without any authentication method.1: PENDING // Access to Portal, but authentication failed.2: AUTHORIZED // Pass through portal, pass other authentication without portal.3: AUTH-FREE // No portal authentication required.4: OFFLINE // Offline client. | [optional] 
**ClientName** | Pointer to **string** | Display name of the client | [optional] 
**ClientType** | Pointer to **string** | Type of the client device | [optional] 
**Health** | Pointer to **int32** | Client health score | [optional] 
**Ip** | Pointer to **string** | IP address of the client | [optional] 
**Mac** | Pointer to **string** | MAC address of the client | [optional] 
**Manager** | Pointer to **bool** | Whether the client is a managing host | [optional] 
**Model** | Pointer to **string** | Client model | [optional] 

## Methods

### NewInfluencingClientVO

`func NewInfluencingClientVO() *InfluencingClientVO`

NewInfluencingClientVO instantiates a new InfluencingClientVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfluencingClientVOWithDefaults

`func NewInfluencingClientVOWithDefaults() *InfluencingClientVO`

NewInfluencingClientVOWithDefaults instantiates a new InfluencingClientVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *InfluencingClientVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InfluencingClientVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InfluencingClientVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InfluencingClientVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAuthStatus

`func (o *InfluencingClientVO) GetAuthStatus() int32`

GetAuthStatus returns the AuthStatus field if non-nil, zero value otherwise.

### GetAuthStatusOk

`func (o *InfluencingClientVO) GetAuthStatusOk() (*int32, bool)`

GetAuthStatusOk returns a tuple with the AuthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthStatus

`func (o *InfluencingClientVO) SetAuthStatus(v int32)`

SetAuthStatus sets AuthStatus field to given value.

### HasAuthStatus

`func (o *InfluencingClientVO) HasAuthStatus() bool`

HasAuthStatus returns a boolean if a field has been set.

### GetClientName

`func (o *InfluencingClientVO) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *InfluencingClientVO) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *InfluencingClientVO) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *InfluencingClientVO) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientType

`func (o *InfluencingClientVO) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *InfluencingClientVO) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *InfluencingClientVO) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *InfluencingClientVO) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetHealth

`func (o *InfluencingClientVO) GetHealth() int32`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *InfluencingClientVO) GetHealthOk() (*int32, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *InfluencingClientVO) SetHealth(v int32)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *InfluencingClientVO) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetIp

`func (o *InfluencingClientVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InfluencingClientVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InfluencingClientVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InfluencingClientVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *InfluencingClientVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InfluencingClientVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InfluencingClientVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InfluencingClientVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManager

`func (o *InfluencingClientVO) GetManager() bool`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *InfluencingClientVO) GetManagerOk() (*bool, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *InfluencingClientVO) SetManager(v bool)`

SetManager sets Manager field to given value.

### HasManager

`func (o *InfluencingClientVO) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetModel

`func (o *InfluencingClientVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InfluencingClientVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InfluencingClientVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InfluencingClientVO) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


