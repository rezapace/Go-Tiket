package config

//NOTE :
// direktory ini berfungsi untuk konfigurasi database, port, dan lain-lain
import (
	"errors"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Config adalah konfigurasi untuk aplikasi
type Config struct {
	Env            string         `env:"ENV" envDefault:"development"`
	Port           string         `env:"PORT" envDefault:"8080"`
	Postgres       PostgresConfig `envPrefix:"POSTGRES_"`
	JWT            JwtConfig      `envPrefix:"JWT_"`
	MidtransConfig MidtransConfig `envPrefix:"MIDTRANS_"`
}

type MidtransConfig struct {
	BaseURL   string `env:"BASE_URL"`
	ClientKey string `env:"CLIENT_KEY"`
	ServerKey string `env:"SERVER_KEY"`
}

// JwtConfig adalah konfigurasi untuk JWT
type JwtConfig struct {
	SecretKey string `env:"SECRET_KEY"`
}

// PostgresConfig adalah konfigurasi untuk koneksi ke database postgres
type PostgresConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Database string `env:"DATABASE" envDefault:"postgres"`
}

// untuk membuat new config
func NewConfig(envPath string) (*Config, error) {
	cfg, err := parseConfig(envPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// untuk parse config
func parseConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, errors.New("failed to load env")
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return nil, errors.New("failed to parse config")
	}
	return cfg, nil
}

// func init() {
// 	// Isi konfigurasi Midtrans Anda di sini
// 	MidtransConfig.ClientKey = ""
// 	MidtransConfig.ServerKey = ""
// 	MidtransConfig.IsProduction = false // Ganti menjadi true jika sudah di production
// }
