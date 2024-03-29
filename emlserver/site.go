package main

const sitehtml = `<!doctype html>
<html>
{{define "items"}}
<div class="list-group">
    {{range .All}}
    {{template "item" .}}
    {{- end}}
</div>
{{end}}

{{define "item"}}
<a class="list-group-item list-group-item-action{{if .Body}} active{{end}}" href="?id={{.ID}}" style="font-size: 80%;">
    <div>
        &#9872;
        {{.To}}
        <br/>
        &#8701;
        {{.From}}
        <br/>
        &#8599;
        {{.Date}}
    </div>
    <div><b>{{.Subject}}</b></div>
</a>
{{end}}

{{define "body"}}{{if .}}
<div class="alert alert-dark" role="alert">
    &#9872;
    {{.To}}
    &#8701;
    {{.From}}
    &#8599;
    {{.Date}}
    <br />
    Betreff: <b>{{.Subject}}</b>
</div>
<pre style="font-size: 72%; white-space: pre-wrap;">{{.Body}}</pre>
<p>
    <a href="?download={{.ID}}" class="btn btn-primary btn-sm" title="{{.ID}}">Download Outlook *.eml</a>
</p>
{{end}}{{end}}

<head>
    <title>Generic Mailbox</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO"
        crossorigin="anonymous">
</head>

<body>
    <div class="container-fluid">
        <div class="row">
            <div class="col-3">{{template "items" .}}</div>
            <div class="col-9">{{template "body" .Selected}}</div>
        </div>
    </div>
</body>

</html>`
