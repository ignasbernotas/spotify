type Header struct {
   ContentType      *string      `protobuf:"bytes,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
   Method           *string      `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
   StatusCode       *int32       `protobuf:"zigzag32,4,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
   Uri              *string      `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
   UserFields       []*UserField `protobuf:"bytes,6,rep,name=user_fields,json=userFields" json:"user_fields,omitempty"`
   XXX_unrecognized []byte       `json:"-"`
}
