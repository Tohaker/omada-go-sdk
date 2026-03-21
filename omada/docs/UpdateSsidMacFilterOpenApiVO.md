# UpdateSsidMacFilterOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacFilterEnable** | **bool** | SSID MAC Filter global config status. True: enable, false: disable. | 
**MacFilterId** | Pointer to **string** | This field represents MAC Group Profile ID. MAC Group Profile can be created using Create a new group profile interface, and MAC Group Profile ID can be obtained from Get group profile list by type interface. | [optional] 
**OuiProfileIdList** | Pointer to **[]string** | This field represents OUI Profile ID list. OUI Profile can be created using Create OUI profile interface, and OUI Profile ID can be obtained from Get OUI profile summary list interface(This configuration applies to the Pro Site of the Omada Pro Controller only). | [optional] 
**Policy** | Pointer to **int32** | SSID MAC Filter policy config mode; It should be a value as follows: 0: Deny List, 1: Allow List. | [optional] 

## Methods

### NewUpdateSsidMacFilterOpenApiVO

`func NewUpdateSsidMacFilterOpenApiVO(macFilterEnable bool, ) *UpdateSsidMacFilterOpenApiVO`

NewUpdateSsidMacFilterOpenApiVO instantiates a new UpdateSsidMacFilterOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSsidMacFilterOpenApiVOWithDefaults

`func NewUpdateSsidMacFilterOpenApiVOWithDefaults() *UpdateSsidMacFilterOpenApiVO`

NewUpdateSsidMacFilterOpenApiVOWithDefaults instantiates a new UpdateSsidMacFilterOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacFilterEnable

`func (o *UpdateSsidMacFilterOpenApiVO) GetMacFilterEnable() bool`

GetMacFilterEnable returns the MacFilterEnable field if non-nil, zero value otherwise.

### GetMacFilterEnableOk

`func (o *UpdateSsidMacFilterOpenApiVO) GetMacFilterEnableOk() (*bool, bool)`

GetMacFilterEnableOk returns a tuple with the MacFilterEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFilterEnable

`func (o *UpdateSsidMacFilterOpenApiVO) SetMacFilterEnable(v bool)`

SetMacFilterEnable sets MacFilterEnable field to given value.


### GetMacFilterId

`func (o *UpdateSsidMacFilterOpenApiVO) GetMacFilterId() string`

GetMacFilterId returns the MacFilterId field if non-nil, zero value otherwise.

### GetMacFilterIdOk

`func (o *UpdateSsidMacFilterOpenApiVO) GetMacFilterIdOk() (*string, bool)`

GetMacFilterIdOk returns a tuple with the MacFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFilterId

`func (o *UpdateSsidMacFilterOpenApiVO) SetMacFilterId(v string)`

SetMacFilterId sets MacFilterId field to given value.

### HasMacFilterId

`func (o *UpdateSsidMacFilterOpenApiVO) HasMacFilterId() bool`

HasMacFilterId returns a boolean if a field has been set.

### GetOuiProfileIdList

`func (o *UpdateSsidMacFilterOpenApiVO) GetOuiProfileIdList() []string`

GetOuiProfileIdList returns the OuiProfileIdList field if non-nil, zero value otherwise.

### GetOuiProfileIdListOk

`func (o *UpdateSsidMacFilterOpenApiVO) GetOuiProfileIdListOk() (*[]string, bool)`

GetOuiProfileIdListOk returns a tuple with the OuiProfileIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuiProfileIdList

`func (o *UpdateSsidMacFilterOpenApiVO) SetOuiProfileIdList(v []string)`

SetOuiProfileIdList sets OuiProfileIdList field to given value.

### HasOuiProfileIdList

`func (o *UpdateSsidMacFilterOpenApiVO) HasOuiProfileIdList() bool`

HasOuiProfileIdList returns a boolean if a field has been set.

### GetPolicy

`func (o *UpdateSsidMacFilterOpenApiVO) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateSsidMacFilterOpenApiVO) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateSsidMacFilterOpenApiVO) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UpdateSsidMacFilterOpenApiVO) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


