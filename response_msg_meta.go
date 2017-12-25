package cellnet

import (
	"bytes"
	"fmt"
	"path"
	"reflect"
)

type ResponseMessageMeta struct {
	Type  reflect.Type
	Name  string
	ID    uint32
	Codec Codec
}

var (
	resMetaByName = map[string]*MessageMeta{}
	resMetaByID   = map[uint32]*MessageMeta{}
	resMetaByType = map[reflect.Type]*MessageMeta{}
)

// 注册消息元信息(代码生成专用)
func RegisterResponseMsgMeta(codecName string, name string, msgType reflect.Type, id uint32) {

	meta := &MessageMeta{
		Type:  msgType,
		Name:  name,
		ID:    id,
		Codec: FetchCodec(codecName),
	}

	if meta.Codec == nil {
		panic("codec not register! " + codecName)
	}

	if _, ok := resMetaByName[name]; ok {
		panic("duplicate message meta register by name: " + name)
	}

	if _, ok := resMetaByID[meta.ID]; ok {
		panic(fmt.Sprintf("duplicate message meta register by id: %d", meta.ID))
	}

	if _, ok := resMetaByType[msgType]; ok {
		panic(fmt.Sprintf("duplicate message meta register by type: %d", meta.ID))
	}

	resMetaByName[name] = meta
	resMetaByID[meta.ID] = meta
	resMetaByType[msgType] = meta
}

// 根据名字查找消息元信息
func ResponseMsgMetaByName(name string) *MessageMeta {
	if v, ok := resMetaByName[name]; ok {
		return v
	}

	return nil
}

// 根据类型查找消息元信息
func ResponseMsgMetaByType(t reflect.Type) *MessageMeta {

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if v, ok := resMetaByType[t]; ok {
		return v
	}

	return nil
}

// 消息全名
func ResponseMessageFullName(rtype reflect.Type) string {

	if rtype == nil {
		panic("empty msg type")
	}

	if rtype.Kind() == reflect.Ptr {
		rtype = rtype.Elem()
	}

	var b bytes.Buffer
	b.WriteString(path.Base(rtype.PkgPath()))
	b.WriteString(".")
	b.WriteString(rtype.Name())

	return b.String()

}

// 根据id查找消息元信息
func ResponseMsgMetaByID(id uint32) *MessageMeta {
	if v, ok := resMetaByID[id]; ok {
		return v
	}

	return nil
}

// 根据id查找消息名, 没找到返回空
func ResponseMsgNameByID(id uint32) string {

	if meta := ResponseMsgMetaByID(id); meta != nil {
		return meta.Name
	}

	return ""
}

// 遍历消息元信息
func VisitResMessageMeta(callback func(*MessageMeta)) {

	for _, meta := range resMetaByName {
		callback(meta)
	}

}
