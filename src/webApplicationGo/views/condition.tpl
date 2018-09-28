<!DOCTYPE html>
<html>
    {{.HtmlHead}}
    <body>
        {{.Header}}
        {{.Section}}
        <article>
            <h2>Condition</h2>
            <div class="condition">
                <div id="if">
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
                        <div>
                            if x := (x + y); x > y { <br>
                                &nbsp;&nbsp;&nbsp;  return {{.a}} <br>
                            } else {<br>
                                &nbsp;&nbsp;&nbsp;  return {{.b}}<br>
                            }<br><br>
                        
                            retourne {{.v}}
                           </div>
                    </div>
                </div>
                <div id="switch">
                    <h3>SWITCH</h3>
                    <br><br><br>
                    <div id="condition">
                        <div>
                                switch os := runtime.GOOS; os {
                                    case "darwin":
                                      return "OS X."
                                    case "linux":
                                        return "Linux."
                                    default:
                                        return = "Other OS"
                                }
                                Retourne {{.os}}
                        </div>
                    </div>
                </div>
            </div>
        </article>
        {{.Footer}}
    </body>
</html>