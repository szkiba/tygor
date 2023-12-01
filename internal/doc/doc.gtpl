{{- $title := list "k6/x" .Name | join "/" -}}
{{- if .GitHub.Repo -}}
{{- $title = .GitHub.RepoName -}}
{{- end -}}

{{ h 1 }} {{ $title }}

{{ doc .Namespace }}
{{ template "example" example .Namespace }}

{{ template "github_related" . }}

{{ h 1 }} API

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
{{- /* title */ -}}
{{- define "github_related" -}}
{{  if .GitHub.Repo }}
{{   if .GitHub.Examples }}
The [examples](https://github.com/{{.GitHub.Repo}}/blob/master/examples) directory contains examples of how to use the {{.GitHub.RepoName}} extension.
A k6 binary containing the {{.GitHub.RepoName}} extension is required to run the examples.
*If the search path also contains the k6 command, don't forget to specify which k6 you want to run (for example `./k6`)*.
{{   end }}
{{   if or .GitHub.Releases .GitHub.Packages }}
**Download**

{{     if .GitHub.Releases -}}
You can download pre-built k6 binaries from the [Releases](https://github.com/{{.GitHub.Repo}}/releases/) page.
{{-     end }}{{     if .GitHub.Packages }}
 Check the [Packages](https://github.com/{{.GitHub.Repo}}/pkgs/container/{{.GitHub.RepoName}}) page for pre-built k6 Docker images.
{{-     end -}}
{{   end }}

<details>
<summary><strong>Build</strong></summary>

The [xk6](https://github.com/grafana/xk6) build tool can be used to build a k6 that will include {{.GitHub.RepoName}} extension:

```bash
$ xk6 build --with github.com/{{.GitHub.Repo}}@latest
```

For more build options and how to use xk6, check out the [xk6 documentation]([xk6](https://github.com/grafana/xk6)).

</details>

{{   end }}
{{- end -}}
{{- /*--------------- template definitions ---------------*/ -}}
