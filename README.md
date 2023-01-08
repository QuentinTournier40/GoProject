# Projet

Projet Go réalisé durant l'UE Architectures distribuées.
Projet visant à découvrir et à approfondir nos connaissances sur la technologie Go et le mécanisme de publication de messages et d’abonnement (publish-subscribe).

## API Documentation

#### Comment générer la Documentation Swagger ?
```cmd
go install github.com/parvez3019/go-swagger3@latest
go-swagger3 --module-path . --main-file-path ./api/api/api.go --output ./api/spec/open-api-spec.yaml --schema-without-pkg --generate-yaml true
```

Importer le contenu du fichier [***/api/spec/open-api-spec.yaml***](https://github.com/QuentinTournier40/GoProject/blob/main/api/spec/open-api-spec.yaml)  ici : https://editor.swagger.io/


## Lancer notre projet

* Lancer un server Redis
* Exécuter **start.sh** et suivre les indications
* Pour lancer le **front**, ouvrir le fichier [***/website/index.html***](https://github.com/QuentinTournier40/GoProject/blob/main/website/index.html)
* L'utilisation du front nécessite d'avoir lancer le server Redis et l'API ...


## 🛠 Compétences
[GoLang](https://go.dev/doc/)
[RediGo](https://github.com/gomodule/redigo)
[Gorilla Mux](https://github.com/gorilla/mux)
[GoCron](https://github.com/go-co-op/gocron)
[Swagger](https://swagger.io/)
[GoPaho](https://github.com/eclipse/paho.mqtt.golang)
[JavaScript](https://developer.mozilla.org/fr/docs/Web/JavaScript)
[HTML](https://developer.mozilla.org/fr/docs/Web/HTML)
[CSS](https://developer.mozilla.org/fr/docs/Web/CSS)


## Auteurs

- [Belicaud Louan](https://github.com/louanbel)
- [Ghoniem Younes](https://github.com/Dhoulnoun)
- [Marche Jules](https://github.com/julesmarche)
- [Tournier Quentin](https://github.com/QuentinTournier40)


![Logo](https://www.imt-atlantique.fr/sites/default/files/ecole/IMT_Atlantique_logo.png)

