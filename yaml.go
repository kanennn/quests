package main

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
)

// yaml for personal use, not exactly standard yaml

// may should use a yaml library like the json library in the future

// yams
// delicious

func MakeYamsFromTask(t Task) []byte {
	rv := reflect.ValueOf(t)
	rt := reflect.TypeOf(t)

	var lines [][]byte

	for i := 0; i < rv.NumField(); i++  {
		fv := rv.Field(i)
		ft := rt.Field(i)

		if ft.Tag.Get("yaml") != "" {
			key := ft.Tag.Get("yaml")
			value := fv.Interface().(string)
			lines = append(lines, []byte(fmt.Sprintf(`%s: %s`, key, value)))
		}

		
	}
	
	yams := bytes.Join(lines, []byte("\n"))
	return yams
}

func ReadYamsFromFrontatter(file []byte) (t Task) {
	//literally no clue if this will work
	re, err := regexp.Compile("---\\n([\\S\\s]+?)\n---")
	Check(err)
	yaml := re.Find(file)
	rvp := reflect.ValueOf(&t)
	rv := rvp.Elem()
	rt := reflect.TypeOf(t)
	


    for i := 0; i < rv.NumField(); i++ {
        fv := rv.Field(i)
		ft := rt.Field(i)
		
		if ft.Tag.Get("yaml") != "" {
			l := ft.Tag.Get("yaml")
			re, err := regexp.Compile(fmt.Sprintf(`%s:\s*(.*?)\n`, l))
			Check(err)

			fv.Set(reflect.ValueOf(string(re.FindSubmatch(yaml)[1])))
		}
	}
	return t
}
