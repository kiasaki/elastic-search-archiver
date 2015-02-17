# Elastic Search Archiver

This little command line tool is built to help you archive old logs in Elastic
Search, most likely logs Logstash put in there.

It will go and **flush indices transactions** then proceed to **close** those indices
so that they don't cost you CPU or memory, all this up to a certain duration,
the default is 2 weeks in the past. On top of that it has the ability to completly
delete (forever) indices up to a certain duration, default is not deletion.

## Usage

```
Help
```

## Developing

```bash
go get .
go run *.go
go build -o elastic-search-archiver .
```

## Liscence

(MIT)

Copyright (c) 2015 Frederic Gingras (frederic@gingras.cc)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
