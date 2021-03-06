// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package sdk

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

func easyjson9e1087fdDecodeGithubComOvhCdsSdk(in *jlexer.Lexer, out *UserPermissionsMap) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make(UserPermissionsMap)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 int
			v1 = int(in.Int())
			(*out)[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9e1087fdEncodeGithubComOvhCdsSdk(out *jwriter.Writer, in UserPermissionsMap) {
	if in == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v2First := true
		for v2Name, v2Value := range in {
			if v2First {
				v2First = false
			} else {
				out.RawByte(',')
			}
			out.String(string(v2Name))
			out.RawByte(':')
			out.Int(int(v2Value))
		}
		out.RawByte('}')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v UserPermissionsMap) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComOvhCdsSdk(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserPermissionsMap) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComOvhCdsSdk(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserPermissionsMap) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComOvhCdsSdk(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserPermissionsMap) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComOvhCdsSdk(l, v)
}
func easyjson9e1087fdDecodeGithubComOvhCdsSdk1(in *jlexer.Lexer, out *UserPermissions) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Groups":
			if in.IsNull() {
				in.Skip()
				out.Groups = nil
			} else {
				in.Delim('[')
				if out.Groups == nil {
					if !in.IsDelim(']') {
						out.Groups = make([]string, 0, 4)
					} else {
						out.Groups = []string{}
					}
				} else {
					out.Groups = (out.Groups)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.Groups = append(out.Groups, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "GroupsAdmin":
			if in.IsNull() {
				in.Skip()
				out.GroupsAdmin = nil
			} else {
				in.Delim('[')
				if out.GroupsAdmin == nil {
					if !in.IsDelim(']') {
						out.GroupsAdmin = make([]string, 0, 4)
					} else {
						out.GroupsAdmin = []string{}
					}
				} else {
					out.GroupsAdmin = (out.GroupsAdmin)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.GroupsAdmin = append(out.GroupsAdmin, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ProjectsPerm":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.ProjectsPerm = make(map[string]int)
				} else {
					out.ProjectsPerm = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v5 int
					v5 = int(in.Int())
					(out.ProjectsPerm)[key] = v5
					in.WantComma()
				}
				in.Delim('}')
			}
		case "ApplicationsPerm":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ApplicationsPerm).UnmarshalJSON(data))
			}
		case "WorkflowsPerm":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.WorkflowsPerm).UnmarshalJSON(data))
			}
		case "PipelinesPerm":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.PipelinesPerm).UnmarshalJSON(data))
			}
		case "EnvironmentsPerm":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EnvironmentsPerm).UnmarshalJSON(data))
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
func easyjson9e1087fdEncodeGithubComOvhCdsSdk1(out *jwriter.Writer, in UserPermissions) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Groups) != 0 {
		const prefix string = ",\"Groups\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.Groups {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	if len(in.GroupsAdmin) != 0 {
		const prefix string = ",\"GroupsAdmin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.GroupsAdmin {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if len(in.ProjectsPerm) != 0 {
		const prefix string = ",\"ProjectsPerm\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v10First := true
			for v10Name, v10Value := range in.ProjectsPerm {
				if v10First {
					v10First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v10Name))
				out.RawByte(':')
				out.Int(int(v10Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.ApplicationsPerm) != 0 {
		const prefix string = ",\"ApplicationsPerm\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ApplicationsPerm).MarshalJSON())
	}
	if len(in.WorkflowsPerm) != 0 {
		const prefix string = ",\"WorkflowsPerm\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.WorkflowsPerm).MarshalJSON())
	}
	if len(in.PipelinesPerm) != 0 {
		const prefix string = ",\"PipelinesPerm\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.PipelinesPerm).MarshalJSON())
	}
	if len(in.EnvironmentsPerm) != 0 {
		const prefix string = ",\"EnvironmentsPerm\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.EnvironmentsPerm).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserPermissions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComOvhCdsSdk1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserPermissions) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComOvhCdsSdk1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserPermissions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComOvhCdsSdk1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserPermissions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComOvhCdsSdk1(l, v)
}
