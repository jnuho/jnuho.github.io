- Local

```sh
# Install Go

# Install nodejs 18
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -

# install Node.js 18.x and npm
sudo apt-get install -y nodejs

## You may also need development tools to build native addons:
sudo apt-get install gcc g++ make

## To install the Yarn package manager, run:
curl -sL https://dl.yarnpkg.com/debian/pubkey.gpg | gpg --dearmor | sudo tee /usr/share/keyrings/yarnkey.gpg >/dev/null
echo "deb [signed-by=/usr/share/keyrings/yarnkey.gpg] https://dl.yarnpkg.com/debian stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
sudo apt update && sudo apt install yarn

# Install go-playground local
git clone https://github.com/x1unix/go-playground.git
cd go-playground
make
sudo make install
```


- Docker

```sh
docker pull x1unix/go-playground
```
