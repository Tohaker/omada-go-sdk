# ExcludedAPsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delete** | **bool** | Whether to delete device(s) from the excluded AP list. | 
**ExcludeAps** | Pointer to **string** | Parameter [excludeAps] should not be null when parameter [delete] is true. It contains device MAC that should not be excluded. MAC should be concatenated with &#39;,&#39;.  | [optional] 
**SelectMacs** | Pointer to [**SelectMacsVO**](SelectMacsVO.md) |  | [optional] 

## Methods

### NewExcludedAPsConfig

`func NewExcludedAPsConfig(delete bool, ) *ExcludedAPsConfig`

NewExcludedAPsConfig instantiates a new ExcludedAPsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExcludedAPsConfigWithDefaults

`func NewExcludedAPsConfigWithDefaults() *ExcludedAPsConfig`

NewExcludedAPsConfigWithDefaults instantiates a new ExcludedAPsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelete

`func (o *ExcludedAPsConfig) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ExcludedAPsConfig) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ExcludedAPsConfig) SetDelete(v bool)`

SetDelete sets Delete field to given value.


### GetExcludeAps

`func (o *ExcludedAPsConfig) GetExcludeAps() string`

GetExcludeAps returns the ExcludeAps field if non-nil, zero value otherwise.

### GetExcludeApsOk

`func (o *ExcludedAPsConfig) GetExcludeApsOk() (*string, bool)`

GetExcludeApsOk returns a tuple with the ExcludeAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAps

`func (o *ExcludedAPsConfig) SetExcludeAps(v string)`

SetExcludeAps sets ExcludeAps field to given value.

### HasExcludeAps

`func (o *ExcludedAPsConfig) HasExcludeAps() bool`

HasExcludeAps returns a boolean if a field has been set.

### GetSelectMacs

`func (o *ExcludedAPsConfig) GetSelectMacs() SelectMacsVO`

GetSelectMacs returns the SelectMacs field if non-nil, zero value otherwise.

### GetSelectMacsOk

`func (o *ExcludedAPsConfig) GetSelectMacsOk() (*SelectMacsVO, bool)`

GetSelectMacsOk returns a tuple with the SelectMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectMacs

`func (o *ExcludedAPsConfig) SetSelectMacs(v SelectMacsVO)`

SetSelectMacs sets SelectMacs field to given value.

### HasSelectMacs

`func (o *ExcludedAPsConfig) HasSelectMacs() bool`

HasSelectMacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


