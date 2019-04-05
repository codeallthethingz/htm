# htm

## Install

Install go. 

https://golang.org/doc/install

Check out the project in to the go home directory (not some other directory as Go expects files to live here)


```
cd $GOPATH
mkdir src/github.com/codeallthethingz
cd src/github.com/codeallthethingz
git clone git@github.com:codeallthethingz/htm
```

Get the dependencies

```
dep ensure
```

Usage

Run the go server

```
gin run ./...
```

Run the UI 

```
cd visualize
npm install
parcel static/index.html
```

Go to http://localhost:1234
