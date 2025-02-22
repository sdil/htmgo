package h

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/hx"
	"strings"
)

type AttributeMap map[string]any

func (m *AttributeMap) ToMap() map[string]string {
	result := make(map[string]string)
	for k, v := range *m {
		switch v.(type) {
		case AttributeMap:
			m2 := v.(*AttributeMap).ToMap()
			for _, a := range m2 {
				result[k] = a
			}
		case string:
			result[k] = v.(string)
		default:
			result[k] = fmt.Sprintf("%v", v)
		}
	}
	return result
}

func Attribute(key string, value string) *AttributeR {
	return &AttributeR{
		Name:  key,
		Value: value,
	}
}

func AttributeList(children ...*AttributeR) *AttributeMap {
	m := make(AttributeMap)
	for _, child := range children {
		m[child.Name] = child.Value
	}
	return &m
}

func Attributes(attrs *AttributeMap) *AttributeMap {
	return attrs
}

func AttributePairs(pairs ...string) *AttributeMap {
	if len(pairs)%2 != 0 {
		return &AttributeMap{}
	}
	m := make(AttributeMap)
	for i := 0; i < len(pairs); i++ {
		m[pairs[i]] = pairs[i+1]
		i++
	}
	return &m
}

func Checked() Ren {
	return Attribute("checked", "")
}

func Id(value string) Ren {
	if strings.HasPrefix(value, "#") {
		value = value[1:]
	}
	return Attribute("id", value)
}

func Disabled() *AttributeR {
	return Attribute("disabled", "")
}

func HxTarget(target string) Ren {
	return Attribute(hx.TargetAttr, target)
}

func Name(name string) Ren {
	return Attribute("name", name)
}

func HxConfirm(message string) Ren {
	return Attribute(hx.ConfirmAttr, message)
}

// HxInclude https://htmx.org/attributes/hx-include/
func HxInclude(selector string) Ren {
	return Attribute(hx.IncludeAttr, selector)
}

func HxIndicator(tag string) *AttributeR {
	return Attribute(hx.IndicatorAttr, tag)
}

func TriggerChildren() Ren {
	return HxExtension("trigger-children")
}

func HxTriggerString(triggers ...string) *AttributeR {
	trigger := hx.NewStringTrigger(strings.Join(triggers, ", "))
	return Attribute(hx.TriggerAttr, trigger.ToString())
}

func HxTrigger(opts ...hx.TriggerEvent) *AttributeR {
	return Attribute(hx.TriggerAttr, hx.NewTrigger(opts...).ToString())
}

func HxTriggerClick(opts ...hx.Modifier) *AttributeR {
	return HxTrigger(hx.OnClick(opts...))
}

func HxExtension(value string) Ren {
	return Attribute(hx.ExtAttr, value)
}

func Href(path string) Ren {
	return Attribute("href", path)
}

func Target(target string) Ren {
	return Attribute("target", target)
}

func D(value string) Ren {
	return Attribute("d", value)
}

func Alt(value string) Ren {
	return Attribute("alt", value)
}

func For(value string) Ren {
	return Attribute("for", value)
}

func Type(name string) Ren {
	return Attribute("type", name)
}

func Placeholder(placeholder string) Ren {
	return Attribute("placeholder", placeholder)
}

func Hidden() Ren {
	return Attribute("style", "display:none")
}

func Class(value ...string) *AttributeR {
	return Attribute("class", MergeClasses(value...))
}

func ClassX(value string, m ClassMap) Ren {
	builder := strings.Builder{}
	builder.WriteString(value)
	builder.WriteString(" ")
	for k, v := range m {
		if v {
			builder.WriteString(k)
			builder.WriteString(" ")
		}
	}
	return Class(builder.String())
}

func MergeClasses(classes ...string) string {
	if len(classes) == 1 {
		return classes[0]
	}
	builder := strings.Builder{}
	for _, s := range classes {
		builder.WriteString(s)
		builder.WriteString(" ")
	}
	return builder.String()
}

func Boost() Ren {
	return Attribute(hx.BoostAttr, "true")
}

func IfQueryParam(key string, node *Element) Ren {
	return Fragment(Attribute("hx-if-qp:"+key, "true"), node)
}

func ReadOnly() *AttributeR {
	return Attribute("readonly", "")
}

func Required() *AttributeR {
	return Attribute("required", "")
}

func Multiple() *AttributeR {
	return Attribute("multiple", "")
}

func Selected() *AttributeR {
	return Attribute("selected", "")
}

func MaxLength(value int) *AttributeR {
	return Attribute("maxlength", fmt.Sprintf("%d", value))
}

func MinLength(value int) *AttributeR {
	return Attribute("minlength", fmt.Sprintf("%d", value))
}

func Size(value int) *AttributeR {
	return Attribute("size", fmt.Sprintf("%d", value))
}

func Width(value int) *AttributeR {
	return Attribute("width", fmt.Sprintf("%d", value))
}

func Height(value int) *AttributeR {
	return Attribute("height", fmt.Sprintf("%d", value))
}

func Download(value bool) *AttributeR {
	return Attribute("download", fmt.Sprintf("%t", value))
}

func Rel(value string) *AttributeR {
	return Attribute("rel", value)
}

func Pattern(value string) *AttributeR {
	return Attribute("pattern", value)
}

func Action(value string) *AttributeR {
	return Attribute("action", value)
}

func Method(value string) *AttributeR {
	return Attribute("method", value)
}

func Enctype(value string) *AttributeR {
	return Attribute("enctype", value)
}

func AutoComplete(value string) *AttributeR {
	return Attribute("autocomplete", value)
}

func AutoFocus() *AttributeR {
	return Attribute("autofocus", "")
}

func NoValidate() *AttributeR {
	return Attribute("novalidate", "")
}

func Step(value string) *AttributeR {
	return Attribute("step", value)
}

func Max(value string) *AttributeR {
	return Attribute("max", value)
}

func Min(value string) *AttributeR {
	return Attribute("min", value)
}

func Cols(value int) *AttributeR {
	return Attribute("cols", fmt.Sprintf("%d", value))
}

func Rows(value int) *AttributeR {
	return Attribute("rows", fmt.Sprintf("%d", value))
}

func Wrap(value string) *AttributeR {
	return Attribute("wrap", value)
}

func Role(value string) *AttributeR {
	return Attribute("role", value)
}

func AriaLabel(value string) *AttributeR {
	return Attribute("aria-label", value)
}

func AriaHidden(value bool) *AttributeR {
	return Attribute("aria-hidden", fmt.Sprintf("%t", value))
}

func TabIndex(value int) *AttributeR {
	return Attribute("tabindex", fmt.Sprintf("%d", value))
}
