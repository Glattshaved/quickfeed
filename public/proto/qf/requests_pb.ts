// @generated by protoc-gen-es v0.1.1 with parameter "target=ts"
// @generated from file qf/requests.proto (package qf, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {Message, proto3, protoInt64} from "@bufbuild/protobuf";
import {Enrollment_UserStatus, Review, Submission_Status, Submissions} from "./types_pb.js";

/**
 * @generated from message qf.CourseSubmissions
 */
export class CourseSubmissions extends Message<CourseSubmissions> {
  /**
   * @generated from field: map<uint64, qf.Submissions> submissions = 1;
   */
  submissions: { [key: string]: Submissions } = {};

  constructor(data?: PartialMessage<CourseSubmissions>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.CourseSubmissions";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "submissions", kind: "map", K: 4 /* ScalarType.UINT64 */, V: {kind: "message", T: Submissions} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CourseSubmissions {
    return new CourseSubmissions().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CourseSubmissions {
    return new CourseSubmissions().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CourseSubmissions {
    return new CourseSubmissions().fromJsonString(jsonString, options);
  }

  static equals(a: CourseSubmissions | PlainMessage<CourseSubmissions> | undefined, b: CourseSubmissions | PlainMessage<CourseSubmissions> | undefined): boolean {
    return proto3.util.equals(CourseSubmissions, a, b);
  }
}

/**
 * @generated from message qf.ReviewRequest
 */
export class ReviewRequest extends Message<ReviewRequest> {
  /**
   * @generated from field: uint64 courseID = 1;
   */
  courseID = protoInt64.zero;

  /**
   * @generated from field: qf.Review review = 2;
   */
  review?: Review;

  constructor(data?: PartialMessage<ReviewRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.ReviewRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "courseID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "review", kind: "message", T: Review },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ReviewRequest {
    return new ReviewRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ReviewRequest {
    return new ReviewRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ReviewRequest {
    return new ReviewRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ReviewRequest | PlainMessage<ReviewRequest> | undefined, b: ReviewRequest | PlainMessage<ReviewRequest> | undefined): boolean {
    return proto3.util.equals(ReviewRequest, a, b);
  }
}

/**
 * @generated from message qf.CourseRequest
 */
export class CourseRequest extends Message<CourseRequest> {
  /**
   * @generated from field: uint64 courseID = 1;
   */
  courseID = protoInt64.zero;

  constructor(data?: PartialMessage<CourseRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.CourseRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "courseID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CourseRequest {
    return new CourseRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CourseRequest {
    return new CourseRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CourseRequest {
    return new CourseRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CourseRequest | PlainMessage<CourseRequest> | undefined, b: CourseRequest | PlainMessage<CourseRequest> | undefined): boolean {
    return proto3.util.equals(CourseRequest, a, b);
  }
}

/**
 * @generated from message qf.GroupRequest
 */
export class GroupRequest extends Message<GroupRequest> {
  /**
   * @generated from field: uint64 courseID = 1;
   */
  courseID = protoInt64.zero;

  /**
   * @generated from field: uint64 userID = 2;
   */
  userID = protoInt64.zero;

  /**
   * @generated from field: uint64 groupID = 3;
   */
  groupID = protoInt64.zero;

  constructor(data?: PartialMessage<GroupRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.GroupRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "courseID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "userID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "groupID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GroupRequest {
    return new GroupRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GroupRequest {
    return new GroupRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GroupRequest {
    return new GroupRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GroupRequest | PlainMessage<GroupRequest> | undefined, b: GroupRequest | PlainMessage<GroupRequest> | undefined): boolean {
    return proto3.util.equals(GroupRequest, a, b);
  }
}

/**
 * @generated from message qf.OrgRequest
 */
export class OrgRequest extends Message<OrgRequest> {
  /**
   * @generated from field: string orgName = 1;
   */
  orgName = "";

  constructor(data?: PartialMessage<OrgRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.OrgRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "orgName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OrgRequest {
    return new OrgRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OrgRequest {
    return new OrgRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OrgRequest {
    return new OrgRequest().fromJsonString(jsonString, options);
  }

  static equals(a: OrgRequest | PlainMessage<OrgRequest> | undefined, b: OrgRequest | PlainMessage<OrgRequest> | undefined): boolean {
    return proto3.util.equals(OrgRequest, a, b);
  }
}

/**
 * @generated from message qf.Organization
 */
export class Organization extends Message<Organization> {
  /**
   * @generated from field: uint64 ID = 1;
   */
  ID = protoInt64.zero;

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string avatar = 3;
   */
  avatar = "";

  /**
   * @generated from field: string paymentPlan = 4;
   */
  paymentPlan = "";

  constructor(data?: PartialMessage<Organization>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.Organization";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "avatar", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "paymentPlan", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Organization {
    return new Organization().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Organization {
    return new Organization().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Organization {
    return new Organization().fromJsonString(jsonString, options);
  }

  static equals(a: Organization | PlainMessage<Organization> | undefined, b: Organization | PlainMessage<Organization> | undefined): boolean {
    return proto3.util.equals(Organization, a, b);
  }
}

/**
 * @generated from message qf.EnrollmentRequest
 */
export class EnrollmentRequest extends Message<EnrollmentRequest> {
  /**
   * @generated from oneof qf.EnrollmentRequest.FetchMode
   */
  FetchMode: {
    /**
     * @generated from field: uint64 courseID = 1;
     */
    value: bigint;
    case: "courseID";
  } | {
    /**
     * @generated from field: uint64 userID = 2;
     */
    value: bigint;
    case: "userID";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * @generated from field: repeated qf.Enrollment.UserStatus statuses = 3;
   */
  statuses: Enrollment_UserStatus[] = [];

  constructor(data?: PartialMessage<EnrollmentRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.EnrollmentRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "courseID", kind: "scalar", T: 4 /* ScalarType.UINT64 */, oneof: "FetchMode" },
    { no: 2, name: "userID", kind: "scalar", T: 4 /* ScalarType.UINT64 */, oneof: "FetchMode" },
    { no: 3, name: "statuses", kind: "enum", T: proto3.getEnumType(Enrollment_UserStatus), repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EnrollmentRequest {
    return new EnrollmentRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EnrollmentRequest {
    return new EnrollmentRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EnrollmentRequest {
    return new EnrollmentRequest().fromJsonString(jsonString, options);
  }

  static equals(a: EnrollmentRequest | PlainMessage<EnrollmentRequest> | undefined, b: EnrollmentRequest | PlainMessage<EnrollmentRequest> | undefined): boolean {
    return proto3.util.equals(EnrollmentRequest, a, b);
  }
}

/**
 * @generated from message qf.SubmissionRequest
 */
export class SubmissionRequest extends Message<SubmissionRequest> {
  /**
   * @generated from field: uint64 CourseID = 1;
   */
  CourseID = protoInt64.zero;

  /**
   * only used for user and group submissions
   *
   * @generated from field: uint64 AssignmentID = 2;
   */
  AssignmentID = protoInt64.zero;

  /**
   * @generated from oneof qf.SubmissionRequest.FetchMode
   */
  FetchMode: {
    /**
     * fetch single user's submissions with build info
     *
     * @generated from field: uint64 UserID = 3;
     */
    value: bigint;
    case: "UserID";
  } | {
    /**
     * fetch single group's submissions with build info
     *
     * @generated from field: uint64 GroupID = 4;
     */
    value: bigint;
    case: "GroupID";
  } | {
    /**
     * fetch single specific submission with build info
     *
     * @generated from field: uint64 SubmissionID = 5;
     */
    value: bigint;
    case: "SubmissionID";
  } | {
    /**
     * fetch all submissions of given type without build info
     *
     * @generated from field: qf.SubmissionRequest.SubmissionType Type = 6;
     */
    value: SubmissionRequest_SubmissionType;
    case: "Type";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<SubmissionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.SubmissionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "CourseID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "AssignmentID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "UserID", kind: "scalar", T: 4 /* ScalarType.UINT64 */, oneof: "FetchMode" },
    { no: 4, name: "GroupID", kind: "scalar", T: 4 /* ScalarType.UINT64 */, oneof: "FetchMode" },
    { no: 5, name: "SubmissionID", kind: "scalar", T: 4 /* ScalarType.UINT64 */, oneof: "FetchMode" },
    { no: 6, name: "Type", kind: "enum", T: proto3.getEnumType(SubmissionRequest_SubmissionType), oneof: "FetchMode" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubmissionRequest {
    return new SubmissionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubmissionRequest {
    return new SubmissionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubmissionRequest {
    return new SubmissionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SubmissionRequest | PlainMessage<SubmissionRequest> | undefined, b: SubmissionRequest | PlainMessage<SubmissionRequest> | undefined): boolean {
    return proto3.util.equals(SubmissionRequest, a, b);
  }
}

/**
 * @generated from enum qf.SubmissionRequest.SubmissionType
 */
export enum SubmissionRequest_SubmissionType {
  /**
   * fetch all submissions
   *
   * @generated from enum value: ALL = 0;
   */
  ALL = 0,

  /**
   * fetch all user submissions
   *
   * @generated from enum value: USER = 1;
   */
  USER = 1,

  /**
   * fetch all group submissions
   *
   * @generated from enum value: GROUP = 2;
   */
  GROUP = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(SubmissionRequest_SubmissionType)
proto3.util.setEnumType(SubmissionRequest_SubmissionType, "qf.SubmissionRequest.SubmissionType", [
  { no: 0, name: "ALL" },
  { no: 1, name: "USER" },
  { no: 2, name: "GROUP" },
]);

/**
 * @generated from message qf.UpdateSubmissionRequest
 */
export class UpdateSubmissionRequest extends Message<UpdateSubmissionRequest> {
  /**
   * @generated from field: uint64 submissionID = 1;
   */
  submissionID = protoInt64.zero;

  /**
   * @generated from field: uint64 courseID = 2;
   */
  courseID = protoInt64.zero;

  /**
   * @generated from field: uint32 score = 3;
   */
  score = 0;

  /**
   * @generated from field: bool released = 4;
   */
  released = false;

  /**
   * @generated from field: qf.Submission.Status status = 5;
   */
  status = Submission_Status.NONE;

  constructor(data?: PartialMessage<UpdateSubmissionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.UpdateSubmissionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "submissionID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "courseID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "score", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 4, name: "released", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "status", kind: "enum", T: proto3.getEnumType(Submission_Status) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateSubmissionRequest {
    return new UpdateSubmissionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateSubmissionRequest {
    return new UpdateSubmissionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateSubmissionRequest {
    return new UpdateSubmissionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateSubmissionRequest | PlainMessage<UpdateSubmissionRequest> | undefined, b: UpdateSubmissionRequest | PlainMessage<UpdateSubmissionRequest> | undefined): boolean {
    return proto3.util.equals(UpdateSubmissionRequest, a, b);
  }
}

/**
 * @generated from message qf.UpdateSubmissionsRequest
 */
export class UpdateSubmissionsRequest extends Message<UpdateSubmissionsRequest> {
  /**
   * @generated from field: uint64 courseID = 1;
   */
  courseID = protoInt64.zero;

  /**
   * @generated from field: uint64 assignmentID = 2;
   */
  assignmentID = protoInt64.zero;

  /**
   * @generated from field: uint32 scoreLimit = 3;
   */
  scoreLimit = 0;

  /**
   * @generated from field: bool release = 4;
   */
  release = false;

  /**
   * @generated from field: bool approve = 5;
   */
  approve = false;

  constructor(data?: PartialMessage<UpdateSubmissionsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.UpdateSubmissionsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "courseID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "assignmentID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "scoreLimit", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 4, name: "release", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "approve", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateSubmissionsRequest {
    return new UpdateSubmissionsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateSubmissionsRequest {
    return new UpdateSubmissionsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateSubmissionsRequest {
    return new UpdateSubmissionsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateSubmissionsRequest | PlainMessage<UpdateSubmissionsRequest> | undefined, b: UpdateSubmissionsRequest | PlainMessage<UpdateSubmissionsRequest> | undefined): boolean {
    return proto3.util.equals(UpdateSubmissionsRequest, a, b);
  }
}

/**
 * used to check whether student/group submission repo is empty
 *
 * @generated from message qf.RepositoryRequest
 */
export class RepositoryRequest extends Message<RepositoryRequest> {
  /**
   * @generated from field: uint64 userID = 1;
   */
  userID = protoInt64.zero;

  /**
   * @generated from field: uint64 groupID = 2;
   */
  groupID = protoInt64.zero;

  /**
   * @generated from field: uint64 courseID = 3;
   */
  courseID = protoInt64.zero;

  constructor(data?: PartialMessage<RepositoryRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.RepositoryRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "userID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "groupID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "courseID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RepositoryRequest {
    return new RepositoryRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RepositoryRequest {
    return new RepositoryRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RepositoryRequest {
    return new RepositoryRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RepositoryRequest | PlainMessage<RepositoryRequest> | undefined, b: RepositoryRequest | PlainMessage<RepositoryRequest> | undefined): boolean {
    return proto3.util.equals(RepositoryRequest, a, b);
  }
}

/**
 * @generated from message qf.Repositories
 */
export class Repositories extends Message<Repositories> {
  /**
   * Map key is the Repository.Type
   *
   * @generated from field: map<uint32, string> URLs = 1;
   */
  URLs: { [key: number]: string } = {};

  constructor(data?: PartialMessage<Repositories>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.Repositories";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "URLs", kind: "map", K: 13 /* ScalarType.UINT32 */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Repositories {
    return new Repositories().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Repositories {
    return new Repositories().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Repositories {
    return new Repositories().fromJsonString(jsonString, options);
  }

  static equals(a: Repositories | PlainMessage<Repositories> | undefined, b: Repositories | PlainMessage<Repositories> | undefined): boolean {
    return proto3.util.equals(Repositories, a, b);
  }
}

/**
 * @generated from message qf.Status
 */
export class Status extends Message<Status> {
  /**
   * @generated from field: uint64 Code = 1;
   */
  Code = protoInt64.zero;

  /**
   * @generated from field: string Error = 2;
   */
  Error = "";

  constructor(data?: PartialMessage<Status>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.Status";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Code", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "Error", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Status {
    return new Status().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Status {
    return new Status().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Status {
    return new Status().fromJsonString(jsonString, options);
  }

  static equals(a: Status | PlainMessage<Status> | undefined, b: Status | PlainMessage<Status> | undefined): boolean {
    return proto3.util.equals(Status, a, b);
  }
}

/**
 * @generated from message qf.RebuildRequest
 */
export class RebuildRequest extends Message<RebuildRequest> {
  /**
   * @generated from field: uint64 courseID = 1;
   */
  courseID = protoInt64.zero;

  /**
   * @generated from field: uint64 assignmentID = 2;
   */
  assignmentID = protoInt64.zero;

  /**
   * @generated from field: uint64 submissionID = 3;
   */
  submissionID = protoInt64.zero;

  constructor(data?: PartialMessage<RebuildRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.RebuildRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "courseID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "assignmentID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "submissionID", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RebuildRequest {
    return new RebuildRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RebuildRequest {
    return new RebuildRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RebuildRequest {
    return new RebuildRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RebuildRequest | PlainMessage<RebuildRequest> | undefined, b: RebuildRequest | PlainMessage<RebuildRequest> | undefined): boolean {
    return proto3.util.equals(RebuildRequest, a, b);
  }
}

/**
 * Void contains no fields. A server response with a Void still contains a gRPC status code,
 * which can be checked for success or failure. Status code 0 indicates that the requested action was successful,
 * whereas any other status code indicates some failure. As such, the status code can be used as a boolean result from
 * the server.
 *
 * @generated from message qf.Void
 */
export class Void extends Message<Void> {
  constructor(data?: PartialMessage<Void>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "qf.Void";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Void {
    return new Void().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Void {
    return new Void().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Void {
    return new Void().fromJsonString(jsonString, options);
  }

  static equals(a: Void | PlainMessage<Void> | undefined, b: Void | PlainMessage<Void> | undefined): boolean {
    return proto3.util.equals(Void, a, b);
  }
}

