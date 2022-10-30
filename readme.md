

# Privy Test Code By Vandy Ahmad

## Before Run You Must Install :
1. Install Golang Migrate with command 
    ```go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest```
2. Install Docker
3. Install Jaeger with command
    ``` docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp  -p 6831:6831/udp   -p 6832:6832/udp -p 5778:5778  -p 16686:16686  -p 14268:14268 -p 14250:14250 -p 9411:9411 jaegertracing/all-in-one:1.20 ```
4. Setting .env with your local enviroment
5. Install make file if you use linux already bundled
    In Windows you can install with choco command :
    ``` choco install make ```

## How to run

1. First You Must Run Migration
    ``` make run-migration ```
2. Running Program with command
    ``` make run ```

### API Documentations

1. https://www.getpostman.com/collections/45e48f7893fc4710b36f