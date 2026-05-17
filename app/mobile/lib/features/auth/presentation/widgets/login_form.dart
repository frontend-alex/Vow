import 'package:flutter/material.dart';

import '../../../../../core/utils/validators.dart';
import '../../../../../shared/widgets/app_text_field.dart';
import '../../../../../shared/widgets/primary_button.dart';

class LoginForm extends StatelessWidget {
  const LoginForm({super.key});

  @override
  Widget build(BuildContext context) {
    return Form(
      child: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: const <Widget>[
          AppTextField(label: 'Email', validator: Validators.email),
          SizedBox(height: 12),
          AppTextField(
            label: 'Password',
            obscureText: true,
            validator: Validators.password,
          ),
          SizedBox(height: 20),
          PrimaryButton(label: 'Log in', onPressed: null),
        ],
      ),
    );
  }
}
