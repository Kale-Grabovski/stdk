<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Stdk</title>
</head>
<body>

    {{range .modules}}
        <div><a href="/versions/{{.Id}}">{{.Name}}</a></div>
    {{end}}

    <form action="/" method="POST">
        <div>
            <h3>Add new module</h3>
        </div>

        <div>
            <div><label>Name</label></div>
            <input type="text" name="name" />
        </div>

        <div>
            <input type="submit" value="Create" />
        </div>
    </form>

</body>
</html>
