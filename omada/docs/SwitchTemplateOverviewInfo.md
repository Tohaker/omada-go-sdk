# SwitchTemplateOverviewInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HwVersion** | Pointer to **string** | Hardware Version | [optional] 
**Id** | Pointer to **string** | Device templateId | [optional] 
**Model** | Pointer to **string** | Model | [optional] 
**PortList** | Pointer to [**[]PortInfo**](PortInfo.md) | Port List | [optional] 

## Methods

### NewSwitchTemplateOverviewInfo

`func NewSwitchTemplateOverviewInfo() *SwitchTemplateOverviewInfo`

NewSwitchTemplateOverviewInfo instantiates a new SwitchTemplateOverviewInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchTemplateOverviewInfoWithDefaults

`func NewSwitchTemplateOverviewInfoWithDefaults() *SwitchTemplateOverviewInfo`

NewSwitchTemplateOverviewInfoWithDefaults instantiates a new SwitchTemplateOverviewInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHwVersion

`func (o *SwitchTemplateOverviewInfo) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *SwitchTemplateOverviewInfo) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *SwitchTemplateOverviewInfo) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *SwitchTemplateOverviewInfo) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetId

`func (o *SwitchTemplateOverviewInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SwitchTemplateOverviewInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SwitchTemplateOverviewInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SwitchTemplateOverviewInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *SwitchTemplateOverviewInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SwitchTemplateOverviewInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SwitchTemplateOverviewInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SwitchTemplateOverviewInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPortList

`func (o *SwitchTemplateOverviewInfo) GetPortList() []PortInfo`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *SwitchTemplateOverviewInfo) GetPortListOk() (*[]PortInfo, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *SwitchTemplateOverviewInfo) SetPortList(v []PortInfo)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *SwitchTemplateOverviewInfo) HasPortList() bool`

HasPortList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


