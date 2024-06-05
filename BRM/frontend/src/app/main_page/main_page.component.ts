import {Component, inject} from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {DalService} from "../DAL/core/dal.service";
import {MainPageResponse} from "../DAL/core/model/MainPageResponse";
import {AuthService} from "../services/auth.service";
import {MatGridListModule} from "@angular/material/grid-list";

@Component({
  selector: 'app-main-page',
  standalone: true,
  imports: [MatCardModule, MatGridListModule],
  templateUrl: './main_page.component.html',
  styleUrl: './main_page.component.scss'
})
export class MainPageComponent {

  dalService = inject(DalService);
  authService = inject(AuthService);
  mainPage?: MainPageResponse

  constructor() {
    this.dalService.getCompanyMainPage(+this.authService.currentUserDataSig()?.
      ["company-id"]!).subscribe(
      // value => this.mainPage = value
      value => {
        if (value && value.data.company_relative_rating !== undefined) {
          value.data.company_relative_rating = +(value.data.company_relative_rating * 100).toFixed(2);
          this.mainPage = value;
        }
      }
    )
  }
}
