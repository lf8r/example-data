// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package data2

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

func easyjson80b93e8cDecodeGithubComLf8rExampleDataPkgData2(in *jlexer.Lexer, out *Base) {
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
		case "ID":
			out.ID = string(in.String())
		case "Name":
			out.Name = string(in.String())
		case "Created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "Modified":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Modified).UnmarshalJSON(data))
			}
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
func easyjson80b93e8cEncodeGithubComLf8rExampleDataPkgData2(out *jwriter.Writer, in Base) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Created\":"
		out.RawString(prefix)
		out.Raw((in.Created).MarshalJSON())
	}
	{
		const prefix string = ",\"Modified\":"
		out.RawString(prefix)
		out.Raw((in.Modified).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Base) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson80b93e8cEncodeGithubComLf8rExampleDataPkgData2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Base) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson80b93e8cEncodeGithubComLf8rExampleDataPkgData2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Base) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson80b93e8cDecodeGithubComLf8rExampleDataPkgData2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Base) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson80b93e8cDecodeGithubComLf8rExampleDataPkgData2(l, v)
}
func easyjson80b93e8cDecodeGithubComLf8rExampleDataPkgData21(in *jlexer.Lexer, out *AccessToken) {
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
		case "Val":
			out.Val = string(in.String())
		case "IssuedTo":
			out.IssuedTo = string(in.String())
		case "Role":
			out.Role = string(in.String())
		case "Issuer":
			out.Issuer = string(in.String())
		case "Expires":
			out.Expires = string(in.String())
		case "ID":
			out.ID = string(in.String())
		case "Name":
			out.Name = string(in.String())
		case "Created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "Modified":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Modified).UnmarshalJSON(data))
			}
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
func easyjson80b93e8cEncodeGithubComLf8rExampleDataPkgData21(out *jwriter.Writer, in AccessToken) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Val\":"
		out.RawString(prefix[1:])
		out.String(string(in.Val))
	}
	{
		const prefix string = ",\"IssuedTo\":"
		out.RawString(prefix)
		out.String(string(in.IssuedTo))
	}
	{
		const prefix string = ",\"Role\":"
		out.RawString(prefix)
		out.String(string(in.Role))
	}
	{
		const prefix string = ",\"Issuer\":"
		out.RawString(prefix)
		out.String(string(in.Issuer))
	}
	{
		const prefix string = ",\"Expires\":"
		out.RawString(prefix)
		out.String(string(in.Expires))
	}
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Created\":"
		out.RawString(prefix)
		out.Raw((in.Created).MarshalJSON())
	}
	{
		const prefix string = ",\"Modified\":"
		out.RawString(prefix)
		out.Raw((in.Modified).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AccessToken) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson80b93e8cEncodeGithubComLf8rExampleDataPkgData21(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AccessToken) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson80b93e8cEncodeGithubComLf8rExampleDataPkgData21(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AccessToken) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson80b93e8cDecodeGithubComLf8rExampleDataPkgData21(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AccessToken) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson80b93e8cDecodeGithubComLf8rExampleDataPkgData21(l, v)
}
