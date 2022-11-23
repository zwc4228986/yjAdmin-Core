package mapping

import (
	"gopkg.in/yaml.v2"
)

func yamlUnmarshal(in []byte, out interface{}) error {
	var res interface{}
	if err := yaml.Unmarshal(in, &res); err != nil {
		return err
	}

	*out.(*interface{}) = cleanupMapValue(res)
	return nil
}

func unmarshalYamlBytes(content []byte, v interface{}, unmarshaler *Unmarshaler) error {
	var o interface{}
	if err := yamlUnmarshal(content, &o); err != nil {
		return err
	}

	return unmarshal(unmarshaler, o, v)
}

func UnmarshalYamlBytes(content []byte, v interface{}) error {
	return unmarshalYamlBytes(content, v, yamlUnmarshaler)
}
