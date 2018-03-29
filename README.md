# Instructions for CentOS Setup
---
##### Make sure curl, git, yum and wget are installed
---
* Setup Go, normal CentOS repos are old, use these fancy ones
   * `rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REP`
   * `curl -s https://mirror.go-repo.io/centos/go-repo.repo | tee /etc/yum.repos.d/go-repo.repo`
   * `yum install golang -y`


* Setup the Go environment
   * `mkdir "$HOME/go"`
   * `export GOPATH=$HOME/go && echo 'export GOPATH=$HOME/go' >> ~/.bashrc`
   * `export PATH=$PATH:$GOPATH/bin && echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc`


* Clone repository
   * `mkdir -p "$GOPATH/src/github.com/sillypears"`
   * `cd "$GOPATH/src/github.com/sillypears"`
   * `git clone https://github.com/sillypears/condor-standings.git`
   * `cd condor-standings`


* Create symlink to project in $HOME (optional, but recommended because easy)
   * `ln -s $HOME/go/src/github.com/sillypears/condor-standings cs`


* Pull all go dependencies
   * `cd src && go get ./... && cd ..`


* Setup environment file and configure
   * `cp .env_template .env`
   * `vim .env`
      * Domain is the public domain name and is required
      * Random 64 digit key is required
      * Database connectivity is required
      * Fill in the rest if wanted, HTTPS is suggested

# To run the server
---
   * `cd "$GOPATH/src/github.com/sillypears/condor-standings` OR `cd $HOME/cs`
   * `go run src/*.go`

* Setting up HTTPS
   * `yum install epel-release -y`
   * `yum install certbot certbot-apache -y`
   * `yum install httpd -y`
   * `echo -e "<VirtualHost *:80>\nServerAdmin admin@test.com\nDocumentRoot "/usr/share/httpd"\nServerName test.com\nServerAlias www.test.com\n</VirtualHost>" > /etc/httpd/conf.d/test.conf`
   * `systemctl start httpd`
   * `certbot certonly --apache -d <domain-name.com> -d <another.domain-name.com> --email <your@email.address> --agree-tos`
   * TLS_CERT_FILE = `/etc/letsencrypt/live/wow.freepizza.how/fullchain.pem` 
   * TLS_KEY_FILE = `/etc/letsencrypt/live/wow.freepizza.how/privkey.pem`
