class AuthInterceptor {
  const AuthInterceptor({this.token});

  final String? token;

  Map<String, String> headers() {
    if (token == null || token!.isEmpty) {
      return const <String, String>{};
    }
    return <String, String>{'Authorization': 'Bearer $token'};
  }
}
