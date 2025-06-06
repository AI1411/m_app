// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: user/v1/user.proto

package userv1

import (
	v11 "github.com/AI1411/m_app/gen/education/v1"
	v1 "github.com/AI1411/m_app/gen/prefecture/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// ユーザー情報取得リクエスト
type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_user_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ユーザー情報レスポンス
type GetUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	mi := &file_user_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// ユーザー検索リクエスト
type SearchUsersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Gender        *string                `protobuf:"bytes,3,opt,name=gender,proto3,oneof" json:"gender,omitempty"`
	MinAge        *int32                 `protobuf:"varint,4,opt,name=min_age,json=minAge,proto3,oneof" json:"min_age,omitempty"`
	MaxAge        *int32                 `protobuf:"varint,5,opt,name=max_age,json=maxAge,proto3,oneof" json:"max_age,omitempty"`
	PrefectureId  *int32                 `protobuf:"varint,6,opt,name=prefecture_id,json=prefectureId,proto3,oneof" json:"prefecture_id,omitempty"` // 都道府県IDで検索
	RegionId      *int32                 `protobuf:"varint,7,opt,name=region_id,json=regionId,proto3,oneof" json:"region_id,omitempty"`             // リージョンIDで検索
	Interests     []string               `protobuf:"bytes,8,rep,name=interests,proto3" json:"interests,omitempty"`
	EducationId   *int32                 `protobuf:"varint,9,opt,name=education_id,json=educationId,proto3,oneof" json:"education_id,omitempty"` // 学歴IDで検索
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchUsersRequest) Reset() {
	*x = SearchUsersRequest{}
	mi := &file_user_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUsersRequest) ProtoMessage() {}

func (x *SearchUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUsersRequest.ProtoReflect.Descriptor instead.
func (*SearchUsersRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *SearchUsersRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchUsersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchUsersRequest) GetGender() string {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return ""
}

func (x *SearchUsersRequest) GetMinAge() int32 {
	if x != nil && x.MinAge != nil {
		return *x.MinAge
	}
	return 0
}

func (x *SearchUsersRequest) GetMaxAge() int32 {
	if x != nil && x.MaxAge != nil {
		return *x.MaxAge
	}
	return 0
}

func (x *SearchUsersRequest) GetPrefectureId() int32 {
	if x != nil && x.PrefectureId != nil {
		return *x.PrefectureId
	}
	return 0
}

func (x *SearchUsersRequest) GetRegionId() int32 {
	if x != nil && x.RegionId != nil {
		return *x.RegionId
	}
	return 0
}

func (x *SearchUsersRequest) GetInterests() []string {
	if x != nil {
		return x.Interests
	}
	return nil
}

func (x *SearchUsersRequest) GetEducationId() int32 {
	if x != nil && x.EducationId != nil {
		return *x.EducationId
	}
	return 0
}

// ユーザー検索レスポンス
type SearchUsersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*UserPreview         `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Page          int32                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchUsersResponse) Reset() {
	*x = SearchUsersResponse{}
	mi := &file_user_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUsersResponse) ProtoMessage() {}

func (x *SearchUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUsersResponse.ProtoReflect.Descriptor instead.
func (*SearchUsersResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *SearchUsersResponse) GetUsers() []*UserPreview {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *SearchUsersResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *SearchUsersResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchUsersResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// ユーザー作成リクエスト
type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Nickname      string                 `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	BirthDate     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Gender        string                 `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	PrefectureId  int32                  `protobuf:"varint,7,opt,name=prefecture_id,json=prefectureId,proto3" json:"prefecture_id,omitempty"` // 都道府県ID
	AboutMe       string                 `protobuf:"bytes,8,opt,name=about_me,json=aboutMe,proto3" json:"about_me,omitempty"`
	Interests     []string               `protobuf:"bytes,9,rep,name=interests,proto3" json:"interests,omitempty"`
	LookingFor    string                 `protobuf:"bytes,10,opt,name=looking_for,json=lookingFor,proto3" json:"looking_for,omitempty"`
	JobTitle      string                 `protobuf:"bytes,11,opt,name=job_title,json=jobTitle,proto3" json:"job_title,omitempty"`           // 職業・職種
	Company       string                 `protobuf:"bytes,12,opt,name=company,proto3" json:"company,omitempty"`                             // 会社名・組織名
	EducationId   int32                  `protobuf:"varint,13,opt,name=education_id,json=educationId,proto3" json:"education_id,omitempty"` // 学歴ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_user_v1_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *CreateUserRequest) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *CreateUserRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CreateUserRequest) GetPrefectureId() int32 {
	if x != nil {
		return x.PrefectureId
	}
	return 0
}

func (x *CreateUserRequest) GetAboutMe() string {
	if x != nil {
		return x.AboutMe
	}
	return ""
}

func (x *CreateUserRequest) GetInterests() []string {
	if x != nil {
		return x.Interests
	}
	return nil
}

func (x *CreateUserRequest) GetLookingFor() string {
	if x != nil {
		return x.LookingFor
	}
	return ""
}

func (x *CreateUserRequest) GetJobTitle() string {
	if x != nil {
		return x.JobTitle
	}
	return ""
}

func (x *CreateUserRequest) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *CreateUserRequest) GetEducationId() int32 {
	if x != nil {
		return x.EducationId
	}
	return 0
}

// ユーザー作成レスポンス
type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_user_v1_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ユーザー更新リクエスト
type UpdateUserRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Nickname        *string                `protobuf:"bytes,3,opt,name=nickname,proto3,oneof" json:"nickname,omitempty"`
	ProfileImageUrl *string                `protobuf:"bytes,4,opt,name=profile_image_url,json=profileImageUrl,proto3,oneof" json:"profile_image_url,omitempty"`
	AboutMe         *string                `protobuf:"bytes,5,opt,name=about_me,json=aboutMe,proto3,oneof" json:"about_me,omitempty"`
	PrefectureId    *int32                 `protobuf:"varint,6,opt,name=prefecture_id,json=prefectureId,proto3,oneof" json:"prefecture_id,omitempty"` // 都道府県ID
	JobTitle        *string                `protobuf:"bytes,7,opt,name=job_title,json=jobTitle,proto3,oneof" json:"job_title,omitempty"`
	Company         *string                `protobuf:"bytes,8,opt,name=company,proto3,oneof" json:"company,omitempty"`
	EducationId     *int32                 `protobuf:"varint,9,opt,name=education_id,json=educationId,proto3,oneof" json:"education_id,omitempty"` // 学歴ID
	Interests       []string               `protobuf:"bytes,10,rep,name=interests,proto3" json:"interests,omitempty"`
	LookingFor      *string                `protobuf:"bytes,11,opt,name=looking_for,json=lookingFor,proto3,oneof" json:"looking_for,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_user_v1_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUserRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateUserRequest) GetNickname() string {
	if x != nil && x.Nickname != nil {
		return *x.Nickname
	}
	return ""
}

func (x *UpdateUserRequest) GetProfileImageUrl() string {
	if x != nil && x.ProfileImageUrl != nil {
		return *x.ProfileImageUrl
	}
	return ""
}

func (x *UpdateUserRequest) GetAboutMe() string {
	if x != nil && x.AboutMe != nil {
		return *x.AboutMe
	}
	return ""
}

func (x *UpdateUserRequest) GetPrefectureId() int32 {
	if x != nil && x.PrefectureId != nil {
		return *x.PrefectureId
	}
	return 0
}

func (x *UpdateUserRequest) GetJobTitle() string {
	if x != nil && x.JobTitle != nil {
		return *x.JobTitle
	}
	return ""
}

func (x *UpdateUserRequest) GetCompany() string {
	if x != nil && x.Company != nil {
		return *x.Company
	}
	return ""
}

func (x *UpdateUserRequest) GetEducationId() int32 {
	if x != nil && x.EducationId != nil {
		return *x.EducationId
	}
	return 0
}

func (x *UpdateUserRequest) GetInterests() []string {
	if x != nil {
		return x.Interests
	}
	return nil
}

func (x *UpdateUserRequest) GetLookingFor() string {
	if x != nil && x.LookingFor != nil {
		return *x.LookingFor
	}
	return ""
}

// ユーザー更新レスポンス
type UpdateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	mi := &file_user_v1_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// ユーザー情報（完全版）
type User struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email           string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name            string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Nickname        string                 `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	BirthDate       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Age             int32                  `protobuf:"varint,6,opt,name=age,proto3" json:"age,omitempty"` // 計算フィールド
	Gender          string                 `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	ProfileImageUrl string                 `protobuf:"bytes,8,opt,name=profile_image_url,json=profileImageUrl,proto3" json:"profile_image_url,omitempty"`
	AboutMe         string                 `protobuf:"bytes,9,opt,name=about_me,json=aboutMe,proto3" json:"about_me,omitempty"`
	Prefecture      *v1.Prefecture         `protobuf:"bytes,10,opt,name=prefecture,proto3" json:"prefecture,omitempty"` // 都道府県情報
	JobTitle        string                 `protobuf:"bytes,11,opt,name=job_title,json=jobTitle,proto3" json:"job_title,omitempty"`
	Company         string                 `protobuf:"bytes,12,opt,name=company,proto3" json:"company,omitempty"`
	Education       *v11.Education         `protobuf:"bytes,13,opt,name=education,proto3" json:"education,omitempty"` // 学歴情報
	Interests       []string               `protobuf:"bytes,14,rep,name=interests,proto3" json:"interests,omitempty"`
	LookingFor      string                 `protobuf:"bytes,15,opt,name=looking_for,json=lookingFor,proto3" json:"looking_for,omitempty"`
	LastActive      *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=last_active,json=lastActive,proto3" json:"last_active,omitempty"`
	IsVerified      bool                   `protobuf:"varint,17,opt,name=is_verified,json=isVerified,proto3" json:"is_verified,omitempty"`
	IsPremium       bool                   `protobuf:"varint,18,opt,name=is_premium,json=isPremium,proto3" json:"is_premium,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // 更新日時を追加
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_user_v1_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *User) GetProfileImageUrl() string {
	if x != nil {
		return x.ProfileImageUrl
	}
	return ""
}

func (x *User) GetAboutMe() string {
	if x != nil {
		return x.AboutMe
	}
	return ""
}

func (x *User) GetPrefecture() *v1.Prefecture {
	if x != nil {
		return x.Prefecture
	}
	return nil
}

func (x *User) GetJobTitle() string {
	if x != nil {
		return x.JobTitle
	}
	return ""
}

func (x *User) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *User) GetEducation() *v11.Education {
	if x != nil {
		return x.Education
	}
	return nil
}

func (x *User) GetInterests() []string {
	if x != nil {
		return x.Interests
	}
	return nil
}

func (x *User) GetLookingFor() string {
	if x != nil {
		return x.LookingFor
	}
	return ""
}

func (x *User) GetLastActive() *timestamppb.Timestamp {
	if x != nil {
		return x.LastActive
	}
	return nil
}

func (x *User) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

func (x *User) GetIsPremium() bool {
	if x != nil {
		return x.IsPremium
	}
	return false
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// ユーザープレビュー（検索結果用の簡易版）
type UserPreview struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Nickname        string                 `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Age             int32                  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Gender          string                 `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	ProfileImageUrl string                 `protobuf:"bytes,6,opt,name=profile_image_url,json=profileImageUrl,proto3" json:"profile_image_url,omitempty"`
	Prefecture      *v1.Prefecture         `protobuf:"bytes,7,opt,name=prefecture,proto3" json:"prefecture,omitempty"` // 都道府県情報
	Interests       []string               `protobuf:"bytes,8,rep,name=interests,proto3" json:"interests,omitempty"`
	JobTitle        string                 `protobuf:"bytes,9,opt,name=job_title,json=jobTitle,proto3" json:"job_title,omitempty"` // 職業情報を追加
	Education       *v11.Education         `protobuf:"bytes,10,opt,name=education,proto3" json:"education,omitempty"`              // 学歴情報を追加
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *UserPreview) Reset() {
	*x = UserPreview{}
	mi := &file_user_v1_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserPreview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPreview) ProtoMessage() {}

func (x *UserPreview) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPreview.ProtoReflect.Descriptor instead.
func (*UserPreview) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{9}
}

func (x *UserPreview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserPreview) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserPreview) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserPreview) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserPreview) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserPreview) GetProfileImageUrl() string {
	if x != nil {
		return x.ProfileImageUrl
	}
	return ""
}

func (x *UserPreview) GetPrefecture() *v1.Prefecture {
	if x != nil {
		return x.Prefecture
	}
	return nil
}

func (x *UserPreview) GetInterests() []string {
	if x != nil {
		return x.Interests
	}
	return nil
}

func (x *UserPreview) GetJobTitle() string {
	if x != nil {
		return x.JobTitle
	}
	return ""
}

func (x *UserPreview) GetEducation() *v11.Education {
	if x != nil {
		return x.Education
	}
	return nil
}

var File_user_v1_user_proto protoreflect.FileDescriptor

const file_user_v1_user_proto_rawDesc = "" +
	"\n" +
	"\x12user/v1/user.proto\x12\auser.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1eprefecture/v1/prefecture.proto\x1a\x1ceducation/v1/education.proto\" \n" +
	"\x0eGetUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"4\n" +
	"\x0fGetUserResponse\x12!\n" +
	"\x04user\x18\x01 \x01(\v2\r.user.v1.UserR\x04user\"\x84\x03\n" +
	"\x12SearchUsersRequest\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x05R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x05R\bpageSize\x12\x1b\n" +
	"\x06gender\x18\x03 \x01(\tH\x00R\x06gender\x88\x01\x01\x12\x1c\n" +
	"\amin_age\x18\x04 \x01(\x05H\x01R\x06minAge\x88\x01\x01\x12\x1c\n" +
	"\amax_age\x18\x05 \x01(\x05H\x02R\x06maxAge\x88\x01\x01\x12(\n" +
	"\rprefecture_id\x18\x06 \x01(\x05H\x03R\fprefectureId\x88\x01\x01\x12 \n" +
	"\tregion_id\x18\a \x01(\x05H\x04R\bregionId\x88\x01\x01\x12\x1c\n" +
	"\tinterests\x18\b \x03(\tR\tinterests\x12&\n" +
	"\feducation_id\x18\t \x01(\x05H\x05R\veducationId\x88\x01\x01B\t\n" +
	"\a_genderB\n" +
	"\n" +
	"\b_min_ageB\n" +
	"\n" +
	"\b_max_ageB\x10\n" +
	"\x0e_prefecture_idB\f\n" +
	"\n" +
	"_region_idB\x0f\n" +
	"\r_education_id\"\x93\x01\n" +
	"\x13SearchUsersResponse\x12*\n" +
	"\x05users\x18\x01 \x03(\v2\x14.user.v1.UserPreviewR\x05users\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\x12\x12\n" +
	"\x04page\x18\x03 \x01(\x05R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x04 \x01(\x05R\bpageSize\"\xa1\x03\n" +
	"\x11CreateUserRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x1a\n" +
	"\bnickname\x18\x04 \x01(\tR\bnickname\x129\n" +
	"\n" +
	"birth_date\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tbirthDate\x12\x16\n" +
	"\x06gender\x18\x06 \x01(\tR\x06gender\x12#\n" +
	"\rprefecture_id\x18\a \x01(\x05R\fprefectureId\x12\x19\n" +
	"\babout_me\x18\b \x01(\tR\aaboutMe\x12\x1c\n" +
	"\tinterests\x18\t \x03(\tR\tinterests\x12\x1f\n" +
	"\vlooking_for\x18\n" +
	" \x01(\tR\n" +
	"lookingFor\x12\x1b\n" +
	"\tjob_title\x18\v \x01(\tR\bjobTitle\x12\x18\n" +
	"\acompany\x18\f \x01(\tR\acompany\x12!\n" +
	"\feducation_id\x18\r \x01(\x05R\veducationId\"$\n" +
	"\x12CreateUserResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"\x8b\x04\n" +
	"\x11UpdateUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\x04name\x18\x02 \x01(\tH\x00R\x04name\x88\x01\x01\x12\x1f\n" +
	"\bnickname\x18\x03 \x01(\tH\x01R\bnickname\x88\x01\x01\x12/\n" +
	"\x11profile_image_url\x18\x04 \x01(\tH\x02R\x0fprofileImageUrl\x88\x01\x01\x12\x1e\n" +
	"\babout_me\x18\x05 \x01(\tH\x03R\aaboutMe\x88\x01\x01\x12(\n" +
	"\rprefecture_id\x18\x06 \x01(\x05H\x04R\fprefectureId\x88\x01\x01\x12 \n" +
	"\tjob_title\x18\a \x01(\tH\x05R\bjobTitle\x88\x01\x01\x12\x1d\n" +
	"\acompany\x18\b \x01(\tH\x06R\acompany\x88\x01\x01\x12&\n" +
	"\feducation_id\x18\t \x01(\x05H\aR\veducationId\x88\x01\x01\x12\x1c\n" +
	"\tinterests\x18\n" +
	" \x03(\tR\tinterests\x12$\n" +
	"\vlooking_for\x18\v \x01(\tH\bR\n" +
	"lookingFor\x88\x01\x01B\a\n" +
	"\x05_nameB\v\n" +
	"\t_nicknameB\x14\n" +
	"\x12_profile_image_urlB\v\n" +
	"\t_about_meB\x10\n" +
	"\x0e_prefecture_idB\f\n" +
	"\n" +
	"_job_titleB\n" +
	"\n" +
	"\b_companyB\x0f\n" +
	"\r_education_idB\x0e\n" +
	"\f_looking_for\".\n" +
	"\x12UpdateUserResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"\xe3\x05\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x1a\n" +
	"\bnickname\x18\x04 \x01(\tR\bnickname\x129\n" +
	"\n" +
	"birth_date\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tbirthDate\x12\x10\n" +
	"\x03age\x18\x06 \x01(\x05R\x03age\x12\x16\n" +
	"\x06gender\x18\a \x01(\tR\x06gender\x12*\n" +
	"\x11profile_image_url\x18\b \x01(\tR\x0fprofileImageUrl\x12\x19\n" +
	"\babout_me\x18\t \x01(\tR\aaboutMe\x129\n" +
	"\n" +
	"prefecture\x18\n" +
	" \x01(\v2\x19.prefecture.v1.PrefectureR\n" +
	"prefecture\x12\x1b\n" +
	"\tjob_title\x18\v \x01(\tR\bjobTitle\x12\x18\n" +
	"\acompany\x18\f \x01(\tR\acompany\x125\n" +
	"\teducation\x18\r \x01(\v2\x17.education.v1.EducationR\teducation\x12\x1c\n" +
	"\tinterests\x18\x0e \x03(\tR\tinterests\x12\x1f\n" +
	"\vlooking_for\x18\x0f \x01(\tR\n" +
	"lookingFor\x12;\n" +
	"\vlast_active\x18\x10 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"lastActive\x12\x1f\n" +
	"\vis_verified\x18\x11 \x01(\bR\n" +
	"isVerified\x12\x1d\n" +
	"\n" +
	"is_premium\x18\x12 \x01(\bR\tisPremium\x129\n" +
	"\n" +
	"created_at\x18\x13 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x14 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"\xd0\x02\n" +
	"\vUserPreview\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bnickname\x18\x03 \x01(\tR\bnickname\x12\x10\n" +
	"\x03age\x18\x04 \x01(\x05R\x03age\x12\x16\n" +
	"\x06gender\x18\x05 \x01(\tR\x06gender\x12*\n" +
	"\x11profile_image_url\x18\x06 \x01(\tR\x0fprofileImageUrl\x129\n" +
	"\n" +
	"prefecture\x18\a \x01(\v2\x19.prefecture.v1.PrefectureR\n" +
	"prefecture\x12\x1c\n" +
	"\tinterests\x18\b \x03(\tR\tinterests\x12\x1b\n" +
	"\tjob_title\x18\t \x01(\tR\bjobTitle\x125\n" +
	"\teducation\x18\n" +
	" \x01(\v2\x17.education.v1.EducationR\teducation2\xab\x02\n" +
	"\vUserService\x12>\n" +
	"\aGetUser\x12\x17.user.v1.GetUserRequest\x1a\x18.user.v1.GetUserResponse\"\x00\x12J\n" +
	"\vSearchUsers\x12\x1b.user.v1.SearchUsersRequest\x1a\x1c.user.v1.SearchUsersResponse\"\x00\x12G\n" +
	"\n" +
	"CreateUser\x12\x1a.user.v1.CreateUserRequest\x1a\x1b.user.v1.CreateUserResponse\"\x00\x12G\n" +
	"\n" +
	"UpdateUser\x12\x1a.user.v1.UpdateUserRequest\x1a\x1b.user.v1.UpdateUserResponse\"\x00B\x81\x01\n" +
	"\vcom.user.v1B\tUserProtoP\x01Z*github.com/AI1411/m_app/gen/user/v1;userv1\xa2\x02\x03UXX\xaa\x02\aUser.V1\xca\x02\aUser\\V1\xe2\x02\x13User\\V1\\GPBMetadata\xea\x02\bUser::V1b\x06proto3"

var (
	file_user_v1_user_proto_rawDescOnce sync.Once
	file_user_v1_user_proto_rawDescData []byte
)

func file_user_v1_user_proto_rawDescGZIP() []byte {
	file_user_v1_user_proto_rawDescOnce.Do(func() {
		file_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_v1_user_proto_rawDesc), len(file_user_v1_user_proto_rawDesc)))
	})
	return file_user_v1_user_proto_rawDescData
}

var file_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_v1_user_proto_goTypes = []any{
	(*GetUserRequest)(nil),        // 0: user.v1.GetUserRequest
	(*GetUserResponse)(nil),       // 1: user.v1.GetUserResponse
	(*SearchUsersRequest)(nil),    // 2: user.v1.SearchUsersRequest
	(*SearchUsersResponse)(nil),   // 3: user.v1.SearchUsersResponse
	(*CreateUserRequest)(nil),     // 4: user.v1.CreateUserRequest
	(*CreateUserResponse)(nil),    // 5: user.v1.CreateUserResponse
	(*UpdateUserRequest)(nil),     // 6: user.v1.UpdateUserRequest
	(*UpdateUserResponse)(nil),    // 7: user.v1.UpdateUserResponse
	(*User)(nil),                  // 8: user.v1.User
	(*UserPreview)(nil),           // 9: user.v1.UserPreview
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*v1.Prefecture)(nil),         // 11: prefecture.v1.Prefecture
	(*v11.Education)(nil),         // 12: education.v1.Education
}
var file_user_v1_user_proto_depIdxs = []int32{
	8,  // 0: user.v1.GetUserResponse.user:type_name -> user.v1.User
	9,  // 1: user.v1.SearchUsersResponse.users:type_name -> user.v1.UserPreview
	10, // 2: user.v1.CreateUserRequest.birth_date:type_name -> google.protobuf.Timestamp
	10, // 3: user.v1.User.birth_date:type_name -> google.protobuf.Timestamp
	11, // 4: user.v1.User.prefecture:type_name -> prefecture.v1.Prefecture
	12, // 5: user.v1.User.education:type_name -> education.v1.Education
	10, // 6: user.v1.User.last_active:type_name -> google.protobuf.Timestamp
	10, // 7: user.v1.User.created_at:type_name -> google.protobuf.Timestamp
	10, // 8: user.v1.User.updated_at:type_name -> google.protobuf.Timestamp
	11, // 9: user.v1.UserPreview.prefecture:type_name -> prefecture.v1.Prefecture
	12, // 10: user.v1.UserPreview.education:type_name -> education.v1.Education
	0,  // 11: user.v1.UserService.GetUser:input_type -> user.v1.GetUserRequest
	2,  // 12: user.v1.UserService.SearchUsers:input_type -> user.v1.SearchUsersRequest
	4,  // 13: user.v1.UserService.CreateUser:input_type -> user.v1.CreateUserRequest
	6,  // 14: user.v1.UserService.UpdateUser:input_type -> user.v1.UpdateUserRequest
	1,  // 15: user.v1.UserService.GetUser:output_type -> user.v1.GetUserResponse
	3,  // 16: user.v1.UserService.SearchUsers:output_type -> user.v1.SearchUsersResponse
	5,  // 17: user.v1.UserService.CreateUser:output_type -> user.v1.CreateUserResponse
	7,  // 18: user.v1.UserService.UpdateUser:output_type -> user.v1.UpdateUserResponse
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_user_v1_user_proto_init() }
func file_user_v1_user_proto_init() {
	if File_user_v1_user_proto != nil {
		return
	}
	file_user_v1_user_proto_msgTypes[2].OneofWrappers = []any{}
	file_user_v1_user_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_v1_user_proto_rawDesc), len(file_user_v1_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_v1_user_proto_goTypes,
		DependencyIndexes: file_user_v1_user_proto_depIdxs,
		MessageInfos:      file_user_v1_user_proto_msgTypes,
	}.Build()
	File_user_v1_user_proto = out.File
	file_user_v1_user_proto_goTypes = nil
	file_user_v1_user_proto_depIdxs = nil
}
