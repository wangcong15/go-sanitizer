var app = angular.module('myApp', []);
app.controller('myCtrl',
function($scope, $http) {
	$http({
		method: 'GET',
		url: '/history/'
	}).then(function successCallback(response) {
		// 收到结果
		parseHtml(response.data);
	},
	function errorCallback(response) {
		// 请求失败执行代码
		console.log("Get Fails");
	});

	// 代码编辑器的初始化
	$scope.editor = ace.edit('editor');
	$scope.editor.getSession().setMode('ace/mode/golang');
	$scope.editor.setTheme('ace/theme/monokai');
	$scope.editor.setHighlightActiveLine(true);
	$scope.editor.$blockScrolling = Infinity;
	$scope.editor.setReadOnly(true);
	document.getElementById('editor').style.fontSize = '16px';
	
	// 变量初始化
	$scope.is_full = 0;
	$scope.history_list = []; // 所有历史提交的名称
	$scope.hlIndex = -1; // 当前选定的历史提交
	$scope.curr_route = []; // 当前项目路径
	$scope.bugs = []; // 当前的BUG
	$scope.currFile = [];
	$scope.currLine = 0;
	$scope.orderCond = '';
	$scope.orderCondSelect = '';
	$scope.currTheme = 'monokai';
	$scope.currCodeTheme = 'monokai'; 
	$scope.searchword = '';
	$scope.fromLargeToSmall = 1;
	$scope.showAllFiles = true;
	$scope.ctCode = [];
	$scope.openFilenames = [];
	$scope.codeDic = {}; // 代码映射字典，Map<filepath, [filename, code]>
	$scope.currTab = '';
	$scope.fontsizes = [10, 11, 12, 13, 14, 15, 16, 17, 18, 19];
	$scope.themes = ["ambiance", "gruvbox", "sqlserver", "chaos", "idle_fingers", "terminal", "chrome", "iplastic", "textmate", "clouds", "katzenmilch", "tomorrow", "clouds_midnight", "kr_theme", "tomorrow_night", "cobalt", "kuroir", "tomorrow_night_blue", "crimson_editor", "merbivore", "tomorrow_night_bright", "dawn", "merbivore_soft", "tomorrow_night_eighties", "dracula", "mono_industrial", "twilight", "dreamweaver", "monokai", "vibrant_ink", "eclipse", "pastel_on_dark", "xcode", "github", "solarized_dark", "gob", "solarized_light"];

	$scope.currFs = 16; // 当前选择字体大小
	$scope.currCodeFs = 16;
	$scope.fullscreenIcon = 'lib/full.png';
	$scope.cwe_desc = {
		"128": "Wrap around errors occur whenever a value is incremented past the maximum value for its type and therefore \"wraps around\" to a very small, negative, or undefined value.",
		"190": "Integer Overflow. The software performs a calculation that can produce an integer overflow or wraparound, when the logic assumes that the resulting value will always be larger than the original value. This can introduce other weaknesses when the calculation is used for resource management or execution control.",
		"191": "Integer Underflow. The product subtracts one value from another, such that the result is less than the minimum allowable integer value, which produces a value that is not equal to the correct result.",
		"466": "A function can return a pointer to memory that is outside of the buffer that the pointer is expected to reference.",
		"478": "The code does not have a default case in a switch statement, which might lead to complex logical errors and resultant weaknesses.",
		"777": "The software uses a regular expression to perform neutralization, but the regular expression is not anchored and may allow malicious or malformed data to slip through.",
		"785": "The software invokes a function for normalizing paths or file names, but it provides an output buffer that is smaller than the maximum possible size, such as PATH_MAX.",
		"824": "The program accesses or uses a pointer that has not been initialized.",
		"1077": "The code performs a comparison such as an equality test between two float (floating point) values, but it uses comparison operators that do not account for the possibility of loss of precision."
	};

	$scope.changeHL = function(hlIndex) {
		if (hlIndex < 0) return;
		// console.log(hlIndex);
		$scope.hlIndex = hlIndex;
		$scope.curr_route = [];
		$scope.bugs = [];
		$scope.curr_route.push($scope.history_list[hlIndex]);
		getProj($scope.history_list[hlIndex]);
	};

	$scope.changeOrderCond = function(newCond) {
		if ($scope.orderCond == newCond) {
			$scope.orderCond = "-" + $scope.orderCond;
		} else {
			$scope.orderCond = newCond;
		}
		$scope.orderCondSelect = newCond;
	};

	$scope.bugClick = function(bug) {
		for (var i = 0; i < $scope.bugs.length; i++) $scope.bugs[i].isSelected = 0;
		bug.isSelected = 1;
		$scope.currLine = bug["lineNo"];

		var file_route = "/history/" + $scope.history_list[$scope.hlIndex] + "/" + bug["filepath"];
		$http({
			method: 'GET',
			url: file_route
		}).then(function successCallback(response) {
			// 收到结果
			$scope.editor.setValue(response.data);
			$scope.editor.gotoLine($scope.currLine-1);
			$scope.editor.navigateTo($scope.currLine-1, 0);
		},
		function errorCallback(response) {
			// 请求失败执行代码
			console.log("Get Fails");
		});
		var treeObj = $.fn.zTree.getZTreeObj("tree");
		var tmpNodes = treeObj.getNodes()[0].children;
		var fileName = bug["filepath"];
		var TId = "";
		for (var i = 0; i < $scope.currFile.length - 1; i++) {
			for (var j = 0; j < tmpNodes.length; j++) {
				if (tmpNodes[j].name == $scope.currFile[i] && tmpNodes[j].isParent) {
					tmpNodes = tmpNodes[j].children;
					break;
				}
			}
		}
		for (var j = 0; j < tmpNodes.length; j++) {
			if (tmpNodes[j].name == fileName && !tmpNodes[j].isParent) {
				TId = tmpNodes[j].tId;
				break;
			}
		}
		var node = treeObj.getNodeByTId(TId);
		if (node) {
			treeObj.selectNode(node);
		}
	};

	$scope.cweclick = function(cweid, desc) {
		$scope.cwe_id = cweid;
		$scope.showCWE = 1;
	};

	$scope.selectFs = function(fsid) {
		$scope.currFs = $scope.fontsizes[fsid];
	};

	$scope.toggleFullScreen = function() {
		if ($scope.is_full == 0) {
			// 平移左右的模块
			$('#right-panel').hide();
			$('#left-panel').animate({
				'width': '0%'
			},
			500);
			$('.left-panel-content').hide();
			$('#middle-panel').animate({
				'width': '100%'
			},
			500);
			$scope.is_full = 1;
			$scope.fullscreenIcon = 'lib/full_exit.png';
		} else {
			// 缩回来
			$('#middle-panel').animate({
				'width': '50%'
			},
			300,
			function() {
				$('#right-panel').fadeIn(200);
				$('.left-panel-content').fadeIn(200);
			});
			$('#left-panel').animate({
				'width': '16.666667%'
			},
			500);
			$scope.is_full = 0;
			$scope.fullscreenIcon = 'lib/full.png';
		}
	};

	$scope.openSetting = function() {
		$('.modal-unity').fadeIn(500);
	};

	$scope.closeSetting = function() {
		$('.modal-unity').fadeOut(500,
		function() {
			$scope.currFs = $scope.currCodeFs;
			$scope.currTheme = $scope.currCodeTheme;
		});
	};

	$scope.saveSetting = function() {
		$scope.currCodeFs = $scope.currFs;
		document.getElementById('editor').style.fontSize = $scope.currCodeFs + 'px';
		$scope.editor.setTheme('ace/theme/'+ $scope.currTheme);
		$('.modal-unity').fadeOut(500);
	};

	$scope.selectTheme = function() {
		$scope.currCodeTheme = $scope.currTheme;
	};

	$scope.changeTab = function(filename){
		$scope.currTab = filename;
		var code_data = $scope.codeDic[filename][1];
		$scope.editor.setValue(code_data);
		$scope.editor.navigateTo(0, 0);
	};

	$scope.openfile = function(filepath, filename, codedata){
		$scope.currTab = filepath;
		if($.inArray(filepath, $scope.openFilenames) < 0){
			$scope.openFilenames.push(filepath);
			$scope.codeDic[filepath] = [filename, codedata];
		}
	};

	$scope.closeTab = function(filepath){
		var tempOpenFilenames = [];
		for(var i = 0; i < $scope.openFilenames.length; i++){
			if ($scope.openFilenames[i] != filepath){
				tempOpenFilenames.push($scope.openFilenames[i]);
			}
		}
		$scope.openFilenames = tempOpenFilenames;
		if ($scope.currTab){
			if ($scope.openFilenames.length > 0){
				$scope.openfile($scope.openFilenames[0]);
			}
			else{
				$scope.editor.setValue("");
			}
		}
	};

	// 文件树
	var zTreeOnClick = function(event, treeId, treeNode) {
		var currNode = treeNode;
		if (!treeNode.isParent) {
			var root_arr = [];
			for (var i = treeNode.level; i >= 1; i--) {
				root_arr.push(treeNode.name);
				treeNode = treeNode.getParentNode();
			}
			var file_route = "/history/" + $scope.history_list[$scope.hlIndex];
			while (root_arr.length > 0) {
				file_route += "/" + root_arr.pop();
			}
			$http({
				method: 'GET',
				url: file_route
			}).then(function successCallback(response) {
				// 收到结果
				$scope.editor.setValue(response.data);
				$scope.editor.navigateTo(0, 0);
				$scope.openfile(file_route, currNode.name, response.data);
			},
			function errorCallback(response) {
				// 请求失败执行代码
				console.log("Get Fails");
			});
		}
	};

	// 代码行获取
	var getCodeLine = function(file_name, line_number, index_number) {
		var file_route = "/history/" + $scope.history_list[$scope.hlIndex] + "/" + file_name;
		$http({
			method: 'GET',
			url: file_route
		}).then(function(response) {
			// 收到结果
			var resp = response.data;
			line_code = $.trim(resp.split('\n')[line_number - 1]);
			if (line_code.length > 30) {
				line_code = line_code.substring(0, 30) + "...";
			}
			console.log(line_code);
			$scope.ctCode[index_number] = line_code;
		});
	};

	var getProj = function(hl_name) {
		var proj_route = "/history/" + hl_name + "/proj_data.json";
		$http({
			method: 'GET',
			url: proj_route
		}).then(function successCallback(response) {
			// 收到结果
			parseProj(response.data);
		},
		function errorCallback(response) {
			// 请求失败执行代码
			console.log("Get Fails");
		});

		var proj_route = "/history/" + hl_name + "/.gosan";
		$http({
			method: 'GET',
			url: proj_route
		}).then(function successCallback(response) {
			// 收到结果
			parseBug(response.data);
		},
		function errorCallback(response) {
			// 请求失败执行代码
			console.log("Get Fails");
		});
	};

	// 处理获得的history文件夹下的结果
	var parseHtml = function(data) {
		// 获得项目文件夹列表
		var reFolder = new RegExp('<span style="background-color: #CEFFCE;">(.*?)/</span>');
		var folder_split = data.split(reFolder);
		var folder_list = [];
		for (var i = 0; i < folder_split.length; i++) {
			if (i % 2 == 1) $scope.history_list.push(folder_split[i]);
		}
	};

	// node排序的比较函数
	var compareNodeName = function (x, y) {
		if (('isParent' in x) && !('isParent' in y)) {
			return -1;
		} 
		else if (('isParent' in y) && !('isParent' in x)) {
			return 1;
		}
		else {
			if (x['name'] > y['name']){
				return 1;
			}
			return -1;
		}
	};

	// 过滤文件树，只保留Go文件
	var filterfile = function(data, showAllFiles) {
		var result = [];
		for (var i = 0; i < data.length; i++) {
			if ('children' in data[i]) {
				var tempChildren = filterfile(data[i]['children'], showAllFiles);
				var tempResult = data[i];
				tempResult['children'] = tempChildren.sort(compareNodeName);
				result.push(tempResult);
			} 
			else if (data[i]['name'].substring(data[i]['name'].length - 6) == ".gosan") {
				return result
			}
			else if (showAllFiles || ([".go"].indexOf(data[i]['name'].substring(data[i]['name'].length - 3)) >= 0)) {
				result.push(data[i]);
			}
		}
		return result;
	};

	// 处理获得路径文件夹下的结果
	var parseProj = function(data) {
		var setting = {
			view: {
				selectedMulti: false
			},
			data: {
				simpleData: {
					enable: true
				}
			},
			callback: {
				onClick: zTreeOnClick
			}
		};
		$scope.zNodes = filterfile(data, $scope.showAllFiles);
		// if ($scope.showAllFiles) {
		// 	$scope.zNodes = data;
		// } else {
		// 	$scope.zNodes = filterfile(data);
		// }
		$.fn.zTree.init($("#tree"), setting, $scope.zNodes);
	};

	// 处理获得路径文件夹下的xml的结果
	var parseBug = function(data) {
		var assert_list = data.split("\n")
		for (var i = 0; i < assert_list.length; i++) {
			if (assert_list[i] == "") {
				continue;
			}
			var msg = assert_list[i].split("\t")
			$scope.bugs.push({
				"cweid": [msg[0]],
				"assertion": msg[2],
				"lineNo": parseInt(msg[1]),
				"filepath": msg[3]
			})
		}
		for (var i = 0; i < $scope.bugs.length; i++) {
			$scope.bugs[i]['temp_id'] = i + 1;
		}
	}
});