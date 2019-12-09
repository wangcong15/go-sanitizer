#!/usr/bin/env python
#coding=utf-8
import os
import sys, getopt
import re
import requests

dirPath = os.getcwd()
opts, args = getopt.getopt(sys.argv[1:], "f:x:s:p:")
project_zip = ""
xml_zip = ""
server_ip = ""
server_port = "8080"
for op, value in opts:
    if op == "-f":
        project_zip = value
    elif op == "-x":
        xml_zip = value
    elif op == "-s":
        server_ip = value
    elif op == "-p":
    	server_port = value
if project_zip == "":
	print "缺少项目文件的指定"
	exit()
if xml_zip == "":
	print "缺少xml文件的指定"
	exit()
project_path = os.path.join(dirPath, project_zip)
if not os.path.exists(project_path) or not project_path.endswith(".zip"):
	print "项目文件不存在或文件格式错误"
	exit()
xml_path = os.path.join(dirPath, xml_zip)
if not os.path.exists(xml_path) or not xml_path.endswith(".zip"):
	print "xml文件不存在或文件格式错误"
	exit()
pattern = re.compile("((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)")
if not pattern.match(server_ip) and not server_ip == "localhost":
	print "IP地址错误"
	exit()
pattern2 = re.compile("^([0-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-5]{2}[0-3][0-5])$")
if not pattern2.match(server_port):
	print "端口号输入错误"
	exit()

r = requests.post("http://" + server_ip + ":" + server_port + "/history/", files={'file':open(project_path, 'rb')})
result = (str(r) == "<Response [200]>")

r = requests.post("http://" + server_ip + ":" + server_port + "/history/", files={'file':open(xml_path, 'rb')})
result2 = (str(r) == "<Response [200]>")

if result and result2:
	print "上传成功"
else:
	print "上传失败，请重新尝试"

