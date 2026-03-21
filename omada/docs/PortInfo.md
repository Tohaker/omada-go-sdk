# PortInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagPort** | **bool** | Whether this port exists in a LAG | 
**Name** | **string** | Port name | 
**PoeMode** | **int32** | PoE mode should be a value as follows: 1: on(802.3at/af); 0: off. | 
**Port** | **int32** | Port ID | 
**ProfileId** | **string** | Profile ID | 
**ProfileName** | **string** | Profile Name | 
**ProfileOverrideEnable** | **bool** | Profile Override Enable | 
**Status** | **int32** | Status should be a value as follows: 0: off; 1: on, only when lagPort is false | 

## Methods

### NewPortInfo

`func NewPortInfo(lagPort bool, name string, poeMode int32, port int32, profileId string, profileName string, profileOverrideEnable bool, status int32, ) *PortInfo`

NewPortInfo instantiates a new PortInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortInfoWithDefaults

`func NewPortInfoWithDefaults() *PortInfo`

NewPortInfoWithDefaults instantiates a new PortInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagPort

`func (o *PortInfo) GetLagPort() bool`

GetLagPort returns the LagPort field if non-nil, zero value otherwise.

### GetLagPortOk

`func (o *PortInfo) GetLagPortOk() (*bool, bool)`

GetLagPortOk returns a tuple with the LagPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagPort

`func (o *PortInfo) SetLagPort(v bool)`

SetLagPort sets LagPort field to given value.


### GetName

`func (o *PortInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortInfo) SetName(v string)`

SetName sets Name field to given value.


### GetPoeMode

`func (o *PortInfo) GetPoeMode() int32`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *PortInfo) GetPoeModeOk() (*int32, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *PortInfo) SetPoeMode(v int32)`

SetPoeMode sets PoeMode field to given value.


### GetPort

`func (o *PortInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PortInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PortInfo) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProfileId

`func (o *PortInfo) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *PortInfo) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *PortInfo) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetProfileName

`func (o *PortInfo) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *PortInfo) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *PortInfo) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.


### GetProfileOverrideEnable

`func (o *PortInfo) GetProfileOverrideEnable() bool`

GetProfileOverrideEnable returns the ProfileOverrideEnable field if non-nil, zero value otherwise.

### GetProfileOverrideEnableOk

`func (o *PortInfo) GetProfileOverrideEnableOk() (*bool, bool)`

GetProfileOverrideEnableOk returns a tuple with the ProfileOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOverrideEnable

`func (o *PortInfo) SetProfileOverrideEnable(v bool)`

SetProfileOverrideEnable sets ProfileOverrideEnable field to given value.


### GetStatus

`func (o *PortInfo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortInfo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortInfo) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


