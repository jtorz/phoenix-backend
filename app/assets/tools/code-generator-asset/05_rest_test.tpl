{{- define "rest_varname"}}
    {{- if eq .GoDataType "float32"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else if eq .GoDataType "float64"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else if eq .GoDataType "int"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else if eq .GoDataType "int8"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else if eq .GoDataType "int16"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else if eq .GoDataType "int32"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else if eq .GoDataType "int64"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else if eq .GoDataType "uint"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else if eq .GoDataType "uint8"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else if eq .GoDataType "uint16"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else if eq .GoDataType "uint32"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else if eq .GoDataType "uint64"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else if eq .GoDataType "bool"}}
        {{- `{{`}}{{.GoVarName}}{{`}}`}}
    {{- else}}"{{`{{`}}{{.GoVarName}}{{`}}`}}"
    {{- end}}
{{- end}}
{{- define "datafield"}}
    {{- if or (.IsPK) (eq .GoField "UpdatedAt")}}{{template "rest_varname" .}}
    {{- else if eq .GoDataType "float32"}}0.0
    {{- else if eq .GoDataType "float64"}}0.0
    {{- else if eq .GoDataType "int"}}0
    {{- else if eq .GoDataType "int8"}}0
    {{- else if eq .GoDataType "int16"}}0
    {{- else if eq .GoDataType "int32"}}0
    {{- else if eq .GoDataType "int64"}}0
    {{- else if eq .GoDataType "uint"}}0
    {{- else if eq .GoDataType "uint8"}}0
    {{- else if eq .GoDataType "uint16"}}0
    {{- else if eq .GoDataType "uint32"}}0
    {{- else if eq .GoDataType "uint64"}}0
    {{- else if eq .GoDataType "bool"}}false
    {{- else}}""
    {{- end}}
{{- end}}
### GoDataType
# @name authenticate
POST {{`{{app-host}}`}}/api/public/foundation/account/login
Content-Type: application/json

{
    "Password": "{{`{{env-password}}`}}",
    "User": "{{`{{env-user}}`}}"
}

@Authorization = Bearer {{`{{authenticate.response.body.$.JWT}}`}}

#############################################################

###
# @name LIST_ALL_{{$.Entity.GoSlice}}
GET {{`{{app-host}}`}}/api/{{$.ServiceName | lowercase}}/{{$.Entity.GoSlice | lowercase}}
Content-Type: application/json
Authorization: {{`{{Authorization}}`}}

###
# @name LIST_ACTIVE_{{$.Entity.GoSlice}}
GET {{`{{app-host}}`}}/api/{{$.ServiceName | lowercase}}/{{$.Entity.GoSlice | lowercase}}/active-records
Content-Type: application/json
Authorization: {{`{{Authorization}}`}}

###
# @name LIST_FILTERED_{{$.Entity.GoSlice}}
POST {{`{{app-host}}`}}/api/{{$.ServiceName | lowercase}}/{{$.Entity.GoSlice | lowercase}}
Content-Type: application/json
Authorization: {{`{{Authorization}}`}}

{
    "limit": 0,
    "offset": 0,
    "sort": [
    {{- range $Col := $.Entity.Columns}}
    {{- if (ne $Col.GoField "Status")}}"+{{$Col.GoField}}",{{- end}}
    {{- end}}
    ]
}



###############################################################
###############################################################
# @name GET_BY_ID_{{$.Entity.GoStruct}}
GET {{`{{app-host}}`}}/api/{{$.ServiceName | lowercase}}/{{$.Entity.GoSlice | lowercase}}/{{$.Entity.GoStruct | lowercase}}/{{range $Col := $.Entity.Columns}}
{{- if $Col.IsPK}}{{`{{`}}{{$Col.GoVarName}}{{`}}`}}/{{end}}
{{- end}}
Content-Type: application/json
Authorization: {{`{{Authorization}}`}}

###
{{- range $Col := $.Entity.Columns}}
{{- if or ($Col.IsPK) }}
@{{$Col.GoVarName}} = {{$Col.GoVarName}}
{{- else if (eq $Col.GoField "UpdatedAt")}}
@{{$Col.GoVarName}} = {{`{{GET_BY_ID_`}}{{$.Entity.GoStruct}}{{`.response.body.$.`}}{{$Col.GoField}}{{`}}`}}
{{- end}}
{{- end}}

###
# @name INSERT_{{$.Entity.GoStruct}}
POST {{`{{app-host}}`}}/api/{{$.ServiceName | lowercase}}/{{$.Entity.GoSlice | lowercase}}/{{$.Entity.GoStruct | lowercase}}
Content-Type: application/json
Authorization: {{`{{Authorization}}`}}

{
{{- range $Col := $.Entity.Columns}}
{{- if and (not $Col.IsPK) (ne $Col.GoField "Status") (ne $Col.GoField "UpdatedAt") (ne $Col.GoField "CreatedAt")}}
    "{{$Col.GoField}}": {{template "datafield" $Col}},
{{- end}}
{{- end}}
}


###
# @name EDIT_{{$.Entity.GoStruct}}
PUT  {{`{{app-host}}`}}/api/{{$.ServiceName | lowercase}}/{{$.Entity.GoSlice | lowercase}}/{{$.Entity.GoStruct | lowercase}}
Content-Type: application/json
Authorization: {{`{{Authorization}}`}}

{
{{- range $Col := $.Entity.Columns}}
{{- if and (ne $Col.GoField "Status") (ne $Col.GoField "CreatedAt")}}
    "{{$Col.GoField}}": {{template "datafield" $Col}},
{{- end}}
{{- end}}
}


###
# @name ACTIVATE_{{$.Entity.GoStruct}}
PUT  {{`{{app-host}}`}}/api/{{$.ServiceName | lowercase}}/{{$.Entity.GoSlice | lowercase}}/{{$.Entity.GoStruct | lowercase}}/activate
Content-Type: application/json
Authorization: {{`{{Authorization}}`}}

{
{{- range $Col := $.Entity.Columns}}
{{- if or ($Col.IsPK) (eq $Col.GoField "UpdatedAt")}}
    "{{$Col.GoField}}": {{template "datafield" $Col}},
{{- end}}
{{- end}}
}


###
# @name INACTIVATE_{{$.Entity.GoStruct}}
PUT  {{`{{app-host}}`}}/api/{{$.ServiceName | lowercase}}/{{$.Entity.GoSlice | lowercase}}/{{$.Entity.GoStruct | lowercase}}/inactivate
Content-Type: application/json
Authorization: {{`{{Authorization}}`}}

{
{{- range $Col := $.Entity.Columns}}
{{- if or ($Col.IsPK) (eq $Col.GoField "UpdatedAt")}}
    "{{$Col.GoField}}": {{template "datafield" $Col}},
{{- end}}
{{- end}}
}

###
# @name SOFT-DELETE_{{$.Entity.GoStruct}}
PUT  {{`{{app-host}}`}}/api/{{$.ServiceName | lowercase}}/{{$.Entity.GoSlice | lowercase}}/{{$.Entity.GoStruct | lowercase}}/soft-delete
Content-Type: application/json
Authorization: {{`{{Authorization}}`}}

{
{{- range $Col := $.Entity.Columns}}
{{- if or ($Col.IsPK) (eq $Col.GoField "UpdatedAt")}}
    "{{$Col.GoField}}": {{template "datafield" $Col}},
{{- end}}
{{- end}}
}

###
# @name HARD-DELETE_{{$.Entity.GoStruct}}
PUT  {{`{{app-host}}`}}/api/{{$.ServiceName | lowercase}}/{{$.Entity.GoSlice | lowercase}}/{{$.Entity.GoStruct | lowercase}}/hard-delete
Content-Type: application/json
Authorization: {{`{{Authorization}}`}}

{
{{- range $Col := $.Entity.Columns}}
{{- if or ($Col.IsPK) (eq $Col.GoField "UpdatedAt")}}
    "{{$Col.GoField}}": {{template "datafield" $Col}},
{{- end}}
{{- end}}
}
