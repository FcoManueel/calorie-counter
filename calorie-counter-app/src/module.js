var $module = angular.module('app', ['ui.router', 'ngCookies']);

$module.constant('API_ROUTE', 'http://127.0.0.1:8000');

$module.config(['$httpProvider', function ($httpProvider) {
    $httpProvider.interceptors.push('AuthInterceptor');
}]);
