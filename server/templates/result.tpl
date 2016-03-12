{{ define "content" }}
<h3 style="float: right">Score: {{ .Result.Score }}</h3>
<h3>{{ .URL }}</h3>
<ul>
    {{ range .Result.Penalties }}<li>{{ .Description }}: {{ .Value }}</li>{{ end }}
</ul>
{{ end }}
