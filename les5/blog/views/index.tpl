<!doctype html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Главная страница блога</title>
</head>
<body>

{{range .Posts}}
<div>
    <a href="/post/{{.Id}}">{{.Title}}</a>
</div>
{{end}}
</body>
</html>