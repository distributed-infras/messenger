import 'dart:convert';

import 'package:HyperChat/entity/connection_hub_config.dart';
import 'package:http/http.dart' as http;

class ChatService {
  String _baseURL;

  ChatService(String baseURL) {
    this._baseURL = baseURL;
  }

  Future<ConnectionHubConfig> allocateConnectionHub() async {
    final response = await http.get('$_baseURL/allocate-connection-hub');
    if (response.statusCode == 200) {
      return ConnectionHubConfig.fromJson(json.decode(response.body));
    }
    throw Exception("Failed to allocate connection hub");
  }
}