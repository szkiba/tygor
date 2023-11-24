{{ h 1 }} k6/x/{{ .Name }}

{{ doc .Namespace }}

{{- /* classes and interfaces */ -}}
{{ range $parent := (concat .Classes .Interfaces) }}

{{ h 2 }} {{ $parent.Name }}

{{  doc $parent }}
{{  template "example" example $parent }}
{{- /* constructors */ -}}
{{  range $ctor := $parent.Constructors}}
{{   template "method" dict "Method" $ctor "Parent" $parent }}
{{  end }}
{{- /* properties */ -}}
{{  range $property := $parent.Properties}}
{{   template "property" dict "Property" $property "Parent" $parent }}
{{  end }}
{{- /* methods */ -}}
{{  range $method := $parent.Methods}}
{{   template "method" dict "Method" $method "Parent" $parent }}
{{  end}}
{{ end}}

{{- /* variables */ -}}
{{  range $variable := .Variables }}
{{   template "variable" . }}
{{  end }}
{{- /* functions */ -}}
{{  range $function := .Functions }}
{{   template "function" . }}
{{  end}}

{{- /*=============== template definitions ===============*/ -}}
{{- /* source */ -}}
{{- define "source" -}}
{{   if . }}
```ts
{{ . }}
```
{{   end }}
{{- end -}}
{{- /* example */ -}}
{{- define "example" }}
{{   if . }}
{{     if eq format "HTML"}}
*Example*
{{.}}
{{     else }}
<details><summary><em>Example</em></summary>

{{.}}

</details>
{{     end }}
{{   end }}
{{- end -}}
{{- /* property */ -}}
{{- define "property" -}}
{{ h 3 }} {{.Parent.Name}}.{{.Property.Name}}

{{   template "source" .Property.Source}}
{{   doc .Property }}
{{  template "example" example .Property }}
{{- end -}}
{{- /* method */ -}}
{{- define "method" -}}
{{ h 3 }} {{.Parent.Name}}{{if .Method.Name}}.{{end}}{{.Method.Name}}()

{{   template "source" .Method.Source}}
{{   $params := index .Method.Tags "param" }}
{{   range $params }}
{{     $parts := splitn " " 2 . }}
- `{{ $parts._0  }}` {{ $parts._1 }}
{{   end }}
{{   doc .Method }}
{{   $returns := index .Method.Tags "returns"}}
{{   if not (empty $returns) }}
*Returns {{ join " " $returns }}*
{{   end }}
{{   template "example" example .Method }}
{{- end -}}
{{- /* variable */ -}}
{{- define "variable" -}}
{{   if not .Modifiers.Default }}
{{ h 3 }} {{.Name}}

{{    template "source" .Source}}
{{    doc . }}
{{    template "example" example . }}
{{   end }}
{{- end -}}
{{- /* function */ -}}
{{- define "function" -}}
{{   if not .Modifiers.Default }}
{{ h 3 }} {{.Name}}()

{{    template "source" .Source}}
{{    doc . }}
{{    template "example" example . }}
{{   end }}
{{- end -}}
{{- /*--------------- template definitions ---------------*/ -}}
