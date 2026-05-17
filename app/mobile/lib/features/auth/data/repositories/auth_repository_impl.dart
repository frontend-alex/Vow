import '../../domain/entities/user.dart';
import '../../domain/repositories/auth_repository.dart';
import '../datasources/auth_remote_datasource.dart';
import '../models/login_request_model.dart';

class AuthRepositoryImpl implements AuthRepository {
  const AuthRepositoryImpl(this.remoteDatasource);

  final AuthRemoteDatasource remoteDatasource;

  @override
  Future<User> login({required String email, required String password}) {
    return remoteDatasource.login(
      LoginRequestModel(email: email, password: password),
    );
  }

  @override
  Future<void> logout() async {}

  @override
  Future<User> register({
    required String name,
    required String email,
    required String password,
  }) async {
    return User(id: 0, email: email, name: name);
  }
}
