package configuration

type PlugingConfiguration struct {
	Slack []struct {
		Token string `yaml:"token"`
		Url   string `yaml:"url"`
	} `yaml:"slack"`
}
