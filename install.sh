#!/bin/bash -e

systemctl stop wrench
cp wrench /home/gregory/bin/wrench
systemctl start wrench

cp nginx.conf /etc/nginx/sites-available/wrench.preferit.se
systemctl reload nginx
