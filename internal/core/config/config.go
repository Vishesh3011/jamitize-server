package config

type AppConfig interface {
	DBConfig() *dbConfig
	LoggerConfig() *loggerConfig
	JwtConfig() *jwtConfig
}

type appConfig struct {
	*dbConfig
	*loggerConfig
	*jwtConfig
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

	jwtConfig, err := newJWTConfig()
	if err != nil {
		return nil, err
	}

	return &appConfig{
		dbConfig:     dbConfig,
		loggerConfig: loggerConfig,
		jwtConfig:    jwtConfig,
	}, nil
}

func (c appConfig) DBConfig() *dbConfig {
	return c.dbConfig
}

func (c appConfig) LoggerConfig() *loggerConfig {
	return c.loggerConfig
}

func (c appConfig) JwtConfig() *jwtConfig {
	return c.jwtConfig
}
