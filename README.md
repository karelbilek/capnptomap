capnptomap
===

capnptomap is a go module that converts capnproto struct to go map[string]any{}, useful mainly for later converting to JSON or YAML for debugging capnproto outputs.

The interface is simple - 1 function - and the code is almost verbatim copied from pogs from go-capnp.

Unlike pogs, this is one-way only; it reads capnproto struct and makes a map, not the other way around.

It doesn't support interfaces and anypointers, simply because I don't understand them.

License - MIT, as go-capnp