// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/modules/wifi/beacon": {
            "post": {
                "description": "Beacon path will send a fake fake management beacons in order to create N access point.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "Beacon attack",
                "parameters": [
                    {
                        "description": "beacon info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/wifi_common.Beaconer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/connectAp/{interfaceName}": {
            "post": {
                "description": "ConnectAp path will connect you to an access point.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "Connect access point",
                "parameters": [
                    {
                        "description": "ap info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/wifi_common.ConnectAp"
                        }
                    },
                    {
                        "type": "string",
                        "description": "interface name",
                        "name": "interfaceName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/cptHandshake": {
            "get": {
                "description": "CaptureHandshake path will use the wireless interface in monitor mode and capture packets and filter handshakes all over the flore.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "Capture handshakes",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/deauth": {
            "post": {
                "description": "Deauth path will deauthenticate user from an access point.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "Deauth client",
                "parameters": [
                    {
                        "description": "deauth info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/wifi_common.Deauther"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/interface": {
            "get": {
                "description": "Interfaces path will list device wirless interfaces, think of it as iwconfig.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "List wireless interfaces",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/wifi_common.WirelessInterface"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/probe": {
            "post": {
                "description": "Probe path will send a fake client probe with the given station BSSID, searching for ESSID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "Probe attack",
                "parameters": [
                    {
                        "description": "probe info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/wifi_common.Prober"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/rogueAp": {
            "post": {
                "description": "RogueAP path will send a fake fake management beacons in order to create rogue access point.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "rogue access point attack",
                "parameters": [
                    {
                        "description": "rogue ap info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/wifi_common.RogueAp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/scanAp/{interfaceName}": {
            "get": {
                "description": "ScanAP path will put wireless interface in monitor mode and capture packets and filter for access point.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "List access points",
                "parameters": [
                    {
                        "type": "string",
                        "description": "interface name",
                        "name": "interfaceName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/scanClient": {
            "get": {
                "description": "ScanClient path will use the wireless interface in monitor mode and capture packets and filter for the already found access point's client.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "List clients of access points",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/stop": {
            "get": {
                "description": "Stop path will kill all process of recon.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "stop recon",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/stopBeaconer": {
            "get": {
                "description": "StopBeaconer path will kill process of sending beacons.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "stop beacon attack",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/stopCptHandshake": {
            "get": {
                "description": "StopCptHandshake path will kill process of searching access points.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "stop capture handshake",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/stopRogueAp": {
            "get": {
                "description": "StopRogueAP path will kill process of sending beacons.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "stop rogue access point attack",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/modules/wifi/stopScanClient": {
            "get": {
                "description": "StopScanClient path will kill process of searching for access point clients.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wifi"
                ],
                "summary": "stop client recon",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_common.ErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1_common.ErrorMessage": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "wifi.InterfaceType": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4,
                5,
                6,
                7,
                8,
                9,
                10,
                11,
                12
            ],
            "x-enum-varnames": [
                "InterfaceTypeUnspecified",
                "InterfaceTypeAdHoc",
                "InterfaceTypeStation",
                "InterfaceTypeAP",
                "InterfaceTypeAPVLAN",
                "InterfaceTypeWDS",
                "InterfaceTypeMonitor",
                "InterfaceTypeMeshPoint",
                "InterfaceTypeP2PClient",
                "InterfaceTypeP2PGroupOwner",
                "InterfaceTypeP2PDevice",
                "InterfaceTypeOCB",
                "InterfaceTypeNAN"
            ]
        },
        "wifi_common.Beaconer": {
            "type": "object",
            "properties": {
                "apChannel": {
                    "type": "integer"
                },
                "apEncryption": {
                    "type": "boolean"
                },
                "apName": {
                    "type": "string"
                },
                "numberOfAP": {
                    "type": "integer"
                }
            }
        },
        "wifi_common.ConnectAp": {
            "type": "object",
            "properties": {
                "apName": {
                    "type": "string"
                },
                "apPass": {
                    "type": "string"
                }
            }
        },
        "wifi_common.Deauther": {
            "type": "object",
            "properties": {
                "apMac": {
                    "type": "string"
                },
                "clientMac": {
                    "type": "string"
                }
            }
        },
        "wifi_common.Prober": {
            "type": "object",
            "properties": {
                "apMac": {
                    "type": "string"
                },
                "apName": {
                    "type": "string"
                }
            }
        },
        "wifi_common.RogueAp": {
            "type": "object",
            "properties": {
                "apChannel": {
                    "type": "integer"
                },
                "apEncryption": {
                    "type": "boolean"
                },
                "apMac": {
                    "type": "string"
                },
                "apName": {
                    "type": "string"
                }
            }
        },
        "wifi_common.WirelessInterface": {
            "type": "object",
            "properties": {
                "device": {
                    "type": "integer"
                },
                "frequency": {
                    "type": "integer"
                },
                "hardwareAddr": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phy": {
                    "type": "integer"
                },
                "type": {
                    "$ref": "#/definitions/wifi.InterfaceType"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "mf-backend",
	Description:      "The backend that holds the raspberry pi HACKING modules.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
