# Dot1xEapInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | EAP MAC address | [optional] 
**Model** | Pointer to **string** | EAP model | [optional] 
**Name** | Pointer to **string** | EAP name | [optional] 
**Ports** | Pointer to [**[]Dot1xEapPortInfoOpenApiVO**](Dot1xEapPortInfoOpenApiVO.md) | EAP port information | [optional] 
**Status** | Pointer to **string** | Device status | [optional] 
**StatusCategory** | Pointer to **int32** | Device status category, 0: Disconnected, 1: Connected, 2: Pending,3: Heartbeat Missed, 4: Isolated | [optional] 
**Version** | Pointer to **string** | EAP firmwareVersion | [optional] 

## Methods

### NewDot1xEapInfoOpenApiVO

`func NewDot1xEapInfoOpenApiVO() *Dot1xEapInfoOpenApiVO`

NewDot1xEapInfoOpenApiVO instantiates a new Dot1xEapInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDot1xEapInfoOpenApiVOWithDefaults

`func NewDot1xEapInfoOpenApiVOWithDefaults() *Dot1xEapInfoOpenApiVO`

NewDot1xEapInfoOpenApiVOWithDefaults instantiates a new Dot1xEapInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *Dot1xEapInfoOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Dot1xEapInfoOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Dot1xEapInfoOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Dot1xEapInfoOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *Dot1xEapInfoOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Dot1xEapInfoOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Dot1xEapInfoOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Dot1xEapInfoOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *Dot1xEapInfoOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dot1xEapInfoOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dot1xEapInfoOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dot1xEapInfoOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *Dot1xEapInfoOpenApiVO) GetPorts() []Dot1xEapPortInfoOpenApiVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Dot1xEapInfoOpenApiVO) GetPortsOk() (*[]Dot1xEapPortInfoOpenApiVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Dot1xEapInfoOpenApiVO) SetPorts(v []Dot1xEapPortInfoOpenApiVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *Dot1xEapInfoOpenApiVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetStatus

`func (o *Dot1xEapInfoOpenApiVO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Dot1xEapInfoOpenApiVO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Dot1xEapInfoOpenApiVO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Dot1xEapInfoOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *Dot1xEapInfoOpenApiVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *Dot1xEapInfoOpenApiVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *Dot1xEapInfoOpenApiVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *Dot1xEapInfoOpenApiVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetVersion

`func (o *Dot1xEapInfoOpenApiVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Dot1xEapInfoOpenApiVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Dot1xEapInfoOpenApiVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Dot1xEapInfoOpenApiVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


