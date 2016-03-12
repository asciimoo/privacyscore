{{ define "content" }}
<h3 style="float: right">Score: {{ .Score }}</h3>
<h3>{{ .URL }}</h3>
<ul>
    {{ range .Penalties }}<li>{{ .Description }}: {{ .Value }}</li>{{ end }}
</ul>
{{ end }}
