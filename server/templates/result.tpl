{{ define "content" }}
<div class="result_header">
    <h3 class="result_url">{{ .Result.BaseURL }}</h3>
    <h3 class="score {{ GetScoreName .Result.Penalties.GetScore }}"><span class="invisible">Score: </span>{{ .Result.Penalties.GetScore }}/100</h3>
</div>
<div class="row">
    {{ if .Result.Penalties.GetAll }}
    <div class="column">
        <table id="check_results">
            <tr><th>Penalty</th><th>Value</th></tr>
                {{ range .Result.Penalties.GetAll }}
                <tr>
                    <td>
                        <span class="penalty_desc">{{ .Description }}</span> <span class="small"><a class="penalty_link" target="_blank" href="{{ .DetailLink }}">(more info)</a></span>
                        {{ if .Notes }}<div>{{ range .Notes }} <span class="penalty_note">{{ . }}</span>{{ end }}</div>{{ end }}
                    </td>
                    <td><span class="penalty_value">-{{ .GetValue }}</span></td>
                </tr>
                {{ end }}
        </table>
    </div>
    {{ else }}
        {{ if not .Result.Errors }}
            <h3 class="nopenalties">Hurray, no penalties</h3>
        {{ end}}
    {{ end }}
    <hr />
</div>
{{ if .Result.Errors }}
<div class="row">
    <div class="column">
        <h4>Errors</h4>
        <table>
            {{ range .Result.Errors }}<tr><td>{{ . }}</td></tr>{{ end }}
        </table>
    </div>
</div>
{{ end }}
{{ if .Resources }}
<div class="row">
    <div class="column">
        <h4>Checked resources</h4>
        <ul>
        {{ range .Resources }}{{ if . }}<li>{{ .URL }}</li>{{ end }}{{ end }}
        </ul>
    </div>
</div>
{{ end }}
{{ end }}
