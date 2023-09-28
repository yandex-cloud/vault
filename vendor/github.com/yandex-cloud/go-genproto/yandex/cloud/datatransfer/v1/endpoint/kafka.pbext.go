// Code generated by protoc-gen-goext. DO NOT EDIT.

package endpoint

type KafkaConnectionOptions_Connection = isKafkaConnectionOptions_Connection

func (m *KafkaConnectionOptions) SetConnection(v KafkaConnectionOptions_Connection) {
	m.Connection = v
}

func (m *KafkaConnectionOptions) SetClusterId(v string) {
	m.Connection = &KafkaConnectionOptions_ClusterId{
		ClusterId: v,
	}
}

func (m *KafkaConnectionOptions) SetOnPremise(v *OnPremiseKafka) {
	m.Connection = &KafkaConnectionOptions_OnPremise{
		OnPremise: v,
	}
}

func (m *OnPremiseKafka) SetBrokerUrls(v []string) {
	m.BrokerUrls = v
}

func (m *OnPremiseKafka) SetTlsMode(v *TLSMode) {
	m.TlsMode = v
}

func (m *OnPremiseKafka) SetSubnetId(v string) {
	m.SubnetId = v
}

type KafkaAuth_Security = isKafkaAuth_Security

func (m *KafkaAuth) SetSecurity(v KafkaAuth_Security) {
	m.Security = v
}

func (m *KafkaAuth) SetSasl(v *KafkaSaslSecurity) {
	m.Security = &KafkaAuth_Sasl{
		Sasl: v,
	}
}

func (m *KafkaAuth) SetNoAuth(v *NoAuth) {
	m.Security = &KafkaAuth_NoAuth{
		NoAuth: v,
	}
}

func (m *KafkaSaslSecurity) SetUser(v string) {
	m.User = v
}

func (m *KafkaSaslSecurity) SetPassword(v *Secret) {
	m.Password = v
}

func (m *KafkaSaslSecurity) SetMechanism(v KafkaMechanism) {
	m.Mechanism = v
}

func (m *KafkaSource) SetConnection(v *KafkaConnectionOptions) {
	m.Connection = v
}

func (m *KafkaSource) SetAuth(v *KafkaAuth) {
	m.Auth = v
}

func (m *KafkaSource) SetSecurityGroups(v []string) {
	m.SecurityGroups = v
}

func (m *KafkaSource) SetTopicName(v string) {
	m.TopicName = v
}

func (m *KafkaSource) SetTransformer(v *DataTransformationOptions) {
	m.Transformer = v
}

func (m *KafkaSource) SetParser(v *Parser) {
	m.Parser = v
}

func (m *KafkaTarget) SetConnection(v *KafkaConnectionOptions) {
	m.Connection = v
}

func (m *KafkaTarget) SetAuth(v *KafkaAuth) {
	m.Auth = v
}

func (m *KafkaTarget) SetSecurityGroups(v []string) {
	m.SecurityGroups = v
}

func (m *KafkaTarget) SetTopicSettings(v *KafkaTargetTopicSettings) {
	m.TopicSettings = v
}

type KafkaTargetTopicSettings_TopicSettings = isKafkaTargetTopicSettings_TopicSettings

func (m *KafkaTargetTopicSettings) SetTopicSettings(v KafkaTargetTopicSettings_TopicSettings) {
	m.TopicSettings = v
}

func (m *KafkaTargetTopicSettings) SetTopic(v *KafkaTargetTopic) {
	m.TopicSettings = &KafkaTargetTopicSettings_Topic{
		Topic: v,
	}
}

func (m *KafkaTargetTopicSettings) SetTopicPrefix(v string) {
	m.TopicSettings = &KafkaTargetTopicSettings_TopicPrefix{
		TopicPrefix: v,
	}
}

func (m *KafkaTargetTopic) SetTopicName(v string) {
	m.TopicName = v
}

func (m *KafkaTargetTopic) SetSaveTxOrder(v bool) {
	m.SaveTxOrder = v
}
