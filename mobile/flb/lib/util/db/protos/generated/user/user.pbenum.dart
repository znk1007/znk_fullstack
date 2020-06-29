///
//  Generated code. Do not modify.
//  source: user.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Permission extends $pb.ProtobufEnum {
  static const Permission super_ = Permission._(0, 'super');
  static const Permission admin = Permission._(1, 'admin');
  static const Permission user = Permission._(2, 'user');
  static const Permission visitor = Permission._(3, 'visitor');

  static const $core.List<Permission> values = <Permission> [
    super_,
    admin,
    user,
    visitor,
  ];

  static final $core.Map<$core.int, Permission> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Permission valueOf($core.int value) => _byValue[value];

  const Permission._($core.int v, $core.String n) : super(v, n);
}

