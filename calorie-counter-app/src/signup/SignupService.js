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