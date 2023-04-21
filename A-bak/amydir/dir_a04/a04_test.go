package dira04

import "testing"

// slice性能

type ObjStruct struct {
	id int
}

func createObjsArray(n int) []*ObjStruct {
	objs := make([]ObjStruct, n) // 0.75
	refs := make([]*ObjStruct, n)
	for i := 0; i < n; i++ {
		refs[i] = &objs[i]
	}
	return refs
}

func createObjs(n int) []*ObjStruct {
	refs := make([]*ObjStruct, n) // 2.46s
	// refs := make([]*ObjStruct, 0)   // 7s
	for i := 0; i < n; i++ {
		obj := ObjStruct{}
		refs = append(refs, &obj)
	}
	return refs
}

func TestXxx(t *testing.T) {
	// createObjs(90000000)
	createObjsArray(90000000)
}
