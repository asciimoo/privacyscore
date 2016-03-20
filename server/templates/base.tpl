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
.navigation { border-bottom: 1px solid #ddd; background-color: #f4f5f6; height: 5.2rem; position: fixed; top: 0; left: 0; right: 0; z-index: 98; }
.navigation h1 { font-size: 1.4em; padding: 0.4em 0; margin: 0; }
.navigation-title, .navigation .title { display: inline; line-height: 5.2rem; margin-right: 5.0rem; }
.github-corner { border: 0; color: #f4f5f6; fill: #606c76; height: 5.2rem; position: fixed; right: 0; z-index: 99; top: 0; width: 5.2rem; }
.github-corner:hover { fill: #2980b9; }
.main { margin-top: 5.2rem; }
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
.stat_container { height: 5em; text-align: center; opacity: 0.6; position: relative; }
.stat_container:hover { opacity: 1; }
.stat_tooltip { transition: opacity 1s ease; opacity: 0; font-size: 0.7em; font-weight: 600; color: #2c3e50; position: absolute; left: 0; right: 0; top: -5.2rem; height: 5.2rem; z-index: 100; background: #f4f5f6; border: 1px solid #ddd; display: flex; align-items: center; justify-content: center; }
.stat_container:hover .stat_tooltip { transition: opacity 1s ease; opacity: 1; }
.stat_col { width: 100%; padding: 0; margin: 0; min-height: 1px; flex-direction: row; }
@media (max-width: 40.0rem) {
    .stat_container, .stat_col { display: none; }
}

td:first-child, th:first-child { padding-left: 0.6em; }
td:last-child, th:last-child { padding-right: 0.6em; }
th:last-child { text-align: right; }

.nopenalties { text-align: center; margin-top: 2em; color: #40d47e; }
.result_header { display: flex; align-items: center; margin-top: 6.2rem; }
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
                <a class="navigation-title" href="./" title="PrivacyScore">
                    <h1 class="title">PrivacyScore</h1>
                </a>
                <a class="navigation-title float-right" href="./about">About</a>
                <a class="hidden-xs github-corner" href="https://github.com/asciimoo/privacyscore" title="PrivacyScore on Github" target="blank">
                    <svg width="80" height="80" viewBox="0 0 250 250" class="github-corner">
                        <path d="M0,0 L115,115 L130,115 L142,142 L250,250 L250,0 Z"></path>
                        <path d="M128.3,109.0 C113.8,99.7 119.0,89.6 119.0,89.6 C122.0,82.7 120.5,78.6 120.5,78.6 C119.2,72.0 123.4,76.3 123.4,76.3 C127.3,80.9 125.5,87.3 125.5,87.3 C122.9,97.6 130.6,101.9 134.4,103.2" fill="currentColor" style="transform-origin: 130px 106px;" class="octo-arm"></path>
                        <path d="M115.0,115.0 C114.9,115.1 118.7,116.5 119.8,115.4 L133.7,101.6 C136.9,99.2 139.9,98.4 142.2,98.6 C133.8,88.0 127.5,74.4 143.8,58.0 C148.5,53.4 154.0,51.2 159.7,51.0 C160.3,49.4 163.2,43.6 171.4,40.1 C171.4,40.1 176.1,42.5 178.8,56.2 C183.1,58.6 187.2,61.8 190.9,65.4 C194.5,69.0 197.7,73.2 200.1,77.6 C213.8,80.2 216.3,84.9 216.3,84.9 C212.7,93.1 206.9,96.0 205.4,96.6 C205.1,102.4 203.0,107.8 198.3,112.5 C181.9,128.9 168.3,122.5 157.7,114.1 C157.9,116.9 156.7,120.9 152.7,124.9 L141.0,136.5 C139.8,137.7 141.6,141.9 141.8,141.8 Z" fill="currentColor" class="octo-body"></path>
                    </svg>
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
