import 'package:equatable/equatable.dart';
import 'package:meta/meta.dart';

@immutable
abstract class AuthState extends Equatable {
  AuthState([List props = const []]): super(props);
}

class Uninitialized extends AuthState {
  @override
  String toString() {
    return 'Uninitialized';
  }
}

class Authenticated extends AuthState {
  final String displayName;
  Authenticated(this.displayName): super([displayName]);
  @override
  String toString() {
    return 'Authenticated {displayName: $displayName}';
  }
}

class UnAuthenticated extends AuthState {
  @override
  String toString() {
    return 'UnAuthenticated';
  }

}