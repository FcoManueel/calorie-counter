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