// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/Febriand1",
            "email": "1214039@std.ulbi.ac.id"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/delete/{id}": {
            "delete": {
                "description": "Hapus data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Delete data presensi.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/delnilai/{id}": {
            "delete": {
                "description": "Hapus data nilai.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Nilai"
                ],
                "summary": "Delete data nilai.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/ins": {
            "post": {
                "description": "Input data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Insert data presensi.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/insnilai": {
            "post": {
                "description": "Input data nilai.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Nilai"
                ],
                "summary": "Insert data nilai.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Nilai"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Nilai"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/nilai": {
            "get": {
                "description": "Mengambil semua data nilai.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Nilai"
                ],
                "summary": "Get All Data Nilai.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Nilai"
                        }
                    }
                }
            }
        },
        "/nilai/{id}": {
            "get": {
                "description": "Ambil per ID data nilai.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Nilai"
                ],
                "summary": "Get By ID Data Nilai.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Nilai"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/presensinew": {
            "get": {
                "description": "Mengambil semua data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Get All Data Presensi.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    }
                }
            }
        },
        "/presensinew/{id}": {
            "get": {
                "description": "Ambil per ID data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Get By ID Data Presensi.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/upd/{id}": {
            "put": {
                "description": "Ubah data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Update data presensi.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/updnilai/{id}": {
            "put": {
                "description": "Ubah data nilai.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Nilai"
                ],
                "summary": "Update data nilai.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Nilai"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Nilai"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Dosen": {
            "type": "object",
            "properties": {
                "namadosen": {
                    "description": "ID            primitive.ObjectID ` + "`" + `bson:\"_id,omitempty\" json:\"_id,omitempty\"` + "`" + `",
                    "type": "string",
                    "example": "Rizkyria"
                },
                "nik": {
                    "type": "string",
                    "example": "123.456.789"
                },
                "phonenumberd": {
                    "type": "string",
                    "example": "089876543210"
                }
            }
        },
        "controller.Grade": {
            "type": "object",
            "properties": {
                "namagrade": {
                    "type": "string",
                    "example": "A"
                },
                "skala": {
                    "type": "string",
                    "example": "75.87"
                }
            }
        },
        "controller.Mahasiswa": {
            "type": "object",
            "properties": {
                "nama": {
                    "description": "ID           primitive.ObjectID ` + "`" + `bson:\"_id,omitempty\" json:\"_id,omitempty\" example:\"123456789\" ` + "`" + `",
                    "type": "string",
                    "example": "Tes Swagger"
                },
                "npm": {
                    "type": "integer",
                    "example": 1234567
                },
                "phonenumber": {
                    "type": "string",
                    "example": "08123456789"
                }
            }
        },
        "controller.Matakuliah": {
            "type": "object",
            "properties": {
                "jadwal": {
                    "$ref": "#/definitions/controller.Waktu"
                },
                "nama_mk": {
                    "description": "ID       primitive.ObjectID ` + "`" + `bson:\"_id,omitempty\" json:\"_id,omitempty\" example:\"123456789\"` + "`" + `",
                    "type": "string",
                    "example": "Pemrograman III"
                },
                "pengampu": {
                    "$ref": "#/definitions/controller.Dosen"
                },
                "sks": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "controller.Nilai": {
            "type": "object",
            "properties": {
                "absensi": {
                    "$ref": "#/definitions/controller.Presensi"
                },
                "alltugas": {
                    "description": "ID        primitive.ObjectID ` + "`" + `bson:\"_id,omitempty\" json:\"_id,omitempty\"` + "`" + `",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controller.Tugas"
                        }
                    ]
                },
                "grade": {
                    "$ref": "#/definitions/controller.Grade"
                },
                "kategori": {
                    "$ref": "#/definitions/controller.Matakuliah"
                },
                "uas": {
                    "type": "integer",
                    "example": 123
                },
                "uts": {
                    "type": "integer",
                    "example": 123
                }
            }
        },
        "controller.Presensi": {
            "type": "object",
            "properties": {
                "biodata": {
                    "$ref": "#/definitions/controller.Mahasiswa"
                },
                "jumlahkehadiran": {
                    "description": "ID               primitive.ObjectID ` + "`" + `bson:\"_id,omitempty\" json:\"_id,omitempty\"` + "`" + `",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "controller.Tugas": {
            "type": "object",
            "properties": {
                "tugas1": {
                    "type": "integer",
                    "example": 123
                },
                "tugas2": {
                    "type": "integer",
                    "example": 123
                },
                "tugas3": {
                    "type": "integer",
                    "example": 123
                },
                "tugas4": {
                    "type": "integer",
                    "example": 123
                },
                "tugas5": {
                    "type": "integer",
                    "example": 123
                }
            }
        },
        "controller.Waktu": {
            "type": "object",
            "properties": {
                "hari": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Senin",
                        "Selasa",
                        "Rabu",
                        "Kamis",
                        "Jumat",
                        "Sabtu",
                        "Minggu"
                    ]
                },
                "jamkeluar": {
                    "type": "string",
                    "example": "16:00"
                },
                "jammasuk": {
                    "type": "string",
                    "example": "08:00"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "ws-nilai.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "TES  SWAG",
	Description:      "This is a sample server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
