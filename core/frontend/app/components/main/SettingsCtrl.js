(function () {
    'use strict';

    angular.module('homescreenApp').controller('SettingsCtrl', ['$stateParams', '$state', '$http', SettingsCtrl]);

    function SettingsCtrl($stateParams, $state, $http) {
        var vm = this;
        vm.modules = [];
        vm.open = $stateParams.tab || 'main';
        vm.openSettings = goto => { $state.go('settings', { tab: goto }); };

        $http.get('/api/v1/modules/list')
            .then(result => {
                vm.modules = result.data;
            });
    }

})();