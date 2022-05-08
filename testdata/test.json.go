package schema

// testdata
type Test struct {
    Arr []string `json:"arr,omitempty" yaml:"arr,omitempty"`
    Boolean bool `json:"boolean,omitempty" yaml:"boolean,omitempty"`
    ComplexArray []*TestComplexArrayItem `json:"complexArray,omitempty" yaml:"complexArray,omitempty"`
    Number float `json:"number,omitempty" yaml:"number,omitempty"`
    Obj *TestObj `json:"obj,omitempty" yaml:"obj,omitempty"`
    String string `json:"string,omitempty" yaml:"string,omitempty"`
}

type TestComplexArrayItem map[string]*TestComplexArrayItemValue

type TestComplexArrayItemValue struct {
    Prop1 string `json:"prop1,omitempty" yaml:"prop1,omitempty"`
}

type TestObj struct {
    EmptyArray1 []interface{} `json:"emptyArray1,omitempty" yaml:"emptyArray1,omitempty"`
    EmptyArray2 []interface{} `json:"emptyArray2,omitempty" yaml:"emptyArray2,omitempty"`
    EmptyObject1 struct{} `json:"emptyObject1,omitempty" yaml:"emptyObject1,omitempty"`
    EmptyObject2 struct{} `json:"emptyObject2,omitempty" yaml:"emptyObject2,omitempty"`
    JustString string `json:"justString,omitempty" yaml:"justString,omitempty"`
    MapObject map[string]*TestObjMapObjectValue `json:"mapObject,omitempty" yaml:"mapObject,omitempty"`
    MapOfEmptyObject map[string]struct{} `json:"mapOfEmptyObject,omitempty" yaml:"mapOfEmptyObject,omitempty"`
    ObjectWithField []*TestObjObjectWithFieldItem `json:"objectWithField,omitempty" yaml:"objectWithField,omitempty"`
}

type TestObjMapObjectValue struct {
    Prop1 string `json:"prop1,omitempty" yaml:"prop1,omitempty"`
    Prop2 []string `json:"prop2,omitempty" yaml:"prop2,omitempty"`
}

type TestObjObjectWithFieldItem struct {
    Prop1 string `json:"prop1,omitempty" yaml:"prop1,omitempty"`
}