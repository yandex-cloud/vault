package configutil

import (
	"context"
	"fmt"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1"
	"os"
	"sync/atomic"

	"github.com/yandex-cloud/go-sdk/iamkey"

	wrapping "github.com/hashicorp/go-kms-wrapping/v2"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

const (
	// Accepted env vars
	EnvYandexCloudOAuthToken            = "YANDEXCLOUD_OAUTH_TOKEN"
	EnvYandexCloudServiceAccountKeyFile = "YANDEXCLOUD_SERVICE_ACCOUNT_KEY_FILE"
	EnvYandexCloudKMSKeyID              = "YANDEXCLOUD_KMS_KEY_ID"
	EnvYandexCloudEndpoint              = "YANDEXCLOUD_ENDPOINT"

	// Accepted config parameters
	CfgYandexCloudOAuthToken            = "oauth_token"
	CfgYandexCloudServiceAccountKeyFile = "service_account_key_file"
	CfgYandexCloudKMSKeyID              = "kms_key_id"
	CfgYandexCloudEndpoint              = "endpoint"

	WrapperTypeYandexCloudKms wrapping.WrapperType = "yandexcloudkms"
)

// Wrapper represents credentials and key information for the KMS Key used to
// encryption and decryption
type Wrapper struct {
	client           kms.SymmetricCryptoServiceClient
	keyID            string
	currentVersionID *atomic.Value
}

// Ensure that we are implementing Wrapper
var _ wrapping.Wrapper = (*Wrapper)(nil)

// NewWrapper creates a new Yandex.Cloud wrapper
func NewWrapper() *Wrapper {
	k := &Wrapper{
		currentVersionID: new(atomic.Value),
	}
	k.currentVersionID.Store("")
	return k
}

// SetConfig sets the fields on the Wrapper object based on
// values from the config parameter.
//
// Order of precedence Yandex.Cloud values:
// * Environment variable
// * Value from Vault configuration file
// * Compute Instance metadata
// func (k *Wrapper) SetConfig(config map[string]string) (map[string]string, error) {
func (wrapper *Wrapper) SetConfig(ctx context.Context, options ...wrapping.Option) (*wrapping.WrapperConfig, error) {
	opts, err := wrapping.GetOpts(options...)
	if err != nil {
		return nil, err
	}

	configMap := opts.WithConfigMap
	if configMap == nil {
		configMap = map[string]string{}
	}

	// Check and set versionId
	wrapper.keyID = coalesce(os.Getenv(EnvYandexCloudKMSKeyID), configMap[CfgYandexCloudKMSKeyID])
	if wrapper.keyID == "" {
		return nil, fmt.Errorf(
			"neither '%s' environment variable nor '%s' config parameter is set",
			EnvYandexCloudKMSKeyID, CfgYandexCloudKMSKeyID,
		)
	}

	// Check and set wrapper.client
	if wrapper.client == nil {
		client, err := getYandexCloudKMSClient(
			coalesce(os.Getenv(EnvYandexCloudOAuthToken), configMap[CfgYandexCloudOAuthToken]),
			coalesce(os.Getenv(EnvYandexCloudServiceAccountKeyFile), configMap[CfgYandexCloudServiceAccountKeyFile]),
			coalesce(os.Getenv(EnvYandexCloudEndpoint), configMap[CfgYandexCloudEndpoint]),
		)
		if err != nil {
			return nil, fmt.Errorf("error initializing Yandex.Cloud KMS client: %w", err)
		}

		if err := wrapper.setClient(client); err != nil {
			return nil, fmt.Errorf("error setting Yandex.Cloud KMS client: %w", err)
		}
	}

	wrapConfig := new(wrapping.WrapperConfig)
	wrapConfig.Metadata = make(map[string]string)
	wrapConfig.Metadata["kms_key_id"] = wrapper.keyID

	return wrapConfig, nil
}

func (wrapper *Wrapper) setClient(client kms.SymmetricCryptoServiceClient) error {
	wrapper.client = client

	// Make sure all the required permissions are granted (also checks if key exists)
	_, err := wrapper.Encrypt(context.Background(), []byte("go-kms-wrapping-test"), nil)
	if err != nil {
		return fmt.Errorf(
			"failed to encrypt with Yandex.Cloud KMS key - ensure the key exists and permission to encrypt the key is granted: %w", err)
	}

	return nil
}

func (wrapper *Wrapper) KeyId(_ context.Context) (string, error) {
	return wrapper.keyID, nil
}

func (wrapper *Wrapper) Type(ctx context.Context) (wrapping.WrapperType, error) {
	return WrapperTypeYandexCloudKms, nil
}

func (wrapper *Wrapper) versionId() string {
	return wrapper.currentVersionID.Load().(string)
}

func (wrapper *Wrapper) Encrypt(ctx context.Context, plaintext []byte, options ...wrapping.Option) (*wrapping.BlobInfo, error) {
	if plaintext == nil {
		return nil, fmt.Errorf("given plaintext for encryption is nil")
	}

	env, err := wrapping.EnvelopeEncrypt(plaintext, options...)
	if err != nil {
		return nil, fmt.Errorf("error wrapping data: %w", err)
	}

	if wrapper.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	encryptResponse, err := wrapper.client.Encrypt(
		ctx,
		&kms.SymmetricEncryptRequest{
			KeyId:     wrapper.keyID,
			Plaintext: env.Key,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error encrypting data: %w", err)
	}

	// Store the current version id
	wrapper.currentVersionID.Store(encryptResponse.VersionId)

	ret := &wrapping.BlobInfo{
		Ciphertext: env.Ciphertext,
		Iv:         env.Iv,
		KeyInfo: &wrapping.KeyInfo{
			// Even though we do not use the version id during decryption, store it
			// to know exactly the specific key version used in encryption in case we
			// want to rewrap older entries
			KeyId:      encryptResponse.VersionId,
			WrappedKey: encryptResponse.Ciphertext,
		},
	}

	return ret, nil
}

func (wrapper *Wrapper) Decrypt(ctx context.Context, ciphertext *wrapping.BlobInfo, options ...wrapping.Option) ([]byte, error) {
	if ciphertext == nil {
		return nil, fmt.Errorf("given input for decryption is nil")
	}

	decryptResponse, err := wrapper.client.Decrypt(
		ctx,
		&kms.SymmetricDecryptRequest{
			KeyId:      wrapper.keyID,
			Ciphertext: ciphertext.KeyInfo.WrappedKey,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error decrypting data encryption key: %w", err)
	}

	envInfo := &wrapping.EnvelopeInfo{
		Key:        decryptResponse.Plaintext,
		Iv:         ciphertext.Iv,
		Ciphertext: ciphertext.Ciphertext,
	}
	plaintext, err := wrapping.EnvelopeDecrypt(envInfo, options...)
	if err != nil {
		return nil, fmt.Errorf("error decrypting data: %w", err)
	}

	return plaintext, nil
}

func getYandexCloudKMSClient(oauthToken string, serviceAccountKeyFile string, endpoint string) (kms.SymmetricCryptoServiceClient, error) {
	credentials, err := getCredentials(oauthToken, serviceAccountKeyFile)
	if err != nil {
		return nil, fmt.Errorf("error getting credentials: %w", err)
	}

	sdk, err := ycsdk.Build(
		context.Background(),
		ycsdk.Config{
			Credentials: credentials,
			Endpoint:    endpoint,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error building Yandex.Cloud SDK instance: %w", err)
	}

	return sdk.KMSCrypto().SymmetricCrypto(), nil
}

func getCredentials(oauthToken string, serviceAccountKeyFile string) (ycsdk.Credentials, error) {
	if oauthToken != "" && serviceAccountKeyFile != "" {
		return nil, fmt.Errorf("error configuring authentication: both OAuth token and service account key file are specified")
	}

	// Yandex account authentication (via Oauth token)
	if oauthToken != "" {
		return ycsdk.OAuthToken(oauthToken), nil
	}

	// Service account authentication (via authorized key)
	if serviceAccountKeyFile != "" {
		key, err := iamkey.ReadFromJSONFile(serviceAccountKeyFile)
		if err != nil {
			return nil, fmt.Errorf("error reading service account key file: %w", err)
		}
		return ycsdk.ServiceAccountKey(key)
	}

	// Compute Instance Service Account authentication
	return ycsdk.InstanceServiceAccount(), nil
}

func coalesce(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
