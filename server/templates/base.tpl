{{ define "base" }}
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8" />
    <meta name="referrer" content="never" />
    <title>Privacy Score</title>
    <link rel="stylesheet" href="./static/milligram.min.css" />
    <style>
body { background-color: #fcfbfa; margin: 0; padding: 0; }
h2 { font-size: 5.2rem; }
a { color: #2980b9; }
.navigation { border-bottom: 1px solid #ddd; background-color: #f4f5f6; }
.navigation h1 { font-size: 1.4em; padding: 0.4em 0; margin: 0; }
#main_form { margin: 2em 0; }
#url_input { color: #606c76; background: #f4f5f6; }
input.big_input { height: 1.6em; padding: 0 0.4em; font-size: 2.2em; line-height: 1.6em; vertical-align: middle; font-weight: 300; }
input.big_input:focus { border: 0.1rem solid #2980b9; }
input[type="submit"].big_input { background-color: #2980b9 !important; border: 0.1rem solid #2980b9; }
.row { margin-left: 0; }
.invisible { display: none; }
.small { font-size: 0.8em; }
.stats { position: relative; }
.stats h4 { font-size: 1.2em; transition: opacity 1s ease; opacity: 0; position: absolute; bottom: 0; left: 0; padding: 0; margin: 0; }
.stats:hover h4 { transition: opacity 2s ease; opacity: 1; }
.stat_container { height: 5em; text-align: center; opacity: 0.6; }
.stat_container:hover { opacity: 1; background: #f4f5f6; }
.stat_col { width: 100%; padding: 0; margin: 0; min-height: 1px; flex-direction: row; }
@media (max-width: 40.0rem) {
    .stat_container, .stat_col { display: none; }
}

td:first-child, th:first-child { padding-left: 0.6em; }
td:last-child, th:last-child { padding-right: 0.6em; }
th:last-child { text-align: right; }

.nopenalties { text-align: center; margin-top: 2em; color: #40d47e; }
.result_header { display: flex; align-items: center; margin-top: 1em; }
.result_url { text-align: center; flex-grow: 1; word-wrap: break-word; }
.score { color: #fcfbfa; border-radius: 50%; height: 4em; width: 4em; line-height: 4em; text-align: center; flex-shrink: 0; }
.penalty_link { white-space: nowrap; }
.penalty_desc { font-size: 1.2em; }
.penalty_note { background: #606c76; color: #fcfbfa; font-size: 0.7em; border-radius: 0.4em; padding: 2px 6px; }
.penalty_value { font-size: 1.7em; font-weight: 600; display: block; text-align: right; }
.good { background-color: #40d47e; }
.medium { background-color: #f1c40f; }
.bad { background-color: #e74c3c; }
.poor { background-color: #2c3e50; }
input[type="submit"] { float: right; }
    </style>
</head>
<body>
    <main class="wrapper">
        <nav class="navigation">
            <section class="container">
                <a class="navigation-title" href="./">
                    <h1 class="title">PrivacyScore</h1>
                </a>
            </section>
        </nav>

        <section class="container main">
        {{ template "content" . }}
        </section>
    </main>
</body>
</html>
{{ end }}
