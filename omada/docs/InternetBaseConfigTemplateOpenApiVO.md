# InternetBaseConfigTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **int32** | Online detection interval(second). 0 means disable. It should be within the range of 0–3600. | [optional] 
**WanPortList** | **[]string** | List of enabled WAN port IDs, the valid port IDs can be obtained from \&quot;Get site template internet basic info\&quot;. | 

## Methods

### NewInternetBaseConfigTemplateOpenApiVO

`func NewInternetBaseConfigTemplateOpenApiVO(wanPortList []string, ) *InternetBaseConfigTemplateOpenApiVO`

NewInternetBaseConfigTemplateOpenApiVO instantiates a new InternetBaseConfigTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetBaseConfigTemplateOpenApiVOWithDefaults

`func NewInternetBaseConfigTemplateOpenApiVOWithDefaults() *InternetBaseConfigTemplateOpenApiVO`

NewInternetBaseConfigTemplateOpenApiVOWithDefaults instantiates a new InternetBaseConfigTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *InternetBaseConfigTemplateOpenApiVO) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InternetBaseConfigTemplateOpenApiVO) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InternetBaseConfigTemplateOpenApiVO) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *InternetBaseConfigTemplateOpenApiVO) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetWanPortList

`func (o *InternetBaseConfigTemplateOpenApiVO) GetWanPortList() []string`

GetWanPortList returns the WanPortList field if non-nil, zero value otherwise.

### GetWanPortListOk

`func (o *InternetBaseConfigTemplateOpenApiVO) GetWanPortListOk() (*[]string, bool)`

GetWanPortListOk returns a tuple with the WanPortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortList

`func (o *InternetBaseConfigTemplateOpenApiVO) SetWanPortList(v []string)`

SetWanPortList sets WanPortList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


