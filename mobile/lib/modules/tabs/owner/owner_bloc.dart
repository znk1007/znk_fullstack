import 'dart:async';
import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter/widgets.dart';
import 'package:znk/modules/tabs/owner/model.dart';
import 'package:znk/utils/database/user.dart';

abstract class OwnerEvent extends Equatable {
  OwnerEvent([List props = const []]) : super(props);
}

class Generate extends OwnerEvent {
  @override
  String toString() {
    return 'Generate';
  }
}

abstract class OwnerState extends Equatable {
  OwnerState([List props = const []]): super(props);
}

class Loaded extends OwnerState {
  final List<OwnerModel> models;
  Loaded({@required this.models}): super([models]);

  @override
  String toString() {
    return 'Loaded { models: ${models.length} }';
  }
}


class Loading extends OwnerState {
  @override
  String toString() {
    return 'Loading';
  }
}

class OwnerBloc extends Bloc<OwnerEvent, OwnerState> {
  // static final OwnerBloc _ownerBlocSingleton = new OwnerBloc._internal();
  // factory OwnerBloc() {
  //   return _ownerBlocSingleton;
  // }
  // OwnerBloc._internal();
  
  final OwnerModel ownerModel;

  OwnerBloc({@required this.ownerModel});

  @override
  OwnerState get initialState => new Loading();

  @override
  Stream<OwnerState> mapEventToState(
    OwnerEvent event,
  ) async* {
    try {
      // yield await event.applyAsync(currentState: currentState, bloc: this);
      if (event is Generate) {
        final models = await ownerModel.generate();
        yield Loaded(models: models);
      }
    } catch (_, stackTrace) {
      print('$_ $stackTrace');
      yield currentState;
    }
  }
}

