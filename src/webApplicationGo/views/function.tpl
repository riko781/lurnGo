<!DOCTYPE html>
<html>
    {{.HtmlHead}}
    <body>
        {{.Header}}
        {{.Section}}
        <article>
            <h2>Fonction</h2>
            <div>
                <h3>Fonction continue</h3>
                <p>
                        func name(nom string) string {
                            return nom
                        }     
                        <br>
                        name("Carneva")
                        <br> retourn {{.name}}
                </p>
            </div>
            <div>
                <h3>Fonction plusieurs résultats</h3>
                <p>
                        func nameMultiple(nom, prenom string) (string, string) {
                            return nom, prenom
                        }
                        <br>
                        nom, prenom = nameMultiple("Carneva", "henrick")
                        <br> retourn {{.nom}} {{.prenom}}
                </p>
            </div>
        </article>
        {{.Footer}}
    </body>
</html>