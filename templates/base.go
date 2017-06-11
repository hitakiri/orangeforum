package templates

const baseSrc = `<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
	html, body {
		margin: 0;
		padding: 0;
		height: 100%;
	}
	#container {
		max-width: 960px;
		line-height: 1.58;
		margin: 0 auto;
		min-height: 100%;
		position: relative;
	}
	#header {
	}
	#content {
		padding: 10px;
		padding-bottom: 60px;
	}
	#footer {
		position: absolute;
		bottom: 0;
		width: 100%;
		height: 40px;
		text-align: center;
	}
	body {
		font-family: Arial, "Helvetica Neue", Helvetica, sans-serif;
		text-rendering: optimizeLegibility;
		-webkit-font-smoothing: antialiased;
	}
        img {
        	max-width: 100%;
        	display: block;
        	margin: 0 auto;
        }
        a {
		text-decoration: none;
        }
        a:link {
		color: #07C;
        }
        a:hover, a:active {
		color: #3af;
        }
        a:visited {
		color: #005999;
        }
        a.nodec, a.nodec:link, a.nodec:hover, a.nodec:active, a.nodec:visited {
        	color: #000;
        }
	</style>
	<title>{{ .Title }}</title>
	{{ block "head" . }}{{ end }}
</head>

<body>
	<div id="container">
		<div id="header"></div>
		<div id="content">
		{{ block "content" . }}{{ end }}
		</div>
		<div id="footer">
		{{ range $i, $e := .ExtraNotesShort }}
		{{ if $i }}&middot;{{ end }}
		<a href="/note?id={{ $e.ID }}">{{ $e.Name }}</a>
		{{ end }}
		</div>
	</div>
</body>
</html>`