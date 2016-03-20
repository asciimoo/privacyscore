{{ define "content" }}
<div class="row stats">
    {{ range .Stats }}
    <div class="column stat_container">
        <div class="{{ GetScoreName .BaseScore }} stat_col" style="height: {{ statHeight .Count $.StatEntryCount }}%"> </div>
        <div class="stat_tooltip">Score {{ .Label }}:<br />{{ .Count }}</div>
    </div>
    {{ end }}
    <h4>Aggregated scores of previous checks</h4>
</div>
<form method="get" action="./check">
    <div class="row" id="main_form">
        <div class="column">
            <label for="url_input"><h2>Check your site now</h2></label>
            <input type="text" placeholder="enter url.." name="url" id="url_input" class="big_input" />
            <input class="button-primary big_input" type="submit" value="Check" />
        </div>
    </div>
</form>
{{ end }}
