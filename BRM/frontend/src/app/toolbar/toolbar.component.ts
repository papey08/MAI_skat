import {Component, Inject, inject, Input, OnDestroy, OnInit, Renderer2,} from '@angular/core';
import {MatIconModule} from '@angular/material/icon';
import {AuthService} from '../services/auth.service';
import {Router} from '@angular/router';
import {MatButtonModule} from '@angular/material/button';
import {Subscription} from 'rxjs';
import {DalService} from '../DAL/core/dal.service';
import {DOCUMENT} from '@angular/common';
import {FormControl} from '@angular/forms';

@Component({
  selector: 'app-toolbar',
  standalone: true,
  imports: [MatIconModule, MatButtonModule],
  templateUrl: './toolbar.component.html',
  styleUrl: './toolbar.component.scss',
})
export class ToolbarComponent implements OnInit, OnDestroy {
  @Input() snav: any;

  authService = inject(AuthService);
  dalService = inject(DalService);

  subscription: Subscription = new Subscription();
  router = inject(Router);
  companyName: string = '';

  switchTheme = new FormControl(false, {nonNullable: true});
  darkClass = 'theme-dark';
  lightClass = 'theme-light';

  constructor(
    @Inject(DOCUMENT) private document: Document,
    private renderer: Renderer2
  ) {
  }

  ngOnInit(): void {
    this.subscription.add(
      this.dalService
        .getCompanyById(+this.authService.currentUserDataSig()?.['company-id']!)
        .subscribe((value) => (this.companyName = value.data.name!))
    );

    this.switchTheme.setValue(
      this.document.body.classList.value !== this.lightClass
    );
    this.switchTheme.valueChanges.subscribe((currentMode) => {
      localStorage.setItem('currentMode', `${currentMode}`);
      const className = currentMode ? this.darkClass : this.lightClass;
      this.renderer.setAttribute(this.document.body, 'class', className);
    });
  }

  ngOnDestroy() {
    this.subscription.unsubscribe();
  }

  logout(): void {
    localStorage.setItem('token', '');
    this.authService.currentUserSig.set(null);
    this.authService.currentUserDataSig.set(null);
    this.snav.close();
    this.router.navigateByUrl('/login');
  }

  redirectToNotifications() {
    this.router.navigateByUrl('/notifications');
  }

  redirectToMainPage() {
    this.router.navigateByUrl('/main-page');
  }

  changeMode() {
    this.switchTheme.setValue(!this.switchTheme.value);
  }
}
