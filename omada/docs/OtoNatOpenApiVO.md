# OtoNatOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description should contain 1 to 64 characters. | [optional] 
**Dmz** | **bool** | Whether to enable the h323 dmz | 
**ExternalIp** | **string** | External IP | 
**InterfaceIds** | **[]string** | This field represents WAN port ID. WAN port ID can be obtained from can be obtained from &#39;Get internet basic info&#39; interface. | 
**InternalIp** | **string** | Internal IP | 
**Name** | **string** | Name, name should contain 1 to 64 characters. | 
**Status** | **bool** | Status | 

## Methods

### NewOtoNatOpenApiVO

`func NewOtoNatOpenApiVO(dmz bool, externalIp string, interfaceIds []string, internalIp string, name string, status bool, ) *OtoNatOpenApiVO`

NewOtoNatOpenApiVO instantiates a new OtoNatOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtoNatOpenApiVOWithDefaults

`func NewOtoNatOpenApiVOWithDefaults() *OtoNatOpenApiVO`

NewOtoNatOpenApiVOWithDefaults instantiates a new OtoNatOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OtoNatOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OtoNatOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OtoNatOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OtoNatOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDmz

`func (o *OtoNatOpenApiVO) GetDmz() bool`

GetDmz returns the Dmz field if non-nil, zero value otherwise.

### GetDmzOk

`func (o *OtoNatOpenApiVO) GetDmzOk() (*bool, bool)`

GetDmzOk returns a tuple with the Dmz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmz

`func (o *OtoNatOpenApiVO) SetDmz(v bool)`

SetDmz sets Dmz field to given value.


### GetExternalIp

`func (o *OtoNatOpenApiVO) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *OtoNatOpenApiVO) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *OtoNatOpenApiVO) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.


### GetInterfaceIds

`func (o *OtoNatOpenApiVO) GetInterfaceIds() []string`

GetInterfaceIds returns the InterfaceIds field if non-nil, zero value otherwise.

### GetInterfaceIdsOk

`func (o *OtoNatOpenApiVO) GetInterfaceIdsOk() (*[]string, bool)`

GetInterfaceIdsOk returns a tuple with the InterfaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIds

`func (o *OtoNatOpenApiVO) SetInterfaceIds(v []string)`

SetInterfaceIds sets InterfaceIds field to given value.


### GetInternalIp

`func (o *OtoNatOpenApiVO) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *OtoNatOpenApiVO) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *OtoNatOpenApiVO) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.


### GetName

`func (o *OtoNatOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OtoNatOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OtoNatOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *OtoNatOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OtoNatOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OtoNatOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


