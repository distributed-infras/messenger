// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'connection_hub_config.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ConnectionHubConfig _$ConnectionHubConfigFromJson(Map<String, dynamic> json) {
  return ConnectionHubConfig(
    host: json['host'] as String,
    port: json['port'] as int,
    protocol: _$enumDecodeNullable(_$ProtocolEnumMap, json['protocol']),
  );
}

Map<String, dynamic> _$ConnectionHubConfigToJson(
        ConnectionHubConfig instance) =>
    <String, dynamic>{
      'host': instance.host,
      'port': instance.port,
      'protocol': _$ProtocolEnumMap[instance.protocol],
    };

T _$enumDecode<T>(
  Map<T, dynamic> enumValues,
  dynamic source, {
  T unknownValue,
}) {
  if (source == null) {
    throw ArgumentError('A value must be provided. Supported values: '
        '${enumValues.values.join(', ')}');
  }

  final value = enumValues.entries
      .singleWhere((e) => e.value == source, orElse: () => null)
      ?.key;

  if (value == null && unknownValue == null) {
    throw ArgumentError('`$source` is not one of the supported values: '
        '${enumValues.values.join(', ')}');
  }
  return value ?? unknownValue;
}

T _$enumDecodeNullable<T>(
  Map<T, dynamic> enumValues,
  dynamic source, {
  T unknownValue,
}) {
  if (source == null) {
    return null;
  }
  return _$enumDecode<T>(enumValues, source, unknownValue: unknownValue);
}

const _$ProtocolEnumMap = {
  Protocol.WebSocket: 0,
  Protocol.MQTT: 1,
  Protocol.XMPP: 2,
  Protocol.LongPoll: 3,
};
