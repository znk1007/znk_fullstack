///
//  Generated code. Do not modify.
//  source: register.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:protobuf/protobuf.dart' as $pb;

class RegistRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegistRequest', package: const $pb.PackageName('protos.register'))
    ..aOS(1, 'account')
    ..aOS(2, 'password')
    ..aOS(3, 'device')
    ..hasRequiredFields = false
  ;

  RegistRequest() : super();
  RegistRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  RegistRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  RegistRequest clone() => RegistRequest()..mergeFromMessage(this);
  RegistRequest copyWith(void Function(RegistRequest) updates) => super.copyWith((message) => updates(message as RegistRequest));
  $pb.BuilderInfo get info_ => _i;
  static RegistRequest create() => RegistRequest();
  RegistRequest createEmptyInstance() => create();
  static $pb.PbList<RegistRequest> createRepeated() => $pb.PbList<RegistRequest>();
  static RegistRequest getDefault() => _defaultInstance ??= create()..freeze();
  static RegistRequest _defaultInstance;

  $core.String get account => $_getS(0, '');
  set account($core.String v) { $_setString(0, v); }
  $core.bool hasAccount() => $_has(0);
  void clearAccount() => clearField(1);

  $core.String get password => $_getS(1, '');
  set password($core.String v) { $_setString(1, v); }
  $core.bool hasPassword() => $_has(1);
  void clearPassword() => clearField(2);

  $core.String get device => $_getS(2, '');
  set device($core.String v) { $_setString(2, v); }
  $core.bool hasDevice() => $_has(2);
  void clearDevice() => clearField(3);
}

class RegistResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegistResponse', package: const $pb.PackageName('protos.register'))
    ..aOS(1, 'account')
    ..aOS(2, 'userId')
    ..aOS(3, 'message')
    ..a<$core.int>(4, 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(5, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  RegistResponse() : super();
  RegistResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  RegistResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  RegistResponse clone() => RegistResponse()..mergeFromMessage(this);
  RegistResponse copyWith(void Function(RegistResponse) updates) => super.copyWith((message) => updates(message as RegistResponse));
  $pb.BuilderInfo get info_ => _i;
  static RegistResponse create() => RegistResponse();
  RegistResponse createEmptyInstance() => create();
  static $pb.PbList<RegistResponse> createRepeated() => $pb.PbList<RegistResponse>();
  static RegistResponse getDefault() => _defaultInstance ??= create()..freeze();
  static RegistResponse _defaultInstance;

  $core.String get account => $_getS(0, '');
  set account($core.String v) { $_setString(0, v); }
  $core.bool hasAccount() => $_has(0);
  void clearAccount() => clearField(1);

  $core.String get userId => $_getS(1, '');
  set userId($core.String v) { $_setString(1, v); }
  $core.bool hasUserId() => $_has(1);
  void clearUserId() => clearField(2);

  $core.String get message => $_getS(2, '');
  set message($core.String v) { $_setString(2, v); }
  $core.bool hasMessage() => $_has(2);
  void clearMessage() => clearField(3);

  $core.int get code => $_get(3, 0);
  set code($core.int v) { $_setSignedInt32(3, v); }
  $core.bool hasCode() => $_has(3);
  void clearCode() => clearField(4);

  $core.int get status => $_get(4, 0);
  set status($core.int v) { $_setSignedInt32(4, v); }
  $core.bool hasStatus() => $_has(4);
  void clearStatus() => clearField(5);
}

