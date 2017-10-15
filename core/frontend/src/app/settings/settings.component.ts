import { Component, OnInit } from '@angular/core';

import { ModuleService } from '../services/module/module.service';
import { IModule } from '../services/module/module-interface';

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
        this.modules = res;
      });
  }

}
