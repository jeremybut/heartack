
var app = angular.module('admin', ['auth', 'flash', 'ngAnimate']);

app.controller('AdminCtrl', ['$scope', '$http', 'Flash',
 function ($scope, $http, Flash) {

	$http.get('/api/users').success(function(result, status, headers) {
		$scope.users = result;
	}).error(function(data,status) {
		Flash.create('danger', data);
	});

}]);
