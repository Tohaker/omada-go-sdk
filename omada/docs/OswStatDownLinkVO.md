# OswStatDownLinkVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientType** | Pointer to **string** | Type of down-link client | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Model of device | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected; 1:Connected; 2:Pending; 3:Heartbeat Missed; 4:Isolated | [optional] 
**Type** | Pointer to **string** | Type of down-link device | [optional] 

## Methods

### NewOswStatDownLinkVO

`func NewOswStatDownLinkVO() *OswStatDownLinkVO`

NewOswStatDownLinkVO instantiates a new OswStatDownLinkVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStatDownLinkVOWithDefaults

`func NewOswStatDownLinkVOWithDefaults() *OswStatDownLinkVO`

NewOswStatDownLinkVOWithDefaults instantiates a new OswStatDownLinkVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientType

`func (o *OswStatDownLinkVO) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *OswStatDownLinkVO) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *OswStatDownLinkVO) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *OswStatDownLinkVO) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetMac

`func (o *OswStatDownLinkVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswStatDownLinkVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswStatDownLinkVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswStatDownLinkVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *OswStatDownLinkVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswStatDownLinkVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswStatDownLinkVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswStatDownLinkVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswStatDownLinkVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswStatDownLinkVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswStatDownLinkVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswStatDownLinkVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *OswStatDownLinkVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStatDownLinkVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStatDownLinkVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswStatDownLinkVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatusCategory

`func (o *OswStatDownLinkVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *OswStatDownLinkVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *OswStatDownLinkVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *OswStatDownLinkVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetType

`func (o *OswStatDownLinkVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswStatDownLinkVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswStatDownLinkVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OswStatDownLinkVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


