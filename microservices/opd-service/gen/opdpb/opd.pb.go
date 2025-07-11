// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.2
// source: opd.proto

package opdpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckAppointmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PatientId     string                 `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	AppointmentId string                 `protobuf:"bytes,2,opt,name=appointment_id,json=appointmentId,proto3" json:"appointment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckAppointmentRequest) Reset() {
	*x = CheckAppointmentRequest{}
	mi := &file_opd_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckAppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAppointmentRequest) ProtoMessage() {}

func (x *CheckAppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opd_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAppointmentRequest.ProtoReflect.Descriptor instead.
func (*CheckAppointmentRequest) Descriptor() ([]byte, []int) {
	return file_opd_proto_rawDescGZIP(), []int{0}
}

func (x *CheckAppointmentRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *CheckAppointmentRequest) GetAppointmentId() string {
	if x != nil {
		return x.AppointmentId
	}
	return ""
}

type CheckAppointmentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Valid         bool                   `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckAppointmentResponse) Reset() {
	*x = CheckAppointmentResponse{}
	mi := &file_opd_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckAppointmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAppointmentResponse) ProtoMessage() {}

func (x *CheckAppointmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opd_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAppointmentResponse.ProtoReflect.Descriptor instead.
func (*CheckAppointmentResponse) Descriptor() ([]byte, []int) {
	return file_opd_proto_rawDescGZIP(), []int{1}
}

func (x *CheckAppointmentResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *CheckAppointmentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StartConsultationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PatientId     string                 `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	DoctorId      string                 `protobuf:"bytes,2,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`
	AppointmentId string                 `protobuf:"bytes,3,opt,name=appointment_id,json=appointmentId,proto3" json:"appointment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartConsultationRequest) Reset() {
	*x = StartConsultationRequest{}
	mi := &file_opd_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartConsultationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartConsultationRequest) ProtoMessage() {}

func (x *StartConsultationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opd_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartConsultationRequest.ProtoReflect.Descriptor instead.
func (*StartConsultationRequest) Descriptor() ([]byte, []int) {
	return file_opd_proto_rawDescGZIP(), []int{2}
}

func (x *StartConsultationRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *StartConsultationRequest) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

func (x *StartConsultationRequest) GetAppointmentId() string {
	if x != nil {
		return x.AppointmentId
	}
	return ""
}

type StartConsultationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VisitId       string                 `protobuf:"bytes,1,opt,name=visit_id,json=visitId,proto3" json:"visit_id,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartConsultationResponse) Reset() {
	*x = StartConsultationResponse{}
	mi := &file_opd_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartConsultationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartConsultationResponse) ProtoMessage() {}

func (x *StartConsultationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opd_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartConsultationResponse.ProtoReflect.Descriptor instead.
func (*StartConsultationResponse) Descriptor() ([]byte, []int) {
	return file_opd_proto_rawDescGZIP(), []int{3}
}

func (x *StartConsultationResponse) GetVisitId() string {
	if x != nil {
		return x.VisitId
	}
	return ""
}

func (x *StartConsultationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RecordDiagnosisRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VisitId       string                 `protobuf:"bytes,1,opt,name=visit_id,json=visitId,proto3" json:"visit_id,omitempty"`
	Diagnosis     string                 `protobuf:"bytes,2,opt,name=diagnosis,proto3" json:"diagnosis,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecordDiagnosisRequest) Reset() {
	*x = RecordDiagnosisRequest{}
	mi := &file_opd_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecordDiagnosisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordDiagnosisRequest) ProtoMessage() {}

func (x *RecordDiagnosisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opd_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordDiagnosisRequest.ProtoReflect.Descriptor instead.
func (*RecordDiagnosisRequest) Descriptor() ([]byte, []int) {
	return file_opd_proto_rawDescGZIP(), []int{4}
}

func (x *RecordDiagnosisRequest) GetVisitId() string {
	if x != nil {
		return x.VisitId
	}
	return ""
}

func (x *RecordDiagnosisRequest) GetDiagnosis() string {
	if x != nil {
		return x.Diagnosis
	}
	return ""
}

type RecordDiagnosisResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecordDiagnosisResponse) Reset() {
	*x = RecordDiagnosisResponse{}
	mi := &file_opd_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecordDiagnosisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordDiagnosisResponse) ProtoMessage() {}

func (x *RecordDiagnosisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opd_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordDiagnosisResponse.ProtoReflect.Descriptor instead.
func (*RecordDiagnosisResponse) Descriptor() ([]byte, []int) {
	return file_opd_proto_rawDescGZIP(), []int{5}
}

func (x *RecordDiagnosisResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PrescribeTestsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VisitId       string                 `protobuf:"bytes,1,opt,name=visit_id,json=visitId,proto3" json:"visit_id,omitempty"`
	TestNames     []string               `protobuf:"bytes,2,rep,name=test_names,json=testNames,proto3" json:"test_names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrescribeTestsRequest) Reset() {
	*x = PrescribeTestsRequest{}
	mi := &file_opd_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrescribeTestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrescribeTestsRequest) ProtoMessage() {}

func (x *PrescribeTestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opd_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrescribeTestsRequest.ProtoReflect.Descriptor instead.
func (*PrescribeTestsRequest) Descriptor() ([]byte, []int) {
	return file_opd_proto_rawDescGZIP(), []int{6}
}

func (x *PrescribeTestsRequest) GetVisitId() string {
	if x != nil {
		return x.VisitId
	}
	return ""
}

func (x *PrescribeTestsRequest) GetTestNames() []string {
	if x != nil {
		return x.TestNames
	}
	return nil
}

type PrescribeTestsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrescribeTestsResponse) Reset() {
	*x = PrescribeTestsResponse{}
	mi := &file_opd_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrescribeTestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrescribeTestsResponse) ProtoMessage() {}

func (x *PrescribeTestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opd_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrescribeTestsResponse.ProtoReflect.Descriptor instead.
func (*PrescribeTestsResponse) Descriptor() ([]byte, []int) {
	return file_opd_proto_rawDescGZIP(), []int{7}
}

func (x *PrescribeTestsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GeneratePrescriptionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VisitId       string                 `protobuf:"bytes,1,opt,name=visit_id,json=visitId,proto3" json:"visit_id,omitempty"`
	Medicines     []string               `protobuf:"bytes,2,rep,name=medicines,proto3" json:"medicines,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeneratePrescriptionRequest) Reset() {
	*x = GeneratePrescriptionRequest{}
	mi := &file_opd_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeneratePrescriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePrescriptionRequest) ProtoMessage() {}

func (x *GeneratePrescriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opd_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePrescriptionRequest.ProtoReflect.Descriptor instead.
func (*GeneratePrescriptionRequest) Descriptor() ([]byte, []int) {
	return file_opd_proto_rawDescGZIP(), []int{8}
}

func (x *GeneratePrescriptionRequest) GetVisitId() string {
	if x != nil {
		return x.VisitId
	}
	return ""
}

func (x *GeneratePrescriptionRequest) GetMedicines() []string {
	if x != nil {
		return x.Medicines
	}
	return nil
}

type GeneratePrescriptionResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	PrescriptionId string                 `protobuf:"bytes,1,opt,name=prescription_id,json=prescriptionId,proto3" json:"prescription_id,omitempty"`
	Message        string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GeneratePrescriptionResponse) Reset() {
	*x = GeneratePrescriptionResponse{}
	mi := &file_opd_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeneratePrescriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePrescriptionResponse) ProtoMessage() {}

func (x *GeneratePrescriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opd_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePrescriptionResponse.ProtoReflect.Descriptor instead.
func (*GeneratePrescriptionResponse) Descriptor() ([]byte, []int) {
	return file_opd_proto_rawDescGZIP(), []int{9}
}

func (x *GeneratePrescriptionResponse) GetPrescriptionId() string {
	if x != nil {
		return x.PrescriptionId
	}
	return ""
}

func (x *GeneratePrescriptionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EndVisitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PatientId     string                 `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	VisitId       string                 `protobuf:"bytes,2,opt,name=visit_id,json=visitId,proto3" json:"visit_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EndVisitRequest) Reset() {
	*x = EndVisitRequest{}
	mi := &file_opd_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndVisitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndVisitRequest) ProtoMessage() {}

func (x *EndVisitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opd_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndVisitRequest.ProtoReflect.Descriptor instead.
func (*EndVisitRequest) Descriptor() ([]byte, []int) {
	return file_opd_proto_rawDescGZIP(), []int{10}
}

func (x *EndVisitRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *EndVisitRequest) GetVisitId() string {
	if x != nil {
		return x.VisitId
	}
	return ""
}

type EndVisitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BillId        string                 `protobuf:"bytes,1,opt,name=bill_id,json=billId,proto3" json:"bill_id,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EndVisitResponse) Reset() {
	*x = EndVisitResponse{}
	mi := &file_opd_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndVisitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndVisitResponse) ProtoMessage() {}

func (x *EndVisitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opd_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndVisitResponse.ProtoReflect.Descriptor instead.
func (*EndVisitResponse) Descriptor() ([]byte, []int) {
	return file_opd_proto_rawDescGZIP(), []int{11}
}

func (x *EndVisitResponse) GetBillId() string {
	if x != nil {
		return x.BillId
	}
	return ""
}

func (x *EndVisitResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_opd_proto protoreflect.FileDescriptor

const file_opd_proto_rawDesc = "" +
	"\n" +
	"\topd.proto\x12\x03opd\"_\n" +
	"\x17CheckAppointmentRequest\x12\x1d\n" +
	"\n" +
	"patient_id\x18\x01 \x01(\tR\tpatientId\x12%\n" +
	"\x0eappointment_id\x18\x02 \x01(\tR\rappointmentId\"J\n" +
	"\x18CheckAppointmentResponse\x12\x14\n" +
	"\x05valid\x18\x01 \x01(\bR\x05valid\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"}\n" +
	"\x18StartConsultationRequest\x12\x1d\n" +
	"\n" +
	"patient_id\x18\x01 \x01(\tR\tpatientId\x12\x1b\n" +
	"\tdoctor_id\x18\x02 \x01(\tR\bdoctorId\x12%\n" +
	"\x0eappointment_id\x18\x03 \x01(\tR\rappointmentId\"P\n" +
	"\x19StartConsultationResponse\x12\x19\n" +
	"\bvisit_id\x18\x01 \x01(\tR\avisitId\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"Q\n" +
	"\x16RecordDiagnosisRequest\x12\x19\n" +
	"\bvisit_id\x18\x01 \x01(\tR\avisitId\x12\x1c\n" +
	"\tdiagnosis\x18\x02 \x01(\tR\tdiagnosis\"3\n" +
	"\x17RecordDiagnosisResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"Q\n" +
	"\x15PrescribeTestsRequest\x12\x19\n" +
	"\bvisit_id\x18\x01 \x01(\tR\avisitId\x12\x1d\n" +
	"\n" +
	"test_names\x18\x02 \x03(\tR\ttestNames\"2\n" +
	"\x16PrescribeTestsResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"V\n" +
	"\x1bGeneratePrescriptionRequest\x12\x19\n" +
	"\bvisit_id\x18\x01 \x01(\tR\avisitId\x12\x1c\n" +
	"\tmedicines\x18\x02 \x03(\tR\tmedicines\"a\n" +
	"\x1cGeneratePrescriptionResponse\x12'\n" +
	"\x0fprescription_id\x18\x01 \x01(\tR\x0eprescriptionId\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"K\n" +
	"\x0fEndVisitRequest\x12\x1d\n" +
	"\n" +
	"patient_id\x18\x01 \x01(\tR\tpatientId\x12\x19\n" +
	"\bvisit_id\x18\x02 \x01(\tR\avisitId\"E\n" +
	"\x10EndVisitResponse\x12\x17\n" +
	"\abill_id\x18\x01 \x01(\tR\x06billId\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2\xe0\x03\n" +
	"\n" +
	"OPDService\x12O\n" +
	"\x10CheckAppointment\x12\x1c.opd.CheckAppointmentRequest\x1a\x1d.opd.CheckAppointmentResponse\x12R\n" +
	"\x11StartConsultation\x12\x1d.opd.StartConsultationRequest\x1a\x1e.opd.StartConsultationResponse\x12L\n" +
	"\x0fRecordDiagnosis\x12\x1b.opd.RecordDiagnosisRequest\x1a\x1c.opd.RecordDiagnosisResponse\x12I\n" +
	"\x0ePrescribeTests\x12\x1a.opd.PrescribeTestsRequest\x1a\x1b.opd.PrescribeTestsResponse\x12[\n" +
	"\x14GeneratePrescription\x12 .opd.GeneratePrescriptionRequest\x1a!.opd.GeneratePrescriptionResponse\x127\n" +
	"\bEndVisit\x12\x14.opd.EndVisitRequest\x1a\x15.opd.EndVisitResponseBGZEgithub.com/somnathbm/hospital-hms/microservices/opd-service/gen;opdpbb\x06proto3"

var (
	file_opd_proto_rawDescOnce sync.Once
	file_opd_proto_rawDescData []byte
)

func file_opd_proto_rawDescGZIP() []byte {
	file_opd_proto_rawDescOnce.Do(func() {
		file_opd_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_opd_proto_rawDesc), len(file_opd_proto_rawDesc)))
	})
	return file_opd_proto_rawDescData
}

var file_opd_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_opd_proto_goTypes = []any{
	(*CheckAppointmentRequest)(nil),      // 0: opd.CheckAppointmentRequest
	(*CheckAppointmentResponse)(nil),     // 1: opd.CheckAppointmentResponse
	(*StartConsultationRequest)(nil),     // 2: opd.StartConsultationRequest
	(*StartConsultationResponse)(nil),    // 3: opd.StartConsultationResponse
	(*RecordDiagnosisRequest)(nil),       // 4: opd.RecordDiagnosisRequest
	(*RecordDiagnosisResponse)(nil),      // 5: opd.RecordDiagnosisResponse
	(*PrescribeTestsRequest)(nil),        // 6: opd.PrescribeTestsRequest
	(*PrescribeTestsResponse)(nil),       // 7: opd.PrescribeTestsResponse
	(*GeneratePrescriptionRequest)(nil),  // 8: opd.GeneratePrescriptionRequest
	(*GeneratePrescriptionResponse)(nil), // 9: opd.GeneratePrescriptionResponse
	(*EndVisitRequest)(nil),              // 10: opd.EndVisitRequest
	(*EndVisitResponse)(nil),             // 11: opd.EndVisitResponse
}
var file_opd_proto_depIdxs = []int32{
	0,  // 0: opd.OPDService.CheckAppointment:input_type -> opd.CheckAppointmentRequest
	2,  // 1: opd.OPDService.StartConsultation:input_type -> opd.StartConsultationRequest
	4,  // 2: opd.OPDService.RecordDiagnosis:input_type -> opd.RecordDiagnosisRequest
	6,  // 3: opd.OPDService.PrescribeTests:input_type -> opd.PrescribeTestsRequest
	8,  // 4: opd.OPDService.GeneratePrescription:input_type -> opd.GeneratePrescriptionRequest
	10, // 5: opd.OPDService.EndVisit:input_type -> opd.EndVisitRequest
	1,  // 6: opd.OPDService.CheckAppointment:output_type -> opd.CheckAppointmentResponse
	3,  // 7: opd.OPDService.StartConsultation:output_type -> opd.StartConsultationResponse
	5,  // 8: opd.OPDService.RecordDiagnosis:output_type -> opd.RecordDiagnosisResponse
	7,  // 9: opd.OPDService.PrescribeTests:output_type -> opd.PrescribeTestsResponse
	9,  // 10: opd.OPDService.GeneratePrescription:output_type -> opd.GeneratePrescriptionResponse
	11, // 11: opd.OPDService.EndVisit:output_type -> opd.EndVisitResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_opd_proto_init() }
func file_opd_proto_init() {
	if File_opd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_opd_proto_rawDesc), len(file_opd_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_opd_proto_goTypes,
		DependencyIndexes: file_opd_proto_depIdxs,
		MessageInfos:      file_opd_proto_msgTypes,
	}.Build()
	File_opd_proto = out.File
	file_opd_proto_goTypes = nil
	file_opd_proto_depIdxs = nil
}
