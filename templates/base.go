// Copyright (c) 2017 Sagar Gubbi. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const baseSrc = `<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" type="text/css" href="/static/css/orangeforum.css?v=140">
	<title>
		{{ if .Common.PageTitle }}
			{{ .Common.PageTitle }}
		{{ else }}
			{{ .Common.ForumName }}
		{{ end }}
	</title>
	{{ block "head" . }}{{ end }}
</head>

<body>
	<div id="container">
		<div id="header" class="clearfix">
			<div id="navleft">
				<a href="/">{{ .Common.ForumName }}</a>{{ if .GroupName }} &gt; <a href="/groups?name={{ .GroupName }}">{{ .GroupName }}</a>{{ end }}
			</div>
			<div id="navright">
				{{ if .Common.UserName }}
				<a href="/users?u={{ .Common.UserName }}">{{ .Common.UserName }}{{ if .Common.IsNotification }}<span class="alert">&#x2757</span>{{ end }}</a>
				{{ else }}
				<a href="/login?next={{ .Common.CurrentURL }}">Login</a>
				{{ end }}
			</div>
		</div>
		<hr>
		<div id="content">
		{{ block "content" . }}{{ end }}
		</div>
		<div id="footer">
		{{ range $i, $e := .Common.ExtraNotesShort }}
			{{ if $i }}&middot;{{ end }}
			<a href="/note?id={{ $e.ID }}">{{ $e.Name }}</a>
		{{ end }}
		</div>
	</div>
	<script src="/static/js/orangeforum.js?v=140"></script>
	{{ .Common.BodyAppendage }}
</body>
</html>`
