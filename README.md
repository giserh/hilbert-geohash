# Hilbert-Geohash [![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

### Dependencies:
+ npm
+ webpack
+ go

### Install:
```bash
npm install
go get ./...
```

### Run:
```bash
go run app.go
```

### Generate Z-order/Hilbert comparison stats:
```bash
go run geohash/stat.go all 1000
```

A live demo can be seen [here](http://hilbert-geohash.herokuapp.com/)
