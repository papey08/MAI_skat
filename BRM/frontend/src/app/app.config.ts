import {ApplicationConfig, LOCALE_ID} from '@angular/core';
import {provideRouter} from '@angular/router';

import {routes} from './app.routes';
import {provideAnimations} from '@angular/platform-browser/animations';
import {provideHttpClient, withInterceptors} from '@angular/common/http';
import {authInterceptor} from './DAL/auth.interceptor';
import {MatPaginatorIntl} from '@angular/material/paginator';
import {MatPaginatorIntlRu} from '../MatPaginatorIntlRu';
import {provideAnimationsAsync} from '@angular/platform-browser/animations/async';
import localeRu from '@angular/common/locales/ru';
import {registerLocaleData} from '@angular/common';

registerLocaleData(localeRu);

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes),
    provideAnimations(),
    {provide: MatPaginatorIntl, useClass: MatPaginatorIntlRu},
    {provide: LOCALE_ID, useValue: 'ru'},
    provideHttpClient(withInterceptors([authInterceptor])),
    provideAnimationsAsync(),
  ],
};
