package mapper

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/node"
	"strings"
)

type OldProvider struct {
	Name string
	Type OldType
}

type OldProviderList []OldProvider

func (l OldProviderList) FindAccessor(name string, typeName string, isGetter bool) string {
	for _, p := range l {
		accessor := p.Name
		if isGetter {
			accessor = fmt.Sprintf("%s()", accessor)
		}

		if p.Type.Name() == typeName {
			if name == "" {
				return accessor
			}
			if strings.ToLower(name) == strings.ToLower(p.Name) {
				return accessor
			}
			continue
		}

		if strings.HasPrefix(strings.ToLower(name), strings.ToLower(p.Name)) {
			name := name[len(p.Name):]
			sub := p.Type.FindAccessor(name, typeName)
			if sub != "" {
				return fmt.Sprintf("%s.%s", accessor, sub)
			}
		}

		sub := p.Type.FindAccessor(name, typeName)
		if sub != "" {
			return fmt.Sprintf("%s.%s", accessor, sub)
		}
	}
	return ""
}

type OldType interface {
	Name() string
	IsPointer() bool // TODO: Remove
	FindAccessor(name string, typeName string) string
	BuildMapper(args OldProviderList) (node.Node, error)
}

type OpaqueType struct {
	Name_ string
}

func (t *OpaqueType) Name() string {
	return t.Name_
}

func (t *OpaqueType) IsPointer() bool {
	return false
}

func (t *OpaqueType) FindAccessor(name string, typeName string) string {
	return ""
}

func (t *OpaqueType) BuildMapper(args OldProviderList) (node.Node, error) {
	accessor := args.FindAccessor("", t.Name_, false)
	if accessor != "" {
		return node.NewValue(accessor), nil
	}

	return nil, fmt.Errorf("no accessor found for type %s", t.Name_)
}

type StructType struct {
	Name_      string
	IsPointer_ bool
	Fields     OldProviderList
	Getters    OldProviderList
}

func (t *StructType) Name() string {
	return t.Name_
}

func (t *StructType) IsPointer() bool {
	return t.IsPointer_
}

func (t *StructType) FindAccessor(name string, typeName string) string {
	res := t.Fields.FindAccessor(name, typeName, false)
	if res != "" {
		return res
	}

	res = t.Getters.FindAccessor(name, typeName, true)
	if res != "" {
		return res
	}

	return ""
}

func (t *StructType) BuildMapper(args OldProviderList) (node.Node, error) {
	accessor := args.FindAccessor("", t.Name_, false)
	if accessor != "" {
		return node.NewValue(accessor), nil
	}

	fields := make([]*node.Field, len(t.Fields))
	for i, v := range t.Fields {
		accessor := args.FindAccessor(v.Name, v.Type.Name(), false)
		if accessor == "" {
			return nil, fmt.Errorf("no accessor found for field %s %s", v.Name, v.Type.Name())
		}
		fields[i] = node.NewField(v.Name, node.NewValue(accessor))
	}

	if t.IsPointer() {
		return node.NewStructPtr(t.Name(), fields...), nil
	} else {
		return node.NewStruct(t.Name(), fields...), nil
	}
}
