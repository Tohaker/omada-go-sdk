# MacFiltering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterMode** | **int32** | Filter mode should be a value as follows: 0: allow; 1: deny. | 
**Id** | Pointer to **string** | ID of the MAC filtering entity. | [optional] 
**MacAddresses** | Pointer to [**[]OsgMacFilterAddressOpenApiVO**](OsgMacFilterAddressOpenApiVO.md) | MAC addresses of the MAC filtering entity. | [optional] 
**MacGroupIds** | Pointer to **[]string** | MAC groups of the MAC filtering entity. MAC group can be created using &#39;Create a new group profile&#39; interface, and MAC group ID can be obtained from &#39;Get group profile list&#39; interface. | [optional] 
**Name** | **string** | Name of the MAC filtering entity. | 
**Type** | **int32** | Type should be a value as follows: 0: macAddresses; 1: macGroupIds. | 

## Methods

### NewMacFiltering

`func NewMacFiltering(filterMode int32, name string, type_ int32, ) *MacFiltering`

NewMacFiltering instantiates a new MacFiltering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacFilteringWithDefaults

`func NewMacFilteringWithDefaults() *MacFiltering`

NewMacFilteringWithDefaults instantiates a new MacFiltering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterMode

`func (o *MacFiltering) GetFilterMode() int32`

GetFilterMode returns the FilterMode field if non-nil, zero value otherwise.

### GetFilterModeOk

`func (o *MacFiltering) GetFilterModeOk() (*int32, bool)`

GetFilterModeOk returns a tuple with the FilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMode

`func (o *MacFiltering) SetFilterMode(v int32)`

SetFilterMode sets FilterMode field to given value.


### GetId

`func (o *MacFiltering) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MacFiltering) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MacFiltering) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MacFiltering) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMacAddresses

`func (o *MacFiltering) GetMacAddresses() []OsgMacFilterAddressOpenApiVO`

GetMacAddresses returns the MacAddresses field if non-nil, zero value otherwise.

### GetMacAddressesOk

`func (o *MacFiltering) GetMacAddressesOk() (*[]OsgMacFilterAddressOpenApiVO, bool)`

GetMacAddressesOk returns a tuple with the MacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddresses

`func (o *MacFiltering) SetMacAddresses(v []OsgMacFilterAddressOpenApiVO)`

SetMacAddresses sets MacAddresses field to given value.

### HasMacAddresses

`func (o *MacFiltering) HasMacAddresses() bool`

HasMacAddresses returns a boolean if a field has been set.

### GetMacGroupIds

`func (o *MacFiltering) GetMacGroupIds() []string`

GetMacGroupIds returns the MacGroupIds field if non-nil, zero value otherwise.

### GetMacGroupIdsOk

`func (o *MacFiltering) GetMacGroupIdsOk() (*[]string, bool)`

GetMacGroupIdsOk returns a tuple with the MacGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacGroupIds

`func (o *MacFiltering) SetMacGroupIds(v []string)`

SetMacGroupIds sets MacGroupIds field to given value.

### HasMacGroupIds

`func (o *MacFiltering) HasMacGroupIds() bool`

HasMacGroupIds returns a boolean if a field has been set.

### GetName

`func (o *MacFiltering) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MacFiltering) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MacFiltering) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MacFiltering) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MacFiltering) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MacFiltering) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


