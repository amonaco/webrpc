// Code generated by statik. DO NOT EDIT.

// Package contains static assets.
package embed

var	Asset = "PK\x03\x04\x14\x00\x08\x00\x00\x00\x01\xa7\x88N\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00client.js.tmplUT\x05\x00\x01c\xb5\xab\\{{define \"client\"}}\n\n{{- if .Services}}\n// Client\n\n{{- range .Services}}\n{{exportKeyword}}class {{.Name}} {\n  constructor(hostname, fetch) {\n    this.path = '/rpc/{{.Name}}/';\n    this.hostname = hostname;\n    this.fetch = fetch;\n  }\n\n  url(name) {\n    return this.hostname + this.path + name;\n  }\n\n  {{range .Methods}}\n  {{.Name | methodName}}({{.Inputs | methodInputs}} = {}) {\n    return this.fetch(\n      this.url('{{.Name}}'),\n      {{if .Inputs | len}}\n      createHTTPRequest(args, headers)\n      {{else}}\n      createHTTPRequest({}, headers)\n      {{end}}\n    ).then((res) => {\n      if (!res.ok) {\n        return throwHTTPError(res);\n      }\n      return res.json().then((_data) => {\n        return {\n        {{- $outputsCount := .Outputs|len -}}\n        {{- range $i, $output := .Outputs}}\n          {{$output | newOutputArgResponse}}{{listComma $i $outputsCount}}\n        {{- end}}\n        }\n      })\n\n    })\n  }\n  {{end}}\n}\n{{end -}}\n\n{{end -}}\n{{end}}\nPK\x07\x08U\xda\xbd\xa2\xc5\x03\x00\x00\xc5\x03\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xf8\xa6\x88N\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0f\x00	\x00helpers.js.tmplUT\x05\x00\x01U\xb5\xab\\{{define \"helpers\"}}\n\n{{exportKeyword}}const throwHTTPError = (resp) => {\n  return resp.json().then((err) => { throw err; });\n}\n\n{{exportKeyword}}const createHTTPRequest = (body = {}, headers = {}) => {\n  return {\n    method: 'POST',\n    headers: { ...headers, 'Content-Type': 'application/json' },\n    body: JSON.stringify(body || {})\n  };\n}\n\n{{end}}\nPK\x07\x08&7\xa6\xc5`\x01\x00\x00`\x01\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xf8\xa6\x88N\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x11\x00	\x00proto.gen.js.tmplUT\x05\x00\x01U\xb5\xab\\{{- define \"proto\" -}}\n// {{.Name}} {{.Version}}\n// --\n// This file has been generated by https://github.com/webrpc/webrpc using gen/javascript\n// Do not edit by hand. Update your webrpc schema and re-generate.\n\n{{template \"types\" .}}\n{{- if .TargetOpts.Client}}\n  {{template \"client\" .}}\n{{- end}}\n{{- if .TargetOpts.Server}}\n  {{template \"server\" .}}\n{{- end}}\n{{template \"helpers\" .}}\n{{- end}}\nPK\x07\x08P\x13\x82\x07\x8e\x01\x00\x00\x8e\x01\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xf8\xa6\x88N\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00server.js.tmplUT\x05\x00\x01U\xb5\xab\\{{define \"server\"}}\n{{- if .Services}}\n// TODO: Server\n{{end -}}\n{{end}}\nPK\x07\x08\x8a@[\xefI\x00\x00\x00I\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xf8\xa6\x88N\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0d\x00	\x00types.js.tmplUT\x05\x00\x01U\xb5\xab\\{{define \"types\"}}\n\n{{- if .Messages -}}\n{{range .Messages -}}\n\n{{if .Type | isEnum -}}\n{{$enumName := .Name}}\n{{exportKeyword}}var {{$enumName}};\n(function ({{$enumName}}) {\n{{- range $i, $field := .Fields}}\n  {{$enumName}}[\"{{$field.Name}}\"] = \"{{$field.Name}}\";\n{{- end}}\n})({{$enumName}} || ({{$enumName}} = {}));\n{{end -}}\n\n{{- if .Type | isStruct  }}\n{{exportKeyword}}class {{.Name}} {\n  constructor(_data) {\n    this._data = {};\n    if (_data) {\n      {{range .Fields -}}\n      this._data['{{. | exportedJSONField}}'] = _data['{{. | exportedJSONField}}'];\n      {{end}}\n    }\n  }\n  {{ range .Fields -}}\n  get {{. | exportedJSONField}}() {\n    return this._data['{{. | exportedJSONField }}'];\n  }\n  set {{. | exportedJSONField}}(value) {\n    this._data['{{. | exportedJSONField}}'] = value;\n  }\n  {{end}}\n  toJSON() {\n    return this._data;\n  }\n}\n{{end -}}\n{{end -}}\n{{end -}}\n\n{{end}}\nPK\x07\x08L\x96\xc7\x8f|\x03\x00\x00|\x03\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x01\xa7\x88NU\xda\xbd\xa2\xc5\x03\x00\x00\xc5\x03\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x00client.js.tmplUT\x05\x00\x01c\xb5\xab\\PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xf8\xa6\x88N&7\xa6\xc5`\x01\x00\x00`\x01\x00\x00\x0f\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\n\x04\x00\x00helpers.js.tmplUT\x05\x00\x01U\xb5\xab\\PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xf8\xa6\x88NP\x13\x82\x07\x8e\x01\x00\x00\x8e\x01\x00\x00\x11\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xb0\x05\x00\x00proto.gen.js.tmplUT\x05\x00\x01U\xb5\xab\\PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xf8\xa6\x88N\x8a@[\xefI\x00\x00\x00I\x00\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x86\x07\x00\x00server.js.tmplUT\x05\x00\x01U\xb5\xab\\PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xf8\xa6\x88NL\x96\xc7\x8f|\x03\x00\x00|\x03\x00\x00\x0d\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x14\x08\x00\x00types.js.tmplUT\x05\x00\x01U\xb5\xab\\PK\x05\x06\x00\x00\x00\x00\x05\x00\x05\x00\\\x01\x00\x00\xd4\x0b\x00\x00\x00\x00"
