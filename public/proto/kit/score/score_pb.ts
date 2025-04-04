// @generated by protoc-gen-es v2.2.5 with parameter "target=ts"
// @generated from file kit/score/score.proto (package score, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file kit/score/score.proto.
 */
export const file_kit_score_score: GenFile = /*@__PURE__*/
  fileDesc("ChVraXQvc2NvcmUvc2NvcmUucHJvdG8SBXNjb3JlItEBCgVTY29yZRIKCgJJRBgBIAEoBBIxCgxTdWJtaXNzaW9uSUQYAiABKARCG8q1AxeiARRnb3JtOiJmb3JlaWduS2V5OklEIhIfCgZTZWNyZXQYAyABKAlCD8q1AwuiAQhnb3JtOiItIhIQCghUZXN0TmFtZRgEIAEoCRIQCghUYXNrTmFtZRgFIAEoCRINCgVTY29yZRgGIAEoBRIQCghNYXhTY29yZRgHIAEoBRIOCgZXZWlnaHQYCCABKAUSEwoLVGVzdERldGFpbHMYCSABKAkitQIKCUJ1aWxkSW5mbxIKCgJJRBgBIAEoBBIxCgxTdWJtaXNzaW9uSUQYAiABKARCG8q1AxeiARRnb3JtOiJmb3JlaWduS2V5OklEIhIQCghCdWlsZExvZxgDIAEoCRIQCghFeGVjVGltZRgEIAEoAxJfCglCdWlsZERhdGUYBSABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wQjDKtQMsogEpZ29ybToic2VyaWFsaXplcjp0aW1lc3RhbXA7dHlwZTpkYXRldGltZSISZAoOU3VibWlzc2lvbkRhdGUYBiABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wQjDKtQMsogEpZ29ybToic2VyaWFsaXplcjp0aW1lc3RhbXA7dHlwZTpkYXRldGltZSJCKlooZ2l0aHViLmNvbS9xdWlja2ZlZWQvcXVpY2tmZWVkL2tpdC9zY29yZWIGcHJvdG8z", [file_google_protobuf_timestamp]);

/**
 * Score give the score for a single test named TestName.
 *
 * @generated from message score.Score
 */
export type Score = Message<"score.Score"> & {
  /**
   * @generated from field: uint64 ID = 1;
   */
  ID: bigint;

  /**
   * @generated from field: uint64 SubmissionID = 2;
   */
  SubmissionID: bigint;

  /**
   * the unique identifier for a scoring session
   *
   * @generated from field: string Secret = 3;
   */
  Secret: string;

  /**
   * name of the test
   *
   * @generated from field: string TestName = 4;
   */
  TestName: string;

  /**
   * name of task this score belongs to
   *
   * @generated from field: string TaskName = 5;
   */
  TaskName: string;

  /**
   * the score obtained
   *
   * @generated from field: int32 Score = 6;
   */
  Score: number;

  /**
   * max score possible to get on this specific test
   *
   * @generated from field: int32 MaxScore = 7;
   */
  MaxScore: number;

  /**
   * the weight of this test; used to compute final grade
   *
   * @generated from field: int32 Weight = 8;
   */
  Weight: number;

  /**
   * if populated, the frontend may display these details
   *
   * @generated from field: string TestDetails = 9;
   */
  TestDetails: string;
};

/**
 * Describes the message score.Score.
 * Use `create(ScoreSchema)` to create a new message.
 */
export const ScoreSchema: GenMessage<Score> = /*@__PURE__*/
  messageDesc(file_kit_score_score, 0);

/**
 * BuildInfo holds build data for an assignment's test execution.
 *
 * @generated from message score.BuildInfo
 */
export type BuildInfo = Message<"score.BuildInfo"> & {
  /**
   * @generated from field: uint64 ID = 1;
   */
  ID: bigint;

  /**
   * @generated from field: uint64 SubmissionID = 2;
   */
  SubmissionID: bigint;

  /**
   * @generated from field: string BuildLog = 3;
   */
  BuildLog: string;

  /**
   * @generated from field: int64 ExecTime = 4;
   */
  ExecTime: bigint;

  /**
   * @generated from field: google.protobuf.Timestamp BuildDate = 5;
   */
  BuildDate?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp SubmissionDate = 6;
   */
  SubmissionDate?: Timestamp;
};

/**
 * Describes the message score.BuildInfo.
 * Use `create(BuildInfoSchema)` to create a new message.
 */
export const BuildInfoSchema: GenMessage<BuildInfo> = /*@__PURE__*/
  messageDesc(file_kit_score_score, 1);

