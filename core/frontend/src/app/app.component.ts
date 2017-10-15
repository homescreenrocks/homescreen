import { Component, OnInit, ViewChild } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { Router, NavigationEnd, ActivatedRoute, ActivatedRouteSnapshot } from '@angular/router';
import * as screenfull from 'screenfull';
import 'rxjs/add/operator/filter';
import { Subscription } from 'rxjs/Subscription';
import { MediaChange, ObservableMedia } from '@angular/flex-layout';
import { MatSidenav, MatSnackBar } from '@angular/material';
import * as domHelper from './shared/helpers/dom.helper';

import { RoutePartsService } from './shared/services/route-parts/route-parts.service';
import { ThemeService } from './shared/services/theme/theme.service';

@Component({
  selector: 'hs-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  fullscreen = screenfull;
  appTitle = 'Homescreen';
  pageTitle = '';
  private isMobile;
  screenSizeWatcher: Subscription;
  isSidenavOpen: Boolean = false;
  @ViewChild(MatSidenav) private sideNave: MatSidenav;

  constructor(public title: Title,
    private router: Router,
    private activeRoute: ActivatedRoute,
    private routePartsService: RoutePartsService,
    private ts: ThemeService,
    private media: ObservableMedia,
    public snackBar: MatSnackBar
  ) {
    // Close sidenav after route change in mobile
    router.events.subscribe((routeChange) => {
      if (routeChange instanceof NavigationEnd && this.isMobile) {
        this.sideNave.close();
      }
    });
    // Watches screen size and open/close sidenav
    this.screenSizeWatcher = media.subscribe((change: MediaChange) => {
      this.isMobile = (change.mqAlias === 'xs') || (change.mqAlias === 'sm');
      this.updateSidenav();
    });
  }

  ngOnInit() {
    this.changePageTitle();
    this.ts.getCurrentTheme();
  }

  changePageTitle() {
    this.router.events.filter(event => event instanceof NavigationEnd).subscribe((routeChange) => {
      const routeParts = this.routePartsService.generateRouteParts(this.activeRoute.snapshot);
      if (!routeParts.length) { return this.title.setTitle(this.appTitle); }
      // Extract title from parts;
      this.pageTitle = routeParts
                      .map((part) => part.title )
                      .reduce((partA, partI) => `${partA} > ${partI}`);
      this.pageTitle += ` | ${this.appTitle}`;
      this.title.setTitle(this.pageTitle);
    });
  }

  updateSidenav() {
    const self = this;
    setTimeout(() => {
      self.isSidenavOpen = !self.isMobile;
      self.sideNave.mode = self.isMobile ? 'over' : 'side';
      if (self.isMobile) { domHelper.removeClass(document.body, 'collapsed-menu'); }
    });
  }
}
