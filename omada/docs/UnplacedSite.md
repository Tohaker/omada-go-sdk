# UnplacedSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | site id | [optional] 
**Name** | Pointer to **string** | site name | [optional] 
**Region** | Pointer to **string** | site region | [optional] 
**Type** | Pointer to **int32** | site type. 0:advanced ; 1:pro | [optional] 

## Methods

### NewUnplacedSite

`func NewUnplacedSite() *UnplacedSite`

NewUnplacedSite instantiates a new UnplacedSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnplacedSiteWithDefaults

`func NewUnplacedSiteWithDefaults() *UnplacedSite`

NewUnplacedSiteWithDefaults instantiates a new UnplacedSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnplacedSite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnplacedSite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnplacedSite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UnplacedSite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UnplacedSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnplacedSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnplacedSite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnplacedSite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *UnplacedSite) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UnplacedSite) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UnplacedSite) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UnplacedSite) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetType

`func (o *UnplacedSite) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnplacedSite) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnplacedSite) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnplacedSite) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


