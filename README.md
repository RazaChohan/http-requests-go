# Using the tool

####  Clone this repository

```bash
$ git clone https://github.com/RazaChohan/http-requests-go.git
```

#### Build

```bash
go build myhttp.go
```

#### Using the tool

```bash
./myhttp -parallel <no. of threads> <urls>
```

#### Examples

```bash
./myhttp -parallel 2 www.google.com www.yahoo.com www.hotmail.com www.github.com

./myhttp http://www.adjust.com www.google.com

./myhttp -parallel 4 www.google.com www.yahoo.com www.hotmail.com www.github.com
```


#### Running test cases

```bash
go test --run ''

```


