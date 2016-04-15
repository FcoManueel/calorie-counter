(function(undefined){

var $module = angular.module('app', ['ui.router', 'ngCookies']);

$module.constant('API_ROUTE', 'http://127.0.0.1:8000');

$module.config(['$httpProvider', function ($httpProvider) {
    $httpProvider.interceptors.push('AuthInterceptor');
}]);

$module.config(['$stateProvider',
	function($stateProvider) {
		var states = {
			'index': {
				url: '/index',
				templateUrl: '/home.html',
				controller: 'HomeController'
			},
			'login': {
				url: '/login',
				templateUrl: '/login.html',
				controller: 'LoginController'
			},
			'signup': {
				url: '/signup',
				templateUrl: '/signup.html',
				controller: 'SignupController'
			},
			'logout': {
				url: '/logout',
				templateUrl: '/logout.html',
				controller: 'LogoutController'
			},
			'dashboard': {
				url: '/dashboard',
				templateUrl: '/dashboard.html',
				controller: 'DashboardController'
			}
		};

		angular.forEach(states, function(config, name) {
			$stateProvider.state(name, config);
		});
	}
]);
$module.factory('AuthInterceptor', ['TokenService',
    function (TokenService) {
        return {
            request: function (config) {
                var accessToken = TokenService.getAccessToken();

                if (accessToken) {
                    config.headers['Authorization'] = 'Bearer ' + accessToken;
                }

                return config;
            }
        };
    }
]);
$module.factory('TokenService', ['$cookieStore',
    function ($cookieStore) {
        var accessToken = "";

        return {
            getAccessToken: function () {
                return accessToken;
            },

            setAccessToken: function (token) {
                $cookieStore.put('x-access-token', token);

                accessToken = token;

                return this;
            },

            loadFromCookie: function () {
                accessToken = $cookieStore.get('x-access-token');
                return this;
            },

            deleteAccessToken: function() {
                $cookieStore.remove('x-access-token');
                return this;
            },

            checkIsAuthenticated: function () {
                return accessToken !== "";
            },
        };
    }
]);

$module.controller('DashboardController', ['$scope', 'IntakesService', 'TokenService', '$state', 'UserService',
    function ($scope, IntakesService, TokenService, $state, UserService) {
        /*
         * Attrs
         */
        $scope.rangeDate = {
            from: "",
            to: ""
        };

        $scope.rangeTime = {
            from: "",
            to: ""
        };

        $scope.moment = moment;

        $scope.intakes = [];

        $scope.newIntake = {
            name: '',
            calories: '',
            time: '',
            date: ''
        };

        $scope.todayCalories = 0;

        $scope.todayIntakes = [];

        $scope.user = {
            id: '',
            email: '',
            role: '',
            name: '',
            goalCalories: 0
        };

        /*
         * Methods
         */

        $scope.init = function init() {
            var accessToken = TokenService.getAccessToken();

            if (!accessToken) {
                $state.go('login');
                return;
            }

            $scope.getUser();
            $scope.getIntakes();
        };

        $scope.getIntakes = function getIntakes() {
            IntakesService.get().then(function (allIntakes) {
                $scope.intakes = allIntakes;

                $scope.todayIntakes = IntakesService.filter(allIntakes, function (intake) {
                    return moment().diff(intake.consumedAt, 'days') === 0;
                });

                $scope.todayCalories = calculateTodayCalories($scope.todayIntakes);
            });
        };

        $scope.getUser = function getUser() {
            UserService.get().then(function (user) {
                $scope.user = user;
            });
        };

        $scope.onAddNewIntake = function onAddNewIntake(newIntake) {

            IntakesService.post(newIntake.name, newIntake.calories, moment()).then(function () {
                $scope.getIntakes();

                $scope.newIntake = {
                    name: '',
                    calories: ''
                };
            });
        };

        function calculateTodayCalories(intakes) {
            var sum = 0;

            intakes.forEach(function (intake) {
                sum += intake.calories;
            });

            return sum;
        }

        $scope.dateFilter = function dateFilter(from, to){
            return function(value, index, array){
                var created = moment(value.createdAt);
                var shouldShow = true;
                if (from !== "" && created.isBefore(from, 'day')) {
                    shouldShow = false
                }
                if (to !== "" && created.isAfter(to, 'day')){
                    shouldShow = false
                }
                return shouldShow;
            }
        };

        $scope.timeFilter = function timeFilter(from, to){
            toHours = function(m){
                return m.minutes() + m.hours()*60
            };

            return function(value, index, array){
                var createdAt = moment(value.createdAt).format("HH:mm");
                var shouldShow = true;
                if (from !== "" && createdAt < from) {
                    console.log("is before");
                    shouldShow = false
                }
                if (to !== "" && createdAt > to){
                    console.log("is after");
                    shouldShow = false
                }
                return shouldShow;
            }
        };
        $scope.formatDate = function(date){
            return moment().diff(date, 'hours') < 24 ? moment(date).startOf('minute').fromNow() : moment(date).format('YYYY-MM-DD [at] hh:mm A')
        }
    }
]);
$module.factory('FooService', ['$http',
	function($http) {
		return {
			getFoo: function() {
				return $http.get('/foo').then(function(response) {
					return response.data;
				});
			}
		};
	}
]);
$module.controller('HomeController', ['$scope',
	function($scope) {
		$scope.awesome = true;
	}
]);
$module.factory('IntakesService', ['$http', 'API_ROUTE',
    function ($http, API_ROUTE) {
        return {
            get: function get() {
                return $http.get(API_ROUTE + '/v1/intakes').then(function (chunk) {
                    return chunk.data.intakes;
                });
            },

            post: function post(intakeName, intakeCalories, intakeConsumedAt) {
                var data = {
                    name: intakeName,
                    calories: parseInt(intakeCalories, 10),
                    consumedAt: intakeConsumedAt
                };

                return $http.post(API_ROUTE + '/v1/intakes', data);
            },

            filter: function filter(intakes, cb) {
                return intakes.filter(cb || fnK);
            }
        };

        function fnK(a) {
            return a;
        }
    }
]);
$module.controller('LoginController', ['$scope', '$state', 'LoginService', 'TokenService',
    function ($scope, $state, LoginService, TokenService) {
        /*
         * Attrs
         */

        $scope.user = {
            username: '',
            passwd: ''
        };

        /*
         * Methods
         */

        $scope.init = function init() {
            var accessToken = TokenService.loadFromCookie().getAccessToken();
            if (accessToken) {
                $state.go('dashboard');
            }
        };

        $scope.onLogin = function onLogin(user) {
            LoginService.postLogin(user.username, user.passwd)
                .then(function () {
                    $state.go('dashboard');
                    $scope.user = {
                        name: '',
                        passwd: '',
                    };
                });
        };
    }
]);


$module.factory('LoginService', ['$http', 'API_ROUTE', 'TokenService',
    function ($http, API_ROUTE, TokenService) {
        return {
            postLogin: function (username, passwd) {
                var data = {
                    email: username,
                    password: passwd
                };

                return $http.post(API_ROUTE + '/auth/login', data)
                    .then(function (chunk) {
                        TokenService.setAccessToken(chunk.data.accessToken);
                    });
            }
        };
    }
]);
$module.controller('LogoutController', ['$scope', 'TokenService', 'LogoutService',
    function ($scope, TokenService, LogoutService){
        $scope.onLogout = LogoutService.onLogout;
    }
]);


$module.factory('LogoutService', ['$http', 'API_ROUTE', 'TokenService', '$state', '$rootScope',
    function ($http, API_ROUTE, TokenService, $state, $rootScope) {
        return {
            onLogout: function () {
                TokenService.deleteAccessToken();
                $state.go('index');
                if(!$scope.$$phase) {
                    $rootScope.$digest();
                }
            }
        };
    }
]);
$module.controller('PageController', ['$scope', 'TokenService', '$state', 'LogoutService',
    function ($scope, TokenService, $state, LogoutService) {
        var isAuthenticated;

        $scope.isAuthenticated = function isAuthenticated(){
                return TokenService.checkIsAuthenticated();
        };

        $scope.signup = function signup(){
            console.log("signup");
            $state.go('signup');
        };

        $scope.login = function login(){
            console.log("login");
            $state.go('login');
        };

        $scope.settings = function settings(){
            console.log("settings");
            $state.go('settings');
        };

        $scope.logout = function logout(){
            console.log("logout");
            LogoutService.onLogout();
        };
    }]
);
$module.controller('SignupController', ['$scope', '$state', 'SignupService', 'TokenService',
    function ($scope, $state, SignupService, TokenService) {
        /*
         * Attrs
         */

        $scope.user = {
            username: '',
            name: '',
            passwd: '',
            goalCalories: ''
        };

        /*
         * Methods
         */

        $scope.init = function init() {
            var accessToken = TokenService.loadFromCookie().getAccessToken();
            if (accessToken) {
                $state.go('dashboard');
            }
        };

        $scope.onSignup = function onLogin(user) {
            SignupService.postSignup(user.username, user.name, user.passwd, user.goalCalories)
                .then(function () {
                    $scope.user = {
                        username: '',
                        name: '',
                        passwd: '',
                        goalCalories: ''
                    };
                    $state.go('dashboard');
                });
        };
    }
]);


$module.factory('SignupService', ['$http', 'API_ROUTE', 'TokenService',
    function ($http, API_ROUTE, TokenService) {
        return {
            postSignup: function (username, name, passwd, goalCalories) {
                var data = {
                    email: username,
                    name: name,
                    password: passwd,
                    goalCalories: parseInt(goalCalories, 10)
                };

                return $http.post(API_ROUTE + '/auth/signup', data)
                    .then(function (chunk) {
                        TokenService.setAccessToken(chunk.data.accessToken);
                    });
            }
        };
    }
]);
$module.factory('UserService', ['$http', 'API_ROUTE',
    function ($http, API_ROUTE) {
        return {
            get: function() {
                return $http.get(API_ROUTE + '/v1/users')
                    .then(function (chunk) {
                        return chunk.data;
                    });
            }
        };
    }
]);
}());