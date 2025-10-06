// sentiric-vertical-finance-service/internal/config/config.go
package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	GRPCPort     string
	HttpPort     string
	CertPath     string
	KeyPath      string
	CaPath       string
	LogLevel     string
	Env          string
    
    // Finans servisi bağımlılıkları (Placeholder)
    CoreBankingSystem string // API Gateway URL'si
    CoreBankingAPIKey string
}

func Load() (*Config, error) {
	godotenv.Load()

	// Harmonik Mimari Portlar (Dikey Servisler, 206XX bloğu atandı)
	return &Config{
		GRPCPort:     GetEnv("VERTICAL_FINANCE_SERVICE_GRPC_PORT", "20611"),
		HttpPort:     GetEnv("VERTICAL_FINANCE_SERVICE_HTTP_PORT", "20610"),
		
		CertPath:     GetEnvOrFail("VERTICAL_FINANCE_SERVICE_CERT_PATH"),
		KeyPath:      GetEnvOrFail("VERTICAL_FINANCE_SERVICE_KEY_PATH"),
		CaPath:       GetEnvOrFail("GRPC_TLS_CA_PATH"),
		LogLevel:     GetEnv("LOG_LEVEL", "info"),
		Env:          GetEnv("ENV", "production"),

        CoreBankingSystem: GetEnv("CORE_BANKING_SYSTEM", "mock_bank_api"),
        CoreBankingAPIKey: GetEnv("CORE_BANKING_API_KEY", ""),
	}, nil
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetEnvOrFail(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal().Str("variable", key).Msg("Gerekli ortam değişkeni tanımlı değil")
	}
	return value
}