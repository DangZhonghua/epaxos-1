// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/sql/jobs.proto
// DO NOT EDIT!

package sql

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_cockroachdb_cockroach_pkg_sql_sqlbase "github.com/cockroachdb/cockroach/pkg/sql/sqlbase"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BackupJobDetails struct {
}

func (m *BackupJobDetails) Reset()                    { *m = BackupJobDetails{} }
func (m *BackupJobDetails) String() string            { return proto.CompactTextString(m) }
func (*BackupJobDetails) ProtoMessage()               {}
func (*BackupJobDetails) Descriptor() ([]byte, []int) { return fileDescriptorJobs, []int{0} }

type RestoreJobDetails struct {
}

func (m *RestoreJobDetails) Reset()                    { *m = RestoreJobDetails{} }
func (m *RestoreJobDetails) String() string            { return proto.CompactTextString(m) }
func (*RestoreJobDetails) ProtoMessage()               {}
func (*RestoreJobDetails) Descriptor() ([]byte, []int) { return fileDescriptorJobs, []int{1} }

type JobPayload struct {
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// For consistency with the SQL timestamp type, which has microsecond
	// precision, we avoid the timestamp.Timestamp WKT, which has nanosecond
	// precision, and use microsecond integers directly.
	StartedMicros     int64                                                 `protobuf:"varint,3,opt,name=started_micros,json=startedMicros,proto3" json:"started_micros,omitempty"`
	FinishedMicros    int64                                                 `protobuf:"varint,4,opt,name=finished_micros,json=finishedMicros,proto3" json:"finished_micros,omitempty"`
	ModifiedMicros    int64                                                 `protobuf:"varint,5,opt,name=modified_micros,json=modifiedMicros,proto3" json:"modified_micros,omitempty"`
	DescriptorIDs     []github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID `protobuf:"varint,6,rep,packed,name=descriptor_ids,json=descriptorIds,casttype=github.com/cockroachdb/cockroach/pkg/sql/sqlbase.ID" json:"descriptor_ids,omitempty"`
	FractionCompleted float32                                               `protobuf:"fixed32,7,opt,name=fraction_completed,json=fractionCompleted,proto3" json:"fraction_completed,omitempty"`
	Error             string                                                `protobuf:"bytes,8,opt,name=error,proto3" json:"error,omitempty"`
	// Types that are valid to be assigned to Details:
	//	*JobPayload_Backup
	//	*JobPayload_Restore
	Details isJobPayload_Details `protobuf_oneof:"details"`
}

func (m *JobPayload) Reset()                    { *m = JobPayload{} }
func (m *JobPayload) String() string            { return proto.CompactTextString(m) }
func (*JobPayload) ProtoMessage()               {}
func (*JobPayload) Descriptor() ([]byte, []int) { return fileDescriptorJobs, []int{2} }

type isJobPayload_Details interface {
	isJobPayload_Details()
	MarshalTo([]byte) (int, error)
	Size() int
}

type JobPayload_Backup struct {
	Backup *BackupJobDetails `protobuf:"bytes,10,opt,name=backup,oneof"`
}
type JobPayload_Restore struct {
	Restore *RestoreJobDetails `protobuf:"bytes,11,opt,name=restore,oneof"`
}

func (*JobPayload_Backup) isJobPayload_Details()  {}
func (*JobPayload_Restore) isJobPayload_Details() {}

func (m *JobPayload) GetDetails() isJobPayload_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *JobPayload) GetBackup() *BackupJobDetails {
	if x, ok := m.GetDetails().(*JobPayload_Backup); ok {
		return x.Backup
	}
	return nil
}

func (m *JobPayload) GetRestore() *RestoreJobDetails {
	if x, ok := m.GetDetails().(*JobPayload_Restore); ok {
		return x.Restore
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*JobPayload) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _JobPayload_OneofMarshaler, _JobPayload_OneofUnmarshaler, _JobPayload_OneofSizer, []interface{}{
		(*JobPayload_Backup)(nil),
		(*JobPayload_Restore)(nil),
	}
}

func _JobPayload_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*JobPayload)
	// details
	switch x := m.Details.(type) {
	case *JobPayload_Backup:
		_ = b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Backup); err != nil {
			return err
		}
	case *JobPayload_Restore:
		_ = b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Restore); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("JobPayload.Details has unexpected type %T", x)
	}
	return nil
}

func _JobPayload_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*JobPayload)
	switch tag {
	case 10: // details.backup
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BackupJobDetails)
		err := b.DecodeMessage(msg)
		m.Details = &JobPayload_Backup{msg}
		return true, err
	case 11: // details.restore
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RestoreJobDetails)
		err := b.DecodeMessage(msg)
		m.Details = &JobPayload_Restore{msg}
		return true, err
	default:
		return false, nil
	}
}

func _JobPayload_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*JobPayload)
	// details
	switch x := m.Details.(type) {
	case *JobPayload_Backup:
		s := proto.Size(x.Backup)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *JobPayload_Restore:
		s := proto.Size(x.Restore)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*BackupJobDetails)(nil), "cockroach.sql.BackupJobDetails")
	proto.RegisterType((*RestoreJobDetails)(nil), "cockroach.sql.RestoreJobDetails")
	proto.RegisterType((*JobPayload)(nil), "cockroach.sql.JobPayload")
}
func (m *BackupJobDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupJobDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *RestoreJobDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RestoreJobDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *JobPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JobPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJobs(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Username) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintJobs(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if m.StartedMicros != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.StartedMicros))
	}
	if m.FinishedMicros != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.FinishedMicros))
	}
	if m.ModifiedMicros != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.ModifiedMicros))
	}
	if len(m.DescriptorIDs) > 0 {
		dAtA2 := make([]byte, len(m.DescriptorIDs)*10)
		var j1 int
		for _, num := range m.DescriptorIDs {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x32
		i++
		i = encodeVarintJobs(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.FractionCompleted != 0 {
		dAtA[i] = 0x3d
		i++
		i = encodeFixed32Jobs(dAtA, i, uint32(math.Float32bits(float32(m.FractionCompleted))))
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintJobs(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	if m.Details != nil {
		nn3, err := m.Details.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn3
	}
	return i, nil
}

func (m *JobPayload_Backup) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Backup != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.Backup.Size()))
		n4, err := m.Backup.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *JobPayload_Restore) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Restore != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.Restore.Size()))
		n5, err := m.Restore.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func encodeFixed64Jobs(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Jobs(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintJobs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BackupJobDetails) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *RestoreJobDetails) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *JobPayload) Size() (n int) {
	var l int
	_ = l
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovJobs(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovJobs(uint64(l))
	}
	if m.StartedMicros != 0 {
		n += 1 + sovJobs(uint64(m.StartedMicros))
	}
	if m.FinishedMicros != 0 {
		n += 1 + sovJobs(uint64(m.FinishedMicros))
	}
	if m.ModifiedMicros != 0 {
		n += 1 + sovJobs(uint64(m.ModifiedMicros))
	}
	if len(m.DescriptorIDs) > 0 {
		l = 0
		for _, e := range m.DescriptorIDs {
			l += sovJobs(uint64(e))
		}
		n += 1 + sovJobs(uint64(l)) + l
	}
	if m.FractionCompleted != 0 {
		n += 5
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovJobs(uint64(l))
	}
	if m.Details != nil {
		n += m.Details.Size()
	}
	return n
}

func (m *JobPayload_Backup) Size() (n int) {
	var l int
	_ = l
	if m.Backup != nil {
		l = m.Backup.Size()
		n += 1 + l + sovJobs(uint64(l))
	}
	return n
}
func (m *JobPayload_Restore) Size() (n int) {
	var l int
	_ = l
	if m.Restore != nil {
		l = m.Restore.Size()
		n += 1 + l + sovJobs(uint64(l))
	}
	return n
}

func sovJobs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozJobs(x uint64) (n int) {
	return sovJobs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BackupJobDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobs
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
			return fmt.Errorf("proto: BackupJobDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupJobDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipJobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJobs
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
func (m *RestoreJobDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobs
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
			return fmt.Errorf("proto: RestoreJobDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RestoreJobDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipJobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJobs
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
func (m *JobPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobs
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
			return fmt.Errorf("proto: JobPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JobPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
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
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
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
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartedMicros", wireType)
			}
			m.StartedMicros = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartedMicros |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinishedMicros", wireType)
			}
			m.FinishedMicros = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinishedMicros |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModifiedMicros", wireType)
			}
			m.ModifiedMicros = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModifiedMicros |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType == 0 {
				var v github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowJobs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.DescriptorIDs = append(m.DescriptorIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowJobs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthJobs
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowJobs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.DescriptorIDs = append(m.DescriptorIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field DescriptorIDs", wireType)
			}
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field FractionCompleted", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.FractionCompleted = float32(math.Float32frombits(v))
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
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
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Backup", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
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
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &BackupJobDetails{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Details = &JobPayload_Backup{v}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Restore", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
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
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RestoreJobDetails{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Details = &JobPayload_Restore{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJobs
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
func skipJobs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJobs
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
					return 0, ErrIntOverflowJobs
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
					return 0, ErrIntOverflowJobs
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
				return 0, ErrInvalidLengthJobs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowJobs
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
				next, err := skipJobs(dAtA[start:])
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
	ErrInvalidLengthJobs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJobs   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/sql/jobs.proto", fileDescriptorJobs) }

var fileDescriptorJobs = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x73, 0x4d, 0x9a, 0x3f, 0x6f, 0xe4, 0xd0, 0x1e, 0x1d, 0xac, 0x0a, 0x1c, 0xab, 0x12,
	0xc2, 0x0b, 0xb6, 0x44, 0x27, 0x24, 0xa6, 0x90, 0x21, 0x89, 0x84, 0x84, 0x3c, 0xb2, 0x44, 0xe7,
	0xbb, 0x4b, 0x72, 0xc4, 0xce, 0xeb, 0xdc, 0x39, 0x03, 0x33, 0x5f, 0x80, 0x8f, 0xd5, 0x91, 0x91,
	0xa9, 0x02, 0xf3, 0x2d, 0x98, 0x50, 0xce, 0x71, 0x52, 0xda, 0xcd, 0xef, 0xf3, 0xfc, 0x1e, 0xcb,
	0x7a, 0x1e, 0xc3, 0x0b, 0x8e, 0x7c, 0xad, 0x91, 0xf1, 0x55, 0x94, 0xaf, 0x97, 0x91, 0xd9, 0xa6,
	0xd1, 0x17, 0x4c, 0x4c, 0x98, 0x6b, 0x2c, 0x90, 0x3a, 0x47, 0x37, 0x34, 0xdb, 0xf4, 0xfa, 0x6a,
	0x89, 0x4b, 0xb4, 0x4e, 0xb4, 0x7f, 0xaa, 0xa0, 0x1b, 0x0a, 0x17, 0x23, 0xc6, 0xd7, 0xbb, 0x7c,
	0x86, 0xc9, 0x58, 0x16, 0x4c, 0xa5, 0xe6, 0xe6, 0x39, 0x5c, 0xc6, 0xd2, 0x14, 0xa8, 0xe5, 0x03,
	0xf1, 0x5b, 0x0b, 0x60, 0x86, 0xc9, 0x27, 0xf6, 0x35, 0x45, 0x26, 0xa8, 0x0f, 0x7d, 0x21, 0x0d,
	0xd7, 0x2a, 0x2f, 0x14, 0x6e, 0x5c, 0xe2, 0x93, 0xa0, 0x17, 0x3f, 0x94, 0xe8, 0x35, 0x74, 0x77,
	0x46, 0xea, 0x0d, 0xcb, 0xa4, 0x7b, 0x66, 0xed, 0xe3, 0x4d, 0x5f, 0xc1, 0xc0, 0x14, 0x4c, 0x17,
	0x52, 0xcc, 0x33, 0xc5, 0x35, 0x1a, 0xb7, 0xe9, 0x93, 0xa0, 0x19, 0x3b, 0x07, 0xf5, 0xa3, 0x15,
	0xe9, 0x6b, 0x78, 0xb6, 0x50, 0x1b, 0x65, 0x56, 0x27, 0xae, 0x65, 0xb9, 0x41, 0x2d, 0x9f, 0xc0,
	0x0c, 0x85, 0x5a, 0xa8, 0x13, 0x78, 0x5e, 0x81, 0xb5, 0x7c, 0x00, 0x11, 0x06, 0xf5, 0x37, 0xa2,
	0x9e, 0x2b, 0x61, 0xdc, 0xb6, 0xdf, 0x0c, 0x9c, 0xd1, 0xa4, 0xbc, 0x1f, 0x3a, 0xe3, 0xa3, 0x33,
	0x1d, 0x9b, 0xbf, 0xf7, 0xc3, 0xdb, 0xa5, 0x2a, 0x56, 0xbb, 0x24, 0xe4, 0x98, 0x45, 0xc7, 0x2e,
	0x45, 0x12, 0x3d, 0x6d, 0xdd, 0x6c, 0xd3, 0x84, 0x19, 0x19, 0x4e, 0xc7, 0xb1, 0x73, 0x7a, 0xff,
	0x54, 0x18, 0xfa, 0x06, 0xe8, 0x42, 0x33, 0xbe, 0x6f, 0x64, 0xce, 0x31, 0xcb, 0x53, 0x59, 0x48,
	0xe1, 0x76, 0x7c, 0x12, 0x9c, 0xc5, 0x97, 0xb5, 0xf3, 0xa1, 0x36, 0xe8, 0x15, 0x9c, 0x4b, 0xad,
	0x51, 0xbb, 0x5d, 0xdb, 0x58, 0x75, 0xd0, 0x77, 0xd0, 0x4e, 0xec, 0x48, 0x2e, 0xf8, 0x24, 0xe8,
	0xbf, 0x1d, 0x86, 0xff, 0x4d, 0x1b, 0x3e, 0x5e, 0x70, 0xd2, 0x88, 0x0f, 0x01, 0xfa, 0x1e, 0x3a,
	0xba, 0xda, 0xd2, 0xed, 0xdb, 0xac, 0xff, 0x28, 0xfb, 0x64, 0xe9, 0x49, 0x23, 0xae, 0x23, 0xa3,
	0x1e, 0x74, 0x44, 0xa5, 0xce, 0x5a, 0xdd, 0xde, 0x05, 0x8c, 0x5e, 0xde, 0xfd, 0xf6, 0x1a, 0x77,
	0xa5, 0x47, 0x7e, 0x94, 0x1e, 0xf9, 0x59, 0x7a, 0xe4, 0x57, 0xe9, 0x91, 0xef, 0x7f, 0xbc, 0xc6,
	0xe7, 0xe6, 0xbe, 0x83, 0xb6, 0xfd, 0xa9, 0x6e, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x73, 0xae,
	0x45, 0x10, 0x99, 0x02, 0x00, 0x00,
}
