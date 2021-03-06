#!/usr/bin/python

import os
import sys, getopt
import argparse, json
import requests
from subprocess import Popen, PIPE
from time import sleep, time

#Bye bye self signed unverified shit certificates warning
requests.packages.urllib3.disable_warnings(requests.packages.urllib3.exceptions.InsecureRequestWarning)

class Logger:
	def __init__(self):
		return

	def log(self, jdict):
		jdict['ts'] = int(time())
		print(json.dumps(jdict))

class Client:
	def __init__(self, target, cert='certificate.pem', key='privkey_open.pem', user='support', passwd='Super5upport!', logger=None):
		self.target = target
		self.cert = cert
		self.key = key
		self.user = user
		self.passwd = passwd
		self.headers = {'Content-Type': 'application/json', 'Accept': 'application/json'}
		self.logger = logger

	def fullurl(self, url):
		return "http{}://{}{}?user={}&password={}".format(
			"s" if self.cert is not None else "", 
			self.target, 
			url, 
			self.user, 
			self.passwd)

	def log(self, uri, method, res):
		logdata={
				"uri": uri,
				"method": method,
				"res": res.status_code,
				"data": res.text
			}
		if self.logger:
			self.logger.log(logdata)

	def put(self, url, data):
		res = requests.put(
			self.fullurl(url),
			data=json.dumps(data),
			headers = self.headers,
			cert=(self.cert, self.key), 
			verify=False)
		self.log(url, "put", res)
		return res

	def post(self, url, data):
		res = requests.post(
			self.fullurl(url),
			data=json.dumps(data),
			headers = self.headers,
			cert=(self.cert, self.key), 
			verify=False)
		self.log(url, "put", res)
		return res

	def get(self, url):
		res = requests.get(
			self.fullurl(url),
			headers = self.headers,
			cert=(self.cert, self.key), 
			verify=False)
		self.log(url, "get", res)
		return res

	def delete(self, url):
		res = requests.delete(
			self.fullurl(url),
			headers = self.headers,
			cert=(self.cert, self.key), 
			verify=False)
		self.log(url, "del", res)
		return res


class Api():
	def __init__(self, client):
		self.client = client

	def set_link(self):
		data = {
		"ifindex": 0,
		"ifname": "string",
		"mtu": 0,
		"linkinfo": {
			"info_kind": "dummy"
		},
		"link_type": "ether",
		"address": "string",
		"addr_info": [
			{
			"prefixlen": 0,
			"broadcast": "string"
			}
		]
		}
		self.client.post("/api/1/config/links", data=data)

	def set_sample_config(self):
		data = {
			"global":{
				"non-local-bind" : True
			},

			"links":[
				{
					"ifname":"dumby",
					"ifindex":4,
					"link_type": "ether",
					"addr_info":[
						{
							"local": "10.1.2.3",
							"prefixlen" : 24,
							"broadcast": "10.1.2.255"
						}
					],
					"linkinfo": {
						"info_kind": "dummy"
					},
					"attributes": {
						"dhcp" : True
					}
				}
			]
		}
		self.client.post("/api/1/config", data=data)

def main(argv):
	parser = argparse.ArgumentParser(description='API Automaton', epilog='Try harder...')

	parser.add_argument('-t', '--target', dest='target', 
		help='Specifies a target host', required=True)

	parser.add_argument('-c', '--cert', dest='cert', 
		help='Specifies the client certificate to use', default=None)

	parser.add_argument('-k', '--key', dest='keyfile', default=None, 
		help='''Specifies the keyfile to use must be open to be supported by requests: 

		openssl rsa -in keyfile.pem -out keyfile_open.pem''')

	parser.add_argument('-u', '--user', dest='user', 
		help='Specifies the API authentication user', default='support')

	parser.add_argument('-p', '--pass', dest='passwd', 
		help='Specifies the API authentication password', default='Super5upport!')
	
	args = parser.parse_args();

	logger = Logger()
	client = Client(args.target, args.cert, args.keyfile, args.user, args.passwd, logger)
	api = Api(client)

	api.set_link()
	exit(0)

if __name__ == "__main__":
	main(sys.argv[1:])

