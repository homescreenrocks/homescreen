import { Component, OnInit } from '@angular/core';

import { ModuleService } from '../shared/services/module/module.service';
import { IModule, IModuleSetting } from '../shared/services/module/module-interface';

@Component({
  selector: 'hs-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent implements OnInit {
  private modules: Array<IModule>;

  constructor(
    private ms: ModuleService
  ) { }

  ngOnInit() {
    this.ms.getSettings()
      .subscribe(res => {
        console.log(res);
        this.modules = res;
      });
  }

  saveSettings(id:  number, settings: IModuleSetting[]) {
    console.log(id, settings);
    this.ms.setModuleValues(id, settings)
      .subscribe(res => res);
  }

}
