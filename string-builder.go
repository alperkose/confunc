package confunc

type StringBuilder struct{
	source Source
	sourceKey string
}

func (sb *StringBuilder) Build() String{
	return func() string {
		return sb.source.Value(sb.sourceKey)
	}
}