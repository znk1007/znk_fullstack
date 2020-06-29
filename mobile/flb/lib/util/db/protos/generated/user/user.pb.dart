///
//  Generated code. Do not modify.
//  source: user.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class User extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('User', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'userID', protoName: 'userID')
    ..aOS(2, 'account')
    ..aOS(3, 'nickname')
    ..aOS(4, 'phone')
    ..aOS(5, 'email')
    ..aOS(6, 'photo')
    ..aOS(7, 'createdAt', protoName: 'createdAt')
    ..aOS(8, 'updatedAt', protoName: 'updatedAt')
    ..hasRequiredFields = false
  ;

  User._() : super();
  factory User() => create();
  factory User.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory User.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  User clone() => User()..mergeFromMessage(this);
  User copyWith(void Function(User) updates) => super.copyWith((message) => updates(message as User));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static User create() => User._();
  User createEmptyInstance() => create();
  static $pb.PbList<User> createRepeated() => $pb.PbList<User>();
  @$core.pragma('dart2js:noInline')
  static User getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<User>(create);
  static User _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userID => $_getSZ(0);
  @$pb.TagNumber(1)
  set userID($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserID() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get account => $_getSZ(1);
  @$pb.TagNumber(2)
  set account($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasAccount() => $_has(1);
  @$pb.TagNumber(2)
  void clearAccount() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get nickname => $_getSZ(2);
  @$pb.TagNumber(3)
  set nickname($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasNickname() => $_has(2);
  @$pb.TagNumber(3)
  void clearNickname() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get phone => $_getSZ(3);
  @$pb.TagNumber(4)
  set phone($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasPhone() => $_has(3);
  @$pb.TagNumber(4)
  void clearPhone() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get email => $_getSZ(4);
  @$pb.TagNumber(5)
  set email($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasEmail() => $_has(4);
  @$pb.TagNumber(5)
  void clearEmail() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get photo => $_getSZ(5);
  @$pb.TagNumber(6)
  set photo($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasPhoto() => $_has(5);
  @$pb.TagNumber(6)
  void clearPhoto() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get createdAt => $_getSZ(6);
  @$pb.TagNumber(7)
  set createdAt($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasCreatedAt() => $_has(6);
  @$pb.TagNumber(7)
  void clearCreatedAt() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get updatedAt => $_getSZ(7);
  @$pb.TagNumber(8)
  set updatedAt($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasUpdatedAt() => $_has(7);
  @$pb.TagNumber(8)
  void clearUpdatedAt() => clearField(8);
}

class UserRgstReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserRgstReq', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserRgstReq._() : super();
  factory UserRgstReq() => create();
  factory UserRgstReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserRgstReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserRgstReq clone() => UserRgstReq()..mergeFromMessage(this);
  UserRgstReq copyWith(void Function(UserRgstReq) updates) => super.copyWith((message) => updates(message as UserRgstReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserRgstReq create() => UserRgstReq._();
  UserRgstReq createEmptyInstance() => create();
  static $pb.PbList<UserRgstReq> createRepeated() => $pb.PbList<UserRgstReq>();
  @$core.pragma('dart2js:noInline')
  static UserRgstReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserRgstReq>(create);
  static UserRgstReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserRgstRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserRgstRes', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserRgstRes._() : super();
  factory UserRgstRes() => create();
  factory UserRgstRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserRgstRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserRgstRes clone() => UserRgstRes()..mergeFromMessage(this);
  UserRgstRes copyWith(void Function(UserRgstRes) updates) => super.copyWith((message) => updates(message as UserRgstRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserRgstRes create() => UserRgstRes._();
  UserRgstRes createEmptyInstance() => create();
  static $pb.PbList<UserRgstRes> createRepeated() => $pb.PbList<UserRgstRes>();
  @$core.pragma('dart2js:noInline')
  static UserRgstRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserRgstRes>(create);
  static UserRgstRes _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserLgnReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserLgnReq', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserLgnReq._() : super();
  factory UserLgnReq() => create();
  factory UserLgnReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserLgnReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserLgnReq clone() => UserLgnReq()..mergeFromMessage(this);
  UserLgnReq copyWith(void Function(UserLgnReq) updates) => super.copyWith((message) => updates(message as UserLgnReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserLgnReq create() => UserLgnReq._();
  UserLgnReq createEmptyInstance() => create();
  static $pb.PbList<UserLgnReq> createRepeated() => $pb.PbList<UserLgnReq>();
  @$core.pragma('dart2js:noInline')
  static UserLgnReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserLgnReq>(create);
  static UserLgnReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserLgnRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserLgnRes', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserLgnRes._() : super();
  factory UserLgnRes() => create();
  factory UserLgnRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserLgnRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserLgnRes clone() => UserLgnRes()..mergeFromMessage(this);
  UserLgnRes copyWith(void Function(UserLgnRes) updates) => super.copyWith((message) => updates(message as UserLgnRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserLgnRes create() => UserLgnRes._();
  UserLgnRes createEmptyInstance() => create();
  static $pb.PbList<UserLgnRes> createRepeated() => $pb.PbList<UserLgnRes>();
  @$core.pragma('dart2js:noInline')
  static UserLgnRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserLgnRes>(create);
  static UserLgnRes _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserUpdatePswReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserUpdatePswReq', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserUpdatePswReq._() : super();
  factory UserUpdatePswReq() => create();
  factory UserUpdatePswReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserUpdatePswReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserUpdatePswReq clone() => UserUpdatePswReq()..mergeFromMessage(this);
  UserUpdatePswReq copyWith(void Function(UserUpdatePswReq) updates) => super.copyWith((message) => updates(message as UserUpdatePswReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserUpdatePswReq create() => UserUpdatePswReq._();
  UserUpdatePswReq createEmptyInstance() => create();
  static $pb.PbList<UserUpdatePswReq> createRepeated() => $pb.PbList<UserUpdatePswReq>();
  @$core.pragma('dart2js:noInline')
  static UserUpdatePswReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserUpdatePswReq>(create);
  static UserUpdatePswReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserUpdatePswRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserUpdatePswRes', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserUpdatePswRes._() : super();
  factory UserUpdatePswRes() => create();
  factory UserUpdatePswRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserUpdatePswRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserUpdatePswRes clone() => UserUpdatePswRes()..mergeFromMessage(this);
  UserUpdatePswRes copyWith(void Function(UserUpdatePswRes) updates) => super.copyWith((message) => updates(message as UserUpdatePswRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserUpdatePswRes create() => UserUpdatePswRes._();
  UserUpdatePswRes createEmptyInstance() => create();
  static $pb.PbList<UserUpdatePswRes> createRepeated() => $pb.PbList<UserUpdatePswRes>();
  @$core.pragma('dart2js:noInline')
  static UserUpdatePswRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserUpdatePswRes>(create);
  static UserUpdatePswRes _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserLgoReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserLgoReq', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserLgoReq._() : super();
  factory UserLgoReq() => create();
  factory UserLgoReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserLgoReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserLgoReq clone() => UserLgoReq()..mergeFromMessage(this);
  UserLgoReq copyWith(void Function(UserLgoReq) updates) => super.copyWith((message) => updates(message as UserLgoReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserLgoReq create() => UserLgoReq._();
  UserLgoReq createEmptyInstance() => create();
  static $pb.PbList<UserLgoReq> createRepeated() => $pb.PbList<UserLgoReq>();
  @$core.pragma('dart2js:noInline')
  static UserLgoReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserLgoReq>(create);
  static UserLgoReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserLgoRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserLgoRes', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserLgoRes._() : super();
  factory UserLgoRes() => create();
  factory UserLgoRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserLgoRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserLgoRes clone() => UserLgoRes()..mergeFromMessage(this);
  UserLgoRes copyWith(void Function(UserLgoRes) updates) => super.copyWith((message) => updates(message as UserLgoRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserLgoRes create() => UserLgoRes._();
  UserLgoRes createEmptyInstance() => create();
  static $pb.PbList<UserLgoRes> createRepeated() => $pb.PbList<UserLgoRes>();
  @$core.pragma('dart2js:noInline')
  static UserLgoRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserLgoRes>(create);
  static UserLgoRes _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserStatusReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserStatusReq', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserStatusReq._() : super();
  factory UserStatusReq() => create();
  factory UserStatusReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserStatusReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserStatusReq clone() => UserStatusReq()..mergeFromMessage(this);
  UserStatusReq copyWith(void Function(UserStatusReq) updates) => super.copyWith((message) => updates(message as UserStatusReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserStatusReq create() => UserStatusReq._();
  UserStatusReq createEmptyInstance() => create();
  static $pb.PbList<UserStatusReq> createRepeated() => $pb.PbList<UserStatusReq>();
  @$core.pragma('dart2js:noInline')
  static UserStatusReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserStatusReq>(create);
  static UserStatusReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserStatusRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserStatusRes', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserStatusRes._() : super();
  factory UserStatusRes() => create();
  factory UserStatusRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserStatusRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserStatusRes clone() => UserStatusRes()..mergeFromMessage(this);
  UserStatusRes copyWith(void Function(UserStatusRes) updates) => super.copyWith((message) => updates(message as UserStatusRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserStatusRes create() => UserStatusRes._();
  UserStatusRes createEmptyInstance() => create();
  static $pb.PbList<UserStatusRes> createRepeated() => $pb.PbList<UserStatusRes>();
  @$core.pragma('dart2js:noInline')
  static UserStatusRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserStatusRes>(create);
  static UserStatusRes _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserUpdatePhoneReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserUpdatePhoneReq', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserUpdatePhoneReq._() : super();
  factory UserUpdatePhoneReq() => create();
  factory UserUpdatePhoneReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserUpdatePhoneReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserUpdatePhoneReq clone() => UserUpdatePhoneReq()..mergeFromMessage(this);
  UserUpdatePhoneReq copyWith(void Function(UserUpdatePhoneReq) updates) => super.copyWith((message) => updates(message as UserUpdatePhoneReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserUpdatePhoneReq create() => UserUpdatePhoneReq._();
  UserUpdatePhoneReq createEmptyInstance() => create();
  static $pb.PbList<UserUpdatePhoneReq> createRepeated() => $pb.PbList<UserUpdatePhoneReq>();
  @$core.pragma('dart2js:noInline')
  static UserUpdatePhoneReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserUpdatePhoneReq>(create);
  static UserUpdatePhoneReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserUpdatePhoneRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserUpdatePhoneRes', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserUpdatePhoneRes._() : super();
  factory UserUpdatePhoneRes() => create();
  factory UserUpdatePhoneRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserUpdatePhoneRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserUpdatePhoneRes clone() => UserUpdatePhoneRes()..mergeFromMessage(this);
  UserUpdatePhoneRes copyWith(void Function(UserUpdatePhoneRes) updates) => super.copyWith((message) => updates(message as UserUpdatePhoneRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserUpdatePhoneRes create() => UserUpdatePhoneRes._();
  UserUpdatePhoneRes createEmptyInstance() => create();
  static $pb.PbList<UserUpdatePhoneRes> createRepeated() => $pb.PbList<UserUpdatePhoneRes>();
  @$core.pragma('dart2js:noInline')
  static UserUpdatePhoneRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserUpdatePhoneRes>(create);
  static UserUpdatePhoneRes _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserUpdateNicknameReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserUpdateNicknameReq', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserUpdateNicknameReq._() : super();
  factory UserUpdateNicknameReq() => create();
  factory UserUpdateNicknameReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserUpdateNicknameReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserUpdateNicknameReq clone() => UserUpdateNicknameReq()..mergeFromMessage(this);
  UserUpdateNicknameReq copyWith(void Function(UserUpdateNicknameReq) updates) => super.copyWith((message) => updates(message as UserUpdateNicknameReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserUpdateNicknameReq create() => UserUpdateNicknameReq._();
  UserUpdateNicknameReq createEmptyInstance() => create();
  static $pb.PbList<UserUpdateNicknameReq> createRepeated() => $pb.PbList<UserUpdateNicknameReq>();
  @$core.pragma('dart2js:noInline')
  static UserUpdateNicknameReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserUpdateNicknameReq>(create);
  static UserUpdateNicknameReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserUpdateNicknameRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserUpdateNicknameRes', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserUpdateNicknameRes._() : super();
  factory UserUpdateNicknameRes() => create();
  factory UserUpdateNicknameRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserUpdateNicknameRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserUpdateNicknameRes clone() => UserUpdateNicknameRes()..mergeFromMessage(this);
  UserUpdateNicknameRes copyWith(void Function(UserUpdateNicknameRes) updates) => super.copyWith((message) => updates(message as UserUpdateNicknameRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserUpdateNicknameRes create() => UserUpdateNicknameRes._();
  UserUpdateNicknameRes createEmptyInstance() => create();
  static $pb.PbList<UserUpdateNicknameRes> createRepeated() => $pb.PbList<UserUpdateNicknameRes>();
  @$core.pragma('dart2js:noInline')
  static UserUpdateNicknameRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserUpdateNicknameRes>(create);
  static UserUpdateNicknameRes _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserUpdateActiveReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserUpdateActiveReq', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserUpdateActiveReq._() : super();
  factory UserUpdateActiveReq() => create();
  factory UserUpdateActiveReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserUpdateActiveReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserUpdateActiveReq clone() => UserUpdateActiveReq()..mergeFromMessage(this);
  UserUpdateActiveReq copyWith(void Function(UserUpdateActiveReq) updates) => super.copyWith((message) => updates(message as UserUpdateActiveReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserUpdateActiveReq create() => UserUpdateActiveReq._();
  UserUpdateActiveReq createEmptyInstance() => create();
  static $pb.PbList<UserUpdateActiveReq> createRepeated() => $pb.PbList<UserUpdateActiveReq>();
  @$core.pragma('dart2js:noInline')
  static UserUpdateActiveReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserUpdateActiveReq>(create);
  static UserUpdateActiveReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class UserUpdateActiveRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UserUpdateActiveRes', package: const $pb.PackageName('user'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  UserUpdateActiveRes._() : super();
  factory UserUpdateActiveRes() => create();
  factory UserUpdateActiveRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserUpdateActiveRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UserUpdateActiveRes clone() => UserUpdateActiveRes()..mergeFromMessage(this);
  UserUpdateActiveRes copyWith(void Function(UserUpdateActiveRes) updates) => super.copyWith((message) => updates(message as UserUpdateActiveRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserUpdateActiveRes create() => UserUpdateActiveRes._();
  UserUpdateActiveRes createEmptyInstance() => create();
  static $pb.PbList<UserUpdateActiveRes> createRepeated() => $pb.PbList<UserUpdateActiveRes>();
  @$core.pragma('dart2js:noInline')
  static UserUpdateActiveRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserUpdateActiveRes>(create);
  static UserUpdateActiveRes _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

