# OnuActivationConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethod** | **string** | Authentication mode used by the automatically discovered ONU.AuthenticationMethod should be a value as follows:PASSWORD_AUTH,SN_AUTH,LOID_AUTH,SN_AND_PASSWORD_AUTH,LOID_AND_PASSWORD_AUTH | 
**Description** | Pointer to **string** | Display the configured port description.Only 1-32 bits numbers, Upper and lower letters, -@_:/. are allowed. | [optional] 
**DiscoveryMode** | Pointer to **string** | ONU auto discovery mode.DiscoveryMode should be a value as follows:ALWAYS_ON,ONCE_ON  | [optional] 
**Keys** | **[]string** | Entry identifier list | 
**LineProfile** | **string** | Line profile id bound to the ONU | 
**OnuId** | Pointer to **int32** | Onu ID should be within the range of 0 to 127 | [optional] 
**ReRegisterAuthentication** | Pointer to **string** | Re-Register-Authentication | [optional] 
**ServicePortProfile** | Pointer to **string** | Service port profile id bound to the ONU | [optional] 
**ServiceProfile** | **string** | Service profile id bound to the ONU | 

## Methods

### NewOnuActivationConfigDTO

`func NewOnuActivationConfigDTO(authenticationMethod string, keys []string, lineProfile string, serviceProfile string, ) *OnuActivationConfigDTO`

NewOnuActivationConfigDTO instantiates a new OnuActivationConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuActivationConfigDTOWithDefaults

`func NewOnuActivationConfigDTOWithDefaults() *OnuActivationConfigDTO`

NewOnuActivationConfigDTOWithDefaults instantiates a new OnuActivationConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationMethod

`func (o *OnuActivationConfigDTO) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *OnuActivationConfigDTO) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *OnuActivationConfigDTO) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetDescription

`func (o *OnuActivationConfigDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OnuActivationConfigDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OnuActivationConfigDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OnuActivationConfigDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscoveryMode

`func (o *OnuActivationConfigDTO) GetDiscoveryMode() string`

GetDiscoveryMode returns the DiscoveryMode field if non-nil, zero value otherwise.

### GetDiscoveryModeOk

`func (o *OnuActivationConfigDTO) GetDiscoveryModeOk() (*string, bool)`

GetDiscoveryModeOk returns a tuple with the DiscoveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMode

`func (o *OnuActivationConfigDTO) SetDiscoveryMode(v string)`

SetDiscoveryMode sets DiscoveryMode field to given value.

### HasDiscoveryMode

`func (o *OnuActivationConfigDTO) HasDiscoveryMode() bool`

HasDiscoveryMode returns a boolean if a field has been set.

### GetKeys

`func (o *OnuActivationConfigDTO) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *OnuActivationConfigDTO) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *OnuActivationConfigDTO) SetKeys(v []string)`

SetKeys sets Keys field to given value.


### GetLineProfile

`func (o *OnuActivationConfigDTO) GetLineProfile() string`

GetLineProfile returns the LineProfile field if non-nil, zero value otherwise.

### GetLineProfileOk

`func (o *OnuActivationConfigDTO) GetLineProfileOk() (*string, bool)`

GetLineProfileOk returns a tuple with the LineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfile

`func (o *OnuActivationConfigDTO) SetLineProfile(v string)`

SetLineProfile sets LineProfile field to given value.


### GetOnuId

`func (o *OnuActivationConfigDTO) GetOnuId() int32`

GetOnuId returns the OnuId field if non-nil, zero value otherwise.

### GetOnuIdOk

`func (o *OnuActivationConfigDTO) GetOnuIdOk() (*int32, bool)`

GetOnuIdOk returns a tuple with the OnuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuId

`func (o *OnuActivationConfigDTO) SetOnuId(v int32)`

SetOnuId sets OnuId field to given value.

### HasOnuId

`func (o *OnuActivationConfigDTO) HasOnuId() bool`

HasOnuId returns a boolean if a field has been set.

### GetReRegisterAuthentication

`func (o *OnuActivationConfigDTO) GetReRegisterAuthentication() string`

GetReRegisterAuthentication returns the ReRegisterAuthentication field if non-nil, zero value otherwise.

### GetReRegisterAuthenticationOk

`func (o *OnuActivationConfigDTO) GetReRegisterAuthenticationOk() (*string, bool)`

GetReRegisterAuthenticationOk returns a tuple with the ReRegisterAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRegisterAuthentication

`func (o *OnuActivationConfigDTO) SetReRegisterAuthentication(v string)`

SetReRegisterAuthentication sets ReRegisterAuthentication field to given value.

### HasReRegisterAuthentication

`func (o *OnuActivationConfigDTO) HasReRegisterAuthentication() bool`

HasReRegisterAuthentication returns a boolean if a field has been set.

### GetServicePortProfile

`func (o *OnuActivationConfigDTO) GetServicePortProfile() string`

GetServicePortProfile returns the ServicePortProfile field if non-nil, zero value otherwise.

### GetServicePortProfileOk

`func (o *OnuActivationConfigDTO) GetServicePortProfileOk() (*string, bool)`

GetServicePortProfileOk returns a tuple with the ServicePortProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortProfile

`func (o *OnuActivationConfigDTO) SetServicePortProfile(v string)`

SetServicePortProfile sets ServicePortProfile field to given value.

### HasServicePortProfile

`func (o *OnuActivationConfigDTO) HasServicePortProfile() bool`

HasServicePortProfile returns a boolean if a field has been set.

### GetServiceProfile

`func (o *OnuActivationConfigDTO) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *OnuActivationConfigDTO) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *OnuActivationConfigDTO) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


