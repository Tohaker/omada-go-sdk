# CreateRadiusProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctServer** | Pointer to [**[]RadiusAcctServerOpenApiVO**](RadiusAcctServerOpenApiVO.md) | Radius accounting server list, valid when parameter [radiusAccountingEnable] is true | [optional] 
**AuthServer** | [**[]RadiusAuthServerOpenApiVO**](RadiusAuthServerOpenApiVO.md) | Radius authentication server list | 
**CoaEnable** | Pointer to **bool** | Radius CoA enable status | [optional] 
**CoaPassword** | Pointer to **string** | Radius CoA password, required when parameter [coaEnable] is true. CoaPassword should contain 1 to 128 characters. The question mark (?), double quote (\&quot;), percent sign (%), and backslash (\\) may cause the RADIUS function to fail and are not recommended. | [optional] 
**InterimUpdateEnable** | Pointer to **bool** | When radius accounting enables, interval update enable status | [optional] 
**InterimUpdateInterval** | Pointer to **int32** | When interval update enables, interval update duration, unit: second. InterimUpdateInterval should be within the range of 60-86400 | [optional] 
**Name** | **string** | Radius profile name should contain 1 to 64 characters | 
**RadiusAccountingEnable** | **bool** | Radius accounting enable status | 
**RequireMessageAuthenticator** | Pointer to **bool** | Message-Authenticator enable status | [optional] 
**WirelessVlanAssignment** | **bool** | VLAN assignment for wireless network enable status | 

## Methods

### NewCreateRadiusProfileOpenApiVO

`func NewCreateRadiusProfileOpenApiVO(authServer []RadiusAuthServerOpenApiVO, name string, radiusAccountingEnable bool, wirelessVlanAssignment bool, ) *CreateRadiusProfileOpenApiVO`

NewCreateRadiusProfileOpenApiVO instantiates a new CreateRadiusProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRadiusProfileOpenApiVOWithDefaults

`func NewCreateRadiusProfileOpenApiVOWithDefaults() *CreateRadiusProfileOpenApiVO`

NewCreateRadiusProfileOpenApiVOWithDefaults instantiates a new CreateRadiusProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctServer

`func (o *CreateRadiusProfileOpenApiVO) GetAcctServer() []RadiusAcctServerOpenApiVO`

GetAcctServer returns the AcctServer field if non-nil, zero value otherwise.

### GetAcctServerOk

`func (o *CreateRadiusProfileOpenApiVO) GetAcctServerOk() (*[]RadiusAcctServerOpenApiVO, bool)`

GetAcctServerOk returns a tuple with the AcctServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctServer

`func (o *CreateRadiusProfileOpenApiVO) SetAcctServer(v []RadiusAcctServerOpenApiVO)`

SetAcctServer sets AcctServer field to given value.

### HasAcctServer

`func (o *CreateRadiusProfileOpenApiVO) HasAcctServer() bool`

HasAcctServer returns a boolean if a field has been set.

### GetAuthServer

`func (o *CreateRadiusProfileOpenApiVO) GetAuthServer() []RadiusAuthServerOpenApiVO`

GetAuthServer returns the AuthServer field if non-nil, zero value otherwise.

### GetAuthServerOk

`func (o *CreateRadiusProfileOpenApiVO) GetAuthServerOk() (*[]RadiusAuthServerOpenApiVO, bool)`

GetAuthServerOk returns a tuple with the AuthServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServer

`func (o *CreateRadiusProfileOpenApiVO) SetAuthServer(v []RadiusAuthServerOpenApiVO)`

SetAuthServer sets AuthServer field to given value.


### GetCoaEnable

`func (o *CreateRadiusProfileOpenApiVO) GetCoaEnable() bool`

GetCoaEnable returns the CoaEnable field if non-nil, zero value otherwise.

### GetCoaEnableOk

`func (o *CreateRadiusProfileOpenApiVO) GetCoaEnableOk() (*bool, bool)`

GetCoaEnableOk returns a tuple with the CoaEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoaEnable

`func (o *CreateRadiusProfileOpenApiVO) SetCoaEnable(v bool)`

SetCoaEnable sets CoaEnable field to given value.

### HasCoaEnable

`func (o *CreateRadiusProfileOpenApiVO) HasCoaEnable() bool`

HasCoaEnable returns a boolean if a field has been set.

### GetCoaPassword

`func (o *CreateRadiusProfileOpenApiVO) GetCoaPassword() string`

GetCoaPassword returns the CoaPassword field if non-nil, zero value otherwise.

### GetCoaPasswordOk

`func (o *CreateRadiusProfileOpenApiVO) GetCoaPasswordOk() (*string, bool)`

GetCoaPasswordOk returns a tuple with the CoaPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoaPassword

`func (o *CreateRadiusProfileOpenApiVO) SetCoaPassword(v string)`

SetCoaPassword sets CoaPassword field to given value.

### HasCoaPassword

`func (o *CreateRadiusProfileOpenApiVO) HasCoaPassword() bool`

HasCoaPassword returns a boolean if a field has been set.

### GetInterimUpdateEnable

`func (o *CreateRadiusProfileOpenApiVO) GetInterimUpdateEnable() bool`

GetInterimUpdateEnable returns the InterimUpdateEnable field if non-nil, zero value otherwise.

### GetInterimUpdateEnableOk

`func (o *CreateRadiusProfileOpenApiVO) GetInterimUpdateEnableOk() (*bool, bool)`

GetInterimUpdateEnableOk returns a tuple with the InterimUpdateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterimUpdateEnable

`func (o *CreateRadiusProfileOpenApiVO) SetInterimUpdateEnable(v bool)`

SetInterimUpdateEnable sets InterimUpdateEnable field to given value.

### HasInterimUpdateEnable

`func (o *CreateRadiusProfileOpenApiVO) HasInterimUpdateEnable() bool`

HasInterimUpdateEnable returns a boolean if a field has been set.

### GetInterimUpdateInterval

`func (o *CreateRadiusProfileOpenApiVO) GetInterimUpdateInterval() int32`

GetInterimUpdateInterval returns the InterimUpdateInterval field if non-nil, zero value otherwise.

### GetInterimUpdateIntervalOk

`func (o *CreateRadiusProfileOpenApiVO) GetInterimUpdateIntervalOk() (*int32, bool)`

GetInterimUpdateIntervalOk returns a tuple with the InterimUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterimUpdateInterval

`func (o *CreateRadiusProfileOpenApiVO) SetInterimUpdateInterval(v int32)`

SetInterimUpdateInterval sets InterimUpdateInterval field to given value.

### HasInterimUpdateInterval

`func (o *CreateRadiusProfileOpenApiVO) HasInterimUpdateInterval() bool`

HasInterimUpdateInterval returns a boolean if a field has been set.

### GetName

`func (o *CreateRadiusProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRadiusProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRadiusProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetRadiusAccountingEnable

`func (o *CreateRadiusProfileOpenApiVO) GetRadiusAccountingEnable() bool`

GetRadiusAccountingEnable returns the RadiusAccountingEnable field if non-nil, zero value otherwise.

### GetRadiusAccountingEnableOk

`func (o *CreateRadiusProfileOpenApiVO) GetRadiusAccountingEnableOk() (*bool, bool)`

GetRadiusAccountingEnableOk returns a tuple with the RadiusAccountingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnable

`func (o *CreateRadiusProfileOpenApiVO) SetRadiusAccountingEnable(v bool)`

SetRadiusAccountingEnable sets RadiusAccountingEnable field to given value.


### GetRequireMessageAuthenticator

`func (o *CreateRadiusProfileOpenApiVO) GetRequireMessageAuthenticator() bool`

GetRequireMessageAuthenticator returns the RequireMessageAuthenticator field if non-nil, zero value otherwise.

### GetRequireMessageAuthenticatorOk

`func (o *CreateRadiusProfileOpenApiVO) GetRequireMessageAuthenticatorOk() (*bool, bool)`

GetRequireMessageAuthenticatorOk returns a tuple with the RequireMessageAuthenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMessageAuthenticator

`func (o *CreateRadiusProfileOpenApiVO) SetRequireMessageAuthenticator(v bool)`

SetRequireMessageAuthenticator sets RequireMessageAuthenticator field to given value.

### HasRequireMessageAuthenticator

`func (o *CreateRadiusProfileOpenApiVO) HasRequireMessageAuthenticator() bool`

HasRequireMessageAuthenticator returns a boolean if a field has been set.

### GetWirelessVlanAssignment

`func (o *CreateRadiusProfileOpenApiVO) GetWirelessVlanAssignment() bool`

GetWirelessVlanAssignment returns the WirelessVlanAssignment field if non-nil, zero value otherwise.

### GetWirelessVlanAssignmentOk

`func (o *CreateRadiusProfileOpenApiVO) GetWirelessVlanAssignmentOk() (*bool, bool)`

GetWirelessVlanAssignmentOk returns a tuple with the WirelessVlanAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessVlanAssignment

`func (o *CreateRadiusProfileOpenApiVO) SetWirelessVlanAssignment(v bool)`

SetWirelessVlanAssignment sets WirelessVlanAssignment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


