import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from '@angular/material';
import { FlexLayoutModule } from '@angular/flex-layout';
import 'hammerjs';

import { AppRoutingModule } from './app-routing.module';
import { ModuleService } from './shared/module.service';
import { AppComponent } from './app.component';
import { SettingsComponent } from './settings/settings.component';
import { DateValueAccessorModule } from 'angular-date-value-accessor';

@NgModule({
  declarations: [
    AppComponent,
    SettingsComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    DateValueAccessorModule,
    MaterialModule,
    FlexLayoutModule
  ],
  providers: [
    ModuleService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
