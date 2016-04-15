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