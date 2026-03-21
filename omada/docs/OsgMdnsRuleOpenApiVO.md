# OsgMdnsRuleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientNetworks** | **[]string** | LAN Network ID list of selected client networks. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface | 
**ServiceNetworks** | **[]string** | LAN Network ID list of selected service networks. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface | 

## Methods

### NewOsgMdnsRuleOpenApiVO

`func NewOsgMdnsRuleOpenApiVO(clientNetworks []string, serviceNetworks []string, ) *OsgMdnsRuleOpenApiVO`

NewOsgMdnsRuleOpenApiVO instantiates a new OsgMdnsRuleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgMdnsRuleOpenApiVOWithDefaults

`func NewOsgMdnsRuleOpenApiVOWithDefaults() *OsgMdnsRuleOpenApiVO`

NewOsgMdnsRuleOpenApiVOWithDefaults instantiates a new OsgMdnsRuleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientNetworks

`func (o *OsgMdnsRuleOpenApiVO) GetClientNetworks() []string`

GetClientNetworks returns the ClientNetworks field if non-nil, zero value otherwise.

### GetClientNetworksOk

`func (o *OsgMdnsRuleOpenApiVO) GetClientNetworksOk() (*[]string, bool)`

GetClientNetworksOk returns a tuple with the ClientNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetworks

`func (o *OsgMdnsRuleOpenApiVO) SetClientNetworks(v []string)`

SetClientNetworks sets ClientNetworks field to given value.


### GetServiceNetworks

`func (o *OsgMdnsRuleOpenApiVO) GetServiceNetworks() []string`

GetServiceNetworks returns the ServiceNetworks field if non-nil, zero value otherwise.

### GetServiceNetworksOk

`func (o *OsgMdnsRuleOpenApiVO) GetServiceNetworksOk() (*[]string, bool)`

GetServiceNetworksOk returns a tuple with the ServiceNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNetworks

`func (o *OsgMdnsRuleOpenApiVO) SetServiceNetworks(v []string)`

SetServiceNetworks sets ServiceNetworks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


