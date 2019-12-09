#!/usr/bin/env python
#coding=utf-8
 
import os, sys, platform
import posixpath
import BaseHTTPServer
from SocketServer import ThreadingMixIn
import threading
import urllib
import cgi
import shutil
import mimetypes
import re
import time
import webbrowser
import json

try:
    from cStringIO import StringIO
except ImportError:
    from StringIO import StringIO
     
 
print ""
print '----------------------------------------------------------------------->> '
try:
    port = int(sys.argv[1])
except Exception, e:
    print '-------->> 警告: 由于没有输入端口，默认使用8080端口 '
    print '-------->> 如果你希望使用其他端口，请执行如下操作: '
    print '-------->> python server.py port '
    print "-------->> port是一个整数，取值范围在: 1024～65535 "
    port = 8080
    
if not 1024 < port < 65535:
    port = 8080
serveraddr = ('', port)
print '-------->> 现在正在监听' + str(port) + '端口...'
print '-------->> 你可以通过这个链接访问:   http://localhost:' + str(port)
print '----------------------------------------------------------------------->> '
print ""

webbrowser.open('http://localhost:' + str(port), new=0, autoraise=True)
 
def sizeof_fmt(num):
    for x in ['bytes','KB','MB','GB']:
        if num < 1024.0:
            return "%3.1f%s" % (num, x)
        num /= 1024.0
    return "%3.1f%s" % (num, 'TB')
 
def modification_date(filename):
    return time.strftime("%Y-%m-%d %H:%M:%S",time.localtime(os.path.getmtime(filename)))
 
class SimpleHTTPRequestHandler(BaseHTTPServer.BaseHTTPRequestHandler):
    def do_GET(self):
        """Serve a GET request."""
        f = self.send_head()
        if f:
            self.copyfile(f, self.wfile)
            f.close()
 
    def do_HEAD(self):
        """Serve a HEAD request."""
        f = self.send_head()
        if f:
            f.close()
 
    def do_POST(self):
        # 检查上传目录
        if self.path != "/history/":
            f = self.send_head()
            if f:
                self.copyfile(f, self.wfile)
                f.close()
            return

        """Serve a POST request."""
        r, info = self.deal_post_data()
        print r, info, "by: ", self.client_address
        f = StringIO()
        f.write('<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">')
        f.write("<html>\n<title>上传结果页面</title>\n")
        f.write('<meta http-equiv="Content-Type" content="text/html;" charset="UTF-8"/>')
        f.write("<body>\n<h2>上传结果页面</h2>\n")
        f.write("<hr>\n")
        if r:
            f.write("<strong>上传成功:</strong>")
        else:
            f.write("<strong>上传失败:</strong>")
        f.write(info)
        f.write("<br><a href=\"/\">返回</a>")
        f.write("</body>\n</html>\n")
        length = f.tell()
        f.seek(0)
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.send_header("Content-Length", str(length))
        self.end_headers()
        if f:
            self.copyfile(f, self.wfile)
            f.close()
         
    def deal_post_data(self):
        boundary = self.headers.plisttext.split("=")[1]
        remainbytes = int(self.headers['content-length'])
        line = self.rfile.readline()
        remainbytes -= len(line)
        if not boundary in line:
            return (False, "Content NOT begin with boundary")
        line = self.rfile.readline()
        remainbytes -= len(line)
        fn = re.findall(r'Content-Disposition.*name="file"; filename="(.*)"', line)
        if not fn:
            return (False, "Can't find out file name...")

        path = self.translate_path(self.path)
        osType = platform.system()
        try:
            if osType == "Linux":
                fn = os.path.join(path, fn[0].decode('gbk').encode('utf-8'))
            else:
                fn = os.path.join(path, fn[0])
        except Exception, e:
            return (False, "文件名请不要用中文，或者使用IE上传中文名的文件。")
        while os.path.exists(fn):
            fn = "_" + fn
        line = self.rfile.readline()
        remainbytes -= len(line)
        if line.startswith("Content-Type"):
            line = self.rfile.readline()
            remainbytes -= len(line)
        try:
            out = open(fn, 'wb')
        except IOError:
            return (False, "Can't create file to write, do you have permission to write?")
                 
        preline = self.rfile.readline()
        remainbytes -= len(preline)
        while remainbytes > 0:
            line = self.rfile.readline()
            remainbytes -= len(line)
            if boundary in line:
                preline = preline[0:-1]
                if preline.endswith('\r'):
                    preline = preline[0:-1]
                out.write(preline)
                out.close()
                if out.name.endswith(".zip"):
                    os.system("cd history/;unzip " + out.name + ";rm " + out.name)
                    self.get_zNodes(out.name[0:-4] + "/")
                return (True, "文件 '%s' 上传成功！" % fn)
            else:
                out.write(preline)
                preline = line
        return (False, "Unexpect Ends of data.")

    def get_zNodes(self, path):
        result = []
        root_node = {'name':"Projects", 'open':'true'}
        sub_array = []
        sub_files = os.listdir(path)
        for file in sub_files:
            file_path = os.path.join(path, file)
            if os.path.isdir(file_path):
                sub_array.append({'name': file, 'isParent': 1, 'children': self.get_sub_zNodes(file_path)})
        for file in sub_files:
            file_path = os.path.join(path, file)
            #if os.path.isfile(file_path) and file_path.endswith('.c'):
            if os.path.isfile(file_path):
                sub_array.append({'name': file})
        root_node['children'] = sub_array
        result.append(root_node)
        json.dump(result, open(path + 'proj_data.json', 'w'))

    def get_sub_zNodes(self, path):
        result = []
        sub_files = os.listdir(path)
        for file in sub_files:
            file_path = os.path.join(path, file)
            if os.path.isdir(file_path):
                result.append({'name': file, 'isParent': 1, 'children': self.get_sub_zNodes(file_path)})
        for file in sub_files:
            file_path = os.path.join(path, file)
            #if os.path.isfile(file_path) and file_path.endswith('.c'):
            if os.path.isfile(file_path):
                result.append({'name': file})
        return result

    def send_head(self):
        path = self.translate_path(self.path)
        f = None
        if os.path.isdir(path):
            if not self.path.endswith('/'):
                # redirect browser - doing basically what apache does
                self.send_response(301)
                self.send_header("Location", self.path + "/")
                self.end_headers()
                return None
            for index in "index.html", "index.htm":
                index = os.path.join(path, index)
                if os.path.exists(index):
                    path = index
                    break
            else:
                return self.list_directory(path)
        ctype = self.guess_type(path)
        try:
            f = open(path, 'rb')
        except IOError:
            self.send_error(404, "File not found")
            return None
        self.send_response(200)
        self.send_header("Content-type", ctype)
        fs = os.fstat(f.fileno())
        self.send_header("Content-Length", str(fs[6]))
        self.send_header("Last-Modified", self.date_time_string(fs.st_mtime))
        self.end_headers()
        return f
 
    def list_directory(self, path):
        try:
            list = os.listdir(path)
        except os.error:
            self.send_error(404, "No permission to list directory")
            return None
        list.sort(key=lambda a: a.lower())
        f = StringIO()
        displaypath = cgi.escape(urllib.unquote(self.path))
        f.write('<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">')
        f.write("<html'>\n<title>文件目录: %s</title>\n" % displaypath)
        f.write('<meta http-equiv="Content-Type" content="text/html;" charset="UTF-8"/>')
        f.write("<body>\n<h2>文件目录: %s</h2>\n" % displaypath)
        f.write("<hr>\n")
        f.write("<form ENCTYPE=\"multipart/form-data\" method=\"post\">")
        f.write("<input name=\"file\" type=\"file\"/>")
        f.write("<input type=\"submit\" value=\"上传\"/>")
        f.write("&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp")
        f.write("<input type=\"button\" value=\"返回根目录\" onClick=\"location='/'\">")
        f.write("<input type=\"button\" value=\"查看历史提交内容\" onClick=\"location='/visualizer.html'\">")
        f.write("</form>\n")
        f.write("<hr>\n<ul>\n")
        for name in list:
            fullname = os.path.join(path, name)
            colorName = displayname = linkname = name
            if os.path.isdir(fullname):
                colorName = '<span style="background-color: #CEFFCE;">' + name + '/</span>'
                displayname = name
                linkname = name + "/"
            if os.path.islink(fullname):
                colorName = '<span style="background-color: #FFBFFF;">' + name + '@</span>'
                displayname = name
            filename = os.getcwd() + '/' + displaypath + displayname
            f.write('<table><tr><td width="60%%"><a href="%s">%s</a></td><td width="20%%">%s</td><td width="20%%">%s</td></tr>\n'
                    % (urllib.quote(linkname), colorName,
                        sizeof_fmt(os.path.getsize(filename)), modification_date(filename)))
        f.write("</table>\n<hr>\n*文件只能在/history/目录中上传，目前只支持zip格式和xml格式，同一组上传请对应相同的文件名。\n</body>\n</html>\n")
        length = f.tell()
        f.seek(0)
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.send_header("Content-Length", str(length))
        self.end_headers()
        return f
 
    def translate_path(self, path):
        path = path.split('?',1)[0]
        path = path.split('#',1)[0]
        path = posixpath.normpath(urllib.unquote(path))
        words = path.split('/')
        words = filter(None, words)
        path = os.getcwd()
        for word in words:
            drive, word = os.path.splitdrive(word)
            head, word = os.path.split(word)
            if word in (os.curdir, os.pardir): continue
            path = os.path.join(path, word)
        return path
 
    def copyfile(self, source, outputfile):
        shutil.copyfileobj(source, outputfile)
 
    def guess_type(self, path):
        base, ext = posixpath.splitext(path)
        if ext in self.extensions_map:
            return self.extensions_map[ext]
        ext = ext.lower()
        if ext in self.extensions_map:
            return self.extensions_map[ext]
        else:
            return self.extensions_map['']
 
    if not mimetypes.inited:
        mimetypes.init()
    extensions_map = mimetypes.types_map.copy()
    extensions_map.update({
        '': 'application/octet-stream', # Default
        '.py': 'text/plain',
        '.c': 'text/plain',
        '.h': 'text/plain',
        })
 
class ThreadingServer(ThreadingMixIn, BaseHTTPServer.HTTPServer):
    pass
 
if __name__ == '__main__':
    #单线程
    # srvr = BaseHTTPServer.HTTPServer(serveraddr, SimpleHTTPRequestHandler)
    #多线程
    srvr = ThreadingServer(serveraddr, SimpleHTTPRequestHandler)
    srvr.serve_forever()