# PortalAccessControlOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeAuthClientEnable** | **bool** | Whether to enable Free-Authentication Client. If parameter [freeAuthClientEnable] is true, parameter [freeAuthClientPolicies] is needed | 
**FreeAuthClientPolicies** | Pointer to [**[]FreeAuthClientPolicyOpenApiVO**](FreeAuthClientPolicyOpenApiVO.md) | List of Free-Authentication Client Policy | [optional] 
**PreAuthAccessEnable** | **bool** | Whether to enable Pre-Authentication Access. If parameter [preAuthAccessEnable] is true, parameter [preAuthAccessPolicies] is needed | 
**PreAuthAccessPolicies** | Pointer to [**[]PreAuthAccessPolicyOpenApiVO**](PreAuthAccessPolicyOpenApiVO.md) | List of Pre-Authentication Access Policy | [optional] 

## Methods

### NewPortalAccessControlOpenApiVO

`func NewPortalAccessControlOpenApiVO(freeAuthClientEnable bool, preAuthAccessEnable bool, ) *PortalAccessControlOpenApiVO`

NewPortalAccessControlOpenApiVO instantiates a new PortalAccessControlOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalAccessControlOpenApiVOWithDefaults

`func NewPortalAccessControlOpenApiVOWithDefaults() *PortalAccessControlOpenApiVO`

NewPortalAccessControlOpenApiVOWithDefaults instantiates a new PortalAccessControlOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeAuthClientEnable

`func (o *PortalAccessControlOpenApiVO) GetFreeAuthClientEnable() bool`

GetFreeAuthClientEnable returns the FreeAuthClientEnable field if non-nil, zero value otherwise.

### GetFreeAuthClientEnableOk

`func (o *PortalAccessControlOpenApiVO) GetFreeAuthClientEnableOk() (*bool, bool)`

GetFreeAuthClientEnableOk returns a tuple with the FreeAuthClientEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeAuthClientEnable

`func (o *PortalAccessControlOpenApiVO) SetFreeAuthClientEnable(v bool)`

SetFreeAuthClientEnable sets FreeAuthClientEnable field to given value.


### GetFreeAuthClientPolicies

`func (o *PortalAccessControlOpenApiVO) GetFreeAuthClientPolicies() []FreeAuthClientPolicyOpenApiVO`

GetFreeAuthClientPolicies returns the FreeAuthClientPolicies field if non-nil, zero value otherwise.

### GetFreeAuthClientPoliciesOk

`func (o *PortalAccessControlOpenApiVO) GetFreeAuthClientPoliciesOk() (*[]FreeAuthClientPolicyOpenApiVO, bool)`

GetFreeAuthClientPoliciesOk returns a tuple with the FreeAuthClientPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeAuthClientPolicies

`func (o *PortalAccessControlOpenApiVO) SetFreeAuthClientPolicies(v []FreeAuthClientPolicyOpenApiVO)`

SetFreeAuthClientPolicies sets FreeAuthClientPolicies field to given value.

### HasFreeAuthClientPolicies

`func (o *PortalAccessControlOpenApiVO) HasFreeAuthClientPolicies() bool`

HasFreeAuthClientPolicies returns a boolean if a field has been set.

### GetPreAuthAccessEnable

`func (o *PortalAccessControlOpenApiVO) GetPreAuthAccessEnable() bool`

GetPreAuthAccessEnable returns the PreAuthAccessEnable field if non-nil, zero value otherwise.

### GetPreAuthAccessEnableOk

`func (o *PortalAccessControlOpenApiVO) GetPreAuthAccessEnableOk() (*bool, bool)`

GetPreAuthAccessEnableOk returns a tuple with the PreAuthAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthAccessEnable

`func (o *PortalAccessControlOpenApiVO) SetPreAuthAccessEnable(v bool)`

SetPreAuthAccessEnable sets PreAuthAccessEnable field to given value.


### GetPreAuthAccessPolicies

`func (o *PortalAccessControlOpenApiVO) GetPreAuthAccessPolicies() []PreAuthAccessPolicyOpenApiVO`

GetPreAuthAccessPolicies returns the PreAuthAccessPolicies field if non-nil, zero value otherwise.

### GetPreAuthAccessPoliciesOk

`func (o *PortalAccessControlOpenApiVO) GetPreAuthAccessPoliciesOk() (*[]PreAuthAccessPolicyOpenApiVO, bool)`

GetPreAuthAccessPoliciesOk returns a tuple with the PreAuthAccessPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthAccessPolicies

`func (o *PortalAccessControlOpenApiVO) SetPreAuthAccessPolicies(v []PreAuthAccessPolicyOpenApiVO)`

SetPreAuthAccessPolicies sets PreAuthAccessPolicies field to given value.

### HasPreAuthAccessPolicies

`func (o *PortalAccessControlOpenApiVO) HasPreAuthAccessPolicies() bool`

HasPreAuthAccessPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


