sudo ./gen-certs.sh certs

sudo chmod 400 certs/
sudo chown 999 certs/

docker-compose up -d --build

wait
sleep 30

docker exec redis_tls_shopview bash -c "redis-server /usr/local/etc/redis/redis.conf"

sudo chmod 555 -R certs/

