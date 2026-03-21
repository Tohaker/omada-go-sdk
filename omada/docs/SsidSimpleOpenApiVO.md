# SsidSimpleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidId** | Pointer to **string** | This field represents SSID ID. SSID can be created using &#39;Create new SSID&#39; (&#39;Create new SSID template&#39;) interface, and SSID ID can be obtained from &#39;Get SSID list&#39; (&#39;Get SSID template list&#39;) interface | [optional] 
**SsidName** | Pointer to **string** | This field represents SSID name | [optional] 

## Methods

### NewSsidSimpleOpenApiVO

`func NewSsidSimpleOpenApiVO() *SsidSimpleOpenApiVO`

NewSsidSimpleOpenApiVO instantiates a new SsidSimpleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidSimpleOpenApiVOWithDefaults

`func NewSsidSimpleOpenApiVOWithDefaults() *SsidSimpleOpenApiVO`

NewSsidSimpleOpenApiVOWithDefaults instantiates a new SsidSimpleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidId

`func (o *SsidSimpleOpenApiVO) GetSsidId() string`

GetSsidId returns the SsidId field if non-nil, zero value otherwise.

### GetSsidIdOk

`func (o *SsidSimpleOpenApiVO) GetSsidIdOk() (*string, bool)`

GetSsidIdOk returns a tuple with the SsidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidId

`func (o *SsidSimpleOpenApiVO) SetSsidId(v string)`

SetSsidId sets SsidId field to given value.

### HasSsidId

`func (o *SsidSimpleOpenApiVO) HasSsidId() bool`

HasSsidId returns a boolean if a field has been set.

### GetSsidName

`func (o *SsidSimpleOpenApiVO) GetSsidName() string`

GetSsidName returns the SsidName field if non-nil, zero value otherwise.

### GetSsidNameOk

`func (o *SsidSimpleOpenApiVO) GetSsidNameOk() (*string, bool)`

GetSsidNameOk returns a tuple with the SsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidName

`func (o *SsidSimpleOpenApiVO) SetSsidName(v string)`

SetSsidName sets SsidName field to given value.

### HasSsidName

`func (o *SsidSimpleOpenApiVO) HasSsidName() bool`

HasSsidName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


