package ec2provider

type ec2ProviderNoop struct {
}

func NewNoopProvider() *ec2ProviderNoop {
	return &ec2ProviderNoop{}
}

func (e *ec2ProviderNoop) GetPrivateDNSName(s string) (string, error) {
	return "", nil
}

func (e *ec2ProviderNoop) StartEc2DescribeBatchProcessing() {
}
