# Install and configure alertmanager

- Download alertmanager

```bash
curl -s https://api.github.com/repos/prometheus/alertmanager/releases/latest | grep browser_download_url | grep linux-amd64 | cut -d '"' -f 4 | wget -qi -
```

- untar (version of alertmanager changes)

```bash
tar xzfv alertmanager-0.25.0.linux-amd64.tar.gz
```

- Copy to /usr/local/bin

```bash
cd alertmanager-0.25.0
cp -r . /usr/local/bin/alertmanager
```

- CSreate a service file

```bash
sudo vi/etc/systemd/system/alertmanager.service
```
- add following contents to the file as [alertmanager.service file](alertmanager.service)

```bash
[Unit]
Description=Prometheus Alert Manager Service
After=network.target
[Service]
Type=simple
ExecStart=/usr/local/bin/alertmanager/alertmanager \
        --config.file=/usr/local/bin/alertmanager/alertmanager.yml 
[Install]
WantedBy=multi-user.target
```

- Run alertmanager as a service

```bash
sudo service alertmanager start
sudo service alertmanager status
```
or

```bash
sudo systemctl daemon-reload
sudo systemctl start alertmanager
sudo systemctl status alertmanager
```

- Configure alertmanager configuration

```bash
sudo vi /usr/local/bin/alertmanager/alertmanager.yml
```
- Add/replace contents like in [alertmanager config file](alertmanager.yml)

- Make sure configs are correct

```bash
/usr/local/bin/alertmanager/amtool check-config
```
- Create a prometheus_rules.yml file to configure alert rules and templates.[Prometheus Rules for Alertmanagement](prometheus_rules.yml)

- Update prometheus configuration as in [updated prometheus config for alertmanager](prometheus.yml)

- Restart prometheus service

```bash
sudo service prometheus stop
sudo service prometheus start
```
or

```bash
sudo systemctl stop prometheus
sudo systemctl start prometheus
sudo systemctl status prometheus
```
