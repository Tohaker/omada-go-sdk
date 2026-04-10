# Dot1xSwitchInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountVrfId** | Pointer to **string** | Account VRF ID | [optional] 
**AddedInAdvanced** | Pointer to **bool** | Whether the device is added offline | [optional] 
**AuthVrfId** | Pointer to **string** | Auth VRF ID | [optional] 
**Mac** | Pointer to **string** | Switch MAC address | [optional] 
**Model** | Pointer to **string** | Switch model | [optional] 
**Name** | Pointer to **string** | Switch name | [optional] 
**PortConfigured** | Pointer to **bool** | Is there any port configured with 802.1x or MAB authentication. | [optional] 
**Ports** | Pointer to [**[]Dot1xPortInfoOpenApiVO**](Dot1xPortInfoOpenApiVO.md) | Switch port information | [optional] 
**Status** | Pointer to **string** | Device status | [optional] 
**StatusCategory** | Pointer to **int32** | Device status category, 0: Disconnected, 1: Connected, 2: Pending,3: Heartbeat Missed, 4: Isolated | [optional] 
**SupportSingleMabAuth** | Pointer to **bool** | Whether the switch supports single MAB Authentication.If not supported, it will not take effect when configuring single MAB authentication for the port of the switch. | [optional] 
**SupportVrf** | Pointer to **bool** | Whether the switch supports VRF | [optional] 
**Version** | Pointer to **string** | Switch firmwareVersion | [optional] 

## Methods

### NewDot1xSwitchInfoOpenApiVO

`func NewDot1xSwitchInfoOpenApiVO() *Dot1xSwitchInfoOpenApiVO`

NewDot1xSwitchInfoOpenApiVO instantiates a new Dot1xSwitchInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDot1xSwitchInfoOpenApiVOWithDefaults

`func NewDot1xSwitchInfoOpenApiVOWithDefaults() *Dot1xSwitchInfoOpenApiVO`

NewDot1xSwitchInfoOpenApiVOWithDefaults instantiates a new Dot1xSwitchInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountVrfId

`func (o *Dot1xSwitchInfoOpenApiVO) GetAccountVrfId() string`

GetAccountVrfId returns the AccountVrfId field if non-nil, zero value otherwise.

### GetAccountVrfIdOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetAccountVrfIdOk() (*string, bool)`

GetAccountVrfIdOk returns a tuple with the AccountVrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountVrfId

`func (o *Dot1xSwitchInfoOpenApiVO) SetAccountVrfId(v string)`

SetAccountVrfId sets AccountVrfId field to given value.

### HasAccountVrfId

`func (o *Dot1xSwitchInfoOpenApiVO) HasAccountVrfId() bool`

HasAccountVrfId returns a boolean if a field has been set.

### GetAddedInAdvanced

`func (o *Dot1xSwitchInfoOpenApiVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *Dot1xSwitchInfoOpenApiVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *Dot1xSwitchInfoOpenApiVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetAuthVrfId

`func (o *Dot1xSwitchInfoOpenApiVO) GetAuthVrfId() string`

GetAuthVrfId returns the AuthVrfId field if non-nil, zero value otherwise.

### GetAuthVrfIdOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetAuthVrfIdOk() (*string, bool)`

GetAuthVrfIdOk returns a tuple with the AuthVrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthVrfId

`func (o *Dot1xSwitchInfoOpenApiVO) SetAuthVrfId(v string)`

SetAuthVrfId sets AuthVrfId field to given value.

### HasAuthVrfId

`func (o *Dot1xSwitchInfoOpenApiVO) HasAuthVrfId() bool`

HasAuthVrfId returns a boolean if a field has been set.

### GetMac

`func (o *Dot1xSwitchInfoOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Dot1xSwitchInfoOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Dot1xSwitchInfoOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *Dot1xSwitchInfoOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Dot1xSwitchInfoOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Dot1xSwitchInfoOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *Dot1xSwitchInfoOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dot1xSwitchInfoOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dot1xSwitchInfoOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortConfigured

`func (o *Dot1xSwitchInfoOpenApiVO) GetPortConfigured() bool`

GetPortConfigured returns the PortConfigured field if non-nil, zero value otherwise.

### GetPortConfiguredOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetPortConfiguredOk() (*bool, bool)`

GetPortConfiguredOk returns a tuple with the PortConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfigured

`func (o *Dot1xSwitchInfoOpenApiVO) SetPortConfigured(v bool)`

SetPortConfigured sets PortConfigured field to given value.

### HasPortConfigured

`func (o *Dot1xSwitchInfoOpenApiVO) HasPortConfigured() bool`

HasPortConfigured returns a boolean if a field has been set.

### GetPorts

`func (o *Dot1xSwitchInfoOpenApiVO) GetPorts() []Dot1xPortInfoOpenApiVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetPortsOk() (*[]Dot1xPortInfoOpenApiVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Dot1xSwitchInfoOpenApiVO) SetPorts(v []Dot1xPortInfoOpenApiVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *Dot1xSwitchInfoOpenApiVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetStatus

`func (o *Dot1xSwitchInfoOpenApiVO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Dot1xSwitchInfoOpenApiVO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Dot1xSwitchInfoOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *Dot1xSwitchInfoOpenApiVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *Dot1xSwitchInfoOpenApiVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *Dot1xSwitchInfoOpenApiVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetSupportSingleMabAuth

`func (o *Dot1xSwitchInfoOpenApiVO) GetSupportSingleMabAuth() bool`

GetSupportSingleMabAuth returns the SupportSingleMabAuth field if non-nil, zero value otherwise.

### GetSupportSingleMabAuthOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetSupportSingleMabAuthOk() (*bool, bool)`

GetSupportSingleMabAuthOk returns a tuple with the SupportSingleMabAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSingleMabAuth

`func (o *Dot1xSwitchInfoOpenApiVO) SetSupportSingleMabAuth(v bool)`

SetSupportSingleMabAuth sets SupportSingleMabAuth field to given value.

### HasSupportSingleMabAuth

`func (o *Dot1xSwitchInfoOpenApiVO) HasSupportSingleMabAuth() bool`

HasSupportSingleMabAuth returns a boolean if a field has been set.

### GetSupportVrf

`func (o *Dot1xSwitchInfoOpenApiVO) GetSupportVrf() bool`

GetSupportVrf returns the SupportVrf field if non-nil, zero value otherwise.

### GetSupportVrfOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetSupportVrfOk() (*bool, bool)`

GetSupportVrfOk returns a tuple with the SupportVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVrf

`func (o *Dot1xSwitchInfoOpenApiVO) SetSupportVrf(v bool)`

SetSupportVrf sets SupportVrf field to given value.

### HasSupportVrf

`func (o *Dot1xSwitchInfoOpenApiVO) HasSupportVrf() bool`

HasSupportVrf returns a boolean if a field has been set.

### GetVersion

`func (o *Dot1xSwitchInfoOpenApiVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Dot1xSwitchInfoOpenApiVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Dot1xSwitchInfoOpenApiVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Dot1xSwitchInfoOpenApiVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


