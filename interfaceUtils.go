package utils

import (
	"fmt"
	// "github.com/Beryl8Plus/be8-sfdc-integration/src/appError"
	"strconv"
	"strings"
)

func ToBool(val interface{}) bool {
	switch value := val.(type) {
	case string:
		if strings.ToLower(value) == "true" {
			return true
		} else {
			return false
		}
	case bool:
		return value
	case int:
		if value > 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func ToInt(val interface{}) (result int) {
	switch value := val.(type) {
	case string:
		var err error
		result, err = strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
	case int:
		result = value
	case bool:
		if value {
			result = 1
		} else {
			result = 0
		}
	default:
		errMsg := fmt.Sprintf("Cannot convert [%v] to int", val)
		panic(errMsg)
	}
	return
}

func ToString(val interface{}) (str string) {
	switch value := val.(type) {
	case nil:
		str = ""
	case string:
		str = value
	case int:
		str = strconv.Itoa(value)
	case bool:
		if value {
			str = "true"
		} else {
			str = "false"
		}
	}

	return
}

func ToInterfaceMapByString(thisInterface interface{}) map[string]interface{} {
	var mapInterfaceByString map[string]interface{}

	if interfaceMap, ok := thisInterface.(map[interface{}]interface{}); ok {
		mapInterfaceByString = make(map[string]interface{})
		for key, val := range interfaceMap {
			mapInterfaceByString[key.(string)] = val
		}
	}

	if interfaceMap, ok := thisInterface.(map[string]interface{}); ok {
		mapInterfaceByString = make(map[string]interface{})
		for key, val := range interfaceMap {
			mapInterfaceByString[key] = val
		}
	}

	if interfaceMap, ok := thisInterface.(map[string]string); ok {
		mapInterfaceByString = make(map[string]interface{})
		for key, val := range interfaceMap {
			mapInterfaceByString[key] = val
		}
	}
	return mapInterfaceByString
}

func ToArrayOfString(thisInterface interface{}) []string {
	var strArray []string
	if thisInterfaces, ok := thisInterface.([]interface{}); ok {
		for _, eachInterface := range thisInterfaces {
			strArray = append(strArray, ToString(eachInterface))
		}
	}

	if stringArray, ok := thisInterface.([]string); ok {
		return stringArray
	}

	return strArray

}

func ToArrayOfInterfaceMapByString(thisInterface interface{}) []map[string]interface{} {
	var mapInterfaceByStrings []map[string]interface{}

	if thisInterfaces, ok := thisInterface.([]interface{}); ok {
		for _, eachInterface := range thisInterfaces {
			mapInterfaceByStrings = append(mapInterfaceByStrings, ToInterfaceMapByString(eachInterface))
		}
	}

	return mapInterfaceByStrings
}

func ToArrayOfStringMapByString(thisInterface interface{}) []map[string]string {
	var mapStringByStrings []map[string]string

	if thisInterfaces, ok := thisInterface.([]interface{}); ok {
		for _, eachInterface := range thisInterfaces {
			mapStringByStrings = append(mapStringByStrings, ToStringMapByString(eachInterface))
		}
	}
	return mapStringByStrings
}

func ToArrayOfInterface(thisInterface interface{}) []interface{} {
	if thisInterfaces, ok := thisInterface.([]interface{}); ok {
		return thisInterfaces
	}
	return nil
}

func ToStringMapByString(eachInterface interface{}) map[string]string {
	mapStringByString := make(map[string]string)

	if interfaceMap, ok := eachInterface.(map[interface{}]interface{}); ok {
		for key, val := range interfaceMap {
			mapStringByString[key.(string)] = val.(string)
		}
	}

	if interfaceMap, ok := eachInterface.(map[string]interface{}); ok {
		for key, val := range interfaceMap {
			mapStringByString[key] = val.(string)
		}
	}

	return mapStringByString
}
