
# 


Table of Contents

1. [API](#api)

<a name="api"></a>

## api

| Specification | Value |
|-----|-----|
| Resource Path | /api |
| API Version | 1.0.0 |
| BasePath for the API | http://wow.freepizza.how/ |
| Consumes | application/json |
| Produces | application/json |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /api/event | [GET](#Events) | Lists all events found by name |
| /api/event/\{event\} | [GET](#Event Listing) | Lists everything found for the event |



<a name="Events"></a>

#### API: /api/event (GET)


Lists all events found by name



| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [ReturnedTables](#github.com.sillypears.condor-standings.src.ReturnedTables) |  |
| 404 | object | [APIError](#github.com.sillypears.condor-standings.src.APIError) | No Events Found |


<a name="Event Listing"></a>

#### API: /api/event/\{event\} (GET)


Lists everything found for the event



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| event | path | string | Event Name | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Event](#github.com.sillypears.condor-standings.src.Event) |  |
| 404 | object | [APIError](#github.com.sillypears.condor-standings.src.APIError) | Event not found |




### Models

<a name="github.com.sillypears.condor-standings.src.APIError"></a>

#### APIError

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| ErrorCode | int |  |
| ErrorMessage | string |  |

<a name="github.com.sillypears.condor-standings.src.Event"></a>

#### Event

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Participants | array |  |
| eventName | string |  |

<a name="github.com.sillypears.condor-standings.src.ReturnedTables"></a>

#### ReturnedTables

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| eventNames | array |  |

<a name="github.com.sillypears.condor-standings.src.models.Participant"></a>

#### Participant

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| discordID | int |  |
| discordUsername | string |  |
| eventPlayed | int |  |
| eventPoints | int |  |
| groupName | string |  |

