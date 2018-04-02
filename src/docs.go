
package main
//This file is generated automatically. Do not try to edit it manually.
var apiDescriptionsJson = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://wow.freepizza.how/",
    "info": {
        "contact": "sillypairs@gmail.com",
        "termsOfServiceUrl": "http://google.com/",
        "license": "BSD",
        "licenseUrl": "http://opensource.org/licenses/BSD-2-Clause"
    },
    "produces": [
        "application/json"
    ],
    "apis": [
        {
            "path": "/api",
            "description": "API"
        },
        {
            "path": "/api/event",
            "description": "Lists all events found by name",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Events",
                    "type": "github.com.sillypears.condor-standings.src.ReturnedTables",
                    "items": {},
                    "summary": "Lists all events found by name",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "github.com.sillypears.condor-standings.src.ReturnedTables"
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
                    "type": "github.com.sillypears.condor-standings.src.Event",
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
                            "responseModel": "github.com.sillypears.condor-standings.src.Event"
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
        "github.com.sillypears.condor-standings.src.Event": {
            "id": "github.com.sillypears.condor-standings.src.Event",
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
        "github.com.sillypears.condor-standings.src.ReturnedTables": {
            "id": "github.com.sillypears.condor-standings.src.ReturnedTables",
            "properties": {
                "eventNames": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "type": "string"
                    },
                    "format": ""
                }
            }
        },
        "github.com.sillypears.condor-standings.src.models.Participant": {
            "id": "github.com.sillypears.condor-standings.src.models.Participant",
            "properties": {
                "discordID": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "discordUsername": {
                    "type": "string",
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
                "groupName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`
