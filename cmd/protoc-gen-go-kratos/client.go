package main

import (
	"fmt"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

const (
	contextPackage       = protogen.GoImportPath("context")
	transportGRPCPackage = protogen.GoImportPath("github.com/go-kratos/kratos/v2/transport/grpc")
	transportHTTPPackage = protogen.GoImportPath("github.com/go-kratos/kratos/v2/transport/http")
	gRPCPackage          = protogen.GoImportPath("google.golang.org/grpc")
	wrrPackage           = protogen.GoImportPath("github.com/go-kratos/kratos/v2/selector/wrr")
)

// generateFile generates a _kratos.pb.go file.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_kratos.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-kratos. DO NOT EDIT")
	g.P("// versions:")
	g.P(fmt.Sprintf("// protoc-gen-go-kratos %s", release))
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(gen, file, g)
	return g
}

// generateFileContent generates file content.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the kratos package it is being compiled against.")
	g.P("var _ = new(", contextPackage.Ident("Context"), ")")
	g.QualifiedGoIdent(transportGRPCPackage.Ident(""))
	g.QualifiedGoIdent(wrrPackage.Ident(""))
	g.QualifiedGoIdent(transportHTTPPackage.Ident(""))
	g.P()
	clientTemplateS := NewClientTemplate()
	for _, service := range file.Services {
		host := proto.GetExtension(service.Desc.Options(), annotations.E_DefaultHost).(string)
		name := service.GoName
		clientTemplateS.AppendClientInfo(name, host)
	}
	g.P(clientTemplateS.execute())
}
