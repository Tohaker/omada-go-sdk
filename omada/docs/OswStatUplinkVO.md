# OswStatUplinkVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Model of device | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**Port** | Pointer to **int32** | Port ID | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected; 1:Connected; 2:Pending; 3:Heartbeat Missed; 4:Isolated | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 

## Methods

### NewOswStatUplinkVO

`func NewOswStatUplinkVO() *OswStatUplinkVO`

NewOswStatUplinkVO instantiates a new OswStatUplinkVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStatUplinkVOWithDefaults

`func NewOswStatUplinkVOWithDefaults() *OswStatUplinkVO`

NewOswStatUplinkVOWithDefaults instantiates a new OswStatUplinkVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *OswStatUplinkVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswStatUplinkVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswStatUplinkVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswStatUplinkVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *OswStatUplinkVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswStatUplinkVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswStatUplinkVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswStatUplinkVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswStatUplinkVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswStatUplinkVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswStatUplinkVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswStatUplinkVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *OswStatUplinkVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStatUplinkVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStatUplinkVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswStatUplinkVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *OswStatUplinkVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswStatUplinkVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswStatUplinkVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswStatUplinkVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatusCategory

`func (o *OswStatUplinkVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *OswStatUplinkVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *OswStatUplinkVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *OswStatUplinkVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetType

`func (o *OswStatUplinkVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswStatUplinkVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswStatUplinkVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OswStatUplinkVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


