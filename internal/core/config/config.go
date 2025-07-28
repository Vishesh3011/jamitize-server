package config

type AppConfig interface {
	DBConfig() *dbConfig
	LoggerConfig() *loggerConfig
}

type appConfig struct {
	*dbConfig
	*loggerConfig
}

func NewAppConfig() (AppConfig, error) {
	dbConfig, err := newDBConfig()
	if err != nil {
		return nil, err
	}

	loggerConfig, err := newLoggerConfig()
	if err != nil {
		return nil, err
	}
	return &appConfig{
		dbConfig:     dbConfig,
		loggerConfig: loggerConfig,
	}, nil
}

func (c appConfig) DBConfig() *dbConfig {
	return c.dbConfig
}

func (c appConfig) LoggerConfig() *loggerConfig {
	return c.loggerConfig
}
