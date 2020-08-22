package db

import (
	"log"
	"os"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// MapUpdateData to bson values for mongo update
func MapUpdateData(data interface{}, rewrite bool) bson.D {
	values := reflect.ValueOf(data)
	fields := reflect.TypeOf(data)

	var setData bson.D
	for i := 0; i < values.NumField(); i++ {
		key := fields.Field(i).Name
		val := values.Field(i).Interface()
		if (val != "" || rewrite) && key != "ID" {
			setData = append(setData, bson.E{Key: strings.ToLower(key), Value: val})
		}
	}

	return setData
}

// GetEnvVar or throw error if missing
func GetEnvVar(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("Missing %v environment variable", key)
	}
	return value
}
