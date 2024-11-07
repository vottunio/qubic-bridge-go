package config

import (
	"os"
	"time"

	"github.com/vottundev/vottun-qubic-bridge-go/utils/crypto"
	"github.com/vottundev/vottun-qubic-bridge-go/utils/log"

	"gopkg.in/yaml.v2"
)

type ChainType string

const (
	CHAIN_ETH ChainType = "ETH"
	CHAIN_ARB ChainType = "ARB"
)

type ChainsMap map[ChainType]ChainInfo

type ChainInfo struct {
	ChainID uint32 `yaml:"chain-id"`
	Name    string `yaml:"name"`
	RpcUrl  string `yaml:"rpc"`
	WssUrl  string `yaml:"wss"`
}

type CacheInfo struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Database int    `yaml:"database"`
}

type config struct {
	ServiceID uint64 `yaml:"service-id"`
	Http      struct {
		Route string `yaml:"route"`
		Port  uint16 `yaml:"port"`
	}
	Cors struct {
		AllowedOrigins []string `yaml:"allowed-origins"`
		AllowedMethods []string `yaml:"allowed-methods"`
		AllowedHeaders []string `yaml:"allowed-headers"`
	} `yaml:"cors"`
	Evm struct {
		InfuraKey string    `yaml:"infura-key"`
		Chains    ChainsMap `yaml:"chains"`
	} `yaml:"evm"`
	Jwt struct {
		PublicKey             string `yaml:"public-key"`
		PrivateKey            string `yaml:"private-key"`
		TokenCreationLifeTime int    `yaml:"token-creation-life-time"`
		TokenSecurityLifeTime int    `yaml:"token-security-life-time"`
	} `yaml:"jwt"`
	MySQL struct {
		Host               string `yaml:"host"`
		Port               uint16 `yaml:"port"`
		User               string `yaml:"user"`
		Password           string `yaml:"password"`
		DatabaseName       string `yaml:"database"`
		Timeout            uint16 `yaml:"timeout"`
		PageMaxRowsDefault uint32 `yaml:"page-max-rows-default"`
		PageMaxRowsLimit   uint32 `yaml:"page-max-rows-limit"`
		MaxOpenConns       uint16 `yaml:"max-open-conns"`
		MaxIdleConns       uint16 `yaml:"max-idle-conns"`
		Jwt                struct {
			User         string `yaml:"user"`
			Password     string `yaml:"password"`
			DatabaseName string `yaml:"database"`
			Timeout      uint16 `yaml:"timeout"`
		} `yaml:"jwt"`
	} `yaml:"mysql"`
	Cache struct {
		Connections        map[string]CacheInfo `yaml:"connections"`
		QubicEventsChannel string               `yaml:"qubic-events-channel"`
	} `yaml:"cache"`
	Telegram struct {
		BotUrl                    string  `yaml:"bot-url"`
		EnableBot                 bool    `yaml:"enable-bot"`
		Token                     string  `yaml:"token"`
		InitDataExpirationSeconds uint64  `yaml:"init-data-expiration-seconds"`
		AdminsAllowed             []int64 `yaml:"admins-allowed"`
	} `yaml:"telegram"`
	Urls struct {
		Images   string `yaml:"images"`
		ImagesS3 string `yaml:"images-s3"`
	} `yaml:"urls"`
	Game struct {
		EnergyPerSecond uint8  `yaml:"energy-per-second"`
		MaxUserLevel    uint64 `yaml:"max-user-level"`
		ReferralPrefix  string `yaml:"referral-prefix"`
	} `yaml:"game"`
	Queue struct {
		Active          bool   `yaml:"active"`
		BotSendMessages string `yaml:"bot-send-messages"`
		Profile         string `yaml:"profile"`
		Region          string `yaml:"region"`
		DelaySeconds    int64  `yaml:"delay-seconds"`
	} `yaml:"queue"`
}

// type vtnConfig struct {
// 	vtnTokenInfo domain.VtnTokenConfigInfo
// }

// Config contains yaml config
var Config config

var Environment uint8

// Secret Key
var secret string

//var VtnConfig vtnConfig

var ExecutionTime time.Time

func GetSecret() string {
	return secret
}

// CreateProperties Creates Properties
func CreateProperties(file string, key string) {

	yamlFile, err := os.ReadFile(file)
	if err != nil {
		log.Errorf("Error reading YAML file: %s\n", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &Config)

	if err != nil {
		log.Errorf("Error parsing YAML file: %s\n", err)
	}
	secret = key

}

func GetEncryptedProperty(propertyValue string) string {
	var result string

	if len(propertyValue) > 4 {

		if propertyValue[0:4] == "ENC(" {
			encrypted := propertyValue[4 : len(propertyValue)-1]
			result = crypto.Decrypt(secret, encrypted)
		} else {
			result = propertyValue
		}
	}

	return result
}