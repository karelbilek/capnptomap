capnptomap
===

capnptomap is a go module that converts capnproto struct to go map[string]any{}, useful mainly for later converting to JSON or YAML for debugging capnproto outputs.

The interface is simple - 1 function - and the code is almost verbatim copied from pogs from go-capnp.

Unlike pogs, this is one-way only; it reads capnproto struct and makes a map, not the other way around.

It doesn't support interfaces and anypointers, simply because I don't understand them.

Usage 
---

```go
import (
    "generatedcode"

    "capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/schemas"
    
    "github.com/karelbilek/capnptomap"

    "gopkg.in/yaml.v3"
)

func main() {
    // register the schema
    generatedcode.RegisterSchema(schemas.DefaultRegistry)

    // get the bytes, read them, parse them
    b := getBytesFromSomewhere()
    
    protoMessage, err := capnp.Unmarshal(someBytes)
    if err != nil {
        panic(err)
    }
    protoFoo, err := generatedcode.ReadFoo(protoMessage)
    if err != nil {
        panic(err)
    }
    
    // this is the actual call
    mp, err := capnptomap.ToMap(schemas.DefaultRegistry, generatedcode.Foo_TypeID, capnp.Struct(protoFoo))
    if err != nil {
        panic(err)
    }

    // for example convert to yaml and print

    yml, err := yaml.Marshal(mp)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(yml))
}
```

License - MIT, as go-capnp