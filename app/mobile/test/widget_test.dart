import 'package:flutter_test/flutter_test.dart';
import 'package:mobile/app/app.dart';

void main() {
  testWidgets('App renders login page', (WidgetTester tester) async {
    await tester.pumpWidget(const VowApp());

    expect(find.text('Welcome to Vow'), findsOneWidget);
  });
}
