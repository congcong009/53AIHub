package docs

import (
	"bytes"
	"encoding/json"
	"strings"
)

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version: "0.1",
	Host:    "",
	BasePath: "/api",
	Title:   "7 AIHub API",
	Description: "This is the API documentation for 7 AIHub",
}

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

func (s swaggerInfo) ReadDoc() string {
	sInfo := SwaggerInfo
	jsonInfo, _ := json.MarshalIndent(sInfo, "", "  ")
	// 将占位符替换为实际值
	doc := `{
  "swagger": "2.0",
  "info": {
    "description": "{{.Description}}",
    "title": "{{.Title}}",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "name": "API Support",
      "email": "admin@53ai.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "{{.Version}}"
  },
  "host": "{{.Host}}",
  "basePath": "{{.BasePath}}",
  "paths": {},
  "definitions": {}
}`
	// 简单替换模板变量
	doc = strings.ReplaceAll(doc, "{{.Description}}", sInfo.Description)
	doc = strings.ReplaceAll(doc, "{{.Title}}", sInfo.Title)
	doc = strings.ReplaceAll(doc, "{{.Version}}", sInfo.Version)
	doc = strings.ReplaceAll(doc, "{{.Host}}", sInfo.Host)
	doc = strings.ReplaceAll(doc, "{{.BasePath}}", sInfo.BasePath)
	
	var buffer bytes.Buffer
	json.Indent(&buffer, jsonInfo, "", "  ")
	return doc
}

func (s swaggerInfo) SwaggerJSON() []byte {
	jsonString := s.ReadDoc()
	return []byte(jsonString)
}