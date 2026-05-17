class ApiClient {
  const ApiClient({required this.baseUrl});

  final String baseUrl;

  Future<Map<String, Object?>> get(String path) async {
    return <String, Object?>{};
  }

  Future<Map<String, Object?>> post(
    String path, {
    Map<String, Object?> body = const <String, Object?>{},
  }) async {
    return <String, Object?>{};
  }
}
