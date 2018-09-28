<!DOCTYPE html>
<html>
    {{.HtmlHead}}
    <body>
        {{.Header}}
        {{.Section}}
        <article>
            <h2>Type and déclaration des variables </h2>
            <div id="variable">
                <div>
                    <h3>Boolean types</h3>
                    <p>
                        var a Bool = true <br>
                        Retourn : {{.a}} <br>  
                    </p>
                    <p> 
                        var a Bool = false <br>
                        Retourn : {{.b}}
                    </p>
                </div>
                <div>
                    <h3>String types</h3>
                    <p>
                        var c string = "Je suis un string" <br>
                        Retourn : {{.c}} <br>
                    </p>
                </div>
                <div>
                    <h3>Integer types</h3>
                    <p>
                        var d int8 = 1 <br>
                        Retourn : {{.d}} <br>
                    </p>
                    <p>
                        var e int8 = -1 <br>
                        Retourn : {{.e}} <br>
                    </p>
                    <p>
                        var f  = float32(-1.22) <br>
                        Retourn : {{.f}} <br>
                    </p>
                    <p>
                        var g  = float32(1.22) <br>
                        Retourn : {{.g}} <br>
                    </p>
                </div>
                <div>
                    <h3>Array types</h3>
                    <p>
                       var h = [2]int8{1, 2}<br>
                        Retourn : {{.h}} <br>
                    </p>
                    <p>
                        var i = [2]string{"Hello", "World"}<br>
                        Retourn : {{.i}} <br>
                    </p>
                    <p>
                        j = [2]float32{1.0, 2.0} <br>
                        Retourn : {{.j}} <br>
                    </p>
                </div>
            </div>
        </article>
        {{.Footer}}
    </body>
</html>