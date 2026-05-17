import '../../../../../core/network/api_client.dart';
import '../models/login_request_model.dart';
import '../models/user_model.dart';

class AuthRemoteDatasource {
  const AuthRemoteDatasource(this.client);

  final ApiClient client;

  Future<UserModel> login(LoginRequestModel request) async {
    final json = await client.post('/v1/auth/login', body: request.toJson());
    return UserModel.fromJson(json);
  }
}
