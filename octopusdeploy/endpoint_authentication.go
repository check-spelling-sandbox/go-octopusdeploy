package octopusdeploy

type EndpointAuthentication struct {
	AccountID                 string `json:"AccountId,omitempty"`
	AdminLogin                string `json:"AdminLogin,omitempty"`
	AssumeRole                bool   `json:"AssumeRole,omitempty"`
	AssumedRoleARN            string `json:"AssumedRoleArn,omitempty"`
	AssumedRoleSession        string `json:"AssumedRoleSession,omitempty"`
	AssumeRoleSessionDuration int    `json:"AssumeRoleSessionDurationSeconds,omitempty"`
	AssumeRoleExternalID      string `json:"AssumeRoleExternalId,omitempty"`
	AuthenticationType        string `json:"AuthenticationType,omitempty"`
	ClientCertificate         string `json:"ClientCertificate,omitempty"`
	ClusterName               string `json:"ClusterName,omitempty"`
	ClusterResourceGroup      string `json:"ClusterResourceGroup,omitempty"`
	UseInstanceRole           bool   `json:"UseInstanceRole,omitempty"`
}
