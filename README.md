# GoProject

## Auteurs: 

[Marche Jules](https://github.com/julesmarche)  
[Tournier Quentin](https://github.com/QuentinTournier40)  
[Belicaud Louan](https://github.com/louanbel)  
[Ghoniem Younes](https://github.com/Dhoulnoun)  

```cmd
mkdir output
go build -o output ./...

go-swagger3 --module-path . --main-file-path ./api/main/main.go --output oas.json --schema-without-pkg --generate-yaml true
```