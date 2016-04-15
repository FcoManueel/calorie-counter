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