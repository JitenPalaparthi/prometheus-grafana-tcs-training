# Install grafana on Ubuntu

```bash
sudo apt-get install -y apt-transport-https
sudo apt-get install -y software-properties-common wget
sudo wget -q -O /usr/share/keyrings/grafana.key https://apt.grafana.com/gpg.key
```
- Stable release

```bash
echo "deb [signed-by=/usr/share/keyrings/grafana.key] https://apt.grafana.com stable main" | sudo tee -a /etc/apt/sources.list.d/grafana.list
```

- Update list of packages

```bash
sudo apt-get update
```

- Install grafana

```bash
sudo apt-get install grafana
```

- Start grafana server

```bash
 sudo /bin/systemctl daemon-reload
 sudo /bin/systemctl enable grafana-server # start when system starts
 sudo /bin/systemctl start grafana-server
```

- Access grafana UI

- open localhost:3000 in the browser
- default username: admin
- default password: admin


# To Uninstall grafana on Ubuntu

- Stop grafana service

```bash
sudo systemctl stop grafana-server
sudo apt-get remove grafana
sudo rm -i /etc/apt/sources.list.d/grafana.list
```


