package providers

type Configuration interface {
	Token string `yaml:"token"`
	Url string `yaml:"url"`
}

func init() {
	// registerProviders = append(registerProviders, func() (Provider, bool){
	// 	return Slack
	// })

	registerProviders = append(registerProviders, func(getenv func(string) string) (Provider, bool) {
	  token := ""
	  return Slack(token), true
	})
}

func Load(configurationPath string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(configurationPath)
	if err != nil {
		return nil, err
	}

	c := Configuration{}
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		log.Printf("Failed to validate syntax: %s", configurationPath)
		return nil, err
	}

	return &c, nil
}
