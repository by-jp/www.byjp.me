package shared

import (
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func LoadConfig(prefix string, config interface{}) error {
	if reflect.TypeOf(config).Kind() != reflect.Ptr {
		return fmt.Errorf("config must be a struct pointer")
	}
	configElem := reflect.ValueOf(config).Elem()
	if configElem.Kind() != reflect.Struct {
		return fmt.Errorf("config must be a struct pointer")
	}

	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := envconfig.Process(prefix, config); err != nil {
		return err
	}

	// rootField, hasRoot := configElem.Type().FieldByName("Root")
	// if !hasRoot {
	// 	return nil
	// }

	// if rootField.Type.Kind() != reflect.String {
	// 	return fmt.Errorf("config field 'Root' must be a string, not %s", rootField.Type)
	// }

	// root := rootField..String()
	// fmt.Printf("%s, %T\n", root, root)

	// if !isDir(path.Join(root, "content")) {
	// 	return fmt.Errorf(
	// 		"root must be the root of your hugo blog (%s/content/ is not a directory)",
	// 		root,
	// 	)
	// }

	return nil
}

func isFile(pathStr string) bool {
	st, err := os.Stat(pathStr)
	if os.IsNotExist(err) {
		return false
	}

	return st.Mode().IsRegular()
}

func isDir(pathStr string) bool {
	st, err := os.Stat(pathStr)
	if os.IsNotExist(err) {
		return false
	}

	return st.IsDir()
}
