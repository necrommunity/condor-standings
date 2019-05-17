
package main
//This file is generated automatically. Do not try to edit it manually.

var resourceListingJson = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://some.pizza/",
    "apis": [
        {
            "path": "/api",
            "description": "API"
        }
    ],
    "info": {
        "contact": "sillypairs@gmail.com",
        "termsOfServiceUrl": "http://google.com/",
        "license": "BSD",
        "licenseUrl": "http://opensource.org/licenses/BSD-2-Clause"
    }
}`
var apiDescriptionsJson = map[string]string{"api":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://some.pizza/",
    "resourcePath": "/api",
    "produces": [
        "application/json"
    ],
    "apis": [
        {
            "path": "/api/event",
            "description": "Lists all events found by name",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Events",
                    "type": "github.com.sillypears.condor-standings.src.models.ReturnedTable",
                    "items": {},
                    "summary": "Lists all events found by name",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.sillypears.condor-standings.src.models.ReturnedTable"
                        },
                        {
                            "code": 404,
                            "message": "No Events Found",
                            "responseType": "object",
                            "responseModel": "github.com.sillypears.condor-standings.src.APIError"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        },
        {
            "path": "/api/event/{event}",
            "description": "Lists everything found for the event",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Event Listing",
                    "type": "github.com.sillypears.condor-standings.src.models.Event",
                    "items": {},
                    "summary": "Lists everything found for the event",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "event",
                            "description": "Event Name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.sillypears.condor-standings.src.models.Event"
                        },
                        {
                            "code": 404,
                            "message": "Event not found",
                            "responseType": "object",
                            "responseModel": "github.com.sillypears.condor-standings.src.APIError"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        },
        {
            "path": "/api/teamresults",
            "description": "Lists everything found for the season 7 teams",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Team Results Listing",
                    "type": "github.com.sillypears.condor-standings.src.models.Result",
                    "items": {},
                    "summary": "Lists everything found for the season 7 teams",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.sillypears.condor-standings.src.models.Result"
                        },
                        {
                            "code": 404,
                            "message": "Nothing found",
                            "responseType": "object",
                            "responseModel": "github.com.sillypears.condor-standings.src.APIError"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        },
        {
            "path": "/api/s",
            "description": "Lists all racers from season 8 specifically",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "S Results Listing",
                    "type": "github.com.sillypears.condor-standings.src.models.Event",
                    "items": {},
                    "summary": "Lists all racers from season 8 specifically",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "event",
                            "description": "Event Name",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.sillypears.condor-standings.src.models.Event"
                        },
                        {
                            "code": 404,
                            "message": "Nothing found",
                            "responseType": "object",
                            "responseModel": "github.com.sillypears.condor-standings.src.APIError"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.sillypears.condor-standings.src.APIError": {
            "id": "github.com.sillypears.condor-standings.src.APIError",
            "properties": {
                "ErrorCode": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "ErrorMessage": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.sillypears.condor-standings.src.models.Event": {
            "id": "github.com.sillypears.condor-standings.src.models.Event",
            "properties": {
                "Participants": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "github.com.sillypears.condor-standings.src.models.Participant"
                    },
                    "format": ""
                },
                "eventName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.sillypears.condor-standings.src.models.Participant": {
            "id": "github.com.sillypears.condor-standings.src.models.Participant",
            "properties": {
                "discordUsername": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "eventLosses": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "eventPlayed": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "eventPoints": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "eventWins": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "groupName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "racerID": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "tierName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "twitchUsername": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.sillypears.condor-standings.src.models.Result": {
            "id": "github.com.sillypears.condor-standings.src.models.Result",
            "properties": {
                "racer1": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "racer2": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "team1": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "team2": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "winner": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "github.com.sillypears.condor-standings.src.models.ReturnedTable": {
            "id": "github.com.sillypears.condor-standings.src.models.ReturnedTable",
            "properties": {
                "eventName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "prettyName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,}
