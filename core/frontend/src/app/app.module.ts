import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
  MdInputModule,
  MdCardModule,
  MdToolbarModule,
  MdButtonModule,
  MdDatepickerModule,
  MdNativeDateModule,
  MdFormFieldModule,
  MdOptionModule,
  MdMenuModule,
  MdRadioModule,
  MdTooltipModule,
  MdCheckboxModule,
  MdSelectModule
} from '@angular/material';
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
    MdInputModule,
    MdCardModule,
    MdToolbarModule,
    MdButtonModule,
    MdDatepickerModule,
    MdNativeDateModule,
    MdFormFieldModule,
    MdOptionModule,
    MdMenuModule,
    MdRadioModule,
    MdTooltipModule,
    FlexLayoutModule,
    MdCheckboxModule,
    MdSelectModule
  ],
  providers: [
    ModuleService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
