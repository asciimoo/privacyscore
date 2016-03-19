{{ define "content" }}
<form method="get" action="./check">
    <div class="column">
        <h2>Something went wrong</h2>
        <h5>{{ .Error }}</h5>
        <a href="./">back</a>
    </div>
</form>
{{ end }}
