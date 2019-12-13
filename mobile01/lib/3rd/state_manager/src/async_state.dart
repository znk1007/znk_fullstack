import 'dart:async';
import 'package:flutter/widgets.dart';
import 'base_state.dart';
import 'delegate_widget.dart';
/// A callback used to build a valid value from an error.  
/// 
/// See also:
/// 
/// * [StreamProvider.catchError] which uses [ErrorBuilder] to handle errors emitted
/// by a [Stream].  
/// * [FutureProvider.catchError] which uses [ErrorBuilder] to handle [Future.catch].  
typedef ErrorBuilder<T> = T Function(BuildContext context, Object error);

/// Listens to a [Stream<T>] and exposes [T] to its descendants. 
/// 
/// Its main use-case is to provide to a large number of a widget the content of
/// a [Stream], without caring ahout reacting to events. 
/// 
/// A typical example whould be to expose the battery level, or a Firebase query. 
/// Trying to use [Stream] to replace [ChangeNotifier] is outside of the scope of
/// this class. 
/// 
/// It is considered an error to pass a stream that can emit errors without providing
/// a [catchError] method. 
/// 
/// {@template provider.streamprovider.initialdata}
/// [initialData] determines the value exposed until the [Stream] emits a value. 
/// If omitted, detaults to `null`. 
/// {@endtemplate}
/// 
/// {@macro provider.updateshouldnotify}
/// 
/// See also:
/// 
/// * [Stream], which is listened by [StreamProvider]. 
/// * [StreamController], to create a [Stream]
class StreamProvider<T> extends ValueDelegateWidget<Stream<T>> 
  implements SingleChildCloneableWidget {
  /// Creates a [Stream] from [create] and subscribes to it.
  /// 
  /// The parameter [create] must not be `null`.
  StreamProvider({
    Key key, 
    @required ValueBuilder<Stream<T>> create, 
    T initialData, 
    ErrorBuilder<T> catchError, 
    UpdateShouldNotify<T> updateShouldNotify, 
    Widget child,
  }) : this._(
    key: key,
    delegate: BuilderStateDelegate<Stream<T>>(create), 
    catchError: catchError, 
    updateShouldNotify: updateShouldNotify, 
    child: child, 
  );

  /// Creates a [StreamController] from [create] and subscribes to its stream.
  /// 
  /// [StreamProvider] will automatically call [StreamController.close] when the
  /// widget is removed from the tree. 
  /// 
  /// The parameter [create] must not be `null`.
  StreamProvider.controller({
    Key key, 
    @required ValueBuilder<StreamController<T>> create, 
    T initialData, 
    ErrorBuilder<T> catchError, 
    UpdateShouldNotify<T> updateShouldNotify, 
    Widget child,
  }) : this._(
    key: key,
    delegate: _StreamControllerBuilderDelegate(create), 
    catchError: catchError, 
    updateShouldNotify: updateShouldNotify, 
    child: child, 
  );

  /// Listens to [value] and expose it to all of [StreamProvider] descendants
  StreamProvider.value({
    Key key, 
    @required Stream<T> value, 
    T initialData, 
    ErrorBuilder<T> catchError, 
    UpdateShouldNotify<T> updateShouldNotify, 
    Widget child,
  }) : this._(
    key: key,
    delegate: SingleValueDelegate(value), 
    catchError: catchError, 
    updateShouldNotify: updateShouldNotify, 
    child: child, 
  );

  StreamProvider._({
    Key key, 
    @required ValueStateDelegate<Stream<T>> delegate, 
    this.initialData, 
    this.catchError,
    this.updateShouldNotify,
    this.child, 
  }):super(
    key: key, 
    delegate: delegate
  );

  /// {@macro provider.streamprovider.initialdata}
  final T initialData;

  /// The widget that is below the current [StreamProvider] widget in the tree.
  /// 
  /// {@macro flutter.wigets.child}
  final Widget child;

  /// An optional function used whenever the [Stream] emits an error.
  /// 
  /// [catchError] will be called with the emitted error and is expected to
  /// return a fallback value without throwing.
  final ErrorBuilder<T> catchError;

  /// @{marcro provider.updatesshouldnotify}
  final UpdateShouldNotify<T> updateShouldNotify;

  @override
  SingleChildCloneableWidget cloneWithChild(Widget child) {
    return StreamProvider._(
      key: key,
      delegate: delegate,
      updateShouldNotify: updateShouldNotify,
      initialData: initialData,
      catchError: catchError,
      child: child,
    );
  }

  @override
  Widget build(BuildContext context) {
    return StreamBuilder<T>(
      stream: delegate.value,
      initialData: initialData,
      builder: (_, snapshot) {
        return InheritedProvider<T>(
          value: _snapshotToValue(snapshot, context, catchError, null),
          child: child,
          updateShouldNotify: updateShouldNotify,
        );
      },
    );
  }

  
}

T _snapshotToValue<T>(
  AsyncSnapshot<T> snapshot, 
  BuildContext context, 
  ErrorBuilder<T> catchError, 
  ValueDelegateWidget owner
) {
  if (snapshot.hasError) {
    if (catchError != null) {
      return catchError(context, snapshot.error);
    }
    throw FlutterError(
      '''
      An exception was throw by ${
              // ignore: invalid_use_of_protected_member
              owner.delegate.value?.runtimeType} listened by
      $owner, but no `catchError` was provided.

      Exception:
      ${snapshot.error}
      '''
    );
  }
  return snapshot.data;
}

class _StreamControllerBuilderDelegate<T> extends 
  ValueStateDelegate<Stream<T>> {
  _StreamControllerBuilderDelegate(this._create) : assert(_create != null);

  StreamController<T> _controller;
  final ValueBuilder<StreamController<T>> _create;

  @override
  Stream<T> value;

  @override
  void initDelegate() {
    super.initDelegate();
    _controller = _create(context);
    value = _controller?.stream;
  }

  @override
  void didUpdateDelegate(_StreamControllerBuilderDelegate<T> old) {
    super.didUpdateDelegate(old);
    value = old.value;
    _controller = old._controller;
  }

  @override
  void dispose() {
    _controller?.close();
    super.dispose();
  }
}

/// Listens to a [Future<T>] and expose [T] to its descendants. 
/// 
/// It is considered an error to pass a future that can emit errors without
/// providing a [catchError] method. 
/// 
/// {@macro provider.updateshouldnotify}
/// 
/// See also:
/// 
/// * [Future], which is listened by [FutureProvider]. 
class FutureProvider<T> extends ValueDelegateWidget<Future<T>> 
  implements SingleChildCloneableWidget {  

  /// Creates a [Future] from [create] and subscribes to it.
  /// 
  /// [create] must not be `null`
  FutureProvider({
    Key key, 
    @required ValueBuilder<Future<T>> create,
    T initalData, 
    ErrorBuilder<T> catchError, 
    UpdateShouldNotify<T> updateShouldNotify, 
    Widget child,
  }): this._(
    key: key,
    initialData: initalData,
    catchError: catchError,
    updateShouldNotify: updateShouldNotify, 
    delegate: BuilderStateDelegate(create),
    child: child,
  );

  /// Listens to [value] and expose it to all of [FutureProvider] descendants. 
  FutureProvider.value({
    Key key,
    @required Future<T> value, 
    T initialData, 
    ErrorBuilder<T> catchError, 
    UpdateShouldNotify<T> updateShouldNotify, 
    Widget child, 
  }): this._(
    key: key,
    initialData: initialData,
    catchError: catchError,
    updateShouldNotify: updateShouldNotify,
    delegate: SingleValueDelegate(value),
    child: child,
  );

  FutureProvider._({
    Key key, 
    @required ValueStateDelegate<Future<T>> delegate, 
    this.initialData,
    this.catchError,
    this.updateShouldNotify,
    this.child,
  }):super(
    key: key, 
    delegate: delegate, 
  );
  
  /// [initialData] determines the value exposed until the [Future] completes.
  /// 
  /// If omitted, defaults to `null`
  final T initialData;

  /// The widget that is below the current [FutureProvider] widget in the tree.
  /// 
  /// {@macro flutter.widgets.child}
  final Widget child;

  /// Optional function used if the [Future] emits an error.
  /// 
  /// [catchError] will be called with the emitted error and is expected to
  /// return a fallback value without throwing.
  /// 
  /// The returned value will then be exposed to the descendants of
  /// [FutureProvider] like any valid value.
  final ErrorBuilder<T> catchError;

  /// {@macro provider.udpateshouldnotify}
  final UpdateShouldNotify<T> updateShouldNotify;

  @override
  SingleChildCloneableWidget cloneWithChild(Widget child) {
    return FutureProvider._(
      key: key,
      delegate: delegate,
      updateShouldNotify: updateShouldNotify,
      initialData: initialData,
      catchError: catchError,
      child: child,
    );
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<T> (
      future: delegate.value,
      initialData: initialData,
      builder: (_, snapshot) {
        return InheritedProvider<T>(
          value: _snapshotToValue(snapshot, context, catchError, this),
          updateShouldNotify: updateShouldNotify,
          child: child,
        );
      },
    );
  }
}
