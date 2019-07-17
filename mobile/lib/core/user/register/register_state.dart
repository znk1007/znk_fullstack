import 'package:meta/meta.dart';

@immutable 
class RegisterState {
  final bool isAccountValid;
  final bool isPasswordValid;
  final bool isSubmitting;
  final bool isSuccess;
  final bool isFailed;
  bool get isFormValid => isAccountValid && isPasswordValid;
  RegisterState({
    @required this.isAccountValid,
    @required this.isPasswordValid,
    @required this.isSubmitting,
    @required this.isSuccess,
    @required this.isFailed,
  });
  // 空状态
  factory RegisterState.empty() {
    return RegisterState(
      isAccountValid: true,
      isPasswordValid: true,
      isSubmitting: false,
      isSuccess: false,
      isFailed: false,
    );
  }
  // 正在请求
  factory RegisterState.loading() {
    return RegisterState(
      isAccountValid: true,
      isPasswordValid: true,
      isSubmitting: true,
      isSuccess: false,
      isFailed: false,
    );
  }
  // 失败状态
  factory RegisterState.failed() {
    return RegisterState(
      isAccountValid: true,
      isPasswordValid: true,
      isSubmitting: false,
      isSuccess: false,
      isFailed: true,
    );
  }
  // 成功状态
  factory RegisterState.success() {
    return RegisterState(
      isAccountValid: true,
      isPasswordValid: true,
      isSubmitting: false,
      isSuccess: true,
      isFailed: false,
    );
  }
  // 更新状态
  RegisterState update({
    bool isAccountValid,
    bool isPasswordValid,
  }) {
    return copy(
      isAccountValid: isAccountValid,
      isPasswordValid: isPasswordValid,
      isSubmitting: false,
      isSuccess: false,
      isFailed: false,
    );
  }
  // 拷贝
  RegisterState copy({
    bool isAccountValid,
    bool isPasswordValid,
    bool isSubmitting,
    bool isSuccess,
    bool isFailed,
  }) {
    return RegisterState(
      isAccountValid: isAccountValid ?? this.isAccountValid,
      isPasswordValid: isPasswordValid ?? this.isPasswordValid,
      isSubmitting: isSubmitting ?? this.isSubmitting,
      isSuccess: isSuccess ?? this.isSuccess,
      isFailed: isFailed ?? this.isFailed,
    );
  }
}

