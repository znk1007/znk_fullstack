///
//  Generated code. Do not modify.
//  source: checkuserId.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:protobuf/protobuf.dart' as $pb;

class CheckUserIdRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('CheckUserIdRequest', package: const $pb.PackageName('protos.checkuserId'))
    ..aOS(1, 'account')
    ..aOS(2, 'device')
    ..hasRequiredFields = false
  ;

  CheckUserIdRequest() : super();
  CheckUserIdRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  CheckUserIdRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  CheckUserIdRequest clone() => CheckUserIdRequest()..mergeFromMessage(this);
  CheckUserIdRequest copyWith(void Function(CheckUserIdRequest) updates) => super.copyWith((message) => updates(message as CheckUserIdRequest));
  $pb.BuilderInfo get info_ => _i;
  static CheckUserIdRequest create() => CheckUserIdRequest();
  CheckUserIdRequest createEmptyInstance() => create();
  static $pb.PbList<CheckUserIdRequest> createRepeated() => $pb.PbList<CheckUserIdRequest>();
  static CheckUserIdRequest getDefault() => _defaultInstance ??= create()..freeze();
  static CheckUserIdRequest _defaultInstance;

  $core.String get account => $_getS(0, '');
  set account($core.String v) { $_setString(0, v); }
  $core.bool hasAccount() => $_has(0);
  void clearAccount() => clearField(1);

  $core.String get device => $_getS(1, '');
  set device($core.String v) { $_setString(1, v); }
  $core.bool hasDevice() => $_has(1);
  void clearDevice() => clearField(2);
}

class CheckUserIdResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('CheckUserIdResponse', package: const $pb.PackageName('protos.checkuserId'))
    ..aOS(1, 'message')
    ..a<$core.int>(2, 'code', $pb.PbFieldType.O3)
    ..aOS(3, 'userId')
    ..a<$core.int>(4, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  CheckUserIdResponse() : super();
  CheckUserIdResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  CheckUserIdResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  CheckUserIdResponse clone() => CheckUserIdResponse()..mergeFromMessage(this);
  CheckUserIdResponse copyWith(void Function(CheckUserIdResponse) updates) => super.copyWith((message) => updates(message as CheckUserIdResponse));
  $pb.BuilderInfo get info_ => _i;
  static CheckUserIdResponse create() => CheckUserIdResponse();
  CheckUserIdResponse createEmptyInstance() => create();
  static $pb.PbList<CheckUserIdResponse> createRepeated() => $pb.PbList<CheckUserIdResponse>();
  static CheckUserIdResponse getDefault() => _defaultInstance ??= create()..freeze();
  static CheckUserIdResponse _defaultInstance;

  $core.String get message => $_getS(0, '');
  set message($core.String v) { $_setString(0, v); }
  $core.bool hasMessage() => $_has(0);
  void clearMessage() => clearField(1);

  $core.int get code => $_get(1, 0);
  set code($core.int v) { $_setSignedInt32(1, v); }
  $core.bool hasCode() => $_has(1);
  void clearCode() => clearField(2);

  $core.String get userId => $_getS(2, '');
  set userId($core.String v) { $_setString(2, v); }
  $core.bool hasUserId() => $_has(2);
  void clearUserId() => clearField(3);

  $core.int get status => $_get(3, 0);
  set status($core.int v) { $_setSignedInt32(3, v); }
  $core.bool hasStatus() => $_has(3);
  void clearStatus() => clearField(4);
}

