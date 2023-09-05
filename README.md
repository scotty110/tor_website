# Setting up Tor Website

## Setup Tor
Update server `sudo apt update && sudo apt upgrade` 

### Install tor
Latest: https://support.torproject.org/apt/tor-deb-repo/
#### Setup Tor Repo
1. `sudo apt install apt-transport-https`
2. Create `/etc/apt/sources.list.d/tor.list` and add:
```
deb     [signed-by=/usr/share/keyrings/tor-archive-keyring.gpg] https://deb.torproject.org/torproject.org <DISTRIBUTION> main
deb-src [signed-by=/usr/share/keyrings/tor-archive-keyring.gpg] https://deb.torproject.org/torproject.org <DISTRIBUTION> main
```
3. `wget -qO- https://deb.torproject.org/torproject.org/A3C4F0F979CAA22CDBA8F512EE8CBC9E886DDD89.asc | gpg --dearmor | sudo tee /usr/share/keyrings/tor-archive-keyring.gpg >/dev/null`
4. `sudo apt update`
5. `sudo apt install tor deb.torproject.org-keyring`


### Edit torrc
`sudo vi /etc/tor/torrc`
Uncomment 
```
HiddenServiceDir /var/lib/tor/hidden_service/
HiddenServicePort 80 127.0.0.1:80
```

Restart Tor
```
sudo service tor stop
sudo service tor start
sudo service tor status
```

### Get Website
`sudo cat /var/lib/tor/hidden_service/hostname`

## Why
I wanted a testbed for setting up a public website with possible endpoints, and tor allowed me to do that easier than setting up a domain.

