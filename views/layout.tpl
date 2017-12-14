<!DOCTYPE html>
<html>

<head>
{{template "head.html" .}}
</head>
<body>
{{template "nav.tpl" .}}
<div class="container">
{{.LayoutContent}}
</div>
{{template "footer.html"}}


</body>
</html>