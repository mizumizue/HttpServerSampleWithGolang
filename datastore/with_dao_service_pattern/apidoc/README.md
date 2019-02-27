# SwaggerUI & SwaggerEditor

## How to start

```
docker-compose up -d
```

## Start URL

SwaggerUI - http://localhost:8080/

SwaggerEditor - http://localhost:8081/

## Share local server

```
brew cask install ngrok

# Share swagger ui local server
ngrok http 8080

# Share swagger editor local server
ngrok http 8081
```

## Code generate from yaml

```
brew install swagger-codegen
```

```
# Generate code from yaml
swagger-codegen generate -l nodejs-server -i http://localhost:8080/swagger.yaml -o build/output/

# Supported code list(2019/2/21)
csharp
csharp-dotnet2
dynamic-html
html
html2
java
jaxrs-cxf-client
jaxrs-cxf
inflector
jaxrs-cxf-cdi
jaxrs-spec
jaxrs-jersey
jaxrs-di
jaxrs-resteasy-eap
jaxrs-resteasy
spring
nodejs-server
openapi
openapi-yaml
kotlin-client
kotlin-server
php
scala
scala-akka-http-server
swift3
swift4
typescript-angular
```