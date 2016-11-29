package gojsonschema

type SubSchema subSchema

func (d *Schema) RootSchema() *SubSchema {
	ss := SubSchema(*d.rootSchema)
	return &ss
}

func (s *subSchema) Exported() *SubSchema {
	ss := SubSchema(*s)
	return &ss
}

func (s *subSchema) Property() string {
	return s.property
}

func (s *SubSchema) unexported() *subSchema {
	ss := subSchema(*s)
	return &ss
}

func (s *SubSchema) AllOf() []*subSchema {
	return s.unexported().allOf
}

func (s *SubSchema) OneOf() []*subSchema {
	return s.unexported().oneOf
}

func (s *SubSchema) ItemsChildren() []*subSchema {
	return s.unexported().itemsChildren
}

func (s *SubSchema) PropertiesChildren() []*subSchema {
	return s.unexported().propertiesChildren
}

func (s *SubSchema) Format() string {
	return s.unexported().format
}

func (s *SubSchema) Property() string {
	return s.unexported().property
}
