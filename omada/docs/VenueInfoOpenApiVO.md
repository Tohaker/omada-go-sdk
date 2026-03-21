# VenueInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **int32** | Venue Group.Parameter group should be a value as follows: [0: Unspecified; 1: Assembly; 2: Business; 3: Educational; 4: Factory and Industrial; 5: Institutional; 6: Mercantile; 7: Residential; 8: Storage; 9: Utility and Miscellaneous; 10: Vehicular; 11: Outdoor] | 
**Name** | Pointer to **string** | Network’s venue name, identifying the physical location of the network.&lt;br /&gt;Note: It should contain between 1 and 64 visible ASCII characters, with no Spaces at the beginning and end, and Spaces in between. | [optional] 
**Type** | **int32** | Venue Type.When Venue Group &#x3D; 0, type should be a value as follows:[0: Unspecified]When Venue Group &#x3D; 1, type should be a value as follows:[0: Unspecified Assembly;1: Arena;2: Stadium;3: Passenger Terminal (e.g., airport, bus, ferry, train station);4: Amphitheater;5: Amusement Park;6: Place of Worship;7: Convention Center;8: Library;9: Museum;10: Restaurant;11: Theater;12: Bar;13: Coffee Shop;14: Zoo or Aquarium;15: Emergency Coordination Center]When Venue Group &#x3D; 2, type should be a value as follows:[0: Unspecified Business;1: Doctor or Dentist office;2: Bank;3: Fire Station;4: Police Station;5: Post Office;6: Professional Office;7: Research and Development Facility;8: Attorney Office]When Venue Group &#x3D; 3, type should be a value as follows:[0: Unspecified Educational;1: School,Primary;2: School, Secondary;3: University or College;]When Venue Group &#x3D; 4, type should be a value as follows:[0: Unspecified Factory and Industrial;1: Factory]When Venue Group &#x3D; 5, type should be a value as follows:[0: Unspecified Institutional;1: Hospital;2: Long-Term Care Facility (e.g. Nursing home, Hospice, etc.);3: Alchohol and Drug Re-habilitation Center;4: Group Home;5: Prison or Jail;]When Venue Group &#x3D; 6, type should be a value as follows:[0: Unspecified Mercantile;1: Retail Store;2: Grocery Market;3: Automotive Service Station;4: Shopping Mall;5: Gas Station;]When Venue Group &#x3D; 7, type should be a value as follows:[0: Unspecified Residential;1: Private Residence;2: Hotel or Motel;3: Dormitory;4: Boarding House;]When Venue Group &#x3D; 8, type should be a value as follows:[0: Unspecified Storage]When Venue Group &#x3D; 9, type should be a value as follows:[0: Unspecified Utility and Miscellaneous]When Venue Group &#x3D; 10, type should be a value as follows:[0: Unspecified Vehicular;1: Automobile or Truck;2: Airplane;3: Bus;4: Ferry;5: Ship or Boat;6: Train;7: Motor Bike;]When Venue Group &#x3D; 11, type should be a value as follows:[0: Unspecified Outdoor;1: Muni-mesh Network;2: City Park;3: Rest Area;4: Traffic Control;5: Bus Stop;6: Kiosk] | 

## Methods

### NewVenueInfoOpenApiVO

`func NewVenueInfoOpenApiVO(group int32, type_ int32, ) *VenueInfoOpenApiVO`

NewVenueInfoOpenApiVO instantiates a new VenueInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVenueInfoOpenApiVOWithDefaults

`func NewVenueInfoOpenApiVOWithDefaults() *VenueInfoOpenApiVO`

NewVenueInfoOpenApiVOWithDefaults instantiates a new VenueInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *VenueInfoOpenApiVO) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VenueInfoOpenApiVO) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VenueInfoOpenApiVO) SetGroup(v int32)`

SetGroup sets Group field to given value.


### GetName

`func (o *VenueInfoOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VenueInfoOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VenueInfoOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VenueInfoOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *VenueInfoOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VenueInfoOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VenueInfoOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


