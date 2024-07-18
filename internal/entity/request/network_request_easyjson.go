// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package request

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson74658edaDecodeGithubComKarimovKamilOtusGoFinalProjectInternalEntityRequest(in *jlexer.Lexer, out *NetworkRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "network":
			out.Network = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson74658edaEncodeGithubComKarimovKamilOtusGoFinalProjectInternalEntityRequest(out *jwriter.Writer, in NetworkRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"network\":"
		out.RawString(prefix[1:])
		out.String(string(in.Network))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NetworkRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74658edaEncodeGithubComKarimovKamilOtusGoFinalProjectInternalEntityRequest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NetworkRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74658edaEncodeGithubComKarimovKamilOtusGoFinalProjectInternalEntityRequest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NetworkRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74658edaDecodeGithubComKarimovKamilOtusGoFinalProjectInternalEntityRequest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NetworkRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74658edaDecodeGithubComKarimovKamilOtusGoFinalProjectInternalEntityRequest(l, v)
}
