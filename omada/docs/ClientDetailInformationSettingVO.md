# ClientDetailInformationSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientDataTrendEnable** | Pointer to **bool** | Whether client data trend record is enabled.When enabled, client trend statistics and charts will be retained, which will take up lots of storage space. | [optional] 
**ClientHealthEnable** | Pointer to **bool** | Whether client health is enabled. When enabled, client health data will be recorded, which may consume a significant amount of storage space. | [optional] 
**ClientHistoryEnable** | Pointer to **bool** | Whether client history is enabled. When enabled, client history, client logs will be recorded. This will occupy much storage space. | [optional] 
**ClientRecognitionEnable** | Pointer to **bool** | Whether client recognition is enabled. With the feature enabled, network devices will report client information in real time to ensure the accuracy of client recognition. Cloud Access is required for client recognition. | [optional] 
**FollowMsp** | Pointer to **bool** | Whether customer follow MSP setting, only effective in MSP mode and only effective for [clientHistoryEnable] and [clientDataTrendEnable]. | [optional] 

## Methods

### NewClientDetailInformationSettingVO

`func NewClientDetailInformationSettingVO() *ClientDetailInformationSettingVO`

NewClientDetailInformationSettingVO instantiates a new ClientDetailInformationSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDetailInformationSettingVOWithDefaults

`func NewClientDetailInformationSettingVOWithDefaults() *ClientDetailInformationSettingVO`

NewClientDetailInformationSettingVOWithDefaults instantiates a new ClientDetailInformationSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientDataTrendEnable

`func (o *ClientDetailInformationSettingVO) GetClientDataTrendEnable() bool`

GetClientDataTrendEnable returns the ClientDataTrendEnable field if non-nil, zero value otherwise.

### GetClientDataTrendEnableOk

`func (o *ClientDetailInformationSettingVO) GetClientDataTrendEnableOk() (*bool, bool)`

GetClientDataTrendEnableOk returns a tuple with the ClientDataTrendEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDataTrendEnable

`func (o *ClientDetailInformationSettingVO) SetClientDataTrendEnable(v bool)`

SetClientDataTrendEnable sets ClientDataTrendEnable field to given value.

### HasClientDataTrendEnable

`func (o *ClientDetailInformationSettingVO) HasClientDataTrendEnable() bool`

HasClientDataTrendEnable returns a boolean if a field has been set.

### GetClientHealthEnable

`func (o *ClientDetailInformationSettingVO) GetClientHealthEnable() bool`

GetClientHealthEnable returns the ClientHealthEnable field if non-nil, zero value otherwise.

### GetClientHealthEnableOk

`func (o *ClientDetailInformationSettingVO) GetClientHealthEnableOk() (*bool, bool)`

GetClientHealthEnableOk returns a tuple with the ClientHealthEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHealthEnable

`func (o *ClientDetailInformationSettingVO) SetClientHealthEnable(v bool)`

SetClientHealthEnable sets ClientHealthEnable field to given value.

### HasClientHealthEnable

`func (o *ClientDetailInformationSettingVO) HasClientHealthEnable() bool`

HasClientHealthEnable returns a boolean if a field has been set.

### GetClientHistoryEnable

`func (o *ClientDetailInformationSettingVO) GetClientHistoryEnable() bool`

GetClientHistoryEnable returns the ClientHistoryEnable field if non-nil, zero value otherwise.

### GetClientHistoryEnableOk

`func (o *ClientDetailInformationSettingVO) GetClientHistoryEnableOk() (*bool, bool)`

GetClientHistoryEnableOk returns a tuple with the ClientHistoryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHistoryEnable

`func (o *ClientDetailInformationSettingVO) SetClientHistoryEnable(v bool)`

SetClientHistoryEnable sets ClientHistoryEnable field to given value.

### HasClientHistoryEnable

`func (o *ClientDetailInformationSettingVO) HasClientHistoryEnable() bool`

HasClientHistoryEnable returns a boolean if a field has been set.

### GetClientRecognitionEnable

`func (o *ClientDetailInformationSettingVO) GetClientRecognitionEnable() bool`

GetClientRecognitionEnable returns the ClientRecognitionEnable field if non-nil, zero value otherwise.

### GetClientRecognitionEnableOk

`func (o *ClientDetailInformationSettingVO) GetClientRecognitionEnableOk() (*bool, bool)`

GetClientRecognitionEnableOk returns a tuple with the ClientRecognitionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRecognitionEnable

`func (o *ClientDetailInformationSettingVO) SetClientRecognitionEnable(v bool)`

SetClientRecognitionEnable sets ClientRecognitionEnable field to given value.

### HasClientRecognitionEnable

`func (o *ClientDetailInformationSettingVO) HasClientRecognitionEnable() bool`

HasClientRecognitionEnable returns a boolean if a field has been set.

### GetFollowMsp

`func (o *ClientDetailInformationSettingVO) GetFollowMsp() bool`

GetFollowMsp returns the FollowMsp field if non-nil, zero value otherwise.

### GetFollowMspOk

`func (o *ClientDetailInformationSettingVO) GetFollowMspOk() (*bool, bool)`

GetFollowMspOk returns a tuple with the FollowMsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowMsp

`func (o *ClientDetailInformationSettingVO) SetFollowMsp(v bool)`

SetFollowMsp sets FollowMsp field to given value.

### HasFollowMsp

`func (o *ClientDetailInformationSettingVO) HasFollowMsp() bool`

HasFollowMsp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


