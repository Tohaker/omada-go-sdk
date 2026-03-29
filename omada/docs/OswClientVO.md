# OswClientVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **string** | Device Type: iPhone, iPod, Android, PC, printer, TV... | [optional] 
**Ip** | Pointer to **string** | IP Address | [optional] 
**LagId** | Pointer to **int32** | (Wired) LAG ID. Exists only when the client is connected to the LAG | [optional] 
**Mac** | Pointer to **string** | Client MAC Address | [optional] 
**Model** | Pointer to **string** | Client Model | [optional] 
**Name** | Pointer to **string** | Client Name, alias | [optional] 
**Port** | Pointer to **int32** | (Wired) Port ID | [optional] 
**StandardPort** | Pointer to **string** | Standard port | [optional] 
**Unit** | Pointer to **int32** | Unit ID | [optional] 

## Methods

### NewOswClientVO

`func NewOswClientVO() *OswClientVO`

NewOswClientVO instantiates a new OswClientVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswClientVOWithDefaults

`func NewOswClientVOWithDefaults() *OswClientVO`

NewOswClientVOWithDefaults instantiates a new OswClientVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *OswClientVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *OswClientVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *OswClientVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *OswClientVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetIp

`func (o *OswClientVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswClientVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswClientVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswClientVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLagId

`func (o *OswClientVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *OswClientVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *OswClientVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *OswClientVO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetMac

`func (o *OswClientVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswClientVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswClientVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswClientVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *OswClientVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswClientVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswClientVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswClientVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *OswClientVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswClientVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswClientVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswClientVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *OswClientVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswClientVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswClientVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswClientVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswClientVO) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswClientVO) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswClientVO) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswClientVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetUnit

`func (o *OswClientVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswClientVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswClientVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswClientVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


