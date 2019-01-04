# Util and library to process json stream

This util may be helpful to analyze json stream by filtering necessary elements and pass them to next pipeline stage.

Features:
* filtering by path
    * string attributes
    * numeric attributes
    * boolean attributes
    * array elements (using "one of" logic)
* extracting (maybe coming soon)
* merging (maybe coming soon)
* I/O
    * Input
        * stdin
        * file (maybe coming soon)
    * Output
        * stdout
        * file (maybe coming soon)

## Install
As binary util:
```bash
$ go build -o ./jsonstream github.com/shnellpavel/json-stream/jsonstream-cli
$ ./jsonstream
usage: jsonstream [<flags>] <command> [<args> ...]

Utils to process and analyze stream of json

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  filter --condition=CONDITION
    Filters json stream by conditions

$ ./jsonstream filter --condition="x = y" < stream.file.json
OR
$ cat stream.file.json | ./jsonstream filter --condition="x = y"
```

As library:
```bash
$ go get -u -t github.com/shnellpavel/json-stream/jsonstream/filter
```

## Filtering by path

Supported compare operations:
* All types
    * = (Equals)
    * != (Not equals)
* Numbers
    * < (Less than)
    * <= (Less than or equal)
    * \> (Greater than)
    * \>= (Greater than or equal)
* Strings
    * < (Less than)
    * <= (Less than or equal)
    * \> (Greater than)
    * \>= (Greater than or equal)

### Examples
Input (tmp.stream.json):
```json
{"id": 1, "name": "John", "emails": ["john@gmail.com", "john@mail.ru"], "children": [{"name": "Alex", "age": 10}, {"name": "Jinny", "age": 5}], "job": {"company": "Some firm"}}
{"id": 2, "name": "Jack", "emails": ["jack@gmail.com"], "children": [], "job": {"company": "Another some firm"}}
{"id": 3, "name": "Ann", "emails": ["ann@gmail.com"], "children": [{"name": "Pit", "age": 8}], "job": {"company": "Some firm"}}
{"error": "Some error msg"}
```

#### Filter by 1st level fields
Command: 
``` bash
$ cat tmp.stream.json | jsonstream filter --condition="id = 3"
```

Output:
```json
{"id": 3, "name": "Ann", "emails": ["ann@gmail.com"], "children": [{"name": "Pit", "age": 8}], "job": {"company": "Some firm"}}
```

#### Filter by fields in nested objects
Command: 
```bash
$ cat tmp.stream.json | jsonstream filter --condition="job.company != 'Some firm'"
```

Output:
```json
{"id": 2, "name": "Jack", "emails": ["jack@gmail.com"], "children": [], "job": {"company": "Another some firm"}}
```

#### Filter by array elements
Command: 
```bash
$ cat tmp.stream.json | jsonstream filter --condition="emails = john@mail.ru"
```

Output:
```json
{"id": 1, "name": "John", "emails": ["john@gmail.com", "john@mail.ru"], "children": [{"name": "Alex", "age": 10}, {"name": "Jinny", "age": 5}], "job": {"company": "Some firm"}}
```

#### Filter by objects - elements of arrays
Command: 
```bash
$ cat tmp.stream.json | jsonstream filter --condition="children.age > 5"
```

Output:
```json
{"id": 1, "name": "John", "emails": ["john@gmail.com", "john@mail.ru"], "children": [{"name": "Alex", "age": 10}, {"name": "Jinny", "age": 5}], "job": {"company": "Some firm"}}
{"id": 3, "name": "Ann", "emails": ["ann@gmail.com"], "children": [{"name": "Pit", "age": 8}], "job": {"company": "Some firm"}}
```

## Performance

```
goos: linux
goarch: amd64
pkg: github.com/shnellpavel/json-stream/jsonstream/filter
BenchmarkProcessElem_200B_FirstLevel-8    	  200000	      5972 ns/op	    2016 B/op	      47 allocs/op
BenchmarkProcessElem_200B_NestedField-8   	  200000	      6494 ns/op	    2176 B/op	      54 allocs/op
BenchmarkProcessElem_16KB_FirstLevel-8    	    5000	    334389 ns/op	   82805 B/op	    1297 allocs/op
BenchmarkProcessElem_16KB_NestedField-8   	    5000	    344009 ns/op	   85317 B/op	    1450 allocs/op
```