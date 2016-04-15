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