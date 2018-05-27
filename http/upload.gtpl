<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<form action="/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="uploadFile">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="submit">
</form>
</body>
</html>
