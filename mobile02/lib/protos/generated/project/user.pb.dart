///
//  Generated code. Do not modify.
//  source: user.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:protobuf/protobuf.dart' as $pb;

class User extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('User', package: const $pb.PackageName('protos.user'))
    ..aOS(1, 'userId')
    ..aOS(2, 'sessionId')
    ..aOS(3, 'account')
    ..aOS(4, 'nickname')
    ..aOS(5, 'phone')
    ..aOS(6, 'email')
    ..aOS(7, 'photo')
    ..aOS(8, 'createdAt')
    ..aOB(9, 'online')
    ..hasRequiredFields = false
  ;

  User() : super();
  User.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  User.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  User clone() => User()..mergeFromMessage(this);
  User copyWith(void Function(User) updates) => super.copyWith((message) => updates(message as User));
  $pb.BuilderInfo get info_ => _i;
  static User create() => User();
  User createEmptyInstance() => create();
  static $pb.PbList<User> createRepeated() => $pb.PbList<User>();
  static User getDefault() => _defaultInstance ??= create()..freeze();
  static User _defaultInstance;

  $core.String get userId => $_getS(0, '');
  set userId($core.String v) { $_setString(0, v); }
  $core.bool hasUserId() => $_has(0);
  void clearUserId() => clearField(1);

  $core.String get sessionId => $_getS(1, '');
  set sessionId($core.String v) { $_setString(1, v); }
  $core.bool hasSessionId() => $_has(1);
  void clearSessionId() => clearField(2);

  $core.String get account => $_getS(2, '');
  set account($core.String v) { $_setString(2, v); }
  $core.bool hasAccount() => $_has(2);
  void clearAccount() => clearField(3);

  $core.String get nickname => $_getS(3, '');
  set nickname($core.String v) { $_setString(3, v); }
  $core.bool hasNickname() => $_has(3);
  void clearNickname() => clearField(4);

  $core.String get phone => $_getS(4, '');
  set phone($core.String v) { $_setString(4, v); }
  $core.bool hasPhone() => $_has(4);
  void clearPhone() => clearField(5);

  $core.String get email => $_getS(5, '');
  set email($core.String v) { $_setString(5, v); }
  $core.bool hasEmail() => $_has(5);
  void clearEmail() => clearField(6);

  $core.String get photo => $_getS(6, '');
  set photo($core.String v) { $_setString(6, v); }
  $core.bool hasPhoto() => $_has(6);
  void clearPhoto() => clearField(7);

  $core.String get createdAt => $_getS(7, '');
  set createdAt($core.String v) { $_setString(7, v); }
  $core.bool hasCreatedAt() => $_has(7);
  void clearCreatedAt() => clearField(8);

  $core.bool get online => $_get(8, false);
  set online($core.bool v) { $_setBool(8, v); }
  $core.bool hasOnline() => $_has(8);
  void clearOnline() => clearField(9);
}

