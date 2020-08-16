import 'package:json_annotation/json_annotation.dart';

part 'connection_hub_config.g.dart';

enum Protocol {
  @JsonValue(0)
  WebSocket,
  @JsonValue(1)
  MQTT,
  @JsonValue(2)
  XMPP,
  @JsonValue(3)
  LongPoll
}

@JsonSerializable()
class ConnectionHubConfig {
  final String host;
  final int port;
  final Protocol protocol;

  ConnectionHubConfig({this.host, this.port, this.protocol});

  factory ConnectionHubConfig.fromJson(Map<String, dynamic> json) => _$ConnectionHubConfigFromJson(json);
  Map<String, dynamic> toJson() => _$ConnectionHubConfigToJson(this);
}