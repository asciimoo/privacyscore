{{ define "content" }}
<form method="get" action="./check">
    <div class="row" id="main_form">
        <div class="column">
            <input type="text" placeholder="Enter url" name="url" />
        </div>
        <div class="column">
            <input class="button-primary" type="submit" value="Check" />
        </div>
    </div>
</form>
{{ end }}
