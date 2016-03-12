{{ define "content" }}
<h3 class="score {{ .GetScoreName }}">{{ .Score }}/100</h3>
<h3 class="result_url">{{ .URL }}</h3>
<div class="row">
    {{ if .Penalties }}
    <div class="column">
        <table>
            <tr><th>Penalty</th><th>Value</th></tr>
                {{ range .Penalties }}<tr><td>{{ .Description }}</td><td>-{{ .Value }}</td></tr>{{ end }}
        </table>
    </div>
    {{ end }}
    {{ if .Errors }}
    <div class="column">
        <table>
            <tr><th>Errors</th></tr>
            {{ range .Errors }}<tr><td>{{ . }}</td></tr>{{ end }}
        </table>
    </div>
    {{ end }}
</div>
{{ end }}
