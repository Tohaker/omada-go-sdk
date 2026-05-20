# SdWanNatInfoConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomMapNetworkList** | Pointer to **[]string** | A list of the customized mapping network | [optional] 
**DefaultMapNetworkList** | Pointer to **[]string** | A list of the default mapping network | [optional] 
**NetworkMapList** | Pointer to [**[]SdWanNatItemConfig**](SdWanNatItemConfig.md) | A list of the network map item | [optional] 

## Methods

### NewSdWanNatInfoConfig

`func NewSdWanNatInfoConfig() *SdWanNatInfoConfig`

NewSdWanNatInfoConfig instantiates a new SdWanNatInfoConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanNatInfoConfigWithDefaults

`func NewSdWanNatInfoConfigWithDefaults() *SdWanNatInfoConfig`

NewSdWanNatInfoConfigWithDefaults instantiates a new SdWanNatInfoConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomMapNetworkList

`func (o *SdWanNatInfoConfig) GetCustomMapNetworkList() []string`

GetCustomMapNetworkList returns the CustomMapNetworkList field if non-nil, zero value otherwise.

### GetCustomMapNetworkListOk

`func (o *SdWanNatInfoConfig) GetCustomMapNetworkListOk() (*[]string, bool)`

GetCustomMapNetworkListOk returns a tuple with the CustomMapNetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMapNetworkList

`func (o *SdWanNatInfoConfig) SetCustomMapNetworkList(v []string)`

SetCustomMapNetworkList sets CustomMapNetworkList field to given value.

### HasCustomMapNetworkList

`func (o *SdWanNatInfoConfig) HasCustomMapNetworkList() bool`

HasCustomMapNetworkList returns a boolean if a field has been set.

### GetDefaultMapNetworkList

`func (o *SdWanNatInfoConfig) GetDefaultMapNetworkList() []string`

GetDefaultMapNetworkList returns the DefaultMapNetworkList field if non-nil, zero value otherwise.

### GetDefaultMapNetworkListOk

`func (o *SdWanNatInfoConfig) GetDefaultMapNetworkListOk() (*[]string, bool)`

GetDefaultMapNetworkListOk returns a tuple with the DefaultMapNetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMapNetworkList

`func (o *SdWanNatInfoConfig) SetDefaultMapNetworkList(v []string)`

SetDefaultMapNetworkList sets DefaultMapNetworkList field to given value.

### HasDefaultMapNetworkList

`func (o *SdWanNatInfoConfig) HasDefaultMapNetworkList() bool`

HasDefaultMapNetworkList returns a boolean if a field has been set.

### GetNetworkMapList

`func (o *SdWanNatInfoConfig) GetNetworkMapList() []SdWanNatItemConfig`

GetNetworkMapList returns the NetworkMapList field if non-nil, zero value otherwise.

### GetNetworkMapListOk

`func (o *SdWanNatInfoConfig) GetNetworkMapListOk() (*[]SdWanNatItemConfig, bool)`

GetNetworkMapListOk returns a tuple with the NetworkMapList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMapList

`func (o *SdWanNatInfoConfig) SetNetworkMapList(v []SdWanNatItemConfig)`

SetNetworkMapList sets NetworkMapList field to given value.

### HasNetworkMapList

`func (o *SdWanNatInfoConfig) HasNetworkMapList() bool`

HasNetworkMapList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


