// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package event

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

func easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent(in *jlexer.Lexer, out *MessagePartVoice) {
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
		case "fileId":
			out.FileID = string(in.String())
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
func easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent(out *jwriter.Writer, in MessagePartVoice) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"fileId\":"
		out.RawString(prefix[1:])
		out.String(string(in.FileID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessagePartVoice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessagePartVoice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessagePartVoice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessagePartVoice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent(l, v)
}
func easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent1(in *jlexer.Lexer, out *MessagePartSticker) {
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
		case "fileId":
			out.FileID = string(in.String())
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
func easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent1(out *jwriter.Writer, in MessagePartSticker) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"fileId\":"
		out.RawString(prefix[1:])
		out.String(string(in.FileID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessagePartSticker) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessagePartSticker) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessagePartSticker) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessagePartSticker) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent1(l, v)
}
func easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent2(in *jlexer.Lexer, out *MessagePartReply) {
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
		case "message":
			easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent3(in, &out.Message)
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
func easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent2(out *jwriter.Writer, in MessagePartReply) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix[1:])
		easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent3(out, in.Message)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessagePartReply) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessagePartReply) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessagePartReply) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessagePartReply) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent2(l, v)
}
func easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent3(in *jlexer.Lexer, out *Message) {
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
		case "msgId":
			out.MessageID = string(in.String())
		case "from":
			(out.From).UnmarshalEasyJSON(in)
		case "text":
			out.Text = string(in.String())
		case "timestamp":
			out.Timestamp = uint64(in.Uint64())
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
func easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent3(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"msgId\":"
		out.RawString(prefix[1:])
		out.String(string(in.MessageID))
	}
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix)
		(in.From).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"timestamp\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Timestamp))
	}
	out.RawByte('}')
}
func easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent4(in *jlexer.Lexer, out *MessagePartMessage) {
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
		case "from":
			(out.From).UnmarshalEasyJSON(in)
		case "msgId":
			out.MessageID = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "timestamp":
			out.Timestamp = uint64(in.Uint64())
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
func easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent4(out *jwriter.Writer, in MessagePartMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix[1:])
		(in.From).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"msgId\":"
		out.RawString(prefix)
		out.String(string(in.MessageID))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"timestamp\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Timestamp))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessagePartMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessagePartMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessagePartMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessagePartMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent4(l, v)
}
func easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent5(in *jlexer.Lexer, out *MessagePartMention) {
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
		case "userId":
			out.UserID = string(in.String())
		case "firstName":
			out.FirstName = string(in.String())
		case "lastName":
			out.LastName = string(in.String())
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
func easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent5(out *jwriter.Writer, in MessagePartMention) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"firstName\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"lastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessagePartMention) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessagePartMention) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessagePartMention) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessagePartMention) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent5(l, v)
}
func easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent6(in *jlexer.Lexer, out *MessagePartForward) {
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
		case "message":
			easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent3(in, &out.Message)
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
func easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent6(out *jwriter.Writer, in MessagePartForward) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix[1:])
		easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent3(out, in.Message)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessagePartForward) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessagePartForward) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessagePartForward) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessagePartForward) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent6(l, v)
}
func easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent7(in *jlexer.Lexer, out *MessagePartFile) {
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
		case "fileId":
			out.FileID = string(in.String())
		case "type":
			out.Type = FileType(in.String())
		case "caption":
			out.Caption = string(in.String())
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
func easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent7(out *jwriter.Writer, in MessagePartFile) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"fileId\":"
		out.RawString(prefix[1:])
		out.String(string(in.FileID))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"caption\":"
		out.RawString(prefix)
		out.String(string(in.Caption))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessagePartFile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessagePartFile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessagePartFile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessagePartFile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent7(l, v)
}
func easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent8(in *jlexer.Lexer, out *MessagePart) {
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
		case "type":
			out.Type = MessagePartType(in.String())
		case "payload":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Payload).UnmarshalJSON(data))
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
func easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent8(out *jwriter.Writer, in MessagePart) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"payload\":"
		out.RawString(prefix)
		out.Raw((in.Payload).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessagePart) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessagePart) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3645a78fEncodeGithubComBm0IcqBotApiEvent8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessagePart) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessagePart) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3645a78fDecodeGithubComBm0IcqBotApiEvent8(l, v)
}
