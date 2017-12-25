// Generated by github.com/davyxu/cellnet/objprotogen
// DO NOT EDIT!
package coredef

import (
	"github.com/jingwanglong/cellnet"
	"reflect"
	"fmt"
)

func (self *SessionAccepted) String() string      { return fmt.Sprintf("%+v", *self) }
func (self *SessionConnected) String() string     { return fmt.Sprintf("%+v", *self) }
func (self *SessionAcceptFailed) String() string  { return fmt.Sprintf("%+v", *self) }
func (self *SessionConnectFailed) String() string { return fmt.Sprintf("%+v", *self) }
func (self *SessionClosed) String() string        { return fmt.Sprintf("%+v", *self) }
func (self *RemoteCallACK) String() string        { return fmt.Sprintf("%+v", *self) }

func init() {

	cellnet.RegisterMessageMeta("binary", "coredef.SessionAccepted", reflect.TypeOf((*SessionAccepted)(nil)).Elem(), 2087448307)
	cellnet.RegisterMessageMeta("binary", "coredef.SessionConnected", reflect.TypeOf((*SessionConnected)(nil)).Elem(), 3159799620)
	cellnet.RegisterMessageMeta("binary", "coredef.SessionAcceptFailed", reflect.TypeOf((*SessionAcceptFailed)(nil)).Elem(), 3257237101)
	cellnet.RegisterMessageMeta("binary", "coredef.SessionConnectFailed", reflect.TypeOf((*SessionConnectFailed)(nil)).Elem(), 290677256)
	cellnet.RegisterMessageMeta("binary", "coredef.SessionClosed", reflect.TypeOf((*SessionClosed)(nil)).Elem(), 1747213404)
	cellnet.RegisterMessageMeta("binary", "coredef.RemoteCallACK", reflect.TypeOf((*RemoteCallACK)(nil)).Elem(), 14297662)
}
