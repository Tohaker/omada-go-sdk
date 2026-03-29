# PortalResOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **int32** | The type of authentication should be a value as follows: 0: No Auth; 1: Simple Password; 2: External Radius; 4: External Portal Server; 11: Hotspot; 15: Ldap; 16: Social Login | [optional] 
**Enable** | Pointer to **bool** | Is the portal enable. | [optional] 
**HotspotTypes** | Pointer to **[]int32** | The enable types of hotspot should be a value as follows: 3: Voucher; 5: Local user 6: SMS; 8: Hotspot Radius; 12: Form Auth | [optional] 
**Id** | Pointer to **string** | Portal ID | [optional] 
**Name** | Pointer to **string** | Portal name | [optional] 
**NetworkList** | Pointer to **[]string** | The network ID list of the portal binding. | [optional] 
**SocialLoginTypes** | Pointer to **[]int32** | The enable types of social login should be a value as follows: 17: Google | [optional] 
**SsidList** | Pointer to **[]string** | The ssid list of the portal binding. | [optional] 

## Methods

### NewPortalResOpenApiVO

`func NewPortalResOpenApiVO() *PortalResOpenApiVO`

NewPortalResOpenApiVO instantiates a new PortalResOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalResOpenApiVOWithDefaults

`func NewPortalResOpenApiVOWithDefaults() *PortalResOpenApiVO`

NewPortalResOpenApiVOWithDefaults instantiates a new PortalResOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *PortalResOpenApiVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *PortalResOpenApiVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *PortalResOpenApiVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *PortalResOpenApiVO) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetEnable

`func (o *PortalResOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PortalResOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PortalResOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PortalResOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetHotspotTypes

`func (o *PortalResOpenApiVO) GetHotspotTypes() []int32`

GetHotspotTypes returns the HotspotTypes field if non-nil, zero value otherwise.

### GetHotspotTypesOk

`func (o *PortalResOpenApiVO) GetHotspotTypesOk() (*[]int32, bool)`

GetHotspotTypesOk returns a tuple with the HotspotTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotspotTypes

`func (o *PortalResOpenApiVO) SetHotspotTypes(v []int32)`

SetHotspotTypes sets HotspotTypes field to given value.

### HasHotspotTypes

`func (o *PortalResOpenApiVO) HasHotspotTypes() bool`

HasHotspotTypes returns a boolean if a field has been set.

### GetId

`func (o *PortalResOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalResOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalResOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortalResOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PortalResOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortalResOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortalResOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PortalResOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkList

`func (o *PortalResOpenApiVO) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *PortalResOpenApiVO) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *PortalResOpenApiVO) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *PortalResOpenApiVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetSocialLoginTypes

`func (o *PortalResOpenApiVO) GetSocialLoginTypes() []int32`

GetSocialLoginTypes returns the SocialLoginTypes field if non-nil, zero value otherwise.

### GetSocialLoginTypesOk

`func (o *PortalResOpenApiVO) GetSocialLoginTypesOk() (*[]int32, bool)`

GetSocialLoginTypesOk returns a tuple with the SocialLoginTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLoginTypes

`func (o *PortalResOpenApiVO) SetSocialLoginTypes(v []int32)`

SetSocialLoginTypes sets SocialLoginTypes field to given value.

### HasSocialLoginTypes

`func (o *PortalResOpenApiVO) HasSocialLoginTypes() bool`

HasSocialLoginTypes returns a boolean if a field has been set.

### GetSsidList

`func (o *PortalResOpenApiVO) GetSsidList() []string`

GetSsidList returns the SsidList field if non-nil, zero value otherwise.

### GetSsidListOk

`func (o *PortalResOpenApiVO) GetSsidListOk() (*[]string, bool)`

GetSsidListOk returns a tuple with the SsidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidList

`func (o *PortalResOpenApiVO) SetSsidList(v []string)`

SetSsidList sets SsidList field to given value.

### HasSsidList

`func (o *PortalResOpenApiVO) HasSsidList() bool`

HasSsidList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


