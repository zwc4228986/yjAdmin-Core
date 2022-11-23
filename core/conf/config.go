package conf

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/mapping"
	"log"
	"os"
	"path"
	"strings"
)

var Loaders = map[string]func([]byte, interface{}) error{
	".yaml": LoadFromYamlBytes,
}

// LoadFromYamlBytes loads config into v from content yaml bytes.
func LoadFromYamlBytes(content []byte, v interface{}) error {
	return mapping.UnmarshalYamlBytes(content, v)
}

// MustLoad loads config into v from path, exits on error.
func MustLoad(path string, v interface{}) {
	if err := Load(path, v); err != nil {
		log.Fatalf("error: config file %s, %s", path, err.Error())
	}

}

func Load(file string, v interface{}) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	loader, ok := Loaders[strings.ToLower(path.Ext(file))]
	if !ok {
		return fmt.Errorf("unrecognized file type: %s", file)
	}
	return loader(content, v)
}
