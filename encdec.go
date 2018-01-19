package aqua

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/stupideity/aero/refl"
)

func encode(r []reflect.Value, typ []string) []byte {

	buf := new(bytes.Buffer)
	encd := json.NewEncoder(buf)

	for i, _ := range r {
		encodeItem(encd, r[i], typ[i])
	}

	return buf.Bytes()
}

func encodeItem(j *json.Encoder, r reflect.Value, t string) {

	switch {
	case t == "int":
		err := j.Encode(r.Int())
		if err != nil {
			panic(err)
		}
	case t == "map":
		err := j.Encode(r.Interface().(map[string]interface{}))
		if err != nil {
			panic(err)
		}
	case t == "string":
		err := j.Encode(r.String())
		if err != nil {
			panic(err)
		}
	case t == "i:.":
		s := refl.ObjSignature(r.Interface())
		err := j.Encode(s)
		if err != nil {
			panic(err)
		}
		encodeItem(j, r, s)
	case strings.HasPrefix(t, "st:"):
		err := j.Encode(r.Interface())
		if err != nil {
			panic(err)
		}
	case strings.HasPrefix(t, "sl:"):
		err := j.Encode(r.Interface())
		if err != nil {
			panic(err)
		}
	default:
		panic(fmt.Sprintf("Can't encode '%s' for endpoint cache", t))
	}
}

func decode(data []byte, typ []string) reflect.Value {
	buf := bytes.NewBuffer(data)
	decd := json.NewDecoder(buf)
	// r := make([]reflect.Value, len(typ))
	var r reflect.Value
	for i, _ := range typ {
		if typ[i] == "string" {
			r = decodeItem(decd, typ[i])
		}
	}
	return r
}

func decodeItem(j *json.Decoder, t string) reflect.Value {
	var r reflect.Value
	switch {
	case t == "int":
		var i int
		err := j.Decode(&i)
		if err != nil {
			panic(err)
		}
		r = reflect.ValueOf(i)
	case t == "map":
		var m map[string]interface{}
		err := j.Decode(&m)
		if err != nil {
			panic(err)
		}
		r = reflect.ValueOf(m)
	case t == "string":
		var s string
		err := j.Decode(&s)
		if err != nil {
			panic(err)
		}
		r = reflect.ValueOf(s)
	case t == "i:.":
		var s string
		err := j.Decode(&s)
		if err != nil {
			panic(err)
		}
		r = decodeItem(j, s)
	case strings.HasPrefix(t, "st:"):
		var m map[string]interface{}
		err := j.Decode(&m)
		if err != nil {
			panic(err)
		}
		r = reflect.ValueOf(m)
	case strings.HasPrefix(t, "sl:"):
		var a []interface{}
		err := j.Decode(&a)
		if err != nil {
			panic(err)
		}
		r = reflect.ValueOf(a)
	default:
		panic(fmt.Sprintf("Can't decdoe '%s' for endpoint cache", t))
	}

	return r
}
