# Go-Sanitizer
Go-Sanitizer：Golang的面向错误的断言生成（ISSRE'19中论文的源代码）

Go-Sanitizer: Bug-Oriented Assertion Generation for Golang (source code for paper in ISSRE'19)

# 介绍 Introduction
Go编程语言（Golang）被广泛使用，并且由于其广泛的应用，安全问题变得越来越重要。大多数现有的验证技术，例如模糊测试和单元测试，主要集中在崩溃检测和覆盖范围改进上。但是，对于测试引擎而言，要理解常见的程序错误（例如精度损失和整数溢出）是一项挑战。

在本文中，我们提出了Go-Sanitizer，这是一种有效的面向Golang的面向错误的断言生成器，它能够在查找程序错误时实现更好的性能。首先，我们手动分析通用弱点枚举（CWE）并总结在Golang上的适用性。其次，我们设计了一个生成器，以将几个面向错误的断言自动插入到目标程序的正确位置。最后，我们可以利用传统的验证技术（例如模糊测试和单元测试）来测试带有已插入断言的程序，而Go-Sanitizer会通过断言失败来报告错误。

Go programming language (Golang) is widely used, and the security issue becomes increasingly important because of its extensive applications. Most existing validation techniques, such as fuzz testing and unit testing, mainly focus on crashes detection and coverage improvements. However, it is challenging for test engines to perceive common program bugs such as loss of precision and integer overflow. 

In this paper, we propose Go-Sanitizer, an effective bug-oriented assertion generator for Golang, which is able to achieve a better performance in finding program bugs. Firstly, we manually analyze the Common Weakness Enumeration (CWE) and summarize the applicabilities on Golang. Secondly, we design a generator to automatically insert several bug-oriented assertions to the proper locations of the target program. Finally, we can utilize the traditional validation techniques such as fuzz and unit testing to test the programs with inserted assertions, and Go-Sanitizer reports bugs via the failures of assertions.

# 用法 Usage
go get github.com/wangcong15/go-sanitizer & go-santizier -h

# 论文 Paper
http://wingtecher.com/themes/WingTecherResearch/assets/papers/issre19_go.pdf
