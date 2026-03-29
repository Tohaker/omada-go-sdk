# InternetBaseInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayModel** | **int32** | Gateway model. The model number refers to function can be obtained from &#39;Get supported gateway model list for pre-configuration&#39; interface. | 
**Interval** | **int32** | Online detection interval(second), 0 means disable. | 
**PortList** | [**[]InternetBasicInfo**](InternetBasicInfo.md) | Port info list | 
**PreConfiguration** | **bool** | You can pre-configure the model-relative settings when it is true. | 

## Methods

### NewInternetBaseInfoOpenApiVO

`func NewInternetBaseInfoOpenApiVO(gatewayModel int32, interval int32, portList []InternetBasicInfo, preConfiguration bool, ) *InternetBaseInfoOpenApiVO`

NewInternetBaseInfoOpenApiVO instantiates a new InternetBaseInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetBaseInfoOpenApiVOWithDefaults

`func NewInternetBaseInfoOpenApiVOWithDefaults() *InternetBaseInfoOpenApiVO`

NewInternetBaseInfoOpenApiVOWithDefaults instantiates a new InternetBaseInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayModel

`func (o *InternetBaseInfoOpenApiVO) GetGatewayModel() int32`

GetGatewayModel returns the GatewayModel field if non-nil, zero value otherwise.

### GetGatewayModelOk

`func (o *InternetBaseInfoOpenApiVO) GetGatewayModelOk() (*int32, bool)`

GetGatewayModelOk returns a tuple with the GatewayModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayModel

`func (o *InternetBaseInfoOpenApiVO) SetGatewayModel(v int32)`

SetGatewayModel sets GatewayModel field to given value.


### GetInterval

`func (o *InternetBaseInfoOpenApiVO) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InternetBaseInfoOpenApiVO) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InternetBaseInfoOpenApiVO) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetPortList

`func (o *InternetBaseInfoOpenApiVO) GetPortList() []InternetBasicInfo`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *InternetBaseInfoOpenApiVO) GetPortListOk() (*[]InternetBasicInfo, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *InternetBaseInfoOpenApiVO) SetPortList(v []InternetBasicInfo)`

SetPortList sets PortList field to given value.


### GetPreConfiguration

`func (o *InternetBaseInfoOpenApiVO) GetPreConfiguration() bool`

GetPreConfiguration returns the PreConfiguration field if non-nil, zero value otherwise.

### GetPreConfigurationOk

`func (o *InternetBaseInfoOpenApiVO) GetPreConfigurationOk() (*bool, bool)`

GetPreConfigurationOk returns a tuple with the PreConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConfiguration

`func (o *InternetBaseInfoOpenApiVO) SetPreConfiguration(v bool)`

SetPreConfiguration sets PreConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


