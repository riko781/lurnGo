<!DOCTYPE html>
<html>
    {{.HtmlHead}}
    <body>
        {{.Header}}
        {{.Section}}
        <article>
            <h2>Condition</h2>
            <div>
                <h3>IF</h3>
                var x uint8 ={{.x}}<br>
                var y uint8 ={{.y}}<br><br>
                <div id="condition">
                    <div>
                        if x < y { <br>
                            &nbsp;&nbsp;&nbsp;   return {{.a}} <br>
                        } else {<br>
                            &nbsp;&nbsp;&nbsp;   return {{.b}}<br>
                        }<br><br>
                        retourne {{.b}}
                    </div>
                    <div>
                         if x > y { <br>
                             &nbsp;&nbsp;&nbsp;  return {{.a}} <br>
                         } else {<br>
                             &nbsp;&nbsp;&nbsp;  return {{.b}}<br>
                         }<br><br>
                         retourne {{.a}}
                    </div>
                </div>
            </div>
        </article>
        {{.Footer}}
    </body>
</html>