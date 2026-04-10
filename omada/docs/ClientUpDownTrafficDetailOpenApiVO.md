# ClientUpDownTrafficDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]ApplicationUpDownTrafficDetailOpenApiVO**](ApplicationUpDownTrafficDetailOpenApiVO.md) | Application uplink and downlink traffic data. | [optional] 
**Categories** | Pointer to [**[]CategoryUpDownTrafficDetailOpenApiVO**](CategoryUpDownTrafficDetailOpenApiVO.md) | Application uplink and downlink categories. | [optional] 
**ClientMac** | Pointer to **string** | The mac of the client. | [optional] 
**ClientName** | Pointer to **string** | The name of the client. | [optional] 
**Down** | Pointer to **int64** | Down traffic. | [optional] 
**Manager** | Pointer to **bool** | Whether it is the client currently being managed. | [optional] 
**Network** | Pointer to **string** | Network name. | [optional] 
**Type** | Pointer to **string** | Type of client. | [optional] 
**Up** | Pointer to **int64** | Up traffic. | [optional] 
**VlanId** | Pointer to **string** | VLAN ID. The VLAN ID range is 1–4096 or Untag. | [optional] 

## Methods

### NewClientUpDownTrafficDetailOpenApiVO

`func NewClientUpDownTrafficDetailOpenApiVO() *ClientUpDownTrafficDetailOpenApiVO`

NewClientUpDownTrafficDetailOpenApiVO instantiates a new ClientUpDownTrafficDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientUpDownTrafficDetailOpenApiVOWithDefaults

`func NewClientUpDownTrafficDetailOpenApiVOWithDefaults() *ClientUpDownTrafficDetailOpenApiVO`

NewClientUpDownTrafficDetailOpenApiVOWithDefaults instantiates a new ClientUpDownTrafficDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetApplications() []ApplicationUpDownTrafficDetailOpenApiVO`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetApplicationsOk() (*[]ApplicationUpDownTrafficDetailOpenApiVO, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ClientUpDownTrafficDetailOpenApiVO) SetApplications(v []ApplicationUpDownTrafficDetailOpenApiVO)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ClientUpDownTrafficDetailOpenApiVO) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCategories

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetCategories() []CategoryUpDownTrafficDetailOpenApiVO`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetCategoriesOk() (*[]CategoryUpDownTrafficDetailOpenApiVO, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ClientUpDownTrafficDetailOpenApiVO) SetCategories(v []CategoryUpDownTrafficDetailOpenApiVO)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ClientUpDownTrafficDetailOpenApiVO) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetClientMac

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *ClientUpDownTrafficDetailOpenApiVO) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *ClientUpDownTrafficDetailOpenApiVO) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetClientName

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ClientUpDownTrafficDetailOpenApiVO) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ClientUpDownTrafficDetailOpenApiVO) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDown

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetDown() int64`

GetDown returns the Down field if non-nil, zero value otherwise.

### GetDownOk

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetDownOk() (*int64, bool)`

GetDownOk returns a tuple with the Down field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDown

`func (o *ClientUpDownTrafficDetailOpenApiVO) SetDown(v int64)`

SetDown sets Down field to given value.

### HasDown

`func (o *ClientUpDownTrafficDetailOpenApiVO) HasDown() bool`

HasDown returns a boolean if a field has been set.

### GetManager

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetManager() bool`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetManagerOk() (*bool, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *ClientUpDownTrafficDetailOpenApiVO) SetManager(v bool)`

SetManager sets Manager field to given value.

### HasManager

`func (o *ClientUpDownTrafficDetailOpenApiVO) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetNetwork

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClientUpDownTrafficDetailOpenApiVO) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ClientUpDownTrafficDetailOpenApiVO) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetType

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientUpDownTrafficDetailOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientUpDownTrafficDetailOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUp

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetUp() int64`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetUpOk() (*int64, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *ClientUpDownTrafficDetailOpenApiVO) SetUp(v int64)`

SetUp sets Up field to given value.

### HasUp

`func (o *ClientUpDownTrafficDetailOpenApiVO) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetVlanId

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ClientUpDownTrafficDetailOpenApiVO) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ClientUpDownTrafficDetailOpenApiVO) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ClientUpDownTrafficDetailOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


