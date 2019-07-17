///
//  Generated code. Do not modify.
//  source: updateonline.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:protobuf/protobuf.dart' as $pb;

class UpdateOnlineRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UpdateOnlineRequest', package: const $pb.PackageName('protos.updateonline'))
    ..aOS(1, 'account')
    ..aOS(2, 'userId')
    ..aOS(3, 'sessionId')
    ..aOB(4, 'online')
    ..aOS(5, 'device')
    ..hasRequiredFields = false
  ;

  UpdateOnlineRequest() : super();
  UpdateOnlineRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  UpdateOnlineRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  UpdateOnlineRequest clone() => UpdateOnlineRequest()..mergeFromMessage(this);
  UpdateOnlineRequest copyWith(void Function(UpdateOnlineRequest) updates) => super.copyWith((message) => updates(message as UpdateOnlineRequest));
  $pb.BuilderInfo get info_ => _i;
  static UpdateOnlineRequest create() => UpdateOnlineRequest();
  UpdateOnlineRequest createEmptyInstance() => create();
  static $pb.PbList<UpdateOnlineRequest> createRepeated() => $pb.PbList<UpdateOnlineRequest>();
  static UpdateOnlineRequest getDefault() => _defaultInstance ??= create()..freeze();
  static UpdateOnlineRequest _defaultInstance;

  $core.String get account => $_getS(0, '');
  set account($core.String v) { $_setString(0, v); }
  $core.bool hasAccount() => $_has(0);
  void clearAccount() => clearField(1);

  $core.String get userId => $_getS(1, '');
  set userId($core.String v) { $_setString(1, v); }
  $core.bool hasUserId() => $_has(1);
  void clearUserId() => clearField(2);

  $core.String get sessionId => $_getS(2, '');
  set sessionId($core.String v) { $_setString(2, v); }
  $core.bool hasSessionId() => $_has(2);
  void clearSessionId() => clearField(3);

  $core.bool get online => $_get(3, false);
  set online($core.bool v) { $_setBool(3, v); }
  $core.bool hasOnline() => $_has(3);
  void clearOnline() => clearField(4);

  $core.String get device => $_getS(4, '');
  set device($core.String v) { $_setString(4, v); }
  $core.bool hasDevice() => $_has(4);
  void clearDevice() => clearField(5);
}

class UpdateOnlineResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UpdateOnlineResponse', package: const $pb.PackageName('protos.updateonline'))
    ..aOS(1, 'message')
    ..a<$core.int>(2, 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(3, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  UpdateOnlineResponse() : super();
  UpdateOnlineResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  UpdateOnlineResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  UpdateOnlineResponse clone() => UpdateOnlineResponse()..mergeFromMessage(this);
  UpdateOnlineResponse copyWith(void Function(UpdateOnlineResponse) updates) => super.copyWith((message) => updates(message as UpdateOnlineResponse));
  $pb.BuilderInfo get info_ => _i;
  static UpdateOnlineResponse create() => UpdateOnlineResponse();
  UpdateOnlineResponse createEmptyInstance() => create();
  static $pb.PbList<UpdateOnlineResponse> createRepeated() => $pb.PbList<UpdateOnlineResponse>();
  static UpdateOnlineResponse getDefault() => _defaultInstance ??= create()..freeze();
  static UpdateOnlineResponse _defaultInstance;

  $core.String get message => $_getS(0, '');
  set message($core.String v) { $_setString(0, v); }
  $core.bool hasMessage() => $_has(0);
  void clearMessage() => clearField(1);

  $core.int get code => $_get(1, 0);
  set code($core.int v) { $_setSignedInt32(1, v); }
  $core.bool hasCode() => $_has(1);
  void clearCode() => clearField(2);

  $core.int get status => $_get(2, 0);
  set status($core.int v) { $_setSignedInt32(2, v); }
  $core.bool hasStatus() => $_has(2);
  void clearStatus() => clearField(3);
}

