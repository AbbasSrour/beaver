package gopher

type Descriptor struct {
	Tag  string // struct tag.
	Size int    // varchar size.
	Name string // field name.
	//Info          *TypeInfo               // field type info.
	ValueScanner  any                     // custom field codec.
	Unique        bool                    // unique index of field.
	Nillable      bool                    // nillable struct field.
	Optional      bool                    // nullable field in database.
	Immutable     bool                    // create only field.
	Default       any                     // default value on create.
	UpdateDefault any                     // default value on update.
	Validators    []any                   // validator functions.
	StorageKey    string                  // sql column or gremlin property.
	Enums         []struct{ N, V string } // enum values.
	Sensitive     bool                    // sensitive info string field.
	SchemaType    map[string]string       // override the schema type.
	//Annotations   []schema.Annotation     // field annotations.
	Comment string // field comment.
	Err     error
}

type FieldInterface interface{}

type Field struct {
	descriptor *Descriptor
}

func Column(tag string) *Field {
	d := &Field{}
	d.descriptor.Tag = tag
	d.descriptor.Name = tag
	return d
}

type IntFieldInterface interface {
	FieldInterface
}

func (f *Field) Int() *Field {
	return f
}

func (f *Field) Unique() *Field {
	f.descriptor.Unique = true
	return f
}
