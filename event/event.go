package event

type Topic struct {
	Domain        string
	Name          string
	ID            string
	Events        map[string]string
	ProtoPackage  string
	ProtoName     string
	Keys          []string
	PartitionKeys []string
	ExtraHeaders  []string
}

func (t Topic) DLQTopic() Topic {
	return Topic{
		Domain:        "DLQ",
		Name:          t.Name + ".dlq",
		ID:            t.ID + ".dlq",
		Events:        t.Events,
		ProtoPackage:  t.ProtoPackage,
		ProtoName:     t.ProtoName,
		Keys:          t.Keys,
		PartitionKeys: t.PartitionKeys,
		ExtraHeaders:  t.ExtraHeaders,
	}
}

func (t Topic) RPQTopic() Topic {
	return Topic{
		Domain:        "RPQ",
		Name:          t.Name + ".rpq",
		ID:            t.ID + ".rpq",
		Events:        t.Events,
		ProtoPackage:  t.ProtoPackage,
		ProtoName:     t.ProtoName,
		Keys:          t.Keys,
		PartitionKeys: t.PartitionKeys,
		ExtraHeaders:  t.ExtraHeaders,
	}
}
