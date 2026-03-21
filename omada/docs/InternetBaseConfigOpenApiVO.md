# InternetBaseConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **int32** | Online detection interval(second). 0 means disable. It should be within the range of 0–3600. | [optional] 
**PreConfiguration** | **bool** | All port and port-related(like ACL) configuration will take effect only when the parameter [preConfiguration] is true. The value will always be true when a gateway is in this site. | 
**WanPortList** | **[]string** | List of enabled WAN port IDs, the valid port IDs can be obtained from \&quot;Get internet basic info\&quot;. | 

## Methods

### NewInternetBaseConfigOpenApiVO

`func NewInternetBaseConfigOpenApiVO(preConfiguration bool, wanPortList []string, ) *InternetBaseConfigOpenApiVO`

NewInternetBaseConfigOpenApiVO instantiates a new InternetBaseConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetBaseConfigOpenApiVOWithDefaults

`func NewInternetBaseConfigOpenApiVOWithDefaults() *InternetBaseConfigOpenApiVO`

NewInternetBaseConfigOpenApiVOWithDefaults instantiates a new InternetBaseConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *InternetBaseConfigOpenApiVO) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InternetBaseConfigOpenApiVO) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InternetBaseConfigOpenApiVO) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *InternetBaseConfigOpenApiVO) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetPreConfiguration

`func (o *InternetBaseConfigOpenApiVO) GetPreConfiguration() bool`

GetPreConfiguration returns the PreConfiguration field if non-nil, zero value otherwise.

### GetPreConfigurationOk

`func (o *InternetBaseConfigOpenApiVO) GetPreConfigurationOk() (*bool, bool)`

GetPreConfigurationOk returns a tuple with the PreConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConfiguration

`func (o *InternetBaseConfigOpenApiVO) SetPreConfiguration(v bool)`

SetPreConfiguration sets PreConfiguration field to given value.


### GetWanPortList

`func (o *InternetBaseConfigOpenApiVO) GetWanPortList() []string`

GetWanPortList returns the WanPortList field if non-nil, zero value otherwise.

### GetWanPortListOk

`func (o *InternetBaseConfigOpenApiVO) GetWanPortListOk() (*[]string, bool)`

GetWanPortListOk returns a tuple with the WanPortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortList

`func (o *InternetBaseConfigOpenApiVO) SetWanPortList(v []string)`

SetWanPortList sets WanPortList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


