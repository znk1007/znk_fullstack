///
//  Generated code. Do not modify.
//  source: register.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class RegistRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegistRequest', package: const $pb.PackageName('protos.register'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'password')
    ..aOS(3, 'device')
    ..hasRequiredFields = false
  ;

  RegistRequest._() : super();
  factory RegistRequest() => create();
  factory RegistRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RegistRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RegistRequest clone() => RegistRequest()..mergeFromMessage(this);
  RegistRequest copyWith(void Function(RegistRequest) updates) => super.copyWith((message) => updates(message as RegistRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RegistRequest create() => RegistRequest._();
  RegistRequest createEmptyInstance() => create();
  static $pb.PbList<RegistRequest> createRepeated() => $pb.PbList<RegistRequest>();
  @$core.pragma('dart2js:noInline')
  static RegistRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RegistRequest>(create);
  static RegistRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get password => $_getSZ(1);
  @$pb.TagNumber(2)
  set password($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPassword() => $_has(1);
  @$pb.TagNumber(2)
  void clearPassword() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get device => $_getSZ(2);
  @$pb.TagNumber(3)
  set device($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDevice() => $_has(2);
  @$pb.TagNumber(3)
  void clearDevice() => clearField(3);
}

class RegistResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegistResponse', package: const $pb.PackageName('protos.register'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'userId', protoName: 'userId')
    ..aOS(3, 'message')
    ..a<$core.int>(4, 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(5, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  RegistResponse._() : super();
  factory RegistResponse() => create();
  factory RegistResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RegistResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RegistResponse clone() => RegistResponse()..mergeFromMessage(this);
  RegistResponse copyWith(void Function(RegistResponse) updates) => super.copyWith((message) => updates(message as RegistResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RegistResponse create() => RegistResponse._();
  RegistResponse createEmptyInstance() => create();
  static $pb.PbList<RegistResponse> createRepeated() => $pb.PbList<RegistResponse>();
  @$core.pragma('dart2js:noInline')
  static RegistResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RegistResponse>(create);
  static RegistResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get userId => $_getSZ(1);
  @$pb.TagNumber(2)
  set userId($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasUserId() => $_has(1);
  @$pb.TagNumber(2)
  void clearUserId() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get message => $_getSZ(2);
  @$pb.TagNumber(3)
  set message($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasMessage() => $_has(2);
  @$pb.TagNumber(3)
  void clearMessage() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get code => $_getIZ(3);
  @$pb.TagNumber(4)
  set code($core.int v) { $_setSignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasCode() => $_has(3);
  @$pb.TagNumber(4)
  void clearCode() => clearField(4);

  @$pb.TagNumber(5)
  $core.int get status => $_getIZ(4);
  @$pb.TagNumber(5)
  set status($core.int v) { $_setSignedInt32(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasStatus() => $_has(4);
  @$pb.TagNumber(5)
  void clearStatus() => clearField(5);
}

