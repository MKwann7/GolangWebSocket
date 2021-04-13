package helper

import (
	"github.com/MKwann7/GolangWebSocket/cmd/app/dtos"
	"reflect"
)

func NewEntityFromReflection(entityType reflect.Type) interface{} {
	switch entityType.String() {
	case "dtos.User":
		return dtos.User{}
	case "dtos.VistiorBrowser":
		return dtos.VisitorBrowser{}
	}
}
