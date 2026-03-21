# PonAutoAuthenticationRuleConfigDeleteDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | **[]string** | Entry identifier list | 
**PonPort** | **string** | Pon port.e.g.GPON 1/0/1 | 

## Methods

### NewPonAutoAuthenticationRuleConfigDeleteDTO

`func NewPonAutoAuthenticationRuleConfigDeleteDTO(keys []string, ponPort string, ) *PonAutoAuthenticationRuleConfigDeleteDTO`

NewPonAutoAuthenticationRuleConfigDeleteDTO instantiates a new PonAutoAuthenticationRuleConfigDeleteDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPonAutoAuthenticationRuleConfigDeleteDTOWithDefaults

`func NewPonAutoAuthenticationRuleConfigDeleteDTOWithDefaults() *PonAutoAuthenticationRuleConfigDeleteDTO`

NewPonAutoAuthenticationRuleConfigDeleteDTOWithDefaults instantiates a new PonAutoAuthenticationRuleConfigDeleteDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *PonAutoAuthenticationRuleConfigDeleteDTO) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *PonAutoAuthenticationRuleConfigDeleteDTO) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *PonAutoAuthenticationRuleConfigDeleteDTO) SetKeys(v []string)`

SetKeys sets Keys field to given value.


### GetPonPort

`func (o *PonAutoAuthenticationRuleConfigDeleteDTO) GetPonPort() string`

GetPonPort returns the PonPort field if non-nil, zero value otherwise.

### GetPonPortOk

`func (o *PonAutoAuthenticationRuleConfigDeleteDTO) GetPonPortOk() (*string, bool)`

GetPonPortOk returns a tuple with the PonPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPort

`func (o *PonAutoAuthenticationRuleConfigDeleteDTO) SetPonPort(v string)`

SetPonPort sets PonPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


