import {Routes} from '@angular/router';
import {ContactsComponent} from './contacts/contacts.component';
import {LoginComponent} from './login/login.component';
import {authGuard} from './guards/auth.guard';
import {RegisterComponent} from './register/register.component';
import {MainPageComponent} from './main_page/main_page.component';
import {AdsComponent} from './ads/ads.component';
import {LeadsComponent} from './leads/leads.component';
import {EmployeesComponent} from './employees/employees.component';
import {NotificationComponent} from './notification/notification.component';
import {CompanyComponent} from "./company/company.component";
import {CompanyResponsesComponent} from "./company-responses/company-responses.component";

export const routes: Routes = [
  {
    path: 'contacts',
    component: ContactsComponent,
    canActivate: [authGuard],
  },
  {path: 'company', component: CompanyComponent, canActivate: [authGuard]},
  {path: 'leads', component: LeadsComponent, canActivate: [authGuard]},
  {path: 'main-page', component: MainPageComponent, canActivate: [authGuard]},
  {path: 'ads', component: AdsComponent, canActivate: [authGuard]},
  {
    path: 'notifications',
    component: NotificationComponent,
    canActivate: [authGuard],
  },
  {
    path: 'employees',
    component: EmployeesComponent,
    canActivate: [authGuard],
  },
  {path: 'login', component: LoginComponent},
  {path: 'sign-up', component: RegisterComponent},
  {path: 'responses', component: CompanyResponsesComponent},
  {path: '**', redirectTo: 'main-page'}
];
