// Auto-generated by avdl-compiler v1.3.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: sserver-avdl/sserver.avdl

package sserver

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type WriteIndexArg struct {
	SecureIndex []byte `codec:"secureIndex" json:"secureIndex"`
	DocID       string `codec:"docID" json:"docID"`
}

type RenameIndexArg struct {
	Orig string `codec:"orig" json:"orig"`
	Curr string `codec:"curr" json:"curr"`
}

type SearchWordArg struct {
	Trapdoors [][]byte `codec:"trapdoors" json:"trapdoors"`
}

type GetSaltsArg struct {
}

type SearchServerInterface interface {
	WriteIndex(context.Context, WriteIndexArg) error
	RenameIndex(context.Context, RenameIndexArg) error
	SearchWord(context.Context, [][]byte) ([]string, error)
	GetSalts(context.Context) ([][]byte, error)
}

func SearchServerProtocol(i SearchServerInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "sserver.searchServer",
		Methods: map[string]rpc.ServeHandlerDescription{
			"writeIndex": {
				MakeArg: func() interface{} {
					ret := make([]WriteIndexArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]WriteIndexArg)
					if !ok {
						err = rpc.NewTypeError((*[]WriteIndexArg)(nil), args)
						return
					}
					err = i.WriteIndex(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"renameIndex": {
				MakeArg: func() interface{} {
					ret := make([]RenameIndexArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RenameIndexArg)
					if !ok {
						err = rpc.NewTypeError((*[]RenameIndexArg)(nil), args)
						return
					}
					err = i.RenameIndex(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"searchWord": {
				MakeArg: func() interface{} {
					ret := make([]SearchWordArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SearchWordArg)
					if !ok {
						err = rpc.NewTypeError((*[]SearchWordArg)(nil), args)
						return
					}
					ret, err = i.SearchWord(ctx, (*typedArgs)[0].Trapdoors)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getSalts": {
				MakeArg: func() interface{} {
					ret := make([]GetSaltsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetSalts(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type SearchServerClient struct {
	Cli rpc.GenericClient
}

func (c SearchServerClient) WriteIndex(ctx context.Context, __arg WriteIndexArg) (err error) {
	err = c.Cli.Call(ctx, "sserver.searchServer.writeIndex", []interface{}{__arg}, nil)
	return
}

func (c SearchServerClient) RenameIndex(ctx context.Context, __arg RenameIndexArg) (err error) {
	err = c.Cli.Call(ctx, "sserver.searchServer.renameIndex", []interface{}{__arg}, nil)
	return
}

func (c SearchServerClient) SearchWord(ctx context.Context, trapdoors [][]byte) (res []string, err error) {
	__arg := SearchWordArg{Trapdoors: trapdoors}
	err = c.Cli.Call(ctx, "sserver.searchServer.searchWord", []interface{}{__arg}, &res)
	return
}

func (c SearchServerClient) GetSalts(ctx context.Context) (res [][]byte, err error) {
	err = c.Cli.Call(ctx, "sserver.searchServer.getSalts", []interface{}{GetSaltsArg{}}, &res)
	return
}
