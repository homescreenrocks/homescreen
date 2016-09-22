(function () {
    'use strict';

    angular.module('homescreenApp').controller('SettingsCtrl', ['$stateParams', '$state', SettingsCtrl]);

    function SettingsCtrl($stateParams, $state) {
        var vm = this;
        vm.open = $stateParams.tab || 'main';
        vm.openSettings = function (goto) { $state.go('settings', { tab: goto }); };
    }

})();