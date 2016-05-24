# website
SmartAs website


### Setup

1. Get source code:

```
go get github.com/smart-as/website
```

2. Get `doc` and `blog` to build contents:

```
cd $GOPATH/src/github.com/beego/website
git clone https://github.com/smart-as/doc
git clone https://github.com/smart-as/blog
```

3. Compile:

```
go build main.go
```

run `website` and visit `http://localhost:8888`
