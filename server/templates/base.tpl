{{ define "base" }}
<html>
<head>
    <title>Privacy Score</title>
    <link rel="stylesheet" href="./static/milligram.min.css" />
</head>
<body>
    <main class="wrapper">
        <section class="container">
        {{ template "content" . }}
        </section>
    </main>
</body>
</html>
{{ end }}
