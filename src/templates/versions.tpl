<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Stdk / Versions</title>
</head>
<body>
    <a href="/"><< Back</a>
    <br /><br />

    <table>
        <tr>
            <th>Id</th>
            <th>Settings</th>
            <th>IsActive</th>
            <th>CreatedAt</th>
            <th>-</th>
        </tr>

        {{range .versions}}
            <tr>
                <td>{{.Id}}</td>
                <td>{{.Settings}}</td>
                <td>{{.IsActive}}</td>
                <td>{{.CreatedAt}}</td>
                <td>
                    <form action="/versions/{{.Id}}/setActive" method="POST">
                        <input type="submit" value="Make Active" />
                    </form>
                </td>
            </tr>
        {{end}}
    </table>

    <form action="/versions/{{.moduleId}}" method="POST" enctype="multipart/form-data">
        <div>
            <h3>Add new version</h3>
        </div>

        <div>
            <div><label>Settings (valid JSON)</label></div>
            <textarea name="settings"></textarea>
        </div>

        <div>
            <div><label>Binary</label></div>
            <input type="file" name="file" />
        </div>

        <div>
            <input type="submit" value="Create" />
        </div>
    </form>

</body>
</html>
