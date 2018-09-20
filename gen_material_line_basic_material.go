package three
// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// material_method_generator -materialName LineBasicMaterial -materialSlug line_basic_material

import (
	"github.com/fatih/structs"
	"github.com/gopherjs/gopherjs/js"
)
	
func (m LineBasicMaterial) OnBeforeCompile() {
	m.Call("onBeforeCompile")
}

func (m LineBasicMaterial) SetValues(values MaterialParameters) {
	m.Call("setValues", structs.Map(values))
}

func (m LineBasicMaterial) ToJSON(meta interface{}) interface{} {
	return m.Call("toJSON", meta)
}

func (m LineBasicMaterial) Clone() {
	m.Call("clone")
}

func (m LineBasicMaterial) Copy(source Object3D) {
	m.Call("copy", source)
}

func (m LineBasicMaterial) Dispose() {
	m.Call("dispose")
}

func (m LineBasicMaterial) getInternalObject() *js.Object {
	return m.Object
}

