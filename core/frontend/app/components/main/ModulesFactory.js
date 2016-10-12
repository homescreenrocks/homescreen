(function () {
    'use strict';

    angular.module('homescreenApp').factory('ModulesFactory', ['$http', '$q', ModulesFactory]);

    function ModulesFactory($http, $q) {
        return {
            getModules: () => {
                var deferred = $q.defer();
                $http.get('/api/v1/modules/').then(
                    resultOkay => {
                        deferred.resolve(resultOkay.data);
                    },
                    resultFailure => {
                        deferred.reject(resultFailure);
                    }
                );
                return deferred.promise;
            }
        };

    }

})();