// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package python

import (
	"bytes"
	"context"
	"sync"
	"fmt"
	thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = sync.Mutex{}
var _ = bytes.Equal
var _ = context.Background

var GoUnusedProtection__ int;

type Hidden struct {
}

func NewHidden() *Hidden {
  return &Hidden{}
}

type HiddenBuilder struct {
  obj *Hidden
}

func NewHiddenBuilder() *HiddenBuilder{
  return &HiddenBuilder{
    obj: NewHidden(),
  }
}

func (p HiddenBuilder) Emit() *Hidden{
  return &Hidden{
  }
}

func (p *Hidden) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Hidden) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("Hidden"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Hidden) String() string {
  if p == nil {
    return "<nil>"
  }

  return fmt.Sprintf("Hidden({})")
}

type Flags struct {
}

func NewFlags() *Flags {
  return &Flags{}
}

type FlagsBuilder struct {
  obj *Flags
}

func NewFlagsBuilder() *FlagsBuilder{
  return &FlagsBuilder{
    obj: NewFlags(),
  }
}

func (p FlagsBuilder) Emit() *Flags{
  return &Flags{
  }
}

func (p *Flags) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Flags) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("Flags"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Flags) String() string {
  if p == nil {
    return "<nil>"
  }

  return fmt.Sprintf("Flags({})")
}

// Attributes:
//  - Name
type Name struct {
  Name string `thrift:"name,1" db:"name" json:"name"`
}

func NewName() *Name {
  return &Name{}
}


func (p *Name) GetName() string {
  return p.Name
}
type NameBuilder struct {
  obj *Name
}

func NewNameBuilder() *NameBuilder{
  return &NameBuilder{
    obj: NewName(),
  }
}

func (p NameBuilder) Emit() *Name{
  return &Name{
    Name: p.obj.Name,
  }
}

func (n *NameBuilder) Name(name string) *NameBuilder {
  n.obj.Name = name
  return n
}

func (n *Name) SetName(name string) *Name {
  n.Name = name
  return n
}

func (p *Name) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Name)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 1: ", err)
  } else {
    p.Name = v
  }
  return nil
}

func (p *Name) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("Name"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Name) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err) }
  if err := oprot.WriteString(string(p.Name)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err) }
  return err
}

func (p *Name) String() string {
  if p == nil {
    return "<nil>"
  }

  nameVal := fmt.Sprintf("%v", p.Name)
  return fmt.Sprintf("Name({Name:%s})", nameVal)
}

type IOBuf struct {
}

func NewIOBuf() *IOBuf {
  return &IOBuf{}
}

type IOBufBuilder struct {
  obj *IOBuf
}

func NewIOBufBuilder() *IOBufBuilder{
  return &IOBufBuilder{
    obj: NewIOBuf(),
  }
}

func (p IOBufBuilder) Emit() *IOBuf{
  return &IOBuf{
  }
}

func (p *IOBuf) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *IOBuf) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("IOBuf"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *IOBuf) String() string {
  if p == nil {
    return "<nil>"
  }

  return fmt.Sprintf("IOBuf({})")
}

// Attributes:
//  - Name
//  - TypeHint
type Adapter struct {
  Name string `thrift:"name,1" db:"name" json:"name"`
  TypeHint string `thrift:"typeHint,2" db:"typeHint" json:"typeHint"`
}

func NewAdapter() *Adapter {
  return &Adapter{}
}


func (p *Adapter) GetName() string {
  return p.Name
}

func (p *Adapter) GetTypeHint() string {
  return p.TypeHint
}
type AdapterBuilder struct {
  obj *Adapter
}

func NewAdapterBuilder() *AdapterBuilder{
  return &AdapterBuilder{
    obj: NewAdapter(),
  }
}

func (p AdapterBuilder) Emit() *Adapter{
  return &Adapter{
    Name: p.obj.Name,
    TypeHint: p.obj.TypeHint,
  }
}

func (a *AdapterBuilder) Name(name string) *AdapterBuilder {
  a.obj.Name = name
  return a
}

func (a *AdapterBuilder) TypeHint(typeHint string) *AdapterBuilder {
  a.obj.TypeHint = typeHint
  return a
}

func (a *Adapter) SetName(name string) *Adapter {
  a.Name = name
  return a
}

func (a *Adapter) SetTypeHint(typeHint string) *Adapter {
  a.TypeHint = typeHint
  return a
}

func (p *Adapter) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Adapter)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 1: ", err)
  } else {
    p.Name = v
  }
  return nil
}

func (p *Adapter)  ReadField2(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 2: ", err)
  } else {
    p.TypeHint = v
  }
  return nil
}

func (p *Adapter) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("Adapter"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Adapter) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err) }
  if err := oprot.WriteString(string(p.Name)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err) }
  return err
}

func (p *Adapter) writeField2(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("typeHint", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:typeHint: ", p), err) }
  if err := oprot.WriteString(string(p.TypeHint)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.typeHint (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:typeHint: ", p), err) }
  return err
}

func (p *Adapter) String() string {
  if p == nil {
    return "<nil>"
  }

  nameVal := fmt.Sprintf("%v", p.Name)
  typeHintVal := fmt.Sprintf("%v", p.TypeHint)
  return fmt.Sprintf("Adapter({Name:%s TypeHint:%s})", nameVal, typeHintVal)
}

