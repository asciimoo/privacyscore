{{ define "content" }}
<form method="get" action="./check">
    <div class="row" id="main_form">
        <div class="column">
            <label for="url"><h2>Enter URL below to check it's privacy score</h2></label>
            <input type="text" placeholder="Enter url" name="url" id="url" />
            <input class="button-primary" type="submit" value="Check" />
        </div>
    </div>
</form>
{{ end }}
