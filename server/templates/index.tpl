{{ define "content" }}
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
