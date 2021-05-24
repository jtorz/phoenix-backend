package {{$.ServiceAbbr | lowercase}}model

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

type {{$.Entity.GoStruct}} struct{
{{- range $Col := $.Entity.Columns}}
	{{- if ne $Col.GoField "Status"}}
		{{$Col.GoField}}  {{$Col.GoDataType}} `rql:"filter,sort,column={{$Col.DBName}}"`
	{{- end}}
{{- end}}
	Status base.Status
	base.RecordActions
}

