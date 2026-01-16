sudo apt-get update
sudo apt-get -y install docker.io

curl -H Metadata-Flavor:Google 169.254.169.254/computeMetadata/v1/instance/service-accounts/default/token | \
cut -f1 -d',' | \
cut -f2 -d':' | \
tr -d '"' | \
sudo docker login --username iam --password-stdin $DOMAIN

curl -H Metadata-Flavor:Google 169.254.169.254/computeMetadata/v1/instance/service-accounts/default/token | \
cut -f1 -d',' | \
cut -f2 -d':' | \
tr -d '"' | \
sudo docker login --username iam --password-stdin $DOCKER_ENDPOINT

sudo docker pull $DOCKER_ENDPOINT/$REGISTRY_ID/vault:$BASE_VERSION-yckms

sudo docker run -p8200:8200 --cap-add=IPC_LOCK -e "VAULT_LOCAL_CONFIG=$(< vault.hcl)" -e "VAULT_ADDR=http://127.0.0.1:8200" -v $(pwd)/startup.sh:/startup.sh -v $(pwd)/vault.hcl:/vault.hcl -v $(pwd)/key.json:/vault/config/key.json $DOMAIN/$REGISTRY_ID/vault:$BASE_VERSION-yckms /bin/sh "/startup.sh"

exit