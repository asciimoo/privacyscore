{{ define "content" }}
<div class="result_header">
    <h3 class="result_url">{{ .URL }}</h3>
    <h3 class="score {{ .GetScoreName }}"><span class="invisible">Score: </span>{{ .Score }}/100</h3>
</div>
<div class="row">
    <div class="column">
    {{ if .Penalties }}
        <table>
            <tr><th>Penalty</th><th>Value</th></tr>
                {{ range .Penalties }}
                <tr>
                    <td>
                        <span class="penalty_desc">{{ .Description }}</span> <span class="small"><a class="penalty_link" target="_blank" href="{{ .DetailLink }}">(more info)</a></span>
                        {{ if .Notes }}<div>{{ range .Notes }} <span class="penalty_note">{{ . }}</span>{{ end }}</div>{{ end }}
                    </td>
                    <td><span class="penalty_value">-{{ .Value }}</span></td>
                </tr>
                {{ end }}
        </table>
    {{ else }}
        <h3 class="nopenalties">Hurray, no penalties</h3>
    {{ end }}
    </div>
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
