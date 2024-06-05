import {
  ChangeDetectorRef,
  Component,
  Inject,
  inject,
  OnChanges,
  OnDestroy,
  OnInit,
  Renderer2,
  SimpleChanges,
  ViewChild,
} from '@angular/core';
import {CommonModule, DOCUMENT} from '@angular/common';
import {Router, RouterOutlet} from '@angular/router';
import {MediaMatcher} from '@angular/cdk/layout';
import {AuthService} from './services/auth.service';
import {MatButtonModule} from '@angular/material/button';
import {MatIconModule} from '@angular/material/icon';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatToolbarModule} from '@angular/material/toolbar';
import {LeftMenuComponent} from './left-menu/left-menu.component';
import {LoginComponent} from './login/login.component';
import {jwtDecode} from 'jwt-decode';
import {DalService} from './DAL/core/dal.service';
import {Subscription} from 'rxjs';
import {ToolbarComponent} from './toolbar/toolbar.component';
import {MatSnackBar} from '@angular/material/snack-bar';
import {HelloSnackbarComponent} from "./hello-snackbar/hello-snackbar.component";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,
    RouterOutlet,
    MatSidenavModule,
    LeftMenuComponent,
    MatButtonModule,
    MatIconModule,
    MatToolbarModule,
    LoginComponent,
    ToolbarComponent,
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent implements OnInit, OnChanges, OnDestroy {
  @ViewChild('snav') snav: any;

  subscription = new Subscription();

  title = 'BRM';
  mobileQuery: MediaQueryList;
  router = inject(Router);

  authService = inject(AuthService);
  dalService = inject(DalService);

  companyName: string = '';

  private _mobileQueryListener: () => void;

  constructor(
    changeDetectorRef: ChangeDetectorRef,
    media: MediaMatcher,
    @Inject(DOCUMENT) private document: Document,
    private renderer: Renderer2,
    private _snackBar: MatSnackBar
  ) {
    this.mobileQuery = media.matchMedia('(max-width: 600px)');
    this._mobileQueryListener = () => changeDetectorRef.detectChanges();
    this.mobileQuery.addListener(this._mobileQueryListener);

    this._snackBar.openFromComponent(HelloSnackbarComponent, {
      duration: 5000,
    });

    const token = localStorage.getItem('token');

    if (token && token != '') {
      this.authService.currentUserSig.set({access: token, refresh: ''});
      this.authService.currentUserDataSig.set(jwtDecode(token));
    } else {
      this.router.navigateByUrl('/login');
    }
  }

  ngOnInit(): void {
    this.setTheme();
  }

  ngOnChanges(changes: SimpleChanges) {
    this.subscription.add(
      this.dalService
        .getCompanyById(+this.authService.currentUserDataSig()?.['company-id']!)
        .subscribe((value) => (this.companyName = value.data.name!))
    );
  }

  ngOnDestroy(): void {
    this.mobileQuery.removeListener(this._mobileQueryListener);
  }

  setTheme() {
    let darkClass = 'theme-dark';
    let lightClass = 'theme-light';
    let currentMode =
      localStorage.getItem('currentMode') === 'true' ? true : false;
    const className = currentMode ? darkClass : lightClass;
    this.renderer.setAttribute(this.document.body, 'class', className);
  }
}
