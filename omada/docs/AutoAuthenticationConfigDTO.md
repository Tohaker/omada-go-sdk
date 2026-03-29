# AutoAuthenticationConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoAuthenticationStatus** | **string** | Auto authentication status should be a value as follows:ENABLE,DISABLE | 

## Methods

### NewAutoAuthenticationConfigDTO

`func NewAutoAuthenticationConfigDTO(autoAuthenticationStatus string, ) *AutoAuthenticationConfigDTO`

NewAutoAuthenticationConfigDTO instantiates a new AutoAuthenticationConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoAuthenticationConfigDTOWithDefaults

`func NewAutoAuthenticationConfigDTOWithDefaults() *AutoAuthenticationConfigDTO`

NewAutoAuthenticationConfigDTOWithDefaults instantiates a new AutoAuthenticationConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoAuthenticationStatus

`func (o *AutoAuthenticationConfigDTO) GetAutoAuthenticationStatus() string`

GetAutoAuthenticationStatus returns the AutoAuthenticationStatus field if non-nil, zero value otherwise.

### GetAutoAuthenticationStatusOk

`func (o *AutoAuthenticationConfigDTO) GetAutoAuthenticationStatusOk() (*string, bool)`

GetAutoAuthenticationStatusOk returns a tuple with the AutoAuthenticationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAuthenticationStatus

`func (o *AutoAuthenticationConfigDTO) SetAutoAuthenticationStatus(v string)`

SetAutoAuthenticationStatus sets AutoAuthenticationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


