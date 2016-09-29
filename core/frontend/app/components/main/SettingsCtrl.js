(function () {
    'use strict';

    angular.module('homescreenApp').controller('SettingsCtrl', ['$stateParams', '$state', '$http', SettingsCtrl]);

    function SettingsCtrl($stateParams, $state, $http) {
        var vm = this;
        vm.modules = [];
        vm.open = $stateParams.tab || 'main';
        vm.openSettings = goto => { $state.go('settings', { tab: goto }); };

        vm.saveSetting = saveSetting;

        function saveSetting(moduleId, setting, value, type, form) {
            if (type === 'time') {
                if (value instanceof Date) {
                    value = value.getHours() + ':' + value.getMinutes();
                }
            } else if (type === 'date') {
                if (value instanceof Date) {
                    value.setMinutes(value.getMinutes() - value.getTimezoneOffset());
                    value = value.toISOString().slice(0, 10);
                }
            }
            $http.put(['/api/v1/storage/module', moduleId, setting].join('/'), JSON.stringify(value)).then(result => {
                form.$setPristine();
            });
        }

        $http.get('/api/v1/modules/')
            .then(result => {
                result.data.forEach(mod => {
                    mod.settings.forEach(setting => {
                        if (setting.type === 'time') {
                            var d = new Date();
                            d.setHours(setting.value.split(':')[0], setting.value.split(':')[1]);
                            setting.value = d;
                        }
                        else if (setting.type === 'date') {
                            var d = new Date(setting.value);
                            setting.value = d;
                        }
                    });
                });
                vm.modules = result.data;
            });
    }

})();