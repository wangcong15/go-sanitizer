# Go-Sanitizer
Go-Sanitizer: Bug-Oriented Assertion Generation for Golang (source code for our paper in ISSRE'19)

# Abstract
Go programming language (Golang) is widely used, and the security issue becomes increasingly important because of its extensive applications. Most existing validation techniques, such as fuzz testing and unit testing, mainly focus on crashes detection and coverage improvements. However, it is challenging for test engines to perceive common program bugs such as loss of precision and integer overflow. 

In this paper, we propose Go-Sanitizer, an effective bug-oriented assertion  generator for Golang, which is able to achieve a better performance in finding program bugs. Firstly, we manually analyze the Common Weakness Enumeration (CWE) and summarize the applicabilities on Golang. Secondly, we design a generator to automatically insert several bug-oriented assertions to the proper locations of the target program. Finally, we can utilize the traditional validation techniques such as fuzz and unit testing to test the programs with inserted assertions, and Go-Sanitizer reports bugs via the failures of assertions. For evaluation, we apply Go-Sanitizer to Badger, a widely-used database software, and successfully discovers 12 previously unreported program bugs, which can not be detected by pure fuzzer such as Go-Fuzz or unit testing methods.

# Usage
go get github.com/wangcong15/go-sanitizer
go-santizier -h

# Paper
http://wingtecher.com/themes/WingTecherResearch/assets/papers/issre19_go.pdf
