// Generated by github.com/davyxu/gosproto/sprotogen
// DO NOT EDIT!

package gamedef

import (
	"reflect"

	"github.com/jingwanglong/cellnet/codec/sproto"
	"fmt"
)

type TestEchoACK struct {
	Content string `sproto:"string,0,name=Content"`
}

func (self *TestEchoACK) String() string { return fmt.Sprintf("%+v", *self) }

var SProtoStructs = []reflect.Type{

	reflect.TypeOf((*TestEchoACK)(nil)).Elem(), // 1899977859
}

var SProtoEnumValue = map[string]map[int32]string{}

func init() {
	sprotocodec.AutoRegisterMessageMeta(SProtoStructs)
}
