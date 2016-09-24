(function () {
    'use strict';

    angular.module('homescreenApp').controller('SettingsCtrl', ['$stateParams', '$state', '$http', SettingsCtrl]);

    function SettingsCtrl($stateParams, $state, $http) {
        var vm = this;
        vm.modules = [];
        vm.open = $stateParams.tab || 'main';
        vm.openSettings = goto => { $state.go('settings', { tab: goto }); };

        vm.saveSetting = saveSetting;

        function saveSetting(moduleId, setting, value, form) {
            $http.put(['/api/v1/storage/module', moduleId, setting].join('/'), JSON.stringify(value)).then(result => {
                form.$setPristine();
            });
        }

        $http.get('/api/v1/modules/')
            .then(result => {
                vm.modules = result.data;
            });
    }

})();