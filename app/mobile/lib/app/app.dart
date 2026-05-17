import 'package:flutter/material.dart';

import '../features/auth/presentation/pages/login_page.dart';
import 'theme.dart';

class VowApp extends StatelessWidget {
  const VowApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Vow',
      theme: AppTheme.light,
      home: const LoginPage(),
    );
  }
}
