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