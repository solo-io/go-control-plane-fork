// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/filter/http/squash.proto

package envoy_api_v2_filter_http

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import google_protobuf4 "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// [#proto-status: experimental]
type Squash struct {
	// The name of the cluster that hosts the Squash server.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// When the filter requests the Squash server to create a DebugAttachment, it will use this
	// structure as template for the body of the request. It can contain reference to environment
	// variables in the form of '{{ ENV_VAR_NAME }}'. These can be used to provide the Squash server
	// with more information to find the process to attach the debugger to. For example, in a
	// Istio/k8s environment, this will contain information on the pod:
	//
	// .. code-block:: json
	//
	//  {
	//    "spec": {
	//      "attachment": {
	//        "pod": "{{ POD_NAME }}",
	//        "namespace": "{{ POD_NAMESPACE }}"
	//      },
	//      "match_request": true
	//    }
	//  }
	//
	// (where POD_NAME, POD_NAMESPACE are configured in the pod via the Downward API)
	AttachmentTemplate *google_protobuf4.Struct `protobuf:"bytes,2,opt,name=attachment_template,json=attachmentTemplate" json:"attachment_template,omitempty"`
	// The timeout for individual requests sent to the Squash cluster. Defaults to 1 second.
	RequestTimeout *google_protobuf.Duration `protobuf:"bytes,3,opt,name=request_timeout,json=requestTimeout" json:"request_timeout,omitempty"`
	// The total timeout Squash will delay a request and wait for it to be attached. Defaults to 60
	// seconds.
	AttachmentTimeout *google_protobuf.Duration `protobuf:"bytes,4,opt,name=attachment_timeout,json=attachmentTimeout" json:"attachment_timeout,omitempty"`
	// Amount of time to poll for the status of the attachment object in the Squash server
	// (to check if has been attached). Defaults to 1 second.
	AttachmentPollPeriod *google_protobuf.Duration `protobuf:"bytes,5,opt,name=attachment_poll_period,json=attachmentPollPeriod" json:"attachment_poll_period,omitempty"`
}

func (m *Squash) Reset()                    { *m = Squash{} }
func (m *Squash) String() string            { return proto.CompactTextString(m) }
func (*Squash) ProtoMessage()               {}
func (*Squash) Descriptor() ([]byte, []int) { return fileDescriptorSquash, []int{0} }

func (m *Squash) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Squash) GetAttachmentTemplate() *google_protobuf4.Struct {
	if m != nil {
		return m.AttachmentTemplate
	}
	return nil
}

func (m *Squash) GetRequestTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.AttachmentTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentPollPeriod() *google_protobuf.Duration {
	if m != nil {
		return m.AttachmentPollPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*Squash)(nil), "envoy.api.v2.filter.http.Squash")
}
func (m *Squash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Squash) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Cluster) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSquash(dAtA, i, uint64(len(m.Cluster)))
		i += copy(dAtA[i:], m.Cluster)
	}
	if m.AttachmentTemplate != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSquash(dAtA, i, uint64(m.AttachmentTemplate.Size()))
		n1, err := m.AttachmentTemplate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.RequestTimeout != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSquash(dAtA, i, uint64(m.RequestTimeout.Size()))
		n2, err := m.RequestTimeout.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.AttachmentTimeout != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSquash(dAtA, i, uint64(m.AttachmentTimeout.Size()))
		n3, err := m.AttachmentTimeout.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.AttachmentPollPeriod != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSquash(dAtA, i, uint64(m.AttachmentPollPeriod.Size()))
		n4, err := m.AttachmentPollPeriod.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeVarintSquash(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Squash) Size() (n int) {
	var l int
	_ = l
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentTemplate != nil {
		l = m.AttachmentTemplate.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.RequestTimeout != nil {
		l = m.RequestTimeout.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentTimeout != nil {
		l = m.AttachmentTimeout.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentPollPeriod != nil {
		l = m.AttachmentPollPeriod.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	return n
}

func sovSquash(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSquash(x uint64) (n int) {
	return sovSquash(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Squash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSquash
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Squash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Squash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentTemplate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentTemplate == nil {
				m.AttachmentTemplate = &google_protobuf4.Struct{}
			}
			if err := m.AttachmentTemplate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RequestTimeout == nil {
				m.RequestTimeout = &google_protobuf.Duration{}
			}
			if err := m.RequestTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentTimeout == nil {
				m.AttachmentTimeout = &google_protobuf.Duration{}
			}
			if err := m.AttachmentTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentPollPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentPollPeriod == nil {
				m.AttachmentPollPeriod = &google_protobuf.Duration{}
			}
			if err := m.AttachmentPollPeriod.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSquash(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSquash
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSquash(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSquash
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSquash
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSquash
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSquash(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSquash = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSquash   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/filter/http/squash.proto", fileDescriptorSquash) }

var fileDescriptorSquash = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xb1, 0x4e, 0x42, 0x31,
	0x14, 0x86, 0x53, 0x40, 0x0c, 0xd5, 0x68, 0xac, 0x46, 0xae, 0x84, 0xdc, 0x10, 0x5d, 0x98, 0xda,
	0x04, 0xdf, 0x80, 0x38, 0xb0, 0x49, 0x80, 0x9d, 0x14, 0x38, 0x40, 0x93, 0x42, 0x4b, 0xef, 0x29,
	0x89, 0xaf, 0xe6, 0xe4, 0xe8, 0xe8, 0xe4, 0x6c, 0xd8, 0x7c, 0x0b, 0x63, 0x6f, 0x09, 0x37, 0x3a,
	0xb0, 0x35, 0xf9, 0xcf, 0xf7, 0xe5, 0x2b, 0x6d, 0x4a, 0xab, 0xc4, 0x5c, 0x69, 0x04, 0x27, 0x96,
	0x88, 0x56, 0x64, 0x1b, 0x2f, 0xb3, 0x25, 0xb7, 0xce, 0xa0, 0x61, 0x09, 0xac, 0xb7, 0xe6, 0x85,
	0x4b, 0xab, 0xf8, 0xb6, 0xc3, 0xf3, 0x33, 0xfe, 0x7b, 0xd6, 0x48, 0x17, 0xc6, 0x2c, 0x34, 0x88,
	0x70, 0x37, 0xf1, 0x73, 0x31, 0xf3, 0x4e, 0xa2, 0x32, 0xeb, 0x9c, 0x6c, 0x34, 0xff, 0xee, 0x19,
	0x3a, 0x3f, 0xc5, 0xb8, 0xd6, 0xb7, 0x52, 0xab, 0x99, 0x44, 0x10, 0xfb, 0x47, 0x3e, 0xdc, 0x7f,
	0x96, 0x68, 0x75, 0x18, 0x0a, 0xd8, 0x03, 0x3d, 0x9d, 0x6a, 0x9f, 0x21, 0xb8, 0x84, 0xb4, 0x48,
	0xbb, 0xd6, 0xad, 0xbd, 0x7e, 0xbf, 0x95, 0x2b, 0xae, 0xd4, 0x22, 0x83, 0xfd, 0xc2, 0x7a, 0xf4,
	0x5a, 0x22, 0xca, 0xe9, 0x72, 0x05, 0x6b, 0x1c, 0x23, 0xac, 0xac, 0x96, 0x08, 0x49, 0xa9, 0x45,
	0xda, 0x67, 0x9d, 0x3a, 0xcf, 0x23, 0xf8, 0x3e, 0x82, 0x0f, 0x43, 0xc4, 0x80, 0x1d, 0x98, 0x51,
	0x44, 0x58, 0x97, 0x5e, 0x3a, 0xd8, 0x78, 0xc8, 0x70, 0x8c, 0x6a, 0x05, 0xc6, 0x63, 0x52, 0x0e,
	0x96, 0xbb, 0x7f, 0x96, 0xa7, 0xf8, 0xd5, 0xc1, 0x45, 0x24, 0x46, 0x39, 0xc0, 0x7a, 0x94, 0x15,
	0x6b, 0xa2, 0xa6, 0x72, 0x4c, 0x73, 0x55, 0xc8, 0x89, 0xa6, 0x67, 0x7a, 0x5b, 0x30, 0x59, 0xa3,
	0xf5, 0xd8, 0x82, 0x53, 0x66, 0x96, 0x9c, 0x1c, 0xb3, 0xdd, 0x1c, 0xc0, 0xbe, 0xd1, 0xba, 0x1f,
	0xb0, 0xee, 0xf9, 0xfb, 0x2e, 0x25, 0x1f, 0xbb, 0x94, 0x7c, 0xed, 0x52, 0x32, 0xa9, 0x06, 0xec,
	0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x45, 0xe0, 0xa1, 0x6a, 0xfe, 0x01, 0x00, 0x00,
}
