<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Go</title>
         <link rel="stylesheet" type="text/css" href="/static/css/index.css">
    </head>
    <body>
        <h1>My First site with Go !</h1>
        <p>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
        <a id="linkSiteGo" href="http://{{.Website}}" target="_blank">Documention BEEGO</a>
        <a id="linkSiteGo" href="http://{{.DocGo}}" target="_blank">Documentation Go</a>
        <a id="email" href="#">{{.Email}}</a>
         <div id="operation">
            <h3>My arhytmétique operation with go!</h3><br>
            <div>
                <h4>Addition</h4>
                <p>{{.x}} + {{.y}} = {{.resultAdittion}}</p>
            </div> 
              <div>
                <h4>Soustraction</h4>
                <p>{{.x}} - {{.y}} = {{.resultSoustraction}}</p>
            </div> 
            <div>
                <h4>Multiplication</h4>
                <p>{{.x}} x {{.y}} = {{.resultMultiplication}}</p>
            </div> 
            <div>
                <h4>Division</h4>
                <p>{{.x}} / {{.y}} = {{.resultDivision}}</p>
            </div> 
        </div> 
    </body>
</html> 