import os
import sys
import json

class variable(object):
    def __init__(self, header):
        self.header = header
        self.front = list()
        self.statements = list()

    def addFront(self, s):
        self.front.append(s)    

    def addStatement(self, s):
        self.statements.append(s)

    def toString(self):
        return "%s{\n\t%s\n\t%s\n}" % (self.header, "\n\t".join(self.front), "\n\t".join(self.statements))

FUNCNAME = ""
BUGTYPE = ""
ASSERT_LINENO = ""
ASSERT_PARAMS = ""

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("A json file needed")
        os._exit(1)
    with open(sys.argv[1], "r") as f:
        cfg = json.load(f)    
    
    # Handle global messages
    FUNCNAME = cfg.get("funcName")
    BUGTYPE = cfg.get("bugType")
    ASSERT_LINENO = cfg.get("genAssert", {}).get("lineNo")
    ASSERT_PARAMS = cfg.get("genAssert", {}).get("params")

    # Handle the configuration
    headers = "package checkers\nimport (\n\t\"go/ast\"\n\t\"github.com/wangcong15/go-sanitizer/code\"\n)"
    localValStr = "type localVars%s struct {\n%s\n\tfuncName string\n\tparams   string\n\tlineNo   int\n\tbugType  string}\n}"
    localVals = ""
    
    funcMap = dict()
    funcN = "c.F"
    header = "func check%s(c *code.Code)" % (BUGTYPE)
    funcMap[funcN] = variable(header)
    funcMap[funcN].addFront("lv := localVars%s{\n\t\tfuncName: \"%s\",\n\t\tbugType:  \"%s\",\n\t}" % (BUGTYPE, FUNCNAME, BUGTYPE))
    for v in cfg.get("localVars"):
        n = v.get("name")
        t = v.get("type")
        rt = v.get("rule", {}).get("type")
        r = v.get("rule", {}).get("root")
        a = v.get("rule", {}).get("attr")
        f = v.get("filter")
        c = v.get("callback")
        i = v.get("init")
        e = v.get("extra")
        localVals += "\t%s %s\n" % (n, t)
        if not e:
            funcN = "func check%s%s(c *code.Code, %s %s, lv *localVars%s)" % (BUGTYPE, n, n, t, BUGTYPE)
            funcMap[n] = variable(funcN)
            funcMap[n].addFront("lv.%s = %s" % (n, n))
        # 
        if rt == "inspect":
            s = "%sList := Inspect(%s, &%s{})\n\tfor _, %s := range %sList {\n\t\tcheck%s%s(c, %s.(%s), %slv)\n\t}" % (n, r, t[1:], n, n, BUGTYPE, n, n, t, "&" if n=="c.F" else "")
            funcMap[r].addStatement(s)
        # 
        elif rt == "element":
            if f:
                s = "if %s {\n\t\t%s := %s\n\t\tif _, ok := %s.(%s); ok {\n\t\t\tcheck%s%s(c, %s.(%s), lv)\n\t\t}\n\t}" % (f, n, "%s.%s" % (r, a) if not a.startswith("lv") else a, n, t, BUGTYPE, n, n, t)
            else:
                s = "%s := %s.%s\n\tif _, ok := %s; ok {\n\t\tcheck%s%s(c, %s.(%s), lv)\n\t}" % (n, "%s.%s" % (r, a) if not a.startswith("lv") else a, n, t, BUGTYPE, n, n, t)
            funcMap[r].addStatement(s)
            if c:
                funcMap[n].addStatement(c)
        # 
        elif rt == "foreach":
            s = "for _, %s := range %s {\n\t\tif _, ok := %s.(%s); ok {\n\t\t\tcheck%s%s(c, %s.(%s), lv)\n\t\t}\n\t}" % (n, r, n, t, BUGTYPE, n, n, t)
            funcMap[r].addStatement(s)
            if c:
                funcMap[n].addStatement(c)
        # 
        elif e:
            funcMap["c.F"].addFront(i)

    localValStr = localValStr % (BUGTYPE, localVals)
    finalCode = "%s\n%s\n%s" % (headers, localValStr, "\n".join([funcMap[k].toString() for k in funcMap]))
    print(finalCode)