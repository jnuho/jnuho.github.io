

## ElastiCache

- Create a dedicated SG (TCP port 6379 open) with a default VPC for ElastiCache cluster
- ElastiCache > Create with SG above && create subnetgroup that uses same VPC as above
- EC2 > Create with a new SG (open port 22 for SSH only)


- Test EC2 - ElastiCache Connection

```sh
# install redis client
wget http://download.redis.io/redis-stable.tar.gz
tar -xvzf redis-stable.tar.gz
cd redis-stable
make && make install
redis-cli -h {ElastiCache_Endpoint} -p {port configured in SG above;6379}

> flushall
> keys *
```
