{{ define "content" }}
<div class="result_header">
    <h3 class="result_url">{{ .BaseURL }}</h3>
    <h3 class="score {{ GetScoreName .Penalties.GetScore }}"><span class="invisible">Score: </span>{{ .Penalties.GetScore }}/100</h3>
</div>
<div class="row">
    {{ if .Penalties.GetAll }}
    <div class="column">
        <table>
            <tr><th>Penalty</th><th>Value</th></tr>
                {{ range .Penalties.GetAll }}
                <tr>
                    <td>
                        <span class="penalty_desc">{{ .Description }}</span> <span class="small"><a class="penalty_link" target="_blank" href="{{ .DetailLink }}">(more info)</a></span>
                        {{ if .Notes }}<div>{{ range .Notes }} <span class="penalty_note">{{ . }}</span>{{ end }}</div>{{ end }}
                    </td>
                    <td><span class="penalty_value">-{{ .GetValue }}</span></td>
                </tr>
                {{ end }}
        </table>
    {{ else }}
        {{ if not .Errors }}
            <h3 class="nopenalties">Hurray, no penalties</h3>
        {{ end}}
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
