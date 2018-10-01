<!DOCTYPE html>
<html>
    {{.HtmlHead}}
    <body>
        {{.Header}}
        {{.Section}}
        <article>
            <h2>Bourcle</h2>
            <div>
                <h3>For with array</h3>
                {{range $key,$val := .somme}}
                Indice : {{$key}} 
                Valeur : {{$val}} <br>
                {{end}}
            </div>
            <div>
                <h3>For with slice</h3>
                {{range $key,$val := .slice}}
                Indice : {{$key}} 
                Valeur : {{$val}} <br>
                {{end}}
                </div>
        </article>
        {{.Footer}}
    </body>
</html>