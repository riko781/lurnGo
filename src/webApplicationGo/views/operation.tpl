<!DOCTYPE html>
<html>
        {{.HtmlHead}}
    <body>
        {{.Header}}
        {{.Section}}
        <article>
            <h2>My arhytmétique operation with go!</h2>
            <div id="operation">
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
        </article>
        {{.Footer}}
    </body>
</html>