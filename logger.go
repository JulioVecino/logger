package logger

import (
    "time"
    "encoding/json"
    "encoding/xml"
    "reflect"
    "github.com/gookit/color"
)

var formatDateTime = "2006-01-02 15:04:05"
var decorator = "----------------------------------"

func Err(errs ...interface{}) {
    now := time.Now()
    for _, err := range errs {
        color.Error.Printf("ERROR: %s : %s\n", now.Format(formatDateTime), err)
    }
}

func Warn(texts ...interface{}) {
    now := time.Now()
    for _, text := range texts {
        color.Warn.Printf("WARN : %s : %s\n",now.Format(formatDateTime), text)
    }
}

func Tittle(texts ...interface{}) {
    now := time.Now()
    for _, text := range texts {
        color.Info.Printf("INFO : %s : %s %s %s\n",now.Format(formatDateTime),decorator, text, decorator)
    }
}

func Var(texts ...interface{}) {
    now := time.Now()
    for i, text := range texts {
        if (i % 2 == 0) {
            color.Info.Printf("INFO : %s : %s = ",now.Format(formatDateTime), text )
        } else {
            color.Info.Printf("%s\n", text )
        }
    }
}

func Info(texts ...interface{}) {
    now := time.Now()
    for _, text := range texts {
        color.Info.Printf("INFO : %s : %s\n",now.Format(formatDateTime), text )
    }
}

func Json(objs ...interface{}) {
    now := time.Now()
    for _, obj := range objs {
        t := reflect.TypeOf(obj)
        if (t.Kind() == reflect.Struct) {
            js, _ := json.MarshalIndent(obj, "", "   ")
            color.Info.Printf("INFO : %s :\n %s\n",now.Format(formatDateTime), js )
        } else {
            color.Info.Printf("INFO : %s : %s %v %s\n", now.Format(formatDateTime),decorator, obj, decorator)
        }
    }
}

func Xml(objs ...interface{}) {
    now := time.Now()
    for _, obj := range objs {
        t := reflect.TypeOf(obj)
        if (t.Kind() == reflect.Struct) {
            js, _ := xml.MarshalIndent(obj, "", "   ")
            color.Info.Printf("INFO : %s :\n %s\n",now.Format(formatDateTime), js )
        } else {
            color.Info.Printf("INFO : %s : %s %v %s\n", now.Format(formatDateTime),decorator, obj, decorator)
        }
    }
}


