import 'package:flutter/material.dart';

import '../../../../../shared/layout/app_scaffold.dart';

class RegisterPage extends StatelessWidget {
  const RegisterPage({super.key});

  @override
  Widget build(BuildContext context) {
    return const AppScaffold(
      title: 'Create Account',
      body: Center(child: Text('Register')),
    );
  }
}
